package game

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	uuid "github.com/satori/go.uuid"
)

type Unit struct {
	ID                  string        `json:"id"`
	X                   float64       `json:"x"`
	Y                   float64       `json:"y"`
	Road                *Coordinate   `json:"coord"`
	SpriteName          string        `json:"sprite_name"`
	Action              string        `json:"action"`
	Frame               int           `json:"frame"`
	HorizontalDirection DirectionType `json:"horizontal"`
}

type Units map[string]*Unit

type World struct {
	MyID     string `json:"-"`
	IsServer bool   `json:"-"`
	Units    `json:"units"`
	Maps     *ebiten.Image
}

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type EventConnection struct {
	Unit
}

type EventMove struct {
	UnitID    string        `json:"unit_id"`
	Direction DirectionType `json:"direction"`
}

type EventIdle struct {
	UnitID string `json:"unit_id"`
}

type EventInit struct {
	PlayerID string `json:"player_id"`
	Units    Units  `json:"units"`
}

const (
	EventTypeConnection = "connection"
	EventTypeMove       = "move"
	EventTypeIdle       = "idle"
	EventTypeInit       = "init"
)

const (
	ActionRun  = "run"
	ActionIdle = "idle"
)

type DirectionType int

const (
	DirectionUp    DirectionType = 0
	DirectionDown  DirectionType = 1
	DirectionLeft  DirectionType = 2
	DirectionRight DirectionType = 3
)

func (world *World) HandleEvent(event *Event) {
	switch event.Type {
	case EventTypeConnection:
		str, _ := json.Marshal(event.Data)
		var ev EventConnection
		json.Unmarshal(str, &ev)

		world.Units[ev.ID] = &ev.Unit

	case EventTypeMove:
		str, _ := json.Marshal(event.Data)
		var ev EventMove
		json.Unmarshal(str, &ev)

		unit := world.Units[ev.UnitID]
		unit.Action = ActionRun

		switch ev.Direction {
		case DirectionUp:
			unit.Y--
		case DirectionDown:
			unit.Y++
		case DirectionLeft:
			unit.X--
			unit.HorizontalDirection = ev.Direction
		case DirectionRight:
			unit.X++
			unit.HorizontalDirection = ev.Direction
		}

	case EventTypeIdle:
		str, _ := json.Marshal(event.Data)
		var ev EventIdle
		json.Unmarshal(str, &ev)

		unit := world.Units[ev.UnitID]
		unit.Action = ActionIdle

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

func (world *World) AddPlayer() *Unit {
	skins := []string{
		"elf_f", "elf_m", "goblin",
		"ice_zombie", "imp", "knight_f",
		"knight_m", "lizard_f", "lizard_m",
		"necromancer",
	}
	id := uuid.NewV4().String()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	unit := &Unit{
		ID:         id,
		Action:     ActionIdle,
		X:          float64(320 / 2),
		Y:          float64(240 / 2),
		Frame:      rnd.Intn(4),
		SpriteName: skins[rnd.Intn(len(skins))],
	}
	world.Units[unit.ID] = unit

	return unit
}
