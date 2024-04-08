package kubernetes

import (
	"context"
	"fmt"
	"log"

	"github.com/zuoyangs/go-kubeCostAnalyzer/config"
	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getResourceInfo(obj interface{}, clusterName string) ResourceInfo {
	var name string
	var replicas int32
	var totalRequestCpu, totalRequestMemory, totalLimitCpu, totalLimitMemory float64 = 0, 0, 0, 0

	switch v := obj.(type) {
	case *v1.Deployment:
		name = v.Name
		replicas = *v.Spec.Replicas
		for _, container := range v.Spec.Template.Spec.Containers {
			requestCpu := container.Resources.Requests.Cpu().MilliValue()
			requestMemory := container.Resources.Requests.Memory().Value()
			limitCpu := container.Resources.Limits.Cpu().MilliValue()
			limitMemory := container.Resources.Limits.Memory().Value()
			totalRequestCpu += float64(requestCpu) / 1000
			totalRequestMemory += float64(requestMemory) / (1024 * 1024)
			totalLimitCpu += float64(limitCpu) / 1000
			totalLimitMemory += float64(limitMemory) / (1024 * 1024)
		}
	case *v1.StatefulSet:
		name = v.Name
		replicas = *v.Spec.Replicas
		for _, container := range v.Spec.Template.Spec.Containers {
			requestCpu := container.Resources.Requests.Cpu().MilliValue()
			requestMemory := container.Resources.Requests.Memory().Value()
			limitCpu := container.Resources.Limits.Cpu().MilliValue()
			limitMemory := container.Resources.Limits.Memory().Value()
			totalRequestCpu += float64(requestCpu) / 1000
			totalRequestMemory += float64(requestMemory) / (1024 * 1024)
			totalLimitCpu += float64(limitCpu) / 1000
			totalLimitMemory += float64(limitMemory) / (1024 * 1024)
		}
	default:
		return ResourceInfo{}
	}

	return ResourceInfo{
		Name:                name,
		ClusterName:         clusterName,
		Replicas:            replicas,
		TotalRequestsCpu:    totalRequestCpu,
		TotalRequestsMemory: totalRequestMemory,
		TotalLimitsCpu:      totalLimitCpu,
		TotalLimitsMemory:   totalLimitMemory,
	}
}
func Kubernete_Querier(dst, timestamp string) {
	envs := []string{"hwc-sh1-dev-cluster", "hwc-sh1-test-cluster"}
	config.Init()

	clusterResources := make(ClusterResources)

	for _, env := range envs {
		_, err := config.GetSectionsAndLabels(env)
		if err != nil {
			fmt.Println("Error getting labels: ", err)
			continue
		}

		kubeconfig, _ := config.GetKey(env, "kubeconfig_path")

		if kubeconfig == "" {
			log.Println("kubeconfig_path 是空的")
		}

		// 1. 构建配置
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}

		// 2. 创建 clientset 访问kubernetes集群
		coreClient := kubernetes.NewForConfigOrDie(config)

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		envResources := make(EnvResources)

		clusterResources[env] = envResources

		// 3. 获取namespace列表
		namespaceList, err := coreClient.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 4. 遍历namespace列表，并遍历
		for _, namespace := range namespaceList.Items {

			nsResources := make(NamespaceResources)
			envResources[namespace.Name] = nsResources

			// 获取并处理 Deployments
			deploymentsClient := clientset.AppsV1().Deployments(namespace.Name)
			deploymentList, err := deploymentsClient.List(context.Background(), metaV1.ListOptions{})
			if err != nil {
				log.Printf("Error getting deployments in namespace '%s': %v", namespace.Name, err)
			}
			for _, deployment := range deploymentList.Items {
				resourceInfo := getResourceInfo(&deployment, env)
				key := "deployment/" + deployment.Name
				nsResources[key] = resourceInfo
			}

			// 获取并处理 StatefulSets
			statefulSetsClient := clientset.AppsV1().StatefulSets(namespace.Name)
			statefulSetList, err := statefulSetsClient.List(context.Background(), metaV1.ListOptions{})
			if err != nil {
				log.Printf("Error getting statefulsets in namespace '%s': %v", namespace.Name, err)
			}
			for _, statefulSet := range statefulSetList.Items {
				resourceInfo := getResourceInfo(&statefulSet, env)
				key := "statefulSet/" + statefulSet.Name
				nsResources[key] = resourceInfo
			}
		}
	}

	// 遍历并打印 clusterResources
	for env, envResources := range clusterResources {
		fmt.Printf("环境: %s\n", env)
		for ns, nsResources := range envResources {
			fmt.Printf("  命名空间: %s\n", ns)
			for resourceKey, resourceInfo := range nsResources {
				fmt.Printf("    资源: %s\n", resourceKey)
				fmt.Printf("      CCE集群名称: %s\n", resourceInfo.ClusterName)
				fmt.Printf("      名称: %s\n", resourceInfo.Name)
				fmt.Printf("      副本数: %d\n", resourceInfo.Replicas)
				fmt.Printf("      请求CPU总量: %.2f\n", resourceInfo.TotalRequestsCpu)
				fmt.Printf("      请求内存总量: %.2f MB\n", resourceInfo.TotalRequestsMemory)
				fmt.Printf("      限制CPU总量: %.2f\n", resourceInfo.TotalLimitsCpu)
				fmt.Printf("      限制内存总量: %.2f MB\n", resourceInfo.TotalLimitsMemory)
			}
		}
	}

	// 写入 Excel
	if err := WriteClusterResourcesToExcel(clusterResources, timestamp); err != nil {
		fmt.Println("Error writing to Excel:", err)
		return
	}

	fmt.Println("Excel file has been saved successfully!")
}
