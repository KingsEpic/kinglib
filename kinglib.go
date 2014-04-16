package kinglib

import (
	"reflect"
)

const (
	// Need to ensure that if I change the assignments that I update the client as well
	MOVE_WALK       = 0
	MOVE_INSTANT    = 1
	MOVE_PROJECTILE = 2
)

type Packet struct {
	SubType int
	Data    []byte
}

type Archetype struct {
	Name       string
	SimpleName string
	Class      int
	Craftable  bool
	Buildable  bool
}

type ArchetypeInstance struct {
	X, Y       int
	Layer      string
	SimpleName string
	ChunkID    int
}

type UnloadChunk struct {
	ChunkID int
	Layer   string
}

type Vessel struct {
	EntityID int
}

type EntityUpdate struct {
	EntityID       int
	FloatX         float32
	FloatY         float32
	Layer          string
	SimpleName     string
	ChunkID        int
	ContainerID    int
	ContainerIndex int
	Quantity       int
	InventorySize  int // For now, changes in this size not supported
}

type EntityDelete struct {
	EntityID int
}

type Move struct {
	EntityID int
	ChunkID  int
	FloatX   float32
	FloatY   float32
	Layer    string
	Duration float32
	MoveType int
}

type TileInstantiation struct {
	ChunkID int
	X, Y    int
	Layer   string
}

type CraftableRequirement struct {
	ArchetypeName   string
	RequirementName string
	Quantity        int
}

type AnimateAttack struct {
	EntityID int
	TargetID int
	Weapon   string
}

type Player struct {
	Name string
}

// Action commands are issued from client to server:
type ActionMove struct {
	X, Y int
}

type JobStatus struct {
	ID           int
	SocietyID    int
	WorkerID     int
	ParentID     int
	Name         string
	ReadableName string
	Status       string
	Succeeded    bool
	Finished     bool
}

// Create a job:
type Job struct {
	Name       string
	Delegated  bool
	Attributes []byte // JSON marshalled array of attributes and values
}

type InventoryMove struct {
	Index       int
	Destination int
}

func GetSubType(obj interface{}) int {
	t := reflect.TypeOf(obj)

	switch t.String() {
	case "*kinglib.ArchetypeInstance":
		return 1
	case "*kinglib.UnloadChunk":
		return 2
	case "*kinglib.Vessel":
		return 3
	case "*kinglib.EntityUpdate":
		return 4
	case "*kinglib.Move":
		return 5
	case "*kinglib.TileInstantiation":
		return 6
	case "*kinglib.EntityDelete":
		return 7
	case "*kinglib.CraftableRequirement":
		return 8
	case "*kinglib.AnimateAttack":
		return 9
	case "*kinglib.Archetype":
		return 10
	case "*kinglib.Player":
		return 11
	case "*kinglib.ActionMove":
		return 12
	case "*kinglib.Job":
		return 13
	case "kinglib.JobStatus":
		return 14
	case "kinglib.InventoryMove":
		return 15
	default:
		return 0
	}
}
