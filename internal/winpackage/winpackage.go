package winpackage

import (
	"log"
	"os/exec"

	"github.com/abergmeier/winsible/internal/winreg"
)

func Run(config map[string]interface{}) {
	path := config["path"].(string)
	productID := config["product_id"].(string)
	arguments := config["arguments"].(string)

	if winreg.Installed(productID) {
		return
	}

	c := exec.Command(path, arguments)

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
