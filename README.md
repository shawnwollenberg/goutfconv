# goutfconv

package main

import (
	convertutf8 "Convertutf8"
	"fmt"
	"strings"
)

func main() {
	xFile := "X:\\filepath.CSV"
	x := convertutf8.ReturnData(xFile)
	y := strings.Split(x, "\n")

	for _, z := range y {
		fmt.Println(z)
	}
}
