package winpackage

import (
	"log"
	"os/exec"

	"github.com/abergmeier/winsible/internal/winreg"
	"github.com/gonuts/go-shellquote"
)

func Run(config map[string]interface{}) error {
	path := config["path"].(string)
	productID := config["product_id"].(string)
	argumentString := config["arguments"].(string)

	if winreg.Installed(productID) {
		return nil
	}

	arguments, err := shellquote.Split(argumentString)
	if err != nil {
		log.Fatal(err)
	}

	c := exec.Command(path, arguments...)

	return c.Run()
}
