package xiezhi

import (
	"strings"
	"sync"
	"xiezhi/cosinesim"
	"xiezhi/fenci"
	"xiezhi/jaccard"
	"xiezhi/minhash"
	"xiezhi/simhash"
	"xiezhi/util/charchar"
)

var (
	//jieba *gojieba.Jieba
	once sync.Once
	gose *fenci.FenCi
)

//func newJieba() *gojieba.Jieba {
//	once.Do(func() {
//		x := gojieba.NewJieba()
//		jieba = x
//	})
//	return jieba
//}

func NewFenCi() *fenci.FenCi {
	once.Do(func() {
		x := fenci.NewFenCi()
		gose = x
	})
	return gose
}

// 分词
func cut(text string) []string {
	text = charchar.RemovePunct(text)        //去除标点符号
	text = charchar.RemoveNonsenseWord(text) //去除无意义词语
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "\t", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	//return newJieba().Cut(text, true)
	//fmt.Println(newFenCi().Cut(text))
	return NewFenCi().Cut(text)
}

// 获取文档min hash签名
func SimHash(text string) uint64 {
	hash := simhash.Simhash(cut(text))
	return hash
}

// 对比两个文档hash相似性
// 返回海明距离和相似性
func SimHashSimilarity(hash1, hash2 uint64) (int, float64) {
	return simhash.Similarity(hash1, hash2)
}

// 获取hash签名
func MinHash(text string) []uint32 {
	return minhash.ComputeMinHashSignature(cut(text))
}

func MinHashSimilarity(hash1, hash2 []uint32) float64 {
	return minhash.ComputeSimilarity(hash1, hash2)
}

// 杰卡德系数计算
func Jaccard(text1, text2 string) float64 {
	coefficient := jaccard.ComputeJaccardCoefficient(cut(text1), cut(text2))
	return coefficient
}

// 余弦相似度计算
func CosineSim(text1, text2 string) float64 {
	frequency1 := cosinesim.CalculateTermFrequency(cut(text1))
	frequency2 := cosinesim.CalculateTermFrequency(cut(text2))
	similarity := cosinesim.ComputeCosineSimilarity(frequency1, frequency2)
	return similarity
}
