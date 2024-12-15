# Go_web-scraper

指定したWebサイトをスクレイピングして、ページ内のすべてのリンク（URL）を取得し、それぞれのリンク先をローカルにダウンロードします。<br>
相対URLも絶対URLに変換し、リンク先のコンテンツをファイルとして保存します。

## Usage
### 必要条件(readlineのインストール)
`readline` (シェルがコマンドを受け付けるときに使用されます。)
```shell
brew install readline
```

## Install 
1. Goをインストールする
	- 1.19以上のバージョンでないと動作しません。
2. 依存パッケージをインストールする
	- [https://github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)

	`goquery`ライブラリを使用しています。依存関係をインストールするため、以下のコマンドを実行します。

	```bash
	go get github.com/PuerkitoBio/goquery
	```
3. ファイル内にスクレイピングしたいWebサイトのURLを指定する
4. 実行する
	```go
	go run web_scraper.go
	```