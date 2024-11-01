package bot

import (
	"log"
	os "os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/isaacgraper/spotfix.git/internal/common/config"
	"github.com/isaacgraper/spotfix.git/internal/page"
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
	u := launcher.New().Bin(path).Headless(false).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect().Trace(true)

	defer browser.MustClose()

	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file:", err)
		return err
	}

	pageInstance := browser.MustPage(os.Getenv("URL")).MustWaitLoad()

	pr.page = &page.Page{
		Page: pageInstance,
	}

	if err := pr.page.Login(c.NewCredential()); err != nil {
		log.Printf("login failed: %v", err)
		return nil
	}

	if err := pr.page.NavigateToInconsistencies(); err != nil {
		log.Printf("navigate to inconsistencies failed: %v", err)
		return nil
	}

	if c.Filter {
		ok, err := pr.page.Filter()
		if err != nil {
			return nil
		}
		if !ok {
			log.Println("[process] filtering failed")
			log.Println("[process] ending process with filter...")
			return nil
		}

		log.Println("[process] starting processor with filter...")
		pr.ProcessFilter(c)
	}

	if !c.Filter {
		log.Println("[process] starting processor with batch...")
		pr.ProcessHandler(c)
	}
	return nil
}
