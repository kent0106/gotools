// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hooks

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/kent0106/gotools/internal/lsp/diff"
	"github.com/kent0106/gotools/internal/span"
)

func ComputeEdits(uri span.URI, before, after string) (edits []diff.TextEdit, err error) {
	// The go-diff library has an unresolved panic (see golang/go#278774).
	// TODO(rstambler): Remove the recover once the issue has been fixed
	// upstream.
	defer func() {
		if r := recover(); r != nil {
			edits = nil
			err = fmt.Errorf("unable to compute edits for %s: %s", uri.Filename(), r)
		}
	}()
	diffs := diffmatchpatch.New().DiffMain(before, after, true)
	edits = make([]diff.TextEdit, 0, len(diffs))
	offset := 0
	for _, d := range diffs {
		start := span.NewPoint(0, 0, offset)
		switch d.Type {
		case diffmatchpatch.DiffDelete:
			offset += len(d.Text)
			edits = append(edits, diff.TextEdit{Span: span.New(uri, start, span.NewPoint(0, 0, offset))})
		case diffmatchpatch.DiffEqual:
			offset += len(d.Text)
		case diffmatchpatch.DiffInsert:
			edits = append(edits, diff.TextEdit{Span: span.New(uri, start, span.Point{}), NewText: d.Text})
		}
	}
	return edits, nil
}
