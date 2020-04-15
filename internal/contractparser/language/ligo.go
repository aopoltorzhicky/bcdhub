package language

import (
	"strconv"
	"strings"

	"github.com/baking-bad/bcdhub/internal/contractparser/node"
)

type ligo struct{}

func (l ligo) DetectInCode(n node.Node) bool {
	if n.HasAnnots() {
		for _, a := range n.Annotations {
			if len(a) < 2 {
				continue
			}
			if a[0] == '%' && isDigit(a[1:]) {
				return true
			}
		}
	}
	str := n.GetString()
	return hasLIGOKeywords(str)
}

func (l ligo) DetectInParameter(n node.Node) bool {
	return false
}

func isDigit(input string) bool {
	_, err := strconv.ParseUint(input, 10, 32)
	return err == nil
}

func hasLIGOKeywords(s string) bool {
	ligoKeywords := []string{
		"GET_FORCE",
		"get_force",
		"MAP FIND",
	}

	for _, keyword := range ligoKeywords {
		if s == keyword {
			return true
		}
	}

	return strings.Contains(s, "get_entrypoint") || strings.Contains(s, "get_contract")
}
