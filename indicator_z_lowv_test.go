package techan

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowvIndicator_Calculate(t *testing.T) {
	series := mockTimeSeriesOCHL(
		[]string{"12", "21", "22", "20.1"}, // 离当前时间远
		[]string{"14", "20", "24", "19"},
		[]string{"19", "18", "22", "17"},
		[]string{"10", "19", "25", "16"},
		[]string{"14", "21", "24", "16.2"},
		[]string{"10", "20", "21", "14"},
		[]string{"12", "25", "29", "13"},
		[]string{"10", "19", "21", "12"},
		[]string{"14", "21", "24", "11"},
		[]string{"10", "19", "21", "10.1"},
		[]string{"12", "10", "22", "9"},
		[]string{"10", "19", "21", "8"}, // 离当前时间近
	)

	// series := mockTimeSeries(typicalPrices...)

	lowv := NewLowvIndicator(series, 9)
	// fmt.Println("lowv: ", lowv)

	// fmt.Println("lowv.Calculate(i).FormattedString(0): ", lowv.Calculate(0))
	// fmt.Println("lowv.Calculate(i).FormattedString(1): ", lowv.Calculate(1))
	// fmt.Print("lowv.Calculate(i).FormattedString(2): ", lowv.Calculate(1).FormattedString(2))

	results := []string{"20.1000", "19.0000", "17.0000", "16.0000", "16.0000",
		"14.0000", "13.0000", "12.0000", "11.0000", "10.1000",
	}
	// fmt.Print("results ", results)

	for i, result := range results {
		// fmt.Println("i: ", i)
		// fmt.Println("result: ", result)
		assert.EqualValues(t, result, lowv.Calculate(i).FormattedString(4))
	}
}
