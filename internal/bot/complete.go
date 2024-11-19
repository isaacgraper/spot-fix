package bot

import (
	"fmt"
	"log"
	"time"
)

func (pr *Process) CompleteNotRegistered(note string) (bool, error) {
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
			return false, fmt.Errorf("[complete] failed to click on %s: %w", selector, err)
		}

		pr.page.Loading()
	}

	noteEl := pr.page.Rod.MustElement(`input#note`)

	noteEl.MustInput(note)

	err := pr.page.Click(`a.btn.button_link.btn-primary.ng-binding`)
	if err != nil {
		return false, fmt.Errorf("[complete] failed to click on submit button: %w", err)
	}

	pr.page.Loading()

	log.Println("[complete] inconsistencies processed!")

	// A modal can appear, so we need to remove it then paginate the page
	hasModal, _, err := pr.page.Rod.Has("div.modal")

	if err != nil {
		return false, fmt.Errorf("[complete] an error has ocurred: %w", err)
	}

	if hasModal {
		log.Println("[complete] an wild modal has apperead!")

		hasBtn, _, err := pr.page.Rod.Has("div.close")

		if err != nil {
			return false, fmt.Errorf("[complete] an error has ocurred: %w", err)
		}

		if hasBtn {
			if err := pr.page.Click(`div.close`); err != nil {
				return false, fmt.Errorf("[complete] error while trying to close modal: %w", err)
			}
		}
	}

	return true, nil
}

func (pr *Process) CompleteWorkSchedule(note string) (bool, error) {
	pr.page.Rod.MustElement(`td.ng-binding`).ScrollIntoView()

	elements := []string{
		`#content > div.app-content-body.nicescroll-continer > div.content-body > div.content-body-header > div.content-body-header-filters > div.filters-right > button`,
		`#app > modal > div > div > div > div > div.modal-body > div > div > div:nth-child(2) > div > multiselect > div > div > div:nth-child(1) > div > i`,
		`[alt="Erro operacional"]`,
	}

	for _, selector := range elements {
		time.Sleep(200 * time.Millisecond)

		err := pr.page.Click(selector)
		if err != nil {
			return false, fmt.Errorf("[complete] failed to click on %s: %w", selector, err)
		}

		pr.page.Loading()
	}

	noteEl := pr.page.Rod.MustElement(`input#note`)

	noteEl.MustInput(note)

	err := pr.page.Click(`a.btn.button_link.btn-primary.ng-binding`)
	if err != nil {
		return false, fmt.Errorf("[complete] failed to click on submit button: %w", err)
	}

	pr.page.Loading()

	log.Println("[complete] inconsistencies processed!")

	// A modal can appear, so we need to remove it then paginate the page
	hasModal, _, err := pr.page.Rod.Has("div.modal")

	if err != nil {
		return false, fmt.Errorf("[complete] an error has ocurred: %w", err)
	}

	if hasModal {
		log.Println("[complete] an wild modal has apperead!")

		hasBtn, _, err := pr.page.Rod.Has("div.close")

		if err != nil {
			return false, fmt.Errorf("[complete] an error has ocurred: %w", err)
		}

		if hasBtn {
			if err := pr.page.Click(`div.close`); err != nil {
				return false, fmt.Errorf("[complete] error while trying to close modal: %w", err)
			}
		}
	}

	return true, nil
}

func (pr *Process) CompleteBatch(note string) (bool, error) {
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
			return false, fmt.Errorf("[complete] failed to click on %s: %w", selector, err)
		}

		pr.page.Loading()
	}

	noteEl := pr.page.Rod.MustElement(`input#note`)

	noteEl.MustInput(note)

	err := pr.page.Click(`a.btn.button_link.btn-primary.ng-binding`)
	if err != nil {
		return false, fmt.Errorf("[complete] failed to click on submit button: %w", err)
	}

	pr.page.Loading()

	log.Println("[complete] inconsistencies processed!")

	pr.page.Loading()

	return true, nil
}
