package page

import (
	"fmt"
	"log"

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

func (p *Page) GetResults(start, end int) (gson.JSON, error) {
	if end != 0 {
		results := p.Rod.MustEval(fmt.Sprintf(`() => {
			const results = [];
				for (let i = %d; i <= %d; i++) {
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
			}`, start, end))

		return results, nil
	} else {
		var qt int

		log.Println("[results] end is 0")

		qt = p.Rod.MustEval(`() => document.querySelectorAll("tr[data-id]").length`).Int() - 1
		results := p.Rod.MustEval(fmt.Sprintf(`() => {
			const results = [];
				for (let i = %d; i <= %d; i++) {
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
			}`, start, qt))

		return results, nil
	}
}
