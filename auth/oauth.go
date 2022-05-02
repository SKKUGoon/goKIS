package auth

type kisHeader struct {
	ContentType string `json:"content-type"`
	AppKey      string `json:"appkey"`
	AppSecret   string `json:"appsecret"`
}

type kisPost struct {
}

type kisResponse struct {
	JsonBody interface{}
	HASH     string
}
