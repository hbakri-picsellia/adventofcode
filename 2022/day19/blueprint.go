package main

import (
	"adventofcode/utils"
	"reflect"
	"regexp"
	"strings"
)

type Blueprint struct {
	Id                int
	OreRobotCost      Inventory
	ClayRobotCost     Inventory
	ObsidianRobotCost Inventory
	GeodeRobotCost    Inventory
}

func MakeBlueprint(rawBlueprint string) (blueprint Blueprint) {
	regex := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	submatches := regex.FindStringSubmatch(rawBlueprint)[1:]
	blueprint.Id = utils.ParseInt(submatches[0])
	blueprint.OreRobotCost = Inventory{Ore: utils.ParseInt(submatches[1])}
	blueprint.ClayRobotCost = Inventory{Ore: utils.ParseInt(submatches[2])}
	blueprint.ObsidianRobotCost = Inventory{Ore: utils.ParseInt(submatches[3]), Clay: utils.ParseInt(submatches[4])}
	blueprint.GeodeRobotCost = Inventory{Ore: utils.ParseInt(submatches[5]), Obsidian: utils.ParseInt(submatches[6])}
	return blueprint
}

func (blueprint Blueprint) MaxRobotCosts() (result Inventory) {
	values := reflect.ValueOf(blueprint)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		if inventory, ok := values.Field(i).Interface().(Inventory); ok &&
			strings.HasSuffix(typesOf.Field(i).Name, "RobotCost") {
			result = result.Max(inventory)
		}
	}
	return result
}
