# Kubernetes

## Tips

### Krew

[Krew](https://github.com/kubernetes-sigs/krew/) is known as plugin manager for `kubectl` command.

#### Installation

You can install `krew` with the following command.

```bash
(
  set -x; cd "$(mktemp -d)" &&
  OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/krew.tar.gz" &&
  tar zxvf krew.tar.gz &&
  KREW=./krew-"${OS}_${ARCH}" &&
  "$KREW" install krew
)
```

After that you need to add `$HOME/.krew/bin` directory to your PATH environment variable.

```bash
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
```

Check the installation.

```bash
kubectl krew version
```

#### Useful plugins

- [ctx](https://github.com/ahmetb/kubectx)
- [resource-capacity](https://github.com/robscott/kube-capacity)
- [score](https://github.com/zegl/kube-score)
