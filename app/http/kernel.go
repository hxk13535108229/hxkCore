package http

import (
	"github.com/gohxk/hxk/framework/gin"
)

// NewHttpEngine is command
func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Routes(r)
	return r, nil
}
