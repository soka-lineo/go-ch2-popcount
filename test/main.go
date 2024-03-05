package main

import (
	"fmt"
	"os"
	"strconv"

	popcount "github.com/soka-lineo/go-ch2-popcount"
)

func main() {
	v, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		os.Exit(1)
	}
	c := popcount.PopCount(v)
	fmt.Printf("PopCount(%v:0b%b)=%d\n", v, v, c)
}
