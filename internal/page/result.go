package page

import (
	"github.com/ysmood/gson"
)

func (p *Page) SetResultsId() error {
	p.Rod.MustEval(`() => {
        const elements = document.querySelectorAll("tr[data-id]");
        elements.forEach((el, index) => {
            el.id = "inconsistency-" + (index + 1);
        });
    }`)

	p.Loading()

	return nil
}

func (p *Page) GetResults() (gson.JSON, error) {
	results := p.Rod.MustEval(`() => {
			const results = [];
				for (let i = 1; i <= 99; i++) {
					const row = document.querySelector('#inconsistency-' + i);
					if (row) {
						results.push({
							index: i,
							name: row.querySelector('td.ng-binding:nth-child(2)').textContent,
							hour: row.querySelector('td.ng-binding:nth-child(6)').textContent,
							category: row.querySelector('td.ng-binding:nth-child(7)').textContent,
						});
					}
				}
			return results;
			}`)

	return results, nil
}
