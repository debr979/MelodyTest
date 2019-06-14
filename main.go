package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	m := melody.New()
	r.LoadHTMLGlob("./html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{"Title":"Melody"})
	})
	r.GET("/ws", func(c *gin.Context) {
		if err := m.HandleRequest(c.Writer, c.Request); err != nil {
			log.Print(err.Error())
		}
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if err := m.Broadcast(msg); err != nil {
			log.Print(err.Error())
		}
	})

	r.Run(":5000")
}
