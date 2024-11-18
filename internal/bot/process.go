package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/isaacgraper/spotfix.git/internal/common/config"
	"github.com/isaacgraper/spotfix.git/internal/report"
)

func (pr *Process) ProcessHandler(c *config.Config) (bool, error) {
	for {
		pr.ProcessResult(c)

		pagination, err := pr.page.Pagination()

		if err != nil {
			return false, fmt.Errorf("[process] error while trying to paginate: %w", err)
		}

		if !pagination {
			log.Panic("[process] no more pages to process")
			break
		}
	}
	return true, nil
}

func (pr *Process) ProcessResult(c *config.Config) {
	if c.Max < 1 {
		log.Println("[process] no results to process")
		return
	}

	batchSize := c.Batch
	for i := 0; i < c.Max; i += batchSize {
		end := i + batchSize
		if end > c.Max {
			end = c.Max
		}
		log.Printf("[process] batch %d-%d initializing\n", i+1, end)
		pr.ProcessBatch(i+1, end, c)
	}

	log.Printf("[process] ending process with %d inconsistencies\n", len(pr.Results))

	if len(pr.Results) == 0 {
		log.Println("[process] no inconsistencies found")
	} else {
		pr.Results = make([]report.ReportData, 0)
		pr.CompleteBatch("Cancelamento automático via Bot")
	}
}

func (pr *Process) ProcessBatch(start, end int, c *config.Config) error {

	pr.page.SetResultsId()

	results, err := pr.page.GetResults(start, end)
	if err != err {
		return fmt.Errorf("[process] error while trying to evaluate inconsistencies: %w", err)
	}

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
			log.Println("[process] inconsistency not found")
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
				return fmt.Errorf("[process] failed to click on inconsistency %w", err)
			}

			log.Printf("[process] found:  %s - %s - %s", name, hour, category)
		}
	}

	return nil
}

func (pr *Process) ProcessNotRegistered() error {
	log.Println("[process] processing inconsistencies...")

	for {
		err := pr.page.Click(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`)
		if err != nil {
			return fmt.Errorf("[process] failed to click filter checkbox: %w", err)
		}

		pr.page.Loading()

		pr.page.SetResultsId()

		results, err := pr.page.GetResults(1, 0)
		if err != nil {
			return fmt.Errorf("[process] error while trying to evaluate inconsistencies: %w", err)
		}

		for _, result := range results.Arr() {
			index := result.Get("index").Int()
			category := result.Get("category").String()
			hour := result.Get("hour").String()
			name := result.Get("name").String()

			pr.Results = append(pr.Results, report.ReportData{
				Index:    index,
				Name:     name,
				Hour:     hour,
				Category: category,
			})

			category = strings.TrimSpace(category)

			if category != "Não registrado" {
				log.Panicf("[process] inconsistence category must not be different from the filter")
			}
		}

		log.Println("[process] saving results")

		report.NewReport(pr.Results).SaveReport()

		pr.page.Loading()

		complete, err := pr.CompleteNotRegistered("Cancelamento automático via Bot: Não Registrado")
		if err != nil {
			return fmt.Errorf("[process] error while trying to complete workSchedule process %w", err)
		}

		if complete {
			pagination, err := pr.page.Pagination()
			if err != nil {
				return fmt.Errorf("[process] error while trying to paginate")
			}

			if pagination {
				log.Println("[process] page paginated...")
				continue
			}

		} else {
			log.Panicf("[process] error ocurred while trying to process and paginated workSchedule...")
			break
		}

		pr.page.Loading()

	}

	return nil
}

func (pr *Process) ProcessWorkSchedule() error {
	log.Println("[process] processing inconsistencies...")

	for {
		hasCheckbox := pr.page.Rod.MustHas(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`)
		if !hasCheckbox {
			return fmt.Errorf("[process] error element was not found")
		}

		err := pr.page.Click(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`)
		if err != nil {
			return fmt.Errorf("[process] failed to click filter checkbox: %w", err)
		}

		pr.page.Loading()

		pr.page.SetResultsId()

		results, err := pr.page.GetResults(1, 0)
		if err != nil {
			return fmt.Errorf("[process] error while trying to evaluate inconsistencies: %w", err)
		}

		for _, result := range results.Arr() {
			index := result.Get("index").Int()
			category := result.Get("category").String()
			hour := result.Get("hour").String()
			name := result.Get("name").String()

			pr.Results = append(pr.Results, report.ReportData{
				Index:    index,
				Name:     name,
				Hour:     hour,
				Category: category,
			})

			category = strings.TrimSpace(category)

			shouldProcess := (category == "Terminal não autorizado") ||
				(category == "Horário inválido") ||
				(category == "Fora do perímetro")

			if shouldProcess {
				continue
			} else {
				log.Panicf("[process] inconsistence category must not be different from the filter")
			}
		}

		log.Println("[process] saving results")

		report.NewReport(pr.Results).SaveReport()

		pr.page.Loading()

		complete, err := pr.CompleteWorkSchedule("Ajustado automaticamente via Bot: Erros de escala")
		if err != nil {
			return fmt.Errorf("[process] error while trying to complete workSchedule process %w", err)
		}

		if complete {
			pagination, err := pr.page.Pagination()
			if err != nil {
				return fmt.Errorf("[process] error while trying to paginate")
			}

			if pagination {
				log.Println("[process] page paginated...")
				continue
			}

		} else {
			log.Panicf("[process] error ocurred while trying to process and paginated workSchedule...")
			break
		}

		pr.page.Loading()

	}

	return nil
}
