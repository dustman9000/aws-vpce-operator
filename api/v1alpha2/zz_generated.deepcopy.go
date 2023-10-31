//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociatedVpc) DeepCopyInto(out *AssociatedVpc) {
	*out = *in
	if in.CredentialsSecretRef != nil {
		in, out := &in.CredentialsSecretRef, &out.CredentialsSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociatedVpc.
func (in *AssociatedVpc) DeepCopy() *AssociatedVpc {
	if in == nil {
		return nil
	}
	out := new(AssociatedVpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsEndpointSelector) DeepCopyInto(out *AwsEndpointSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsEndpointSelector.
func (in *AwsEndpointSelector) DeepCopy() *AwsEndpointSelector {
	if in == nil {
		return nil
	}
	out := new(AwsEndpointSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDns) DeepCopyInto(out *CustomDns) {
	*out = *in
	in.Route53PrivateHostedZone.DeepCopyInto(&out.Route53PrivateHostedZone)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDns.
func (in *CustomDns) DeepCopy() *CustomDns {
	if in == nil {
		return nil
	}
	out := new(CustomDns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnsSelector) DeepCopyInto(out *DnsSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnsSelector.
func (in *DnsSelector) DeepCopy() *DnsSelector {
	if in == nil {
		return nil
	}
	out := new(DnsSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainName) DeepCopyInto(out *DomainName) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(DomainNameSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainName.
func (in *DomainName) DeepCopy() *DomainName {
	if in == nil {
		return nil
	}
	out := new(DomainName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNameSource) DeepCopyInto(out *DomainNameSource) {
	*out = *in
	if in.DnsRef != nil {
		in, out := &in.DnsRef, &out.DnsRef
		*out = new(DnsSelector)
		**out = **in
	}
	if in.HostedControlPlaneRef != nil {
		in, out := &in.HostedControlPlaneRef, &out.HostedControlPlaneRef
		*out = new(HostedControlPlaneSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNameSource.
func (in *DomainNameSource) DeepCopy() *DomainNameSource {
	if in == nil {
		return nil
	}
	out := new(DomainNameSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNameService) DeepCopyInto(out *ExternalNameService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNameService.
func (in *ExternalNameService) DeepCopy() *ExternalNameService {
	if in == nil {
		return nil
	}
	out := new(ExternalNameService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedControlPlaneSelector) DeepCopyInto(out *HostedControlPlaneSelector) {
	*out = *in
	if in.NamespaceFieldRef != nil {
		in, out := &in.NamespaceFieldRef, &out.NamespaceFieldRef
		*out = new(ObjectFieldSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedControlPlaneSelector.
func (in *HostedControlPlaneSelector) DeepCopy() *HostedControlPlaneSelector {
	if in == nil {
		return nil
	}
	out := new(HostedControlPlaneSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectFieldSelector) DeepCopyInto(out *ObjectFieldSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectFieldSelector.
func (in *ObjectFieldSelector) DeepCopy() *ObjectFieldSelector {
	if in == nil {
		return nil
	}
	out := new(ObjectFieldSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HostedZoneRecord) DeepCopyInto(out *Route53HostedZoneRecord) {
	*out = *in
	out.ExternalNameService = in.ExternalNameService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HostedZoneRecord.
func (in *Route53HostedZoneRecord) DeepCopy() *Route53HostedZoneRecord {
	if in == nil {
		return nil
	}
	out := new(Route53HostedZoneRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53PrivateHostedZone) DeepCopyInto(out *Route53PrivateHostedZone) {
	*out = *in
	if in.AssociatedVpcs != nil {
		in, out := &in.AssociatedVpcs, &out.AssociatedVpcs
		*out = make([]AssociatedVpc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DomainNameRef != nil {
		in, out := &in.DomainNameRef, &out.DomainNameRef
		*out = new(DomainName)
		(*in).DeepCopyInto(*out)
	}
	out.Record = in.Record
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53PrivateHostedZone.
func (in *Route53PrivateHostedZone) DeepCopy() *Route53PrivateHostedZone {
	if in == nil {
		return nil
	}
	out := new(Route53PrivateHostedZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
	if in.EgressRules != nil {
		in, out := &in.EgressRules, &out.EgressRules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRule) DeepCopyInto(out *SecurityGroupRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRule.
func (in *SecurityGroupRule) DeepCopy() *SecurityGroupRule {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceName) DeepCopyInto(out *ServiceName) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ServiceNameSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceName.
func (in *ServiceName) DeepCopy() *ServiceName {
	if in == nil {
		return nil
	}
	out := new(ServiceName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceNameSource) DeepCopyInto(out *ServiceNameSource) {
	*out = *in
	if in.AwsEndpointServiceRef != nil {
		in, out := &in.AwsEndpointServiceRef, &out.AwsEndpointServiceRef
		*out = new(AwsEndpointSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceNameSource.
func (in *ServiceNameSource) DeepCopy() *ServiceNameSource {
	if in == nil {
		return nil
	}
	out := new(ServiceNameSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vpc) DeepCopyInto(out *Vpc) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ids != nil {
		in, out := &in.Ids, &out.Ids
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.SubnetTags != nil {
		in, out := &in.SubnetTags, &out.SubnetTags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpc.
func (in *Vpc) DeepCopy() *Vpc {
	if in == nil {
		return nil
	}
	out := new(Vpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpoint) DeepCopyInto(out *VpcEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpoint.
func (in *VpcEndpoint) DeepCopy() *VpcEndpoint {
	if in == nil {
		return nil
	}
	out := new(VpcEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointList) DeepCopyInto(out *VpcEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointList.
func (in *VpcEndpointList) DeepCopy() *VpcEndpointList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSpec) DeepCopyInto(out *VpcEndpointSpec) {
	*out = *in
	if in.ServiceNameRef != nil {
		in, out := &in.ServiceNameRef, &out.ServiceNameRef
		*out = new(ServiceName)
		(*in).DeepCopyInto(*out)
	}
	in.SecurityGroup.DeepCopyInto(&out.SecurityGroup)
	if in.AWSCredentialOverrideRef != nil {
		in, out := &in.AWSCredentialOverrideRef, &out.AWSCredentialOverrideRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	in.Vpc.DeepCopyInto(&out.Vpc)
	in.CustomDns.DeepCopyInto(&out.CustomDns)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSpec.
func (in *VpcEndpointSpec) DeepCopy() *VpcEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointStatus) DeepCopyInto(out *VpcEndpointStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointStatus.
func (in *VpcEndpointStatus) DeepCopy() *VpcEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointTemplate) DeepCopyInto(out *VpcEndpointTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointTemplate.
func (in *VpcEndpointTemplate) DeepCopy() *VpcEndpointTemplate {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointTemplateList) DeepCopyInto(out *VpcEndpointTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpointTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointTemplateList.
func (in *VpcEndpointTemplateList) DeepCopy() *VpcEndpointTemplateList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointTemplateSpec) DeepCopyInto(out *VpcEndpointTemplateSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointTemplateSpec.
func (in *VpcEndpointTemplateSpec) DeepCopy() *VpcEndpointTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointTemplateStatus) DeepCopyInto(out *VpcEndpointTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointTemplateStatus.
func (in *VpcEndpointTemplateStatus) DeepCopy() *VpcEndpointTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpceTemplateSpec) DeepCopyInto(out *VpceTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpceTemplateSpec.
func (in *VpceTemplateSpec) DeepCopy() *VpceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VpceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
