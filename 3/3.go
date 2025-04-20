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
		var n, k, sum int
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}

		dif := make([]int, n+2)
		resp := true
		for i := 0; i < n; i++ {
			sum += dif[i]
			if a[i]+sum < 0 || (a[i]+sum > 0 && i+k > n) {
				resp = false
				break
			}
			if a[i]+sum > 0 {
				use := a[i] + sum
				sum -= use
				dif[i+k] += use
			}
		}

		if resp {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
