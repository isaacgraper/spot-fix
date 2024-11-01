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

	element, err := p.Page.Element(`select#clockingTypes`)
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

	ok, err := p.DateFilter()
	if err != nil {
		return false, fmt.Errorf("failed to apply a date filter: %w", err)
	}

	if !ok {
		log.Println("[page] filter not applied, date is not 1 week ago")
		return false, nil
	}

	if err := p.Click(`#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding`, false); err != nil {
		return false, fmt.Errorf("failed to apply filter: %w", err)
	}

	log.Println("[page] filter applied!")
	return true, nil
}

func (p *Page) DateFilter() (bool, error) {
	el, err := p.Page.Element("input[name=\"finishDate\"]")
	if err != nil {
		log.Printf("Error finding element: %v\n", err)
		return false, err
	}

	date := time.Now()
	newDate := date.AddDate(0, 0, -7)

	if !p.CheckDateFilter(newDate.Format("02-01-2006")) {
		return false, nil
	} else {
		el.MustInputTime(newDate)
	}

	log.Printf("[page] date: %s passed to the filter", newDate.Format("02-01-2006"))
	return true, nil
}

func (p *Page) CheckDateFilter(dateFilter string) bool {
	p.Page.MustEval(`() => document.querySelectorAll("tr[data-id] > td.ng-binding:nth-child(5)")[0].id = "inconsistency-date"`)

	log.Println("[page] evaluating inconsistency-date")

	el := p.Page.MustElement("td#inconsistency-date.ng-binding")
	log.Println("Element text:", el.MustText())
	log.Println("Element HTML:", el.MustHTML())
	log.Println("Element attributes:", el.MustAttribute("id"), el.MustAttribute("class"))

	date := p.Page.MustElement("td#inconsistency-date.ng-binding").MustText()
	log.Println("Data:", date)

	// dateSplit := strings.Split(date, " ")
	// date = strings.TrimSpace(dateSplit[0])

	// log.Println(dateFilter)
	// log.Println(date)

	// if dateFilter != date {
	// 	log.Println("[page] date rejected by the CheckDateFilter func")
	// 	return false
	// }
	return false // true
}

func (p *Page) Loading() {
	p.Page.MustWaitLoad().MustWaitStable().MustWaitDOMStable()
}
