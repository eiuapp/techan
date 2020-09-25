package techan

// 计算最近window个周期内的RSV值。KDJ指标中会使用到。

import (
	// "fmt"

	"github.com/sdcoffey/big"
)

type rsvIndicator struct {
	lowv   Indicator
	highv  Indicator
	close  Indicator
	window int
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewRsvIndicator(series *TimeSeries, window int) Indicator {
	return rsvIndicator{
		lowv:   NewLowvIndicator(series, window),
		highv:  NewHighvIndicator(series, window),
		close:  NewClosePriceIndicator(series),
		window: window,
	}
}

func (rsv rsvIndicator) Calculate(index int) big.Decimal {
	// if index < rsv.window-1 {
	// 	return big.ZERO
	// }
	// fmt.Println(" ++++++++++++++++++++++ index: ", index)
	// fmt.Println("rsv.close.Calculate(index): ", rsv.close.Calculate(index))
	// fmt.Println("rsv.highv.Calculate(index): ", rsv.highv.Calculate(index))
	// fmt.Println("rsv.lowv.Calculate(index): ", rsv.lowv.Calculate(index))
	fenzi := rsv.close.Calculate(index).Sub(rsv.lowv.Calculate(index))
	// fmt.Println("fenzi: ", fenzi)
	fenmu := rsv.highv.Calculate(index).Sub(rsv.lowv.Calculate(index))
	// fmt.Println("fenmu: ", fenmu)
	oneHundred := big.TEN.Mul(big.TEN)
	rsvValue := fenzi.Div(fenmu).Mul(oneHundred)
	// fmt.Println("rsvValue: ", rsvValue)
	return rsvValue
}
