package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "授权头缺失"})
			return
		}

		// 提取 token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析 token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 确保 token 签名方法符合预期
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// 返回用于验证签名的密钥
			return []byte("你的JWT密钥"), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			return
		}

		// 从 token 的 claims 中获取 userID
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["userID"].(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无法解析用户ID"})
				return
			} else {
				fmt.Printf("user id is : ", userID)
			}

			// ... 现在你可以使用 userID 来获取用户信息，或者进一步验证用户的角色等 ...
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无法解析用户ID"})
			return
		}

		c.Next()
	}
}
