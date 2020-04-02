package main

import "testing"

func TestStateFirst(t *testing.T) {
	got := state("110")
	if got != "S0 = 0 \n" {
		t.Error("S0 = 0 \n", got)
	}
}

func TestStateSecond(t *testing.T) {
	got := state("1010")
	if got != "S1 = 1 \n" {
		t.Error("S1 = 1 \n", got)
	}
}

func TestStateThird(t *testing.T) {
	got := state("1012")
	if got != "invalid input 2 at position 4 \n" {
		t.Error("invalid input 2 at position 4 \n", got)
	}
}
