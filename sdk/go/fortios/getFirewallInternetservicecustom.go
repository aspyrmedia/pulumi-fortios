// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicecustom
func LookupFirewallInternetservicecustom(ctx *pulumi.Context, args *LookupFirewallInternetservicecustomArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetservicecustomResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetservicecustomResult
	err := ctx.Invoke("fortios:index/getFirewallInternetservicecustom:getFirewallInternetservicecustom", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallInternetservicecustom.
type LookupFirewallInternetservicecustomArgs struct {
	// Specify the name of the desired firewall internetservicecustom.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallInternetservicecustom.
type LookupFirewallInternetservicecustomResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
	Entries []GetFirewallInternetservicecustomEntry `pulumi:"entries"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Select the destination address6 or address group object from available options.
	Name string `pulumi:"name"`
	// Reputation level of the custom Internet Service.
	Reputation int     `pulumi:"reputation"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

func LookupFirewallInternetservicecustomOutput(ctx *pulumi.Context, args LookupFirewallInternetservicecustomOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetservicecustomResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetservicecustomResult, error) {
			args := v.(LookupFirewallInternetservicecustomArgs)
			r, err := LookupFirewallInternetservicecustom(ctx, &args, opts...)
			var s LookupFirewallInternetservicecustomResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallInternetservicecustomResultOutput)
}

// A collection of arguments for invoking getFirewallInternetservicecustom.
type LookupFirewallInternetservicecustomOutputArgs struct {
	// Specify the name of the desired firewall internetservicecustom.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetservicecustomOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicecustomArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallInternetservicecustom.
type LookupFirewallInternetservicecustomResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetservicecustomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicecustomResult)(nil)).Elem()
}

func (o LookupFirewallInternetservicecustomResultOutput) ToLookupFirewallInternetservicecustomResultOutput() LookupFirewallInternetservicecustomResultOutput {
	return o
}

func (o LookupFirewallInternetservicecustomResultOutput) ToLookupFirewallInternetservicecustomResultOutputWithContext(ctx context.Context) LookupFirewallInternetservicecustomResultOutput {
	return o
}

// Comment.
func (o LookupFirewallInternetservicecustomResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicecustomResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
func (o LookupFirewallInternetservicecustomResultOutput) Entries() GetFirewallInternetservicecustomEntryArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicecustomResult) []GetFirewallInternetservicecustomEntry {
		return v.Entries
	}).(GetFirewallInternetservicecustomEntryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetservicecustomResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicecustomResult) string { return v.Id }).(pulumi.StringOutput)
}

// Select the destination address6 or address group object from available options.
func (o LookupFirewallInternetservicecustomResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicecustomResult) string { return v.Name }).(pulumi.StringOutput)
}

// Reputation level of the custom Internet Service.
func (o LookupFirewallInternetservicecustomResultOutput) Reputation() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicecustomResult) int { return v.Reputation }).(pulumi.IntOutput)
}

func (o LookupFirewallInternetservicecustomResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicecustomResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetservicecustomResultOutput{})
}