package filter

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/isaacgraper/spotfix.git/internal/page"
)

func FilterWorkSchedule(p *page.Page) error {
	if err := p.Click(`i#inconsistenciesFilter`); err != nil {
		return fmt.Errorf("[filter] failed to click inconsistencies filter: %w", err)
	}

	element, err := p.Rod.Element(`select#clockingTypes`)
	if err != nil {
		return fmt.Errorf("[filter] failed to find clocking types element: %w", err)
	}

	if err := p.Click(`select#clockingTypes`); err != nil {
		return fmt.Errorf("[filter] failed to click clocking types: %w", err)
	}

	if err := ApplyFilterWorkSchedule(element); err != nil {
		return fmt.Errorf("[filter] error while trying to apply filter: %w", err)
	}

	p.Rod.Eval(`document.querySelector("#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding").id = "filter-btn"`)

	if err := p.Click("#filter-btn"); err != nil {
		return fmt.Errorf("[filter] error while trying to click into filter: %w", err)
	}

	p.Loading()

	return nil
}

func ApplyFilterWorkSchedule(element *rod.Element) error {
	element.MustWaitVisible().
		MustSelect("Terminal não autorizado").
		MustSelect("Horário inválido").
		MustSelect("Fora do perímetro")
	return nil
}
