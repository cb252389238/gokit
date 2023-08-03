package fenci

import "github.com/go-ego/gse"

type FenCi struct {
	seg gse.Segmenter
}

func NewFenCi() *FenCi {
	f := &FenCi{
		seg: gse.Segmenter{},
	}
	dict := []string{
		"dict/zh/s_1.txt",
		"dict/zh/idf.txt",
		"dict/zh/stop_tokens.txt",
		"dict/zh/stop_word.txt",
		"dict/zh/t_1.txt",
		"dict/en/dict.txt",
	}
	err := f.seg.LoadDict(dict...)
	if err != nil {
		panic(err)
	}
	return f
}

func (f *FenCi) Cut(text string) []string {
	cut := f.seg.Cut(text, true)
	return cut
}
