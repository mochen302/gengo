// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by defaulter-gen. DO NOT EDIT.

package marker

import (
	"encoding/json"

	external "github.com/mochen302/gengox/examples/defaulter-gen/_output_tests/marker/external"
	externalexternal "github.com/mochen302/gengox/examples/defaulter-gen/_output_tests/marker/external/external"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Defaulted{}, func(obj interface{}) { SetObjectDefaults_Defaulted(obj.(*Defaulted)) })
	scheme.AddTypeDefaultingFunc(&DefaultedWithFunction{}, func(obj interface{}) { SetObjectDefaults_DefaultedWithFunction(obj.(*DefaultedWithFunction)) })
	return nil
}

func SetObjectDefaults_Defaulted(in *Defaulted) {
	if in.StringDefault == "" {
		in.StringDefault = "bar"
	}
	if in.StringPointer == nil {
		var ptrVar1 string = "default"
		in.StringPointer = &ptrVar1
	}
	if in.IntDefault == 0 {
		in.IntDefault = 1
	}
	if in.FloatDefault == 0 {
		in.FloatDefault = 0.5
	}
	if in.List == nil {
		if err := json.Unmarshal([]byte(`["foo", "bar"]`), &in.List); err != nil {
			panic(err)
		}
	}
	for i := range in.List {
		if in.List[i] == nil {
			var ptrVar1 string = "apple"
			in.List[i] = &ptrVar1
		}
	}
	if in.Sub == nil {
		if err := json.Unmarshal([]byte(`{"s": "foo", "i": 5}`), &in.Sub); err != nil {
			panic(err)
		}
	}
	if in.Sub != nil {
		if in.Sub.I == 0 {
			in.Sub.I = 1
		}
	}
	if in.StructList == nil {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.StructList); err != nil {
			panic(err)
		}
	}
	for i := range in.StructList {
		a := &in.StructList[i]
		if a.I == 0 {
			a.I = 1
		}
	}
	if in.PtrStructList == nil {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.PtrStructList); err != nil {
			panic(err)
		}
	}
	for i := range in.PtrStructList {
		a := in.PtrStructList[i]
		if a != nil {
			if a.I == 0 {
				a.I = 1
			}
		}
	}
	if in.StringList == nil {
		if err := json.Unmarshal([]byte(`["foo"]`), &in.StringList); err != nil {
			panic(err)
		}
	}
	if in.OtherSub.I == 0 {
		in.OtherSub.I = 1
	}
	if in.Map == nil {
		if err := json.Unmarshal([]byte(`{"foo": "bar"}`), &in.Map); err != nil {
			panic(err)
		}
	}
	for i_Map := range in.Map {
		if in.Map[i_Map] == nil {
			var ptrVar1 string = "apple"
			in.Map[i_Map] = &ptrVar1
		}
	}
	if in.StructMap == nil {
		if err := json.Unmarshal([]byte(`{"foo": {"S": "string", "I": 1}}`), &in.StructMap); err != nil {
			panic(err)
		}
	}
	if in.PtrStructMap == nil {
		if err := json.Unmarshal([]byte(`{"foo": {"S": "string", "I": 1}}`), &in.PtrStructMap); err != nil {
			panic(err)
		}
	}
	if in.AliasPtr == nil {
		var ptrVar1 string = "banana"
		in.AliasPtr = &ptrVar1
	}
	if in.SymbolReference == "" {
		in.SymbolReference = SomeDefault
	}
	if in.QualifiedSymbolReference == "" {
		in.QualifiedSymbolReference = v1.TerminationMessagePathDefault
	}
	if in.SameNamePackageSymbolReference1 == "" {
		in.SameNamePackageSymbolReference1 = external.AConstant
	}
	if in.SameNamePackageSymbolReference2 == "" {
		in.SameNamePackageSymbolReference2 = externalexternal.AnotherConstant
	}
}

func SetObjectDefaults_DefaultedWithFunction(in *DefaultedWithFunction) {
	SetDefaults_DefaultedWithFunction(in)
	if in.S1 == "" {
		in.S1 = "default_marker"
	}
	if in.S2 == "" {
		in.S2 = "default_marker"
	}
}
