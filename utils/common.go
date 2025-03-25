package utils

import "math"

func Percentage(x float64, sum float64) float64 {
	return RoundP((x/sum)*100.0, 3)
}

func RoundP(x float64, prec uint) float64 {
	ratio := math.Pow(10, float64(prec))
	return math.Round(x*ratio) / ratio
}
func Avg(v ...interface{}) float64 {
	return Total(v) / float64(len(v))
}

func Total(v ...interface{}) float64 {
	var total float64

	for _, val := range v {
		switch val.(type) {
		case uint, uint8, uint16, uint32, uint64:
			total += float64(val.(uint64))
		case int, int8, int16, int32, int64:
			total += float64(val.(int64))
		case float32, float64:
			total += float64(val.(float64))
		default:
			continue
		}
	}

	return total
}
