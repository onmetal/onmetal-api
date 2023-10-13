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

package storage

import (
	"context"
	"fmt"

	corev1beta1 "github.com/onmetal/onmetal-api/api/core/v1beta1"
	storagev1beta1 "github.com/onmetal/onmetal-api/api/storage/v1beta1"
	"github.com/onmetal/onmetal-api/internal/apis/storage"
	internalstoragev1beta1 "github.com/onmetal/onmetal-api/internal/apis/storage/v1beta1"
	"github.com/onmetal/onmetal-api/internal/quota/evaluator/generic"
	"github.com/onmetal/onmetal-api/utils/quota"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	bucketResource          = storagev1beta1.Resource("buckets")
	bucketCountResourceName = corev1beta1.ObjectCountQuotaResourceNameFor(bucketResource)

	BucketResourceNames = sets.New(
		bucketCountResourceName,
		corev1beta1.ResourceRequestsStorage,
	)
)

type bucketEvaluator struct {
	capabilities generic.CapabilitiesReader
}

func NewBucketEvaluator(capabilities generic.CapabilitiesReader) quota.Evaluator {
	return &bucketEvaluator{
		capabilities: capabilities,
	}
}

func (m *bucketEvaluator) Type() client.Object {
	return &storagev1beta1.Bucket{}
}

func (m *bucketEvaluator) MatchesResourceName(name corev1beta1.ResourceName) bool {
	return BucketResourceNames.Has(name)
}

func (m *bucketEvaluator) MatchesResourceScopeSelectorRequirement(item client.Object, req corev1beta1.ResourceScopeSelectorRequirement) (bool, error) {
	bucket := item.(*storagev1beta1.Bucket)

	switch req.ScopeName {
	case corev1beta1.ResourceScopeBucketClass:
		return bucketMatchesBucketClassScope(bucket, req.Operator, req.Values), nil
	default:
		return false, nil
	}
}

func bucketMatchesBucketClassScope(bucket *storagev1beta1.Bucket, op corev1beta1.ResourceScopeSelectorOperator, values []string) bool {
	bucketClassRef := bucket.Spec.BucketClassRef

	switch op {
	case corev1beta1.ResourceScopeSelectorOperatorExists:
		return bucketClassRef != nil
	case corev1beta1.ResourceScopeSelectorOperatorDoesNotExist:
		return bucketClassRef == nil
	case corev1beta1.ResourceScopeSelectorOperatorIn:
		return slices.Contains(values, bucketClassRef.Name)
	case corev1beta1.ResourceScopeSelectorOperatorNotIn:
		if bucketClassRef == nil {
			return false
		}
		return !slices.Contains(values, bucketClassRef.Name)
	default:
		return false
	}
}

func toExternalBucketOrError(obj client.Object) (*storagev1beta1.Bucket, error) {
	switch t := obj.(type) {
	case *storagev1beta1.Bucket:
		return t, nil
	case *storage.Bucket:
		bucket := &storagev1beta1.Bucket{}
		if err := internalstoragev1beta1.Convert_storage_Bucket_To_v1beta1_Bucket(t, bucket, nil); err != nil {
			return nil, err
		}
		return bucket, nil
	default:
		return nil, fmt.Errorf("expect *storage.Bucket or *storagev1beta1.Bucket but got %v", t)
	}
}

func (m *bucketEvaluator) Usage(ctx context.Context, item client.Object) (corev1beta1.ResourceList, error) {
	_, err := toExternalBucketOrError(item)
	if err != nil {
		return nil, err
	}

	return corev1beta1.ResourceList{
		// TODO: return more detailed usage
		bucketCountResourceName: resource.MustParse("1"),
	}, nil
}
