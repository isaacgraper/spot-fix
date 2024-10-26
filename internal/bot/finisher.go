package bot

import (
	"log"
	"time"
)

func (pr *Process) EndProcess() bool {
	pr.page.Page.MustElement(`td.ng-binding`).ScrollIntoView()

	elements := []string{
		`#content > div.app-content-body.nicescroll-continer > div.content-body > div.content-body-header > div.content-body-header-filters > div.filters-right > button`,
		`[btn-radio="\'CANCELED\'"]`,
		`#app > modal > div > div > div > div > div.modal-body > div > div > div:nth-child(2) > div > multiselect > div > div > div:nth-child(1) > div > i`,
		`[alt="Erro operacional"]`,
	}

	for _, selector := range elements {
		time.Sleep(time.Millisecond * 250)

		if err := pr.page.Click(selector, false); err != nil {
			log.Printf("Failed to click on %s: %v", selector, err)
			return false
		}
		pr.page.Loading()
	}

	note := pr.page.Page.MustElement(`input#note`)
	note.MustInput("Cancelamento automático via bot")

	if err := pr.page.Click(`a.btn.button_link.btn-primary.ng-binding`, false); err != nil {
		log.Printf("Failed to click on submit button: %v", err)
		return false
	}
	pr.page.Loading()

	log.Println("[process] inconsistencies processed")
	return true
}
