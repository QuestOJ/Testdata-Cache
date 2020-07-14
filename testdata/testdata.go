package testdata

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/QuestOJ/testdata-cache/typedef"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var dataPath string

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func cacheNotExpire(filename string, id string, OSS typedef.OSSConfig) bool {
	return true
}

func cacheRead(filename string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(filename)
	return fileData, err
}

func cacheExist(id string, fileType string, OSS typedef.OSSConfig) bool {
	switch fileType {
	case "testdata":
		filepath := dataPath + "/testdata/" + id + "/testdata.zip"
		return fileExist(filepath) && cacheNotExpire(filepath, id, OSS)
	default:
		return false
	}
}

func cacheWrite(id string, fileType string, writer http.ResponseWriter) error {
	switch fileType {
	case "testdata":
		res, err := cacheRead(dataPath + "/testdata/" + id + "/testdata.zip")
		if err != nil {
			return err
		}
		writer.Write(res)
		return nil
	default:
		return errors.New("Invaild filetype")
	}
}

func cacheMissed(id string, fileType string, OSS typedef.OSSConfig, writer http.ResponseWriter) error {
	var name string

	switch fileType {
	case "testdata":
		name = "data/" + id + "/testdata.zip"
		break
	default:
		return errors.New("Invaild filetype")
	}

	client, err := oss.New(OSS.EndPoint, OSS.Key, OSS.Secret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(OSS.BucketName)
	if err != nil {
		return err
	}

	body, err := bucket.GetObject(name)
	if err != nil {
		return err
	}
	defer body.Close()

	fd, err := os.OpenFile("data/testdata/"+id+"/testdata.zip", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer fd.Close()

	io.Copy(fd, body)
	return nil
}

func Get(id string, fileType string, datapath string, OSS typedef.OSSConfig, writer http.ResponseWriter) error {
	dataPath = datapath

	os.Mkdir(dataPath+"/testdata/"+id, 0770)

	if cacheExist(id, fileType, OSS) {
		err := cacheWrite(id, fileType, writer)
		if err != nil {
			return err
		}
	} else {
		err := cacheMissed(id, fileType, OSS, writer)
		if err != nil {
			return err
		}
		err = cacheWrite(id, fileType, writer)
		if err != nil {
			return err
		}
	}

	return nil
}
