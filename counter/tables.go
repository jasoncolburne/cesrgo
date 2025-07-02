package counter

import (
	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/common/util"
	"github.com/jasoncolburne/cesrgo/counter/one"
	counter "github.com/jasoncolburne/cesrgo/counter/sizage"
	"github.com/jasoncolburne/cesrgo/counter/two"
	"github.com/jasoncolburne/cesrgo/types"
)

var CounterCodex = map[uint32][]types.Code{
	common.VERSION_1_0.Major: one.CounterCodex,
	common.VERSION_2_0.Major: two.CounterCodex,
}

var SpecialUniversalCodex = map[uint32][]types.Code{
	common.VERSION_1_0.Major: one.SpecialUniversalCodex,
	common.VERSION_2_0.Major: two.SpecialUniversalCodex,
}

var MessageUniversalCodex = map[uint32][]types.Code{
	common.VERSION_1_0.Major: one.MessageUniversalCodex,
	common.VERSION_2_0.Major: two.MessageUniversalCodex,
}

var Sizes = map[uint32]map[types.Code]counter.Sizage{
	common.VERSION_1_0.Major: one.Sizes,
	common.VERSION_2_0.Major: two.Sizes,
}

var Hards = map[string]int{}
var Bards = map[string]int{}

func generateHards() {
	if len(Hards) > 0 {
		return
	}

	for i := 65; i < 65+26; i++ {
		Hards["-"+string(byte(i))] = 2
	}

	for i := 97; i < 97+26; i++ {
		Hards["-"+string(byte(i))] = 2
	}

	Hards["--"] = 3
	Hards["-_"] = 5
}

func generateBards() error {
	if len(Bards) > 0 {
		return nil
	}

	generateHards()

	for hard, i := range Hards {
		bard, err := util.CodeB64ToB2(hard)
		if err != nil {
			return err
		}
		Bards[string(bard)] = i
	}

	return nil
}

func Hardage(s string) (int, bool) {
	generateHards()

	n, ok := Hards[s]
	return n, ok
}

func Bardage(b []byte) (int, bool) {
	err := generateBards()
	if err != nil {
		return -1, false
	}

	n, ok := Bards[string(b)]
	return n, ok
}
