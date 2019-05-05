package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"project-builder/mapper"
	"project-builder/service"
)

func UserAdd(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		name := ctx.PostForm("name")
		ageTemp, cErr := strconv.Atoi(ctx.PostForm("age"))
		if cErr != nil {
			ctx.HTML(http.StatusOK, "user_add.html", gin.H{
				"name": name,
				"age":  ageTemp,
			})
			return
		}

		user := mapper.User{Name: name}
		_, insertErr := service.InsertUser(&user)
		if insertErr != nil {
			ctx.HTML(http.StatusInternalServerError, "user_add.html", gin.H{
				"name": name,
				"age":  ageTemp,
			})
		}

		ctx.HTML(http.StatusOK, "user_add.html", gin.H{
			"name": name,
			"age":  ageTemp,
		})
	} else {
		ctx.HTML(http.StatusOK, "user_add.html", gin.H{
			"title": "用户添加",
		})
	}

}

func QueryUserList(ctx *gin.Context)  {
	service.QueryUserList()
}