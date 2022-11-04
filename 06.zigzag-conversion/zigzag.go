package main

import (
	"bytes"
	"fmt"
)

func zigzagConversion(text string, nRows int) string {
	if nRows <= 1 || len(text) <= nRows {
		return text
	}

	res := bytes.Buffer{}
	pace := 2*nRows - 2

	// process first row
	for i := 0; i < len(text); i += pace {
		res.WriteByte(text[i])
	}

	// process 2nd ... last-1 row
	for row := 1; row <= nRows-2; row++ {
		res.WriteByte(text[row])
		for k := pace; k-row < len(text); k += pace {
			res.WriteByte(text[k-row])
			if k+row < len(text) {
				res.WriteByte(text[k+row])
			}
		}
	}

	// process last row
	for i := nRows - 1; i < len(text); i += pace {
		res.WriteByte(text[i])
	}
	return res.String()
}

func solveProblem(text string, nRows int) {
	res := zigzagConversion("PAYPALISHIRING", 3)
	fmt.Printf("string %s, row: %d, result: %s\n", text, nRows, res)
	if res == "PAHNAPLSIIGYIR" {
		fmt.Println("Conversion is correct")
	} else {
		fmt.Println("Failed  to zigzag convert")
	}
}

func main() {
	solveProblem("PAYPALISHIRING", 3)

}
