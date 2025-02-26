package pattern

// Pipeline 管道模式
func pipeline(pipe <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range pipe {
			out <- n * n
		}
		close(out)
	}()
	return out
}
