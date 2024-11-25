package task

import "github.com/gin-gonic/gin"

func CreateTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST Task",
	})
}

func GetTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET Task",
	})
}

func UpdateTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "UPDATE Task",
	})
}

func DeleteTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE Task",
	})
}
