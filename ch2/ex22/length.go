package main

import (
	"fmt"
	"gobook/ch2/ex22/lengthconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}
		m := lengthconv.Meters(l)
		f := lengthconv.Foot(l)
		fmt.Printf("%s = %s, %s =%s\n", m, lengthconv.MToF(m), f, lengthconv.FToM(f))
	}
}
