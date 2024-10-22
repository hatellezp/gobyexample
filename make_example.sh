#!/bin/sh

EXAMPLE=$1
EDITOR=$2

mkdir src/$EXAMPLE
cd src/$EXAMPLE
go mod init gobyexample/$EXAMPLE
touch $EXAMPLE.go

echo 'package main \n\n import (\n\t"fmt"\n)\n\nfunc main() {\n\tfmt.Println("Hello World!");\n}\n' > $EXAMPLE.go

$EDITOR $EXAMPLE.go
