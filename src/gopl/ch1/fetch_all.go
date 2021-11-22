// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := os.Args[1:]

	for _, url := range urls {
		// twice
		for i := 0; i < 2; i++ {
			go fetch(url, ch)
		}
	}

	for i := 0; i <= len(urls)*2; i++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	parse, _ := url2.Parse(url)
	filePath := fmt.Sprintf("/Users/rat/Downloads/%s-%d.txt", parse.Hostname(), time.Now())
	out, err := os.Create(filePath)
	if err != nil {
		panic(out)
	}

	nbytes, err := io.Copy(out, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
