package ep

import "sort"

func SliceAsc(arry []int){
	sort.Ints(arry)
}

func SliceDesc(arry []int){
	sort.Sort(intSlice(arry))
}

type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
