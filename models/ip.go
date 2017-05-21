package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/parnurzeal/gorequest"
	"time"
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
	resp, _, errs := gorequest.New().Proxy("http://" + ip.Data).Get(pollURL).End()
	end := time.Now().Unix()
	if errs != nil {
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
