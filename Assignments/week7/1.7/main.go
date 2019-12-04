// 055-copy-to-stdout
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		switch {
		case !strings.HasPrefix(url, "http://"):
			url = "https://" + strings.TrimPrefix(url, "http://")
		case !strings.HasPrefix(url, "http://"):
			url = "https://" + url

		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		fmt.Println(n)
		if err != nil {
			err := fmt.Errorf("could not copy using io.copy - here the error %w", err)
			log.Panic(err)

		}

		err = resp.Body.Close()
		if err != nil {
			err := fmt.Errorf("fetch: reading %s: %w", url, err)
			log.Panic(err)
		}
	}
}
