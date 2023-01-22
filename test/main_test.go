package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("前処理")
	status := m.Run()
  fmt.Println("後処理")

  os.Exit(status)
}