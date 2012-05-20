/*
Simple RSS parser, tested with Wordpress feeds.
*/
package rss

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate Date   `xml:"lastBuildDate"`
	Item          []Item `xml:"item"`
}

type ItemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Comments    string        `xml:"comments"`
	PubDate     string          `xml:"pubDate"`
	GUID        string        `xml:"guid"`
	Category    []string      `xml:"category"`
	Enclosure   ItemEnclosure `xml:"enclosure"`
	Description string        `xml:"description"`
	Content     string        `xml:"content"`
}

func Read(url string) (*Channel, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	text, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var rss struct {
		Channel Channel `xml:"channel"`
	}

	err = xml.Unmarshal(text, &rss)

	if err != nil {
		return nil, err
	}

	return &rss.Channel, nil
}
