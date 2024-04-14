# 路由设置与Token验证中间件

## 项目介绍

该项目是一个简单的Go Web服务，使用Gin Web框架进行路由设置，并包含一个自定义的Token验证中间件。此示例展示了如何定义GET和POST路由，以及如何为所有路由应用中间件来验证请求头中的Token。

## 主要功能

- **路由设置**：使用Gin框架设置了多个路由，包括获取部署信息、获取StatefulSet信息和创建新部署、创建StatefulSet的路由。
- **Token验证中间件**：实现了一个中间件来验证请求头中的Token。如果Token无效或缺失，将返回401未授权状态码。

## 如何使用

1. **安装依赖**：首先，确保你已经安装了Go环境。然后，通过`go get`命令安装Gin框架：


```bash
go get -u github.com/gin-gonic/gin
```
2. **运行项目**：在项目的根目录下运行`main.go`文件（或其他入口文件），启动Web服务。
3. **测试路由**：使用curl、Postman或其他HTTP客户端向定义的路由发送请求，并查看响应。
4. **Token验证**：在发送请求时，确保在请求头中包含有效的`Authorization`字段和对应的Token值。

## 路由列表

- **GET /api/deployments**：获取部署信息
- **GET /api/statefulsets**：获取StatefulSet信息
- **POST /api/deployments**：创建新部署
- **POST /api/statefulsets**：创建StatefulSet

## 注意事项

- 示例中的Token验证是简单的字符串比较，仅用于演示目的。在实际应用中，应使用更安全的Token生成和验证机制。
- 业务逻辑部分在示例中仅用简单的字符串响应代替，实际应用中应根据具体需求实现相应的功能。