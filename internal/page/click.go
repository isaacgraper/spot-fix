package page

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const (
	retries = 6
)

func (p *Page) Click(selector string) error {
	var err error

	for i := 0; i < retries; i++ {
		err = rod.Try(func() {
			element, err := p.Rod.Timeout(200 * time.Millisecond).Element(selector)
			if err != nil {
				log.Printf("[click] element not found: %s", selector)
				return
			}

			err = element.Click(proto.InputMouseButtonLeft, 1)
			if err != nil {
				log.Printf("[click] failed to click element: %s", selector)
				return
			}
		})

		if err == nil {
			return nil
		}

		log.Printf("[click] attempt %d failed to click on %s: %v", i+1, selector, err)

		time.Sleep(200 * time.Millisecond)
	}

	return fmt.Errorf("[click] failed to click on %s after %d attempts: %v", selector, retries, err)
}
