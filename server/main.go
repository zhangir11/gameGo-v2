package main

import (
	"fmt"
	"rpg/game"
	"rpg/game/units"

	"github.com/gin-gonic/gin"
)

func main() {
	world := &game.World{
		IsServer: true,
		Units:    units.Units{},
	}
	hub := newHub()
	go hub.run()
	r := gin.New()
	r.GET("/ws", func(hub *Hub, world *game.World) gin.HandlerFunc {
		return gin.HandlerFunc(func(c *gin.Context) {
			serveWs(hub, world, c.Writer, c.Request)
		})
	}(hub, world))
	r.Run(":8080")
	fmt.Println(r)
}
