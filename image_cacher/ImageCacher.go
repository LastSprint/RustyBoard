package image_cacher

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AsyncImageCacher struct {
	PathToFolderWithImages string
	UrlPathToImages        string
	// creds for service which stores images
	User string
	Pass string
}

func (c *AsyncImageCacher) Cache(url string) string {
	md5hash := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	resultPath := c.PathToFolderWithImages + "/" + md5hash + ".png"

	resultUrl := c.UrlPathToImages + "/" + md5hash + ".png"

	_, err := os.Stat(resultPath)

	if os.IsNotExist(err) {
		f, err := os.OpenFile(resultPath, os.O_CREATE|os.O_RDWR, os.ModePerm)

		if err != nil {
			log.Println("[ERR] Cant open file for", resultPath, err.Error())
			return resultUrl
		}

		defer func() {
			if err := f.Close(); err != nil {
				log.Println("[WARN] Cant close file at path", resultPath, err.Error())
			}
		}()

		data, err := c.load(url)

		if err != nil {
			log.Println("[ERR] error while loading image", url, err.Error())
		}

		f.Write(data)
	}

	return resultUrl
}

func (c *AsyncImageCacher) load(url string) ([]byte, error) {
	rq, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	rq.SetBasicAuth(c.User, c.Pass)

	data, err := http.DefaultClient.Do(rq)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := data.Body.Close(); err != nil {
			log.Println("[WARN] Couldn't close response body in image loading", err.Error())
		}
	}()

	res, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return nil, err
	}

	return res, nil
}
