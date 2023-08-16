/*
 * @Author: p_hanxichen
 * @Date: 2023-08-16 18:02:48
 * @LastEditors: p_hanxichen
 * @FilePath: /gin/main.go
 * @Description:
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"message": name,
		})
	})
	r.GET("/view/*.html", func(ctx *gin.Context) {
		page := ctx.Param(".html")
		ctx.JSON(200, gin.H{
			"message": page,
		})
	})
	r.GET("/order", func(ctx *gin.Context) {
		oid := ctx.Query("id")
		ctx.JSON(200, gin.H{
			"message": oid,
		})
	})
	r.POST("/post", func(ctx *gin.Context) {
		param := ctx.Request.Body
		ctx.JSON(200, gin.H{
			"message": param,
		})
	})
	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务

}
