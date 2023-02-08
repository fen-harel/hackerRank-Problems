package main

import "fmt"

func arr(max int) {
	var lim int = max + max - 1
	var origin int = max - 1
	var val int = 1

	// var myVec = make([]int, lim)
	var myVec_2d [][]int

	// var sliced_Vec = myVec_2d[0:2]

	for t := 0; t < lim; t++ {
		var myVec = make([]int, lim)
		myVec_2d = append(myVec_2d, myVec)
	}
	fmt.Printf("%d\n", myVec_2d[origin])

	myVec_2d[origin][origin] = val
	fmt.Printf("b4 for() --> %d\n", myVec_2d[origin])

	for i := 0; i < max; i++ {

		// left
		myVec_2d[origin][origin-i] = val + i
		// top-left
		myVec_2d[origin-i][origin-i] = val + i
		// top
		myVec_2d[origin-i][origin] = val + i
		// top-right
		myVec_2d[origin-i][origin+i] = val + i
		// right
		myVec_2d[origin][origin+i] = val + i
		// bot-right
		myVec_2d[origin+i][origin+i] = val + i
		// bot
		myVec_2d[origin+i][origin] = val + i
		// bo-left
		myVec_2d[origin+i][origin-i] = val + i
	}

	for p := 0; p < lim; p++ {
		fmt.Printf("--> %d\n", myVec_2d[p])
	}

}

func main() {
	arr(6)
}
