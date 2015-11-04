// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_with_list_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("StructWithList",
			[]types.Field{
				types.Field{"l", types.MakeCompoundTypeRef(types.ListKind, types.MakePrimitiveTypeRef(types.UInt8Kind)), false},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"i", types.MakePrimitiveTypeRef(types.Int64Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_with_list_CachedRef = types.RegisterPackage(&p)
}

// StructWithList

type StructWithList struct {
	_l ListOfUInt8
	_b bool
	_s string
	_i int64

	ref *ref.Ref
}

func NewStructWithList() StructWithList {
	return StructWithList{
		_l: NewListOfUInt8(),
		_b: false,
		_s: "",
		_i: int64(0),

		ref: &ref.Ref{},
	}
}

type StructWithListDef struct {
	L ListOfUInt8Def
	B bool
	S string
	I int64
}

func (def StructWithListDef) New() StructWithList {
	return StructWithList{
		_l:  def.L.New(),
		_b:  def.B,
		_s:  def.S,
		_i:  def.I,
		ref: &ref.Ref{},
	}
}

func (s StructWithList) Def() (d StructWithListDef) {
	d.L = s._l.Def()
	d.B = s._b
	d.S = s._s
	d.I = s._i
	return
}

var __typeRefForStructWithList types.TypeRef
var __typeDefForStructWithList types.TypeRef

func (m StructWithList) TypeRef() types.TypeRef {
	return __typeRefForStructWithList
}

func init() {
	__typeRefForStructWithList = types.MakeTypeRef(__genPackageInFile_struct_with_list_CachedRef, 0)
	__typeDefForStructWithList = types.MakeStructTypeRef("StructWithList",
		[]types.Field{
			types.Field{"l", types.MakeCompoundTypeRef(types.ListKind, types.MakePrimitiveTypeRef(types.UInt8Kind)), false},
			types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
			types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
			types.Field{"i", types.MakePrimitiveTypeRef(types.Int64Kind), false},
		},
		types.Choices{},
	)
	types.RegisterStructBuilder(__typeRefForStructWithList, builderForStructWithList)
}

func (s StructWithList) InternalImplementation() types.Struct {
	// TODO: Remove this
	m := map[string]types.Value{
		"l": s._l,
		"b": types.Bool(s._b),
		"s": types.NewString(s._s),
		"i": types.Int64(s._i),
	}
	return types.NewStruct(__typeRefForStructWithList, __typeDefForStructWithList, m)
}

func builderForStructWithList() chan types.Value {
	c := make(chan types.Value)
	s := StructWithList{ref: &ref.Ref{}}
	go func() {
		s._l = (<-c).(ListOfUInt8)
		s._b = bool((<-c).(types.Bool))
		s._s = (<-c).(types.String).String()
		s._i = int64((<-c).(types.Int64))

		c <- s
	}()
	return c
}

func (s StructWithList) Equals(other types.Value) bool {
	return other != nil && __typeRefForStructWithList.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s StructWithList) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructWithList) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForStructWithList.Chunks()...)
	chunks = append(chunks, s._l.Chunks()...)
	return
}

func (s StructWithList) L() ListOfUInt8 {
	return s._l
}

func (s StructWithList) SetL(val ListOfUInt8) StructWithList {
	s._l = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithList) B() bool {
	return s._b
}

func (s StructWithList) SetB(val bool) StructWithList {
	s._b = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithList) S() string {
	return s._s
}

func (s StructWithList) SetS(val string) StructWithList {
	s._s = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithList) I() int64 {
	return s._i
}

func (s StructWithList) SetI(val int64) StructWithList {
	s._i = val
	s.ref = &ref.Ref{}
	return s
}