package main

import (
	"log"
	"net/http"

	"apiserver/router"
	"github.com/gin-gonic/gin"
)

func main() {

	//创建Gin
	g := gin.New()
	//
	middlewares := []gin.HandlerFunc{}

	router.Load(

		g,

		middlewares...,
	)
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}
