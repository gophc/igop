// export by github.com/goplus/gossa/cmd/qexp

package dsa

import (
	q "crypto/dsa"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "dsa",
		Path: "crypto/dsa",
		Deps: map[string]string{
			"crypto/internal/randutil": "randutil",
			"errors":                   "errors",
			"io":                       "io",
			"math/big":                 "big",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"ParameterSizes": {reflect.TypeOf((*q.ParameterSizes)(nil)).Elem(), "", ""},
			"Parameters":     {reflect.TypeOf((*q.Parameters)(nil)).Elem(), "", ""},
			"PrivateKey":     {reflect.TypeOf((*q.PrivateKey)(nil)).Elem(), "", ""},
			"PublicKey":      {reflect.TypeOf((*q.PublicKey)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrInvalidPublicKey": reflect.ValueOf(&q.ErrInvalidPublicKey),
		},
		Funcs: map[string]reflect.Value{
			"GenerateKey":        reflect.ValueOf(q.GenerateKey),
			"GenerateParameters": reflect.ValueOf(q.GenerateParameters),
			"Sign":               reflect.ValueOf(q.Sign),
			"Verify":             reflect.ValueOf(q.Verify),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"L1024N160": {reflect.TypeOf(q.L1024N160), constant.MakeInt64(0)},
			"L2048N224": {reflect.TypeOf(q.L2048N224), constant.MakeInt64(1)},
			"L2048N256": {reflect.TypeOf(q.L2048N256), constant.MakeInt64(2)},
			"L3072N256": {reflect.TypeOf(q.L3072N256), constant.MakeInt64(3)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
