package techan

import (
	// "fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRsvIndicator_Calculate(t *testing.T) {
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

	rsv := NewRsvIndicator(series, 9)
	// fmt.Println("rsv.Calculate(i).FormattedString(0): ", rsv.Calculate(0))
	// fmt.Println("rsv.Calculate(i).FormattedString(1): ", rsv.Calculate(1))
	// fmt.Print("rsv.Calculate(i).FormattedString(2): ", rsv.Calculate(1).FormattedString(2))

	// results := []string{"22.0000", "24.0000", "24.0000", "25.0000", "25.0000",
	// 	"25.0000", "29.0000", "29.0000", "29.0000", "29.0000",
	// }
	results := []string{"47.3684", "20.0000", "14.2857", "33.3333"}
	// fmt.Print("results ", results)

	for i, result := range results {
		// fmt.Println("i: ", i)
		// fmt.Println("rsv.Calculate(i).FormattedString(i): ", rsv.Calculate(i).FormattedString(4))
		// fmt.Println("result: ", result)
		assert.EqualValues(t, result, rsv.Calculate(i).FormattedString(4))
	}
}
