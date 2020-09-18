package executor

import (
	"fmt"

	"github.com/abergmeier/winsible/internal/gcstorage"
	"github.com/abergmeier/winsible/internal/winpackage"
)

func MustRun(config []interface{}) {

	for _, c := range config {
		taskConfig := c.(map[string]interface{})

		var vConfig map[string]interface{}
		var run func(map[string]interface{})

		for k, v := range taskConfig {
			vConfig = v.(map[string]interface{})
			if k == "gc_storage" {
				run = gcstorage.Run
			} else if k == "win_package" {
				run = winpackage.Run
			}

			if run != nil {
				break
			}
		}

		fmt.Printf("TASK [%s] ***\n", taskConfig["name"])
		run(vConfig)
	}

}
