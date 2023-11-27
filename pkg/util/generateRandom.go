package util

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GenerateRandom(numDigits int) string {
	rand.Seed(time.Now().UnixNano())
	randonNumer := rand.Int31n(int32(math.Pow10(numDigits)))
	// 将随机数转换为numDigits位字符串
	formatString := fmt.Sprintf("%%0%dd", numDigits)
	result := fmt.Sprintf(formatString, randonNumer)
	return result
}
