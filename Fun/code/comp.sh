#!/usr/bin/env bash

set -e
b=$(g++ "$1")

mv 2>>/dev/null a.out "$2"
