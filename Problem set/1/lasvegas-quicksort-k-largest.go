package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	baseCase = 1
)

func main() {
	var list = []int{0, 4, 1, -2, 30, 12, 43, -123, 0, 321, 36, 65}
	k, _ := strconv.Atoi(os.Args[1])

	qsVegas(list[:], len(list)-1-k)
}

func qsVegas(list []int, k int) {
	quicksort(list, 0, len(list)-1, k)
}

func quicksort(list []int, lower int, upper int, k int) {
	if len(list) == baseCase || lower >= upper {
		print(list, list[lower])
		return
	}

	piv, interrupt := pivot(list[:], lower, upper, k)

	if interrupt {
		return
	}

	if piv > k {
		quicksort(list[:], lower, piv-1, k)
	} else {
		quicksort(list[:], piv+1, upper, k)
	}
}

func pivot(list []int, lower int, upper int, k int) (int, bool) {
	randomIndex := lower + rand.Intn(upper-lower)
	lasVegas := list[randomIndex]
	down := lower
	up := upper

	for down < up {
		for (list[down] < lasVegas && down < up) || down == randomIndex {
			down++
		}
		for (list[up] >= lasVegas && up > down) || up == randomIndex {
			up--
		}

		swap(list[:], down, up)
	}

	// Off by one
	if down > 0 && list[down] > lasVegas && list[down-1] >= lasVegas {
		down--
	}

	swap(list[:], randomIndex, down)

	if down == k {
		print(list, list[down])
		return down, true
	}

	return down, false
}

func print(list []int, element int) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Println("Element found: ", element)
}

func swap(list []int, a int, b int) {
	tmp := list[a]

	list[a] = list[b]
	list[b] = tmp
}
