package easy

import "testing"

func TestMbStrlen(t *testing.T) {
	t.Log(MbStrlen("你好"))
	t.Log(MbStrlen("你好，我是中国人"))
	t.Log(MbStrlen("你好，我是中国人，我是中国人,abc"))
}

func TestStrRev(t *testing.T) {
	t.Log(StrRev("abc"))
	t.Log(StrRev("abcdef"))
}

func TestStrShuffle(t *testing.T) {
	t.Log(StrShuffle("abcdef"))
	t.Log(StrShuffle("abcdefghijklmnopqrstuvwxyz"))
}

func TestParseStr(t *testing.T) {
	res := map[string]any{}
	err := ParseStr("a=1&b=2&c=3", res)
	t.Log(res, err)
	res = map[string]any{}
	err = ParseStr("a=1&b=2&c=3&a=4", res)
	t.Log(res, err)
	res = map[string]any{}
	err = ParseStr("a[]=1&a[]=2&a[]=3", res)
	t.Log(res, err)
	res = map[string]any{}
	err = ParseStr("a[b]=1&a[c]=2&a[d]=3", res)
	t.Log(res, err)
	res = map[string]any{}
	err = ParseStr("a[b][]=1&a[b][]=2&a[b][]=3", res)
	t.Log(res, err)
}

func TestCrc32(t *testing.T) {
	t.Log(Crc32("abc"))
	t.Log(Crc32("abcdef"))
	t.Log(Crc32("你好"))
}
