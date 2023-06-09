// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallAddrgrp`.
func GetFirewallAddrgrplist(ctx *pulumi.Context, args *GetFirewallAddrgrplistArgs, opts ...pulumi.InvokeOption) (*GetFirewallAddrgrplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallAddrgrplistResult
	err := ctx.Invoke("fortios:index/getFirewallAddrgrplist:getFirewallAddrgrplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallAddrgrplist.
type GetFirewallAddrgrplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallAddrgrplist.
type GetFirewallAddrgrplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallAddrgrp`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallAddrgrplistOutput(ctx *pulumi.Context, args GetFirewallAddrgrplistOutputArgs, opts ...pulumi.InvokeOption) GetFirewallAddrgrplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallAddrgrplistResult, error) {
			args := v.(GetFirewallAddrgrplistArgs)
			r, err := GetFirewallAddrgrplist(ctx, &args, opts...)
			var s GetFirewallAddrgrplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallAddrgrplistResultOutput)
}

// A collection of arguments for invoking getFirewallAddrgrplist.
type GetFirewallAddrgrplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallAddrgrplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddrgrplistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallAddrgrplist.
type GetFirewallAddrgrplistResultOutput struct{ *pulumi.OutputState }

func (GetFirewallAddrgrplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddrgrplistResult)(nil)).Elem()
}

func (o GetFirewallAddrgrplistResultOutput) ToGetFirewallAddrgrplistResultOutput() GetFirewallAddrgrplistResultOutput {
	return o
}

func (o GetFirewallAddrgrplistResultOutput) ToGetFirewallAddrgrplistResultOutputWithContext(ctx context.Context) GetFirewallAddrgrplistResultOutput {
	return o
}

func (o GetFirewallAddrgrplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddrgrplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallAddrgrplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallAddrgrplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallAddrgrp`.
func (o GetFirewallAddrgrplistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallAddrgrplistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallAddrgrplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddrgrplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallAddrgrplistResultOutput{})
}
