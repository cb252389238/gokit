package sharedCode

type SharedCodeService struct {
	maxValidID    uint32
	feistelRounds int
	base62Chars   string
}

func NewSharedCodeService() *SharedCodeService {
	return &SharedCodeService{
		maxValidID:    916132832,
		feistelRounds: 3,
		base62Chars:   "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
}

// Feistel网络实现
func (s *SharedCodeService) feistelNetwork(id uint32) uint32 {
	left := uint16(id >> 16)
	right := uint16(id & 0xFFFF)
	for i := 0; i < s.feistelRounds; i++ {
		temp := left
		f := uint16((uint64(right)*11400714819323199549 + 1442695040888963407) >> 48)
		left = right
		right = temp ^ f
	}
	return (uint32(left) << 16) | uint32(right)
}
