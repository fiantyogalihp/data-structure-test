package data_strcuture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckTheValueIsExistingWithNoForLoop(t *testing.T) {
	arr := []string{"admin", "operator", "finance"}
	currentProfile := "finance"

	isExist := isTheValueIsExisting(arr, currentProfile)

	assert.Equal(t, true, isExist, "the finance's profile is not found")
}

func isTheValueIsExisting(arr []string, currentProfile string) bool {
	// code here
	// do not use for loop

	return false
}

func TestCheckTheValueIsDuplicate(t *testing.T) {
	inputData := []string{"019238019283", "6182736187236", "126381760442553"}
	databaseData := []string{"6182736187237", "019238019284", "126381760442553"}

	isExist := isTheValuesIsDuplicate(inputData, databaseData)

	assert.Equal(t, false, isExist, "the duplicate value is not found")
}

func isTheValuesIsDuplicate(inputData []string, databaseData []string) bool {
	// update these code correctly
	for i, v := range inputData {
		if v == databaseData[i] {
			return true
		}
	}

	return false
}
