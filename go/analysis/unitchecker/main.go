// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This file provides an example command for static checkers
// conforming to the github.com/kent0106/gotools/go/analysis API.
// It serves as a model for the behavior of the cmd/vet tool in $GOROOT.
// Being based on the unitchecker driver, it must be run by go vet:
//
//   $ go build -o unitchecker main.go
//   $ go vet -vettool=unitchecker my/project/...
//
// For a checker also capable of running standalone, use multichecker.
package main

import (
	"github.com/kent0106/gotools/go/analysis/unitchecker"

	"github.com/kent0106/gotools/go/analysis/passes/asmdecl"
	"github.com/kent0106/gotools/go/analysis/passes/assign"
	"github.com/kent0106/gotools/go/analysis/passes/atomic"
	"github.com/kent0106/gotools/go/analysis/passes/bools"
	"github.com/kent0106/gotools/go/analysis/passes/buildtag"
	"github.com/kent0106/gotools/go/analysis/passes/cgocall"
	"github.com/kent0106/gotools/go/analysis/passes/composite"
	"github.com/kent0106/gotools/go/analysis/passes/copylock"
	"github.com/kent0106/gotools/go/analysis/passes/errorsas"
	"github.com/kent0106/gotools/go/analysis/passes/httpresponse"
	"github.com/kent0106/gotools/go/analysis/passes/loopclosure"
	"github.com/kent0106/gotools/go/analysis/passes/lostcancel"
	"github.com/kent0106/gotools/go/analysis/passes/nilfunc"
	"github.com/kent0106/gotools/go/analysis/passes/printf"
	"github.com/kent0106/gotools/go/analysis/passes/shift"
	"github.com/kent0106/gotools/go/analysis/passes/stdmethods"
	"github.com/kent0106/gotools/go/analysis/passes/structtag"
	"github.com/kent0106/gotools/go/analysis/passes/tests"
	"github.com/kent0106/gotools/go/analysis/passes/unmarshal"
	"github.com/kent0106/gotools/go/analysis/passes/unreachable"
	"github.com/kent0106/gotools/go/analysis/passes/unsafeptr"
	"github.com/kent0106/gotools/go/analysis/passes/unusedresult"
)

func main() {
	unitchecker.Main(
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		errorsas.Analyzer,
		httpresponse.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		stdmethods.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)
}
