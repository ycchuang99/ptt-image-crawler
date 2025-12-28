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

	c := colly.NewCollector()

	c.OnHTML(".b-ent", func(e *colly.HTMLElement) {
		boards = append(boards, Board{
			name:  e.ChildText(".board-name"),
			class: e.ChildText(".board-class"),
		})
	})

	err := c.Visit(pptUrl + "/bbs/")

	return boards, err
}
