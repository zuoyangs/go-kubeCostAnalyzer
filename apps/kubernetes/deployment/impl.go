package deployment

import (
	"context"
	"fmt"

	"github.com/yourdomain/kubeCostAnalyzer/apps/kubernetes/util"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListAll() ([]v1.Deployment, error) {
	// 获取KubeConfig配置
	config, err := util.GetKubeConfig()
	if err != nil {
		return nil, fmt.Errorf("获取kubeconfig失败: %w", err)
	}

	// 根据KubeConfig配置创建Kubernetes客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("创建Kubernetes客户端失败: %w", err)
	}

	// 获取所有命名空间的Deployment资源客户端
	deploymentsClient := clientset.AppsV1().Deployments(metav1.NamespaceAll)
	// 列出所有Deployment资源
	deploymentList, err := deploymentsClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("列出deployments失败: %w", err)
	}

	// 返回所有Deployment资源项
	return deploymentList.Items, nil
}

func ExportToExcel(timestamp string) (string, error) {
	// 获取deployment资源并填充ClusterResources结构
	// ...

	// 使用WriteClusterResourcesToExcel函数将ClusterResources写入Excel
	// ...

	// 返回生成的Excel文件路径
	return filePath, nil
}
