package jsonim

import (
	"encoding/json"

	"github.com/chewxy/math32"
)

// JsonNilFloat32 knows how to store a float to JSON
type JsonNilFloat32 float32

func (jf JsonNilFloat32) MarshalJSON() ([]byte, error) {
	f := float32(jf)
	if math32.IsInf(f, 0) || math32.IsNaN(f) {
		return json.Marshal(nil)
	} else {
		return json.Marshal(f)
	}
}
