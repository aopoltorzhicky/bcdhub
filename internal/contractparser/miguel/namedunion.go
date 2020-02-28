package miguel

import (
	"strings"

	"github.com/aopoltorzhicky/bcdhub/internal/contractparser/meta"
	"github.com/tidwall/gjson"
)

type namedUnionDecoder struct{}

// Decode -
func (l *namedUnionDecoder) Decode(node gjson.Result, path string, nm *meta.NodeMetadata, metadata meta.Metadata) (interface{}, error) {
	res := make(map[string]interface{})
	for i, arg := range nm.Args {
		argPath := strings.TrimPrefix(arg, path+"/")
		unionPath := getGJSONPathUnion(argPath, node)
		argNode := node.Get(unionPath)

		if argNode.Exists() {
			data, err := michelineNodeToMiguel(argNode, arg, metadata)
			if err != nil {
				return nil, err
			}
			name := metadata[arg].GetName(i)
			res[name] = data
			return res, nil
		}
	}

	name := metadata[path].GetName(-1)
	res[name] = nil
	return res, nil
}
