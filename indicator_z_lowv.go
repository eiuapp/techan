package techan

// 计算最近window个周期内的最低价。KDJ指标中会使用到。

import (
	// "fmt"

	"github.com/sdcoffey/big"
)

type lowvIndicator struct {
	lowPriceIndicator Indicator
	window            int
}

// NewLowvIndicator returns an Indicator which returns the lowv price of a candle for a given index.
// The lowv price is the lowest of the low prices for a given candle.
func NewLowvIndicator(series *TimeSeries, window int) Indicator {
	return lowvIndicator{
		// series: series,
		lowPriceIndicator: NewLowPriceIndicator(series),
		window:            window,
	}
}

func (lowv lowvIndicator) Calculate(index int) big.Decimal {
	// if index < 1 {
	// 	// return big.ZERO
	// 	return big.TEN.Mul(big.TEN).Mul(big.TEN).Mul(big.TEN).Mul(big.TEN).Mul(big.TEN)
	// }

	lv := lowv.lowPriceIndicator.Calculate(0)
	// lowIndex := 0
	if index < lowv.window {
		// for i := 0; i <= lowv.window; i++ {
		for i := 0; i <= index; i++ {
			value := lowv.lowPriceIndicator.Calculate(i)
			if value.LT(lv) {
				lv = value
				// lowIndex = i
			}
		}
	} else {
		for i := (index + 1) - lowv.window; i <= index; i++ {
			value := lowv.lowPriceIndicator.Calculate(i)
			if value.LT(lv) {
				lv = value
				// lowIndex = i
			}
		}
	}
	return lv
}

// [1,2,3,4,5,6,7,8,9,10,11,12]

// [12,11,10,9,8,7,6,5,4,3,2,1]
//　远，＝＝＝＝＝＝＝＝＝＝＝》 近
// window: 9, index : 11, i: from (11 + 1) - 9 to 11, lowest: 2
// window: 9, index : 10, i: from (10 + 1) - 9 to 10, lowest: 2
// window: 9, index : 9, i: from (9 + 1) - 9 to 9, lowest: 2
// window: 9, index : 8, i: from (8 + 1) - 9 to 8, lowest: 2
// window: 9, index : 7, i: from (7 + 1) - 9 to 8, lowest: 2
// window: 9, index : 0, i: 0-0, lowest: 1
// window: 9, index : 1, i: 0-1, lowest: 2
// window: 9, index : 8, i: 0-8, lowest: 2
