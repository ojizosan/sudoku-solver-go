package main

import (
	"errors"
	"fmt"
	"strings"
)

func cross(a string, b string) []string {
	var slice []string
	for _, ca := range a {
		for _, cb := range b {
			slice = append(slice, string([]rune{ca, cb}))
		}
	}
	return slice
}

const digits = "123456789"
const rows = "ABCDEFGHI"
const cols = digits

var squares = cross(rows, digits)
var unitlist [][]string
var units = make(map[string][][]string)
var peers = make(map[string]map[string]struct{})

func initialize() {
	for _, c := range cols {
		unitlist = append(unitlist, cross(rows, string(c)))
	}
	for _, r := range rows {
		unitlist = append(unitlist, cross(string(r), cols))
	}
	for _, r := range []string{"ABC", "DEF", "GHI"} {
		for _, c := range []string{"123", "456", "789"} {
			unitlist = append(unitlist, cross(r, c))
		}
	}

	for _, unit := range unitlist {
		for _, s := range unit {
			units[s] = append(units[s], unit)
		}
	}

	for _, square := range squares {
		for _, unit := range units[square] {
			for _, peer := range unit {
				if peer == square {
					continue
				}
				if _, ok := peers[square]; !ok {
					peers[square] = make(map[string]struct{})
				}
				peers[square][peer] = struct{}{}
			}
		}
	}
}

func parseGrid(grid string) (map[string]string, error) {
	values := make(map[string]string)
	for _, s := range squares {
		values[s] = digits
	}

	for s, d := range gridValues(grid) {
		if d != "0" && d != "." {
			if err := assign(values, s, d); err != nil {
				return nil, err
			}
		}
	}
	return values, nil
}

func gridValues(grid string) map[string]string {
	chars := make(map[string]string)
	for i, c := range grid {
		chars[squares[i]] = string(c)
	}
	return chars
}

func assign(values map[string]string, s string, d string) error {
	otherValues := strings.Replace(values[s], d, "", 1)
	for _, d2 := range otherValues {
		err := eliminate(values, s, string(d2))
		if err != nil {
			return err
		}
	}
	return nil
}

func eliminate(values map[string]string, s string, d string) error {
	// TODO: 実装する
	if strings.Index(values[s], d) == -1 {
		return nil
	}
	values[s] = strings.Replace(values[s], d, "", 1)

	// (1) If a square s is reduced to one value d2, then eliminate d2 from the peers.
	if len(values[s]) == 0 {
		return errors.New("contradiction detected: some squares cannot be assigned to any number")
	} else if len(values[s]) == 1 {
		d2 := values[s]
		for s2, _ := range peers[s] {
			if err := eliminate(values, s2, d2); err != nil {
				return err
			}
		}
	}

	// (2) If a unit u is reduced to only one place for a value d, then put it there.
	for _, u := range units[s] {
		var dplaces []string
		for _, s2 := range u {
			if strings.Index(values[s2], d) != -1 {
				dplaces = append(dplaces, s2)
			}
		}
		if len(dplaces) == 0 {
			return errors.New("contradiction detected: number cannot be assigned to any square in a unit exists")
		} else if len(dplaces) == 1 {
			if err := assign(values, dplaces[0], d); err != nil {
				return err
			}
		}
	}

	return nil
}

func display(values map[string]string) {
	width := 1
	for _, square := range squares {
		if len(values[square])+1 > width {
			width = len(values[square]) + 1
		}
	}
	bar := []string{strings.Repeat("-", width*3)}
	bar = append(bar, bar[0], bar[0])
	line := strings.Join(bar, "+")
	for i, s := range squares {
		fmt.Printf("%*s", width, values[s])
		if i%9 == 8 {
			fmt.Print("\n")
		} else if i%3 == 2 {
			fmt.Print("|")
		}
		if i%27 == 26 && i%81 != 80 {
			fmt.Printf("%s\n", line)
		}
	}
}

func main() {
	initialize()
	_ = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	grid2 := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"

	values, err := parseGrid(grid2)
	if err != nil {
		fmt.Println("error", err)
	}

	display(values)
}
