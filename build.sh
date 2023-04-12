#!/usr/bin/env bash
# script for building weather-go binaries

case $1 in
  tui)
    echo "----- Building tui binary -----"
    cd cmd/tui/
    go build -v -o ../../bin/tui
    echo "----- Done building tui binary -----"
  ;;

  widget)
    echo "----- Building widget binary -----"
    cd cmd/widget/
    go build -o ../../bin/widget
    echo "----- Done building widget binary -----"
  ;;

  *)
    printf "Script for building weather-go binaries.\nUsage:\n\tbash build.sh widget | tui\n"
  ;;

esac