package bot

import (
	"fmt"
	"log"
	"strings"
	"time"

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

	batchSize := c.BatchSize
	for i := 0; i < c.Max; i += batchSize {
		end := i + batchSize
		if end > c.Max {
			end = c.Max
		}
		log.Println("[processor] batch initializing")
		pr.ProcessBatch(i+1, end, c)
	}

	// Implements email logic here
	log.Println("[finisher] ending processor")
	pr.EndProcess()
}

func (pr *Process) ProcessBatch(start, end int, c *config.Config) error {
	pr.page.Loading()

	pr.page.Page.MustEval(`() => {
        const elements = document.querySelectorAll("tr[data-id]");
        elements.forEach((el, index) => {
            el.id = "inconsistency-" + (index + 1);
        });
    }`)

	pr.page.Loading()

	results := pr.page.Page.MustEval(fmt.Sprintf(`() => {
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

	var data []report.ReportData

	for _, result := range results.Arr() {
		index := result.Get("index").Int()
		category := result.Get("category").String()
		hour := result.Get("hour").String()
		name := result.Get("name").String()

		hourSplit := strings.Split(hour, " ")
		hour = strings.TrimSpace(hourSplit[1])

		shouldProcess := (c.Hour == "" || hour == c.Hour) &&
			(c.Category == "" || category == c.Category)
			// && category != "Não registrado"

		if !shouldProcess {
			log.Println("[processor] inconsistency not found")
		}

		if shouldProcess {
			log.Printf("[processor] found:  %s - %s - %s", name, hour, category)

			data = append(data, report.ReportData{
				Index:    index,
				Name:     name,
				Hour:     hour,
				Category: category,
			})

			filename := fmt.Sprintf("relatório-inconsistências-%v.txt", time.Now().Format("02012006"))
			report.NewReport(filename, data).SaveReport()

			pr.page.Loading()
			time.Sleep(time.Millisecond * 250)

			if err := pr.page.ClickWithRetry(fmt.Sprintf(`#inconsistency-%d.ng-scope i`, index), 3); err != nil {
				log.Printf("failed to click on inconsistency %v", err)
			}

			log.Println("[processor] clicked into inconsistency")
		}
	}
	return nil
}

func (pr *Process) ProcessFilter(c *config.Config) {
	for {
		if err := pr.page.Click(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`, false); err != nil {
			log.Printf("Failed to click filter checkbox: %v", err)
			break
		}
		pr.page.Loading()

		if pr.EndProcess() {
			if pr.page.Pagination() {
				log.Println("[processor] pagination started")
				continue
			}
		} else {
			break
		}
		pr.page.Loading()
	}
}
