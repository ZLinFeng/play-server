package initial

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ZLinFeng/play-server/config"
)

const Banner = `            
   / __ \/ /___ ___  __     Os: %s
  / /_/ / / __ \/ / / /     Config: %s
 / ____/ / /_/ / /_/ /      App: %s
/_/   /_/\__,_/\__, /       Port: %d
              /____/        Version: %s
                            User: %s
Running[%d]...
`

func PrintBanner(c *config.Config) {
	username, _ := user.Current()
	fmt.Printf(Banner,
		c.Env.Os,
		c.Env.ConfigPath,
		c.Server.AppName,
		c.Server.Port,
		c.Server.Version,
		username.Username,
		os.Getppid())
}
