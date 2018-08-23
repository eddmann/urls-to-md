package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"net/url"
	"strings"
)

func isValidUri(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	return err == nil
}

func toUrlList(input string) []string {
	list := strings.Split(strings.TrimSpace(input), "\n")
	urls := make([]string, 0)

	for _, url := range list {
		if isValidUri(url) {
			urls = append(urls, url)
		}
	}

	return urls
}

type UrlTitle struct {
	idx   int
	url   string
	title string
}

func fetchUrlTitles(urls []string) []*UrlTitle {
	ch := make(chan *UrlTitle, len(urls))

	for idx, url := range urls {
		go func(idx int, url string) {
			doc, err := goquery.NewDocument(url)

			if err != nil {
				ch <- &UrlTitle{idx, url, ""}
			} else {
				ch <- &UrlTitle{idx, url, doc.Find("title").Text()}
			}
		}(idx, url)
	}

	urlsWithTitles := make([]*UrlTitle, len(urls))

	for _ = range urls {
		urlWithTitle := <-ch
		urlsWithTitles[urlWithTitle.idx] = urlWithTitle
	}

	return urlsWithTitles
}

func toMarkdownList(urlsWithTitles []*UrlTitle) string {
	markdown := ""

	for _, urlWithTitle := range urlsWithTitles {
		markdown += fmt.Sprintf("- [%s](%s)\n", urlWithTitle.title, urlWithTitle.url)
	}

	return strings.TrimSpace(markdown)
}

func main() {
	input, _ := clipboard.ReadAll()

	urls := toUrlList(input)

	if len(urls) == 0 {
		fmt.Println("No URLs found in clipboard.")
		return
	}

	urlsWithTitles := fetchUrlTitles(urls)

	markdown := toMarkdownList(urlsWithTitles)

	fmt.Println(markdown)

	clipboard.WriteAll(markdown)
}
