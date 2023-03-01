package main

import (
	config "first-app/configuration"
	"first-app/controllers"
	"first-app/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func init() {
	config.ConnectToDB()
	config.LoadEnv()
}

func main() {
	//route
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     false,
		ValidateHeaders: false,
	}))
	r.GET("/posts", controllers.PostsRead)
	r.GET("/posts/:id", controllers.PostShow)
	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	//auth
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.CheckAuth, controllers.Validate)
	r.Run()

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.

	// handler := cors.Default().Handler(mux)
	// http.ListenAndServe(":8080", handler)
}
