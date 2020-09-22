package winpackage

import (
	"os/exec"

	"github.com/abergmeier/winsible/internal/winreg"
	"github.com/gonuts/go-shellquote"
)

func argumentsFromConfig(config map[string]interface{}) ([]string, error) {
	argumentConfig, ok := config["arguments"]
	if !ok || argumentConfig == nil {
		return []string{}, nil
	}
	argumentString := argumentConfig.(string)
	if argumentString == "" {
		return []string{}, nil
	}

	arguments, err := shellquote.Split(argumentString)
	return arguments, err
}

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

	arguments, err := argumentsFromConfig(config)
	if err != nil {
		return err
	}

	c := exec.Command(path, arguments...)

	return c.Run()
}
