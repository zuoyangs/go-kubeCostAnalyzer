package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 应用TokenMiddleware中间件到所有路由
	r.Use(TokenMiddleware())

	// 定义GET和POST路由
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/deployments", GetDeployments)
		apiGroup.GET("/statefulsets", GetStatefulSets)
		apiGroup.POST("/deployments", CreateDeployment)
		apiGroup.POST("/statefulsets", CreateStatefulSet)
	}

	return r
}

// TokenMiddleware 验证token的中间件
func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "your_valid_token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: 401"})
			return
		}
		c.Next()
	}
}

// GetDeployments 处理获取部署信息的GET请求
func GetDeployments(c *gin.Context) {
	// 业务逻辑
	c.String(http.StatusOK, "GetDeployments was called")
}

// GetStatefulSets 处理获取StatefulSet信息的GET请求
func GetStatefulSets(c *gin.Context) {
	// 业务逻辑
	c.String(http.StatusOK, "GetStatefulSets was called")
}

// CreateDeployment 处理创建新部署的POST请求
func CreateDeployment(c *gin.Context) {
	// 业务逻辑
	c.String(http.StatusOK, "CreateDeployment was called")
}

// CreateStatefulSet 处理创建StatefulSet的POST请求
func CreateStatefulSet(c *gin.Context) {
	// 业务逻辑
	c.String(http.StatusOK, "CreateStatefulSet was called")
}
