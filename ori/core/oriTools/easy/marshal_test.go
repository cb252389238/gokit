package easy

import "testing"

type Resp struct {
	Code int
	Data any
}

func TestMarshal(t *testing.T) {
	res1 := Resp{
		Code: 200,
		Data: "hello",
	}
	t.Log(Marshal(res1))
	res2 := Resp{
		Code: 200,
		Data: map[string]any{
			"name": "hello",
		},
	}
	t.Log(Marshal(res2))
	var data3 []string
	res3 := Resp{
		Code: 200,
		Data: data3,
	}
	t.Log(Marshal(res3))
	var data4 map[string]int
	res4 := Resp{
		Code: 200,
		Data: data4,
	}
	t.Log(Marshal(res4))
}
