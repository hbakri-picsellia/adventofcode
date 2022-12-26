package main

type Input struct {
	blueprint Blueprint
	minutes   int
	robots    Inventory
	inventory Inventory
}

type Output struct {
	result int
}

type Function func(blueprint Blueprint, minutes int, robots Inventory, inventory Inventory) int

var cache = make(map[Input]Output)

func cached(function Function) Function {
	return func(blueprint Blueprint, minutes int, robots Inventory, inventory Inventory) int {
		input := Input{blueprint: blueprint, minutes: minutes, robots: robots, inventory: inventory}
		if output, found := cache[input]; found {
			return output.result
		}

		result := function(blueprint, minutes, robots, inventory)
		cache[input] = Output{result: result}
		return result
	}
}
