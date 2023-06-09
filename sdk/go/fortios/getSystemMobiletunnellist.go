// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemMobiletunnel`.
func GetSystemMobiletunnellist(ctx *pulumi.Context, args *GetSystemMobiletunnellistArgs, opts ...pulumi.InvokeOption) (*GetSystemMobiletunnellistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemMobiletunnellistResult
	err := ctx.Invoke("fortios:index/getSystemMobiletunnellist:getSystemMobiletunnellist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemMobiletunnellist.
type GetSystemMobiletunnellistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemMobiletunnellist.
type GetSystemMobiletunnellistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemMobiletunnel`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemMobiletunnellistOutput(ctx *pulumi.Context, args GetSystemMobiletunnellistOutputArgs, opts ...pulumi.InvokeOption) GetSystemMobiletunnellistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemMobiletunnellistResult, error) {
			args := v.(GetSystemMobiletunnellistArgs)
			r, err := GetSystemMobiletunnellist(ctx, &args, opts...)
			var s GetSystemMobiletunnellistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemMobiletunnellistResultOutput)
}

// A collection of arguments for invoking getSystemMobiletunnellist.
type GetSystemMobiletunnellistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemMobiletunnellistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemMobiletunnellistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemMobiletunnellist.
type GetSystemMobiletunnellistResultOutput struct{ *pulumi.OutputState }

func (GetSystemMobiletunnellistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemMobiletunnellistResult)(nil)).Elem()
}

func (o GetSystemMobiletunnellistResultOutput) ToGetSystemMobiletunnellistResultOutput() GetSystemMobiletunnellistResultOutput {
	return o
}

func (o GetSystemMobiletunnellistResultOutput) ToGetSystemMobiletunnellistResultOutputWithContext(ctx context.Context) GetSystemMobiletunnellistResultOutput {
	return o
}

func (o GetSystemMobiletunnellistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemMobiletunnellistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemMobiletunnellistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemMobiletunnellistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemMobiletunnel`.
func (o GetSystemMobiletunnellistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemMobiletunnellistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemMobiletunnellistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemMobiletunnellistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemMobiletunnellistResultOutput{})
}
