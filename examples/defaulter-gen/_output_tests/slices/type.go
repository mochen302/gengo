/*
Copyright 2018 The Kubernetes Authors.

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

package slices

import (
	"github.com/mochen302/gengox/examples/defaulter-gen/_output_tests/empty"
)

// Only test
type Ttest struct {
	empty.TypeMeta
	BoolField *bool
}

type TtestList struct {
	empty.TypeMeta
	Items []Ttest
}

type TtestPointerList struct {
	empty.TypeMeta
	Items []*Ttest
}
