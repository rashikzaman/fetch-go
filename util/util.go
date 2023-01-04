package util

import "os"

func AddHttpToUrlString(url string) string {
	if len(url) > 4 {
		if url[:4] != "http" {
			return "http://" + url
		}
	}
	return url
}

func StoreFile(path string, body []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err2 := f.Write(body)
	return err2
}
