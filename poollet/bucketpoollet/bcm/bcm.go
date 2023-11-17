// Copyright 2022 IronCore authors
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

package bcm

import (
	"context"
	"errors"

	ori "github.com/ironcore-dev/ironcore/ori/apis/bucket/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	ErrNoMatchingBucketClass        = errors.New("no matching bucket class")
	ErrAmbiguousMatchingBucketClass = errors.New("ambiguous matching bucket classes")
)

type BucketClassMapper interface {
	manager.Runnable
	GetBucketClassFor(ctx context.Context, name string, capabilities *ori.BucketClassCapabilities) (*ori.BucketClass, error)
	WaitForSync(ctx context.Context) error
}
