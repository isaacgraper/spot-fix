package page

import "fmt"

func (p *Page) NavigateToInconsistencies() error {
	if err := p.Click(`a[href="/#/inconsistencies"]`); err != nil {
		return fmt.Errorf("[navigation] error while trying to click into inconsistencies %w", err)
	}

	p.Loading()

	if err := p.Click(`button.btn.btn-default[data-toggle]`); err != nil {
		return fmt.Errorf("[navigation] error while trying to click into data-toggle %w", err)
	}

	p.Rod.MustEval(`() => document.querySelector("div.app-content-body.nicescroll-continer > div.content-body > div.content-body-header > div.content-body-header-filters > div.filters-right > div > div > ul > li:nth-child(4) > a").id = "hundred-lines"`)

	if err := p.Click(`#hundred-lines`); err != nil {
		return fmt.Errorf("[navigation] error while trying to select a hundred-lines %w", err)
	}

	p.Loading()

	return nil
}
