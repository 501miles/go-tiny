package main

//import logger "github.com/sirupsen/logrus"
import "github.com/gin-gonic/gin"


func main() {
	//str := "9990001"
	//logger.Info(str[:3])
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
