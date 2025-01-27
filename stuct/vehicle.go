package main

import (
	"fmt"
)

type Vehicle struct {
	Name     string
	MaxSpeed int
	speed    int
}

type Car struct {
	Vehicle
}

type Truch struct {
	LoadingCapacity int
}

type Drive interface {
	Move() string
}

var ErrOutOfTea = fmt.Errorf("Going over the speed limit")

func (v Vehicle) Move() string {
	if v.speed < v.MaxSpeed {
		fmt.Errorf("%w", ErrOutOfTea)
	}
	return fmt.Sprintf("%s is moving at %d km/h", v.Name, v.speed)
}
func DriveVehicle(d Drive) {
	fmt.Println(d.Move())
}

func main() {
	car1 := Car{
		Vehicle: Vehicle{
			Name:     "Benx",
			MaxSpeed: 220,
			speed:    400,
		},
	}

	DriveVehicle((car1))
}
