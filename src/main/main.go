package main

import (
	"github.com/EwertonMendes/billog-backend/src/main/resources"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := ":5000"

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		resources.Routes(v1)
	}

	router.Run(PORT)
}
