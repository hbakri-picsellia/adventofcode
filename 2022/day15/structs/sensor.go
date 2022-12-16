package structs

import (
	"adventofcode/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"math"
	"strings"
)

type Sensor struct {
	Position                  models.Position
	ClosestBeaconPosition     models.Position
	DistanceWithClosestBeacon float64
}

func getStringBetween(s string, sep1 string, sep2 string) string {
	return strings.TrimSpace(strings.SplitN(strings.SplitAfterN(s, sep1, 2)[1], sep2, 2)[0])
}

func (sensor *Sensor) Decode(s string) {
	positionJ := utils.ParseStringToInt(getStringBetween(s, "Sensor at x=", ","))
	positionI := utils.ParseStringToInt(getStringBetween(strings.Split(s, ",")[1], "y=", ":"))
	sensor.Position = models.Position{I: positionI, J: positionJ}

	closestBeaconPositionJ := utils.ParseStringToInt(getStringBetween(strings.Split(s, ":")[1], "closest beacon is at x=", ","))
	closestBeaconPositionI := utils.ParseStringToInt(getStringBetween(strings.Split(s, ",")[2], "y=", "\n"))
	sensor.ClosestBeaconPosition = models.Position{I: closestBeaconPositionI, J: closestBeaconPositionJ}

	sensor.DistanceWithClosestBeacon = sensor.Position.Distance(sensor.ClosestBeaconPosition)
}

type SensorList []Sensor

func (sensorList SensorList) Decode(s string) ([]Sensor, int, int) {
	minJ, maxJ := math.MaxInt, math.MinInt
	return operators.Map(strings.Split(s, "\n"), func(sensorInput string) (sensor Sensor) {
		sensor.Decode(sensorInput)
		minJ = int(math.Min(float64(sensor.Position.J)-sensor.DistanceWithClosestBeacon, float64(minJ)))
		maxJ = int(math.Max(float64(sensor.Position.J)+sensor.DistanceWithClosestBeacon, float64(maxJ)))
		minJ = int(math.Min(float64(sensor.ClosestBeaconPosition.J)-sensor.DistanceWithClosestBeacon, float64(minJ)))
		maxJ = int(math.Max(float64(sensor.ClosestBeaconPosition.J)+sensor.DistanceWithClosestBeacon, float64(maxJ)))
		return sensor
	}), minJ, maxJ
}
