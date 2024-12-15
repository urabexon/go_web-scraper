package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	doc, err := goquery.NewDocument("URL")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		download(url)
	})
	t = time.Now()
	fmt.Println(t)
}

func download(url string) (err error) {
	filename := path.Base(url)
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error!: downloading %s: %v\n", url, err)
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error!: creating file %s: %v\n", filename, err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Printf("Error!: copying data to file %s: %v\n", filename, err)
		return err
	}
}	