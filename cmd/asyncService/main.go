package main

import (
	"awesomeProject/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PK struct {
	PK int `json:"id"`
}

func main() {
	r := gin.Default()
	r.POST("/set_status", func(c *gin.Context) {
		var request PK
		if err := c.BindJSON(&request); err != nil {
			// DO SOMETHING WITH THE ERROR
		}
		// Отправка PUT-запроса к основному серверу
		url := "http://127.0.0.1:8000/applications/putQuantityOfVotes/" // Замените на ваш реальный URL
		go api.SendStatus(request.PK, url)
		c.JSON(http.StatusOK, gin.H{"message": "Status update initiated"})
	})

	r.Run()
}
