package winreg

import (
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func Installed(productID string) bool {
	key := fmt.Sprintf(`Software\Microsoft\Windows\CurrentVersion\Uninstall\%s`, productID)
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		log.Fatal(err)
	}
	defer k.Close()

	return true
}
