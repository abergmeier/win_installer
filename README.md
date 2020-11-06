![Test & Build](https://github.com/abergmeier/winsible/workflows/Test%20&%20Build/badge.svg)

# winsible
Installer similar to Ansible but running on Windows!

```cmd
winsible.exe --tasks C:\temp\tasks.yaml
```
Supported Modules:
- ansible.builtin.git
- ansible.builtin.unarchive
- ansible.windows.win_package
- community.general.gc_storage

Some more facts:
- Runs local only - no remote exec
- Written in Go
- No dependencies
- Runs on Wine
