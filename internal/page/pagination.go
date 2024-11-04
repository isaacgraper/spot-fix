package page

import "log"

func (p *Page) Pagination() bool {
	hasNextPage := p.Rod.MustHas(`[ng-click="changePage('next')"]`)
	if !hasNextPage {
		return false
	}

	p.Loading()

	err := p.Click(`[ng-click="changePage('next')"]`)
	if err != nil {
		log.Println("[pagination] error while trying to click in the element")
		return false
	}

	p.Loading()

	log.Println("[pagination] paginated to the next page")

	p.Loading()

	return true
}
