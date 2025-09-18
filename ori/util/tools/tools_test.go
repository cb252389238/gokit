package tools

import "testing"

func TestGenerateRedPackets(t *testing.T) {
	for i := 0; i < 100; i++ {
		redPackets, err := GenerateRedPackets(1000, i, 1000*3/5)
		if err != nil {
			t.Errorf("Error generating red packets: %v", err)
		}
		t.Logf("Generated red packets: %v", redPackets)
	}
}
