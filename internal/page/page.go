package page

import (
	"github.com/go-rod/rod"
)

type Page struct {
	Rod *rod.Page
}

func (p *Page) NewPage() *Page {
	return &Page{}
}
