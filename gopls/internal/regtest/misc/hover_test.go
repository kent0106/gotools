// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package misc

import (
	"strings"
	"testing"

	"github.com/kent0106/gotools/internal/lsp/fake"
	. "github.com/kent0106/gotools/internal/lsp/regtest"
	"github.com/kent0106/gotools/internal/testenv"
)

func TestHoverUnexported(t *testing.T) {
	const proxy = `
-- golang.org/x/structs@v1.0.0/go.mod --
module golang.org/x/structs

go 1.12

-- golang.org/x/structs@v1.0.0/types.go --
package structs

type Mixed struct {
	Exported   int
	unexported string
}

func printMixed(m Mixed) {
	println(m)
}
`
	const mod = `
-- go.mod --
module mod.com

go 1.12

require golang.org/x/structs v1.0.0
-- go.sum --
golang.org/x/structs v1.0.0 h1:3DlrFfd3OsEen7FnCHfqtnJvjBZ8ZFKmrD/+HjpdJj0=
golang.org/x/structs v1.0.0/go.mod h1:47gkSIdo5AaQaWJS0upVORsxfEr1LL1MWv9dmYF3iq4=
-- main.go --
package main

import "golang.org/x/structs"

func main() {
	var _ structs.Mixed
}
`
	// TODO: use a nested workspace folder here.
	WithOptions(
		ProxyFiles(proxy),
	).Run(t, mod, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		mixedPos := env.RegexpSearch("main.go", "Mixed")
		got, _ := env.Hover("main.go", mixedPos)
		if !strings.Contains(got.Value, "unexported") {
			t.Errorf("Workspace hover: missing expected field 'unexported'. Got:\n%q", got.Value)
		}
		cacheFile, _ := env.GoToDefinition("main.go", mixedPos)
		argPos := env.RegexpSearch(cacheFile, "printMixed.*(Mixed)")
		got, _ = env.Hover(cacheFile, argPos)
		if !strings.Contains(got.Value, "unexported") {
			t.Errorf("Non-workspace hover: missing expected field 'unexported'. Got:\n%q", got.Value)
		}
	})
}

func TestHoverIntLiteral(t *testing.T) {
	testenv.NeedsGo1Point(t, 13)
	const source = `
-- main.go --
package main

var (
	bigBin = 0b1001001
)

var hex = 0xe34e

func main() {
}
`
	Run(t, source, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		hexExpected := "58190"
		got, _ := env.Hover("main.go", env.RegexpSearch("main.go", "hex"))
		if got != nil && !strings.Contains(got.Value, hexExpected) {
			t.Errorf("Hover: missing expected field '%s'. Got:\n%q", hexExpected, got.Value)
		}

		binExpected := "73"
		got, _ = env.Hover("main.go", env.RegexpSearch("main.go", "bigBin"))
		if got != nil && !strings.Contains(got.Value, binExpected) {
			t.Errorf("Hover: missing expected field '%s'. Got:\n%q", binExpected, got.Value)
		}
	})
}

// Tests that hovering does not trigger the panic in golang/go#48249.
func TestPanicInHoverBrokenCode(t *testing.T) {
	testenv.NeedsGo1Point(t, 13)
	const source = `
-- main.go --
package main

type Example struct`
	Run(t, source, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		env.Editor.Hover(env.Ctx, "main.go", env.RegexpSearch("main.go", "Example"))
	})
}

func TestHoverRune_48492(t *testing.T) {
	const files = `
-- go.mod --
module mod.com

go 1.18
-- main.go --
package main
`
	Run(t, files, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		env.EditBuffer("main.go", fake.NewEdit(0, 0, 1, 0, "package main\nfunc main() {\nconst x = `\nfoo\n`\n}"))
		env.Editor.Hover(env.Ctx, "main.go", env.RegexpSearch("main.go", "foo"))
	})
}
