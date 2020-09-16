/*
Copyright 2017 The Kubernetes Authors.

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

//go:generate apiserver-runtime-gen
package main

import (
	"github.com/pwittrock/apiserver-runtime-sample/pkg/apis/sample/v1alpha1"
	"github.com/pwittrock/apiserver-runtime-sample/pkg/generated/openapi"
	"github.com/pwittrock/apiserver-runtime/pkg/builder"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // register auth plugins
	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	err := builder.APIServer.
		WithOpenAPIDefinitions("sample", "v0.0.0", openapi.GetOpenAPIDefinitions).
		WithResource(&v1alpha1.Flunder{}). // namespaced resource
		WithResource(&v1alpha1.Fischer{}). // non-namespaced resource
		WithResource(&v1alpha1.Fortune{}). // resource with custom rest.Storage implementation
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
