package page

func (p *Page) Filter() (bool, error) {
	return true, nil
}

func (p *Page) ApplyFiler() error {
	return nil
}

func (p *Page) ApplyDateFilter() error {
	return nil
}

func (p *Page) ValidateDateFilter() error {
	return nil
}

// func (p *Page) Filter() (bool, error) {
// 	if err := p.Click(`#inconsistenciesFilter`, false); err != nil {
// 		return false, fmt.Errorf("failed to click inconsistencies filter: %w", err)
// 	}

// 	element, err := p.Rod.Element(`select#clockingTypes`)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to find clocking types element: %w", err)
// 	}

// 	element.MustWaitStable()

// 	if err := p.Click(`#clockingTypes`, false); err != nil {
// 		return false, fmt.Errorf("failed to click clocking types: %w", err)
// 	}

// 	p.Loading()

// 	// filter category logic enters here

// 	err = element.Type(input.ArrowDown)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to type arrow down: %w", err)
// 	}

// 	err = element.Click(proto.InputMouseButtonLeft, 1)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to click selected option: %w", err)
// 	}

// 	date, ok, err := p.DateFilter()
// 	if err != nil {
// 		return false, fmt.Errorf("failed to apply a date filter: %w", err)
// 	}

// 	if !ok {
// 		log.Println("[page] filter not applied")
// 		return false, nil
// 	}

// 	if err := p.Click(`#app > searchfilterinconsistencies > div > div.row.overscreen_child > div.filter_container > div.hbox.filter_button.ng-scope > a.btn.button_link.btn-dark.ng-binding`, false); err != nil {
// 		return false, fmt.Errorf("failed to apply filter: %w", err)
// 	}

// 	log.Println("[page] filter applied!")

// 	if !p.CheckDateFilter(date) {
// 		return false, fmt.Errorf("failed")
// 	}

// 	return true, nil
// }
// func (p *Page) DateFilter() (string, bool, error) {
// 	el, err := p.Rod.Element("input[name=\"finishDate\"]")
// 	if err != nil {
// 		log.Printf("Error finding element: %v\n", err)
// 		return " ", false, err
// 	}

// 	date := time.Now()
// 	newDate := date.AddDate(0, 0, -7)

// 	el.MustInputTime(newDate)

// 	p.Loading()

// 	log.Printf("[page] date: %s passed to the filter", newDate.Format("02-01-2006"))
// 	return newDate.Format("02-01-2006"), true, nil
// }

// func (p *Page) CheckDateFilter(dateFilter string) bool {
// 	p.Rod.MustEval(`() => document.querySelectorAll("tr[data-id] > td.ng-binding:nth-child(6)")[0].id = "i-date"`)

// 	date := p.Rod.MustElement("td#i-date.ng-binding").MustText()

// 	dateSplit := strings.Split(date, " ")
// 	date = strings.TrimSpace(dateSplit[0])

// 	dateTime, err := time.Parse("02/01/2006", date)
// 	if err != nil {
// 		return false
// 	}

// 	log.Println(dateTime.Format("02-01-2006"))
// 	log.Println(dateFilter)

// 	if dateFilter != dateTime.Format("02-01-2006") {
// 		return false
// 	}
// 	return true
// }
