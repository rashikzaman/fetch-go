package scrapper

import (
	"bytes"
	"fetch/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type MetaData struct {
	Site      string
	NumLinks  int
	Images    int
	LastFetch time.Time
}

func GetHtmlFromUrl(inputUrl string, includeMetadata bool) (*MetaData, error) {
	inputUrl = util.AddHttpToUrlString(inputUrl)
	url, err := url.Parse(inputUrl)
	if err != nil {
		fmt.Println("error parsing url", err)
	}

	htmlDataBody, err := GetHtml(inputUrl)
	if err != nil {
		fmt.Println("error getting html", err)
		return nil, err
	}

	urlPath := fmt.Sprintf("./%s.html", url.Hostname())
	err2 := util.StoreFile(urlPath, htmlDataBody)
	if err2 != nil {
		fmt.Println("error storing html", err2)
		return nil, err
	}

	if includeMetadata {
		metadata, err := ParseHtml(htmlDataBody)
		if err != nil {
			fmt.Println("error parsing html", err)
			return nil, err
		}
		metadata.Site = inputUrl
		return metadata, nil
	}
	return nil, nil
}

func GetHtml(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ParseHtml(data []byte) (*MetaData, error) {
	reader := bytes.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("cannnot parse html data, error: ", err)
		return nil, err
	}

	linkLength := doc.Find("a").Length()
	imagesLength := doc.Find("img").Length()
	result := &MetaData{
		NumLinks:  linkLength,
		Images:    imagesLength,
		LastFetch: time.Now(),
	}
	return result, nil
}
