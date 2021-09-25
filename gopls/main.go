// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gopls (pronounced “go please”) is an LSP server for Go.
// The Language Server Protocol allows any text editor
// to be extended with IDE-like features;
// see https://langserver.org/ for details.
//
// See https://github.com/golang/tools/blob/master/gopls/README.md
// for the most up-to-date documentation.
package main // import "github.com/kent0106/gotools/gopls"

import (
	"context"
	"os"

	"github.com/kent0106/gotools/gopls/internal/hooks"
	"github.com/kent0106/gotools/internal/lsp/cmd"
	"github.com/kent0106/gotools/internal/tool"
)

func main() {
	ctx := context.Background()
	tool.Main(ctx, cmd.New("gopls", "", nil, hooks.Options), os.Args[1:])
}
