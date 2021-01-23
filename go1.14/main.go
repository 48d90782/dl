// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.14 command runs the go command from Go 1.14.
//
// To install, run:
//
//     $ go get github.com/48d90782/dl/go1.14
//     $ go1.14 download
//
// And then use the go1.14 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/go1.14
//
// File bugs at https://golang.org/issues/new
package main

import "github.com/48d90782/dl/internal/version"

func main() {
	version.Run("go1.14")
}
