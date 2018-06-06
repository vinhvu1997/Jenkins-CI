package main

import "testing"
import "fmt"

func TestSum(t *testing.T) {
    fmt.print("\n*************\n****TEST*****\n*************\n")
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
    total = Sum(8,8)
    if total != 16 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 16)
    }
    total = Sum(0,0)
    if total != 0 {
	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
    }
}
