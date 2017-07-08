package go_koans

func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; i < 100; /* infinite loop */ i++ {
		// your code goes here

		assert(i < 100) // i is afraid of heights
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	go func() {

		ch <- 2
		ch <- 3
		ch <- 5
		ch <- 7
		ch <- 11
	}()
	// concurrency can be almost trivial
	// your code goes here

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
