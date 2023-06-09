// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemSessionhelper`.
func GetSystemSessionhelperlist(ctx *pulumi.Context, args *GetSystemSessionhelperlistArgs, opts ...pulumi.InvokeOption) (*GetSystemSessionhelperlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemSessionhelperlistResult
	err := ctx.Invoke("fortios:index/getSystemSessionhelperlist:getSystemSessionhelperlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemSessionhelperlist.
type GetSystemSessionhelperlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemSessionhelperlist.
type GetSystemSessionhelperlistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `SystemSessionhelper`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemSessionhelperlistOutput(ctx *pulumi.Context, args GetSystemSessionhelperlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemSessionhelperlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemSessionhelperlistResult, error) {
			args := v.(GetSystemSessionhelperlistArgs)
			r, err := GetSystemSessionhelperlist(ctx, &args, opts...)
			var s GetSystemSessionhelperlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemSessionhelperlistResultOutput)
}

// A collection of arguments for invoking getSystemSessionhelperlist.
type GetSystemSessionhelperlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemSessionhelperlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemSessionhelperlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemSessionhelperlist.
type GetSystemSessionhelperlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemSessionhelperlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemSessionhelperlistResult)(nil)).Elem()
}

func (o GetSystemSessionhelperlistResultOutput) ToGetSystemSessionhelperlistResultOutput() GetSystemSessionhelperlistResultOutput {
	return o
}

func (o GetSystemSessionhelperlistResultOutput) ToGetSystemSessionhelperlistResultOutputWithContext(ctx context.Context) GetSystemSessionhelperlistResultOutput {
	return o
}

func (o GetSystemSessionhelperlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemSessionhelperlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `SystemSessionhelper`.
func (o GetSystemSessionhelperlistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemSessionhelperlistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemSessionhelperlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemSessionhelperlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemSessionhelperlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemSessionhelperlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemSessionhelperlistResultOutput{})
}
