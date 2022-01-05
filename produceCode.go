package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func produceCodeSubsetString() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
  b := make([]byte, 4)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}
func createProduceCode() string {
	return fmt.Sprintf(`%s-%s-%s-%s`, produceCodeSubsetString(), produceCodeSubsetString(), produceCodeSubsetString(), produceCodeSubsetString(),)
}