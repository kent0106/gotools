// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplifyslice_test

import (
	"testing"

	"github.com/kent0106/gotools/go/analysis/analysistest"
	"github.com/kent0106/gotools/internal/lsp/analysis/simplifyslice"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, simplifyslice.Analyzer, "a")
}
