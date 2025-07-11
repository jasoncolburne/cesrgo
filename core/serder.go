package cesr

import "github.com/jasoncolburne/cesrgo/core/types"

type Serder struct {
	raw   types.Raw
	sad   types.Map
	proto types.Proto
	pvrsn types.Version
	genus string
	gvrsn types.Version
	kind  types.Kind
	size  types.Size
}
