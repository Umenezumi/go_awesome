// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var httpPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//b, err := ioutil.ReadAll(resp.Body)
		written, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()

		fmt.Println("\n-------------------------------")
		fmt.Printf("resp code: %d", resp.StatusCode)
		runes := []rune(strconv.FormatInt(written, 10))[0:30]
		fmt.Printf("\n%s", string(runes))
	}
}
