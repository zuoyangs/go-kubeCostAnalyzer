package utils

import (
	"fmt"
	"log"
	"os"
)

// ensureDirExists 检查目录是否存在，如果不存在则创建它
func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 尝试创建目录
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建目录 %s 失败: %w", dir, err)
		}
		log.Println("目录创建成功:", dir)
	}
	return nil
}
