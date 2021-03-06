// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthTypes) DeepCopyInto(out *AuthTypes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthTypes.
func (in *AuthTypes) DeepCopy() *AuthTypes {
	if in == nil {
		return nil
	}
	out := new(AuthTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonStatus) DeepCopyInto(out *CommonStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonStatus.
func (in *CommonStatus) DeepCopy() *CommonStatus {
	if in == nil {
		return nil
	}
	out := new(CommonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericStatus) DeepCopyInto(out *GenericStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericStatus.
func (in *GenericStatus) DeepCopy() *GenericStatus {
	if in == nil {
		return nil
	}
	out := new(GenericStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keycloak) DeepCopyInto(out *Keycloak) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keycloak.
func (in *Keycloak) DeepCopy() *Keycloak {
	if in == nil {
		return nil
	}
	out := new(Keycloak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Keycloak) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClient) DeepCopyInto(out *KeycloakClient) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClient.
func (in *KeycloakClient) DeepCopy() *KeycloakClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakList) DeepCopyInto(out *KeycloakList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Keycloak, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakList.
func (in *KeycloakList) DeepCopy() *KeycloakList {
	if in == nil {
		return nil
	}
	out := new(KeycloakList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealm) DeepCopyInto(out *KeycloakRealm) {
	*out = *in
	if in.AuthTypes != nil {
		in, out := &in.AuthTypes, &out.AuthTypes
		*out = make([]AuthTypes, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]KeycloakUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Clients != nil {
		in, out := &in.Clients, &out.Clients
		*out = make([]KeycloakClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealm.
func (in *KeycloakRealm) DeepCopy() *KeycloakRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
	if in.Realms != nil {
		in, out := &in.Realms, &out.Realms
		*out = make([]KeycloakRealm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakStatus) DeepCopyInto(out *KeycloakStatus) {
	*out = *in
	out.GenericStatus = in.GenericStatus
	out.SharedConfig = in.SharedConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakStatus.
func (in *KeycloakStatus) DeepCopy() *KeycloakStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUser) DeepCopyInto(out *KeycloakUser) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUser.
func (in *KeycloakUser) DeepCopy() *KeycloakUser {
	if in == nil {
		return nil
	}
	out := new(KeycloakUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedService) DeepCopyInto(out *SharedService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedService.
func (in *SharedService) DeepCopy() *SharedService {
	if in == nil {
		return nil
	}
	out := new(SharedService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceAction) DeepCopyInto(out *SharedServiceAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceAction.
func (in *SharedServiceAction) DeepCopy() *SharedServiceAction {
	if in == nil {
		return nil
	}
	out := new(SharedServiceAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServiceAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceActionList) DeepCopyInto(out *SharedServiceActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedServiceAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceActionList.
func (in *SharedServiceActionList) DeepCopy() *SharedServiceActionList {
	if in == nil {
		return nil
	}
	out := new(SharedServiceActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServiceActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceActionSpec) DeepCopyInto(out *SharedServiceActionSpec) {
	*out = *in
	if in.ProvidedParams != nil {
		in, out := &in.ProvidedParams, &out.ProvidedParams
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceActionSpec.
func (in *SharedServiceActionSpec) DeepCopy() *SharedServiceActionSpec {
	if in == nil {
		return nil
	}
	out := new(SharedServiceActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceActionStatus) DeepCopyInto(out *SharedServiceActionStatus) {
	*out = *in
	out.CommonStatus = in.CommonStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceActionStatus.
func (in *SharedServiceActionStatus) DeepCopy() *SharedServiceActionStatus {
	if in == nil {
		return nil
	}
	out := new(SharedServiceActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceList) DeepCopyInto(out *SharedServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceList.
func (in *SharedServiceList) DeepCopy() *SharedServiceList {
	if in == nil {
		return nil
	}
	out := new(SharedServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServicePlan) DeepCopyInto(out *SharedServicePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServicePlan.
func (in *SharedServicePlan) DeepCopy() *SharedServicePlan {
	if in == nil {
		return nil
	}
	out := new(SharedServicePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServicePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServicePlanList) DeepCopyInto(out *SharedServicePlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedServicePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServicePlanList.
func (in *SharedServicePlanList) DeepCopy() *SharedServicePlanList {
	if in == nil {
		return nil
	}
	out := new(SharedServicePlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServicePlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServicePlanSpec) DeepCopyInto(out *SharedServicePlanSpec) {
	*out = *in
	if in.InstanceParams != nil {
		in, out := &in.InstanceParams, &out.InstanceParams
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.BindParams != nil {
		in, out := &in.BindParams, &out.BindParams
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServicePlanSpec.
func (in *SharedServicePlanSpec) DeepCopy() *SharedServicePlanSpec {
	if in == nil {
		return nil
	}
	out := new(SharedServicePlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServicePlanStatus) DeepCopyInto(out *SharedServicePlanStatus) {
	*out = *in
	out.CommonStatus = in.CommonStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServicePlanStatus.
func (in *SharedServicePlanStatus) DeepCopy() *SharedServicePlanStatus {
	if in == nil {
		return nil
	}
	out := new(SharedServicePlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceSlice) DeepCopyInto(out *SharedServiceSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceSlice.
func (in *SharedServiceSlice) DeepCopy() *SharedServiceSlice {
	if in == nil {
		return nil
	}
	out := new(SharedServiceSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServiceSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceSliceList) DeepCopyInto(out *SharedServiceSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedServiceSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceSliceList.
func (in *SharedServiceSliceList) DeepCopy() *SharedServiceSliceList {
	if in == nil {
		return nil
	}
	out := new(SharedServiceSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedServiceSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceSliceSpec) DeepCopyInto(out *SharedServiceSliceSpec) {
	*out = *in
	if in.ProvidedParams != nil {
		in, out := &in.ProvidedParams, &out.ProvidedParams
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceSliceSpec.
func (in *SharedServiceSliceSpec) DeepCopy() *SharedServiceSliceSpec {
	if in == nil {
		return nil
	}
	out := new(SharedServiceSliceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceSliceStatus) DeepCopyInto(out *SharedServiceSliceStatus) {
	*out = *in
	out.CommonStatus = in.CommonStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceSliceStatus.
func (in *SharedServiceSliceStatus) DeepCopy() *SharedServiceSliceStatus {
	if in == nil {
		return nil
	}
	out := new(SharedServiceSliceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceSpec) DeepCopyInto(out *SharedServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceSpec.
func (in *SharedServiceSpec) DeepCopy() *SharedServiceSpec {
	if in == nil {
		return nil
	}
	out := new(SharedServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedServiceStatus) DeepCopyInto(out *SharedServiceStatus) {
	*out = *in
	out.CommonStatus = in.CommonStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedServiceStatus.
func (in *SharedServiceStatus) DeepCopy() *SharedServiceStatus {
	if in == nil {
		return nil
	}
	out := new(SharedServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusSharedConfig) DeepCopyInto(out *StatusSharedConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusSharedConfig.
func (in *StatusSharedConfig) DeepCopy() *StatusSharedConfig {
	if in == nil {
		return nil
	}
	out := new(StatusSharedConfig)
	in.DeepCopyInto(out)
	return out
}
