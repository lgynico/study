package main

func main() {
	ch := make(chan []string)
	s := []string{"nico"}
	go func() {
		ch <- s
	}()
}
