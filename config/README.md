# 配置文件处理库

## 简介

这个Go语言库提供了一个简单的方式来处理配置文件，特别是`ini`格式的配置文件。它使用了`spf13/viper`这个流行的Go配置管理库作为基础，并对其进行了封装，以便于更方便地读取和使用配置文件中的数据。

## 功能特点

- **简化配置文件的加载**：通过`Init`函数，可以轻松加载指定路径下的`ini`配置文件。
- **获取配置节和标签**：`GetSectionsAndLabels`函数允许你获取指定环境下的所有配置节和对应的标签。
- **获取配置项**：通过`GetKey`函数，可以方便地获取指定配置节中的某个配置项的值。

## 使用说明

1. **安装库**：首先，你需要将这个库安装到你的Go项目中。

2. **初始化**：在你的代码中调用`config.Init()`来初始化并加载配置文件。

3. **获取配置**：使用`GetSectionsAndLabels`或`GetKey`函数来获取你需要的配置信息。

## 示例代码

```go
package main

import (
    "fmt"
    "your_project/config" // 替换成你的项目路径
)

func main() {
    // 初始化配置文件
    config.Init()
    
    // 获取指定环境下的配置节和标签
    labels, err := config.GetSectionsAndLabels("development")
    if err != nil {
        // 处理错误
    }
    fmt.Println(labels) // 输出配置节和标签的映射关系
    
    // 获取特定配置项的值
    key, err := config.GetKey("database", "user")
    if err != nil {
        // 处理错误
    }
    fmt.Println("Database User:", key) // 输出配置项的值
}
```

## 注意事项

- 确保你的配置文件路径和名称与库中的默认设置相匹配，或者你可以在`Init`函数中进行自定义设置。
- 如果配置文件不存在或无法读取，库会提供相应的错误信息。