package v1

import "math"

func GetTotalPage(countHistories, pageSize int64) int {
	d := float64(countHistories) / float64(pageSize)
	return int(math.Ceil(d))
}

func GetHasMore(currentPage int64, totalCount int64, pageSize int64) bool {
	return currentPage < int64(math.Ceil(float64(totalCount)/float64(pageSize)))
}
