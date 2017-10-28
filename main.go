// Buffon's Needle

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	const (
		ITERATIONS int = 1000
		LINES      int = 7
	)

	var (
		randomAngleRads float64
		randomX         int
		x2              int
	)

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// x-pos of lines
	lines_x := make([]int, LINES)

	// populate lines_x
	for i := 0; i < len(lines_x); i++ {
		lines_x[i] = i * 100
	}


	var hits int = 0
	for i := 0; i <= ITERATIONS; i++ {
		randomAngleRads = (rand.Float64() * 360) * float64(math.Pi / 180)

		// end point 1
		randomX = rand.Intn(600)
		// end point 2
		x2 = randomX + int(100 * math.Cos(randomAngleRads))

		// check if needle hits any lines
		for _, line_x := range lines_x {
			if x2 > line_x && randomX < line_x {
                fmt.Println("HIT")
				hits++
			} else if x2 < line_x && randomX > line_x {
                fmt.Println("HIT")
				hits++
			} else {
                fmt.Println("NO HIT")
            }
		}
	}

	fmt.Println("HITS: ", hits)
    fmt.Println("ITERATIONS: ", ITERATIONS)

	// fmt.Println(prob)

}
