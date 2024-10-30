package page

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/isaacgraper/spotfix.git/internal/common/config"
	"github.com/joho/godotenv"
)

func (p *Page) Login(c *config.Credential) error {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file:", err)
		return err
	}

	c.Username = os.Getenv("USERNAME")
	c.Password = os.Getenv("PASSWORD")

	if c.Username == "" || c.Password == "" {
		return fmt.Errorf("password and username must be set in environment variables")
	}

	_ = proto.NetworkSetCacheDisabled{CacheDisabled: true}.Call(p.Page)
	_ = proto.NetworkClearBrowserCache{}.Call(p.Page)

	p.Page.Reload()

	name, err := p.Page.Element("#inputUsername")
	if err != nil {
		log.Printf("Error finding element: %v\n", err)
		return err
	}
	// hardcode is not the usual way to do this
	name.MustInput("Jorge").MustType(input.Tab)

	pwd, err := p.Page.Element("#inputPassword")
	if err != nil {
		log.Printf("Error finding element: %v\n", err)
		return err
	}
	pwd.MustInput(c.Password).MustType(input.Enter)

	p.Loading()

	log.Println("[login] bot logged in successfully")

	return nil
}
