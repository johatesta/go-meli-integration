package tp5_consumo_apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumoApi() {
	r := gin.Default()
	r.GET("https://api.mercadolibre.com/sites/MLA", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run(":3000")
}
