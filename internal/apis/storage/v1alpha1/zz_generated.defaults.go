//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (c) 2022 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1alpha1.Bucket{}, func(obj interface{}) { SetObjectDefaults_Bucket(obj.(*v1alpha1.Bucket)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.BucketList{}, func(obj interface{}) { SetObjectDefaults_BucketList(obj.(*v1alpha1.BucketList)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.Volume{}, func(obj interface{}) { SetObjectDefaults_Volume(obj.(*v1alpha1.Volume)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.VolumeClass{}, func(obj interface{}) { SetObjectDefaults_VolumeClass(obj.(*v1alpha1.VolumeClass)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.VolumeClassList{}, func(obj interface{}) { SetObjectDefaults_VolumeClassList(obj.(*v1alpha1.VolumeClassList)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.VolumeList{}, func(obj interface{}) { SetObjectDefaults_VolumeList(obj.(*v1alpha1.VolumeList)) })
	return nil
}

func SetObjectDefaults_Bucket(in *v1alpha1.Bucket) {
	SetDefaults_BucketStatus(&in.Status)
}

func SetObjectDefaults_BucketList(in *v1alpha1.BucketList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Bucket(a)
	}
}

func SetObjectDefaults_Volume(in *v1alpha1.Volume) {
	SetDefaults_VolumeStatus(&in.Status)
}

func SetObjectDefaults_VolumeClass(in *v1alpha1.VolumeClass) {
	SetDefaults_VolumeClass(in)
}

func SetObjectDefaults_VolumeClassList(in *v1alpha1.VolumeClassList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VolumeClass(a)
	}
}

func SetObjectDefaults_VolumeList(in *v1alpha1.VolumeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Volume(a)
	}
}
