package winreg

import (
	"errors"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func Installed(productId string) bool {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Uninstall`, registry.QUERY_VALUE)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		log.Fatal(err)
	}
	defer k.Close()

	return true
}
