package techan

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKdjKIndicator_Calculate(t *testing.T) {
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

	kdjk := NewKdjKIndicator(series, 9)
	// fmt.Println("kdjk.Calculate(i).FormattedString(0): ", kdjk.Calculate(0))
	// fmt.Println("kdjk.Calculate(i).FormattedString(1): ", kdjk.Calculate(1))
	// fmt.Print("kdjk.Calculate(i).FormattedString(2): ", kdjk.Calculate(1).FormattedString(2))

	// results := []string{"47.3684", "20.0000", "14.2857", "33.3333", "25.0000",
	// 	"25.0000", "29.0000", "29.0000", "29.0000", "29.0000",
	// }
	results := []string{"50.00", "40.00", "31.43", "32.06"}
	// fmt.Print("results ", results)

	for i, result := range results {
		// fmt.Println("*************************************************************i: ", i)
		// fmt.Println("kdjk.Calculate(i).FormattedString(i): ", kdjk.Calculate(i).FormattedString(2))
		// fmt.Println("result: ", result)
		assert.EqualValues(t, result, kdjk.Calculate(i).FormattedString(2))
	}
}
