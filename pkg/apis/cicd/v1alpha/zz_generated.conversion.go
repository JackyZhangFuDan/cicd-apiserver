//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha

import (
	cicd "cicd-apiserver/pkg/apis/cicd"
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*JenkinsServerInstance)(nil), (*cicd.JenkinsServerInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha_JenkinsServerInstance_To_cicd_JenkinsServerInstance(a.(*JenkinsServerInstance), b.(*cicd.JenkinsServerInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cicd.JenkinsServerInstance)(nil), (*JenkinsServerInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cicd_JenkinsServerInstance_To_v1alpha_JenkinsServerInstance(a.(*cicd.JenkinsServerInstance), b.(*JenkinsServerInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JenkinsService)(nil), (*cicd.JenkinsService)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha_JenkinsService_To_cicd_JenkinsService(a.(*JenkinsService), b.(*cicd.JenkinsService), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cicd.JenkinsService)(nil), (*JenkinsService)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cicd_JenkinsService_To_v1alpha_JenkinsService(a.(*cicd.JenkinsService), b.(*JenkinsService), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JenkinsServiceList)(nil), (*cicd.JenkinsServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha_JenkinsServiceList_To_cicd_JenkinsServiceList(a.(*JenkinsServiceList), b.(*cicd.JenkinsServiceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cicd.JenkinsServiceList)(nil), (*JenkinsServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cicd_JenkinsServiceList_To_v1alpha_JenkinsServiceList(a.(*cicd.JenkinsServiceList), b.(*JenkinsServiceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JenkinsServiceSpec)(nil), (*cicd.JenkinsServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha_JenkinsServiceSpec_To_cicd_JenkinsServiceSpec(a.(*JenkinsServiceSpec), b.(*cicd.JenkinsServiceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cicd.JenkinsServiceSpec)(nil), (*JenkinsServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cicd_JenkinsServiceSpec_To_v1alpha_JenkinsServiceSpec(a.(*cicd.JenkinsServiceSpec), b.(*JenkinsServiceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JenkinsServiceStatus)(nil), (*cicd.JenkinsServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha_JenkinsServiceStatus_To_cicd_JenkinsServiceStatus(a.(*JenkinsServiceStatus), b.(*cicd.JenkinsServiceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cicd.JenkinsServiceStatus)(nil), (*JenkinsServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cicd_JenkinsServiceStatus_To_v1alpha_JenkinsServiceStatus(a.(*cicd.JenkinsServiceStatus), b.(*JenkinsServiceStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha_JenkinsServerInstance_To_cicd_JenkinsServerInstance(in *JenkinsServerInstance, out *cicd.JenkinsServerInstance, s conversion.Scope) error {
	out.Cpu = in.Cpu
	out.Running = in.Running
	return nil
}

// Convert_v1alpha_JenkinsServerInstance_To_cicd_JenkinsServerInstance is an autogenerated conversion function.
func Convert_v1alpha_JenkinsServerInstance_To_cicd_JenkinsServerInstance(in *JenkinsServerInstance, out *cicd.JenkinsServerInstance, s conversion.Scope) error {
	return autoConvert_v1alpha_JenkinsServerInstance_To_cicd_JenkinsServerInstance(in, out, s)
}

func autoConvert_cicd_JenkinsServerInstance_To_v1alpha_JenkinsServerInstance(in *cicd.JenkinsServerInstance, out *JenkinsServerInstance, s conversion.Scope) error {
	out.Cpu = in.Cpu
	out.Running = in.Running
	return nil
}

// Convert_cicd_JenkinsServerInstance_To_v1alpha_JenkinsServerInstance is an autogenerated conversion function.
func Convert_cicd_JenkinsServerInstance_To_v1alpha_JenkinsServerInstance(in *cicd.JenkinsServerInstance, out *JenkinsServerInstance, s conversion.Scope) error {
	return autoConvert_cicd_JenkinsServerInstance_To_v1alpha_JenkinsServerInstance(in, out, s)
}

func autoConvert_v1alpha_JenkinsService_To_cicd_JenkinsService(in *JenkinsService, out *cicd.JenkinsService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha_JenkinsServiceSpec_To_cicd_JenkinsServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha_JenkinsServiceStatus_To_cicd_JenkinsServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha_JenkinsService_To_cicd_JenkinsService is an autogenerated conversion function.
func Convert_v1alpha_JenkinsService_To_cicd_JenkinsService(in *JenkinsService, out *cicd.JenkinsService, s conversion.Scope) error {
	return autoConvert_v1alpha_JenkinsService_To_cicd_JenkinsService(in, out, s)
}

func autoConvert_cicd_JenkinsService_To_v1alpha_JenkinsService(in *cicd.JenkinsService, out *JenkinsService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cicd_JenkinsServiceSpec_To_v1alpha_JenkinsServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cicd_JenkinsServiceStatus_To_v1alpha_JenkinsServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cicd_JenkinsService_To_v1alpha_JenkinsService is an autogenerated conversion function.
func Convert_cicd_JenkinsService_To_v1alpha_JenkinsService(in *cicd.JenkinsService, out *JenkinsService, s conversion.Scope) error {
	return autoConvert_cicd_JenkinsService_To_v1alpha_JenkinsService(in, out, s)
}

func autoConvert_v1alpha_JenkinsServiceList_To_cicd_JenkinsServiceList(in *JenkinsServiceList, out *cicd.JenkinsServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cicd.JenkinsService)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha_JenkinsServiceList_To_cicd_JenkinsServiceList is an autogenerated conversion function.
func Convert_v1alpha_JenkinsServiceList_To_cicd_JenkinsServiceList(in *JenkinsServiceList, out *cicd.JenkinsServiceList, s conversion.Scope) error {
	return autoConvert_v1alpha_JenkinsServiceList_To_cicd_JenkinsServiceList(in, out, s)
}

func autoConvert_cicd_JenkinsServiceList_To_v1alpha_JenkinsServiceList(in *cicd.JenkinsServiceList, out *JenkinsServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]JenkinsService)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cicd_JenkinsServiceList_To_v1alpha_JenkinsServiceList is an autogenerated conversion function.
func Convert_cicd_JenkinsServiceList_To_v1alpha_JenkinsServiceList(in *cicd.JenkinsServiceList, out *JenkinsServiceList, s conversion.Scope) error {
	return autoConvert_cicd_JenkinsServiceList_To_v1alpha_JenkinsServiceList(in, out, s)
}

func autoConvert_v1alpha_JenkinsServiceSpec_To_cicd_JenkinsServiceSpec(in *JenkinsServiceSpec, out *cicd.JenkinsServiceSpec, s conversion.Scope) error {
	out.InstanceAmount = in.InstanceAmount
	out.InstanceCpu = in.InstanceCpu
	return nil
}

// Convert_v1alpha_JenkinsServiceSpec_To_cicd_JenkinsServiceSpec is an autogenerated conversion function.
func Convert_v1alpha_JenkinsServiceSpec_To_cicd_JenkinsServiceSpec(in *JenkinsServiceSpec, out *cicd.JenkinsServiceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha_JenkinsServiceSpec_To_cicd_JenkinsServiceSpec(in, out, s)
}

func autoConvert_cicd_JenkinsServiceSpec_To_v1alpha_JenkinsServiceSpec(in *cicd.JenkinsServiceSpec, out *JenkinsServiceSpec, s conversion.Scope) error {
	out.InstanceAmount = in.InstanceAmount
	out.InstanceCpu = in.InstanceCpu
	return nil
}

// Convert_cicd_JenkinsServiceSpec_To_v1alpha_JenkinsServiceSpec is an autogenerated conversion function.
func Convert_cicd_JenkinsServiceSpec_To_v1alpha_JenkinsServiceSpec(in *cicd.JenkinsServiceSpec, out *JenkinsServiceSpec, s conversion.Scope) error {
	return autoConvert_cicd_JenkinsServiceSpec_To_v1alpha_JenkinsServiceSpec(in, out, s)
}

func autoConvert_v1alpha_JenkinsServiceStatus_To_cicd_JenkinsServiceStatus(in *JenkinsServiceStatus, out *cicd.JenkinsServiceStatus, s conversion.Scope) error {
	out.ApprovalStatus = in.ApprovalStatus
	out.Instances = *(*[]cicd.JenkinsServerInstance)(unsafe.Pointer(&in.Instances))
	return nil
}

// Convert_v1alpha_JenkinsServiceStatus_To_cicd_JenkinsServiceStatus is an autogenerated conversion function.
func Convert_v1alpha_JenkinsServiceStatus_To_cicd_JenkinsServiceStatus(in *JenkinsServiceStatus, out *cicd.JenkinsServiceStatus, s conversion.Scope) error {
	return autoConvert_v1alpha_JenkinsServiceStatus_To_cicd_JenkinsServiceStatus(in, out, s)
}

func autoConvert_cicd_JenkinsServiceStatus_To_v1alpha_JenkinsServiceStatus(in *cicd.JenkinsServiceStatus, out *JenkinsServiceStatus, s conversion.Scope) error {
	out.ApprovalStatus = in.ApprovalStatus
	out.Instances = *(*[]JenkinsServerInstance)(unsafe.Pointer(&in.Instances))
	return nil
}

// Convert_cicd_JenkinsServiceStatus_To_v1alpha_JenkinsServiceStatus is an autogenerated conversion function.
func Convert_cicd_JenkinsServiceStatus_To_v1alpha_JenkinsServiceStatus(in *cicd.JenkinsServiceStatus, out *JenkinsServiceStatus, s conversion.Scope) error {
	return autoConvert_cicd_JenkinsServiceStatus_To_v1alpha_JenkinsServiceStatus(in, out, s)
}
