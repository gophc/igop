// export by github.com/goplus/gossa/cmd/qexp

package types

import (
	q "go/types"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "types",
		Path: "go/types",
		Deps: map[string]string{
			"bytes":          "bytes",
			"container/heap": "heap",
			"errors":         "errors",
			"fmt":            "fmt",
			"go/ast":         "ast",
			"go/constant":    "constant",
			"go/parser":      "parser",
			"go/token":       "token",
			"io":             "io",
			"math":           "math",
			"sort":           "sort",
			"strconv":        "strconv",
			"strings":        "strings",
			"unicode":        "unicode",
		},
		Interfaces: map[string]reflect.Type{
			"Importer":     reflect.TypeOf((*q.Importer)(nil)).Elem(),
			"ImporterFrom": reflect.TypeOf((*q.ImporterFrom)(nil)).Elem(),
			"Object":       reflect.TypeOf((*q.Object)(nil)).Elem(),
			"Sizes":        reflect.TypeOf((*q.Sizes)(nil)).Elem(),
			"Type":         reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Array":         {reflect.TypeOf((*q.Array)(nil)).Elem(), "", "Elem,Len,String,Underlying"},
			"Basic":         {reflect.TypeOf((*q.Basic)(nil)).Elem(), "", "Info,Kind,Name,String,Underlying"},
			"BasicInfo":     {reflect.TypeOf((*q.BasicInfo)(nil)).Elem(), "", ""},
			"BasicKind":     {reflect.TypeOf((*q.BasicKind)(nil)).Elem(), "", ""},
			"Builtin":       {reflect.TypeOf((*q.Builtin)(nil)).Elem(), "", "String"},
			"Chan":          {reflect.TypeOf((*q.Chan)(nil)).Elem(), "", "Dir,Elem,String,Underlying"},
			"ChanDir":       {reflect.TypeOf((*q.ChanDir)(nil)).Elem(), "", ""},
			"Checker":       {reflect.TypeOf((*q.Checker)(nil)).Elem(), "", "Files,addDeclDep,addMethodDecls,addUnusedDotImport,argument,arguments,arityMatch,arrayLength,assertableTo,assignVar,assignVars,assignment,atEnd,binary,blockBranches,builtin,call,canConvertUntyped,caseTypes,caseValues,checkFiles,closeScope,collectObjects,collectParams,comparison,completeInterface,consolidateMultiples,constDecl,conversion,convertUntyped,cycle,cycleError,declStmt,declare,declareInSet,declarePkgObj,definedType,dump,err,error,errorf,expr,exprInternal,exprOrType,exprWithHint,filename,funcBody,funcDecl,funcType,handleBailout,ident,identical,identical0,identicalIgnoreTags,implicitType,importPackage,index,indexedElts,initConst,initFiles,initOrder,initVar,initVars,interfaceType,invalidAST,invalidArg,invalidOp,isRepresentable,isTerminating,isTerminatingList,isTerminatingSwitch,labels,later,lookupFieldOrMethod,lookupType,missingMethod,multiExpr,multipleDefaults,newError,newErrorf,objDecl,op,openScope,packageObjects,pop,processDelayed,processFinals,push,qualifier,rawExpr,rawLookupFieldOrMethod,recordBuiltinType,recordCommaOkTypes,recordDef,recordImplicit,recordScope,recordSelection,recordTypeAndValue,recordUntyped,recordUse,rememberUntyped,reportAltDecl,reportCycle,representable,resolveBaseTypeName,selector,shift,shortVarDecl,simpleStmt,singleValue,softErrorf,sprintf,stmt,stmtList,structType,suspendedCall,tag,trace,typ,typInternal,typOrNil,typeAssertion,typeDecl,unary,underlying,unusedImports,updateExprType,updateExprVal,usage,use,useGetter,useLHS,validType,varDecl,walkDecl,walkDecls"},
			"Config":        {reflect.TypeOf((*q.Config)(nil)).Elem(), "", "Check,alignof,offsetof,offsetsof,sizeof"},
			"Const":         {reflect.TypeOf((*q.Const)(nil)).Elem(), "", "String,Val,isDependency"},
			"Error":         {reflect.TypeOf((*q.Error)(nil)).Elem(), "Error", ""},
			"Func":          {reflect.TypeOf((*q.Func)(nil)).Elem(), "", "FullName,Scope,String,isDependency"},
			"ImportMode":    {reflect.TypeOf((*q.ImportMode)(nil)).Elem(), "", ""},
			"Info":          {reflect.TypeOf((*q.Info)(nil)).Elem(), "", "ObjectOf,TypeOf"},
			"Initializer":   {reflect.TypeOf((*q.Initializer)(nil)).Elem(), "", "String"},
			"Interface":     {reflect.TypeOf((*q.Interface)(nil)).Elem(), "", "Complete,Embedded,EmbeddedType,Empty,ExplicitMethod,Method,NumEmbeddeds,NumExplicitMethods,NumMethods,String,Underlying,assertCompleteness"},
			"Label":         {reflect.TypeOf((*q.Label)(nil)).Elem(), "", "String"},
			"Map":           {reflect.TypeOf((*q.Map)(nil)).Elem(), "", "Elem,Key,String,Underlying"},
			"MethodSet":     {reflect.TypeOf((*q.MethodSet)(nil)).Elem(), "", "At,Len,Lookup,String"},
			"Named":         {reflect.TypeOf((*q.Named)(nil)).Elem(), "", "AddMethod,Method,NumMethods,Obj,SetUnderlying,String,Underlying,setUnderlying"},
			"Nil":           {reflect.TypeOf((*q.Nil)(nil)).Elem(), "", "String"},
			"Package":       {reflect.TypeOf((*q.Package)(nil)).Elem(), "", "Complete,Imports,MarkComplete,Name,Path,Scope,SetImports,SetName,String"},
			"PkgName":       {reflect.TypeOf((*q.PkgName)(nil)).Elem(), "", "Imported,String"},
			"Pointer":       {reflect.TypeOf((*q.Pointer)(nil)).Elem(), "", "Elem,String,Underlying"},
			"Qualifier":     {reflect.TypeOf((*q.Qualifier)(nil)).Elem(), "", ""},
			"Scope":         {reflect.TypeOf((*q.Scope)(nil)).Elem(), "", "Child,Contains,End,Innermost,Insert,Len,Lookup,LookupParent,Names,NumChildren,Parent,Pos,String,WriteTo"},
			"Selection":     {reflect.TypeOf((*q.Selection)(nil)).Elem(), "", "Index,Indirect,Kind,Obj,Recv,String,Type"},
			"SelectionKind": {reflect.TypeOf((*q.SelectionKind)(nil)).Elem(), "", ""},
			"Signature":     {reflect.TypeOf((*q.Signature)(nil)).Elem(), "", "Params,Recv,Results,String,Underlying,Variadic"},
			"Slice":         {reflect.TypeOf((*q.Slice)(nil)).Elem(), "", "Elem,String,Underlying"},
			"StdSizes":      {reflect.TypeOf((*q.StdSizes)(nil)).Elem(), "", "Alignof,Offsetsof,Sizeof"},
			"Struct":        {reflect.TypeOf((*q.Struct)(nil)).Elem(), "", "Field,NumFields,String,Tag,Underlying"},
			"Tuple":         {reflect.TypeOf((*q.Tuple)(nil)).Elem(), "", "At,Len,String,Underlying"},
			"TypeAndValue":  {reflect.TypeOf((*q.TypeAndValue)(nil)).Elem(), "Addressable,Assignable,HasOk,IsBuiltin,IsNil,IsType,IsValue,IsVoid", ""},
			"TypeName":      {reflect.TypeOf((*q.TypeName)(nil)).Elem(), "", "IsAlias,String"},
			"Var":           {reflect.TypeOf((*q.Var)(nil)).Elem(), "", "Anonymous,Embedded,IsField,String,isDependency"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Typ":      reflect.ValueOf(&q.Typ),
			"Universe": reflect.ValueOf(&q.Universe),
			"Unsafe":   reflect.ValueOf(&q.Unsafe),
		},
		Funcs: map[string]reflect.Value{
			"AssertableTo":            reflect.ValueOf(q.AssertableTo),
			"AssignableTo":            reflect.ValueOf(q.AssignableTo),
			"CheckExpr":               reflect.ValueOf(q.CheckExpr),
			"Comparable":              reflect.ValueOf(q.Comparable),
			"ConvertibleTo":           reflect.ValueOf(q.ConvertibleTo),
			"DefPredeclaredTestFuncs": reflect.ValueOf(q.DefPredeclaredTestFuncs),
			"Default":                 reflect.ValueOf(q.Default),
			"Eval":                    reflect.ValueOf(q.Eval),
			"ExprString":              reflect.ValueOf(q.ExprString),
			"Id":                      reflect.ValueOf(q.Id),
			"Identical":               reflect.ValueOf(q.Identical),
			"IdenticalIgnoreTags":     reflect.ValueOf(q.IdenticalIgnoreTags),
			"Implements":              reflect.ValueOf(q.Implements),
			"IsInterface":             reflect.ValueOf(q.IsInterface),
			"LookupFieldOrMethod":     reflect.ValueOf(q.LookupFieldOrMethod),
			"MissingMethod":           reflect.ValueOf(q.MissingMethod),
			"NewArray":                reflect.ValueOf(q.NewArray),
			"NewChan":                 reflect.ValueOf(q.NewChan),
			"NewChecker":              reflect.ValueOf(q.NewChecker),
			"NewConst":                reflect.ValueOf(q.NewConst),
			"NewField":                reflect.ValueOf(q.NewField),
			"NewFunc":                 reflect.ValueOf(q.NewFunc),
			"NewInterface":            reflect.ValueOf(q.NewInterface),
			"NewInterfaceType":        reflect.ValueOf(q.NewInterfaceType),
			"NewLabel":                reflect.ValueOf(q.NewLabel),
			"NewMap":                  reflect.ValueOf(q.NewMap),
			"NewMethodSet":            reflect.ValueOf(q.NewMethodSet),
			"NewNamed":                reflect.ValueOf(q.NewNamed),
			"NewPackage":              reflect.ValueOf(q.NewPackage),
			"NewParam":                reflect.ValueOf(q.NewParam),
			"NewPkgName":              reflect.ValueOf(q.NewPkgName),
			"NewPointer":              reflect.ValueOf(q.NewPointer),
			"NewScope":                reflect.ValueOf(q.NewScope),
			"NewSignature":            reflect.ValueOf(q.NewSignature),
			"NewSlice":                reflect.ValueOf(q.NewSlice),
			"NewStruct":               reflect.ValueOf(q.NewStruct),
			"NewTuple":                reflect.ValueOf(q.NewTuple),
			"NewTypeName":             reflect.ValueOf(q.NewTypeName),
			"NewVar":                  reflect.ValueOf(q.NewVar),
			"ObjectString":            reflect.ValueOf(q.ObjectString),
			"RelativeTo":              reflect.ValueOf(q.RelativeTo),
			"SelectionString":         reflect.ValueOf(q.SelectionString),
			"SizesFor":                reflect.ValueOf(q.SizesFor),
			"TypeString":              reflect.ValueOf(q.TypeString),
			"WriteExpr":               reflect.ValueOf(q.WriteExpr),
			"WriteSignature":          reflect.ValueOf(q.WriteSignature),
			"WriteType":               reflect.ValueOf(q.WriteType),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"Bool":           {reflect.TypeOf(q.Bool), constant.MakeInt64(1)},
			"Byte":           {reflect.TypeOf(q.Byte), constant.MakeInt64(8)},
			"Complex128":     {reflect.TypeOf(q.Complex128), constant.MakeInt64(16)},
			"Complex64":      {reflect.TypeOf(q.Complex64), constant.MakeInt64(15)},
			"FieldVal":       {reflect.TypeOf(q.FieldVal), constant.MakeInt64(0)},
			"Float32":        {reflect.TypeOf(q.Float32), constant.MakeInt64(13)},
			"Float64":        {reflect.TypeOf(q.Float64), constant.MakeInt64(14)},
			"Int":            {reflect.TypeOf(q.Int), constant.MakeInt64(2)},
			"Int16":          {reflect.TypeOf(q.Int16), constant.MakeInt64(4)},
			"Int32":          {reflect.TypeOf(q.Int32), constant.MakeInt64(5)},
			"Int64":          {reflect.TypeOf(q.Int64), constant.MakeInt64(6)},
			"Int8":           {reflect.TypeOf(q.Int8), constant.MakeInt64(3)},
			"Invalid":        {reflect.TypeOf(q.Invalid), constant.MakeInt64(0)},
			"IsBoolean":      {reflect.TypeOf(q.IsBoolean), constant.MakeInt64(1)},
			"IsComplex":      {reflect.TypeOf(q.IsComplex), constant.MakeInt64(16)},
			"IsConstType":    {reflect.TypeOf(q.IsConstType), constant.MakeInt64(59)},
			"IsFloat":        {reflect.TypeOf(q.IsFloat), constant.MakeInt64(8)},
			"IsInteger":      {reflect.TypeOf(q.IsInteger), constant.MakeInt64(2)},
			"IsNumeric":      {reflect.TypeOf(q.IsNumeric), constant.MakeInt64(26)},
			"IsOrdered":      {reflect.TypeOf(q.IsOrdered), constant.MakeInt64(42)},
			"IsString":       {reflect.TypeOf(q.IsString), constant.MakeInt64(32)},
			"IsUnsigned":     {reflect.TypeOf(q.IsUnsigned), constant.MakeInt64(4)},
			"IsUntyped":      {reflect.TypeOf(q.IsUntyped), constant.MakeInt64(64)},
			"MethodExpr":     {reflect.TypeOf(q.MethodExpr), constant.MakeInt64(2)},
			"MethodVal":      {reflect.TypeOf(q.MethodVal), constant.MakeInt64(1)},
			"RecvOnly":       {reflect.TypeOf(q.RecvOnly), constant.MakeInt64(2)},
			"Rune":           {reflect.TypeOf(q.Rune), constant.MakeInt64(5)},
			"SendOnly":       {reflect.TypeOf(q.SendOnly), constant.MakeInt64(1)},
			"SendRecv":       {reflect.TypeOf(q.SendRecv), constant.MakeInt64(0)},
			"String":         {reflect.TypeOf(q.String), constant.MakeInt64(17)},
			"Uint":           {reflect.TypeOf(q.Uint), constant.MakeInt64(7)},
			"Uint16":         {reflect.TypeOf(q.Uint16), constant.MakeInt64(9)},
			"Uint32":         {reflect.TypeOf(q.Uint32), constant.MakeInt64(10)},
			"Uint64":         {reflect.TypeOf(q.Uint64), constant.MakeInt64(11)},
			"Uint8":          {reflect.TypeOf(q.Uint8), constant.MakeInt64(8)},
			"Uintptr":        {reflect.TypeOf(q.Uintptr), constant.MakeInt64(12)},
			"UnsafePointer":  {reflect.TypeOf(q.UnsafePointer), constant.MakeInt64(18)},
			"UntypedBool":    {reflect.TypeOf(q.UntypedBool), constant.MakeInt64(19)},
			"UntypedComplex": {reflect.TypeOf(q.UntypedComplex), constant.MakeInt64(23)},
			"UntypedFloat":   {reflect.TypeOf(q.UntypedFloat), constant.MakeInt64(22)},
			"UntypedInt":     {reflect.TypeOf(q.UntypedInt), constant.MakeInt64(20)},
			"UntypedNil":     {reflect.TypeOf(q.UntypedNil), constant.MakeInt64(25)},
			"UntypedRune":    {reflect.TypeOf(q.UntypedRune), constant.MakeInt64(21)},
			"UntypedString":  {reflect.TypeOf(q.UntypedString), constant.MakeInt64(24)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
