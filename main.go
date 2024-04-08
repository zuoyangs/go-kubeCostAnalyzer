package main

import (
	"log"
	"time"

	"github.com/zuoyangs/go-kubeCostAnalyzer/apps/kubernetes"
	"github.com/zuoyangs/go-kubeCostAnalyzer/utils"
)

const (
	outputDir = "./output"
)

func main() {

	// 调用EnsureDirExists函数来确保目录存在
	if err := utils.EnsureDirExists(outputDir); err != nil {
		log.Fatalf("目录处理失败: %v", err)
	}

	// 获取当前时间
	currentTime := time.Now()

	// 格式化当前时间为指定的字符串格式
	timestamp := currentTime.Format("2006-01-02_15-04-05")

	// 拼接输出文件路径
	kubernetes.Kubernete_Querier(outputDir, timestamp)
}
