package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCrossUpIndicatorRule(t *testing.T) {
// 	upInd := NewFixedIndicator(3, 4, 5, 6)
// 	dnInd := NewFixedIndicator(6, 5, 4, 3)

// 	rule := NewCrossUpIndicatorRule(dnInd, upInd)

// 	t.Run("always returns false when index == 0", func(t *testing.T) {
// 		assert.False(t, rule.IsSatisfied(0, nil))
// 	})

// 	t.Run("Returns true when lower indicator crosses upper indicator", func(t *testing.T) {
// 		assert.False(t, rule.IsSatisfied(1, nil))
// 		assert.True(t, rule.IsSatisfied(2, nil))
// 		assert.True(t, rule.IsSatisfied(3, nil))
// 	})
// }

func TestCrossDownIndicatorRule(t *testing.T) {
	// dnInd := NewFixedIndicator(3, 4, 5, 6)
	// upInd := NewFixedIndicator(6, 5, 4, 3)
	upInd := NewFixedIndicator(3, 4, 5, 6)
	dnInd := NewFixedIndicator(6.1, 5.1, 4.1, 3.1)

	rule := NewCrossDownIndicatorRule(dnInd, upInd)

	t.Run("returns false when index == 0", func(t *testing.T) {
		assert.False(t, rule.IsSatisfied(0, nil))
	})

	t.Run("returns true when upper indicator crosses below lower indicator", func(t *testing.T) {
		assert.False(t, rule.IsSatisfied(1, nil))
		assert.True(t, rule.IsSatisfied(2, nil))
		assert.True(t, rule.IsSatisfied(3, nil))
	})

	closePriceIndicator := NewFixedIndicator(
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2000, 9927.56, 9924.17, 3)
	bollingerBandsUpIndicator := NewFixedIndicator(
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1000000, 9943.454276, 9925.627737, 6)
	rule2 := NewCrossDownIndicatorRule(closePriceIndicator, bollingerBandsUpIndicator)
	// fmt.Println("000000000000000")
	t.Run("returns true when upper indicator crosses below lower indicator", func(t *testing.T) {
		assert.True(t, rule2.IsSatisfied(31, nil))
		// assert.True(t, rule2.IsSatisfied(32, nil))
		// assert.True(t, rule2.IsSatisfied(33, nil))
	})
}
