package rander

import (
	"rpg/game"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Rander(world game.World, frame int, screen *ebiten.Image) {
	var myUnit *game.Unit
	var relativeDifferenceX, relativeDifferenceY float64

	screenX, screenY := screen.Size()
	unitList := []*game.Unit{}

	for _, unit := range world.Units {
		unitList = append(unitList, unit)
		if unit.ID == world.MyID {
			myUnit = unit
			continue
		}
	}

	sort.Slice(unitList, func(i, j int) bool {
		return unitList[i].Y < unitList[j].Y
	})

	relativeDifferenceX = float64(screenX)/2 - myUnit.X
	relativeDifferenceY = float64(screenY)/2 - myUnit.Y

	for _, unit := range unitList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(relativeDifferenceX+unit.X, relativeDifferenceY+unit.Y)
		spriteIndex := (frame/12 + unit.Frame) % 4
		img, _, _ := ebitenutil.NewImageFromFile("frames/" +
			unit.SpriteName + "_" + unit.Action + "_anim_f" + strconv.Itoa(spriteIndex) + ".png")
		screen.DrawImage(img, op)
		// ebiten.SetWindowPosition(int(unit.X), int(unit.Y))

	}
}
