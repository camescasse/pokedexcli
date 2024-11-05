package main

import (
	"errors"
	"math"
)

func calculateCatchRate(baseExperience int) (float64, error) {
	if baseExperience <= 0 {
		return 0, errors.New("base experience must be larger than 0")
	}

	const (
		minCatchRate = 5.0
		maxCatchRate = 90.0
		decayFactor  = 0.003
	)

	// inverse exponential decay formula: y = a * e^(-bx)
	catchRate := maxCatchRate * math.Exp(-decayFactor*float64(baseExperience))

	// normalize result to stay within bounds
	catchRate = math.Min(math.Max(catchRate, minCatchRate), maxCatchRate)

	return catchRate, nil
}
