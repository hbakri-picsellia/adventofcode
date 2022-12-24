package main

import (
	"adventofcode/mathInt"
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type Input struct {
	blueprint Blueprint
	minutes   int
	robots    Inventory
	inventory Inventory
}

type Output struct {
	result int
}

type Func func(blueprint Blueprint, minutes int, robots Inventory, inventory Inventory) int

var cache = make(map[Input]Output)

func cached(fn Func) Func {
	return func(blueprint Blueprint, minutes int, robots Inventory, inventory Inventory) int {
		input := Input{blueprint: blueprint, minutes: minutes, robots: robots, inventory: inventory}
		if output, found := cache[input]; found {
			return output.result
		}

		result := fn(blueprint, minutes, robots, inventory)
		cache[input] = Output{result: result}
		return result
	}
}

func NotEnoughMinerals(blueprint Blueprint, minutes int, robots Inventory, inventory Inventory) (result int) {
	if minutes == 0 {
		return inventory.Geode
	}
	waitMinutes := 0

	// if we create no more robot
	result = inventory.Geode + robots.Geode*minutes

	// if we create an ore robot
	if robots.Ore > 0 {
		waitMinutes = int(math.Max(math.Ceil(float64(blueprint.OreRobotCost.Ore-inventory.Ore)/float64(robots.Ore)), 0))
		if minutes-(waitMinutes+1) >= 0 {
			result = mathInt.Max(result, cached(NotEnoughMinerals)(
				blueprint,
				minutes-(waitMinutes+1),
				robots.Add(Inventory{Ore: 1}),
				inventory.Add(robots.Multiply(waitMinutes+1)).Subtract(blueprint.OreRobotCost),
			))
		}
	}

	// if we create a clay robot
	if robots.Ore > 0 {
		waitMinutes = int(math.Max(math.Ceil(float64(blueprint.ClayRobotCost.Ore-inventory.Ore)/float64(robots.Ore)), 0))
		if minutes-(waitMinutes+1) >= 0 {
			result = mathInt.Max(result, cached(NotEnoughMinerals)(
				blueprint,
				minutes-(waitMinutes+1),
				robots.Add(Inventory{Clay: 1}),
				inventory.Add(robots.Multiply(waitMinutes+1)).Subtract(blueprint.ClayRobotCost),
			))
		}
	}

	// if we create an obsidian robot
	if robots.Ore > 0 && robots.Clay > 0 {
		waitMinutesOre := int(math.Max(math.Ceil(float64(blueprint.ObsidianRobotCost.Ore-inventory.Ore)/float64(robots.Ore)), 0))
		waitMinutesClay := int(math.Max(math.Ceil(float64(blueprint.ObsidianRobotCost.Clay-inventory.Clay)/float64(robots.Clay)), 0))
		waitMinutes = mathInt.Max(waitMinutesOre, waitMinutesClay)
		if minutes-(waitMinutes+1) >= 0 {
			result = mathInt.Max(result, cached(NotEnoughMinerals)(
				blueprint,
				minutes-(waitMinutes+1),
				robots.Add(Inventory{Obsidian: 1}),
				inventory.Add(robots.Multiply(waitMinutes+1)).Subtract(blueprint.ObsidianRobotCost),
			))
		}
	}

	// if we create a geode robot
	if robots.Ore > 0 && robots.Obsidian > 0 {
		waitMinutesOre := int(math.Max(math.Ceil(float64(blueprint.GeodeRobotCost.Ore-inventory.Ore)/float64(robots.Ore)), 0))
		waitMinutesObsidian := int(math.Max(math.Ceil(float64(blueprint.GeodeRobotCost.Obsidian-inventory.Obsidian)/float64(robots.Obsidian)), 0))
		waitMinutes = mathInt.Max(waitMinutesOre, waitMinutesObsidian)
		if minutes-(waitMinutes+1) >= 0 {
			result = mathInt.Max(result, cached(NotEnoughMinerals)(
				blueprint,
				minutes-(waitMinutes+1),
				robots.Add(Inventory{Geode: 1}),
				inventory.Add(robots.Multiply(waitMinutes+1)).Subtract(blueprint.GeodeRobotCost),
			))
		}
	}
	return result
}

func step1(input string) int {
	blueprints := List[Blueprint](operators.Map(strings.Split(input, "\n"), MakeBlueprint))
	//blueprint := blueprints[0].ClayRobotCost
	//values := reflect.ValueOf(blueprint)
	//typesOf := values.Type()
	//for i := 0; i < values.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typesOf.Field(i).Name, values.Field(i).Interface())
	//}
	return operators.Sum(operators.Map(blueprints, func(blueprint Blueprint) int {
		result := cached(NotEnoughMinerals)(blueprint, 24, Inventory{Ore: 1}, Inventory{})
		fmt.Println(blueprint.Id, result)
		return blueprint.Id * result
	}))
}

func step2(input string) int {
	blueprints := List[Blueprint](operators.Map(strings.Split(input, "\n"), MakeBlueprint))
	blueprints = blueprints[:mathInt.Min(len(blueprints), 3)]
	return operators.Multiply(operators.Map(blueprints, func(blueprint Blueprint) int {
		result := cached(NotEnoughMinerals)(blueprint, 32, Inventory{Ore: 1}, Inventory{})
		fmt.Println(blueprint.Id, result)
		return result
	}))
}

func main() {
	const title, day = "--- Day 19: Not Enough Minerals ---", "2022/day19/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 33, "example step1")
	utils.AssertEqual(step2(example), 3472, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 1466, "step1")
	utils.AssertEqual(step2(input), 8250, "step2")
}
