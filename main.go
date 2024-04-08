package main

import (
	"fmt"
	"log"
	"time"
)

const (
	outputDir   = "./output"
	fileNameFmt = "./%s/华为云CCE集群_资源成本精细化分账报表_%s.xlsx"
)

func main() {

	// 调用ensureDirExists函数来确保目录存在
	if err := utils.ensureDirExists(outputDir); err != nil {
		log.Fatalf("目录处理失败: %v", err)
	}

	// 获取当前时间
	currentTime := time.Now()

	// 格式化当前时间为指定的字符串格式
	timestamp := currentTime.Format("2006-01-02_15-04-05")

	// 使用格式化后的时间戳来构造文件名
	dst := fmt.Sprintf("./output/华为云CCE集群_资源成本精细化分账报表_%s.xlsx", timestamp)

	// 打印构造好的文件名
	fmt.Println(dst)
}
