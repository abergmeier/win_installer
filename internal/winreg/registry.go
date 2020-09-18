package winreg

import (
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func Installed(productID string) bool {
	keys := []string{
		fmt.Sprintf(`Software\Microsoft\Windows\CurrentVersion\Uninstall\%s`, productID),
		fmt.Sprintf(`Software\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall\%s`, productID),
	}
	for _, key := range keys {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			log.Fatal(err)
		}
		k.Close()

		return true
	}

	return false
}
