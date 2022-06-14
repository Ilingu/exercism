package armstrong

import (
	"math"
	"strconv"
	"strings"
	"sync"
)

/* WITHOUT CONCURRENCY --> more performent
func IsNumber(n int) bool {
	strNum := strconv.Itoa(n)
	digits := strings.Split(strNum, "")

	var sum int
	for _, digit := range digits {
		digitNum, _ := strconv.Atoi(digit)
		sum += int(math.Pow(float64(digitNum), float64(len(digits))))
	}
	return sum == n
}
*/

/* WITH CONCURRENCY --> ~1s faster */
var sumMut sync.Mutex
var wg sync.WaitGroup

func Add(sum *int, digit string, length float64) {
	defer wg.Done()
	digitNum, _ := strconv.Atoi(digit)
	sumMut.Lock()
	*sum += int(math.Pow(float64(digitNum), length))
	sumMut.Unlock()
}

func IsNumber(n int) bool {
	strNum := strconv.Itoa(n)
	digits := strings.Split(strNum, "")

	var sum, length = 0, len(digits)

	wg.Add(length)
	for _, digit := range digits {
		go Add(&sum, digit, float64(length))
	}
	wg.Wait()

	return sum == n
}
