#! /bin/bash

# default ENV is dev
env=dev

while test $# -gt 0; do
    case "$1" in
        -env)
            shift
            if test $# -gt 0; then
                env=$1
            fi
            # shift
            ;;
        *)
        break
        ;;
    esac
done

cd ../../blog
source .env
go build -o cmd/go-blog/go-blog cmd/go-blog/main.go
cmd/go-blog/go-blog -env $env &