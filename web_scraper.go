package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	doc, err := goquery.NewDocument("URLを指定")
	if err != nil {
		fmt.Println("Error!: URL scraping failed:", err)
		return
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if !exists {
			return
		}
		// 相対URLの場合は絶対URLに変換(パス変換)
		if !strings.HasPrefix(url, "http") {
			url = "URLを指定" + url
		}
		err := download(url)
		if err != nil {
			fmt.Printf("Error!: Failed to download %s: %v\n", url, err)
		}
	})
	t = time.Now()
	fmt.Println("End time", t)
}

func download(url string) (err error) {
	// ファイル名取得
	filename := path.Base(url)
	fmt.Println("Downloading ", url, " to ", filename)

	// HTTPリクエストでファイル取得
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error!: downloading %s: %v\n", url, err)
		return err
	}
	defer resp.Body.Close()

	// ローカルにファイル生成する
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error!: creating file %s: %v\n", filename, err)
		return err
	}
	defer f.Close()

	// データをファイルにコピーする
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Printf("Error!: copying data to file %s: %v\n", filename, err)
		return err
	}
	return nil
}