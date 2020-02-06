package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	ReadEnv()
}

func main() {
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:5000/").Start()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.File("./client/index.html")
	})

	router.Use(static.Serve("/", static.LocalFile("./client", true)))

	fmt.Println("Server started at port", Env.Port)
	router.Run(":5000")
}
