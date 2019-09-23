package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"hrinfo/config"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("I am before next")
		token := c.Query("token")
		if token != config.Token {

		}
		/*
		   c.Next()后就执行真实的路由函数，路由函数执行完成之后继续执行后续的代码
		*/
		c.Next()
		fmt.Println("I am after next")
	}
}
