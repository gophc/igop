// export by github.com/goplus/gossa/cmd/qexp

package rand

import (
	q "math/rand"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "rand",
		Path: "math/rand",
		Deps: map[string]string{
			"math": "math",
			"sync": "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Source":   reflect.TypeOf((*q.Source)(nil)).Elem(),
			"Source64": reflect.TypeOf((*q.Source64)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Rand": {reflect.TypeOf((*q.Rand)(nil)).Elem(), "", "ExpFloat64,Float32,Float64,Int,Int31,Int31n,Int63,Int63n,Intn,NormFloat64,Perm,Read,Seed,Shuffle,Uint32,Uint64,int31n"},
			"Zipf": {reflect.TypeOf((*q.Zipf)(nil)).Elem(), "", "Uint64,h,hinv"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ExpFloat64":  reflect.ValueOf(q.ExpFloat64),
			"Float32":     reflect.ValueOf(q.Float32),
			"Float64":     reflect.ValueOf(q.Float64),
			"Int":         reflect.ValueOf(q.Int),
			"Int31":       reflect.ValueOf(q.Int31),
			"Int31n":      reflect.ValueOf(q.Int31n),
			"Int63":       reflect.ValueOf(q.Int63),
			"Int63n":      reflect.ValueOf(q.Int63n),
			"Intn":        reflect.ValueOf(q.Intn),
			"New":         reflect.ValueOf(q.New),
			"NewSource":   reflect.ValueOf(q.NewSource),
			"NewZipf":     reflect.ValueOf(q.NewZipf),
			"NormFloat64": reflect.ValueOf(q.NormFloat64),
			"Perm":        reflect.ValueOf(q.Perm),
			"Read":        reflect.ValueOf(q.Read),
			"Seed":        reflect.ValueOf(q.Seed),
			"Shuffle":     reflect.ValueOf(q.Shuffle),
			"Uint32":      reflect.ValueOf(q.Uint32),
			"Uint64":      reflect.ValueOf(q.Uint64),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
