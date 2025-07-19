package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseIntegerParam(paramName string, c *gin.Context) int {
	param, err := strconv.Atoi(c.Params.ByName(paramName))
	if err != nil {
		return -1
	}
	return param
}

func GetNewObjectID[T any](objects []T, getID func(T) int) int {
	maxID := 0
	for _, obj := range objects {
		id := getID(obj)
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
