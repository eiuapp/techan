package techan

// 计算最近window个周期内的K值。KDJ指标中会使用到。

import (
	// "fmt" 

	"github.com/sdcoffey/big"
)

type kdjkIndicator struct {
	rsv    Indicator
	window int
}

// NewKdjKIndicator returns an Indicator which returns the k value of a candle for a given index.
// The k value is an average of the high, low, and close prices for a given candle.
func NewKdjKIndicator(series *TimeSeries, window int) Indicator {
	return kdjkIndicator{
		rsv:    NewRsvIndicator(series, window),
		window: window,
	}
}

func (kdjk kdjkIndicator) Calculate(index int) big.Decimal {
	if index < 1 {
		return big.TEN.Mul(big.NewDecimal(5))
	}
	// fmt.Println(" ++++++++++++++++++++++ index: ", index)
	// fmt.Println("kdjk.rsv.Calculate(index): ", kdjk.rsv.Calculate(index))
	preK := kdjk.Calculate(index - 1)
	// fmt.Println("preK: ", preK)
	part1 := preK.Mul(big.NewDecimal(2 / 3.0))
	// fmt.Println("part1: ", part1)
	part2 := kdjk.rsv.Calculate(index).Mul(big.NewDecimal(1 / 3.0))
	// fmt.Println("part2: ", part2)
	k := part1.Add(part2)
	// fmt.Println("k: ", k)
	return k
}
