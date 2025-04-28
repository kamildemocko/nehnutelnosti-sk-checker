package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct {
	doc *goquery.Document
}

func NewParser(html string) (*Parser, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	p := Parser{doc: doc}

	return &p, err
}

// parse all flats from webpage
func (p *Parser) ParseFlats() []*Flat {
	var flats []*Flat

	p.doc.Find(
		".MuiContainer-root > .MuiContainer-root > .MuiBox-root > .MuiBox-root > .MuiBox-root").
		Each(func(i int, item *goquery.Selection) {
			title := item.Find("h2").Text()
			link := item.Find("a").First().AttrOr(
				"href",
				"https://seosherpa.com/wp-content/uploads/2020/12/custom-404-error-page.png",
			)
			address := item.Find("p").Eq(0).Text()
			size := item.Find("p").Eq(1).Text()
			area := item.Find("p").Eq(2).Text()
			price := item.Find("p").Eq(11).Text()

			if title == "" {
				return
			}

			flats = append(flats, &Flat{
				Title:   title,
				Address: address,
				Size:    size,
				Area:    parseArea(area),
				Price:   parsePrice(price),
				Link:    link,
			})
		})

	return flats
}

func parseArea(area string) int {
	areaStr := strings.Split(area, " ")[0]
	areaInt, err := strconv.Atoi(areaStr)
	if err != nil {
		areaInt = 0
	}

	return areaInt
}

func parsePrice(price string) int {
	lastSpaceIndex := strings.LastIndex(price, " ")
	if lastSpaceIndex == -1 {
		return 0
	}

	priceStr := price[:lastSpaceIndex]
	priceStr = strings.ReplaceAll(priceStr, "\u00A0", "")
	priceInt, err := strconv.Atoi(priceStr)
	if err != nil {
		priceInt = 0
	}

	return priceInt
}
