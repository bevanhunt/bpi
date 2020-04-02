package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	currentState := "S0"
	resultState := ""
	states := [][]string{{"S0", "0", "S0"}, {"S0", "1", "S1"}, {"S1", "0", "S2"}, {"S1", "1", "S0"}, {"S2", "0", "S1"}, {"S2", "1", "S2"}}

	var input string
	flag.StringVar(&input, "i", "", "string of zeros and ones")
	flag.Parse()

	values := strings.Split(input, "")
out:
	for i := 0; i < len(values); i++ {
		if values[i] != "0" && values[i] != "1" {
			fmt.Printf("invalid input %v at position %v \n", values[i], i+1)
			os.Exit(4)
		}
		for v := 0; v < len(states); v++ {
			if currentState == states[v][0] && states[v][1] == values[i] {
				currentState = states[v][0]
				resultState = states[v][2]
				if i == len(values)-1 {
					fmt.Printf("%v = %v \n", resultState, values[i])
				}
				currentState = states[v][2]
				continue out
			}
		}
	}
}
