package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		// users, err := model.GetAll()
		//if err != nil {
		//	context.JSON(http.StatusInternalServerError, gin.H{
		//		"detail": err,
		//	})
		//} else {
		//	context.JSON(http.StatusOK, gin.H{
		//		"Status": "OK",
		//		"data": users,
		//		"total": len(users),
		//	})
		//}
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
