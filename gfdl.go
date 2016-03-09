package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	regex = regexp.MustCompile(`https?://.+\.\w+`)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Usage: gfdl src [dest]`)
		os.Exit(1)
	}
	src := os.Args[1]
	var dest string
	if len(os.Args) > 2 {
		dest = os.Args[2]
	}
	dir := filepath.Dir(dest)
	err := os.MkdirAll(dir, 0755)
	check(err)

	resp, err := http.Get(src)
	check(err)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	css := string(bytes[:])
	urls := regex.FindAllString(css, -1)
	if urls == nil {
		fmt.Println("No fonts found!")
		return
	}

	var wg sync.WaitGroup
	semas := make(chan string)
	for i := range urls {
		wg.Add(1)
		go func(url string) {
			fmt.Println("fetch ", url)
			defer wg.Done()
			resp, err := http.Get(url)

			defer resp.Body.Close()
			bytes, err = ioutil.ReadAll(resp.Body)
			check(err)

			err = ioutil.WriteFile(filepath.Join(dir, shortName(url)), bytes, 0644)
			check(err)

			semas <- url
		}(urls[i])
	}

	go func() {
		wg.Wait()
		close(semas)
	}()

	for url := range semas {
		short := shortName(url)
		css = strings.Replace(css, url, short, -1)
	}

	err = ioutil.WriteFile(dest, []byte(css), 0644)
	check(err)
}

func shortName(s string) string {
	return s[strings.LastIndex(s, "/")+1:]
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
