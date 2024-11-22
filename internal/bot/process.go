package bot

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/isaacgraper/spotfix.git/internal/report"
)

func (pr *Process) ProcessNotRegistered() error {
	log.Println("[process] processing inconsistencies...")

	for {
		err := pr.page.Click(`#content > div.app-content-body.nicescroll-continer > div.content-body > div.app-content-body > div.tab-lis > div.content-table > table > thead > tr > th:nth-child(1) > label > i`)
		if err != nil {
			return fmt.Errorf("[process] failed to click filter checkbox: %w", err)
		}

		pr.page.Loading()

		pr.page.SetResultsId()

		results, err := pr.page.GetResults()
		if err != nil {
			return fmt.Errorf("[process] error while trying to evaluate inconsistencies: %w", err)
		}

		for _, result := range results.Arr() {
			index := result.Get("index").Int()
			category := strings.TrimSpace(result.Get("category").String())
			hour := result.Get("hour").String()
			name := result.Get("name").String()

			pr.Results = append(pr.Results, report.ReportData{
				Index:    index,
				Name:     name,
				Hour:     hour,
				Category: category,
			})

			date := strings.Split(hour, " ")[0]
			date = strings.ReplaceAll(date, "/", "-")

			parsedDate, err := time.Parse("01-02-2006", date)
			if err != nil {
				log.Panicf("[process] error while trying to parse string to date: %s", date)
			}

			oneWeekAgo := time.Now().AddDate(0, 0, -7)

			if parsedDate.After(oneWeekAgo) {
				log.Panicln("[process] date must not be different from the filter")
			}

			if category != "Não registrado" {
				log.Panicln("[process] category must not be different from the filter")
			}
		}

		log.Println("[process] saving results")

		report.NewReport(pr.Results).SaveReport("não-registrado")

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

		results, err := pr.page.GetResults()
		if err != nil {
			return fmt.Errorf("[process] error while trying to evaluate inconsistencies: %w", err)
		}

		for _, result := range results.Arr() {
			index := result.Get("index").Int()
			category := strings.TrimSpace(result.Get("category").String())
			hour := result.Get("hour").String()
			name := result.Get("name").String()

			pr.Results = append(pr.Results, report.ReportData{
				Index:    index,
				Name:     name,
				Hour:     hour,
				Category: category,
			})

			shouldProcess := (category == "Terminal não autorizado") ||
				(category == "Horário inválido") ||
				(category == "Fora do perímetro")

			if shouldProcess {
				continue
			} else {
				log.Panicf("[process] inconsistency category must not be different from the filter")
			}
		}

		log.Println("[process] saving results")

		report.NewReport(pr.Results).SaveReport("erros-de-escala")

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
