package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"net/http"
	"net/url"
)

// IP struct
type IP struct {
	ID    bson.ObjectId `bson:"_id" json:"-"`
	Data  string        `bson:"data" json:"ip"`
	Type  string        `bson:"type" json:"type"`
	Delay int64         `bson:"delay" json:"delay"`
}

// NewIP .
func NewIP() *IP {
	return &IP{
		ID: bson.NewObjectId(),
	}
}

func NewIPAndCheck(ip, schema string) *IP {
	i := &IP{
		ID: bson.NewObjectId(),
		Data: ip,
		Type: schema,
	}

	if i.CheckIP() {
		return i
	} else {
		return nil
	}
}

// CheckIP is to check the ip work or not
func (ip *IP) CheckIP() bool {
	pollURL := "http://httpbin.org/get"
	start := time.Now().Unix()
	resp, err := proxyGet(pollURL, "http://" + ip.Data)
	end := time.Now().Unix()
	if err != nil {
		return false
	}
	ip.Delay = end - start
	if ip.Delay > 5 {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func proxyGet(testUrl string, proxy string) (*http.Response, error) {
	u, _ := url.Parse(testUrl)
	transport := &http.Transport{Proxy: func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxy)
	}}
    client := &http.Client{Transport: transport}
    return client.Get(u.String())
}
