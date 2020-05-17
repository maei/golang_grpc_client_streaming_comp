package service

var AverageService averageServiceInterface = &averageService{}

type averageServiceInterface interface {
	Average(data []int32) float32
}

type averageService struct{}

func (*averageService) Average(data []int32) float32 {
	var result int32
	for _, values := range data {
		result += values
	}

	return float32(result) / float32(len(data))
}
