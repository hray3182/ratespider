//從台銀獲取匯率數據並輸出成json格式

package ratespider

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//從台銀獲取匯率數據並輸出成json格式
type Ratedata struct {
	currency, cashbuy, cashsell, spotbuy, spotsell string
}

var ratedata []Ratedata

func Fetch() ([]Ratedata, error) {

	req, _ := http.NewRequest("GET", "https://rate.bot.com.tw/xrt?Lang=zh-TW", nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)

	}

	if resp.StatusCode != 200 {
		fmt.Println("http status code:", resp.StatusCode)

	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fn := func(s string) []string {
		fn := func(i int, s *goquery.Selection) string {
			return strings.TrimSpace(s.Text())
		}
		return doc.Find(s).Map(fn)
	}

	cu := fn("div.hidden-phone.print_show")
	cb := fn("div > table > tbody > tr > td:nth-child(2)")
	cs := fn("div > table > tbody > tr > td:nth-child(3)")
	sb := fn("div > table > tbody > tr > td:nth-child(4)")
	ss := fn("div > table > tbody > tr > td:nth-child(5)")

	for i := range cu {
		single := Ratedata{cu[i], cb[i], cs[i], sb[i], ss[i]}
		ratedata = append(ratedata, single)
	}

	return ratedata, err
}