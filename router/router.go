package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func Router() {

	router := gin.Default()

	router.GET("/", home)

	fmt.Println("http://localhost" + PORT)
	router.Run(PORT)
}

func home(context *gin.Context) {
	context.String(http.StatusOK, "message748243")
}
