#!/bin/bash

set -e

go build

cd /workspaces/gh-schedule
gh extension install .
cd -

# Test that the extension is installed
gh schedule
