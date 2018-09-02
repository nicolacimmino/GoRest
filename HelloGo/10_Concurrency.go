package HelloGo

import (
	"fmt"
	"time"
)

func Chapter_10_Concurrency() string {

	if ChapterIndex() &&
		Chapter_10_01_GoRoutines() &&
		Chapter_10_02_Channels() &&
		Chapter_10_03_BufferedChannels() &&
		Chapter_10_03_MultipleChannelsAndTimeouts() {
		return "L"
	}

	panic("OMG! Something is wrong in chapter 10!")
}

func Chapter_10_01_GoRoutines() bool {
	aValue := 0

	// Just a regular function that takes 2 seconds to execute.
	aSlowFunction := func() {
		time.Sleep(2 * time.Second)
		aValue++
	}

	startTime := time.Now()

	// You can run a go routine, sort of light weight thread, by invoking
	// a function with "go". These won't block and will execute concurrently.
	go aSlowFunction()
	go aSlowFunction()
	go aSlowFunction()

	// We can prove this by checking that it took less than a second to
	// got through the above calls.
	if duration := time.Since(startTime); duration.Seconds() > 1 {
		return false
	}

	// Also at this point none of the slow functions has executed the
	// line incrementing aValue. In fact, if the program ends here
	// they will never have a chance to execute.
	if aValue > 0 {
		return false
	}

	// You don't need to declare a function to start a go routine,
	// it can be started with an anonymous function.
	go func() {
		time.Sleep(2 * time.Second)
	}()

	// Again We can prove the above call didn't block.
	if duration := time.Since(startTime); duration.Seconds() > 1 {
		return false
	}

	return true
}

func Chapter_10_02_Channels() bool {
	// A channel is a pipe through which you can send messages
	// to a go routine. A channel has a specific type.
	aChannel := make(chan int)

	// Sending and receiving from a channel are blocking operations
	// so we need to do this in a different go routine.
	// Also this means channels are a way to synchronise go routines.
	go func() {
		time.Sleep(1 * time.Second)
		// Send an int into the channel, this will block
		// till the value is read.
		aChannel <- 1
	}()

	startTime := time.Now()

	// This will block till a value is received from aChannel.
	if <-aChannel != 1 {
		return false
	}

	// We can prove this cause the above if statement took more than 1 seconds.
	if duration := time.Since(startTime); duration.Seconds() < 1 {
		return false
	}

	return true
}

func Chapter_10_03_BufferedChannels() bool {
	// A buffered channel doesn't block when writing unless it's full
	// and doesn't block on reading unless it's empty.
	aBufferedChannel := make(chan int, 10)

	// Send 5 values down the channel and then close it.
	go func() {
		aBufferedChannel <- 1
		aBufferedChannel <- 2

		// You can signal to the channel you are done by closing the channel.
		// This does not change immediately the state of the channel, think it
		// more like a "bye" message that is queued behind the the others.
		close(aBufferedChannel)
	}()

	// Ensure the above go routine is done so we can prove the channel
	// close is just a message in the queue.
	time.Sleep(1 * time.Second)

	// We can get a value and a bool indicating whether the channel is open.
	aValue, channelOpen := <-aBufferedChannel

	// In this case it's still open, even though the go routine above has
	// already called close(), as we have pulled only the first value.
	if !channelOpen || aValue != 1 {
		return false
	}

	// This is the second value, after which the channel was closed.
	// Here we still get channelOpen true.
	aValue, channelOpen = <-aBufferedChannel

	// In this case we expect the channel to be still open.
	if !channelOpen || aValue != 2 {
		return false
	}

	// Now the channel is closed as the last value was read. This will
	// return a 0 value and false channelOpen.
	aValue, channelOpen = <-aBufferedChannel

	// In this case we expect the channel to be closed.
	if channelOpen || aValue != 0 {
		return false
	}

	return true
}

func Chapter_10_03_MultipleChannelsAndTimeouts() bool {

	newsChannel := make(chan string)
	weatherChannel := make(chan string)

	go func() {
		// Fetch news from a very slow server
		time.Sleep(2 * time.Second)
		newsChannel <- "Something new happened today."
	}()

	go func() {
		// Fetch weather from a very slow server
		time.Sleep(2 * time.Second)
		weatherChannel <- "Extremely hot weather with scattered snow blizzards."
	}()

	// All 3 calls in the case:s will block and the taken action will depend on
	// the one that will unblock first. Clearly only one of the cases will be executed.
	select {
	case news := <-newsChannel:
		fmt.Println("Something is wrong, news came back too fast: " + news)
		return false
	case weather := <-weatherChannel:
		fmt.Println("Something is wrong, weather came back too fast: " + weather)
		return false
	case <-time.After(100 * time.Millisecond):
	}

	return true
}
