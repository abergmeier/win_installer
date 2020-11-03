![Test & Build](https://github.com/abergmeier/winsible/workflows/Test%20&%20Build/badge.svg)

# winsible
Installer similar to Ansible but running on Windows!

```cmd
winsible.exe --tasks C:\temp\tasks.yaml
```
Supported Modules:
- gc_storage
- git
- unarchive
- win_package

Some more facts:
- Runs local only - no remote exec
- Written in Go
- No dependencies
- Runs on Wine
