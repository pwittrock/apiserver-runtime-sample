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

package main

import (
	// register auth mechanisms
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/pwittrock/apiserver-runtime-sample/pkg/apis/sample"
	"github.com/pwittrock/apiserver-runtime-sample/pkg/apis/sample/v1alpha1"
	"github.com/pwittrock/apiserver-runtime-sample/pkg/generated/openapi"
	"github.com/pwittrock/apiserver-runtime/pkg/builder"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	err := builder.APIServer.
		WithSchemeInstallers(sample.Install).
		WithOpenAPIDefinitions("sample", "v0.0.0", openapi.GetOpenAPIDefinitions).
		WithResource(&v1alpha1.Flunder{}).
		WithResource(&v1alpha1.Fischer{}).

		// Start the apiserver
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
