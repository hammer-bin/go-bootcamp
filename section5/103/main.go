package main

import (
	"fmt"
	"github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	const pageSize = 4

	l := len(items)
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		prettyslice.Show(head, currentPage)
	}

	/*	prettyslice.Show("0:4", items[0:4])
		prettyslice.Show("4:8", items[4:8])
		prettyslice.Show("8:12", items[8:12])
		prettyslice.Show("12:13", items[12:13])*/

}
