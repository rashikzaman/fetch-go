package scrapper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"url-saver/util"
)

func GetHtmlFromUrl(url string) error {
	htmlDataBody, err := GetHtml(url)
	if err != nil {
		fmt.Println("error getting html", err)
		return err
	}
	urlPath := fmt.Sprintf("./%s.html", url)
	err2 := util.StoreFile(urlPath, htmlDataBody)
	if err2 != nil {
		fmt.Println("error storing html", err2)
		return err
	}
	return nil
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
