# dockutil
general docker utilities

## Setup

```
mkdir -p ~/goworkspace/bin ~/goworkspace/src ~/goworkspace/pkg
cd ~/goworkspace

# add go settings. you can add this to your .bash_profile file
export GOPATH=$HOME/goworkspace
export PATH=$PATH:$GOPATH/bin

cd $GOPATH
mkdir -p src/github.com/unicredit
cd src/github.com/unicredit
git clone git@github.com:unicredit/dockutil.git

```

Installed packages utils

```
go get -u github.com/golang/lint/golint
go get -u sourcegraph.com/sqs/goreturns
go get -u github.com/nsf/gocode #autocomplete
go get -u github.com/rogpeppe/godef
go get github.com/tpng/gopkgs
```