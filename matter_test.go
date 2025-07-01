package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type TestMatter struct {
	code types.Code
	size types.Size
	raw  types.Raw
	soft *string
}

func (m *TestMatter) SetCode(code types.Code) {
	m.code = code
}

func (m *TestMatter) GetCode() types.Code {
	return m.code
}

func (m *TestMatter) SetRaw(raw types.Raw) {
	m.raw = raw
}

func (m *TestMatter) GetRaw() types.Raw {
	return m.raw
}

func (m *TestMatter) SetSize(size types.Size) {
	m.size = size
}

func (m *TestMatter) GetSize() types.Size {
	return m.size
}

func (m *TestMatter) Hard() string {
	return string(m.code)
}

func (m *TestMatter) SetSoft(soft *string) {
	m.soft = soft
}

func (m *TestMatter) GetSoft() string {
	return *m.soft
}

func (m *TestMatter) Both() (string, error) {
	return "", nil
}

func (m *TestMatter) Qb2() (types.Qb2, error) {
	return types.Qb2{}, nil
}

func (m *TestMatter) Qb64() (types.Qb64, error) {
	return types.Qb64(""), nil
}

func (m *TestMatter) Qb64b() (types.Qb64b, error) {
	return types.Qb64b{}, nil
}

func TestNewMatter(t *testing.T) {
	m := &TestMatter{}

	if err := cesrgo.NewMatter(m); err == nil {
		t.Fatalf("no options should fail")
	}

	if err := cesrgo.NewMatter(
		m,
		options.WithCode(codex.Blake3_256),
		options.WithRaw(types.Raw{}),
		options.WithQb2(types.Qb2{}),
	); err == nil {
		t.Fatalf("code + raw + qb2 should fail")
	}

	if err := cesrgo.NewMatter(
		m,
		options.WithQb2(types.Qb2{19, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}),
	); err == nil {
		t.Fatalf("? should fail")
	}
}
