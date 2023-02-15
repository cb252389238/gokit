package common

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
