package main

import (
	"reflect"
	"testing"
)

func TestGetImageLayers(t *testing.T) {
	input := "123456789012"
	width := 3
	height := 2
	expectedResult := [][]string{{"123", "456"}, {"789", "012"}}

	result := getImageLayers(input, width, height)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}

func TestResolvedLayer(t *testing.T) {
	input := "0222112222120000"
	width := 2
	height := 2
	expectedResult := []string{"01", "10"}
	result := resolvedLayer(getImageLayers(input, width, height))
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}
