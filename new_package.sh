#!/bin/bash

package_name=$1

mkdir $package_name

cat <<EOF > $package_name/main.go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("go")
}
EOF

code $package_name/main.go