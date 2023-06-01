// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall addrgrp
func LookupFirewallAddrgrp(ctx *pulumi.Context, args *LookupFirewallAddrgrpArgs, opts ...pulumi.InvokeOption) (*LookupFirewallAddrgrpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallAddrgrpResult
	err := ctx.Invoke("fortios:index/getFirewallAddrgrp:getFirewallAddrgrp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallAddrgrp.
type LookupFirewallAddrgrpArgs struct {
	// Specify the name of the desired firewall addrgrp.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallAddrgrp.
type LookupFirewallAddrgrpResult struct {
	// Enable/disable use of this group in the static route configuration.
	AllowRouting string `pulumi:"allowRouting"`
	// Tag category.
	Category string `pulumi:"category"`
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// Enable/disable address exclusion.
	Exclude string `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []GetFirewallAddrgrpExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []GetFirewallAddrgrpMember `pulumi:"members"`
	// Tag name.
	Name string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []GetFirewallAddrgrpTagging `pulumi:"taggings"`
	// Address group type.
	Type string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupFirewallAddrgrpOutput(ctx *pulumi.Context, args LookupFirewallAddrgrpOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallAddrgrpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallAddrgrpResult, error) {
			args := v.(LookupFirewallAddrgrpArgs)
			r, err := LookupFirewallAddrgrp(ctx, &args, opts...)
			var s LookupFirewallAddrgrpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallAddrgrpResultOutput)
}

// A collection of arguments for invoking getFirewallAddrgrp.
type LookupFirewallAddrgrpOutputArgs struct {
	// Specify the name of the desired firewall addrgrp.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallAddrgrpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallAddrgrpArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallAddrgrp.
type LookupFirewallAddrgrpResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallAddrgrpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallAddrgrpResult)(nil)).Elem()
}

func (o LookupFirewallAddrgrpResultOutput) ToLookupFirewallAddrgrpResultOutput() LookupFirewallAddrgrpResultOutput {
	return o
}

func (o LookupFirewallAddrgrpResultOutput) ToLookupFirewallAddrgrpResultOutputWithContext(ctx context.Context) LookupFirewallAddrgrpResultOutput {
	return o
}

// Enable/disable use of this group in the static route configuration.
func (o LookupFirewallAddrgrpResultOutput) AllowRouting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.AllowRouting }).(pulumi.StringOutput)
}

// Tag category.
func (o LookupFirewallAddrgrpResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Category }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o LookupFirewallAddrgrpResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupFirewallAddrgrpResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Enable/disable address exclusion.
func (o LookupFirewallAddrgrpResultOutput) Exclude() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Exclude }).(pulumi.StringOutput)
}

// Address exclusion member. The structure of `excludeMember` block is documented below.
func (o LookupFirewallAddrgrpResultOutput) ExcludeMembers() GetFirewallAddrgrpExcludeMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) []GetFirewallAddrgrpExcludeMember { return v.ExcludeMembers }).(GetFirewallAddrgrpExcludeMemberArrayOutput)
}

// Security Fabric global object setting.
func (o LookupFirewallAddrgrpResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallAddrgrpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Address objects contained within the group. The structure of `member` block is documented below.
func (o LookupFirewallAddrgrpResultOutput) Members() GetFirewallAddrgrpMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) []GetFirewallAddrgrpMember { return v.Members }).(GetFirewallAddrgrpMemberArrayOutput)
}

// Tag name.
func (o LookupFirewallAddrgrpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o LookupFirewallAddrgrpResultOutput) Taggings() GetFirewallAddrgrpTaggingArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) []GetFirewallAddrgrpTagging { return v.Taggings }).(GetFirewallAddrgrpTaggingArrayOutput)
}

// Address group type.
func (o LookupFirewallAddrgrpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupFirewallAddrgrpResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallAddrgrpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable address visibility in the GUI.
func (o LookupFirewallAddrgrpResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrpResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallAddrgrpResultOutput{})
}