package executor

import (
	"fmt"

	"github.com/abergmeier/winsible/internal/gcstorage"
)

func MustRun(config []interface{}) {

	for _, c := range config {
		taskConfig := c.(map[string]interface{})

		var vConfig map[string]interface{}

		for k, v := range taskConfig {
			if k == "gc_storage" {
				vConfig = v.(map[string]interface{})
			}

			if vConfig != nil {
				break
			}
		}

		fmt.Printf("TASK [%s] ***\n", taskConfig["name"])
		gcstorage.Run(vConfig)
	}

}
