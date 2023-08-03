package xiezhi

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	text1 = `1927年4月21日，路易·艾黎踏上了中国的土地，第一站便是上海。当时的上海正处于“四一二”反革命政变后的社会动乱之中。4月22日，路易·艾黎开始在上海求职。
他带着新西兰一位朋友所写的介绍信，寻找到一位酒商。这位商人马上给上海工部局火政处的副处长打了个电话。适逢火政处正在招聘，副处长看了路易·艾黎的证件，发现他非常符合火政处的工作要求，便将他引荐给了火政处处长道森上尉。
在回忆录中，路易·艾黎这样叙述当时的情景`
	text2 = `1927年4月21日，路易·艾黎踏上了中国的土地，第一站便是上海。当时的上海正处于“四一二”反革命政变后的社会动乱之中。4月22日，路易·艾黎开始在上海求职。
他带着新西兰一位朋友所写的介绍信，寻找到一位酒商。副处长看了路易·艾黎的证件，发现他非常符合火政处的工作要求，便将他引荐给了火政处处长道森上尉。
在回忆录中，路易·艾黎这样叙述当时的情景：“上尉与我谈了一会儿，让我到外面的小房间里等着，我听见这位处长在副处长的办公室里谈论我的声音”，副处长说：“我相信这个人很有用处，我们应该高兴地接受他。”处长嘟囔了一句“年轻的愣小子”，但随即同意了。
副处长把路易·艾黎叫进去，说：“今天下午，虹口救火会将派车把你接到那里去。`
)

func TestMinHash(t *testing.T) {
	hash1 := MinHash(text1)
	hash2 := MinHash(text2)
	fmt.Println(hash1)
	fmt.Println(hash2)
	similarity := MinHashSimilarity(hash1, hash2)
	fmt.Println(similarity)
}

func BenchmarkMinHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//hash1 := MinHash(text1)
		//hash2 := MinHash(text2)
		//MinHashSimilarity(hash1, hash2)
		hash1 := []uint32{21786534, 3813651, 16491142, 2760681, 9994562, 33660666, 3625849, 19068251, 4967947, 13097248, 2278835, 4982908, 22058703, 12431822, 12942044, 20791561, 26708247, 13803401, 1402089, 10877545, 8272193, 6607987, 7469685, 11493166, 19324761, 2609975, 1124155, 21462898, 3044391, 8063556, 54948517, 6836503, 18129241, 15558320, 4233603, 1181630, 12670454, 28852647, 25153343, 1135842, 318534, 19409078, 24839731, 2240189, 2420116, 11726433, 16155452, 14535957, 1719871, 11952521, 14460205, 2837559, 4313463, 11049864, 19108135, 4206265, 8873296, 4617667, 14646660, 11380295, 14735902, 22482221, 5287110, 21933407, 12769836, 19460998, 9514836, 8266152, 18506565, 28503976, 10430642, 25027779, 32201938, 15951839, 34199406, 779626, 7829923, 13797329, 189841, 3372463, 10060218, 2105443, 44582558, 136640, 3286445, 9133171, 21245538, 27936188, 24190666, 8120745, 9391144, 12162198, 441428, 19652928, 20509507, 17248708, 8832302, 54387721, 11621792, 8229802}
		MinHashSimilarity(hash1, hash1)
	}
}

func TestSimHash(t *testing.T) {
	hash1 := SimHash(text1)
	hash2 := SimHash(text2)
	fmt.Println(hash1)
	fmt.Println(hash2)
	fmt.Println(strconv.FormatInt(int64(hash1), 2))
	fmt.Println(strconv.FormatInt(int64(hash2), 2))
	similarity, f := SimHashSimilarity(hash1, hash2)
	fmt.Println(similarity, f)
}

func BenchmarkSimHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimHashSimilarity(4572973431209003220, 4572973431209003220)
	}

}

func TestJaccard(t *testing.T) {
	jaccard := Jaccard(text1, text2)
	fmt.Println(jaccard)
}

func BenchmarkJaccard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jaccard(text1, text2)
	}
}

func TestCosineSim(t *testing.T) {
	sim := CosineSim(text1, text2)
	fmt.Println(sim)
}

func BenchmarkCosineSim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CosineSim(text1, text2)
	}
}
