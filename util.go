package cesrgo

import (
	"encoding/binary"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/jasoncolburne/cesrgo/types"
)

func validateCode(code types.Code, validCodes []types.Code) bool {
	return slices.Contains(validCodes, code)
}

var b64Runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

//nolint:dupl
func b64CharToIndex(c byte) (uint8, error) {
	index := strings.Index(b64Runes, string(c))
	if index == -1 {
		return 0, fmt.Errorf("invalid base64 character: %c", c)
	}

	if index > 63 {
		return 0, fmt.Errorf("programmer error:invalid base64 character: %c", c)
	}

	return uint8(index), nil
}

//nolint:dupl
func b64IndexToChar(i uint8) (byte, error) {
	if i > 63 {
		return 0, fmt.Errorf("programmer error:invalid base64 index: %d", i)
	}

	return b64Runes[i], nil
}

func nabSextets(binary []byte, count int) ([]byte, error) {
	n := ((count + 1) * 3) / 4

	if n > len(binary) {
		return nil, fmt.Errorf("binary is too small")
	}

	bps := 3 - (len(binary) % 3)
	padded := make([]byte, len(binary)+bps)
	copy(padded, binary)

	out := make([]byte, len(padded)*4/3)
	i := 0
	j := 0
	for {
		n := uint32(padded[i])<<16 | uint32(padded[i+1])<<8 | uint32(padded[i+2])

		out[j] = byte((n & 0xfc0000) >> 18)
		out[j+1] = byte((n & 0x03f000) >> 12)
		out[j+2] = byte((n & 0x000fc0) >> 6)
		out[j+3] = byte(n & 0x00003f)

		j += 4
		i += 3
		if i >= len(padded) {
			break
		}
	}

	return out[:count], nil
}

func codeB2ToB64(b2 []byte, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	n := int(math.Ceil(float64(length) * 3 / 4))

	if n > len(b2) {
		return "", fmt.Errorf("not enough bytes")
	}

	tbs := 2 * (length % 4)
	if length <= 4 {
		bytes := [4]byte{}
		copy(bytes[:], b2[:n])

		i := binary.BigEndian.Uint32(bytes[:])
		return u32ToB64(i>>tbs, length)
	} else if length <= 8 {
		bytes := [8]byte{}
		copy(bytes[:], b2[:n])

		i := binary.BigEndian.Uint64(bytes[:])
		return u64ToB64(i>>tbs, length)
	} else {
		return "", fmt.Errorf("unexpected length")
	}
}

// func codeB64ToB2(code string) ([]byte, error) {
// 	i, err := b64ToU64(code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	i <<= 2 * (len(code) % 4)
// 	n := ((len(code) + 1) * 3) / 4
// 	return binary.BigEndian.AppendUint64(make([]byte, 0, 8), i)[8-n:], nil
// }

func b64ToU32(b64 string) (uint32, error) {
	var out uint32 = 0

	for _, c := range b64 {
		i, err := b64CharToIndex(byte(c))
		if err != nil {
			return 0, err
		}
		out = (out << 6) + uint32(i)
	}

	return out, nil
}

func b64ToU64(b64 string) (uint64, error) {
	var out uint64 = 0

	for _, c := range b64 {
		i, err := b64CharToIndex(byte(c))
		if err != nil {
			return 0, err
		}
		out = (out << 6) + uint64(i)
	}

	return out, nil
}

func u32ToB64(n uint32, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	var x uint32 = n
	out := ""

	var overflow float64 = float64(length) - math.Log2(float64(n))/math.Log2(64)
	for x > 0 {
		if overflow >= 0.0 {
			i, err := b64IndexToChar(byte(x % 64))
			if err != nil {
				return "", err
			}
			out = string(i) + out
		} else {
			overflow += 1
		}
		x /= 64
	}

	for len(out) < length {
		out = "A" + out
	}

	return out, nil
}

func u64ToB64(n uint64, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	var x uint64 = n
	out := ""

	var overflow float64 = float64(length) - math.Log2(float64(n))/math.Log2(64)
	for x > 0 {
		if overflow >= 0.0 {
			c, err := b64IndexToChar(byte(x % 64))
			if err != nil {
				return "", err
			}
			out = string([]byte{c}) + out
		} else {
			overflow += 1
		}
		x /= 64
	}

	for len(out) < length {
		out = "A" + out
	}

	return out, nil
}
