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
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pwittrock/apiserver-runtime-sample/pkg/apis/sample/v1alpha1"
	"github.com/pwittrock/apiserver-runtime-sample/pkg/generated/openapi"
	"github.com/pwittrock/apiserver-runtime/pkg/builder"
	"github.com/pwittrock/apiserver-runtime/pkg/builder/resource"
	"github.com/pwittrock/apiserver-runtime/pkg/builder/rest"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // register auth plugins
	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	err := builder.APIServer.
		WithOpenAPIDefinitions("sample", "v0.0.0", openapi.GetOpenAPIDefinitions).
		WithResource(&v1alpha1.Flunder{}).
		WithResource(&v1alpha1.Fischer{}).
		WithResourceAndHandler(&v1alpha1.Fortune{}, newFortuneResourceHandler).

		// Start the apiserver
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}

func newFortuneResourceHandler(s *runtime.Scheme, g generic.RESTOptionsGetter) (rest.Storage, error) {
	return &fortuneResourceHandler{}, nil
}

var _ rest.Getter = &fortuneResourceHandler{}
var _ rest.Lister = &fortuneResourceHandler{}

type fortuneResourceHandler struct{ v1alpha1.Fortune }

func (f *fortuneResourceHandler) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*v1.Table, error) {
	c, ok := object.(resource.TableConverter)
	if !ok {
		return nil, fmt.Errorf("table printing not supoorted for %T", object)
	}
	return c.ConvertToTable(ctx, tableOptions)
}

func (f *fortuneResourceHandler) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	fl := &v1alpha1.FortuneList{}
	// return 5 random fortunes
	for i := 0; i < 5; i++ {
		obj, err := f.Get(ctx, "", &v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		fl.Items = append(fl.Items, *obj.(*v1alpha1.Fortune))
	}
	return fl, nil
}

func (f *fortuneResourceHandler) Get(ctx context.Context, name string, options *v1.GetOptions) (runtime.Object, error) {
	obj := &v1alpha1.Fortune{}
	var out []byte
	// fortune exits non-zero on success
	if name == "" {
		out, _ = exec.Command("/usr/games/fortune", "-s").Output()
	} else {
		/* #nosec */
		out, _ = exec.Command("/usr/games/fortune", "-s", "-m", name).Output()
		fortunes := strings.Split(string(out), "\n%\n")
		if len(fortunes) > 0 {
			out = []byte(fortunes[0])
		}
	}
	if len(strings.TrimSpace(string(out))) == 0 {
		return nil, errors.NewNotFound(v1alpha1.Fortune{}.GetGroupVersionResource().GroupResource(), name)
	}
	obj.Value = strings.TrimSpace(string(out))
	return obj, nil
}
