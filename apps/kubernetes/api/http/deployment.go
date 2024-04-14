package http

import (
	"encoding/json"
	"net/http"

	"github.com/yourdomain/kubeCostAnalyzer/apps/kubernetes/deployment"
	"github.com/yourdomain/kubeCostAnalyzer/apps/kubernetes/util"
)

func HandleDeploymentList(w http.ResponseWriter, r *http.Request) {
	// 从 deployment 服务中获取deployment列表
	deployments, err := deployment.ListAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将 deployment 序列化为JSON并写入响应
	jsonData, err := json.Marshal(deployments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandleDeploymentExport(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中提取时间戳或生成默认时间戳
	timestamp := r.URL.Query().Get("timestamp")
	if timestamp == "" {
		timestamp = util.GetCurrentTimestamp()
	}

	// 将 deployment 资源导出到Excel文件
	filePath, err := deployment.ExportToExcel(timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 提供生成的Excel文件
	http.ServeFile(w, r, filePath)
}
