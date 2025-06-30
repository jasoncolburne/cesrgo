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

func b64CharToIndex(c byte) (uint8, error) {
	switch c {
	case 'A':
		return 0, nil
	case 'B':
		return 1, nil
	case 'C':
		return 2, nil
	case 'D':
		return 3, nil
	case 'E':
		return 4, nil
	case 'F':
		return 5, nil
	case 'G':
		return 6, nil
	case 'H':
		return 7, nil
	case 'I':
		return 8, nil
	case 'J':
		return 9, nil
	case 'K':
		return 10, nil
	case 'L':
		return 11, nil
	case 'M':
		return 12, nil
	case 'N':
		return 13, nil
	case 'O':
		return 14, nil
	case 'P':
		return 15, nil
	case 'Q':
		return 16, nil
	case 'R':
		return 17, nil
	case 'S':
		return 18, nil
	case 'T':
		return 19, nil
	case 'U':
		return 20, nil
	case 'V':
		return 21, nil
	case 'W':
		return 22, nil
	case 'X':
		return 23, nil
	case 'Y':
		return 24, nil
	case 'Z':
		return 25, nil
	case 'a':
		return 26, nil
	case 'b':
		return 27, nil
	case 'c':
		return 28, nil
	case 'd':
		return 29, nil
	case 'e':
		return 30, nil
	case 'f':
		return 31, nil
	case 'g':
		return 32, nil
	case 'h':
		return 33, nil
	case 'i':
		return 34, nil
	case 'j':
		return 35, nil
	case 'k':
		return 36, nil
	case 'l':
		return 37, nil
	case 'm':
		return 38, nil
	case 'n':
		return 39, nil
	case 'o':
		return 40, nil
	case 'p':
		return 41, nil
	case 'q':
		return 42, nil
	case 'r':
		return 43, nil
	case 's':
		return 44, nil
	case 't':
		return 45, nil
	case 'u':
		return 46, nil
	case 'v':
		return 47, nil
	case 'w':
		return 48, nil
	case 'x':
		return 49, nil
	case 'y':
		return 50, nil
	case 'z':
		return 51, nil
	case '0':
		return 52, nil
	case '1':
		return 53, nil
	case '2':
		return 54, nil
	case '3':
		return 55, nil
	case '4':
		return 56, nil
	case '5':
		return 57, nil
	case '6':
		return 58, nil
	case '7':
		return 59, nil
	case '8':
		return 60, nil
	case '9':
		return 61, nil
	case '-':
		return 62, nil
	case '_':
		return 63, nil
	default:
		return 0, fmt.Errorf("invalid base64 character: %c", c)
	}
}

func b64IndexToChar(i uint8) (byte, error) {
	switch i {
	case 0:
		return 'A', nil
	case 1:
		return 'B', nil
	case 2:
		return 'C', nil
	case 3:
		return 'D', nil
	case 4:
		return 'E', nil
	case 5:
		return 'F', nil
	case 6:
		return 'G', nil
	case 7:
		return 'H', nil
	case 8:
		return 'I', nil
	case 9:
		return 'J', nil
	case 10:
		return 'K', nil
	case 11:
		return 'L', nil
	case 12:
		return 'M', nil
	case 13:
		return 'N', nil
	case 14:
		return 'O', nil
	case 15:
		return 'P', nil
	case 16:
		return 'Q', nil
	case 17:
		return 'R', nil
	case 18:
		return 'S', nil
	case 19:
		return 'T', nil
	case 20:
		return 'U', nil
	case 21:
		return 'V', nil
	case 22:
		return 'W', nil
	case 23:
		return 'X', nil
	case 24:
		return 'Y', nil
	case 25:
		return 'Z', nil
	case 26:
		return 'a', nil
	case 27:
		return 'b', nil
	case 28:
		return 'c', nil
	case 29:
		return 'd', nil
	case 30:
		return 'e', nil
	case 31:
		return 'f', nil
	case 32:
		return 'g', nil
	case 33:
		return 'h', nil
	case 34:
		return 'i', nil
	case 35:
		return 'j', nil
	case 36:
		return 'k', nil
	case 37:
		return 'l', nil
	case 38:
		return 'm', nil
	case 39:
		return 'n', nil
	case 40:
		return 'o', nil
	case 41:
		return 'p', nil
	case 42:
		return 'q', nil
	case 43:
		return 'r', nil
	case 44:
		return 's', nil
	case 45:
		return 't', nil
	case 46:
		return 'u', nil
	case 47:
		return 'v', nil
	case 48:
		return 'w', nil
	case 49:
		return 'x', nil
	case 50:
		return 'y', nil
	case 51:
		return 'z', nil
	case 52:
		return '0', nil
	case 53:
		return '1', nil
	case 54:
		return '2', nil
	case 55:
		return '3', nil
	case 56:
		return '4', nil
	case 57:
		return '5', nil
	case 58:
		return '6', nil
	case 59:
		return '7', nil
	case 60:
		return '8', nil
	case 61:
		return '9', nil
	case 62:
		return '-', nil
	case 63:
		return '_', nil
	default:
		return 0, fmt.Errorf("invalid base64 index: %d", i)
	}
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

	n := ((length + 1) * 3) / 4

	if n > len(b2) {
		return "", fmt.Errorf("not enough bytes")
	}

	if length <= 4 {
		bytes := [4]byte{0, 0, 0, 0}
		copy(bytes[:], b2[:n])

		i := binary.BigEndian.Uint32(bytes[:])
		tbs := 2*(length%4) + (4-n)*8
		return u32ToB64(i>>tbs, length)
	} else if length <= 8 {
		bytes := [8]byte{}
		copy(bytes[:], b2[:n])

		i := binary.BigEndian.Uint64(bytes[:])
		tbs := 2*(length%4) + (8-n)*8
		return u64ToB64(i>>tbs, length)
	} else {
		return "", fmt.Errorf("unexpected length")
	}
}

func codeB64ToB2(code string) ([]byte, error) {
	i, err := b64ToU64(code)
	if err != nil {
		return nil, err
	}

	i <<= 2 * (len(code) % 4)
	n := ((len(code) + 1) * 3) / 4
	return binary.BigEndian.AppendUint64(make([]byte, 0, 8), i)[8-n:], nil
}

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
	var out string = strings.Repeat("A", length)

	var overflow float64 = float64(length) - math.Log2(float64(n))/math.Log2(64)
	for x > 0 {
		if overflow >= 0.0 {
			i, err := b64CharToIndex(byte(x % 64))
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
	var out string = strings.Repeat("A", length)

	var overflow float64 = float64(length) - math.Log2(float64(n))/math.Log2(64)
	for x > 0 {
		if overflow >= 0.0 {
			i, err := b64CharToIndex(byte(x % 64))
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
