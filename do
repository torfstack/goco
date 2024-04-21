#! /bin/bash

function task_lint() {
  echo "Running linter..."
  golangci-lint run
}

function check_installed() {
  if ! command -v "$1" &> /dev/null; then
    echo "Command $1 not found. Please install it."
    exit 1
  fi
}

check_installed golangci-lint

if [ $# -eq 0 ]; then
  echo "Usage: $0 <task>"
  echo "Tasks:"
  echo "  lint"
  exit 1
fi

case $1 in
  lint)
    task_lint
    ;;
  *)
    echo "Unknown task: $1"
    exit 1
    ;;
esac
