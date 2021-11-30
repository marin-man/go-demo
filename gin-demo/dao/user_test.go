package dao

import (
	"fmt"
	"testing"
)

func TestGetUserByPhone(t *testing.T) {
	user, _ := GetUserByPhone("12345678901")
	fmt.Printf("user = %v", user)
}