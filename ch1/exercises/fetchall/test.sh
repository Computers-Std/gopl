#!/usr/bin/env bash

# NOTE: Don't end the url with '/', because the filename is based on
# the its URL
go run main.go https://go.dev http://eloquentjavascript.net http://example.com
