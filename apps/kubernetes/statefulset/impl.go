package statefulset

import (
	"context"
	"fmt"

	"github.com/yourdomain/kubeCostAnalyzer/apps/kubernetes/util"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListAll 从Kubernetes集群中获取所有的statefulsets
func ListAll() ([]v1.StatefulSet, error) {
	config, err := util.GetKubeConfig()
	if err != nil {
		return nil, fmt.Errorf("无法获取kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("无法创建Kubernetes客户端: %w", err)
	}

	statefulSetsClient := clientset.AppsV1().StatefulSets(metav1.NamespaceAll)
	statefulSetList, err := statefulSetsClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("无法列出stateful sets: %w", err)
	}

	return statefulSetList.Items, nil
}

// ExportToExcel 将statefulset资源导出到Excel文件
// 提供的时间戳将用于生成文件名
func ExportToExcel(timestamp string) (string, error) {
	// 获取stateful set资源并填充ClusterResources结构

	// 使用WriteClusterResourcesToExcel函数将ClusterResources写入Excel

	// 返回生成的Excel文件路径
	return filePath, nil
}
