//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleset) DeepCopyInto(out *FirewallRuleset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleset.
func (in *FirewallRuleset) DeepCopy() *FirewallRuleset {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallRuleset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRulesetInitParameters) DeepCopyInto(out *FirewallRulesetInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]map[string]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]*string, len(*in))
				for key, val := range *in {
					var outVal *string
					if val == nil {
						(*out)[key] = nil
					} else {
						inVal := (*in)[key]
						in, out := &inVal, &outVal
						*out = new(string)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRulesetInitParameters.
func (in *FirewallRulesetInitParameters) DeepCopy() *FirewallRulesetInitParameters {
	if in == nil {
		return nil
	}
	out := new(FirewallRulesetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRulesetList) DeepCopyInto(out *FirewallRulesetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirewallRuleset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRulesetList.
func (in *FirewallRulesetList) DeepCopy() *FirewallRulesetList {
	if in == nil {
		return nil
	}
	out := new(FirewallRulesetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallRulesetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRulesetObservation) DeepCopyInto(out *FirewallRulesetObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]map[string]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]*string, len(*in))
				for key, val := range *in {
					var outVal *string
					if val == nil {
						(*out)[key] = nil
					} else {
						inVal := (*in)[key]
						in, out := &inVal, &outVal
						*out = new(string)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRulesetObservation.
func (in *FirewallRulesetObservation) DeepCopy() *FirewallRulesetObservation {
	if in == nil {
		return nil
	}
	out := new(FirewallRulesetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRulesetParameters) DeepCopyInto(out *FirewallRulesetParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]map[string]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]*string, len(*in))
				for key, val := range *in {
					var outVal *string
					if val == nil {
						(*out)[key] = nil
					} else {
						inVal := (*in)[key]
						in, out := &inVal, &outVal
						*out = new(string)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRulesetParameters.
func (in *FirewallRulesetParameters) DeepCopy() *FirewallRulesetParameters {
	if in == nil {
		return nil
	}
	out := new(FirewallRulesetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRulesetSpec) DeepCopyInto(out *FirewallRulesetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRulesetSpec.
func (in *FirewallRulesetSpec) DeepCopy() *FirewallRulesetSpec {
	if in == nil {
		return nil
	}
	out := new(FirewallRulesetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRulesetStatus) DeepCopyInto(out *FirewallRulesetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRulesetStatus.
func (in *FirewallRulesetStatus) DeepCopy() *FirewallRulesetStatus {
	if in == nil {
		return nil
	}
	out := new(FirewallRulesetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnel) DeepCopyInto(out *GRETunnel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnel.
func (in *GRETunnel) DeepCopy() *GRETunnel {
	if in == nil {
		return nil
	}
	out := new(GRETunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GRETunnel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnelInitParameters) DeepCopyInto(out *GRETunnelInitParameters) {
	*out = *in
	if in.CloudflareGreEndpoint != nil {
		in, out := &in.CloudflareGreEndpoint, &out.CloudflareGreEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomerGreEndpoint != nil {
		in, out := &in.CustomerGreEndpoint, &out.CustomerGreEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEnabled != nil {
		in, out := &in.HealthCheckEnabled, &out.HealthCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheckTarget != nil {
		in, out := &in.HealthCheckTarget, &out.HealthCheckTarget
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckType != nil {
		in, out := &in.HealthCheckType, &out.HealthCheckType
		*out = new(string)
		**out = **in
	}
	if in.InterfaceAddress != nil {
		in, out := &in.InterfaceAddress, &out.InterfaceAddress
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnelInitParameters.
func (in *GRETunnelInitParameters) DeepCopy() *GRETunnelInitParameters {
	if in == nil {
		return nil
	}
	out := new(GRETunnelInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnelList) DeepCopyInto(out *GRETunnelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GRETunnel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnelList.
func (in *GRETunnelList) DeepCopy() *GRETunnelList {
	if in == nil {
		return nil
	}
	out := new(GRETunnelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GRETunnelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnelObservation) DeepCopyInto(out *GRETunnelObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.CloudflareGreEndpoint != nil {
		in, out := &in.CloudflareGreEndpoint, &out.CloudflareGreEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomerGreEndpoint != nil {
		in, out := &in.CustomerGreEndpoint, &out.CustomerGreEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEnabled != nil {
		in, out := &in.HealthCheckEnabled, &out.HealthCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheckTarget != nil {
		in, out := &in.HealthCheckTarget, &out.HealthCheckTarget
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckType != nil {
		in, out := &in.HealthCheckType, &out.HealthCheckType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InterfaceAddress != nil {
		in, out := &in.InterfaceAddress, &out.InterfaceAddress
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnelObservation.
func (in *GRETunnelObservation) DeepCopy() *GRETunnelObservation {
	if in == nil {
		return nil
	}
	out := new(GRETunnelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnelParameters) DeepCopyInto(out *GRETunnelParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudflareGreEndpoint != nil {
		in, out := &in.CloudflareGreEndpoint, &out.CloudflareGreEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomerGreEndpoint != nil {
		in, out := &in.CustomerGreEndpoint, &out.CustomerGreEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEnabled != nil {
		in, out := &in.HealthCheckEnabled, &out.HealthCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheckTarget != nil {
		in, out := &in.HealthCheckTarget, &out.HealthCheckTarget
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckType != nil {
		in, out := &in.HealthCheckType, &out.HealthCheckType
		*out = new(string)
		**out = **in
	}
	if in.InterfaceAddress != nil {
		in, out := &in.InterfaceAddress, &out.InterfaceAddress
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnelParameters.
func (in *GRETunnelParameters) DeepCopy() *GRETunnelParameters {
	if in == nil {
		return nil
	}
	out := new(GRETunnelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnelSpec) DeepCopyInto(out *GRETunnelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnelSpec.
func (in *GRETunnelSpec) DeepCopy() *GRETunnelSpec {
	if in == nil {
		return nil
	}
	out := new(GRETunnelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRETunnelStatus) DeepCopyInto(out *GRETunnelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRETunnelStatus.
func (in *GRETunnelStatus) DeepCopy() *GRETunnelStatus {
	if in == nil {
		return nil
	}
	out := new(GRETunnelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnel) DeepCopyInto(out *IPsecTunnel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnel.
func (in *IPsecTunnel) DeepCopy() *IPsecTunnel {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPsecTunnel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnelInitParameters) DeepCopyInto(out *IPsecTunnelInitParameters) {
	*out = *in
	if in.AllowNullCipher != nil {
		in, out := &in.AllowNullCipher, &out.AllowNullCipher
		*out = new(bool)
		**out = **in
	}
	if in.CloudflareEndpoint != nil {
		in, out := &in.CloudflareEndpoint, &out.CloudflareEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomerEndpoint != nil {
		in, out := &in.CustomerEndpoint, &out.CustomerEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FqdnID != nil {
		in, out := &in.FqdnID, &out.FqdnID
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckDirection != nil {
		in, out := &in.HealthCheckDirection, &out.HealthCheckDirection
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEnabled != nil {
		in, out := &in.HealthCheckEnabled, &out.HealthCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheckRate != nil {
		in, out := &in.HealthCheckRate, &out.HealthCheckRate
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckTarget != nil {
		in, out := &in.HealthCheckTarget, &out.HealthCheckTarget
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckType != nil {
		in, out := &in.HealthCheckType, &out.HealthCheckType
		*out = new(string)
		**out = **in
	}
	if in.HexID != nil {
		in, out := &in.HexID, &out.HexID
		*out = new(string)
		**out = **in
	}
	if in.InterfaceAddress != nil {
		in, out := &in.InterfaceAddress, &out.InterfaceAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RemoteID != nil {
		in, out := &in.RemoteID, &out.RemoteID
		*out = new(string)
		**out = **in
	}
	if in.ReplayProtection != nil {
		in, out := &in.ReplayProtection, &out.ReplayProtection
		*out = new(bool)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnelInitParameters.
func (in *IPsecTunnelInitParameters) DeepCopy() *IPsecTunnelInitParameters {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnelInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnelList) DeepCopyInto(out *IPsecTunnelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPsecTunnel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnelList.
func (in *IPsecTunnelList) DeepCopy() *IPsecTunnelList {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPsecTunnelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnelObservation) DeepCopyInto(out *IPsecTunnelObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AllowNullCipher != nil {
		in, out := &in.AllowNullCipher, &out.AllowNullCipher
		*out = new(bool)
		**out = **in
	}
	if in.CloudflareEndpoint != nil {
		in, out := &in.CloudflareEndpoint, &out.CloudflareEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomerEndpoint != nil {
		in, out := &in.CustomerEndpoint, &out.CustomerEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FqdnID != nil {
		in, out := &in.FqdnID, &out.FqdnID
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckDirection != nil {
		in, out := &in.HealthCheckDirection, &out.HealthCheckDirection
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEnabled != nil {
		in, out := &in.HealthCheckEnabled, &out.HealthCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheckRate != nil {
		in, out := &in.HealthCheckRate, &out.HealthCheckRate
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckTarget != nil {
		in, out := &in.HealthCheckTarget, &out.HealthCheckTarget
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckType != nil {
		in, out := &in.HealthCheckType, &out.HealthCheckType
		*out = new(string)
		**out = **in
	}
	if in.HexID != nil {
		in, out := &in.HexID, &out.HexID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InterfaceAddress != nil {
		in, out := &in.InterfaceAddress, &out.InterfaceAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RemoteID != nil {
		in, out := &in.RemoteID, &out.RemoteID
		*out = new(string)
		**out = **in
	}
	if in.ReplayProtection != nil {
		in, out := &in.ReplayProtection, &out.ReplayProtection
		*out = new(bool)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnelObservation.
func (in *IPsecTunnelObservation) DeepCopy() *IPsecTunnelObservation {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnelParameters) DeepCopyInto(out *IPsecTunnelParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowNullCipher != nil {
		in, out := &in.AllowNullCipher, &out.AllowNullCipher
		*out = new(bool)
		**out = **in
	}
	if in.CloudflareEndpoint != nil {
		in, out := &in.CloudflareEndpoint, &out.CloudflareEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomerEndpoint != nil {
		in, out := &in.CustomerEndpoint, &out.CustomerEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FqdnID != nil {
		in, out := &in.FqdnID, &out.FqdnID
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckDirection != nil {
		in, out := &in.HealthCheckDirection, &out.HealthCheckDirection
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckEnabled != nil {
		in, out := &in.HealthCheckEnabled, &out.HealthCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HealthCheckRate != nil {
		in, out := &in.HealthCheckRate, &out.HealthCheckRate
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckTarget != nil {
		in, out := &in.HealthCheckTarget, &out.HealthCheckTarget
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckType != nil {
		in, out := &in.HealthCheckType, &out.HealthCheckType
		*out = new(string)
		**out = **in
	}
	if in.HexID != nil {
		in, out := &in.HexID, &out.HexID
		*out = new(string)
		**out = **in
	}
	if in.InterfaceAddress != nil {
		in, out := &in.InterfaceAddress, &out.InterfaceAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PskSecretRef != nil {
		in, out := &in.PskSecretRef, &out.PskSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.RemoteID != nil {
		in, out := &in.RemoteID, &out.RemoteID
		*out = new(string)
		**out = **in
	}
	if in.ReplayProtection != nil {
		in, out := &in.ReplayProtection, &out.ReplayProtection
		*out = new(bool)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnelParameters.
func (in *IPsecTunnelParameters) DeepCopy() *IPsecTunnelParameters {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnelSpec) DeepCopyInto(out *IPsecTunnelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnelSpec.
func (in *IPsecTunnelSpec) DeepCopy() *IPsecTunnelSpec {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPsecTunnelStatus) DeepCopyInto(out *IPsecTunnelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPsecTunnelStatus.
func (in *IPsecTunnelStatus) DeepCopy() *IPsecTunnelStatus {
	if in == nil {
		return nil
	}
	out := new(IPsecTunnelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRoute) DeepCopyInto(out *StaticRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRoute.
func (in *StaticRoute) DeepCopy() *StaticRoute {
	if in == nil {
		return nil
	}
	out := new(StaticRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StaticRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRouteInitParameters) DeepCopyInto(out *StaticRouteInitParameters) {
	*out = *in
	if in.ColoNames != nil {
		in, out := &in.ColoNames, &out.ColoNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ColoRegions != nil {
		in, out := &in.ColoRegions, &out.ColoRegions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Nexthop != nil {
		in, out := &in.Nexthop, &out.Nexthop
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRouteInitParameters.
func (in *StaticRouteInitParameters) DeepCopy() *StaticRouteInitParameters {
	if in == nil {
		return nil
	}
	out := new(StaticRouteInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRouteList) DeepCopyInto(out *StaticRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StaticRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRouteList.
func (in *StaticRouteList) DeepCopy() *StaticRouteList {
	if in == nil {
		return nil
	}
	out := new(StaticRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StaticRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRouteObservation) DeepCopyInto(out *StaticRouteObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.ColoNames != nil {
		in, out := &in.ColoNames, &out.ColoNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ColoRegions != nil {
		in, out := &in.ColoRegions, &out.ColoRegions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Nexthop != nil {
		in, out := &in.Nexthop, &out.Nexthop
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRouteObservation.
func (in *StaticRouteObservation) DeepCopy() *StaticRouteObservation {
	if in == nil {
		return nil
	}
	out := new(StaticRouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRouteParameters) DeepCopyInto(out *StaticRouteParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ColoNames != nil {
		in, out := &in.ColoNames, &out.ColoNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ColoRegions != nil {
		in, out := &in.ColoRegions, &out.ColoRegions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Nexthop != nil {
		in, out := &in.Nexthop, &out.Nexthop
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRouteParameters.
func (in *StaticRouteParameters) DeepCopy() *StaticRouteParameters {
	if in == nil {
		return nil
	}
	out := new(StaticRouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRouteSpec) DeepCopyInto(out *StaticRouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRouteSpec.
func (in *StaticRouteSpec) DeepCopy() *StaticRouteSpec {
	if in == nil {
		return nil
	}
	out := new(StaticRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRouteStatus) DeepCopyInto(out *StaticRouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRouteStatus.
func (in *StaticRouteStatus) DeepCopy() *StaticRouteStatus {
	if in == nil {
		return nil
	}
	out := new(StaticRouteStatus)
	in.DeepCopyInto(out)
	return out
}
