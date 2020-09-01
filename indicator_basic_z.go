package techan

import "github.com/sdcoffey/big"

type hlhPriceIndicator struct {
	*TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewHighlowhalfPriceIndicator(series *TimeSeries) Indicator {
	return hlhPriceIndicator{series}
}

func (hlh hlhPriceIndicator) Calculate(index int) big.Decimal {
	numerator := hlh.Candles[index].MaxPrice.Add(hlh.Candles[index].MinPrice) // .Add(tpi.Candles[index].ClosePrice)
	return numerator.Div(big.NewFromString("2"))
}
