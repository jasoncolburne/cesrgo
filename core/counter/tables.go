package counter

import (
	"github.com/jasoncolburne/cesrgo"
	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/core/counter/one"
	counter "github.com/jasoncolburne/cesrgo/core/counter/sizage"
	"github.com/jasoncolburne/cesrgo/core/counter/two"
	"github.com/jasoncolburne/cesrgo/core/types"
)

var CounterCodex = map[uint32][]types.Code{
	cesrgo.VERSION_1_0.Major: one.CounterCodex,
	cesrgo.VERSION_2_0.Major: two.CounterCodex,
}

var SpecialUniversalCodex = map[uint32][]types.Code{
	cesrgo.VERSION_1_0.Major: one.SpecialUniversalCodex,
	cesrgo.VERSION_2_0.Major: two.SpecialUniversalCodex,
}

var MessageUniversalCodex = map[uint32][]types.Code{
	cesrgo.VERSION_1_0.Major: one.MessageUniversalCodex,
	cesrgo.VERSION_2_0.Major: two.MessageUniversalCodex,
}

var Sizes = map[uint32]map[types.Code]counter.Sizage{
	cesrgo.VERSION_1_0.Major: one.Sizes,
	cesrgo.VERSION_2_0.Major: two.Sizes,
}

var Hards = map[string]int{}
var Bards = map[[2]byte]int{}

func generateHards() {
	if len(Hards) > 0 {
		return
	}

	for i := 65; i < 65+26; i++ {
		key := "-" + string(byte(i))
		Hards[key] = 2
	}

	for i := 97; i < 97+26; i++ {
		key := "-" + string(byte(i))
		Hards[key] = 2
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
		bard, err := common.CodeB64ToB2(hard)
		if err != nil {
			return err
		}

		nabbed, err := common.NabSextets(bard, 2)
		key := [2]byte{nabbed[0], nabbed[1]}
		Bards[key] = i
	}

	return nil
}

func Hardage(s string) (int, bool) {
	generateHards()

	n, ok := Hards[s]
	return n, ok
}

func Bardage(b [2]byte) (int, bool) {
	err := generateBards()
	if err != nil {
		return -1, false
	}

	n, ok := Bards[b]
	return n, ok
}
