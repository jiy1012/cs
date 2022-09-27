#!/usr/bin/env bash

 go build -o ./Release/cs ./

 tar -czf cs-mac.tar.gz Release

 shasum -a 256 cs-mac.tar.gz | awk '{print $1}' > cs-mac-sha256.txt