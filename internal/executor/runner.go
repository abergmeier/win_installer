package executor

import (
	"fmt"
	"strings"

	"github.com/abergmeier/winsible/internal/gcstorage"
	"github.com/abergmeier/winsible/internal/winpackage"
)

func MustRun(config []interface{}) {

	maxNameLength := 0

	for _, c := range config {
		taskConfig := c.(map[string]interface{})
		newLength := len(taskConfig["name"].(string))
		if newLength > maxNameLength {
			maxNameLength = newLength
		}
	}

	for _, c := range config {
		taskConfig := c.(map[string]interface{})

		var vConfig map[string]interface{}
		var run func(map[string]interface{})

		for k, v := range taskConfig {
			if k == "gc_storage" {
				run = gcstorage.Run
			} else if k == "win_package" {
				run = winpackage.Run
			}

			if run != nil {
				vConfig = v.(map[string]interface{})
				break
			}
		}

		name := taskConfig["name"].(string)
		fillCount := maxNameLength - len(name)
		fmt.Printf("TASK [%s] ***%s\n", name, strings.Repeat("*", fillCount))
		run(vConfig)
	}

}
