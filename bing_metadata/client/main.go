package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"metadata/metadata"
	"net/http"
	"net/url"
	"os"
)

func handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}
	fmt.Printf("%d: %s\n", i, url)

	res, err := http.Get(url)

	if err != nil {
		return
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf(
		"%25s %25s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}
func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing required args")
	}
	domain := os.Args[1]
	fileType := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title): %s",
		domain,
		fileType,
		fileType)
	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))

	doc, err := goquery.NewDocument(search)

	if err != nil {
		log.Panic(err)
	}

	s := "html body div#b_content ol#b_result li.b_algo div.b_title h2"
	doc.Find(s).Each(handler)
}
