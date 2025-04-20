package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Relation struct {
	From, To string
	Action   string
	Age      int
}

func main() {
	var in *bufio.Scanner
	var out *bufio.Writer
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Scan()
	t, _ := strconv.Atoi(in.Text())
	for i := 0; i < t; i++ {
		in.Scan()
		target := strings.TrimSuffix(strings.TrimPrefix(in.Text(), "How old is "), "?")
		var facts []string
		for j := 0; j < 3; j++ {
			in.Scan()
			facts = append(facts, in.Text())
		}

		var (
			ages     = make(map[string]int)
			checked  = make(map[string]bool)
			relation []Relation
		)

		for _, line := range facts {
			var r Relation
			parts := strings.Fields(line)
			if parts[2] == "the" {
				r = Relation{From: parts[0], To: parts[6], Action: "=", Age: 0}
			}

			age, _ := strconv.Atoi(parts[2])
			if parts[4] == "older" {
				r = Relation{From: parts[0], To: parts[6], Action: "+", Age: age}
			}
			if parts[4] == "younger" {
				r = Relation{From: parts[0], To: parts[6], Action: "-", Age: age}
			}

			if r.To == "" {
				ages[parts[0]], checked[parts[0]] = age, true
			} else {
				relation = append(relation, r)
			}
		}

		for updated := true; updated; {
			updated = false
			for _, r := range relation {
				fChecked, fAge, tChecked, tAge := checked[r.From], ages[r.From], checked[r.To], ages[r.To]
				switch r.Action {
				case "=":
					if fChecked && !tChecked {
						ages[r.To], checked[r.To], updated = fAge, true, true
					} else if !fChecked && tChecked {
						ages[r.From], checked[r.From], updated = tAge, true, true
					}
				case "+":
					if fChecked && !tChecked {
						ages[r.To], checked[r.To], updated = fAge-r.Age, true, true
					} else if !fChecked && tChecked {
						ages[r.From], checked[r.From], updated = tAge+r.Age, true, true
					}
				case "-":
					if fChecked && !tChecked {
						ages[r.To], checked[r.To], updated = fAge+r.Age, true, true
					} else if !fChecked && tChecked {
						ages[r.From], checked[r.From], updated = tAge-r.Age, true, true
					}
				}
			}
		}
		fmt.Fprintln(out, ages[target])
	}
}
