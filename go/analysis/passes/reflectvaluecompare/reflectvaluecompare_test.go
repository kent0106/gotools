// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectvaluecompare_test

import (
	"testing"

	"github.com/kent0106/gotools/go/analysis/analysistest"
	"github.com/kent0106/gotools/go/analysis/passes/reflectvaluecompare"
)

func TestReflectValueCompare(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, reflectvaluecompare.Analyzer, "a")
}
