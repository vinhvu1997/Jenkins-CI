package main

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {
    fmt.Print("\n**************\n**  JENKINS!  **\n**************\n\n")
    total := Sum(6, 5)
    if total != 11 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 11)
    }
    total = Sum(8,8)
    if total != 16 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 16)
    }
    total = Sum(0,0)
    if total != 0 {
	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
    }
    fmt.Println("Done!")
}
