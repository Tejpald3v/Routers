package route

import (
	"Routers/routes/gin/database"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	users := database.New()
	r.GET("/", home())
	r.GET("/get", get(users))
	r.POST("/post", post(users))
	r.PUT("/put/:id", put(users))
	r.DELETE("/delete/:id", delete(users))

	r.Run()
}
