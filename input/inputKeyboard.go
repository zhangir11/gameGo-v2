package input

import (
	"rpg/game"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

//HandleKeyPress handler key pressed
func HandleKeyPress(c *websocket.Conn, world *game.World) {
	// Write your game's logical update.
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.WriteJSON(
			game.Event{
				Type: game.EventTypeMove,
				Data: game.EventMove{
					UnitID:    world.MyID,
					Direction: game.DirectionLeft,
				},
			})

		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.WriteJSON(
			game.Event{
				Type: game.EventTypeMove,
				Data: game.EventMove{
					UnitID:    world.MyID,
					Direction: game.DirectionRight,
				},
			})

		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		c.WriteJSON(
			game.Event{
				Type: game.EventTypeMove,
				Data: game.EventMove{
					UnitID:    world.MyID,
					Direction: game.DirectionUp,
				},
			})

		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		c.WriteJSON(
			game.Event{
				Type: game.EventTypeMove,
				Data: game.EventMove{
					UnitID:    world.MyID,
					Direction: game.DirectionDown,
				},
			})

		return
	}

	if world.Units[world.MyID].Action == game.ActionRun {
		c.WriteJSON(game.Event{
			Type: game.EventTypeIdle,
			Data: game.EventIdle{
				UnitID: world.MyID,
			},
		})
	}
}
