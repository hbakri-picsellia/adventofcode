package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"math"
	"strings"
)

type Valve struct {
	Name     string
	FlowRate float64
	Valves   []*Valve
	Open     bool
}

func MakeValve(s string) (valve Valve) {
	valve.Name = utils.GetStringBetween(s, "Valve", "has")
	valve.FlowRate = float64(utils.ParseStringToInt(utils.GetStringBetween(s, "flow rate=", ";")))
	return
}

func (valve *Valve) addValves(valve2 []*Valve) {
	valve.Valves = append(valve.Valves, valve2...)
}

type Valves struct {
	List[Valve]
}

func MakeValves(s string) Valves {
	rowValves := strings.Split(s, "\n")
	var valves List[Valve] = operators.Map(rowValves, MakeValve)
	for index := range valves {
		valves[index].addValves(operators.Map(strings.Split(utils.GetStringBetween(rowValves[index], "valve", "\n"), ","), func(valveName string) *Valve {
			return valves.Find(func(valve Valve) bool {
				return valve.Name == strings.Trim(valveName, "s ")
			})
		}))
	}
	return Valves{valves}
}

func (valves *Valves) GetAdjacencyMatrix() Matrix[float64] {
	adjacencyMatrix := MakeMatrix[float64](len(valves.List), len(valves.List), math.Inf(1))
	for x := range valves.List {
		for _, linkedValve := range valves.List[x].Valves {
			y := valves.FindIndex(func(valve Valve) bool {
				return valve.Name == linkedValve.Name
			})
			adjacencyMatrix[x][y] = 1
		}
	}
	return adjacencyMatrix
}
