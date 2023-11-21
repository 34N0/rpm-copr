#  ⏳ rpm-copr

⚠️ This utility serves as a workaround until the functionality is implemented in ```rpm-otsree```.

## Install

```bash
curl -L https://github.com/34N0/rpm-copr/releases/download/v0.8-alpha/rpm-copr-v0.8-alpha-linux-amd64.tar.gz | sudo tar zx -C /usr/local/bin
```

## Usage
```bash
rpm-copr is a Command Line Interface that ports the COPR dnf plugin to immutable (OSTree) images.

Usage:
  rpm-copr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  disable     Disable the copr repository
  enable      Enable the copr repository
  help        Help about any command
  remove      Remove the copr repository

Flags:
  -h, --help   help for rpm-copr

Use "rpm-copr [command] --help" for more information about a command.
```

### Contributing

Feel free to open issues or pull requests for improvements, bug fixes, or new features. We welcome your contributions!
