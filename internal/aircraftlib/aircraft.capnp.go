package aircraftlib

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	math "math"
	strconv "strconv"
	C "zombiezen.com/go/capnproto"
	server "zombiezen.com/go/capnproto/server"
)

type Zdate struct{ C.Struct }

func NewZdate(s *C.Segment) (Zdate, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Zdate{}, err
	}
	return Zdate{st}, nil
}

func NewRootZdate(s *C.Segment) (Zdate, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Zdate{}, err
	}
	return Zdate{st}, nil
}

func ReadRootZdate(msg *C.Message) (Zdate, error) {
	root, err := msg.Root()
	if err != nil {
		return Zdate{}, err
	}
	st := C.ToStruct(root)
	return Zdate{st}, nil
}

func (s Zdate) Year() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s Zdate) SetYear(v int16) {

	s.Struct.SetUint16(0, uint16(v))
}

func (s Zdate) Month() uint8 {
	return s.Struct.Uint8(2)
}

func (s Zdate) SetMonth(v uint8) {

	s.Struct.SetUint8(2, v)
}

func (s Zdate) Day() uint8 {
	return s.Struct.Uint8(3)
}

func (s Zdate) SetDay(v uint8) {

	s.Struct.SetUint8(3, v)
}

// Zdate_List is a list of Zdate.
type Zdate_List struct{ C.List }

// NewZdate creates a new list of Zdate.
func NewZdate_List(s *C.Segment, sz int32) (Zdate_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Zdate_List{}, err
	}
	return Zdate_List{l}, nil
}

func (s Zdate_List) At(i int) Zdate           { return Zdate{s.List.Struct(i)} }
func (s Zdate_List) Set(i int, v Zdate) error { return s.List.SetStruct(i, v.Struct) }

// Zdate_Promise is a wrapper for a Zdate promised by a client call.
type Zdate_Promise struct{ *C.Pipeline }

func (p Zdate_Promise) Struct() (Zdate, error) {
	s, err := p.Pipeline.Struct()
	return Zdate{s}, err
}

type Zdata struct{ C.Struct }

func NewZdata(s *C.Segment) (Zdata, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Zdata{}, err
	}
	return Zdata{st}, nil
}

func NewRootZdata(s *C.Segment) (Zdata, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Zdata{}, err
	}
	return Zdata{st}, nil
}

func ReadRootZdata(msg *C.Message) (Zdata, error) {
	root, err := msg.Root()
	if err != nil {
		return Zdata{}, err
	}
	st := C.ToStruct(root)
	return Zdata{st}, nil
}

func (s Zdata) Data() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(C.ToData(p)), nil

}

func (s Zdata) SetData(v []byte) error {

	d, err := C.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

// Zdata_List is a list of Zdata.
type Zdata_List struct{ C.List }

// NewZdata creates a new list of Zdata.
func NewZdata_List(s *C.Segment, sz int32) (Zdata_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Zdata_List{}, err
	}
	return Zdata_List{l}, nil
}

func (s Zdata_List) At(i int) Zdata           { return Zdata{s.List.Struct(i)} }
func (s Zdata_List) Set(i int, v Zdata) error { return s.List.SetStruct(i, v.Struct) }

// Zdata_Promise is a wrapper for a Zdata promised by a client call.
type Zdata_Promise struct{ *C.Pipeline }

func (p Zdata_Promise) Struct() (Zdata, error) {
	s, err := p.Pipeline.Struct()
	return Zdata{s}, err
}

type Airport uint16

// Values of Airport.
const (
	Airport_none Airport = 0
	Airport_jfk  Airport = 1
	Airport_lax  Airport = 2
	Airport_sfo  Airport = 3
	Airport_luv  Airport = 4
	Airport_dfw  Airport = 5
	Airport_test Airport = 6
)

// String returns the enum's constant name.
func (c Airport) String() string {
	switch c {
	case Airport_none:
		return "none"
	case Airport_jfk:
		return "jfk"
	case Airport_lax:
		return "lax"
	case Airport_sfo:
		return "sfo"
	case Airport_luv:
		return "luv"
	case Airport_dfw:
		return "dfw"
	case Airport_test:
		return "test"

	default:
		return ""
	}
}

// AirportFromString returns the enum value with a name,
// or the zero value if there's no such value.
func AirportFromString(c string) Airport {
	switch c {
	case "none":
		return Airport_none
	case "jfk":
		return Airport_jfk
	case "lax":
		return Airport_lax
	case "sfo":
		return Airport_sfo
	case "luv":
		return Airport_luv
	case "dfw":
		return Airport_dfw
	case "test":
		return Airport_test

	default:
		return 0
	}
}

type Airport_List struct{ C.List }

func NewAirport_List(s *C.Segment, sz int32) (Airport_List, error) {
	l, err := C.NewUInt16List(s, sz)
	if err != nil {
		return Airport_List{}, err
	}
	return Airport_List{l.List}, nil
}

func (l Airport_List) At(i int) Airport {
	ul := C.UInt16List{List: l.List}
	return Airport(ul.At(i))
}

func (l Airport_List) Set(i int, v Airport) {
	ul := C.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type PlaneBase struct{ C.Struct }

func NewPlaneBase(s *C.Segment) (PlaneBase, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 32, PointerCount: 2})
	if err != nil {
		return PlaneBase{}, err
	}
	return PlaneBase{st}, nil
}

func NewRootPlaneBase(s *C.Segment) (PlaneBase, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 32, PointerCount: 2})
	if err != nil {
		return PlaneBase{}, err
	}
	return PlaneBase{st}, nil
}

func ReadRootPlaneBase(msg *C.Message) (PlaneBase, error) {
	root, err := msg.Root()
	if err != nil {
		return PlaneBase{}, err
	}
	st := C.ToStruct(root)
	return PlaneBase{st}, nil
}

func (s PlaneBase) Name() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s PlaneBase) SetName(v string) error {

	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s PlaneBase) Homes() (Airport_List, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Airport_List{}, err
	}

	l := C.ToList(p)

	return Airport_List{List: l}, nil
}

func (s PlaneBase) SetHomes(v Airport_List) error {

	return s.Struct.SetPointer(1, v.List)
}

func (s PlaneBase) Rating() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s PlaneBase) SetRating(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s PlaneBase) CanFly() bool {
	return s.Struct.Bit(64)
}

func (s PlaneBase) SetCanFly(v bool) {

	s.Struct.SetBit(64, v)
}

func (s PlaneBase) Capacity() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s PlaneBase) SetCapacity(v int64) {

	s.Struct.SetUint64(16, uint64(v))
}

func (s PlaneBase) MaxSpeed() float64 {
	return math.Float64frombits(s.Struct.Uint64(24))
}

func (s PlaneBase) SetMaxSpeed(v float64) {

	s.Struct.SetUint64(24, math.Float64bits(v))
}

// PlaneBase_List is a list of PlaneBase.
type PlaneBase_List struct{ C.List }

// NewPlaneBase creates a new list of PlaneBase.
func NewPlaneBase_List(s *C.Segment, sz int32) (PlaneBase_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 32, PointerCount: 2}, sz)
	if err != nil {
		return PlaneBase_List{}, err
	}
	return PlaneBase_List{l}, nil
}

func (s PlaneBase_List) At(i int) PlaneBase           { return PlaneBase{s.List.Struct(i)} }
func (s PlaneBase_List) Set(i int, v PlaneBase) error { return s.List.SetStruct(i, v.Struct) }

// PlaneBase_Promise is a wrapper for a PlaneBase promised by a client call.
type PlaneBase_Promise struct{ *C.Pipeline }

func (p PlaneBase_Promise) Struct() (PlaneBase, error) {
	s, err := p.Pipeline.Struct()
	return PlaneBase{s}, err
}

type B737 struct{ C.Struct }

func NewB737(s *C.Segment) (B737, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return B737{}, err
	}
	return B737{st}, nil
}

func NewRootB737(s *C.Segment) (B737, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return B737{}, err
	}
	return B737{st}, nil
}

func ReadRootB737(msg *C.Message) (B737, error) {
	root, err := msg.Root()
	if err != nil {
		return B737{}, err
	}
	st := C.ToStruct(root)
	return B737{st}, nil
}

func (s B737) Base() (PlaneBase, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return PlaneBase{}, err
	}

	ss := C.ToStruct(p)

	return PlaneBase{Struct: ss}, nil
}

func (s B737) SetBase(v PlaneBase) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s B737) NewBase() (PlaneBase, error) {

	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// B737_List is a list of B737.
type B737_List struct{ C.List }

// NewB737 creates a new list of B737.
func NewB737_List(s *C.Segment, sz int32) (B737_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return B737_List{}, err
	}
	return B737_List{l}, nil
}

func (s B737_List) At(i int) B737           { return B737{s.List.Struct(i)} }
func (s B737_List) Set(i int, v B737) error { return s.List.SetStruct(i, v.Struct) }

// B737_Promise is a wrapper for a B737 promised by a client call.
type B737_Promise struct{ *C.Pipeline }

func (p B737_Promise) Struct() (B737, error) {
	s, err := p.Pipeline.Struct()
	return B737{s}, err
}

func (p B737_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type A320 struct{ C.Struct }

func NewA320(s *C.Segment) (A320, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return A320{}, err
	}
	return A320{st}, nil
}

func NewRootA320(s *C.Segment) (A320, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return A320{}, err
	}
	return A320{st}, nil
}

func ReadRootA320(msg *C.Message) (A320, error) {
	root, err := msg.Root()
	if err != nil {
		return A320{}, err
	}
	st := C.ToStruct(root)
	return A320{st}, nil
}

func (s A320) Base() (PlaneBase, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return PlaneBase{}, err
	}

	ss := C.ToStruct(p)

	return PlaneBase{Struct: ss}, nil
}

func (s A320) SetBase(v PlaneBase) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s A320) NewBase() (PlaneBase, error) {

	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// A320_List is a list of A320.
type A320_List struct{ C.List }

// NewA320 creates a new list of A320.
func NewA320_List(s *C.Segment, sz int32) (A320_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return A320_List{}, err
	}
	return A320_List{l}, nil
}

func (s A320_List) At(i int) A320           { return A320{s.List.Struct(i)} }
func (s A320_List) Set(i int, v A320) error { return s.List.SetStruct(i, v.Struct) }

// A320_Promise is a wrapper for a A320 promised by a client call.
type A320_Promise struct{ *C.Pipeline }

func (p A320_Promise) Struct() (A320, error) {
	s, err := p.Pipeline.Struct()
	return A320{s}, err
}

func (p A320_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type F16 struct{ C.Struct }

func NewF16(s *C.Segment) (F16, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return F16{}, err
	}
	return F16{st}, nil
}

func NewRootF16(s *C.Segment) (F16, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return F16{}, err
	}
	return F16{st}, nil
}

func ReadRootF16(msg *C.Message) (F16, error) {
	root, err := msg.Root()
	if err != nil {
		return F16{}, err
	}
	st := C.ToStruct(root)
	return F16{st}, nil
}

func (s F16) Base() (PlaneBase, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return PlaneBase{}, err
	}

	ss := C.ToStruct(p)

	return PlaneBase{Struct: ss}, nil
}

func (s F16) SetBase(v PlaneBase) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s F16) NewBase() (PlaneBase, error) {

	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// F16_List is a list of F16.
type F16_List struct{ C.List }

// NewF16 creates a new list of F16.
func NewF16_List(s *C.Segment, sz int32) (F16_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return F16_List{}, err
	}
	return F16_List{l}, nil
}

func (s F16_List) At(i int) F16           { return F16{s.List.Struct(i)} }
func (s F16_List) Set(i int, v F16) error { return s.List.SetStruct(i, v.Struct) }

// F16_Promise is a wrapper for a F16 promised by a client call.
type F16_Promise struct{ *C.Pipeline }

func (p F16_Promise) Struct() (F16, error) {
	s, err := p.Pipeline.Struct()
	return F16{s}, err
}

func (p F16_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Regression struct{ C.Struct }

func NewRegression(s *C.Segment) (Regression, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 24, PointerCount: 3})
	if err != nil {
		return Regression{}, err
	}
	return Regression{st}, nil
}

func NewRootRegression(s *C.Segment) (Regression, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 24, PointerCount: 3})
	if err != nil {
		return Regression{}, err
	}
	return Regression{st}, nil
}

func ReadRootRegression(msg *C.Message) (Regression, error) {
	root, err := msg.Root()
	if err != nil {
		return Regression{}, err
	}
	st := C.ToStruct(root)
	return Regression{st}, nil
}

func (s Regression) Base() (PlaneBase, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return PlaneBase{}, err
	}

	ss := C.ToStruct(p)

	return PlaneBase{Struct: ss}, nil
}

func (s Regression) SetBase(v PlaneBase) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s Regression) NewBase() (PlaneBase, error) {

	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Regression) B0() float64 {
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s Regression) SetB0(v float64) {

	s.Struct.SetUint64(0, math.Float64bits(v))
}

func (s Regression) Beta() (C.Float64List, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return C.Float64List{}, err
	}

	l := C.ToList(p)

	return C.Float64List{List: l}, nil
}

func (s Regression) SetBeta(v C.Float64List) error {

	return s.Struct.SetPointer(1, v.List)
}

func (s Regression) Planes() (Aircraft_List, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return Aircraft_List{}, err
	}

	l := C.ToList(p)

	return Aircraft_List{List: l}, nil
}

func (s Regression) SetPlanes(v Aircraft_List) error {

	return s.Struct.SetPointer(2, v.List)
}

func (s Regression) Ymu() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Regression) SetYmu(v float64) {

	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Regression) Ysd() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s Regression) SetYsd(v float64) {

	s.Struct.SetUint64(16, math.Float64bits(v))
}

// Regression_List is a list of Regression.
type Regression_List struct{ C.List }

// NewRegression creates a new list of Regression.
func NewRegression_List(s *C.Segment, sz int32) (Regression_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	if err != nil {
		return Regression_List{}, err
	}
	return Regression_List{l}, nil
}

func (s Regression_List) At(i int) Regression           { return Regression{s.List.Struct(i)} }
func (s Regression_List) Set(i int, v Regression) error { return s.List.SetStruct(i, v.Struct) }

// Regression_Promise is a wrapper for a Regression promised by a client call.
type Regression_Promise struct{ *C.Pipeline }

func (p Regression_Promise) Struct() (Regression, error) {
	s, err := p.Pipeline.Struct()
	return Regression{s}, err
}

func (p Regression_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Aircraft struct{ C.Struct }
type Aircraft_Which uint16

const (
	Aircraft_Which_void Aircraft_Which = 0
	Aircraft_Which_b737 Aircraft_Which = 1
	Aircraft_Which_a320 Aircraft_Which = 2
	Aircraft_Which_f16  Aircraft_Which = 3
)

func (w Aircraft_Which) String() string {
	const s = "voidb737a320f16"
	switch w {
	case Aircraft_Which_void:
		return s[0:4]
	case Aircraft_Which_b737:
		return s[4:8]
	case Aircraft_Which_a320:
		return s[8:12]
	case Aircraft_Which_f16:
		return s[12:15]

	}
	return "Aircraft_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewAircraft(s *C.Segment) (Aircraft, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Aircraft{}, err
	}
	return Aircraft{st}, nil
}

func NewRootAircraft(s *C.Segment) (Aircraft, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Aircraft{}, err
	}
	return Aircraft{st}, nil
}

func ReadRootAircraft(msg *C.Message) (Aircraft, error) {
	root, err := msg.Root()
	if err != nil {
		return Aircraft{}, err
	}
	st := C.ToStruct(root)
	return Aircraft{st}, nil
}

func (s Aircraft) Which() Aircraft_Which {
	return Aircraft_Which(s.Struct.Uint16(0))
}

func (s Aircraft) SetVoid() {
	s.Struct.SetUint16(0, 0)
}

func (s Aircraft) B737() (B737, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return B737{}, err
	}

	ss := C.ToStruct(p)

	return B737{Struct: ss}, nil
}

func (s Aircraft) SetB737(v B737) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewB737 sets the b737 field to a newly
// allocated B737 struct, preferring placement in s's segment.
func (s Aircraft) NewB737() (B737, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewB737(s.Struct.Segment())
	if err != nil {
		return B737{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Aircraft) A320() (A320, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return A320{}, err
	}

	ss := C.ToStruct(p)

	return A320{Struct: ss}, nil
}

func (s Aircraft) SetA320(v A320) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewA320 sets the a320 field to a newly
// allocated A320 struct, preferring placement in s's segment.
func (s Aircraft) NewA320() (A320, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewA320(s.Struct.Segment())
	if err != nil {
		return A320{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Aircraft) F16() (F16, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return F16{}, err
	}

	ss := C.ToStruct(p)

	return F16{Struct: ss}, nil
}

func (s Aircraft) SetF16(v F16) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewF16 sets the f16 field to a newly
// allocated F16 struct, preferring placement in s's segment.
func (s Aircraft) NewF16() (F16, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewF16(s.Struct.Segment())
	if err != nil {
		return F16{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// Aircraft_List is a list of Aircraft.
type Aircraft_List struct{ C.List }

// NewAircraft creates a new list of Aircraft.
func NewAircraft_List(s *C.Segment, sz int32) (Aircraft_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Aircraft_List{}, err
	}
	return Aircraft_List{l}, nil
}

func (s Aircraft_List) At(i int) Aircraft           { return Aircraft{s.List.Struct(i)} }
func (s Aircraft_List) Set(i int, v Aircraft) error { return s.List.SetStruct(i, v.Struct) }

// Aircraft_Promise is a wrapper for a Aircraft promised by a client call.
type Aircraft_Promise struct{ *C.Pipeline }

func (p Aircraft_Promise) Struct() (Aircraft, error) {
	s, err := p.Pipeline.Struct()
	return Aircraft{s}, err
}

func (p Aircraft_Promise) B737() B737_Promise {
	return B737_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Aircraft_Promise) A320() A320_Promise {
	return A320_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Aircraft_Promise) F16() F16_Promise {
	return F16_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Z struct{ C.Struct }
type Z_Which uint16

const (
	Z_Which_void        Z_Which = 0
	Z_Which_zz          Z_Which = 1
	Z_Which_f64         Z_Which = 2
	Z_Which_f32         Z_Which = 3
	Z_Which_i64         Z_Which = 4
	Z_Which_i32         Z_Which = 5
	Z_Which_i16         Z_Which = 6
	Z_Which_i8          Z_Which = 7
	Z_Which_u64         Z_Which = 8
	Z_Which_u32         Z_Which = 9
	Z_Which_u16         Z_Which = 10
	Z_Which_u8          Z_Which = 11
	Z_Which_bool        Z_Which = 12
	Z_Which_text        Z_Which = 13
	Z_Which_blob        Z_Which = 14
	Z_Which_f64vec      Z_Which = 15
	Z_Which_f32vec      Z_Which = 16
	Z_Which_i64vec      Z_Which = 17
	Z_Which_i32vec      Z_Which = 18
	Z_Which_i16vec      Z_Which = 19
	Z_Which_i8vec       Z_Which = 20
	Z_Which_u64vec      Z_Which = 21
	Z_Which_u32vec      Z_Which = 22
	Z_Which_u16vec      Z_Which = 23
	Z_Which_u8vec       Z_Which = 24
	Z_Which_zvec        Z_Which = 25
	Z_Which_zvecvec     Z_Which = 26
	Z_Which_zdate       Z_Which = 27
	Z_Which_zdata       Z_Which = 28
	Z_Which_aircraftvec Z_Which = 29
	Z_Which_aircraft    Z_Which = 30
	Z_Which_regression  Z_Which = 31
	Z_Which_planebase   Z_Which = 32
	Z_Which_airport     Z_Which = 33
	Z_Which_b737        Z_Which = 34
	Z_Which_a320        Z_Which = 35
	Z_Which_f16         Z_Which = 36
	Z_Which_zdatevec    Z_Which = 37
	Z_Which_zdatavec    Z_Which = 38
	Z_Which_boolvec     Z_Which = 39
)

func (w Z_Which) String() string {
	const s = "voidzzf64f32i64i32i16i8u64u32u16u8booltextblobf64vecf32veci64veci32veci16veci8vecu64vecu32vecu16vecu8veczveczvecveczdatezdataaircraftvecaircraftregressionplanebaseairportb737a320f16zdateveczdatavecboolvec"
	switch w {
	case Z_Which_void:
		return s[0:4]
	case Z_Which_zz:
		return s[4:6]
	case Z_Which_f64:
		return s[6:9]
	case Z_Which_f32:
		return s[9:12]
	case Z_Which_i64:
		return s[12:15]
	case Z_Which_i32:
		return s[15:18]
	case Z_Which_i16:
		return s[18:21]
	case Z_Which_i8:
		return s[21:23]
	case Z_Which_u64:
		return s[23:26]
	case Z_Which_u32:
		return s[26:29]
	case Z_Which_u16:
		return s[29:32]
	case Z_Which_u8:
		return s[32:34]
	case Z_Which_bool:
		return s[34:38]
	case Z_Which_text:
		return s[38:42]
	case Z_Which_blob:
		return s[42:46]
	case Z_Which_f64vec:
		return s[46:52]
	case Z_Which_f32vec:
		return s[52:58]
	case Z_Which_i64vec:
		return s[58:64]
	case Z_Which_i32vec:
		return s[64:70]
	case Z_Which_i16vec:
		return s[70:76]
	case Z_Which_i8vec:
		return s[76:81]
	case Z_Which_u64vec:
		return s[81:87]
	case Z_Which_u32vec:
		return s[87:93]
	case Z_Which_u16vec:
		return s[93:99]
	case Z_Which_u8vec:
		return s[99:104]
	case Z_Which_zvec:
		return s[104:108]
	case Z_Which_zvecvec:
		return s[108:115]
	case Z_Which_zdate:
		return s[115:120]
	case Z_Which_zdata:
		return s[120:125]
	case Z_Which_aircraftvec:
		return s[125:136]
	case Z_Which_aircraft:
		return s[136:144]
	case Z_Which_regression:
		return s[144:154]
	case Z_Which_planebase:
		return s[154:163]
	case Z_Which_airport:
		return s[163:170]
	case Z_Which_b737:
		return s[170:174]
	case Z_Which_a320:
		return s[174:178]
	case Z_Which_f16:
		return s[178:181]
	case Z_Which_zdatevec:
		return s[181:189]
	case Z_Which_zdatavec:
		return s[189:197]
	case Z_Which_boolvec:
		return s[197:204]

	}
	return "Z_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewZ(s *C.Segment) (Z, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Z{}, err
	}
	return Z{st}, nil
}

func NewRootZ(s *C.Segment) (Z, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Z{}, err
	}
	return Z{st}, nil
}

func ReadRootZ(msg *C.Message) (Z, error) {
	root, err := msg.Root()
	if err != nil {
		return Z{}, err
	}
	st := C.ToStruct(root)
	return Z{st}, nil
}

func (s Z) Which() Z_Which {
	return Z_Which(s.Struct.Uint16(0))
}

func (s Z) SetVoid() {
	s.Struct.SetUint16(0, 0)
}

func (s Z) Zz() (Z, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Z{}, err
	}

	ss := C.ToStruct(p)

	return Z{Struct: ss}, nil
}

func (s Z) SetZz(v Z) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewZz sets the zz field to a newly
// allocated Z struct, preferring placement in s's segment.
func (s Z) NewZz() (Z, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewZ(s.Struct.Segment())
	if err != nil {
		return Z{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) F64() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Z) SetF64(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Z) F32() float32 {
	return math.Float32frombits(s.Struct.Uint32(8))
}

func (s Z) SetF32(v float32) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint32(8, math.Float32bits(v))
}

func (s Z) I64() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Z) SetI64(v int64) {
	s.Struct.SetUint16(0, 4)
	s.Struct.SetUint64(8, uint64(v))
}

func (s Z) I32() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s Z) SetI32(v int32) {
	s.Struct.SetUint16(0, 5)
	s.Struct.SetUint32(8, uint32(v))
}

func (s Z) I16() int16 {
	return int16(s.Struct.Uint16(8))
}

func (s Z) SetI16(v int16) {
	s.Struct.SetUint16(0, 6)
	s.Struct.SetUint16(8, uint16(v))
}

func (s Z) I8() int8 {
	return int8(s.Struct.Uint8(8))
}

func (s Z) SetI8(v int8) {
	s.Struct.SetUint16(0, 7)
	s.Struct.SetUint8(8, uint8(v))
}

func (s Z) U64() uint64 {
	return s.Struct.Uint64(8)
}

func (s Z) SetU64(v uint64) {
	s.Struct.SetUint16(0, 8)
	s.Struct.SetUint64(8, v)
}

func (s Z) U32() uint32 {
	return s.Struct.Uint32(8)
}

func (s Z) SetU32(v uint32) {
	s.Struct.SetUint16(0, 9)
	s.Struct.SetUint32(8, v)
}

func (s Z) U16() uint16 {
	return s.Struct.Uint16(8)
}

func (s Z) SetU16(v uint16) {
	s.Struct.SetUint16(0, 10)
	s.Struct.SetUint16(8, v)
}

func (s Z) U8() uint8 {
	return s.Struct.Uint8(8)
}

func (s Z) SetU8(v uint8) {
	s.Struct.SetUint16(0, 11)
	s.Struct.SetUint8(8, v)
}

func (s Z) Bool() bool {
	return s.Struct.Bit(64)
}

func (s Z) SetBool(v bool) {
	s.Struct.SetUint16(0, 12)
	s.Struct.SetBit(64, v)
}

func (s Z) Text() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s Z) SetText(v string) error {
	s.Struct.SetUint16(0, 13)
	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Z) Blob() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(C.ToData(p)), nil

}

func (s Z) SetBlob(v []byte) error {
	s.Struct.SetUint16(0, 14)
	d, err := C.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s Z) F64vec() (C.Float64List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.Float64List{}, err
	}

	l := C.ToList(p)

	return C.Float64List{List: l}, nil
}

func (s Z) SetF64vec(v C.Float64List) error {
	s.Struct.SetUint16(0, 15)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) F32vec() (C.Float32List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.Float32List{}, err
	}

	l := C.ToList(p)

	return C.Float32List{List: l}, nil
}

func (s Z) SetF32vec(v C.Float32List) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) I64vec() (C.Int64List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.Int64List{}, err
	}

	l := C.ToList(p)

	return C.Int64List{List: l}, nil
}

func (s Z) SetI64vec(v C.Int64List) error {
	s.Struct.SetUint16(0, 17)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) I32vec() (C.Int32List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.Int32List{}, err
	}

	l := C.ToList(p)

	return C.Int32List{List: l}, nil
}

func (s Z) SetI32vec(v C.Int32List) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) I16vec() (C.Int16List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.Int16List{}, err
	}

	l := C.ToList(p)

	return C.Int16List{List: l}, nil
}

func (s Z) SetI16vec(v C.Int16List) error {
	s.Struct.SetUint16(0, 19)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) I8vec() (C.Int8List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.Int8List{}, err
	}

	l := C.ToList(p)

	return C.Int8List{List: l}, nil
}

func (s Z) SetI8vec(v C.Int8List) error {
	s.Struct.SetUint16(0, 20)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) U64vec() (C.UInt64List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.UInt64List{}, err
	}

	l := C.ToList(p)

	return C.UInt64List{List: l}, nil
}

func (s Z) SetU64vec(v C.UInt64List) error {
	s.Struct.SetUint16(0, 21)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) U32vec() (C.UInt32List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.UInt32List{}, err
	}

	l := C.ToList(p)

	return C.UInt32List{List: l}, nil
}

func (s Z) SetU32vec(v C.UInt32List) error {
	s.Struct.SetUint16(0, 22)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) U16vec() (C.UInt16List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.UInt16List{}, err
	}

	l := C.ToList(p)

	return C.UInt16List{List: l}, nil
}

func (s Z) SetU16vec(v C.UInt16List) error {
	s.Struct.SetUint16(0, 23)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) U8vec() (C.UInt8List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.UInt8List{}, err
	}

	l := C.ToList(p)

	return C.UInt8List{List: l}, nil
}

func (s Z) SetU8vec(v C.UInt8List) error {
	s.Struct.SetUint16(0, 24)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) Zvec() (Z_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Z_List{}, err
	}

	l := C.ToList(p)

	return Z_List{List: l}, nil
}

func (s Z) SetZvec(v Z_List) error {
	s.Struct.SetUint16(0, 25)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) Zvecvec() (C.PointerList, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.PointerList{}, err
	}

	l := C.ToList(p)

	return C.PointerList{List: l}, nil
}

func (s Z) SetZvecvec(v C.PointerList) error {
	s.Struct.SetUint16(0, 26)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) Zdate() (Zdate, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Zdate{}, err
	}

	ss := C.ToStruct(p)

	return Zdate{Struct: ss}, nil
}

func (s Z) SetZdate(v Zdate) error {
	s.Struct.SetUint16(0, 27)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewZdate sets the zdate field to a newly
// allocated Zdate struct, preferring placement in s's segment.
func (s Z) NewZdate() (Zdate, error) {
	s.Struct.SetUint16(0, 27)
	ss, err := NewZdate(s.Struct.Segment())
	if err != nil {
		return Zdate{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) Zdata() (Zdata, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Zdata{}, err
	}

	ss := C.ToStruct(p)

	return Zdata{Struct: ss}, nil
}

func (s Z) SetZdata(v Zdata) error {
	s.Struct.SetUint16(0, 28)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewZdata sets the zdata field to a newly
// allocated Zdata struct, preferring placement in s's segment.
func (s Z) NewZdata() (Zdata, error) {
	s.Struct.SetUint16(0, 28)
	ss, err := NewZdata(s.Struct.Segment())
	if err != nil {
		return Zdata{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) Aircraftvec() (Aircraft_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Aircraft_List{}, err
	}

	l := C.ToList(p)

	return Aircraft_List{List: l}, nil
}

func (s Z) SetAircraftvec(v Aircraft_List) error {
	s.Struct.SetUint16(0, 29)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) Aircraft() (Aircraft, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Aircraft{}, err
	}

	ss := C.ToStruct(p)

	return Aircraft{Struct: ss}, nil
}

func (s Z) SetAircraft(v Aircraft) error {
	s.Struct.SetUint16(0, 30)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewAircraft sets the aircraft field to a newly
// allocated Aircraft struct, preferring placement in s's segment.
func (s Z) NewAircraft() (Aircraft, error) {
	s.Struct.SetUint16(0, 30)
	ss, err := NewAircraft(s.Struct.Segment())
	if err != nil {
		return Aircraft{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) Regression() (Regression, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Regression{}, err
	}

	ss := C.ToStruct(p)

	return Regression{Struct: ss}, nil
}

func (s Z) SetRegression(v Regression) error {
	s.Struct.SetUint16(0, 31)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewRegression sets the regression field to a newly
// allocated Regression struct, preferring placement in s's segment.
func (s Z) NewRegression() (Regression, error) {
	s.Struct.SetUint16(0, 31)
	ss, err := NewRegression(s.Struct.Segment())
	if err != nil {
		return Regression{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) Planebase() (PlaneBase, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return PlaneBase{}, err
	}

	ss := C.ToStruct(p)

	return PlaneBase{Struct: ss}, nil
}

func (s Z) SetPlanebase(v PlaneBase) error {
	s.Struct.SetUint16(0, 32)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewPlanebase sets the planebase field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s Z) NewPlanebase() (PlaneBase, error) {
	s.Struct.SetUint16(0, 32)
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) Airport() Airport {
	return Airport(s.Struct.Uint16(8))
}

func (s Z) SetAirport(v Airport) {
	s.Struct.SetUint16(0, 33)
	s.Struct.SetUint16(8, uint16(v))
}

func (s Z) B737() (B737, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return B737{}, err
	}

	ss := C.ToStruct(p)

	return B737{Struct: ss}, nil
}

func (s Z) SetB737(v B737) error {
	s.Struct.SetUint16(0, 34)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewB737 sets the b737 field to a newly
// allocated B737 struct, preferring placement in s's segment.
func (s Z) NewB737() (B737, error) {
	s.Struct.SetUint16(0, 34)
	ss, err := NewB737(s.Struct.Segment())
	if err != nil {
		return B737{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) A320() (A320, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return A320{}, err
	}

	ss := C.ToStruct(p)

	return A320{Struct: ss}, nil
}

func (s Z) SetA320(v A320) error {
	s.Struct.SetUint16(0, 35)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewA320 sets the a320 field to a newly
// allocated A320 struct, preferring placement in s's segment.
func (s Z) NewA320() (A320, error) {
	s.Struct.SetUint16(0, 35)
	ss, err := NewA320(s.Struct.Segment())
	if err != nil {
		return A320{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) F16() (F16, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return F16{}, err
	}

	ss := C.ToStruct(p)

	return F16{Struct: ss}, nil
}

func (s Z) SetF16(v F16) error {
	s.Struct.SetUint16(0, 36)
	return s.Struct.SetPointer(0, v.Struct)
}

// NewF16 sets the f16 field to a newly
// allocated F16 struct, preferring placement in s's segment.
func (s Z) NewF16() (F16, error) {
	s.Struct.SetUint16(0, 36)
	ss, err := NewF16(s.Struct.Segment())
	if err != nil {
		return F16{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Z) Zdatevec() (Zdate_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Zdate_List{}, err
	}

	l := C.ToList(p)

	return Zdate_List{List: l}, nil
}

func (s Z) SetZdatevec(v Zdate_List) error {
	s.Struct.SetUint16(0, 37)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) Zdatavec() (Zdata_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Zdata_List{}, err
	}

	l := C.ToList(p)

	return Zdata_List{List: l}, nil
}

func (s Z) SetZdatavec(v Zdata_List) error {
	s.Struct.SetUint16(0, 38)
	return s.Struct.SetPointer(0, v.List)
}

func (s Z) Boolvec() (C.BitList, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.BitList{}, err
	}

	l := C.ToList(p)

	return C.BitList{List: l}, nil
}

func (s Z) SetBoolvec(v C.BitList) error {
	s.Struct.SetUint16(0, 39)
	return s.Struct.SetPointer(0, v.List)
}

// Z_List is a list of Z.
type Z_List struct{ C.List }

// NewZ creates a new list of Z.
func NewZ_List(s *C.Segment, sz int32) (Z_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	if err != nil {
		return Z_List{}, err
	}
	return Z_List{l}, nil
}

func (s Z_List) At(i int) Z           { return Z{s.List.Struct(i)} }
func (s Z_List) Set(i int, v Z) error { return s.List.SetStruct(i, v.Struct) }

// Z_Promise is a wrapper for a Z promised by a client call.
type Z_Promise struct{ *C.Pipeline }

func (p Z_Promise) Struct() (Z, error) {
	s, err := p.Pipeline.Struct()
	return Z{s}, err
}

func (p Z_Promise) Zz() Z_Promise {
	return Z_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Zdate() Zdate_Promise {
	return Zdate_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Zdata() Zdata_Promise {
	return Zdata_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Aircraft() Aircraft_Promise {
	return Aircraft_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Regression() Regression_Promise {
	return Regression_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Planebase() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) B737() B737_Promise {
	return B737_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) A320() A320_Promise {
	return A320_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) F16() F16_Promise {
	return F16_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Counter struct{ C.Struct }

func NewCounter(s *C.Segment) (Counter, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Counter{}, err
	}
	return Counter{st}, nil
}

func NewRootCounter(s *C.Segment) (Counter, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Counter{}, err
	}
	return Counter{st}, nil
}

func ReadRootCounter(msg *C.Message) (Counter, error) {
	root, err := msg.Root()
	if err != nil {
		return Counter{}, err
	}
	st := C.ToStruct(root)
	return Counter{st}, nil
}

func (s Counter) Size() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Counter) SetSize(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s Counter) Words() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s Counter) SetWords(v string) error {

	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Counter) Wordlist() (C.TextList, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return C.TextList{}, err
	}

	l := C.ToList(p)

	return C.TextList{List: l}, nil
}

func (s Counter) SetWordlist(v C.TextList) error {

	return s.Struct.SetPointer(1, v.List)
}

// Counter_List is a list of Counter.
type Counter_List struct{ C.List }

// NewCounter creates a new list of Counter.
func NewCounter_List(s *C.Segment, sz int32) (Counter_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Counter_List{}, err
	}
	return Counter_List{l}, nil
}

func (s Counter_List) At(i int) Counter           { return Counter{s.List.Struct(i)} }
func (s Counter_List) Set(i int, v Counter) error { return s.List.SetStruct(i, v.Struct) }

// Counter_Promise is a wrapper for a Counter promised by a client call.
type Counter_Promise struct{ *C.Pipeline }

func (p Counter_Promise) Struct() (Counter, error) {
	s, err := p.Pipeline.Struct()
	return Counter{s}, err
}

type Bag struct{ C.Struct }

func NewBag(s *C.Segment) (Bag, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Bag{}, err
	}
	return Bag{st}, nil
}

func NewRootBag(s *C.Segment) (Bag, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Bag{}, err
	}
	return Bag{st}, nil
}

func ReadRootBag(msg *C.Message) (Bag, error) {
	root, err := msg.Root()
	if err != nil {
		return Bag{}, err
	}
	st := C.ToStruct(root)
	return Bag{st}, nil
}

func (s Bag) Counter() (Counter, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Counter{}, err
	}

	ss := C.ToStruct(p)

	return Counter{Struct: ss}, nil
}

func (s Bag) SetCounter(v Counter) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewCounter sets the counter field to a newly
// allocated Counter struct, preferring placement in s's segment.
func (s Bag) NewCounter() (Counter, error) {

	ss, err := NewCounter(s.Struct.Segment())
	if err != nil {
		return Counter{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// Bag_List is a list of Bag.
type Bag_List struct{ C.List }

// NewBag creates a new list of Bag.
func NewBag_List(s *C.Segment, sz int32) (Bag_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Bag_List{}, err
	}
	return Bag_List{l}, nil
}

func (s Bag_List) At(i int) Bag           { return Bag{s.List.Struct(i)} }
func (s Bag_List) Set(i int, v Bag) error { return s.List.SetStruct(i, v.Struct) }

// Bag_Promise is a wrapper for a Bag promised by a client call.
type Bag_Promise struct{ *C.Pipeline }

func (p Bag_Promise) Struct() (Bag, error) {
	s, err := p.Pipeline.Struct()
	return Bag{s}, err
}

func (p Bag_Promise) Counter() Counter_Promise {
	return Counter_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Zserver struct{ C.Struct }

func NewZserver(s *C.Segment) (Zserver, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Zserver{}, err
	}
	return Zserver{st}, nil
}

func NewRootZserver(s *C.Segment) (Zserver, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Zserver{}, err
	}
	return Zserver{st}, nil
}

func ReadRootZserver(msg *C.Message) (Zserver, error) {
	root, err := msg.Root()
	if err != nil {
		return Zserver{}, err
	}
	st := C.ToStruct(root)
	return Zserver{st}, nil
}

func (s Zserver) Waitingjobs() (Zjob_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Zjob_List{}, err
	}

	l := C.ToList(p)

	return Zjob_List{List: l}, nil
}

func (s Zserver) SetWaitingjobs(v Zjob_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// Zserver_List is a list of Zserver.
type Zserver_List struct{ C.List }

// NewZserver creates a new list of Zserver.
func NewZserver_List(s *C.Segment, sz int32) (Zserver_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Zserver_List{}, err
	}
	return Zserver_List{l}, nil
}

func (s Zserver_List) At(i int) Zserver           { return Zserver{s.List.Struct(i)} }
func (s Zserver_List) Set(i int, v Zserver) error { return s.List.SetStruct(i, v.Struct) }

// Zserver_Promise is a wrapper for a Zserver promised by a client call.
type Zserver_Promise struct{ *C.Pipeline }

func (p Zserver_Promise) Struct() (Zserver, error) {
	s, err := p.Pipeline.Struct()
	return Zserver{s}, err
}

type Zjob struct{ C.Struct }

func NewZjob(s *C.Segment) (Zjob, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Zjob{}, err
	}
	return Zjob{st}, nil
}

func NewRootZjob(s *C.Segment) (Zjob, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Zjob{}, err
	}
	return Zjob{st}, nil
}

func ReadRootZjob(msg *C.Message) (Zjob, error) {
	root, err := msg.Root()
	if err != nil {
		return Zjob{}, err
	}
	st := C.ToStruct(root)
	return Zjob{st}, nil
}

func (s Zjob) Cmd() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s Zjob) SetCmd(v string) error {

	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Zjob) Args() (C.TextList, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return C.TextList{}, err
	}

	l := C.ToList(p)

	return C.TextList{List: l}, nil
}

func (s Zjob) SetArgs(v C.TextList) error {

	return s.Struct.SetPointer(1, v.List)
}

// Zjob_List is a list of Zjob.
type Zjob_List struct{ C.List }

// NewZjob creates a new list of Zjob.
func NewZjob_List(s *C.Segment, sz int32) (Zjob_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Zjob_List{}, err
	}
	return Zjob_List{l}, nil
}

func (s Zjob_List) At(i int) Zjob           { return Zjob{s.List.Struct(i)} }
func (s Zjob_List) Set(i int, v Zjob) error { return s.List.SetStruct(i, v.Struct) }

// Zjob_Promise is a wrapper for a Zjob promised by a client call.
type Zjob_Promise struct{ *C.Pipeline }

func (p Zjob_Promise) Struct() (Zjob, error) {
	s, err := p.Pipeline.Struct()
	return Zjob{s}, err
}

type VerEmpty struct{ C.Struct }

func NewVerEmpty(s *C.Segment) (VerEmpty, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return VerEmpty{}, err
	}
	return VerEmpty{st}, nil
}

func NewRootVerEmpty(s *C.Segment) (VerEmpty, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return VerEmpty{}, err
	}
	return VerEmpty{st}, nil
}

func ReadRootVerEmpty(msg *C.Message) (VerEmpty, error) {
	root, err := msg.Root()
	if err != nil {
		return VerEmpty{}, err
	}
	st := C.ToStruct(root)
	return VerEmpty{st}, nil
}

// VerEmpty_List is a list of VerEmpty.
type VerEmpty_List struct{ C.List }

// NewVerEmpty creates a new list of VerEmpty.
func NewVerEmpty_List(s *C.Segment, sz int32) (VerEmpty_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return VerEmpty_List{}, err
	}
	return VerEmpty_List{l}, nil
}

func (s VerEmpty_List) At(i int) VerEmpty           { return VerEmpty{s.List.Struct(i)} }
func (s VerEmpty_List) Set(i int, v VerEmpty) error { return s.List.SetStruct(i, v.Struct) }

// VerEmpty_Promise is a wrapper for a VerEmpty promised by a client call.
type VerEmpty_Promise struct{ *C.Pipeline }

func (p VerEmpty_Promise) Struct() (VerEmpty, error) {
	s, err := p.Pipeline.Struct()
	return VerEmpty{s}, err
}

type VerOneData struct{ C.Struct }

func NewVerOneData(s *C.Segment) (VerOneData, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return VerOneData{}, err
	}
	return VerOneData{st}, nil
}

func NewRootVerOneData(s *C.Segment) (VerOneData, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return VerOneData{}, err
	}
	return VerOneData{st}, nil
}

func ReadRootVerOneData(msg *C.Message) (VerOneData, error) {
	root, err := msg.Root()
	if err != nil {
		return VerOneData{}, err
	}
	st := C.ToStruct(root)
	return VerOneData{st}, nil
}

func (s VerOneData) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerOneData) SetVal(v int16) {

	s.Struct.SetUint16(0, uint16(v))
}

// VerOneData_List is a list of VerOneData.
type VerOneData_List struct{ C.List }

// NewVerOneData creates a new list of VerOneData.
func NewVerOneData_List(s *C.Segment, sz int32) (VerOneData_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return VerOneData_List{}, err
	}
	return VerOneData_List{l}, nil
}

func (s VerOneData_List) At(i int) VerOneData           { return VerOneData{s.List.Struct(i)} }
func (s VerOneData_List) Set(i int, v VerOneData) error { return s.List.SetStruct(i, v.Struct) }

// VerOneData_Promise is a wrapper for a VerOneData promised by a client call.
type VerOneData_Promise struct{ *C.Pipeline }

func (p VerOneData_Promise) Struct() (VerOneData, error) {
	s, err := p.Pipeline.Struct()
	return VerOneData{s}, err
}

type VerTwoData struct{ C.Struct }

func NewVerTwoData(s *C.Segment) (VerTwoData, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return VerTwoData{}, err
	}
	return VerTwoData{st}, nil
}

func NewRootVerTwoData(s *C.Segment) (VerTwoData, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return VerTwoData{}, err
	}
	return VerTwoData{st}, nil
}

func ReadRootVerTwoData(msg *C.Message) (VerTwoData, error) {
	root, err := msg.Root()
	if err != nil {
		return VerTwoData{}, err
	}
	st := C.ToStruct(root)
	return VerTwoData{st}, nil
}

func (s VerTwoData) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoData) SetVal(v int16) {

	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoData) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoData) SetDuo(v int64) {

	s.Struct.SetUint64(8, uint64(v))
}

// VerTwoData_List is a list of VerTwoData.
type VerTwoData_List struct{ C.List }

// NewVerTwoData creates a new list of VerTwoData.
func NewVerTwoData_List(s *C.Segment, sz int32) (VerTwoData_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	if err != nil {
		return VerTwoData_List{}, err
	}
	return VerTwoData_List{l}, nil
}

func (s VerTwoData_List) At(i int) VerTwoData           { return VerTwoData{s.List.Struct(i)} }
func (s VerTwoData_List) Set(i int, v VerTwoData) error { return s.List.SetStruct(i, v.Struct) }

// VerTwoData_Promise is a wrapper for a VerTwoData promised by a client call.
type VerTwoData_Promise struct{ *C.Pipeline }

func (p VerTwoData_Promise) Struct() (VerTwoData, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoData{s}, err
}

type VerOnePtr struct{ C.Struct }

func NewVerOnePtr(s *C.Segment) (VerOnePtr, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return VerOnePtr{}, err
	}
	return VerOnePtr{st}, nil
}

func NewRootVerOnePtr(s *C.Segment) (VerOnePtr, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return VerOnePtr{}, err
	}
	return VerOnePtr{st}, nil
}

func ReadRootVerOnePtr(msg *C.Message) (VerOnePtr, error) {
	root, err := msg.Root()
	if err != nil {
		return VerOnePtr{}, err
	}
	st := C.ToStruct(root)
	return VerOnePtr{st}, nil
}

func (s VerOnePtr) Ptr() (VerOneData, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerOneData{}, err
	}

	ss := C.ToStruct(p)

	return VerOneData{Struct: ss}, nil
}

func (s VerOnePtr) SetPtr(v VerOneData) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewPtr sets the ptr field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerOnePtr) NewPtr() (VerOneData, error) {

	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// VerOnePtr_List is a list of VerOnePtr.
type VerOnePtr_List struct{ C.List }

// NewVerOnePtr creates a new list of VerOnePtr.
func NewVerOnePtr_List(s *C.Segment, sz int32) (VerOnePtr_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return VerOnePtr_List{}, err
	}
	return VerOnePtr_List{l}, nil
}

func (s VerOnePtr_List) At(i int) VerOnePtr           { return VerOnePtr{s.List.Struct(i)} }
func (s VerOnePtr_List) Set(i int, v VerOnePtr) error { return s.List.SetStruct(i, v.Struct) }

// VerOnePtr_Promise is a wrapper for a VerOnePtr promised by a client call.
type VerOnePtr_Promise struct{ *C.Pipeline }

func (p VerOnePtr_Promise) Struct() (VerOnePtr, error) {
	s, err := p.Pipeline.Struct()
	return VerOnePtr{s}, err
}

func (p VerOnePtr_Promise) Ptr() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type VerTwoPtr struct{ C.Struct }

func NewVerTwoPtr(s *C.Segment) (VerTwoPtr, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return VerTwoPtr{}, err
	}
	return VerTwoPtr{st}, nil
}

func NewRootVerTwoPtr(s *C.Segment) (VerTwoPtr, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return VerTwoPtr{}, err
	}
	return VerTwoPtr{st}, nil
}

func ReadRootVerTwoPtr(msg *C.Message) (VerTwoPtr, error) {
	root, err := msg.Root()
	if err != nil {
		return VerTwoPtr{}, err
	}
	st := C.ToStruct(root)
	return VerTwoPtr{st}, nil
}

func (s VerTwoPtr) Ptr1() (VerOneData, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerOneData{}, err
	}

	ss := C.ToStruct(p)

	return VerOneData{Struct: ss}, nil
}

func (s VerTwoPtr) SetPtr1(v VerOneData) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoPtr) NewPtr1() (VerOneData, error) {

	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s VerTwoPtr) Ptr2() (VerOneData, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return VerOneData{}, err
	}

	ss := C.ToStruct(p)

	return VerOneData{Struct: ss}, nil
}

func (s VerTwoPtr) SetPtr2(v VerOneData) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoPtr) NewPtr2() (VerOneData, error) {

	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

// VerTwoPtr_List is a list of VerTwoPtr.
type VerTwoPtr_List struct{ C.List }

// NewVerTwoPtr creates a new list of VerTwoPtr.
func NewVerTwoPtr_List(s *C.Segment, sz int32) (VerTwoPtr_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return VerTwoPtr_List{}, err
	}
	return VerTwoPtr_List{l}, nil
}

func (s VerTwoPtr_List) At(i int) VerTwoPtr           { return VerTwoPtr{s.List.Struct(i)} }
func (s VerTwoPtr_List) Set(i int, v VerTwoPtr) error { return s.List.SetStruct(i, v.Struct) }

// VerTwoPtr_Promise is a wrapper for a VerTwoPtr promised by a client call.
type VerTwoPtr_Promise struct{ *C.Pipeline }

func (p VerTwoPtr_Promise) Struct() (VerTwoPtr, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoPtr{s}, err
}

func (p VerTwoPtr_Promise) Ptr1() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoPtr_Promise) Ptr2() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type VerTwoDataTwoPtr struct{ C.Struct }

func NewVerTwoDataTwoPtr(s *C.Segment) (VerTwoDataTwoPtr, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 16, PointerCount: 2})
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	return VerTwoDataTwoPtr{st}, nil
}

func NewRootVerTwoDataTwoPtr(s *C.Segment) (VerTwoDataTwoPtr, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 16, PointerCount: 2})
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	return VerTwoDataTwoPtr{st}, nil
}

func ReadRootVerTwoDataTwoPtr(msg *C.Message) (VerTwoDataTwoPtr, error) {
	root, err := msg.Root()
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	st := C.ToStruct(root)
	return VerTwoDataTwoPtr{st}, nil
}

func (s VerTwoDataTwoPtr) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoDataTwoPtr) SetVal(v int16) {

	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoDataTwoPtr) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoDataTwoPtr) SetDuo(v int64) {

	s.Struct.SetUint64(8, uint64(v))
}

func (s VerTwoDataTwoPtr) Ptr1() (VerOneData, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerOneData{}, err
	}

	ss := C.ToStruct(p)

	return VerOneData{Struct: ss}, nil
}

func (s VerTwoDataTwoPtr) SetPtr1(v VerOneData) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoDataTwoPtr) NewPtr1() (VerOneData, error) {

	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s VerTwoDataTwoPtr) Ptr2() (VerOneData, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return VerOneData{}, err
	}

	ss := C.ToStruct(p)

	return VerOneData{Struct: ss}, nil
}

func (s VerTwoDataTwoPtr) SetPtr2(v VerOneData) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoDataTwoPtr) NewPtr2() (VerOneData, error) {

	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

// VerTwoDataTwoPtr_List is a list of VerTwoDataTwoPtr.
type VerTwoDataTwoPtr_List struct{ C.List }

// NewVerTwoDataTwoPtr creates a new list of VerTwoDataTwoPtr.
func NewVerTwoDataTwoPtr_List(s *C.Segment, sz int32) (VerTwoDataTwoPtr_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	if err != nil {
		return VerTwoDataTwoPtr_List{}, err
	}
	return VerTwoDataTwoPtr_List{l}, nil
}

func (s VerTwoDataTwoPtr_List) At(i int) VerTwoDataTwoPtr { return VerTwoDataTwoPtr{s.List.Struct(i)} }
func (s VerTwoDataTwoPtr_List) Set(i int, v VerTwoDataTwoPtr) error {
	return s.List.SetStruct(i, v.Struct)
}

// VerTwoDataTwoPtr_Promise is a wrapper for a VerTwoDataTwoPtr promised by a client call.
type VerTwoDataTwoPtr_Promise struct{ *C.Pipeline }

func (p VerTwoDataTwoPtr_Promise) Struct() (VerTwoDataTwoPtr, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoDataTwoPtr{s}, err
}

func (p VerTwoDataTwoPtr_Promise) Ptr1() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoDataTwoPtr_Promise) Ptr2() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HoldsVerEmptyList struct{ C.Struct }

func NewHoldsVerEmptyList(s *C.Segment) (HoldsVerEmptyList, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerEmptyList{}, err
	}
	return HoldsVerEmptyList{st}, nil
}

func NewRootHoldsVerEmptyList(s *C.Segment) (HoldsVerEmptyList, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerEmptyList{}, err
	}
	return HoldsVerEmptyList{st}, nil
}

func ReadRootHoldsVerEmptyList(msg *C.Message) (HoldsVerEmptyList, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerEmptyList{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerEmptyList{st}, nil
}

func (s HoldsVerEmptyList) Mylist() (VerEmpty_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerEmpty_List{}, err
	}

	l := C.ToList(p)

	return VerEmpty_List{List: l}, nil
}

func (s HoldsVerEmptyList) SetMylist(v VerEmpty_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerEmptyList_List is a list of HoldsVerEmptyList.
type HoldsVerEmptyList_List struct{ C.List }

// NewHoldsVerEmptyList creates a new list of HoldsVerEmptyList.
func NewHoldsVerEmptyList_List(s *C.Segment, sz int32) (HoldsVerEmptyList_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerEmptyList_List{}, err
	}
	return HoldsVerEmptyList_List{l}, nil
}

func (s HoldsVerEmptyList_List) At(i int) HoldsVerEmptyList {
	return HoldsVerEmptyList{s.List.Struct(i)}
}
func (s HoldsVerEmptyList_List) Set(i int, v HoldsVerEmptyList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerEmptyList_Promise is a wrapper for a HoldsVerEmptyList promised by a client call.
type HoldsVerEmptyList_Promise struct{ *C.Pipeline }

func (p HoldsVerEmptyList_Promise) Struct() (HoldsVerEmptyList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerEmptyList{s}, err
}

type HoldsVerOneDataList struct{ C.Struct }

func NewHoldsVerOneDataList(s *C.Segment) (HoldsVerOneDataList, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerOneDataList{}, err
	}
	return HoldsVerOneDataList{st}, nil
}

func NewRootHoldsVerOneDataList(s *C.Segment) (HoldsVerOneDataList, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerOneDataList{}, err
	}
	return HoldsVerOneDataList{st}, nil
}

func ReadRootHoldsVerOneDataList(msg *C.Message) (HoldsVerOneDataList, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerOneDataList{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerOneDataList{st}, nil
}

func (s HoldsVerOneDataList) Mylist() (VerOneData_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerOneData_List{}, err
	}

	l := C.ToList(p)

	return VerOneData_List{List: l}, nil
}

func (s HoldsVerOneDataList) SetMylist(v VerOneData_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerOneDataList_List is a list of HoldsVerOneDataList.
type HoldsVerOneDataList_List struct{ C.List }

// NewHoldsVerOneDataList creates a new list of HoldsVerOneDataList.
func NewHoldsVerOneDataList_List(s *C.Segment, sz int32) (HoldsVerOneDataList_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerOneDataList_List{}, err
	}
	return HoldsVerOneDataList_List{l}, nil
}

func (s HoldsVerOneDataList_List) At(i int) HoldsVerOneDataList {
	return HoldsVerOneDataList{s.List.Struct(i)}
}
func (s HoldsVerOneDataList_List) Set(i int, v HoldsVerOneDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerOneDataList_Promise is a wrapper for a HoldsVerOneDataList promised by a client call.
type HoldsVerOneDataList_Promise struct{ *C.Pipeline }

func (p HoldsVerOneDataList_Promise) Struct() (HoldsVerOneDataList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerOneDataList{s}, err
}

type HoldsVerTwoDataList struct{ C.Struct }

func NewHoldsVerTwoDataList(s *C.Segment) (HoldsVerTwoDataList, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoDataList{}, err
	}
	return HoldsVerTwoDataList{st}, nil
}

func NewRootHoldsVerTwoDataList(s *C.Segment) (HoldsVerTwoDataList, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoDataList{}, err
	}
	return HoldsVerTwoDataList{st}, nil
}

func ReadRootHoldsVerTwoDataList(msg *C.Message) (HoldsVerTwoDataList, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerTwoDataList{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerTwoDataList{st}, nil
}

func (s HoldsVerTwoDataList) Mylist() (VerTwoData_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoData_List{}, err
	}

	l := C.ToList(p)

	return VerTwoData_List{List: l}, nil
}

func (s HoldsVerTwoDataList) SetMylist(v VerTwoData_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerTwoDataList_List is a list of HoldsVerTwoDataList.
type HoldsVerTwoDataList_List struct{ C.List }

// NewHoldsVerTwoDataList creates a new list of HoldsVerTwoDataList.
func NewHoldsVerTwoDataList_List(s *C.Segment, sz int32) (HoldsVerTwoDataList_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerTwoDataList_List{}, err
	}
	return HoldsVerTwoDataList_List{l}, nil
}

func (s HoldsVerTwoDataList_List) At(i int) HoldsVerTwoDataList {
	return HoldsVerTwoDataList{s.List.Struct(i)}
}
func (s HoldsVerTwoDataList_List) Set(i int, v HoldsVerTwoDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoDataList_Promise is a wrapper for a HoldsVerTwoDataList promised by a client call.
type HoldsVerTwoDataList_Promise struct{ *C.Pipeline }

func (p HoldsVerTwoDataList_Promise) Struct() (HoldsVerTwoDataList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoDataList{s}, err
}

type HoldsVerOnePtrList struct{ C.Struct }

func NewHoldsVerOnePtrList(s *C.Segment) (HoldsVerOnePtrList, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerOnePtrList{}, err
	}
	return HoldsVerOnePtrList{st}, nil
}

func NewRootHoldsVerOnePtrList(s *C.Segment) (HoldsVerOnePtrList, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerOnePtrList{}, err
	}
	return HoldsVerOnePtrList{st}, nil
}

func ReadRootHoldsVerOnePtrList(msg *C.Message) (HoldsVerOnePtrList, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerOnePtrList{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerOnePtrList{st}, nil
}

func (s HoldsVerOnePtrList) Mylist() (VerOnePtr_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerOnePtr_List{}, err
	}

	l := C.ToList(p)

	return VerOnePtr_List{List: l}, nil
}

func (s HoldsVerOnePtrList) SetMylist(v VerOnePtr_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerOnePtrList_List is a list of HoldsVerOnePtrList.
type HoldsVerOnePtrList_List struct{ C.List }

// NewHoldsVerOnePtrList creates a new list of HoldsVerOnePtrList.
func NewHoldsVerOnePtrList_List(s *C.Segment, sz int32) (HoldsVerOnePtrList_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerOnePtrList_List{}, err
	}
	return HoldsVerOnePtrList_List{l}, nil
}

func (s HoldsVerOnePtrList_List) At(i int) HoldsVerOnePtrList {
	return HoldsVerOnePtrList{s.List.Struct(i)}
}
func (s HoldsVerOnePtrList_List) Set(i int, v HoldsVerOnePtrList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerOnePtrList_Promise is a wrapper for a HoldsVerOnePtrList promised by a client call.
type HoldsVerOnePtrList_Promise struct{ *C.Pipeline }

func (p HoldsVerOnePtrList_Promise) Struct() (HoldsVerOnePtrList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerOnePtrList{s}, err
}

type HoldsVerTwoPtrList struct{ C.Struct }

func NewHoldsVerTwoPtrList(s *C.Segment) (HoldsVerTwoPtrList, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoPtrList{}, err
	}
	return HoldsVerTwoPtrList{st}, nil
}

func NewRootHoldsVerTwoPtrList(s *C.Segment) (HoldsVerTwoPtrList, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoPtrList{}, err
	}
	return HoldsVerTwoPtrList{st}, nil
}

func ReadRootHoldsVerTwoPtrList(msg *C.Message) (HoldsVerTwoPtrList, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerTwoPtrList{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerTwoPtrList{st}, nil
}

func (s HoldsVerTwoPtrList) Mylist() (VerTwoPtr_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoPtr_List{}, err
	}

	l := C.ToList(p)

	return VerTwoPtr_List{List: l}, nil
}

func (s HoldsVerTwoPtrList) SetMylist(v VerTwoPtr_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerTwoPtrList_List is a list of HoldsVerTwoPtrList.
type HoldsVerTwoPtrList_List struct{ C.List }

// NewHoldsVerTwoPtrList creates a new list of HoldsVerTwoPtrList.
func NewHoldsVerTwoPtrList_List(s *C.Segment, sz int32) (HoldsVerTwoPtrList_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerTwoPtrList_List{}, err
	}
	return HoldsVerTwoPtrList_List{l}, nil
}

func (s HoldsVerTwoPtrList_List) At(i int) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList{s.List.Struct(i)}
}
func (s HoldsVerTwoPtrList_List) Set(i int, v HoldsVerTwoPtrList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoPtrList_Promise is a wrapper for a HoldsVerTwoPtrList promised by a client call.
type HoldsVerTwoPtrList_Promise struct{ *C.Pipeline }

func (p HoldsVerTwoPtrList_Promise) Struct() (HoldsVerTwoPtrList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoPtrList{s}, err
}

type HoldsVerTwoTwoList struct{ C.Struct }

func NewHoldsVerTwoTwoList(s *C.Segment) (HoldsVerTwoTwoList, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoTwoList{}, err
	}
	return HoldsVerTwoTwoList{st}, nil
}

func NewRootHoldsVerTwoTwoList(s *C.Segment) (HoldsVerTwoTwoList, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoTwoList{}, err
	}
	return HoldsVerTwoTwoList{st}, nil
}

func ReadRootHoldsVerTwoTwoList(msg *C.Message) (HoldsVerTwoTwoList, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerTwoTwoList{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerTwoTwoList{st}, nil
}

func (s HoldsVerTwoTwoList) Mylist() (VerTwoDataTwoPtr_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoDataTwoPtr_List{}, err
	}

	l := C.ToList(p)

	return VerTwoDataTwoPtr_List{List: l}, nil
}

func (s HoldsVerTwoTwoList) SetMylist(v VerTwoDataTwoPtr_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerTwoTwoList_List is a list of HoldsVerTwoTwoList.
type HoldsVerTwoTwoList_List struct{ C.List }

// NewHoldsVerTwoTwoList creates a new list of HoldsVerTwoTwoList.
func NewHoldsVerTwoTwoList_List(s *C.Segment, sz int32) (HoldsVerTwoTwoList_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerTwoTwoList_List{}, err
	}
	return HoldsVerTwoTwoList_List{l}, nil
}

func (s HoldsVerTwoTwoList_List) At(i int) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList{s.List.Struct(i)}
}
func (s HoldsVerTwoTwoList_List) Set(i int, v HoldsVerTwoTwoList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoTwoList_Promise is a wrapper for a HoldsVerTwoTwoList promised by a client call.
type HoldsVerTwoTwoList_Promise struct{ *C.Pipeline }

func (p HoldsVerTwoTwoList_Promise) Struct() (HoldsVerTwoTwoList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoTwoList{s}, err
}

type HoldsVerTwoTwoPlus struct{ C.Struct }

func NewHoldsVerTwoTwoPlus(s *C.Segment) (HoldsVerTwoTwoPlus, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoTwoPlus{}, err
	}
	return HoldsVerTwoTwoPlus{st}, nil
}

func NewRootHoldsVerTwoTwoPlus(s *C.Segment) (HoldsVerTwoTwoPlus, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HoldsVerTwoTwoPlus{}, err
	}
	return HoldsVerTwoTwoPlus{st}, nil
}

func ReadRootHoldsVerTwoTwoPlus(msg *C.Message) (HoldsVerTwoTwoPlus, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsVerTwoTwoPlus{}, err
	}
	st := C.ToStruct(root)
	return HoldsVerTwoTwoPlus{st}, nil
}

func (s HoldsVerTwoTwoPlus) Mylist() (VerTwoTwoPlus_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoTwoPlus_List{}, err
	}

	l := C.ToList(p)

	return VerTwoTwoPlus_List{List: l}, nil
}

func (s HoldsVerTwoTwoPlus) SetMylist(v VerTwoTwoPlus_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HoldsVerTwoTwoPlus_List is a list of HoldsVerTwoTwoPlus.
type HoldsVerTwoTwoPlus_List struct{ C.List }

// NewHoldsVerTwoTwoPlus creates a new list of HoldsVerTwoTwoPlus.
func NewHoldsVerTwoTwoPlus_List(s *C.Segment, sz int32) (HoldsVerTwoTwoPlus_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HoldsVerTwoTwoPlus_List{}, err
	}
	return HoldsVerTwoTwoPlus_List{l}, nil
}

func (s HoldsVerTwoTwoPlus_List) At(i int) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus{s.List.Struct(i)}
}
func (s HoldsVerTwoTwoPlus_List) Set(i int, v HoldsVerTwoTwoPlus) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoTwoPlus_Promise is a wrapper for a HoldsVerTwoTwoPlus promised by a client call.
type HoldsVerTwoTwoPlus_Promise struct{ *C.Pipeline }

func (p HoldsVerTwoTwoPlus_Promise) Struct() (HoldsVerTwoTwoPlus, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoTwoPlus{s}, err
}

type VerTwoTwoPlus struct{ C.Struct }

func NewVerTwoTwoPlus(s *C.Segment) (VerTwoTwoPlus, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 24, PointerCount: 3})
	if err != nil {
		return VerTwoTwoPlus{}, err
	}
	return VerTwoTwoPlus{st}, nil
}

func NewRootVerTwoTwoPlus(s *C.Segment) (VerTwoTwoPlus, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 24, PointerCount: 3})
	if err != nil {
		return VerTwoTwoPlus{}, err
	}
	return VerTwoTwoPlus{st}, nil
}

func ReadRootVerTwoTwoPlus(msg *C.Message) (VerTwoTwoPlus, error) {
	root, err := msg.Root()
	if err != nil {
		return VerTwoTwoPlus{}, err
	}
	st := C.ToStruct(root)
	return VerTwoTwoPlus{st}, nil
}

func (s VerTwoTwoPlus) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoTwoPlus) SetVal(v int16) {

	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoTwoPlus) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoTwoPlus) SetDuo(v int64) {

	s.Struct.SetUint64(8, uint64(v))
}

func (s VerTwoTwoPlus) Ptr1() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}

	ss := C.ToStruct(p)

	return VerTwoDataTwoPtr{Struct: ss}, nil
}

func (s VerTwoTwoPlus) SetPtr1(v VerTwoDataTwoPtr) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewPtr1() (VerTwoDataTwoPtr, error) {

	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s VerTwoTwoPlus) Ptr2() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}

	ss := C.ToStruct(p)

	return VerTwoDataTwoPtr{Struct: ss}, nil
}

func (s VerTwoTwoPlus) SetPtr2(v VerTwoDataTwoPtr) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewPtr2() (VerTwoDataTwoPtr, error) {

	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s VerTwoTwoPlus) Tre() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s VerTwoTwoPlus) SetTre(v int64) {

	s.Struct.SetUint64(16, uint64(v))
}

func (s VerTwoTwoPlus) Lst3() (C.Int64List, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return C.Int64List{}, err
	}

	l := C.ToList(p)

	return C.Int64List{List: l}, nil
}

func (s VerTwoTwoPlus) SetLst3(v C.Int64List) error {

	return s.Struct.SetPointer(2, v.List)
}

// VerTwoTwoPlus_List is a list of VerTwoTwoPlus.
type VerTwoTwoPlus_List struct{ C.List }

// NewVerTwoTwoPlus creates a new list of VerTwoTwoPlus.
func NewVerTwoTwoPlus_List(s *C.Segment, sz int32) (VerTwoTwoPlus_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	if err != nil {
		return VerTwoTwoPlus_List{}, err
	}
	return VerTwoTwoPlus_List{l}, nil
}

func (s VerTwoTwoPlus_List) At(i int) VerTwoTwoPlus           { return VerTwoTwoPlus{s.List.Struct(i)} }
func (s VerTwoTwoPlus_List) Set(i int, v VerTwoTwoPlus) error { return s.List.SetStruct(i, v.Struct) }

// VerTwoTwoPlus_Promise is a wrapper for a VerTwoTwoPlus promised by a client call.
type VerTwoTwoPlus_Promise struct{ *C.Pipeline }

func (p VerTwoTwoPlus_Promise) Struct() (VerTwoTwoPlus, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoTwoPlus{s}, err
}

func (p VerTwoTwoPlus_Promise) Ptr1() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoTwoPlus_Promise) Ptr2() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HoldsText struct{ C.Struct }

func NewHoldsText(s *C.Segment) (HoldsText, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HoldsText{}, err
	}
	return HoldsText{st}, nil
}

func NewRootHoldsText(s *C.Segment) (HoldsText, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HoldsText{}, err
	}
	return HoldsText{st}, nil
}

func ReadRootHoldsText(msg *C.Message) (HoldsText, error) {
	root, err := msg.Root()
	if err != nil {
		return HoldsText{}, err
	}
	st := C.ToStruct(root)
	return HoldsText{st}, nil
}

func (s HoldsText) Txt() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s HoldsText) SetTxt(v string) error {

	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s HoldsText) Lst() (C.TextList, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return C.TextList{}, err
	}

	l := C.ToList(p)

	return C.TextList{List: l}, nil
}

func (s HoldsText) SetLst(v C.TextList) error {

	return s.Struct.SetPointer(1, v.List)
}

func (s HoldsText) Lstlst() (C.PointerList, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return C.PointerList{}, err
	}

	l := C.ToList(p)

	return C.PointerList{List: l}, nil
}

func (s HoldsText) SetLstlst(v C.PointerList) error {

	return s.Struct.SetPointer(2, v.List)
}

// HoldsText_List is a list of HoldsText.
type HoldsText_List struct{ C.List }

// NewHoldsText creates a new list of HoldsText.
func NewHoldsText_List(s *C.Segment, sz int32) (HoldsText_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return HoldsText_List{}, err
	}
	return HoldsText_List{l}, nil
}

func (s HoldsText_List) At(i int) HoldsText           { return HoldsText{s.List.Struct(i)} }
func (s HoldsText_List) Set(i int, v HoldsText) error { return s.List.SetStruct(i, v.Struct) }

// HoldsText_Promise is a wrapper for a HoldsText promised by a client call.
type HoldsText_Promise struct{ *C.Pipeline }

func (p HoldsText_Promise) Struct() (HoldsText, error) {
	s, err := p.Pipeline.Struct()
	return HoldsText{s}, err
}

type WrapEmpty struct{ C.Struct }

func NewWrapEmpty(s *C.Segment) (WrapEmpty, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WrapEmpty{}, err
	}
	return WrapEmpty{st}, nil
}

func NewRootWrapEmpty(s *C.Segment) (WrapEmpty, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WrapEmpty{}, err
	}
	return WrapEmpty{st}, nil
}

func ReadRootWrapEmpty(msg *C.Message) (WrapEmpty, error) {
	root, err := msg.Root()
	if err != nil {
		return WrapEmpty{}, err
	}
	st := C.ToStruct(root)
	return WrapEmpty{st}, nil
}

func (s WrapEmpty) MightNotBeReallyEmpty() (VerEmpty, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerEmpty{}, err
	}

	ss := C.ToStruct(p)

	return VerEmpty{Struct: ss}, nil
}

func (s WrapEmpty) SetMightNotBeReallyEmpty(v VerEmpty) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerEmpty struct, preferring placement in s's segment.
func (s WrapEmpty) NewMightNotBeReallyEmpty() (VerEmpty, error) {

	ss, err := NewVerEmpty(s.Struct.Segment())
	if err != nil {
		return VerEmpty{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// WrapEmpty_List is a list of WrapEmpty.
type WrapEmpty_List struct{ C.List }

// NewWrapEmpty creates a new list of WrapEmpty.
func NewWrapEmpty_List(s *C.Segment, sz int32) (WrapEmpty_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WrapEmpty_List{}, err
	}
	return WrapEmpty_List{l}, nil
}

func (s WrapEmpty_List) At(i int) WrapEmpty           { return WrapEmpty{s.List.Struct(i)} }
func (s WrapEmpty_List) Set(i int, v WrapEmpty) error { return s.List.SetStruct(i, v.Struct) }

// WrapEmpty_Promise is a wrapper for a WrapEmpty promised by a client call.
type WrapEmpty_Promise struct{ *C.Pipeline }

func (p WrapEmpty_Promise) Struct() (WrapEmpty, error) {
	s, err := p.Pipeline.Struct()
	return WrapEmpty{s}, err
}

func (p WrapEmpty_Promise) MightNotBeReallyEmpty() VerEmpty_Promise {
	return VerEmpty_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Wrap2x2 struct{ C.Struct }

func NewWrap2x2(s *C.Segment) (Wrap2x2, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Wrap2x2{}, err
	}
	return Wrap2x2{st}, nil
}

func NewRootWrap2x2(s *C.Segment) (Wrap2x2, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Wrap2x2{}, err
	}
	return Wrap2x2{st}, nil
}

func ReadRootWrap2x2(msg *C.Message) (Wrap2x2, error) {
	root, err := msg.Root()
	if err != nil {
		return Wrap2x2{}, err
	}
	st := C.ToStruct(root)
	return Wrap2x2{st}, nil
}

func (s Wrap2x2) MightNotBeReallyEmpty() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}

	ss := C.ToStruct(p)

	return VerTwoDataTwoPtr{Struct: ss}, nil
}

func (s Wrap2x2) SetMightNotBeReallyEmpty(v VerTwoDataTwoPtr) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s Wrap2x2) NewMightNotBeReallyEmpty() (VerTwoDataTwoPtr, error) {

	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// Wrap2x2_List is a list of Wrap2x2.
type Wrap2x2_List struct{ C.List }

// NewWrap2x2 creates a new list of Wrap2x2.
func NewWrap2x2_List(s *C.Segment, sz int32) (Wrap2x2_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Wrap2x2_List{}, err
	}
	return Wrap2x2_List{l}, nil
}

func (s Wrap2x2_List) At(i int) Wrap2x2           { return Wrap2x2{s.List.Struct(i)} }
func (s Wrap2x2_List) Set(i int, v Wrap2x2) error { return s.List.SetStruct(i, v.Struct) }

// Wrap2x2_Promise is a wrapper for a Wrap2x2 promised by a client call.
type Wrap2x2_Promise struct{ *C.Pipeline }

func (p Wrap2x2_Promise) Struct() (Wrap2x2, error) {
	s, err := p.Pipeline.Struct()
	return Wrap2x2{s}, err
}

func (p Wrap2x2_Promise) MightNotBeReallyEmpty() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Wrap2x2plus struct{ C.Struct }

func NewWrap2x2plus(s *C.Segment) (Wrap2x2plus, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Wrap2x2plus{}, err
	}
	return Wrap2x2plus{st}, nil
}

func NewRootWrap2x2plus(s *C.Segment) (Wrap2x2plus, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Wrap2x2plus{}, err
	}
	return Wrap2x2plus{st}, nil
}

func ReadRootWrap2x2plus(msg *C.Message) (Wrap2x2plus, error) {
	root, err := msg.Root()
	if err != nil {
		return Wrap2x2plus{}, err
	}
	st := C.ToStruct(root)
	return Wrap2x2plus{st}, nil
}

func (s Wrap2x2plus) MightNotBeReallyEmpty() (VerTwoTwoPlus, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerTwoTwoPlus{}, err
	}

	ss := C.ToStruct(p)

	return VerTwoTwoPlus{Struct: ss}, nil
}

func (s Wrap2x2plus) SetMightNotBeReallyEmpty(v VerTwoTwoPlus) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerTwoTwoPlus struct, preferring placement in s's segment.
func (s Wrap2x2plus) NewMightNotBeReallyEmpty() (VerTwoTwoPlus, error) {

	ss, err := NewVerTwoTwoPlus(s.Struct.Segment())
	if err != nil {
		return VerTwoTwoPlus{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// Wrap2x2plus_List is a list of Wrap2x2plus.
type Wrap2x2plus_List struct{ C.List }

// NewWrap2x2plus creates a new list of Wrap2x2plus.
func NewWrap2x2plus_List(s *C.Segment, sz int32) (Wrap2x2plus_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Wrap2x2plus_List{}, err
	}
	return Wrap2x2plus_List{l}, nil
}

func (s Wrap2x2plus_List) At(i int) Wrap2x2plus           { return Wrap2x2plus{s.List.Struct(i)} }
func (s Wrap2x2plus_List) Set(i int, v Wrap2x2plus) error { return s.List.SetStruct(i, v.Struct) }

// Wrap2x2plus_Promise is a wrapper for a Wrap2x2plus promised by a client call.
type Wrap2x2plus_Promise struct{ *C.Pipeline }

func (p Wrap2x2plus_Promise) Struct() (Wrap2x2plus, error) {
	s, err := p.Pipeline.Struct()
	return Wrap2x2plus{s}, err
}

func (p Wrap2x2plus_Promise) MightNotBeReallyEmpty() VerTwoTwoPlus_Promise {
	return VerTwoTwoPlus_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type VoidUnion struct{ C.Struct }
type VoidUnion_Which uint16

const (
	VoidUnion_Which_a VoidUnion_Which = 0
	VoidUnion_Which_b VoidUnion_Which = 1
)

func (w VoidUnion_Which) String() string {
	const s = "ab"
	switch w {
	case VoidUnion_Which_a:
		return s[0:1]
	case VoidUnion_Which_b:
		return s[1:2]

	}
	return "VoidUnion_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewVoidUnion(s *C.Segment) (VoidUnion, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return VoidUnion{}, err
	}
	return VoidUnion{st}, nil
}

func NewRootVoidUnion(s *C.Segment) (VoidUnion, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return VoidUnion{}, err
	}
	return VoidUnion{st}, nil
}

func ReadRootVoidUnion(msg *C.Message) (VoidUnion, error) {
	root, err := msg.Root()
	if err != nil {
		return VoidUnion{}, err
	}
	st := C.ToStruct(root)
	return VoidUnion{st}, nil
}

func (s VoidUnion) Which() VoidUnion_Which {
	return VoidUnion_Which(s.Struct.Uint16(0))
}

func (s VoidUnion) SetA() {
	s.Struct.SetUint16(0, 0)
}

func (s VoidUnion) SetB() {
	s.Struct.SetUint16(0, 1)
}

// VoidUnion_List is a list of VoidUnion.
type VoidUnion_List struct{ C.List }

// NewVoidUnion creates a new list of VoidUnion.
func NewVoidUnion_List(s *C.Segment, sz int32) (VoidUnion_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return VoidUnion_List{}, err
	}
	return VoidUnion_List{l}, nil
}

func (s VoidUnion_List) At(i int) VoidUnion           { return VoidUnion{s.List.Struct(i)} }
func (s VoidUnion_List) Set(i int, v VoidUnion) error { return s.List.SetStruct(i, v.Struct) }

// VoidUnion_Promise is a wrapper for a VoidUnion promised by a client call.
type VoidUnion_Promise struct{ *C.Pipeline }

func (p VoidUnion_Promise) Struct() (VoidUnion, error) {
	s, err := p.Pipeline.Struct()
	return VoidUnion{s}, err
}

type Nester1Capn struct{ C.Struct }

func NewNester1Capn(s *C.Segment) (Nester1Capn, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Nester1Capn{}, err
	}
	return Nester1Capn{st}, nil
}

func NewRootNester1Capn(s *C.Segment) (Nester1Capn, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Nester1Capn{}, err
	}
	return Nester1Capn{st}, nil
}

func ReadRootNester1Capn(msg *C.Message) (Nester1Capn, error) {
	root, err := msg.Root()
	if err != nil {
		return Nester1Capn{}, err
	}
	st := C.ToStruct(root)
	return Nester1Capn{st}, nil
}

func (s Nester1Capn) Strs() (C.TextList, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.TextList{}, err
	}

	l := C.ToList(p)

	return C.TextList{List: l}, nil
}

func (s Nester1Capn) SetStrs(v C.TextList) error {

	return s.Struct.SetPointer(0, v.List)
}

// Nester1Capn_List is a list of Nester1Capn.
type Nester1Capn_List struct{ C.List }

// NewNester1Capn creates a new list of Nester1Capn.
func NewNester1Capn_List(s *C.Segment, sz int32) (Nester1Capn_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Nester1Capn_List{}, err
	}
	return Nester1Capn_List{l}, nil
}

func (s Nester1Capn_List) At(i int) Nester1Capn           { return Nester1Capn{s.List.Struct(i)} }
func (s Nester1Capn_List) Set(i int, v Nester1Capn) error { return s.List.SetStruct(i, v.Struct) }

// Nester1Capn_Promise is a wrapper for a Nester1Capn promised by a client call.
type Nester1Capn_Promise struct{ *C.Pipeline }

func (p Nester1Capn_Promise) Struct() (Nester1Capn, error) {
	s, err := p.Pipeline.Struct()
	return Nester1Capn{s}, err
}

type RWTestCapn struct{ C.Struct }

func NewRWTestCapn(s *C.Segment) (RWTestCapn, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return RWTestCapn{}, err
	}
	return RWTestCapn{st}, nil
}

func NewRootRWTestCapn(s *C.Segment) (RWTestCapn, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return RWTestCapn{}, err
	}
	return RWTestCapn{st}, nil
}

func ReadRootRWTestCapn(msg *C.Message) (RWTestCapn, error) {
	root, err := msg.Root()
	if err != nil {
		return RWTestCapn{}, err
	}
	st := C.ToStruct(root)
	return RWTestCapn{st}, nil
}

func (s RWTestCapn) NestMatrix() (C.PointerList, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return C.PointerList{}, err
	}

	l := C.ToList(p)

	return C.PointerList{List: l}, nil
}

func (s RWTestCapn) SetNestMatrix(v C.PointerList) error {

	return s.Struct.SetPointer(0, v.List)
}

// RWTestCapn_List is a list of RWTestCapn.
type RWTestCapn_List struct{ C.List }

// NewRWTestCapn creates a new list of RWTestCapn.
func NewRWTestCapn_List(s *C.Segment, sz int32) (RWTestCapn_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return RWTestCapn_List{}, err
	}
	return RWTestCapn_List{l}, nil
}

func (s RWTestCapn_List) At(i int) RWTestCapn           { return RWTestCapn{s.List.Struct(i)} }
func (s RWTestCapn_List) Set(i int, v RWTestCapn) error { return s.List.SetStruct(i, v.Struct) }

// RWTestCapn_Promise is a wrapper for a RWTestCapn promised by a client call.
type RWTestCapn_Promise struct{ *C.Pipeline }

func (p RWTestCapn_Promise) Struct() (RWTestCapn, error) {
	s, err := p.Pipeline.Struct()
	return RWTestCapn{s}, err
}

type ListStructCapn struct{ C.Struct }

func NewListStructCapn(s *C.Segment) (ListStructCapn, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ListStructCapn{}, err
	}
	return ListStructCapn{st}, nil
}

func NewRootListStructCapn(s *C.Segment) (ListStructCapn, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ListStructCapn{}, err
	}
	return ListStructCapn{st}, nil
}

func ReadRootListStructCapn(msg *C.Message) (ListStructCapn, error) {
	root, err := msg.Root()
	if err != nil {
		return ListStructCapn{}, err
	}
	st := C.ToStruct(root)
	return ListStructCapn{st}, nil
}

func (s ListStructCapn) Vec() (Nester1Capn_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Nester1Capn_List{}, err
	}

	l := C.ToList(p)

	return Nester1Capn_List{List: l}, nil
}

func (s ListStructCapn) SetVec(v Nester1Capn_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// ListStructCapn_List is a list of ListStructCapn.
type ListStructCapn_List struct{ C.List }

// NewListStructCapn creates a new list of ListStructCapn.
func NewListStructCapn_List(s *C.Segment, sz int32) (ListStructCapn_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return ListStructCapn_List{}, err
	}
	return ListStructCapn_List{l}, nil
}

func (s ListStructCapn_List) At(i int) ListStructCapn           { return ListStructCapn{s.List.Struct(i)} }
func (s ListStructCapn_List) Set(i int, v ListStructCapn) error { return s.List.SetStruct(i, v.Struct) }

// ListStructCapn_Promise is a wrapper for a ListStructCapn promised by a client call.
type ListStructCapn_Promise struct{ *C.Pipeline }

func (p ListStructCapn_Promise) Struct() (ListStructCapn, error) {
	s, err := p.Pipeline.Struct()
	return ListStructCapn{s}, err
}

type Echo struct{ Client C.Client }

func (c Echo) Echo(ctx context.Context, params func(Echo_echo_Params) error, opts ...C.CallOption) Echo_echo_Results_Promise {
	if c.Client == nil {
		return Echo_echo_Results_Promise{Pipeline: C.NewPipeline(C.ErrorAnswer(C.ErrNullClient))}
	}
	return Echo_echo_Results_Promise{Pipeline: C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0x8e5322c1e9282534,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:Echo",
			MethodName:    "echo",
		},
		ParamsSize: C.ObjectSize{DataSize: 0, PointerCount: 1},
		ParamsFunc: func(s C.Struct) error { return params(Echo_echo_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	}))}
}

type Echo_Server interface {
	Echo(Echo_echo) error
}

func Echo_ServerToClient(s Echo_Server) Echo {
	c, _ := s.(server.Closer)
	return Echo{Client: server.New(Echo_Methods(nil, s), c)}
}

func Echo_Methods(methods []server.Method, s Echo_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0x8e5322c1e9282534,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:Echo",
			MethodName:    "echo",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := Echo_echo{c, opts, Echo_echo_Params{Struct: p}, Echo_echo_Results{Struct: r}}
			return s.Echo(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Echo_echo holds the arguments for a server call to Echo.echo.
type Echo_echo struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  Echo_echo_Params
	Results Echo_echo_Results
}

type Echo_echo_Params struct{ C.Struct }

func NewEcho_echo_Params(s *C.Segment) (Echo_echo_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echo_echo_Params{}, err
	}
	return Echo_echo_Params{st}, nil
}

func NewRootEcho_echo_Params(s *C.Segment) (Echo_echo_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echo_echo_Params{}, err
	}
	return Echo_echo_Params{st}, nil
}

func ReadRootEcho_echo_Params(msg *C.Message) (Echo_echo_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return Echo_echo_Params{}, err
	}
	st := C.ToStruct(root)
	return Echo_echo_Params{st}, nil
}

func (s Echo_echo_Params) In() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s Echo_echo_Params) SetIn(v string) error {

	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// Echo_echo_Params_List is a list of Echo_echo_Params.
type Echo_echo_Params_List struct{ C.List }

// NewEcho_echo_Params creates a new list of Echo_echo_Params.
func NewEcho_echo_Params_List(s *C.Segment, sz int32) (Echo_echo_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Echo_echo_Params_List{}, err
	}
	return Echo_echo_Params_List{l}, nil
}

func (s Echo_echo_Params_List) At(i int) Echo_echo_Params { return Echo_echo_Params{s.List.Struct(i)} }
func (s Echo_echo_Params_List) Set(i int, v Echo_echo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Echo_echo_Params_Promise is a wrapper for a Echo_echo_Params promised by a client call.
type Echo_echo_Params_Promise struct{ *C.Pipeline }

func (p Echo_echo_Params_Promise) Struct() (Echo_echo_Params, error) {
	s, err := p.Pipeline.Struct()
	return Echo_echo_Params{s}, err
}

type Echo_echo_Results struct{ C.Struct }

func NewEcho_echo_Results(s *C.Segment) (Echo_echo_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echo_echo_Results{}, err
	}
	return Echo_echo_Results{st}, nil
}

func NewRootEcho_echo_Results(s *C.Segment) (Echo_echo_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Echo_echo_Results{}, err
	}
	return Echo_echo_Results{st}, nil
}

func ReadRootEcho_echo_Results(msg *C.Message) (Echo_echo_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return Echo_echo_Results{}, err
	}
	st := C.ToStruct(root)
	return Echo_echo_Results{st}, nil
}

func (s Echo_echo_Results) Out() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return C.ToText(p), nil

}

func (s Echo_echo_Results) SetOut(v string) error {

	t, err := C.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// Echo_echo_Results_List is a list of Echo_echo_Results.
type Echo_echo_Results_List struct{ C.List }

// NewEcho_echo_Results creates a new list of Echo_echo_Results.
func NewEcho_echo_Results_List(s *C.Segment, sz int32) (Echo_echo_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Echo_echo_Results_List{}, err
	}
	return Echo_echo_Results_List{l}, nil
}

func (s Echo_echo_Results_List) At(i int) Echo_echo_Results {
	return Echo_echo_Results{s.List.Struct(i)}
}
func (s Echo_echo_Results_List) Set(i int, v Echo_echo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Echo_echo_Results_Promise is a wrapper for a Echo_echo_Results promised by a client call.
type Echo_echo_Results_Promise struct{ *C.Pipeline }

func (p Echo_echo_Results_Promise) Struct() (Echo_echo_Results, error) {
	s, err := p.Pipeline.Struct()
	return Echo_echo_Results{s}, err
}

type Hoth struct{ C.Struct }

func NewHoth(s *C.Segment) (Hoth, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Hoth{}, err
	}
	return Hoth{st}, nil
}

func NewRootHoth(s *C.Segment) (Hoth, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Hoth{}, err
	}
	return Hoth{st}, nil
}

func ReadRootHoth(msg *C.Message) (Hoth, error) {
	root, err := msg.Root()
	if err != nil {
		return Hoth{}, err
	}
	st := C.ToStruct(root)
	return Hoth{st}, nil
}

func (s Hoth) Base() (EchoBase, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return EchoBase{}, err
	}

	ss := C.ToStruct(p)

	return EchoBase{Struct: ss}, nil
}

func (s Hoth) SetBase(v EchoBase) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewBase sets the base field to a newly
// allocated EchoBase struct, preferring placement in s's segment.
func (s Hoth) NewBase() (EchoBase, error) {

	ss, err := NewEchoBase(s.Struct.Segment())
	if err != nil {
		return EchoBase{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// Hoth_List is a list of Hoth.
type Hoth_List struct{ C.List }

// NewHoth creates a new list of Hoth.
func NewHoth_List(s *C.Segment, sz int32) (Hoth_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Hoth_List{}, err
	}
	return Hoth_List{l}, nil
}

func (s Hoth_List) At(i int) Hoth           { return Hoth{s.List.Struct(i)} }
func (s Hoth_List) Set(i int, v Hoth) error { return s.List.SetStruct(i, v.Struct) }

// Hoth_Promise is a wrapper for a Hoth promised by a client call.
type Hoth_Promise struct{ *C.Pipeline }

func (p Hoth_Promise) Struct() (Hoth, error) {
	s, err := p.Pipeline.Struct()
	return Hoth{s}, err
}

func (p Hoth_Promise) Base() EchoBase_Promise {
	return EchoBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EchoBase struct{ C.Struct }

func NewEchoBase(s *C.Segment) (EchoBase, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EchoBase{}, err
	}
	return EchoBase{st}, nil
}

func NewRootEchoBase(s *C.Segment) (EchoBase, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EchoBase{}, err
	}
	return EchoBase{st}, nil
}

func ReadRootEchoBase(msg *C.Message) (EchoBase, error) {
	root, err := msg.Root()
	if err != nil {
		return EchoBase{}, err
	}
	st := C.ToStruct(root)
	return EchoBase{st}, nil
}

func (s EchoBase) Echo() Echo {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return Echo{}
	}
	c := C.ToInterface(p).Client()
	return Echo{Client: c}
}

func (s EchoBase) SetEcho(v Echo) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	ci := seg.Message().AddCap(v.Client)
	return s.Struct.SetPointer(0, C.NewInterface(seg, ci))
}

// EchoBase_List is a list of EchoBase.
type EchoBase_List struct{ C.List }

// NewEchoBase creates a new list of EchoBase.
func NewEchoBase_List(s *C.Segment, sz int32) (EchoBase_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EchoBase_List{}, err
	}
	return EchoBase_List{l}, nil
}

func (s EchoBase_List) At(i int) EchoBase           { return EchoBase{s.List.Struct(i)} }
func (s EchoBase_List) Set(i int, v EchoBase) error { return s.List.SetStruct(i, v.Struct) }

// EchoBase_Promise is a wrapper for a EchoBase promised by a client call.
type EchoBase_Promise struct{ *C.Pipeline }

func (p EchoBase_Promise) Struct() (EchoBase, error) {
	s, err := p.Pipeline.Struct()
	return EchoBase{s}, err
}

func (p EchoBase_Promise) Echo() Echo {
	return Echo{Client: p.Pipeline.GetPipeline(0).Client()}
}

type StackingRoot struct{ C.Struct }

func NewStackingRoot(s *C.Segment) (StackingRoot, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return StackingRoot{}, err
	}
	return StackingRoot{st}, nil
}

func NewRootStackingRoot(s *C.Segment) (StackingRoot, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return StackingRoot{}, err
	}
	return StackingRoot{st}, nil
}

func ReadRootStackingRoot(msg *C.Message) (StackingRoot, error) {
	root, err := msg.Root()
	if err != nil {
		return StackingRoot{}, err
	}
	st := C.ToStruct(root)
	return StackingRoot{st}, nil
}

func (s StackingRoot) A() (StackingA, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return StackingA{}, err
	}

	ss := C.ToStruct(p)

	return StackingA{Struct: ss}, nil
}

func (s StackingRoot) SetA(v StackingA) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewA sets the a field to a newly
// allocated StackingA struct, preferring placement in s's segment.
func (s StackingRoot) NewA() (StackingA, error) {

	ss, err := NewStackingA(s.Struct.Segment())
	if err != nil {
		return StackingA{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s StackingRoot) AWithDefault() (StackingA, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return StackingA{}, err
	}

	ss, err := C.ToStructDefault(p, x_832bcc6686a26d56[0:32])
	if err != nil {
		return StackingA{}, err
	}

	return StackingA{Struct: ss}, nil
}

func (s StackingRoot) SetAWithDefault(v StackingA) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewAWithDefault sets the aWithDefault field to a newly
// allocated StackingA struct, preferring placement in s's segment.
func (s StackingRoot) NewAWithDefault() (StackingA, error) {

	ss, err := NewStackingA(s.Struct.Segment())
	if err != nil {
		return StackingA{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// StackingRoot_List is a list of StackingRoot.
type StackingRoot_List struct{ C.List }

// NewStackingRoot creates a new list of StackingRoot.
func NewStackingRoot_List(s *C.Segment, sz int32) (StackingRoot_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return StackingRoot_List{}, err
	}
	return StackingRoot_List{l}, nil
}

func (s StackingRoot_List) At(i int) StackingRoot           { return StackingRoot{s.List.Struct(i)} }
func (s StackingRoot_List) Set(i int, v StackingRoot) error { return s.List.SetStruct(i, v.Struct) }

// StackingRoot_Promise is a wrapper for a StackingRoot promised by a client call.
type StackingRoot_Promise struct{ *C.Pipeline }

func (p StackingRoot_Promise) Struct() (StackingRoot, error) {
	s, err := p.Pipeline.Struct()
	return StackingRoot{s}, err
}

func (p StackingRoot_Promise) A() StackingA_Promise {
	return StackingA_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p StackingRoot_Promise) AWithDefault() StackingA_Promise {
	return StackingA_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_832bcc6686a26d56[32:64])}
}

type StackingA struct{ C.Struct }

func NewStackingA(s *C.Segment) (StackingA, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return StackingA{}, err
	}
	return StackingA{st}, nil
}

func NewRootStackingA(s *C.Segment) (StackingA, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return StackingA{}, err
	}
	return StackingA{st}, nil
}

func ReadRootStackingA(msg *C.Message) (StackingA, error) {
	root, err := msg.Root()
	if err != nil {
		return StackingA{}, err
	}
	st := C.ToStruct(root)
	return StackingA{st}, nil
}

func (s StackingA) Num() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s StackingA) SetNum(v int32) {

	s.Struct.SetUint32(0, uint32(v))
}

func (s StackingA) B() (StackingB, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return StackingB{}, err
	}

	ss := C.ToStruct(p)

	return StackingB{Struct: ss}, nil
}

func (s StackingA) SetB(v StackingB) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewB sets the b field to a newly
// allocated StackingB struct, preferring placement in s's segment.
func (s StackingA) NewB() (StackingB, error) {

	ss, err := NewStackingB(s.Struct.Segment())
	if err != nil {
		return StackingB{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// StackingA_List is a list of StackingA.
type StackingA_List struct{ C.List }

// NewStackingA creates a new list of StackingA.
func NewStackingA_List(s *C.Segment, sz int32) (StackingA_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return StackingA_List{}, err
	}
	return StackingA_List{l}, nil
}

func (s StackingA_List) At(i int) StackingA           { return StackingA{s.List.Struct(i)} }
func (s StackingA_List) Set(i int, v StackingA) error { return s.List.SetStruct(i, v.Struct) }

// StackingA_Promise is a wrapper for a StackingA promised by a client call.
type StackingA_Promise struct{ *C.Pipeline }

func (p StackingA_Promise) Struct() (StackingA, error) {
	s, err := p.Pipeline.Struct()
	return StackingA{s}, err
}

func (p StackingA_Promise) B() StackingB_Promise {
	return StackingB_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type StackingB struct{ C.Struct }

func NewStackingB(s *C.Segment) (StackingB, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return StackingB{}, err
	}
	return StackingB{st}, nil
}

func NewRootStackingB(s *C.Segment) (StackingB, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return StackingB{}, err
	}
	return StackingB{st}, nil
}

func ReadRootStackingB(msg *C.Message) (StackingB, error) {
	root, err := msg.Root()
	if err != nil {
		return StackingB{}, err
	}
	st := C.ToStruct(root)
	return StackingB{st}, nil
}

func (s StackingB) Num() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s StackingB) SetNum(v int32) {

	s.Struct.SetUint32(0, uint32(v))
}

// StackingB_List is a list of StackingB.
type StackingB_List struct{ C.List }

// NewStackingB creates a new list of StackingB.
func NewStackingB_List(s *C.Segment, sz int32) (StackingB_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return StackingB_List{}, err
	}
	return StackingB_List{l}, nil
}

func (s StackingB_List) At(i int) StackingB           { return StackingB{s.List.Struct(i)} }
func (s StackingB_List) Set(i int, v StackingB) error { return s.List.SetStruct(i, v.Struct) }

// StackingB_Promise is a wrapper for a StackingB promised by a client call.
type StackingB_Promise struct{ *C.Pipeline }

func (p StackingB_Promise) Struct() (StackingB, error) {
	s, err := p.Pipeline.Struct()
	return StackingB{s}, err
}

type CallSequence struct{ Client C.Client }

func (c CallSequence) GetNumber(ctx context.Context, params func(CallSequence_getNumber_Params) error, opts ...C.CallOption) CallSequence_getNumber_Results_Promise {
	if c.Client == nil {
		return CallSequence_getNumber_Results_Promise{Pipeline: C.NewPipeline(C.ErrorAnswer(C.ErrNullClient))}
	}
	return CallSequence_getNumber_Results_Promise{Pipeline: C.NewPipeline(c.Client.Call(&C.Call{
		Ctx: ctx,
		Method: C.Method{

			InterfaceID:   0xabaedf5f7817c820,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:CallSequence",
			MethodName:    "getNumber",
		},
		ParamsSize: C.ObjectSize{DataSize: 0, PointerCount: 0},
		ParamsFunc: func(s C.Struct) error { return params(CallSequence_getNumber_Params{Struct: s}) },
		Options:    C.NewCallOptions(opts),
	}))}
}

type CallSequence_Server interface {
	GetNumber(CallSequence_getNumber) error
}

func CallSequence_ServerToClient(s CallSequence_Server) CallSequence {
	c, _ := s.(server.Closer)
	return CallSequence{Client: server.New(CallSequence_Methods(nil, s), c)}
}

func CallSequence_Methods(methods []server.Method, s CallSequence_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: C.Method{

			InterfaceID:   0xabaedf5f7817c820,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:CallSequence",
			MethodName:    "getNumber",
		},
		Impl: func(c context.Context, opts C.CallOptions, p, r C.Struct) error {
			call := CallSequence_getNumber{c, opts, CallSequence_getNumber_Params{Struct: p}, CallSequence_getNumber_Results{Struct: r}}
			return s.GetNumber(call)
		},
		ResultsSize: C.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// CallSequence_getNumber holds the arguments for a server call to CallSequence.getNumber.
type CallSequence_getNumber struct {
	Ctx     context.Context
	Options C.CallOptions
	Params  CallSequence_getNumber_Params
	Results CallSequence_getNumber_Results
}

type CallSequence_getNumber_Params struct{ C.Struct }

func NewCallSequence_getNumber_Params(s *C.Segment) (CallSequence_getNumber_Params, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return CallSequence_getNumber_Params{}, err
	}
	return CallSequence_getNumber_Params{st}, nil
}

func NewRootCallSequence_getNumber_Params(s *C.Segment) (CallSequence_getNumber_Params, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return CallSequence_getNumber_Params{}, err
	}
	return CallSequence_getNumber_Params{st}, nil
}

func ReadRootCallSequence_getNumber_Params(msg *C.Message) (CallSequence_getNumber_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return CallSequence_getNumber_Params{}, err
	}
	st := C.ToStruct(root)
	return CallSequence_getNumber_Params{st}, nil
}

// CallSequence_getNumber_Params_List is a list of CallSequence_getNumber_Params.
type CallSequence_getNumber_Params_List struct{ C.List }

// NewCallSequence_getNumber_Params creates a new list of CallSequence_getNumber_Params.
func NewCallSequence_getNumber_Params_List(s *C.Segment, sz int32) (CallSequence_getNumber_Params_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return CallSequence_getNumber_Params_List{}, err
	}
	return CallSequence_getNumber_Params_List{l}, nil
}

func (s CallSequence_getNumber_Params_List) At(i int) CallSequence_getNumber_Params {
	return CallSequence_getNumber_Params{s.List.Struct(i)}
}
func (s CallSequence_getNumber_Params_List) Set(i int, v CallSequence_getNumber_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// CallSequence_getNumber_Params_Promise is a wrapper for a CallSequence_getNumber_Params promised by a client call.
type CallSequence_getNumber_Params_Promise struct{ *C.Pipeline }

func (p CallSequence_getNumber_Params_Promise) Struct() (CallSequence_getNumber_Params, error) {
	s, err := p.Pipeline.Struct()
	return CallSequence_getNumber_Params{s}, err
}

type CallSequence_getNumber_Results struct{ C.Struct }

func NewCallSequence_getNumber_Results(s *C.Segment) (CallSequence_getNumber_Results, error) {
	st, err := C.NewStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return CallSequence_getNumber_Results{}, err
	}
	return CallSequence_getNumber_Results{st}, nil
}

func NewRootCallSequence_getNumber_Results(s *C.Segment) (CallSequence_getNumber_Results, error) {
	st, err := C.NewRootStruct(s, C.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return CallSequence_getNumber_Results{}, err
	}
	return CallSequence_getNumber_Results{st}, nil
}

func ReadRootCallSequence_getNumber_Results(msg *C.Message) (CallSequence_getNumber_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return CallSequence_getNumber_Results{}, err
	}
	st := C.ToStruct(root)
	return CallSequence_getNumber_Results{st}, nil
}

func (s CallSequence_getNumber_Results) N() uint32 {
	return s.Struct.Uint32(0)
}

func (s CallSequence_getNumber_Results) SetN(v uint32) {

	s.Struct.SetUint32(0, v)
}

// CallSequence_getNumber_Results_List is a list of CallSequence_getNumber_Results.
type CallSequence_getNumber_Results_List struct{ C.List }

// NewCallSequence_getNumber_Results creates a new list of CallSequence_getNumber_Results.
func NewCallSequence_getNumber_Results_List(s *C.Segment, sz int32) (CallSequence_getNumber_Results_List, error) {
	l, err := C.NewCompositeList(s, C.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return CallSequence_getNumber_Results_List{}, err
	}
	return CallSequence_getNumber_Results_List{l}, nil
}

func (s CallSequence_getNumber_Results_List) At(i int) CallSequence_getNumber_Results {
	return CallSequence_getNumber_Results{s.List.Struct(i)}
}
func (s CallSequence_getNumber_Results_List) Set(i int, v CallSequence_getNumber_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// CallSequence_getNumber_Results_Promise is a wrapper for a CallSequence_getNumber_Results promised by a client call.
type CallSequence_getNumber_Results_Promise struct{ *C.Pipeline }

func (p CallSequence_getNumber_Results_Promise) Struct() (CallSequence_getNumber_Results, error) {
	s, err := p.Pipeline.Struct()
	return CallSequence_getNumber_Results{s}, err
}

var x_832bcc6686a26d56 = []byte{
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}
