package lengthconv

import "fmt"

type Meters float64
type Foot float64

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string   { return fmt.Sprintf("%gft", f) }

func MToF(m Meters) Foot {
	return Foot(m * 3.28084)
}

func FToM(f Foot) Meters {
	return Meters(f * 0.3048)
}
