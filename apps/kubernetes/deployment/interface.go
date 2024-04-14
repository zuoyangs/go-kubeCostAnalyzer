package deployment

import (
	v1 "k8s.io/api/apps/v1"
)

// ListAll 从Kubernetes集群中检索所有deployments。
func ListAll() ([]v1.Deployment, error)

// ExportToExcel 将deployment资源导出到Excel文件。
// 提供的时间戳用于文件名。
func ExportToExcel(timestamp string) (string, error)
