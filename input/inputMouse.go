package input

import (
	"fmt"
	"rpg/game"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

//HandleMousePress handler key pressed
func HandleMousePress(c *websocket.Conn, world *game.World) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		fmt.Println(ebiten.CursorPosition())
	}
}
