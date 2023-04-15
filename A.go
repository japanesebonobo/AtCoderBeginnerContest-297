package main

import (
	"fmt"
)

func main() {
	var N, D int
	fmt.Scan(&N, &D)
	var T = make([]int, N)
	var ans int = -1
	for i := 0; i < N; i++ {
		fmt.Scan(&T[i])
		if (i > 0) && (T[i]-T[i-1] <= D) {
			ans = T[i]
			break
		}
	}
	fmt.Println(ans)
}
