package http

import (
	"github.com/gohxk/hxk/app/http/module/demo"
	"github.com/gohxk/hxk/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
