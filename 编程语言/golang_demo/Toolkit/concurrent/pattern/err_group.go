package pattern

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// ErrGroup 错误组模式
func errGroup() {
	var g errgroup.Group
	urls := []string{
		"https://golang.org", "https://google.com", "https://badhost",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			fmt.Printf("Fetch %s\n", url)

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("Successfully fetched all URLs")
	}
}
