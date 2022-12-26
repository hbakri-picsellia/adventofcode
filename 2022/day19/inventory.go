package main

import "adventofcode/mathInt"

type Inventory struct {
	Ore      int
	Clay     int
	Obsidian int
	Geode    int
}

func (inventory Inventory) Add(inventory2 Inventory) Inventory {
	inventory3 := inventory
	inventory3.Ore += inventory2.Ore
	inventory3.Clay += inventory2.Clay
	inventory3.Obsidian += inventory2.Obsidian
	inventory3.Geode += inventory2.Geode
	return inventory3
}

func (inventory Inventory) Subtract(inventory2 Inventory) Inventory {
	inventory3 := inventory
	inventory3.Ore -= inventory2.Ore
	inventory3.Clay -= inventory2.Clay
	inventory3.Obsidian -= inventory2.Obsidian
	inventory3.Geode -= inventory2.Geode
	return inventory3
}

func (inventory Inventory) Multiply(i int) Inventory {
	inventory3 := inventory
	inventory3.Ore *= i
	inventory3.Clay *= i
	inventory3.Obsidian *= i
	inventory3.Geode *= i
	return inventory3
}

func (inventory Inventory) LessThan(inventory2 Inventory) bool {
	return inventory.Ore <= inventory2.Ore &&
		inventory.Clay <= inventory2.Clay &&
		inventory.Obsidian <= inventory2.Obsidian &&
		inventory.Geode <= inventory2.Geode
}

func (inventory Inventory) Max(inventory2 Inventory) (inventory3 Inventory) {
	inventory3.Ore = mathInt.Max(inventory.Ore, inventory2.Ore)
	inventory3.Clay = mathInt.Max(inventory.Clay, inventory2.Clay)
	inventory3.Obsidian = mathInt.Max(inventory.Obsidian, inventory2.Obsidian)
	inventory3.Geode = mathInt.Max(inventory.Geode, inventory2.Geode)
	return inventory3
}
