# 目录创建工具函数库

## 简介

这是一个包含 `EnsureDirExists` 函数的Go语言工具库，该函数用于检查指定目录是否存在，如果不存在则尝试创建它。这是一个在文件操作中经常需要用到的实用功能。

## 功能

- **EnsureDirExists**：此函数接收一个目录路径字符串作为参数。如果该路径指向的目录不存在，函数会尝试创建它。如果创建成功，将在日志中记录一条消息。如果目录已存在或创建成功，函数返回 `nil`，否则返回一个描述错误的 `error` 对象。

## 使用说明

1. 首先，你需要在你的Go项目中导入这个包。
2. 调用 `EnsureDirExists` 函数，并传入你想要检查或创建的目录路径。
3. 函数将检查该目录是否存在，如果不存在，将尝试创建它。
4. 根据函数的返回值判断操作是否成功。

## 示例代码

```go
package main

import (
    "fmt"
    "your_project/utils" // 替换成你的项目路径
)

func main() {
    dir := "/path/to/your/directory"
    err := utils.EnsureDirExists(dir)
    if err != nil {
        fmt.Println("目录创建失败:", err)
    } else {
        fmt.Println("目录已存在或已成功创建")
    }
}
```

## 注意事项

- 确保你有足够的权限在指定的路径上创建目录。
- 如果目录创建失败，将返回一个包含错误信息的 `error` 对象。