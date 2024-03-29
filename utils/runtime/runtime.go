// Copyright 2023 OnMetal authors
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

package runtime

import "github.com/onmetal/onmetal-api/utils/slices"

type DeepCopier[E any] interface {
	DeepCopy() E
}

type RefDeepCopier[E any] interface {
	*E
	DeepCopier[*E]
}

func DeepCopySlice[E DeepCopier[E], S ~[]E](slice S) S {
	return slices.Map(slice, func(e E) E {
		return e.DeepCopy()
	})
}

// DeepCopySliceRefs runs DeepCopy on the references of the elements of the slice and returns the created structs.
func DeepCopySliceRefs[E any, D RefDeepCopier[E], S ~[]E](slice S) []E {
	return slices.MapRef(slice, func(e *E) E {
		return *(D(e)).DeepCopy()
	})
}
