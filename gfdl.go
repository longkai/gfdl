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

func fetchBytes(src string) ([]byte, error) {
	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func fetchCss(src string) (string, error) {
	bytes, err := fetchBytes(src)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func writeFile(src, dest string, mode os.FileMode) error {
	fmt.Println("fetch ", src)
	bytes, err := fetchBytes(src)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(dest, bytes, mode)
}

func shortName(s string) string {
	return s[strings.LastIndex(s, "/")+1:]
}

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
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("MkdirAll fail, %v\n", err)
	}

	css, err := fetchCss(src)
	if err != nil {
		log.Fatalf("fetch css fail, %v\n", err)
	}

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
			defer wg.Done()

			if err := writeFile(url, filepath.Join(dir, shortName(url)), 0644); err != nil {
				log.Fatalf("write font fail, %v\n", err)
			}

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

	if err := ioutil.WriteFile(dest, []byte(css), 0644); err != nil {
		log.Fatalf("fail to write dest css file, %v\n", err)
	}
	fmt.Println("done :)")
}
