package pipeline

import "sync"

func LaunchPipeline(amount int) int {
	createListChannel := generator(amount)
	powerChannel := power(createListChannel, amount)
	resultChannel := sum(powerChannel, amount)

	return resultChannel
}

func generator(max int) <-chan int {
	outChInt := make(chan int, max)

	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}

		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int, max int) <-chan int {
	outChInt := make(chan int, max)

	go func() {
		for c := range in {
			outChInt <- c * c
		}

		close(outChInt)
	}()

	return outChInt
}

func sum(in <-chan int, max int) int {
	//sumChannel := make(chan int, max)
	var sum int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		for v := range in {
			println(v)
			sum += v
		}
		//sumChannel <- sum
		//close(sumChannel)
		wg.Done()

	}()

	wg.Wait()
	return sum
}
