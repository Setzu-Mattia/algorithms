package main

import "fmt"

func main() {
	l := []int{0, -1, 4, 45, 6, -12, 43, -1234, 12, 34}
	mergesort(l)
	fmt.Println("list", l)
}

func mergesort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	breakPoint := len(list) / 2
	merge(mergesort(list[:breakPoint]), mergesort(list[breakPoint:]))

	return list
}


func merge(bottom, top []int) {
	switch len(bottom) + len(top) {
		case 0:
		case 1:
		case 2:
			if len(bottom) == 1 {
				if bottom[0] > top[0] {
					tmp := top[0]
					top[0] = bottom[0]
					bottom[0] = tmp
				}
				return
			}

			if len(bottom) == 2 && bottom[0] > bottom[1] {
				tmp := bottom[1]
				bottom[1] = bottom[0]
				bottom[0] = tmp
				return
			}

			if len(top) == 2 && top[0] > top[1] {
				tmp := top[1]
				top[1]= top[0]
				top[0] = tmp
				return
			}
		default:
			bottomLen := len(bottom)
			topLen := len(top)
			bottomI, topI := 0, 0

			for bottomI < bottomLen && topI < topLen {
				if bottom[bottomI] > top[topI] {
					tmp := top[topI]
					top[topI] = bottom[bottomI]
					bottom[bottomI] = tmp

					bottomI++
				} else {
					topI++
				}
			}

			if topI < topLen {
				for ;topI < topLen - 1; topI++ {
					if top[topI] >  top[topI + 1] {
						tmp := top[topI + 1]
						top[topI + 1] = top[topI]
						top[topI] = tmp
					}
				}
			}

			if bottomI < bottomLen {
				for ;bottomI < bottomLen - 1; bottomI++ {
					if bottom[topI] >  top[topI + 1] {
						tmp := bottom[topI + 1]
						bottom[topI + 1] = top[topI]
						bottom[topI] = tmp
					}
				}
			}
	}
}
