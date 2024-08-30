#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <number>"
  exit 1
fi

if ! [[ "$1" =~ ^[0-9]+$ ]]; then
  echo "Error: The parameter must be a number."
  exit 1
fi

NUMBER=$(printf "%03d" "$1")
DIRECTORY="exercise${NUMBER}"

if [ -d "$DIRECTORY" ]; then
  echo "Directory $DIRECTORY already exists."
  exit 1
fi

mkdir "$DIRECTORY"
cd "$DIRECTORY"

GO_MOD_PATH="github.com/asebexen/go-exercises/exercise${NUMBER}"
go mod init "$GO_MOD_PATH"

cat <<EOF > makefile
build:
	go build -o out
EOF

cat <<EOF > main.go
package main

func main() {

}
EOF

echo "Setup complete."