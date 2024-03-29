package game

import (
	"encoding/json"
	"log"
	"rpg/game/player"
	"rpg/game/units"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	MyID        string `json:"-"`
	IsServer    bool   `json:"-"`
	units.Units `json:"units"`
	Maps        *ebiten.Image
}

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type EventConnection struct {
	Player player.Player `json:"player"`
}

type EventMove struct {
	UnitID    string              `json:"unit_id"`
	Direction units.DirectionType `json:"direction"`
}

type EventIdle struct {
	UnitID string `json:"unit_id"`
}

type EventInit struct {
	PlayerID string      `json:"player_id"`
	Units    units.Units `json:"units"`
}

const (
	EventTypeConnection = "connection"
	EventTypeMove       = "move"
	EventTypeIdle       = "idle"
	EventTypeInit       = "init"
)

func (world *World) HandleEvent(event *Event) {
	switch event.Type {
	case EventTypeConnection:
		str, err := json.Marshal(event.Data)
		if err != nil {
			log.Panic(err)
		}
		var ev EventConnection
		err = json.Unmarshal(str, &ev)

		if err != nil {
			log.Panic(err)
		}

		world.Units[ev.Player.ID] = &ev.Player

	case EventTypeMove:
		str, _ := json.Marshal(event.Data)
		var ev EventMove
		json.Unmarshal(str, &ev)

		unit := world.Units[ev.UnitID]

		unit.UpdateAction(units.ActionRun)

		switch ev.Direction {
		case units.DirectionUp:
			unit.UpdateCoordinate(0, -1)
		case units.DirectionDown:
			unit.UpdateCoordinate(0, 1)
		case units.DirectionLeft:
			unit.UpdateCoordinate(-1, 0)
			// unit.HorizontalDirection = ev.Direction
		case units.DirectionRight:
			unit.UpdateCoordinate(1, 0)
			// unit.HorizontalDirection = ev.Direction
		}
		world.Units[ev.UnitID] = unit

	case EventTypeIdle:
		str, _ := json.Marshal(event.Data)
		var ev EventIdle
		json.Unmarshal(str, &ev)

		unit := world.Units[ev.UnitID]
		unit.UpdateAction(units.ActionIdle)

	case EventTypeInit:
		str, _ := json.Marshal(event.Data)
		var ev EventInit
		json.Unmarshal(str, &ev)

		if !world.IsServer {
			world.MyID = ev.PlayerID
			world.Units = ev.Units
		}

	}
}

func (world *World) AddUnit(unit units.IsUnit) *units.IsUnit {
	unit.Create()
	world.Units[unit.GetID()] = unit
	return &unit
}
