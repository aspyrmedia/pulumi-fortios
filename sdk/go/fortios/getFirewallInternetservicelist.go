// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallInternetservice`.
func LookupFirewallInternetservicelist(ctx *pulumi.Context, args *LookupFirewallInternetservicelistArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetservicelistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetservicelistResult
	err := ctx.Invoke("fortios:index/getFirewallInternetservicelist:getFirewallInternetservicelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallInternetservicelist.
type LookupFirewallInternetservicelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallInternetservicelist.
type LookupFirewallInternetservicelistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `FirewallInternetservice`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallInternetservicelistOutput(ctx *pulumi.Context, args LookupFirewallInternetservicelistOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetservicelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetservicelistResult, error) {
			args := v.(LookupFirewallInternetservicelistArgs)
			r, err := LookupFirewallInternetservicelist(ctx, &args, opts...)
			var s LookupFirewallInternetservicelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallInternetservicelistResultOutput)
}

// A collection of arguments for invoking getFirewallInternetservicelist.
type LookupFirewallInternetservicelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetservicelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicelistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallInternetservicelist.
type LookupFirewallInternetservicelistResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetservicelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicelistResult)(nil)).Elem()
}

func (o LookupFirewallInternetservicelistResultOutput) ToLookupFirewallInternetservicelistResultOutput() LookupFirewallInternetservicelistResultOutput {
	return o
}

func (o LookupFirewallInternetservicelistResultOutput) ToLookupFirewallInternetservicelistResultOutputWithContext(ctx context.Context) LookupFirewallInternetservicelistResultOutput {
	return o
}

func (o LookupFirewallInternetservicelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `FirewallInternetservice`.
func (o LookupFirewallInternetservicelistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicelistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetservicelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicelistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetservicelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetservicelistResultOutput{})
}