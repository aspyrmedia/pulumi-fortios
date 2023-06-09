// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicegroup
func LookupFirewallInternetservicegroup(ctx *pulumi.Context, args *LookupFirewallInternetservicegroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetservicegroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetservicegroupResult
	err := ctx.Invoke("fortios:index/getFirewallInternetservicegroup:getFirewallInternetservicegroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallInternetservicegroup.
type LookupFirewallInternetservicegroupArgs struct {
	// Specify the name of the desired firewall internetservicegroup.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallInternetservicegroup.
type LookupFirewallInternetservicegroupResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// How this service may be used (source, destination or both).
	Direction string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Internet Service group member. The structure of `member` block is documented below.
	Members []GetFirewallInternetservicegroupMember `pulumi:"members"`
	// Internet Service name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallInternetservicegroupOutput(ctx *pulumi.Context, args LookupFirewallInternetservicegroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetservicegroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetservicegroupResult, error) {
			args := v.(LookupFirewallInternetservicegroupArgs)
			r, err := LookupFirewallInternetservicegroup(ctx, &args, opts...)
			var s LookupFirewallInternetservicegroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallInternetservicegroupResultOutput)
}

// A collection of arguments for invoking getFirewallInternetservicegroup.
type LookupFirewallInternetservicegroupOutputArgs struct {
	// Specify the name of the desired firewall internetservicegroup.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetservicegroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicegroupArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallInternetservicegroup.
type LookupFirewallInternetservicegroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetservicegroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicegroupResult)(nil)).Elem()
}

func (o LookupFirewallInternetservicegroupResultOutput) ToLookupFirewallInternetservicegroupResultOutput() LookupFirewallInternetservicegroupResultOutput {
	return o
}

func (o LookupFirewallInternetservicegroupResultOutput) ToLookupFirewallInternetservicegroupResultOutputWithContext(ctx context.Context) LookupFirewallInternetservicegroupResultOutput {
	return o
}

// Comment.
func (o LookupFirewallInternetservicegroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicegroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// How this service may be used (source, destination or both).
func (o LookupFirewallInternetservicegroupResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicegroupResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetservicegroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicegroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Internet Service group member. The structure of `member` block is documented below.
func (o LookupFirewallInternetservicegroupResultOutput) Members() GetFirewallInternetservicegroupMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicegroupResult) []GetFirewallInternetservicegroupMember {
		return v.Members
	}).(GetFirewallInternetservicegroupMemberArrayOutput)
}

// Internet Service name.
func (o LookupFirewallInternetservicegroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicegroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetservicegroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicegroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetservicegroupResultOutput{})
}
