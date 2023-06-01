// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemArptable`.
func GetSystemArptablelist(ctx *pulumi.Context, args *GetSystemArptablelistArgs, opts ...pulumi.InvokeOption) (*GetSystemArptablelistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemArptablelistResult
	err := ctx.Invoke("fortios:index/getSystemArptablelist:getSystemArptablelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemArptablelist.
type GetSystemArptablelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemArptablelist.
type GetSystemArptablelistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `SystemArptable`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemArptablelistOutput(ctx *pulumi.Context, args GetSystemArptablelistOutputArgs, opts ...pulumi.InvokeOption) GetSystemArptablelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemArptablelistResult, error) {
			args := v.(GetSystemArptablelistArgs)
			r, err := GetSystemArptablelist(ctx, &args, opts...)
			var s GetSystemArptablelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemArptablelistResultOutput)
}

// A collection of arguments for invoking getSystemArptablelist.
type GetSystemArptablelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemArptablelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemArptablelistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemArptablelist.
type GetSystemArptablelistResultOutput struct{ *pulumi.OutputState }

func (GetSystemArptablelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemArptablelistResult)(nil)).Elem()
}

func (o GetSystemArptablelistResultOutput) ToGetSystemArptablelistResultOutput() GetSystemArptablelistResultOutput {
	return o
}

func (o GetSystemArptablelistResultOutput) ToGetSystemArptablelistResultOutputWithContext(ctx context.Context) GetSystemArptablelistResultOutput {
	return o
}

func (o GetSystemArptablelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemArptablelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `SystemArptable`.
func (o GetSystemArptablelistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemArptablelistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemArptablelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemArptablelistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemArptablelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemArptablelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemArptablelistResultOutput{})
}