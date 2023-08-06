package main

import "strings"

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
	return nil
}

func main() {
	initialize()
}
