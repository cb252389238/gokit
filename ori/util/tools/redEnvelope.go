package tools

import (
	"errors"
	"math/rand"
	"time"
)

// GenerateRedPackets 生成指定总金额（单位：分）和数量的红包
func GenerateRedPackets(totalAmount int, count int, avgAmountBoundary int) ([]int, error) {
	// 参数校验
	if totalAmount < count {
		return nil, errors.New("总金额不能小于红包数量")
	}
	if count <= 0 {
		return nil, errors.New("红包数量必须大于0")
	}
	if count == 1 {
		return []int{totalAmount}, nil
	}

	// 计算单个红包最大金额（总金额的2/3）
	maxPer := avgAmountBoundary
	if maxPer < 1 {
		maxPer = 1
	}

	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 初始化红包数组
	packets := make([]int, count)
	remaining := totalAmount

	// 使用二倍均值法生成红包
	for i := 0; i < count-1; i++ {
		// 计算当前红包的最大可能金额
		// 取以下两者的较小值：1) 最大金额限制 2) 剩余金额的均值*2
		avg := remaining / (count - i)
		maxCurrent := min(maxPer, 2*avg)

		// 确保后续红包至少有1分钱
		maxCurrent = min(maxCurrent, remaining-(count-i-1))

		// 红包最小金额至少1分钱
		minCurrent := 1

		// 如果计算出的最大值小于最小值，则调整
		if maxCurrent < minCurrent {
			maxCurrent = minCurrent
		}

		// 生成随机金额
		money := minCurrent
		if maxCurrent > minCurrent {
			money = minCurrent + rand.Intn(maxCurrent-minCurrent+1)
		}

		packets[i] = money
		remaining -= money
	}

	// 最后一个红包等于剩余金额
	packets[count-1] = remaining

	// 验证最后一个红包是否超过限制
	if packets[count-1] > maxPer {
		// 如果超过限制，重新生成（递归调用，但有终止条件）
		if count > 1 && totalAmount > 1 {
			return GenerateRedPackets(totalAmount, count, avgAmountBoundary)
		}
		return nil, errors.New("无法生成满足条件的红包")
	}
	return packets, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
