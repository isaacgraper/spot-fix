package page

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
)

type Page struct {
	Rod *rod.Page
}

func (p *Page) Click(selector string, screenshot bool) error {
	err := rod.Try(func() {
		element, err := p.Rod.Element(selector)
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
			p.Rod.MustScreenshot(fmt.Sprintf("screenshot_%d.png", time.Now().Unix()))
		}
	})

	return err
}

func (p *Page) ClickWithRetry(selector string, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = rod.Try(func() {
			element, err := p.Rod.Timeout(250 * time.Millisecond).Element(selector)
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
	p.Rod.Eval(fmt.Sprintf(`() => {
		const el = %s;
		if (el) {
		el.id = "%s";
		} else {
			console.error("Element not found:", %s);
		}
	}`, selector, id, selector))

}

func (p *Page) ScrollToElement(selector string) {
	p.Rod.Eval(fmt.Sprintf(`() => {
		const element = document.querySelector('%s');
		if (element) {
			element.scrollIntoView({ behavior: 'smooth', block: 'center' });
		} else {
			console.error("Element not found:", '%s');
		}
	}`, selector, selector))

}

func (p *Page) Pagination() bool {
	hasNextPage := p.Rod.MustHas(`[ng-click="changePage('next')"]`)
	if !hasNextPage {
		return false
	}

	p.Loading()

	if err := p.ClickWithRetry(`[ng-click="changePage('next')"]`, 6); err != nil {
		return false
	}

	p.Loading()

	log.Println("[page] paginated to the next page")

	p.Loading()

	return true
}

func (p *Page) Filter() (bool, error) {
	if err := p.Click(`#inconsistenciesFilter`, false); err != nil {
		return false, fmt.Errorf("failed to click inconsistencies filter: %w", err)
	}

	element, err := p.Rod.Element(`select#clockingTypes`)
	if err != nil {
		return false, fmt.Errorf("failed to find clocking types element: %w", err)
	}

	element.MustWaitStable()

	if err := p.Click(`#clockingTypes`, false); err != nil {
		return false, fmt.Errorf("failed to click clocking types: %w", err)
	}

	p.Loading()

	// filter category logic enters here

	err = element.Type(input.ArrowDown)
	if err != nil {
		return false, fmt.Errorf("failed to type arrow down: %w", err)
	}

	err = element.Click(proto.InputMouseButtonLeft, 1)
	if err != nil {
		return false, fmt.Errorf("failed to click selected option: %w", err)
	}

	date, ok, err := p.DateFilter()
	if err != nil {
		return false, fmt.Errorf("failed to apply a date filter: %w", err)
	}

	if !ok {
		log.Println("[page] filter not applied")
		return false, nil
	}

	if err := p.Click(`#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding`, false); err != nil {
		return false, fmt.Errorf("failed to apply filter: %w", err)
	}

	log.Println("[page] filter applied!")

	if !p.CheckDateFilter(date) {
		return false, fmt.Errorf("failed")
	}

	return true, nil
}

func (p *Page) DateFilter() (string, bool, error) {
	el, err := p.Rod.Element("input[name=\"finishDate\"]")
	if err != nil {
		log.Printf("Error finding element: %v\n", err)
		return " ", false, err
	}

	date := time.Now()
	newDate := date.AddDate(0, 0, -7)

	el.MustInputTime(newDate)

	p.Loading()

	log.Printf("[page] date: %s passed to the filter", newDate.Format("02-01-2006"))
	return newDate.Format("02-01-2006"), true, nil
}

func (p *Page) CheckDateFilter(dateFilter string) bool {
	p.Rod.MustEval(`() => document.querySelectorAll("tr[data-id] > td.ng-binding:nth-child(6)")[0].id = "i-date"`)

	date := p.Rod.MustElement("td#i-date.ng-binding").MustText()

	dateSplit := strings.Split(date, " ")
	date = strings.TrimSpace(dateSplit[0])

	dateTime, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false
	}

	log.Println(dateTime.Format("02-01-2006"))
	log.Println(dateFilter)

	if dateFilter != dateTime.Format("02-01-2006") {
		return false
	}
	return true
}

func (p *Page) Loading() {
	p.Rod.MustWaitLoad().MustWaitStable().MustWaitDOMStable()
}
