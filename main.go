package main

import (
	"log"
	"rpg/game"
	"sort"
	"strconv"

	_ "image/png"

	"github.com/gorilla/websocket"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var world game.World
var frame int
var img *ebiten.Image
var c *websocket.Conn

func init() {
	world = game.World{
		IsServer: false,
		Units:    game.Units{},
	}
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
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

		return nil
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

		return nil
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

		return nil
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

		return nil
	}

	if world.Units[world.MyID].Action == game.ActionRun {
		c.WriteJSON(game.Event{
			Type: game.EventTypeIdle,
			Data: game.EventIdle{
				UnitID: world.MyID,
			},
		})
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	frame++
	// img, _, _ := ebitenutil.NewImageFromFile("frames/big_demon_idle_anim_f0.png")

	// screen.DrawImage(img, nil)

	unitList := []*game.Unit{}
	for _, unit := range world.Units {
		unitList = append(unitList, unit)
	}
	sort.Slice(unitList, func(i, j int) bool {
		return unitList[i].Y < unitList[j].Y
	})
	for _, unit := range unitList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(unit.X, unit.Y)
		spriteIndex := (frame/12 + unit.Frame) % 4
		img, _, _ = ebitenutil.NewImageFromFile("frames/" +
			unit.SpriteName + "_" + unit.Action + "_anim_f" + strconv.Itoa(spriteIndex) + ".png")
		screen.DrawImage(img, op)

	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	var err error
	c, _, err = websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		log.Println(err)
		return
	}
	go func(c *websocket.Conn) {
		defer c.Close()

		for {
			var event game.Event
			c.ReadJSON(&event)
			world.HandleEvent(&event)
		}
	}(c)

	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
