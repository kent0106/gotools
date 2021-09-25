// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The lostcancel command applies the github.com/kent0106/gotools/go/analysis/passes/lostcancel
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/kent0106/gotools/go/analysis/passes/lostcancel"
	"github.com/kent0106/gotools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
