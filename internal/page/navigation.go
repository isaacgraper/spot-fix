package page

func (p *Page) NavigateToInconsistencies() error {
	if err := p.Click(`a[href="/#/inconsistencies"]`); err != nil {
		return nil
	}

	if err := p.Click(`.btn.btn-default[data-toggle]`); err != nil {
		return nil
	}

	p.Rod.MustEval(`() => document.querySelector("div.app-content-body.nicescroll-continer > div.content-body > div.content-body-header > div.content-body-header-filters > div.filters-right > div > div > ul > li:nth-child(4) > a").id = "hundred-lines"`)

	if err := p.Click(`#hundred-lines`); err != nil {
		return nil
	}

	p.Loading()

	return nil
}
