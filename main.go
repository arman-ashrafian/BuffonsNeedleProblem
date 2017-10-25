// Buffon's Needle

package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func main() {
    const (
        ITERATIONS int = 1000
        LINES      int = 7
    )

    var (
        randomAngleRads float64
        randomX int
        x2 int
    )

    // Seed random number generator
    rand.Seed(time.Now().UnixNano())

    // x-pos of lines
    lines_x := make([]int, LINES)

    // populate lines_x
    for i := 0; i < len(lines_x); i++ {
        lines_x[i] = i*10
    }

    hits := 0
    loops := 0
    for i := 0; i <= ITERATIONS; i++ {
        randomAngleRads = float64(rand.Intn(360) * (math.Pi/180))

        // end point 1
        randomX = rand.Intn(600)
        // end point 2
        x2 = int(randomX + (100*math.Cos(randomAngleRads)))

        // check if needle hits any lines
        for _, line_x := range lines_x {
            if x2 > line_x && randomX < line_x {
                hits++
            } else if x2 < line_x && randomX > line_x{
                hits++
            }
        }
    }

    prob := hits/ITERATIONS

    fmt.Println(prob)


}
