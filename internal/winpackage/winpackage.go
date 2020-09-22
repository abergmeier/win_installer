package winpackage

import (
	"os/exec"
	"strings"

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

	return run(path, arguments...)

}

func run(path string, arguments ...string) error {

	if strings.HasSuffix(strings.ToLower(path), ".msi") {
		return runMsi(path, arguments...)
	}

	c := exec.Command(path, arguments...)
	return c.Run()
}

func runMsi(path string, arguments ...string) error {
	arguments = append([]string{"/i", path, "/quiet"}, arguments...)
	c := exec.Command("msiexec", arguments...)
	return c.Run()
}
