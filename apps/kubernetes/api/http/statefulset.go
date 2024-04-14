package http

import (
	"encoding/json"
	"net/http"

	"github.com/yourdomain/kubeCostAnalyzer/apps/kubernetes/statefulset"
	"github.com/yourdomain/kubeCostAnalyzer/apps/kubernetes/util"
)

// HandleStatefulSetList 处理获取状态型集合列表的请求。
// 参数:
// - w http.ResponseWriter: 用于向客户端发送响应的接口。
// - r *http.Request: 表示客户端请求的结构体。
// 该函数不返回任何值。
func HandleStatefulSetList(w http.ResponseWriter, r *http.Request) {
	// 从状态型集合服务获取状态型集合列表
	statefulSets, err := statefulset.ListAll()
	if err != nil {
		// 若获取列表时出现错误，返回500内部服务器错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将状态型集合列表序列化为JSON，并写入响应中
	jsonData, err := json.Marshal(statefulSets)
	if err != nil {
		// 若序列化时出现错误，返回500内部服务器错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置响应头部为JSON类型
	w.Header().Set("Content-Type", "application/json")
	// 向客户端发送JSON数据
	w.Write(jsonData)
}

// HandleStatefulSetExport 处理 statefulset 的导出请求
// 将指定时间戳的 statefulset 资源导出为Excel文件，并返回给客户端。
//
// 参数:
// w http.ResponseWriter - 用于向客户端发送响应的接口
// r *http.Request - 表示客户端请求的结构体
//
// 无返回值，但会通过http.ResponseWriter向客户端发送导出的Excel文件。
func HandleStatefulSetExport(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中提取时间戳或生成默认时间戳
	timestamp := r.URL.Query().Get("timestamp")
	if timestamp == "" {
		timestamp = util.GetCurrentTimestamp()
	}

	// 将 statefulset 资源导出到Excel文件
	filePath, err := statefulset.ExportToExcel(timestamp)
	if err != nil {
		// 如果导出过程中出现错误，返回500 Internal Server Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 服务生成的Excel文件
	http.ServeFile(w, r, filePath)
}
