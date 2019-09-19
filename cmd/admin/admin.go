/*
 * Copyright 2017 caicloud authors. All rights reserved.
 */

package main

import (
	"fmt"

	"github.com/caicloud/go-common/version"
	"github.com/caicloud/golang-template-project/pkg/utils/net"
)

func main() {
	fmt.Println("Hello world admin")
	fmt.Printf("XXX build information: %v\n", version.Get().Pretty())

	net.Helper()
}
