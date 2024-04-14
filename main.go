package main

import (
	"log"
	"net/http"

	"github.com/zuoyangs/go-kubeCostAnalyzer/routers"
	"github.com/zuoyangs/go-kubeCostAnalyzer/utils"
)

const outputDir = "./output"

func main() {

	// 确保output目录存在
	if err := utils.EnsureDirExists(outputDir); err != nil {
		log.Fatalf("无法确保output目录存在: %s。错误详情: %v", outputDir, err)
	}

	/*
		需求后期实现：
			后期实现一个功能，该功能将检查 config.ini 文件的存在性，并确保在 华为云CCE 集群的 /config/kubeconfig/ 目录下存在与 config.ini 相对应的配置文件。
			当前，我们主要关注逻辑框架的搭建和核心功能的实现，具体的检查和维护代码将在后续的开发过程中逐步添加和完善。
			这样做旨在确保我们的系统在配置管理上保持高度的一致性和可靠性，从而为后续的开发和运维工作提供坚实的基础。
	*/

	// 创建路由
	router := routers.SetupRouter()

	// 启动HTTP服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on :8080")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
