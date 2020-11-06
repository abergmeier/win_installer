module github.com/abergmeier/winsible

go 1.14

replace github.com/go-git/go-git/v5 v5.2.0 => github.com/abergmeier/go-git/v5 v5.0.0-20201106103301-6be754ba63e9

require (
	cloud.google.com/go/storage v1.11.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/gonuts/go-shellquote v0.0.0-20180428030007-95032a82bc51
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/sys v0.0.0-20200917073148-efd3b9a0ff20
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
