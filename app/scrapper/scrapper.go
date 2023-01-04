package scrapper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"url-saver/util"

	"github.com/PuerkitoBio/goquery"
)

type MetaData struct {
	Site      string
	NumLinks  int
	Images    int
	LastFetch time.Time
}

func GetHtmlFromUrl(url string, includeMetadata bool) (*MetaData, error) {
	htmlDataBody, err := GetHtml(url)
	if err != nil {
		fmt.Println("error getting html", err)
		return nil, err
	}
	urlPath := fmt.Sprintf("./%s.html", url)
	fmt.Println("size", len(htmlDataBody))
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
		return metadata, nil
	}
	return nil, nil
}

func GetHtml(url string) ([]byte, error) {
	client := &http.Client{}
	url = util.AddHttpToUrlString(url)
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
