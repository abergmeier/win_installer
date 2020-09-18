# winsible
Installer similar to Ansible but running on Windows!

```cmd
winsible.exe --tasks C:\temp\tasks.yaml
```
Supported Modules:
- gc_storage
- win_package

Some more facts:
- Runs local only - no remote exec
- Written in Go
- No dependencies
