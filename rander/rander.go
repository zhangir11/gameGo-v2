package rander

import (
	"rpg/game"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rander struct {
	World               *game.World
	Frame               int
	Screen              *ebiten.Image
	RelativeDifferenceX float64
	RelativeDifferenceY float64
}

func RanderBuild(world *game.World, frame int, screen *ebiten.Image) Rander {
	var r Rander
	r.World = world
	r.Frame = frame
	r.Screen = screen
	return r
}

func (rander *Rander) RanderUnits() {
	var myUnit *game.Unit

	screenX, screenY := rander.Screen.Size()
	unitList := []*game.Unit{}

	for _, unit := range rander.World.Units {
		unitList = append(unitList, unit)
		if unit.ID == rander.World.MyID {
			myUnit = unit
			continue
		}
	}

	sort.Slice(unitList, func(i, j int) bool {
		return unitList[i].Y < unitList[j].Y
	})

	rander.RelativeDifferenceX = float64(screenX)/2 - myUnit.X
	rander.RelativeDifferenceY = float64(screenY)/2 - myUnit.Y

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(rander.RelativeDifferenceX, rander.RelativeDifferenceY)
	rander.Screen.DrawImage(rander.World.Maps, op)

	for _, unit := range unitList {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(rander.RelativeDifferenceX+unit.X, rander.RelativeDifferenceY+unit.Y)
		spriteIndex := (rander.Frame/7 + unit.Frame) % 4
		img, _, _ := ebitenutil.NewImageFromFile("frames/" +
			unit.SpriteName + "_" + unit.Action + "_anim_f" + strconv.Itoa(spriteIndex) + ".png")
		// if unit.Action == game.ActionRun {
		// 	if unit.HorizontalDirection == game.DirectionLeft {
		// 		op.GeoM.Invert()
		// 	}
		// }
		rander.Screen.DrawImage(img, op)
		// ebiten.SetWindowPosition(int(unit.X), int(unit.Y))

	}
}

func (rander *Rander) RanderMaps() {
	var x, y float64
	maps := game.LoadMap()
	mapsImage := ebiten.NewImage(len(maps[0])*16, len(maps)*16)
	for _, line := range maps {
		for _, sprite := range line {

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)

			img, _, _ := ebitenutil.NewImageFromFile("frames/" + sprite + ".png")
			mapsImage.DrawImage(img, op)
			x += 16
		}
		y += 16
		x = 0
	}
	rander.World.Maps = mapsImage
}
