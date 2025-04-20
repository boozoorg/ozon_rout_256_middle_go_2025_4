package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
		var n int
		fmt.Fscan(in, &n)
		mp := make(map[string]int)
		action := ""
		for j := 0; j < n; j++ {
			var (
				speaker,
				action1,
				action2,
				action3 string
				isNot bool
			)
			fmt.Fscan(in, &speaker, &action1, &action2, &action3)
			if _, ok := mp[speaker[:len(speaker)-1]]; !ok {
				mp[speaker[:len(speaker)-1]] = 0
			}
			if action3 == "not" {
				isNot = true
				fmt.Fscan(in, &action)
			} else if action == "" {
				action = action3
			}
			if action2 == "is" {
				if isNot {
					mp[action1]--
				} else {
					mp[action1]++
				}
			} else {
				if isNot {
					mp[speaker[:len(speaker)-1]]--
				} else {
					mp[speaker[:len(speaker)-1]] += 2
				}
			}
		}
		resp := " is " + action[:len(action)-1] + "."
		m := math.MinInt
		for _, v := range mp {
			m = max(m, v)
		}
		var suspects []string
		for k, v := range mp {
			if v == m {
				suspects = append(suspects, k)
			}
		}
		sort.Strings(suspects)
		for i := 0; i < len(suspects); i++ {
			fmt.Fprintln(out, suspects[i]+resp)
		}
	}
}
