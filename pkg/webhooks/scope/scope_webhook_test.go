/*
 * Copyright (c) 2021 by the OnMetal authors.
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

package scope

import (
	"context"
	api "github.com/onmetal/onmetal-api/apis/core/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Scope webhook", func() {

	const (
		scopeName        = "myscope"
		scopeDescription = "myaccount description"
		scopeRegion      = "myregion"
		scopeNameSpace   = "default"

		//timeout  = time.Second * 10
		//interval = time.Second * 1
	)

	var scope *api.Scope
	//var scopeLookUpKey types.NamespacedName

	Context("When creating a Scope", func() {
		It("Should accept a Scope creation", func() {
			ctx := context.Background()
			By("Creating a new Scope")
			scope = &api.Scope{
				TypeMeta: metav1.TypeMeta{
					Kind:       api.ScopeGK.Kind,
					APIVersion: api.ScopeGK.Group,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      scopeName,
					Namespace: scopeNameSpace,
				},
				Spec: api.ScopeSpec{
					Description: scopeDescription,
					Region:      scopeRegion,
				},
			}
			//scopeLookUpKey = types.NamespacedName{
			//	Name:      scopeName,
			//	Namespace: scopeNameSpace,
			//}
			Expect(k8sClient.Create(ctx, scope)).Should(Succeed())

			//By(fmt.Sprintf("Expecting created and State %s", v1alpha1.ScopeStateInitial))
			//Eventually(func() bool {
			//	s := &api.Scope{}
			//	if err := k8sClient.Get(context.Background(), scopeLookUpKey, s); err != nil {
			//		return false
			//	}
			//	if s.Status.State == api.ScopeStateInitial {
			//		return true
			//	}
			//	return false
			//}, timeout, interval).Should(BeTrue())
		})
	})

})