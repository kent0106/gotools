// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package inspect defines an Analyzer that provides an AST inspector
// (github.com/kent0106/gotools/go/ast/inspector.Inspector) for the syntax trees
// of a package. It is only a building block for other analyzers.
//
// Example of use in another analysis:
//
//	import (
//		"github.com/kent0106/gotools/go/analysis"
//		"github.com/kent0106/gotools/go/analysis/passes/inspect"
//		"github.com/kent0106/gotools/go/ast/inspector"
//	)
//
//	var Analyzer = &analysis.Analyzer{
//		...
//		Requires:       []*analysis.Analyzer{inspect.Analyzer},
//	}
//
// 	func run(pass *analysis.Pass) (interface{}, error) {
// 		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
// 		inspect.Preorder(nil, func(n ast.Node) {
// 			...
// 		})
// 		return nil
// 	}
//
package inspect

import (
	"reflect"

	"github.com/kent0106/gotools/go/analysis"
	"github.com/kent0106/gotools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:             "inspect",
	Doc:              "optimize AST traversal for later passes",
	Run:              run,
	RunDespiteErrors: true,
	ResultType:       reflect.TypeOf(new(inspector.Inspector)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	return inspector.New(pass.Files), nil
}
