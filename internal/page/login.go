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
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("[login] error loading .env file: %w", err)
	}

	c.Username = os.Getenv("USERNAME")
	c.Password = os.Getenv("PASSWORD")

	if c.Username == "" || c.Password == "" {
		return fmt.Errorf("[login] password and username must be set in environment variables")
	}

	_ = proto.NetworkSetCacheDisabled{CacheDisabled: true}.Call(p.Rod)
	_ = proto.NetworkClearBrowserCache{}.Call(p.Rod)

	p.Rod.Reload()

	p.Loading()

	name, err := p.Rod.Element("#inputUsername")
	if err != nil {
		return fmt.Errorf("[login] error finding element: %w", err)
	}

	name.MustInput("bot@icop").MustType(input.Tab)

	pwd, err := p.Rod.Element("#inputPassword")
	if err != nil {
		return fmt.Errorf("[login] error finding element: %w", err)
	}

	pwd.MustInput(c.Password).MustType(input.Enter)

	p.Loading()

	log.Println("[login] logged in successfully!")

	return nil
}
