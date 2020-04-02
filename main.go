package main

import (
	"flag"
	"fmt"
	"strings"
)

func state(input string) string {
	currentState := "S0"
	resultState := ""
	states := [][]string{{"S0", "0", "S0", "0"}, {"S0", "1", "S1", "1"}, {"S1", "0", "S2", "2"}, {"S1", "1", "S0", "0"}, {"S2", "0", "S1", "1"}, {"S2", "1", "S2", "2"}}

	values := strings.Split(input, "")
out:
	for i := 0; i < len(values); i++ {
		if values[i] != "0" && values[i] != "1" {
			return fmt.Sprintf("invalid input %v at position %v \n", values[i], i+1)
		}
		for v := 0; v < len(states); v++ {
			if currentState == states[v][0] && states[v][1] == values[i] {
				currentState = states[v][0]
				resultState = states[v][2]
				if i == len(values)-1 {
					return fmt.Sprintf("%v = %v \n", resultState, states[v][3])
				}
				currentState = states[v][2]
				continue out
			}
		}
	}
	return ""
}

func main() {
	var input string
	flag.StringVar(&input, "i", "", "string of zeros and ones")
	flag.Parse()

	result := state(input)
	fmt.Printf(result)
}
