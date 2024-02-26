package jsonim

import (
	"fmt"

	"github.com/chewxy/math32"
)

// JsonStrFloat32 supports marshaling of a float (as a string) so we can support NaN and Inf
type JsonStrFloat32 float32

const NegativeInf = "\"-INF\""
const PositiveInf = "\"+INF\""
const NotNumber = "\"NaN\""

//goland:noinspection GoMixedReceiverTypes -- this is the "Go Way" to implement a Marshaller
func (jf JsonStrFloat32) MarshalJSON() ([]byte, error) {
	v := float32(jf)
	if math32.IsInf(v, -1) {
		return []byte(NegativeInf), nil
	} else if math32.IsInf(v, +1) {
		return []byte(PositiveInf), nil
	} else if math32.IsNaN(v) {
		return []byte(NotNumber), nil
	} else {
		return []byte(fmt.Sprintf("\"%v\"", v)), nil //json.Marshal(v)
	}
}

//goland:noinspection GoMixedReceiverTypes -- this is the "Go Way" to implement a Marshaller
func (jf *JsonStrFloat32) UnmarshalJSON(v []byte) error {
	s := string(v)
	s = s[1 : len(s)-1] // remove wrapping quotes
	if s == NegativeInf {
		*jf = JsonStrFloat32(math32.Inf(-1))
	} else if s == PositiveInf {
		*jf = JsonStrFloat32(math32.Inf(1))
	} else if s == NotNumber {
		*jf = JsonStrFloat32(math32.NaN())
	} else {
		var regularFloat float32
		_, err := fmt.Sscanf(s, "%v", &regularFloat)
		if err != nil {
			return err
		}
		*jf = JsonStrFloat32(regularFloat)
	}
	return nil
}
