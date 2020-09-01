package techan

import "github.com/sdcoffey/big"

type awesomeOscillatorIndicator struct {
	series  *TimeSeries
	window1 int
	window2 int
}

// NewAOIndicator Returns a new Awesome Oscillator Indicator
// https://www.tradingview.com/support/solutions/43000501826-awesome-oscillator-ao/
func NewAOIndicator(ts *TimeSeries, window1, window2 int) Indicator {
	return awesomeOscillatorIndicator{
		series:  ts,
		window1: window1,
		window2: window2,
	}
}

func (aoi awesomeOscillatorIndicator) Calculate(index int) big.Decimal {
	hlhPrice := NewHighlowhalfPriceIndicator(aoi.series)
	hlhPriceSma1 := NewSimpleMovingAverage(hlhPrice, aoi.window1)
	hlhPriceSma2 := NewSimpleMovingAverage(hlhPrice, aoi.window2)
	// meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(aoi.series), aoi.window1)
	return hlhPriceSma1.Calculate(index).Sub(hlhPriceSma2.Calculate(index))
}

// Calculation
// lengthAO1=input(5, minval=1) //5 periods
// lengthAO2=input(34, minval=1) //34 periods
// AO = sma((high+low)/2, lengthAO1) - sma((high+low)/2, lengthAO2)

// 网上一个的通达信的版本，计算方式，有点不一样
// 通达信bai系统不带ao指标,所以需du要用户自己新建一个指标公zhi式使用.代码如下dao,效果如图.
// Y:=(H+L)/2;
// AO:EMA(Y,5)-EMA(Y,34);
// AO1:=REF(AO,1);
// STICKLINE(AO>AO1,0,AO,0,0),COLORRED;
// STICKLINE(AO<AO1,0,AO,0,0),COLORGREEN;
// 原理：动量震荡指标（AO）是从5根价格线的中点的移动平均线值减去34根价格线的中点的移动平均线值得来的。
// 通过将一系列所得结果组成柱状图能准确发现当前动量的变化。中点=(最高价+最低价)/2
