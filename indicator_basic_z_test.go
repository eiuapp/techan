package techan

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestHighlowhalfPriceIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	})
	candle.MinPrice = big.NewFromString("1.31")
	candle.MaxPrice = big.NewFromString("1.11")
	// candle.ClosePrice = big.NewFromString("1.215")

	series.AddCandle(candle)

	typicalPrice := NewHighlowhalfPriceIndicator(series).Calculate(0)

	assert.EqualValues(t, "1.21", typicalPrice.FormattedString(2))
	assert.EqualValues(t, "1.2100", typicalPrice.FormattedString(4))
}
