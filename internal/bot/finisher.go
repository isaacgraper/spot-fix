package bot

import (
	"log"
	"time"
)

func (pr *Process) Finisher(note string) bool {
	pr.page.Rod.MustElement(`td.ng-binding`).ScrollIntoView()

	elements := []string{
		`#content > div.app-content-body.nicescroll-continer > div.content-body > div.content-body-header > div.content-body-header-filters > div.filters-right > button`,
		`[btn-radio="\'CANCELED\'"]`,
		`#app > modal > div > div > div > div > div.modal-body > div > div > div:nth-child(2) > div > multiselect > div > div > div:nth-child(1) > div > i`,
		`[alt="Erro operacional"]`,
	}

	for _, selector := range elements {
		time.Sleep(200 * time.Millisecond)

		err := pr.page.Click(selector)
		if err != nil {
			log.Printf("[finisher] failed to click on %s: %v", selector, err)
			return false
		}

		pr.page.Loading()
	}

	noteEl := pr.page.Rod.MustElement(`input#note`)

	noteEl.MustInput(note)

	err := pr.page.Click(`a.btn.button_link.btn-primary.ng-binding`)
	if err != nil {
		log.Printf("[finisher] failed to click on submit button: %v", err)
		return false
	}

	pr.page.Loading()

	log.Println("[finisher] inconsistencies processed!")

	return true
}
