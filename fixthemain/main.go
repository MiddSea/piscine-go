package main

import (
	"github.com/01-edu/z01"
)

// Defining the Door struct
type Door struct {
    state int
}

// Defining the constants for door states
const (
    OPEN  = 1
    CLOSE = 0
)

func PrintStr(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}

func CloseDoor(ptrDoor Door) bool {
    PrintStr("Door Closing...")
    ptrDoor.state = CLOSE
    return true
}

func IsDoorOpen(doorDoor) bool {
    PrintStr("is the Door opened ?")
    return door.state == OPEN
}

func IsDoorClose(ptrDoor Door) bool {
    PrintStr("is the Door closed ?")
    return ptrDoor.state == CLOSE
}

func OpenDoor(ptrDoorDoor) bool {
    PrintStr("Door Opening...")
    ptrDoor.state = OPEN
    return true
}

func main() {
    door := &Door{}

    OpenDoor(door)
    if IsDoorClose(door) {
        OpenDoor(door)
    }
    if IsDoorOpen(door) {
        CloseDoor(door)
    }
    if door.state == OPEN {
        CloseDoor(door)
    }
}