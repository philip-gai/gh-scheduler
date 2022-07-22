#!/bin/bash

set -e

cd /workspaces/gh-scheduler

go build
gh extension remove gh-scheduler &> /dev/null
gh extension install .

cd - > /dev/null
