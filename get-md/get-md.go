package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	goconfig "github.com/iglin/go-config"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	config := goconfig.NewConfig("./get-md/config.json", goconfig.Json)

	url := fmt.Sprint(config.GetProp("url"))

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	converter := md.NewConverter("", true, nil)

	markdown := converter.Convert(doc.Find(".post-full-header, .post-full-author-header, .post-full-image,  .post-and-sidebar .post-content "))
	if err != nil {
		log.Fatal(err)
	}

	// create the file
	split := strings.Split(url, "/")
	var fileName string
	if strings.Trim(split[len(split)-1], " ") != "" {
		fileName = split[len(split)-1]
	} else {
		fileName = split[len(split)-2]
	}

	//title := doc.Find(".post-full-title").First().Text()
	f, err := os.Create("./get-md/" + fileName + ".md")
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	//write directly into file
	_, err = f.WriteString(markdown)
	if err != nil {
		log.Fatal(err)
	}
}
