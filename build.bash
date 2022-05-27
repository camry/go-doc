#!/usr/bin/env bash

go build -o godoc godoc.go wire_gen.go && \
upx -qvf ./godoc

exit 0
