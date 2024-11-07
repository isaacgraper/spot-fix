package page

import (
	"fmt"
	"log"
)

func (p *Page) Pagination() (bool, error) {
	hasNextPage := p.Rod.MustHas(`[ng-click="changePage('next')"]`)
	if !hasNextPage {
		return false, fmt.Errorf("[pagination] error element was not found")
	}

	p.Loading()

	err := p.Click(`[ng-click="changePage('next')"]`)
	if err != nil {
		return false, fmt.Errorf("[pagination] error while trying to click into element: %w", err)
	}

	p.Loading()

	log.Println("[pagination] paginated to the next page")

	p.Loading()

	return true, nil
}
