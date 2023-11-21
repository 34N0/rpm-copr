#  ⏳ rpm-copr

rpm-copr is a Command Line Interface that ports the COPR dnf command to immutable (OSTree) images. This makes it possible for Fedora Silverblue and Fedora Kyonite users to use COPR repositories in a safe and intuitive way.

## Install

```bash
curl -L https://github.com/34N0/rpm-copr/releases/download/v0.9-beta/rpm-copr-v0.9-beta-linux-amd64.tar.gz | sudo tar zx -C /usr/local/bin
```

## Usage
```
Usage:
  rpm-copr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  disable     Disable the copr repository
  enable      Enable the copr repository
  help        Help about any command
  list        List local copr repositories
  remove      Remove the copr repository

Flags:
  -h, --help      help for rpm-copr
  -v, --version   version for rpm-copr
```

## Contributing

Feel free to open issues or pull requests for improvements, bug fixes, or new features. We welcome your contributions!
