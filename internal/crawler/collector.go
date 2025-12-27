package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

var pptUrl string = "https://www.ptt.cc"

type Collector interface {
	CollectBoardList() ([]Board, error)
}

type Board struct {
	Name  string
	Class string
}

func (b Board) String() string {
	return fmt.Sprintf("%-12s [%s]", b.Name, b.Class)
}

func CollectBoardList() ([]Board, error) {
	boards := []Board{}

	c := colly.NewCollector()

	c.OnHTML(".b-ent", func(e *colly.HTMLElement) {
		boards = append(boards, Board{
			Name:  e.ChildText(".board-name"),
			Class: e.ChildText(".board-class"),
		})
	})

	err := c.Visit(pptUrl + "/bbs/")

	return boards, err
}

