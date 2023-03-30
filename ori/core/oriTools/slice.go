package oriTools

func RemoveSliceRepeat(sli []string) []string {
	result := []string{}
	m := make(map[string]bool)
	for _, v := range sli {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

func InArray[T comparable](val T, arr []T) bool {
	ok := false
	for _, v := range arr {
		if v == val {
			ok = true
			break
		}
	}
	return ok
}
