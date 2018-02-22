package main

import (
	"strconv"
	"testing"
)

func TestSampleOne(t *testing.T) {
	expected := "11"
	result := strconv.Itoa(10)
	compare(expected, result, t)
}

func TestSampleTwo(t *testing.T) {
	expected := "11"
	result := strconv.Itoa(10)
	compareWithHelper(expected, result, t)
}

func TestSampleThree(t *testing.T) {
	expected := "10"
	result := strconv.Itoa(10)
	compare(expected, result, t)
}

func compareWithHelper(expected, result string, t *testing.T) {
	t.Helper()
	if expected != result {
		t.Fatalf("Expected result %v does not match result %v",
			expected,
			result)
	}
}

func compare(expected, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("Fail: Expected result %v does not match result %v",
			expected,
			result)
	}
	t.Logf("OK: Expected result %v = %v",
		expected,
		result)
}
