package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"strings"
)

func main() {

	var urlList []crawler.Document

	urls := []string{
		"https://go.dev",
		"https://golang.org",
	}

	scanner := spider.New()

	for _, url := range urls {
		data, err := scanner.Scan(url, 2)
		if err != nil {
			fmt.Print(err)
		}
		urlList = append(urlList, data...)
	}




	searchWord := flag.String("s", "", "--help")
	flag.Parse()

	if *searchWord == "" {
		fmt.Println("Не указано слово для поиска")
		return
	} else {
		match := checkUrlsForWord(urlList, *searchWord)
		if len(match) == 0 {
			fmt.Println("Совпадений не найдено")
		}
	for _, value := range match {
		fmt.Printf("%s\n", value)
	  }
	}

	}

			
	func checkUrlsForWord(arr []crawler.Document, w string) []string  {
		var matchList []string
		for _, v := range arr {
			if strings.Contains(v.URL, w) {
				matchList = append(matchList, v.URL)
			}
		}
		
		return matchList
	}