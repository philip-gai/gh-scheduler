#!/bin/bash

set -e

cd /workspaces/gh-schedule

go build
gh extension remove gh-schedule &> /dev/null
gh extension install .

cd - > /dev/null
