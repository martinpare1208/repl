package pokeapi

import (
	"math/rand"
)

func GenerateSuccessRate(baseExperience int) (bool) {

	baseRate := float64(baseExperience)

	// Legendaries
	if baseExperience > 300 {
		catchRate := baseRate * 0.90
		catch := GenerateCatch(baseRate)
		if catch > catchRate {
			return true
		}

	}
	//Pseudo Legendaries

	if baseExperience > 200 {
		catchRate := baseRate * 0.80
		catch := GenerateCatch(baseRate)
		if catch > catchRate {
			return true
		}

	//Non legendaries
	} else {
		catchRate := baseRate * 0.40
		catch := GenerateCatch(baseRate)
		if catch > catchRate {
			return true
		}
	}

	return false
}


func GenerateCatch(rate float64) (float64) {
	return rand.Float64()*(rate)
}