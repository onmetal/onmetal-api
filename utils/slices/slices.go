// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slices

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

func Map[S ~[]E, E, F any](s S, f func(e E) F) []F {
	res := make([]F, len(s))
	for i, e := range s {
		res[i] = f(e)
	}
	return res
}

// MapRef maps the references of the values of the slice to a result type.
func MapRef[S ~[]E, E, F any](s S, f func(e *E) F) []F {
	res := make([]F, len(s))
	for i := range s {
		res[i] = f(&s[i])
	}
	return res
}

func ToMap[S ~[]E, E any, K comparable, V any](s S, f func(E) (K, V)) map[K]V {
	res := make(map[K]V)
	for _, e := range s {
		k, v := f(e)
		res[k] = v
	}
	return res
}

func ToMapByKey[S ~[]V, K comparable, V any](s S, f func(v V) K) map[K]V {
	return ToMap(s, func(v V) (K, V) {
		return f(v), v
	})
}

func FilterNot[S ~[]E, E comparable](s S, e E) []E {
	return FilterFunc(s, func(it E) bool {
		return e != it
	})
}

func Filter[S ~[]E, E comparable](s S, e E) []E {
	return FilterFunc(s, func(it E) bool {
		return e == it
	})
}

func FilterFunc[S ~[]E, E any](s S, f func(e E) bool) []E {
	var res []E
	for i := range s {
		if f(s[i]) {
			res = append(res, s[i])
		}
	}
	return res
}

func ToSetFunc[S ~[]V, K comparable, V any](s S, f func(v V) K) sets.Set[K] {
	res := sets.New[K]()
	for _, v := range s {
		res.Insert(f(v))
	}
	return res
}

func FindFunc[S ~[]V, V any](s S, f func(v V) bool) (V, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

func FindRefFunc[S ~[]V, V any](s S, f func(v V) bool) *V {
	for i, v := range s {
		if f(v) {
			return &s[i]
		}
	}
	return nil
}
