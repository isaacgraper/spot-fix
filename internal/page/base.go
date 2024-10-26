package page

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
)

type Page struct {
	Page *rod.Page
}

func (p *Page) Click(selector string, screenshot bool) error {
	err := rod.Try(func() {
		element, err := p.Page.Element(selector)
		if err != nil {
			log.Printf("Element not found: %s", selector)
			return
		}

		p.Loading()

		err = element.Click(proto.InputMouseButtonLeft, 1)
		if err != nil {
			log.Printf("Failed to click element: %s", selector)
			return
		}

		time.Sleep(time.Millisecond * 200)

		if screenshot {
			p.Page.MustScreenshot(fmt.Sprintf("screenshot_%d.png", time.Now().Unix()))
		}
	})

	return err
}

func (p *Page) ClickWithRetry(selector string, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = rod.Try(func() {
			element, err := p.Page.Timeout(250 * time.Millisecond).Element(selector)
			if err != nil {
				log.Printf("Element not found: %s", selector)
				return
			}
			err = element.Click(proto.InputMouseButtonLeft, 1)
			if err != nil {
				log.Printf("Failed to click element: %s", selector)
				return
			}
		})

		if err == nil {
			return nil
		}

		log.Printf("Attempt %d failed to click on %s: %v", i+1, selector, err)

		time.Sleep(time.Second)
	}

	return fmt.Errorf("failed to click on %s after %d attempts: %v", selector, maxRetries, err)
}

func (p *Page) AddElementId(selector, id string) {
	p.Page.Eval(fmt.Sprintf(`() => {
		const el = %s;
		if (el) {
		el.id = "%s";
		} else {
			console.error("Element not found:", %s);
		}
	}`, selector, id, selector))

}

func (p *Page) ScrollToElement(selector string) {
	p.Page.Eval(fmt.Sprintf(`() => {
		const element = document.querySelector('%s');
		if (element) {
			element.scrollIntoView({ behavior: 'smooth', block: 'center' });
		} else {
			console.error("Element not found:", '%s');
		}
	}`, selector, selector))

}

func (p *Page) Pagination() bool {
	hasNextPage := p.Page.MustHas(`[ng-click="changePage('next')"]`)
	if !hasNextPage {
		return false
	}

	if err := p.ClickWithRetry(`[ng-click="changePage('next')"]`, 3); err != nil {
		return false
	}

	p.Loading()

	log.Println("[page] moved to the next page")
	return true
}

func (p *Page) Filter() error {
	if err := p.Click(`#inconsistenciesFilter`, false); err != nil {
		return fmt.Errorf("failed to click inconsistencies filter: %w", err)
	}

	element, err := p.Page.Element(`select#clockingTypes`)
	if err != nil {
		return fmt.Errorf("failed to find clocking types element: %w", err)
	}

	element.MustWaitStable()

	if err := p.Click(`#clockingTypes`, false); err != nil {
		return fmt.Errorf("failed to click clocking types: %w", err)
	}

	p.Loading()

	// filter category logic enters here

	err = element.Type(input.ArrowDown)
	if err != nil {
		return fmt.Errorf("failed to type arrow down: %w", err)
	}

	err = element.Click(proto.InputMouseButtonLeft, 1)
	if err != nil {
		return fmt.Errorf("failed to click selected option: %w", err)
	}

	if err := p.DateFilter(); err != nil {
		return fmt.Errorf("failed to apply a date filter: %w", err)
	}

	if err := p.Click(`#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding`, false); err != nil {
		return fmt.Errorf("failed to apply filter: %w", err)
	}

	log.Println("[page] filter applied")
	return nil
}

func (p *Page) DateFilter() error {
	el, err := p.Page.Element("input[name=\"finishDate\"]")
	if err != nil {
		log.Printf("Error finding element: %v\n", err)
		return err
	}

	date := time.Now()
	newDate := date.AddDate(0, 0, -7)

	el.MustInputTime(newDate)
	return nil
}

func (p *Page) Loading() {
	p.Page.MustWaitLoad().MustWaitStable().MustWaitDOMStable()
}
