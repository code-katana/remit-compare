package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
	"time"
	"golang.org/x/net/html"
	"github.com/yhat/scrape"
	"log"
	"fmt"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	timestamp := int32(time.Now().Unix())
	url := fmt.Sprintf("https://www.westernunion.com/retailpresentationservice/rest/api/v1.0/FeeInquiryEstimated?timestamp=%d", timestamp)
	log.Print(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Cannot get response", err)
	}
	root , err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("Cannot parse response", err)
	}
	title , ok := scrape.Find(root, scrape.ByClass("head"))
	if ok {
		c.Set("title", scrape.Text(title))
	}
	return c.Render(200, r.HTML("index.html"))
}
