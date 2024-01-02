package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / 31557600
	case "Mercury":
		return seconds / (0.2408467 / 1 * 31557600)
	case "Venus":
		return seconds / (0.61519726 / 1 * 31557600)
	case "Mars":
		return seconds / (1.8808158 / 1 * 31557600)
	case "Jupiter":
		return seconds / (11.862615 / 1 * 31557600)
	case "Saturn":
		return seconds / (29.447498 / 1 * 31557600)
	case "Uranus":
		return seconds / (84.016846 / 1 * 31557600)
	case "Neptune":
		return seconds / (164.79132 / 1 * 31557600)
	default:
		return -1
	}
}
