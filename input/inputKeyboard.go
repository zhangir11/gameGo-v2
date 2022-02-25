package input

import (
	"rpg/game"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

//HandleKeyPress handler key pressed
func HandleKeyPress(c *websocket.Conn, world *game.World) {
	// Write your game's logical update.

	var crEvent createEventKey

	crEvent.c = c
	crEvent.world = world

	//handler left
	crEvent.keys = []ebiten.Key{ebiten.KeyA, ebiten.KeyLeft}
	crEvent.dir = game.DirectionLeft
	if crEvent.createEventMove() {
		return
	}

	//handler right
	crEvent.keys = []ebiten.Key{ebiten.KeyD, ebiten.KeyRight}
	crEvent.dir = game.DirectionRight
	if crEvent.createEventMove() {
		return
	}

	//handler Up
	crEvent.keys = []ebiten.Key{ebiten.KeyW, ebiten.KeyUp}
	crEvent.dir = game.DirectionUp
	if crEvent.createEventMove() {
		return
	}

	//handler Down
	crEvent.keys = []ebiten.Key{ebiten.KeyS, ebiten.KeyDown}
	crEvent.dir = game.DirectionDown

	if crEvent.createEventMove() {
		return
	}

	//handle Idle
	if world.Units[world.MyID].Action == game.ActionRun && world.Units[world.MyID].Road == nil {
		c.WriteJSON(game.Event{
			Type: game.EventTypeIdle,
			Data: game.EventIdle{
				UnitID: world.MyID,
			},
		})
	}
}

type createEventKey struct {
	c     *websocket.Conn
	world *game.World
	keys  []ebiten.Key
	dir   game.DirectionType
}

func (cr *createEventKey) createEventMove() bool {
	for _, key := range cr.keys {
		if ebiten.IsKeyPressed(key) {
			cr.world.Units[cr.world.MyID].Road = nil
			cr.c.WriteJSON(
				game.Event{
					Type: game.EventTypeMove,
					Data: game.EventMove{
						UnitID:    cr.world.MyID,
						Direction: cr.dir,
					},
				})

			return true
		}
	}
	return false
}
