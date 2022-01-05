package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}
func createProduceCode() string {
	return fmt.Sprintf(`%s-%s-%s-%s`, StringWithCharset(4), StringWithCharset(4), StringWithCharset(4), StringWithCharset(4),)
}