package structs

import (
	"adventofcode/operators"
	"adventofcode/structs"
	"adventofcode/utils"
	"math"
	"strings"
)

type Sensor struct {
	Position                  structs.Position
	ClosestBeaconPosition     structs.Position
	DistanceWithClosestBeacon float64
}

func getStringBetween(s string, sep1 string, sep2 string) string {
	return strings.TrimSpace(strings.SplitN(strings.SplitAfterN(s, sep1, 2)[1], sep2, 2)[0])
}

func (sensor *Sensor) Decode(s string) {
	positionY := utils.ParseInt(getStringBetween(s, "Sensor at x=", ","))
	positionX := utils.ParseInt(getStringBetween(strings.Split(s, ",")[1], "y=", ":"))
	sensor.Position = structs.Position{X: positionX, Y: positionY}

	closestBeaconPositionY := utils.ParseInt(getStringBetween(strings.Split(s, ":")[1], "closest beacon is at x=", ","))
	closestBeaconPositionX := utils.ParseInt(getStringBetween(strings.Split(s, ",")[2], "y=", "\n"))
	sensor.ClosestBeaconPosition = structs.Position{X: closestBeaconPositionX, Y: closestBeaconPositionY}

	sensor.DistanceWithClosestBeacon = sensor.Position.Distance(sensor.ClosestBeaconPosition)
}

func (sensor *Sensor) GetExternalBorder() (points []structs.Position) {
	borderDistance := int(sensor.DistanceWithClosestBeacon) + 1
	var dY, X, Y int
	for dX := 0; dX <= borderDistance; dX++ {
		dY = borderDistance - dX
		for _, direction := range [4][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
			X = sensor.Position.X + dX*direction[0]
			Y = sensor.Position.Y + dY*direction[1]
			points = append(points, structs.Position{X: X, Y: Y})
		}
	}
	return points
}

type SensorList []Sensor

func (sensorList SensorList) Decode(s string) ([]Sensor, int, int) {
	minY, maxY := math.MaxInt, math.MinInt
	return operators.Map(strings.Split(s, "\n"), func(sensorInput string) (sensor Sensor) {
		sensor.Decode(sensorInput)
		minY = int(math.Min(float64(sensor.Position.Y)-sensor.DistanceWithClosestBeacon, float64(minY)))
		maxY = int(math.Max(float64(sensor.Position.Y)+sensor.DistanceWithClosestBeacon, float64(maxY)))
		minY = int(math.Min(float64(sensor.ClosestBeaconPosition.Y), float64(minY)))
		maxY = int(math.Max(float64(sensor.ClosestBeaconPosition.Y), float64(maxY)))
		return sensor
	}), minY, maxY
}
