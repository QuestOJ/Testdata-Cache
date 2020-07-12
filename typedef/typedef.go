package typedef

import "net/http"

type OSSConfig struct {
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	BucketName string `json:"bucket-name"`
	EndPoint   string `json:"endpoint"`
}

type Config struct {
	Server       string    `json:"server"`
	ClientID     string    `json:"id"`
	ClientSecret string    `json:"secret"`
	OSS          OSSConfig `json:"oss"`
}

//Request describes the http request settings
type Request struct {
	URL         string
	Method      string
	Header      map[string]string
	Data        map[string]string
	NotRedirect bool

	Cookie []*http.Cookie
}

//The Response contains the response of the http request
type Response struct {
	ResponseBody   []byte
	RedirectStatus bool
}
