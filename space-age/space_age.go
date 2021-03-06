// Package space provides a solution to the space age exercise of the Go track in https://exercism.io
package space

// Planet is a representation of a planet with a name
type Planet string

var earthOrbitalPeriod = 365.25
var secondsInEarthYear = 60 * 60 * 24 * earthOrbitalPeriod
var orbitalPeriodMap = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculate how old someone would be on different planets, given an age in seconds
func Age(seconds float64, planet Planet) float64 {
	return seconds / (orbitalPeriodMap[planet] * secondsInEarthYear)
}
