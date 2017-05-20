package getter

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/freewh/ProxyPool/models"
	"github.com/parnurzeal/gorequest"
)

//PLP get ip from proxylistplus.com
func PLP() (result []*models.IP) {
	pollURL := "https://list.proxylistplus.com/"
	_, body, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("#page > table.bg > tbody > tr").Each(func(i int, s *goquery.Selection) {
		node := strconv.Itoa(i + 1)
		ss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(2)").Html()
		sss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(3)").Html()
		ssss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(7)").Html()
		if ssss == "yes" {
			ssss = "http,https"
		} else if ssss == "no" {
			ssss = "http"
		}
		ip := models.NewIPAndCheck(ss + ":" + sss, ssss)
		if ip != nil {
			result = append(result, ip)
		}
	})
	log.Println("PLP done.")
	if len(result) > 2 {
		return result[2:]
	} else {
		return
	}
}
