package main 

import (
	"fmt"
	"net/http"
	"log"
	"regexp"
	"io/ioutil"
	"time"
)

var cleanUrlList []string

func urlGetter() {
	resp, err := http.Get("https://golang.org/pkg/")
	if err != nil {
		log.Fatal("http.Get: %v", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll 2: %v", err)
	}


	re := regexp.MustCompile("(href=\")[a-z]+/[a-z]*/*")
	re2 := regexp.MustCompile("(href=\")")

	pkgLIstHtml := re.FindAllString(string(bytes), -1)

	for _, res := range pkgLIstHtml {
		clean := re2.ReplaceAllString(res, "$1W")
		cleanUrlList = append(cleanUrlList, clean)
	}

	fmt.Println(pkgLIstHtml)
	fmt.Println(cleanUrlList)

}

func textGetter() {

	for _, partUrl := range cleanUrlList {
		resp, err := http.Get("https://golang.org/pkg/" + partUrl)
		if err != nil {
			log.Fatal("http.Get 3: %v", err)
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("ioutil.ReadAll 4: %v", err)
		}

		fmt.Println(string(bytes))

		time.Sleep(time.Millisecond * 1000)
	}
}


func main() {
	urlGetter()

	textGetter()

}