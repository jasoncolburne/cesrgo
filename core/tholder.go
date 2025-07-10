package cesr

import (
	"encoding/json"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Tholder struct {
	thold    any
	weighted bool
	size     types.Size
	satisfy  func(t *Tholder, indices []types.Index) bool
	number   *Number
	bexter   *Bexter
}

func (t *Tholder) Satisfy(indices []types.Index) bool {
	return t.satisfy(t, indices)
}

func (t *Tholder) Weighted() bool {
	return t.weighted
}

func (t *Tholder) Thold() any {
	return t.thold
}

func (t *Tholder) Limen() (types.Qb64, error) {
	if t.bexter != nil {
		qb64, err := t.bexter.Qb64()
		if err != nil {
			return types.Qb64(""), err
		}

		return qb64, nil
	}

	if t.number != nil {
		qb64, err := t.number.Qb64()
		if err != nil {
			return types.Qb64(""), err
		}

		return qb64, nil
	}

	return types.Qb64(""), fmt.Errorf("no limen")
}

func (t *Tholder) Sith() (any, error) {
	var sith any

	if t.weighted {
		tholdList, ok := t.thold.([]any)
		if !ok {
			return "", fmt.Errorf("thold is not a list")
		}

		sithList := []any{}
		for _, c := range tholdList {
			clauseList, ok := c.([]any)
			if !ok {
				return "", fmt.Errorf("clause is not a list")
			}

			clause := []any{}
			for _, e := range clauseList {
				switch eType := e.(type) {
				case map[string]any:
					if len(eType) != 1 {
						return "", fmt.Errorf("nested clause too long")
					}

					for k, _v := range eType {
						vList, ok := _v.([]any)
						if !ok {
							return "", fmt.Errorf("invalid clause, value is a %T, not a list", _v)
						}

						v := []any{}
						for _, _v := range vList {
							s, ok := _v.(*big.Rat)
							if !ok {
								return "", fmt.Errorf("invalid clause, value is a %T, not a rational", _v)
							}

							v = append(v, extractIntOrRational(s, "/"))
						}

						clause = append(clause, map[string]any{k: v})
					}
				case *big.Rat:
					clause = append(clause, extractIntOrRational(eType, "/"))
				default:
					return "", fmt.Errorf("invalid clause, member is a %t", eType)
				}
			}
			sithList = append(sithList, clause)
		}

		if len(sithList) == 1 {
			sith = sithList[0]
		} else {
			sith = sithList
		}
	} else {
		tholdInt, ok := t.thold.(int)
		if !ok {
			return "", fmt.Errorf("thold is not an int")
		}

		sith = fmt.Sprintf("%x", tholdInt)
	}

	return sith, nil
}

func (t *Tholder) Size() types.Size {
	return t.size
}

func extractIntOrRational(r *big.Rat, sep string) string {
	if r.Cmp(big.NewRat(1, 1)) >= 0 || r.Cmp(big.NewRat(0, 1)) == 0 {
		value, _ := r.Float64()
		return fmt.Sprintf("%d", int(value))
	} else {
		return fmt.Sprintf("%d%s%d", r.Num(), sep, r.Denom())
	}
}

func NewTholder(thold any, limen *types.Qb64, sith any, opts ...options.MatterOption) (*Tholder, error) {
	t := &Tholder{}

	if thold != nil && limen == nil && sith == nil {
		err := t.processThold(thold)
		if err != nil {
			return nil, err
		}
	} else if thold == nil && limen != nil && sith == nil {
		err := t.processLimen(*limen, opts...)
		if err != nil {
			return nil, err
		}
	} else if thold == nil && limen == nil && sith != nil {
		sithStr, ok := sith.(string)
		if ok {
			if sithStr == "" {
				return nil, fmt.Errorf("empty sith")
			}

			sith = sithStr
		}

		err := t.processSith(sith)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("invalid input")
	}

	return t, nil
}

func (t *Tholder) processThold(thold any) error {
	switch x := thold.(type) {
	case int:
		return t.processUnweighted(*big.NewInt(int64(x)))
	case big.Int:
		return t.processUnweighted(x)
	case []any:
		return t.processWeighted(x)
	default:
		return fmt.Errorf("invalid thold type %T", x)
	}
}

func (t *Tholder) processLimen(limen types.Qb64, opts ...options.MatterOption) error {
	m := &matter{}
	opts = append(opts, options.WithQb64(limen))
	err := NewMatter(m, opts...)
	if err != nil {
		return err
	}

	if slices.Contains(codex.NumCodex, m.GetCode()) {
		number, err := NewNumber(nil, nil, options.WithCode(m.GetCode()), options.WithRaw(m.GetRaw()))
		if err != nil {
			return err
		}

		bigNum := number.Number()
		return t.processUnweighted(bigNum)
	} else if slices.Contains(codex.BextCodex, m.GetCode()) {
		bexter, err := NewBexter(nil, options.WithCode(m.GetCode()), options.WithRaw(m.GetRaw()))
		if err != nil {
			return err
		}

		b, err := bexter.Bext()
		if err != nil {
			return err
		}

		b = strings.ReplaceAll(b, "s", "/")
		parts := strings.Split(b, "a")
		clauses := [][]string{}
		for _, c := range parts {
			clauses = append(clauses, strings.Split(c, "c"))
		}

		thold := []any{}
		for _, c := range clauses {
			clause := []any{}
			for _, e := range c {
				i := strings.Index(e, "k")
				if i != -1 {
					k := e[:i]
					v := []any{}
					weights := strings.SplitSeq(e[i+1:], "v")

					for w := range weights {
						r, err := weight(w)
						if err != nil {
							return err
						}
						v = append(v, r)
					}

					clause = append(clause, map[string]any{k: v})
				} else {
					r, err := weight(e)
					if err != nil {
						return err
					}

					clause = append(clause, r)
				}
			}

			thold = append(thold, clause)
		}

		return t.processWeighted(thold)
	} else {
		return fmt.Errorf("invalid code for limen")
	}
}

func (t *Tholder) processSith(sith any) error {
	var err error

	switch x := sith.(type) {
	case int:
		err = t.processUnweighted(*big.NewInt(int64(x)))
	case big.Int:
		err = t.processUnweighted(x)
	case *big.Int:
		err = t.processUnweighted(*x)
	case string:
		if !strings.Contains(x, "[") {
			thold, err := strconv.ParseInt(x, 16, 64)
			if err != nil {
				return err
			}

			bigNum := big.NewInt(thold)
			err = t.processUnweighted(*bigNum)
		} else {
			var thold []any

			err = json.Unmarshal([]byte(x), &thold)
			if err != nil {
				return err
			}

			if len(thold) == 0 {
				return fmt.Errorf("empty weight list")
			}

			err = t.processInternalThold(thold)
		}
	case []any:
		err = t.processInternalThold(x)
	default:
		return fmt.Errorf("invalid sith type %T", sith)
	}

	return err
}

func (t *Tholder) processInternalThold(thold []any) error {
	sith := thold
	mask := []bool{}
	for _, potentialSequence := range thold {
		mask = append(mask, isSequence(potentialSequence))
	}
	if slices.Contains(mask, false) {
		sith = []any{sith}
	}

	mask = []bool{}
	for _, c := range sith {
		slice, ok := c.([]any)
		if !ok {
			return fmt.Errorf("c is a %T, not a sequence", c)
		}

		for _, e := range slice {
			_, isString := e.(string)
			_, isMap := e.(map[string]any)

			mask = append(mask, isString || isMap)
		}

	}

	if slices.Contains(mask, false) {
		return fmt.Errorf("invalid sith - some weights in %+v are non strings", sith)
	}

	thold = []any{}

	for _, c := range sith {
		clause := []any{}

		cList, ok := c.([]any)
		if !ok {
			return fmt.Errorf("c is a %T, not a sequence", c)
		}
		for _, e := range cList {
			switch _e := e.(type) {
			case map[string]any:
				if len(_e) != 1 {
					return fmt.Errorf("nested clause too long")
				}
				for k, v := range _e {
					vList, ok := v.([]any)
					if !ok {
						return fmt.Errorf("invalid sith, nested value not a list")
					}

					toAppend := []any{}

					for _, v := range vList {
						s, ok := v.(string)
						if !ok {
							return fmt.Errorf("invalid sith, nested value not a string")
						}

						w, err := weight(s)
						if err != nil {
							return fmt.Errorf("invalid sith, nested value not a weight")
						}
						toAppend = append(toAppend, w)
					}
					subClause := map[string]any{k: toAppend}
					clause = append(clause, subClause)
				}
			default:
				s, ok := e.(string)
				if !ok {
					return fmt.Errorf("invalid sith, weight not a string")
				}
				w, err := weight(s)
				if err != nil {
					return fmt.Errorf("invalid sith, weight not a weight")
				}
				clause = append(clause, w)
			}
		}
		thold = append(thold, clause)
	}

	return t.processWeighted(thold)
}

func isSequence(sequence any) bool {
	switch sequence.(type) {
	case []string, []any:
		return true
	default:
		return false
	}
}

func (t *Tholder) processUnweighted(thold big.Int) error {
	tholdInt := int(thold.Int64())

	if tholdInt >= 0 {
		t.thold = tholdInt
		t.weighted = false
		t.size = types.Size(tholdInt)
		t.satisfy = satisfyNumeric
		number, err := NewNumber(&thold, nil)
		if err != nil {
			return err
		}
		t.number = number
		t.bexter = nil
	} else {
		return fmt.Errorf("invalid thold")
	}

	return nil
}

func (t *Tholder) processWeighted(thold any) error {
	if thold == nil {
		thold = []any{}
	}

	tholdList, ok := thold.([]any)
	if !ok {
		return fmt.Errorf("invalid thold, not a list")
	}

	sum, err := sumClause(tholdList, 0)
	if err != nil {
		return err
	}

	if sum.Cmp(big.NewRat(1, 1)) < 0 {
		return fmt.Errorf("invalid thold, does not sum to 1+")
	}

	t.thold = tholdList
	t.weighted = true
	s := 0
	for _, clause := range tholdList {
		clauseList, ok := clause.([]any)
		if !ok {
			return fmt.Errorf("clause is a %T, not a list", clause)
		}

		for _, e := range clauseList {
			switch eType := e.(type) {
			case map[string]any:
				if len(eType) > 1 {
					return fmt.Errorf("invalid clause, map length > 1")
				}

				for _, v := range eType {
					vList, ok := v.([]any)
					if !ok {
						return fmt.Errorf("invalid clause, value is a %T, not a list", v)
					}

					s += len(vList)
				}
			default:
				s += 1
			}
		}
	}
	t.size = types.Size(s)
	t.satisfy = satisfyWeighted

	ta := [][]string{}

	for _, c := range tholdList {
		bc := []string{}

		switch cType := c.(type) {
		case []any:
			v := []string{}
			for _, e := range cType {
				switch eType := e.(type) {
				case *big.Rat:
					v = append(v, extractIntOrRational(eType, "s"))
				case map[string]any:
					if len(eType) > 1 {
						return fmt.Errorf("invalid clause, map length > 1")
					}

					rational, list, err := decomposeMapClause(eType)
					if err != nil {
						return err
					}

					var k string
					_v := []string{}

					k = extractIntOrRational(rational, "s")

					for _, e := range list {
						_v = append(_v, extractIntOrRational(e, "s"))
					}

					kv := k + "k" + strings.Join(_v, "v")
					v = append(v, kv)
				default:
					return fmt.Errorf("invalid clause type %T", eType)
				}

			}

			bc = append(bc, strings.Join(v, "c"))
		default:
			return fmt.Errorf("invalid clause type %T", cType)
		}

		ta = append(ta, bc)
	}

	taPrime := []string{}
	for _, bc := range ta {
		taPrime = append(taPrime, strings.Join(bc, "c"))
	}

	bext := strings.Join(taPrime, "a")

	t.number = nil
	t.bexter, err = NewBexter(&bext)
	if err != nil {
		return err
	}

	return nil
}

func decomposeMapClause(clause map[string]any) (*big.Rat, []*big.Rat, error) {
	if len(clause) > 1 {
		return nil, nil, fmt.Errorf("invalid clause, map length > 1")
	}

	var k string
	var v []*big.Rat
	for _k, _v := range clause {
		k = _k
		vAny, ok := _v.([]any)
		if !ok {
			return nil, nil, fmt.Errorf("invalid clause, map value is a %T, not a list", _v)
		}

		for _, _r := range vAny {
			r, ok := _r.(*big.Rat)
			if !ok {
				return nil, nil, fmt.Errorf("invalid clause, value is a %T, not a rational", _r)
			}

			v = append(v, r)
		}
	}

	r, err := weight(k)
	if err != nil {
		return nil, nil, err
	}

	return r, v, nil
}

func sumClause(clause any, depth int) (*big.Rat, error) {
	if depth > 3 {
		return nil, fmt.Errorf("invalid clause depth")
	}

	sum := big.NewRat(0, 1)

	switch t := clause.(type) {
	case map[string]any:
		rational, list, err := decomposeMapClause(t)
		if err != nil {
			return nil, err
		}
		sum.Add(sum, rational)

		listSum, err := sumClause(list, depth+1)
		if err != nil {
			return nil, err
		}

		if listSum.Cmp(big.NewRat(1, 1)) < 0 {
			return nil, fmt.Errorf("invalid clause - nested clause weight sums must be >= 1")
		}
	case *big.Rat:
		sum.Add(sum, t)
	case []*big.Rat:
		for _, item := range t {
			sum.Add(sum, item)
		}
	case []any:
		for _, item := range t {
			itemSum, err := sumClause(item, depth+1)
			if err != nil {
				return nil, err
			}
			sum.Add(sum, itemSum)
		}
	case string:
		w, err := weight(t)
		if err != nil {
			return nil, err
		}
		sum.Add(sum, w)
	default:
		return nil, fmt.Errorf("invalid clause type %T", t)
	}

	return sum, nil
}

// this is a bit different than KERIpy in that it
func satisfyNumeric(t *Tholder, indices []types.Index) bool {
	tholdInt, ok := t.thold.(int)
	if !ok {
		return false
	}

	unique := []types.Index{}

	for _, index := range indices {
		if !slices.Contains(unique, index) {
			unique = append(unique, index)
		}
	}

	return len(unique) >= tholdInt
}

func satisfyWeighted(t *Tholder, indices []types.Index) bool {
	unique := []types.Index{}

	for _, index := range indices {
		if !slices.Contains(unique, index) {
			unique = append(unique, index)
		}
	}

	slices.Sort(unique)
	sats := make([]bool, t.size)
	for _, index := range unique {
		sats[index] = true
	}

	tholdList, ok := t.thold.([]any)
	if !ok {
		return false
	}

	wio := 0
	for _, clause := range tholdList {
		cw := big.NewRat(0, 1)
		clauseList, ok := clause.([]any)
		if !ok {
			return false
		}

		for _, e := range clauseList {
			switch eType := e.(type) {
			case map[string]any:
				if len(eType) > 1 {
					return false
				}

				var k *big.Rat
				vw := big.NewRat(0, 1)
				for _k, v := range eType {
					valueList := v.([]any)

					for _, val := range valueList {
						value, ok := val.(*big.Rat)
						if !ok {
							return false
						}

						if sats[wio] {
							vw.Add(vw, value)
						}
						wio += 1
					}

					var err error
					k, err = weight(_k)
					if err != nil {
						return false
					}
				}

				if vw.Cmp(big.NewRat(1, 1)) >= 0 {
					cw.Add(cw, k)
				}
			case *big.Rat:
				if sats[wio] {
					cw.Add(cw, eType)
				}
				wio += 1
			default:
				return false
			}
		}

		if cw.Cmp(big.NewRat(1, 1)) < 0 {
			return false
		}
	}

	return true
}

func weight(w string) (*big.Rat, error) {
	parts := strings.Split(w, "/")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid weight")
	}

	if len(parts) == 1 {
		n, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		return big.NewRat(n, 1), nil
	}

	n, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}

	d, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return big.NewRat(n, d), nil
}
