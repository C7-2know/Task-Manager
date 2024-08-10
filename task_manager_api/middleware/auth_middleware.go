package middleware

import (
	"fmt"
	"strings"
	"task_manager/data"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserMiddleware(c *gin.Context) {
	headers,_ := c.Cookie("Authorization")
	if headers == "" {
		c.JSON(401, gin.H{"error": "Unauthorized user"})
		c.Abort()
		return
	}

	authParts := strings.Split(headers, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.JSON(401, gin.H{"error": "Invalid Authorization header"})
		c.Abort()
	}
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(data.JwtSecret), nil
	})
	if err!=nil{
		c.JSON(401,gin.H{"error":err.Error()})
		c.Abort()
		return
	}
	if claims,ok:= token.Claims.(jwt.MapClaims); ok && token.Valid{
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.JSON(401,gin.H{"error":"Token expired"})
			c.Abort()
			return
		}
	}

	c.Next()

}

func AdminMiddleWare(c *gin.Context) {
	
	headers,_ := c.Cookie("Authorization")
	if headers == "" {
		c.JSON(401, gin.H{"error": "Unauthorized user"})
		c.Abort()
		return
	}

	authParts := strings.Split(headers, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.JSON(401, gin.H{"error": "Invalid Authorization header"})
		c.Abort()
	}
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(data.JwtSecret), nil
	})
	if err!=nil{
		c.JSON(401,gin.H{"error":err.Error()})
		c.Abort()
		return
	}
	if claims,ok:= token.Claims.(jwt.MapClaims); ok && token.Valid{
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.JSON(401,gin.H{"error":"Token expired"})
			c.Abort()
			return
		}
		if claims["role"]!="admin"{
			c.JSON(401,gin.H{"error":"Unauthorized user"})
			c.Abort()
			return
		}
	}

	c.Next()
}
