// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemGretunnel`.
func GetSystemGretunnellist(ctx *pulumi.Context, args *GetSystemGretunnellistArgs, opts ...pulumi.InvokeOption) (*GetSystemGretunnellistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemGretunnellistResult
	err := ctx.Invoke("fortios:index/getSystemGretunnellist:getSystemGretunnellist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemGretunnellist.
type GetSystemGretunnellistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemGretunnellist.
type GetSystemGretunnellistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemGretunnel`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemGretunnellistOutput(ctx *pulumi.Context, args GetSystemGretunnellistOutputArgs, opts ...pulumi.InvokeOption) GetSystemGretunnellistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemGretunnellistResult, error) {
			args := v.(GetSystemGretunnellistArgs)
			r, err := GetSystemGretunnellist(ctx, &args, opts...)
			var s GetSystemGretunnellistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemGretunnellistResultOutput)
}

// A collection of arguments for invoking getSystemGretunnellist.
type GetSystemGretunnellistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemGretunnellistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemGretunnellistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemGretunnellist.
type GetSystemGretunnellistResultOutput struct{ *pulumi.OutputState }

func (GetSystemGretunnellistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemGretunnellistResult)(nil)).Elem()
}

func (o GetSystemGretunnellistResultOutput) ToGetSystemGretunnellistResultOutput() GetSystemGretunnellistResultOutput {
	return o
}

func (o GetSystemGretunnellistResultOutput) ToGetSystemGretunnellistResultOutputWithContext(ctx context.Context) GetSystemGretunnellistResultOutput {
	return o
}

func (o GetSystemGretunnellistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemGretunnellistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemGretunnellistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemGretunnellistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemGretunnel`.
func (o GetSystemGretunnellistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemGretunnellistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemGretunnellistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemGretunnellistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemGretunnellistResultOutput{})
}