package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
)

// https://judge.beecrowd.com/pt/problems/view/1168
func main() {
        scanner := bufio.NewScanner(os.Stdin)

        scanner.Scan()
        lines := scanner.Text()

        linesNumber, _ := strconv.Atoi(lines)

        m := make(map[rune]int)

        m['1'] = 2
        m['2'] = 5
        m['3'] = 5
        m['4'] = 4
        m['5'] = 5
        m['6'] = 6
        m['7'] = 3
        m['8'] = 7
        m['9'] = 6
        m['0'] = 6

        for i := 0; i < linesNumber; i++ {
                scanner.Scan()
                number := scanner.Text()

                nums := []rune(number)

                sum := 0
                for _, v := range nums {

                        sum += m[v]
                }
                fmt.Printf("%d leds\n", sum)
        }
}
