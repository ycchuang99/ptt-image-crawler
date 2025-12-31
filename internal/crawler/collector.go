package crawler

import (
	"github.com/gocolly/colly/v2"
)

const pptUrl string = "https://www.ptt.cc"

type Collector interface {
	CollectBoardList() ([]Board, error)
}

type Board struct {
	name  string
	class string
}

func (b Board) Title() string       { return b.name }
func (b Board) Description() string { return b.class }
func (b Board) FilterValue() string { return b.name }

func CollectBoardList() ([]Board, error) {
	boards := []Board{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "over18=1")
	})

	c.OnHTML(".b-ent", func(e *colly.HTMLElement) {
		boards = append(boards, Board{
			name:  e.ChildText(".board-name"),
			class: e.ChildText(".board-class"),
		})
	})

	err := c.Visit(pptUrl + "/bbs/")

	return boards, err
}
