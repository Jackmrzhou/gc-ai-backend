package utils

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	tStr := GenerateToken(1, "test@mail.com")
	claim, err := ParseToken(tStr)
	if err != nil{
		t.Fatal(err)
	}else{
		fmt.Println(claim)
	}
}
