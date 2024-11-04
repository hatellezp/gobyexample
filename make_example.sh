#!/bin/sh

EXAMPLE=$1


if [ -z "$2" ]; then
    EDITOR=/home/tellez/.local/bin/nvim.appimage
else
    EDITOR=$2
fi

mkdir src/$EXAMPLE
cd src/$EXAMPLE
go mod init gobyexample/$EXAMPLE
touch $EXAMPLE.go

echo 'package main \n\n import (\n\t"fmt"\n)\n\nfunc main() {\n\tfmt.Println("Hello World!");\n}\n' > $EXAMPLE.go

$EDITOR $EXAMPLE.go
