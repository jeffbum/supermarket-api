package main

import (
	"regexp"
	"testing"
)

func TestCreationOfProduceCodeSubsetString(t *testing.T) {
	substring := produceCodeSubsetString()
	length := len(substring)
	if  length != 4 {
		t.Errorf("Got %d, expected %d. They should be the same.", length, 4)
	}
}

func TestCreationOfProduceCode(t *testing.T) {
	produceCode := createProduceCode()
	length := len(produceCode)
	if  length != 19 {
		t.Errorf("Got %d, expected %d. They should be the same.", length, 19)
	}
}

func TestCreationOfProduceCodeNumberOfHypens(t *testing.T) {
	produceCode := createProduceCode()
	re, err := regexp.Compile("-")
	if err != nil {
		t.Errorf("there were no hypens in the produce code.")
	}
	replaced := re.ReplaceAllString(produceCode, "")
	length := 19 -len(replaced)
	if  length != 3 {
		t.Errorf("Got %d, expected %d. They should be the same.", length, 3)
	}
}