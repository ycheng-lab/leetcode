# [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

## Problem

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```text
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

```go
func convert(text string, nRows int) string
```

convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

