package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/isaacgraper/spotfix.git/internal/common/config"
	"github.com/isaacgraper/spotfix.git/internal/report"
)

func (pr *Process) ProcessHandler(c *config.Config) (error, bool) {
	for {
		pr.ProcessResult(c)

		if !pr.page.Pagination() {
			log.Println("[processor] no more pages to process")
			break
		}
	}
	return nil, true
}

func (pr *Process) ProcessResult(c *config.Config) {
	if c.Max < 1 {
		log.Println("[processor] no results to process")
		return
	}

	batchSize := c.Batch
	for i := 0; i < c.Max; i += batchSize {
		end := i + batchSize
		if end > c.Max {
			end = c.Max
		}
		log.Printf("[processor] batch %d-%d initializing\n", i+1, end)
		pr.ProcessBatch(i+1, end, c)
	}

	log.Printf("[processor] ending processor with %d inconsistencies\n", len(pr.Results))

	if len(pr.Results) == 0 {
		log.Println("[processor] no inconsistencies found")
	} else {
		pr.Results = make([]report.ReportData, 0)
		pr.CompleteBatch("Cancelamento automático via Bot")
	}
}

func (pr *Process) ProcessBatch(start, end int, c *config.Config) error {
	pr.page.Rod.MustEval(`() => {
        const elements = document.querySelectorAll("tr[data-id]");
        elements.forEach((el, index) => {
            el.id = "inconsistency-" + (index + 1);
        });
    }`)

	pr.page.Loading()

	results := pr.page.Rod.MustEval(fmt.Sprintf(`() => {
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

	pr.page.Loading()

	for _, result := range results.Arr() {
		index := result.Get("index").Int()
		category := result.Get("category").String()
		hour := result.Get("hour").String()
		name := result.Get("name").String()

		hourSplit := strings.Split(hour, " ")
		hour = strings.TrimSpace(hourSplit[1])

		shouldProcess := (c.Hour == "" || hour == c.Hour) &&
			(c.Category == "" || category == c.Category) &&
			category != "Não registrado"

		if !shouldProcess {
			log.Println("[processor] inconsistency not found")
		}

		if shouldProcess {
			pr.Results = append(pr.Results, report.ReportData{
				Index:    index,
				Name:     name,
				Hour:     hour,
				Category: category,
			})

			report.NewReport(pr.Results).SaveReport()

			pr.page.Loading()

			err := pr.page.Click(fmt.Sprintf(`#inconsistency-%d.ng-scope i`, index))
			if err != nil {
				return fmt.Errorf("[processor] failed to click on inconsistency %w", err)
			}

			log.Printf("[processor] found:  %s - %s - %s", name, hour, category)
		}
	}

	return nil
}

func (pr *Process) ProcessNotRegistered() error {
	for {
		err := pr.page.Click(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`)
		if err != nil {
			return fmt.Errorf("[processor] failed to click filter checkbox: %w", err)
		}

		pr.page.Loading()

		if pr.CompleteNotRegistered("Cancelamento automático via Bot: Não registrados") {
			if pr.page.Pagination() {
				log.Println("[processor] page paginated...")
				continue
			}
		} else {
			break
		}

		pr.page.Loading()

	}

	return nil
}

func (pr *Process) ProcessWorkSchedule() error {
	for {
		err := pr.page.Click(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`)
		if err != nil {
			return fmt.Errorf("[processor] failed to click filter checkbox: %w", err)
		}

		pr.page.Loading()

		if pr.CompleteWorkSchedule("Cancelamento automático via Bot: Erros de escala") {
			if pr.page.Pagination() {
				log.Println("[processor] page paginated...")
				continue
			}
		} else {
			break
		}

		pr.page.Loading()

	}

	return nil
}
