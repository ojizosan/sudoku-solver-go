package main

func cross(a string, b string) []string {
	var hoge []string
	for _, ca := range a {
		for _, cb := range b {
			hoge = append(hoge, string([]rune{ca, cb}))
		}
	}
	return hoge
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

func main() {
	initialize()
}
