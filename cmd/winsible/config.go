package main

import (
	"strings"

	"github.com/abergmeier/winsible/internal/config"
)

const (
	taskConfig = `
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
)

func mustReadConfig(bucket string) []interface{} {

	f := strings.NewReader(taskConfig)
	return config.ReadYaml(f)
}
