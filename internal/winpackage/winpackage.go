package winpackage

import (
	"log"
	"os/exec"

	"github.com/abergmeier/winsible/internal/winreg"
	"github.com/gonuts/go-shellquote"
)

func isProductIDInstalled(config map[string]interface{}) bool {
	productID := ""
	productConfig, ok := config["product_id"]
	if !ok || productConfig == nil {
		return false
	}
	productID = productConfig.(string)

	if productID == "" {
		return false
	}

	return winreg.Installed(productID)
}

func Run(config map[string]interface{}) error {
	path := config["path"].(string)

	if isProductIDInstalled(config) {
		return nil
	}

	argumentString := config["arguments"].(string)

	arguments, err := shellquote.Split(argumentString)
	if err != nil {
		log.Fatal(err)
	}

	c := exec.Command(path, arguments...)

	return c.Run()
}
