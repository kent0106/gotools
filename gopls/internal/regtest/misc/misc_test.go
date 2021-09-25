// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package misc

import (
	"testing"

	"github.com/kent0106/gotools/gopls/internal/hooks"
	"github.com/kent0106/gotools/internal/lsp/regtest"
)

func TestMain(m *testing.M) {
	regtest.Main(m, hooks.Options)
}
