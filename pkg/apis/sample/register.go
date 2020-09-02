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

package sample

import (
	"github.com/pwittrock/apiserver-runtime-sample/pkg/apis/sample/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Fischer is the internal type used for Fischer
type Fischer = v1alpha1.Fischer

// FischerList is the internal type used for FischerList
type FischerList = v1alpha1.FischerList

// Flunder is the internal type used for Flunder
type Flunder = v1alpha1.Flunder

// FlunderList is the internal type used for FlunderList
type FlunderList = v1alpha1.FlunderList

var (
	// SchemeGroupVersion configures the scheme for internal representations of types
	// the v1alpha1 version is used as the internal version
	SchemeGroupVersion = schema.GroupVersion{
		Group:   v1alpha1.SchemeGroupVersion.Group,
		Version: runtime.APIVersionInternal,
	}

	// SchemeBuilder collects functions that add things to a scheme. It's to allow
	// code to compile without explicitly referencing generated types. You should
	// declare one in each package that will have generated deep copy or conversion
	// functions.
	SchemeBuilder = runtime.NewSchemeBuilder(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(SchemeGroupVersion, &Fischer{}, &FischerList{}, &Flunder{}, &FlunderList{})
		return nil
	})

	// AddToScheme applies all the stored functions to the scheme. A non-nil error
	// indicates that one function failed and the attempt was abandoned.
	AddToScheme = SchemeBuilder.AddToScheme
)

// Install registers all types in the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) error {
	if err := AddToScheme(scheme); err != nil {
		return err
	}
	if err := v1alpha1.AddToScheme(scheme); err != nil {
		return err
	}
	return scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion)
}
