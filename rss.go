package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

var rss string
var rssObject Rss2

type Rss2 struct {
	Version     string `xml:"version,attr"`
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	PubDate     string `xml:"channel>pubDate"`
	ItemList    []Item `xml:"channel>item"`
}

type newsList []Item

func (p newsList) Len() int {
	return len(p)
}

func (p newsList) Less(i, j int) bool {
	date1, _ := time.Parse(time.RFC1123, p[i].PubDate)
	date2, _ := time.Parse(time.RFC1123, p[j].PubDate)
	return date1.After(date2)
}

func (p newsList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	Content     template.HTML `xml:"encoded"`
	PubDate     string        `xml:"pubDate"`
}

func fetchFeed() {
	response, err := http.Get("http://feeds.bbci.co.uk/news/business/rss.xml?edition=uk")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		v := Rss2{}
		err2 := xml.Unmarshal(contents, &v)
		if err2 != nil {
			log.Println(err2)
		}
		changeTimeStamps(&v.ItemList)
		json, err3 := json.Marshal(v)
		if err3 != nil {
			fmt.Println("error:", err)
		}
		rss = string(json)
		rssObject = v
	}
}

func changeTimeStamps(itemList *[]Item) {
	sort.Sort(newsList(*itemList))
	for i := 0; i < len(*itemList); i++ {
		date, _ := time.Parse(time.RFC1123, (*itemList)[i].PubDate)
		(*itemList)[i].PubDate = date.Format("2.1.2006 15:04")
	}
}
