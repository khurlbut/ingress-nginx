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

package annotations

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.IngressNginxDescribe("Annotations - Log", func() {
	f := framework.NewDefaultFramework("log")

	BeforeEach(func() {
		err := f.NewEchoDeploymentWithReplicas(2)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
	})

	It("set access_log off", func() {
		host := "log.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/enable-access-log": "false",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.IngressController.Namespace, "http-svc", 80, &annotations)
		_, err := f.EnsureIngress(ing)

		Expect(err).NotTo(HaveOccurred())
		Expect(ing).NotTo(BeNil())

		err = f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring(`access_log off;`))
			})
		Expect(err).NotTo(HaveOccurred())
	})

	It("set rewrite_log on", func() {
		host := "log.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/enable-rewrite-log": "true",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.IngressController.Namespace, "http-svc", 80, &annotations)
		_, err := f.EnsureIngress(ing)

		Expect(err).NotTo(HaveOccurred())
		Expect(ing).NotTo(BeNil())

		err = f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring(`rewrite_log on;`))
			})
		Expect(err).NotTo(HaveOccurred())
	})
})
