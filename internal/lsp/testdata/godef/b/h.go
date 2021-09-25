package b

import . "github.com/kent0106/gotools/internal/lsp/godef/a"

func _() {
	// variable of type a.A
	var _ A //@mark(AVariable, "_"),hoverdef("_", AVariable)

	AStuff() //@hoverdef("AStuff", AStuff)
}
