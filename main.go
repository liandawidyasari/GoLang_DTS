package main

import (
	"net/http"

	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-MVC-Golang-Concept/app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/statys", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}

func getSomething(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"body1": "Get Something Success",
	})

	return
}
