package filter

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/isaacgraper/spotfix.git/internal/page"
)

func FilterNotRegistered(p *page.Page) (bool, error) {
	if err := p.Click(`i#inconsistenciesFilter`); err != nil {
		return false, fmt.Errorf("[filter] failed to click inconsistencies filter: %w", err)
	}

	element, err := p.Rod.Element(`select#clockingTypes`)
	if err != nil {
		return false, fmt.Errorf("[filter] failed to find clocking types element: %w", err)
	}

	element.MustWaitStable()

	if err := p.Click(`select#clockingTypes`); err != nil {
		return false, fmt.Errorf("[filter] failed to click clocking types: %w", err)
	}

	if err := ApplyFilerNotRegistered(element); err != nil {
		return false, fmt.Errorf("[filter] error while trying to apply filter: %w", err)
	}

	dateFilter, err := ApplyDateFilter(p)
	if err != nil {
		return false, fmt.Errorf("[filter] error while trying to apply date filter: %w", err)
	}

	log.Printf("[filter] filter data: %s", dateFilter)

	p.Rod.Eval(`document.querySelector("#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding").id = "filter-btn"`)

	if err := p.Click("#filter-btn"); err != nil {
		return false, fmt.Errorf("[filter] error while trying to click into filter: %w", err)
	}

	p.Loading()

	p.Rod.MustWaitRequestIdle()
	time.Sleep(time.Second * 90)

	// ok, err := ValidateDateFilter(dateFilter, p)
	// if err != nil {
	// 	return false, fmt.Errorf("[filter] error while trying to validate date filter: %w", err)
	// }

	// if !ok {
	// 	log.Println("date expected is not from 1 week ago...")
	// 	return false, nil
	// }

	validate, err := ValidateDataNotRegistered(p)

	if err != nil {
		return false, fmt.Errorf("[filter] error while trying to check if data was not found: %w", err)
	}

	if validate {
		log.Println("[filter] no inconsistencies found")
		return false, nil
	}

	p.Loading()

	return true, nil
}

func ApplyFilerNotRegistered(element *rod.Element) error {
	element.MustWaitVisible().
		MustSelect("Não registrado")
	return nil
}

func ApplyDateFilter(p *page.Page) (string, error) {
	el, err := p.Rod.Element("input#finishDate")
	if err != nil {
		return " ", fmt.Errorf("[filter] error finding element: %w", err)
	}

	date := time.Now()
	newDate := date.AddDate(0, 0, -8)

	el.MustInputTime(newDate)

	p.Loading()

	return newDate.Format("02-01-2006"), nil
}

// func ValidateDateFilter(dateFilter string, p *page.Page) (bool, error) {

// 	p.Rod.MustEval(`() => document.querySelectorAll("tr[data-id] > td.ng-binding:nth-child(6)")[0].id = "first-date"`)

// 	date := p.Rod.MustElement("td#first-date.ng-binding").MustText()

// 	datesplit := strings.Split(date, " ")
// 	date = strings.TrimSpace(datesplit[0])

// 	datetime, err := time.Parse("02/01/2006", date)
// 	if err != nil {
// 		return false, nil
// 	}

// 	if dateFilter != datetime.Format("02-01-2006") {
// 		log.Printf("date filer: %s - date expected: %s", dateFilter, datetime.Format("02-01-2006"))
// 		return false, nil
// 	}
// 	return true, nil
// }

func ValidateDataNotRegistered(p *page.Page) (bool, error) {
	has := p.Rod.MustHas("td>p")
	if has {
		el, err := p.Rod.Element("td>p")

		if err != nil {
			return false, fmt.Errorf("[filter] error while trying to get element: %w", err)
		}

		if el == nil {
			return false, nil
		}

		if el.MustText() == "Nenhum registro encontrado" {
			return true, nil
		}
	}
	return false, nil
}
