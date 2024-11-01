package page

func (p *Page) Pagination() bool {
	hasNextPage := p.Rod.MustHas(`[ng-click="changePage('next')"]`)
	if !hasNextPage {
		return false
	}

	p.Loading()

	err := p.Click(`[ng-click="changePage('next')"]`)
	if err != nil {
		p.logger.Println("[pagination] error while trying to click in the element")
		return false
	}

	p.Loading()

	p.logger.Println("[pagination] paginated to the next page")

	p.Loading()

	return true
}
