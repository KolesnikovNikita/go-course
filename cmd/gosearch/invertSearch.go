// package main

// import (
// 	"flag"
// 	"fmt"
// 	"gosearch/pkg/crawler/spider"
// 	"sort"
// 	"strings"
// 	"unicode"
// )

// type InvertedIndex map[string][]string

// func main() {

// 	var urlList []string

// 	url_1 := "https://go.dev"
// 	url_2 := "https://golang.org"
// 	scanner := spider.New()

// 	data_1, err := scanner.Scan(url_1, 2)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	data_2, err := scanner.Scan(url_2, 2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	summary := append(data_1, data_2...)

// 	for _, v := range summary {
// 		urlList = append(urlList, v.URL)
// 	}
		
// 			res := CreateInvertedIndex(urlList)
// 			sortedStruct := sortInvertIndex(res)
			

// 			searchWord := flag.String("s", "", "--help")
// 	    flag.Parse()

// 		result := SearchInvertedIndex(sortedStruct, *searchWord)

// 	  for _, value := range result {
// 		fmt.Println(value)
// 	  }
// 	}
			
	
// 	func CreateInvertedIndex(links []string) InvertedIndex {
// 		 index := make(InvertedIndex)

// 		for _, link := range links {
// 			terms := extractTerms(link)

// 			for _, term := range terms {
// 				index[term] = append(index[term], link)
// 			}
// 		}
		
// 		return index

// 	}

// 	func extractTerms(link string) []string {
// 		terms := strings.FieldsFunc(link, func(r rune)bool{
// 			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// 		})

// 		return terms
// 	}
 
// 	func sortInvertIndex(index InvertedIndex) InvertedIndex {
// 		var sortedKeys []string

// 		for term := range index {
// 			sortedKeys = append(sortedKeys, term)
// 		}

// 		sort.Strings(sortedKeys)

// 		sortedIndex := make(InvertedIndex)

// 		for _, term := range sortedKeys {
// 			sortedIndex[term] = index[term]
// 		}

// 		return sortedIndex
//  	}

// 	func SearchInvertedIndex(index InvertedIndex, word string) []string {
//     links, found := index[word]
//     if found {
//         return links
//     }
//     return nil
// }