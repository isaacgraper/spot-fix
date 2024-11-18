package filter

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/isaacgraper/spotfix.git/internal/page"
)

func FilterWorkSchedule(p *page.Page) (bool, error) {
	if err := p.Click(`i#inconsistenciesFilter`); err != nil {
		return false, fmt.Errorf("[filter] failed to click inconsistencies filter: %w", err)
	}

	element, err := p.Rod.Element(`select#clockingTypes`)
	if err != nil {
		return false, fmt.Errorf("[filter] failed to find clocking types element: %w", err)
	}

	if err := p.Click(`select#clockingTypes`); err != nil {
		return false, fmt.Errorf("[filter] failed to click clocking types: %w", err)
	}

	if err := ApplyFilterWorkSchedule(element); err != nil {
		return false, fmt.Errorf("[filter] error while trying to apply filter: %w", err)
	}

	p.Rod.Eval(`document.querySelector("#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding").id = "filter-btn"`)

	if err := p.Click("#filter-btn"); err != nil {
		return false, fmt.Errorf("[filter] error while trying to click into filter: %w", err)
	}

	p.Loading()

	p.Rod.MustWaitRequestIdle()
	time.Sleep(time.Second * 60)

	validate, err := ValidateDataWorkSchedule(p)

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

func ApplyFilterWorkSchedule(element *rod.Element) error {
	element.MustWaitVisible().
		MustSelect("Terminal não autorizado").
		MustSelect("Horário inválido").
		MustSelect("Fora do perímetro")
	return nil
}

func ValidateDataWorkSchedule(p *page.Page) (bool, error) {
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
