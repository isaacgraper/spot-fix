package page

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func (p *Page) Click(selector string) error {
	err := rod.Try(func() {
		element, err := p.Rod.Element(selector)
		if err != nil {
			log.Printf("[click] element not found: %s", err)
			panic(err)
		}

		p.Loading()

		err = element.Click(proto.InputMouseButtonLeft, 1)
		if err != nil {
			log.Printf("[click] failed to click element: %s", err)
			panic(err)
		}

		time.Sleep(time.Millisecond * 500)
	})

	if err != nil {
		return fmt.Errorf("[click] failed while trying to click: %w", err)
	}

	return nil
}
