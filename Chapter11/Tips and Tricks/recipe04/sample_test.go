package main

import (
	"fmt"
	"strconv"
	"testing"
)

var testData = []int{10, 11, 017}

func TestSampleOne(t *testing.T) {
	expected := "10"
	for _, val := range testData {
		tc := val
		t.Run(fmt.Sprintf("input = %d", tc), func(t *testing.T) {
			if expected != strconv.Itoa(tc) {
				t.Fail()
			}
		})
	}
}
