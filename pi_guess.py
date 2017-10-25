# Arman Ashrafian
# Estimate PI

import math, random, time

def main():
    # number of sticks to drop
    ITERATIONS = 100000

    # seed random number generator
    random.seed(time.time())

    # x position of lines
    lines_x = [x*100 for x in range(7)]

    hits = 0
    loops = 0
    while loops < ITERATIONS:
        randomAngle = random.randint(0,360)

        # end point 1
        randomX = random.randint(0, 600)
        # end point 2
        x2 = randomX + 100*math.cos(math.radians(randomAngle))

        # check if needle hits any lines
        for x in lines_x:
            if x2 > x and randomX < x:
                hits += 1
            elif x2 < x and randomX > x:
                hits += 1

        loops += 1

    prob = hits/loops

    print("PI: " + str(2/prob))



if __name__ == '__main__':
    main()
