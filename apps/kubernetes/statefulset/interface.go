package statefulset

import (
	v1 "k8s.io/api/apps/v1"
)

// ListAll 从Kubernetes集群中检索所有的statefulsets
func ListAll() ([]v1.StatefulSet, error)

// ExportToExcel 将statefulset资源导出到Excel文件
// 提供的时间戳将用于文件名。
func ExportToExcel(timestamp string) (string, error)
