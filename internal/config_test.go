package internal

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := NewConfig("../config/config.ini")
	fmt.Println(config)
}
