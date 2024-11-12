package easy

import "testing"

func TestEmpty(t *testing.T) {
	t.Log(Empty(1))
	t.Log(Empty(""))
	t.Log(Empty(nil))
	t.Log(Empty(0))
	t.Log(Empty(false))
	t.Log(Empty([]int{}))
	t.Log(Empty(map[string]string{}))
	t.Log(Empty(make(chan int)))
	t.Log(Empty(make(chan int, 1)))
	t.Log(Empty(make([]int, 0)))
	t.Log(Empty(make(map[string]string)))
	t.Log(Empty(make(map[string]string, 0)))
	t.Log(Empty(make(map[string]string, 1)))
	t.Log(Empty(make(map[string]string, 2)))
	t.Log(Empty([]int{1, 3, 4}))
	t.Log(Empty(map[string]string{"a": "b"}))
}

func TestIsNumeric(t *testing.T) {
	t.Log(IsNumeric(1))
	t.Log(IsNumeric("1"))
	t.Log(IsNumeric(1.1))
	t.Log(IsNumeric("1.1"))
	t.Log(IsNumeric("a"))
	t.Log(IsNumeric([]int{}))
	t.Log(IsNumeric(map[string]string{}))
	t.Log(IsNumeric(make(chan int)))
	t.Log(IsNumeric(make(chan int, 1)))
	t.Log(IsNumeric(make([]int, 0)))
	t.Log(IsNumeric(make(map[string]string)))
	t.Log(IsNumeric(make(map[string]string, 0)))

}

func TestVersionCompare(t *testing.T) {
	t.Log(VersionCompare("1.2.3", "1.2.3", "=="))
	t.Log(VersionCompare("1.2.3", "1.2.3", ">="))
	t.Log(VersionCompare("1.2.3", "1.2.3", ">"))
	t.Log(VersionCompare("1.2.3", "1.2.3", "<="))
	t.Log(VersionCompare("1.2.3", "1.2.3", "<"))
	t.Log(VersionCompare("1.2.3", "1.2.3", "!="))
	t.Log(VersionCompare("1.2.3", "1.2.4", "=="))
	t.Log(VersionCompare("1.2.3", "1.2.4", ">="))
	t.Log(VersionCompare("1.2.3", "1.2.4", ">"))
	t.Log(VersionCompare("1.2.3", "1.2.4", "<="))
	t.Log(VersionCompare("1.2.3", "1.2.4", "<"))
	t.Log(VersionCompare("1.2.3", "1.2.4", "!="))
	t.Log(VersionCompare("1.2.3", "1.2.5", "=="))
	t.Log(VersionCompare("1.2.3", "1.2.5", ">="))
	t.Log(VersionCompare("1.2.3", "1.2.5", ">"))
	t.Log(VersionCompare("1.2.3", "1.2.5", "<="))
	t.Log(VersionCompare("1.2.3", "1.2.5", "<"))
	t.Log(VersionCompare("1.2.3", "1.2.5", "!="))
	t.Log(VersionCompare("1.2.3.1", "1.2.3.1", "=="))
}
