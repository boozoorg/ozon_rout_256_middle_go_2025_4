package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		histograms := make([][]int, n)
		for j := 0; j < n; j++ {
			line := make([]int, m)
			for k := 0; k < m; k++ {
				fmt.Fscan(in, &line[k])
				histograms[j] = line
			}
		}

		c := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if isMatching(histograms[i], histograms[j]) {
					c++
				}
			}
		}

		fmt.Fprintln(out, c)
	}
}

func isMatching(a, b []int) bool {
	m := len(a)
	sum := a[0] + b[m-1]
	for i := 0; i < m; i++ {
		if a[i]+b[m-1-i] != sum {
			return false
		}
	}
	return true
}
