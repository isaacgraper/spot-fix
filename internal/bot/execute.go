package bot

import (
	"fmt"
	"log"
	os "os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/isaacgraper/spotfix.git/internal/common/config"
	"github.com/isaacgraper/spotfix.git/internal/page"
	"github.com/isaacgraper/spotfix.git/internal/page/filter"
	"github.com/isaacgraper/spotfix.git/internal/report"
	"github.com/joho/godotenv"
)

type Process struct {
	config  *config.Config
	page    *page.Page
	Results []report.ReportData
}

func NewProcess() *Process {
	return &Process{
		config: &config.Config{},
		page:   &page.Page{},
	}
}

func (pr *Process) Execute(c *config.Config) error {
	path, _ := launcher.LookPath()

	u := launcher.
		New().
		Bin(path).
		Headless(false).
		Set("start-maximized").
		MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect().
		Trace(false)

	defer browser.MustClose()

	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file:", err)
		return err
	}

	pageInstance := browser.MustPage(os.Getenv("URL")).MustWaitLoad()

	pr.page = &page.Page{
		Rod: pageInstance,
	}

	if err := pr.page.Login(c.NewCredential()); err != nil {
		log.Printf("[execute] login failed: %v", err)
		return nil
	}

	if err := pr.page.NavigateToInconsistencies(); err != nil {
		log.Printf("[execute] navigate to inconsistencies failed: %v", err)
		return nil
	}

	if c.NotRegistered {
		ok, err := filter.FilterNotRegistered(pr.page)
		if err != nil {
			return fmt.Errorf("[execute] error while trying to filter: %w", err)
		}

		if !ok {
			log.Println("[execute] filtering failed")
			log.Println("[execute] ending process with filter...")
			return nil
		}

		log.Println("[execute] starting process with notRegistered filter...")
		pr.ProcessNotRegistered()
	}

	if c.WorkSchedule {
		ok, err := filter.FilterWorkSchedule(pr.page)
		if err != nil {
			return fmt.Errorf("[execute] error while trying to filter: %w", err)
		}

		if !ok {
			log.Println("[execute] filtering failed")
			log.Println("[execute] ending process with filter...")
			return nil
		}

		log.Println("[execute] starting process with workSchedule filter...")
		pr.ProcessWorkSchedule()
	}

	if c.ProcessBatch {
		log.Println("[execute] starting process with batch...")
		pr.ProcessHandler(c)
	}
	return nil
}
