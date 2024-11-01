package page

import (
	"log"

	"github.com/go-rod/rod"
)

type Page struct {
	Rod    *rod.Page
	logger *log.Logger
}

func (p *Page) NewPage() *Page {
	return &Page{}
}
