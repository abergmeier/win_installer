package config

import (
	"reflect"
	"strings"
	"testing"
)

var (
	input = `
- name: Download PowerDesigner
  gc_storage:
    bucket: bar
    object: PowerDesigner165SP04x64.exe
    dest:   'C:\temp\PowerDesigner165SP04x64.exe'
    mode:   get
- name: Download pd.iss
  gc_storage:
    bucket: foo
    object: pd.iss
    dest:   'C:\temp\pd.iss'
    mode:   get
- name: Install PowerDesigner
  win_package:
    path:       'C:\temp\PowerDesigner165SP04x64.exe'
    product_id: '{D174290F-9A4E-48E3-9EB5-1B6A8AB67E9B}'
    arguments:  /s /f1'C:\temp\pd.iss'
`

	downloadPowerDesigner = map[string]interface{}{
		"name": "Download PowerDesigner",
		"gc_storage": map[string]interface{}{
			"bucket": "bar",
			"object": "PowerDesigner165SP04x64.exe",
			"dest":   `C:\temp\PowerDesigner165SP04x64.exe`,
			"mode":   "get",
		},
	}

	downloadIss = map[string]interface{}{
		"name": "Download pd.iss",
		"gc_storage": map[string]interface{}{
			"bucket": "foo",
			"object": "pd.iss",
			"dest":   `C:\temp\pd.iss`,
			"mode":   "get",
		},
	}

	installPowerDesigner = map[string]interface{}{
		"name": "Install PowerDesigner",
		"win_package": map[string]interface{}{
			"path":       `C:\temp\PowerDesigner165SP04x64.exe`,
			"product_id": "{D174290F-9A4E-48E3-9EB5-1B6A8AB67E9B}",
			"arguments":  `/s /f1'C:\temp\pd.iss'`,
		},
	}

	expectedTasks = []interface{}{
		downloadPowerDesigner,
		downloadIss,
		installPowerDesigner,
	}
)

func TestParsing(t *testing.T) {
	r := strings.NewReader(input)
	tasks := ReadYaml(r)

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Fatalf("Reading config failed:\nExpected:\n%#v\nActual:\n%#v\n", expectedTasks, tasks)
	}
}
