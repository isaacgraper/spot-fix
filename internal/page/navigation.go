package page

func (p *Page) NavigateToInconsistencies() error {
	if err := p.Click(`a[href="/#/inconsistencies"]`, false); err != nil {
		return nil
	}

	if err := p.Click(`.btn.btn-default[data-toggle]`, false); err != nil {
		return nil
	}

	p.Page.MustEval(`() => document.querySelector("div.app-content-body.nicescroll-continer > div.content-body > div.content-body-header > div.content-body-header-filters > div.filters-right > div > div > ul > li:nth-child(4) > a").id = "hundred-lines"`)

	if err := p.Click(`#hundred-lines`, false); err != nil {
		return nil
	}

	p.Loading()

	has, _, _ := p.Page.Has(`document.querySelector('.beamerAnnouncementSnippet')`)
	if !has {
		p.Page.MustEval(`() => document.querySelector('.beamerAnnouncementSnippet').style.display="none"`)
	}

	return nil
}
