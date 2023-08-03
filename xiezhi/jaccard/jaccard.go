package jaccard

func ComputeJaccardCoefficient(setA, setB []string) float64 {
	intersection := make(map[string]bool)
	union := make(map[string]bool)

	for _, item := range setA {
		union[item] = true
	}

	for _, item := range setB {
		if union[item] {
			intersection[item] = true
		}
		union[item] = true
	}

	return float64(len(intersection)) / float64(len(union))
}
