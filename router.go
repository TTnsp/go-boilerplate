package main

import (
	"net/http"
	"strings"

	"github.com/ttnsp/go-boilerplate/auth"
	"github.com/ttnsp/go-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoute(r *gin.Engine) {

	r.GET("/foos", isJWTAuthorized, controllers.FindFoos)
	r.GET("/foos/:id", isJWTAuthorized, controllers.FindFoo)
	r.POST("/foos", isJWTAuthorized, controllers.CreateFoo)
	r.PATCH("/foos/:id", isJWTAuthorized, controllers.UpdateFoo)

	r.POST("/login", controllers.Login)
}

func isJWTAuthorized(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization != "" {
		fields := strings.Fields(authorization)
		// fields[0] is Bearer keyword
		// fields[1] is jwt
		if fields[1] != "" {
			valid, payload := auth.IsAuthorized(fields[1])
			if valid {
				c.Set("name", payload.Name)
				return
			}
		}
	}

	c.AbortWithStatus(http.StatusForbidden)
}
