package main

import (
	"fmt"
	"math/rand"
)

const (
	baseCase = 1
)

func main() {
	var list = []int{0, 4, 1, -2, 30, 12, 43, -123, 0, 321, 36, 65}

	qsVegas(list[:])

	for i := 0; i < len(list); i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Println("")
}

func qsVegas(list []int) {
	quicksort(list, 0, len(list)-1)
}

func quicksort(list []int, lower int, upper int) {
	if len(list) <= baseCase || lower >= upper {
		return
	}

	piv := pivot(list[:], lower, upper)

	quicksort(list[:], lower, piv-1)
	quicksort(list[:], piv+1, upper)
}

func pivot(list []int, lower int, upper int) int {
	// Las Vegas index
	randomIndex := lower + rand.Intn(upper-lower)
	down := lower
	up := upper
	lasVegas := list[randomIndex]

	for down < up {
		for (list[down] < lasVegas && down < up) || down == randomIndex {
			down++
		}
		for (list[up] >= lasVegas && up > down) || up == randomIndex {
			up--
		}

		swap(list[:], down, up)
	}

	// Off by one down
	if down > 0 && list[down] > lasVegas && list[down-1] >= lasVegas {
		down--
	}

	swap(list[:], randomIndex, down)
	return down
}

func swap(list []int, a int, b int) {
	tmp := list[a]

	list[a] = list[b]
	list[b] = tmp
}
