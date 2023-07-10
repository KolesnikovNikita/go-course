package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler/spider"
	"strings"
)

func main() {
	var urlList []string



	url_1 := "https://go.dev"
	url_2 := "https://golang.org"
	scanner := spider.New()

	data_1, err := scanner.Scan(url_1, 2)
	if err != nil {
		panic(err)
	}
	
	data_2, err := scanner.Scan(url_2, 2)
	if err != nil {
		panic(err)
	}

	result := append(data_1, data_2...)

	for _, v := range result {
		urlList = append(urlList, v.URL)
	}


	searchWord := flag.String("s", "", "--help")
	flag.Parse()

	if *searchWord == "" {
		fmt.Println("Не указано слово для поиска!")
		return
	} else {
		match := checkUrlsForWord(urlList, *searchWord)
		if len(match) == 0 {
			fmt.Println("Совпадений не найдено")
		}
	for _, value := range match {
		fmt.Println(value)
	  }
	}

	}

			
	func checkUrlsForWord(arr []string, w string) []string  {
		var matchList []string
		for _, v := range arr {
			if strings.Contains(v, w) {
				matchList = append(matchList, v)
			}
		}
		
		return matchList
	}