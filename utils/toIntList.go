package utils

import (
	"errors"
	"strconv"
)

func ToIntList(stringList []string) ([]int, error) {
	var result []int
	for _, stringItem := range stringList {
		item, err := strconv.Atoi(stringItem)
		if err != nil {
			return nil, errors.New("error while converting " + stringItem + " : " + err.Error())
		}

		result = append(result, item)
	}

	return result, nil
}