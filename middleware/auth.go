package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/database"
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey string = "niowdnaWnOKSp142kawo19523kOkbMASj"

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Harap Login terlebih dahulu!",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expired"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Harap Login terlebih dahulu!",
			})
		}

		var user models.User
		database.DB.First(&user, claims["userid"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Harap Login terlebih dahulu!",
			})
		}
		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println(err)
	}

}
