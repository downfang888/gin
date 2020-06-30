package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

)


func  AdminGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "admin.html", gin.H{"title": "后台首页"})
}

