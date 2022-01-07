export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin

if [[ $PATH != *$GOPATH* ]]; then
    export PATH="${GOPATH}/bin:${PATH}"
fi

if [[ $PATH != *$GOROOT* ]]; then
    export PATH="${GOROOT}/bin:${PATH}"
fi

export OS=OSX
