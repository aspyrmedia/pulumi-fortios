// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterMulticastflow`.
func GetRouterMulticastflowlist(ctx *pulumi.Context, args *GetRouterMulticastflowlistArgs, opts ...pulumi.InvokeOption) (*GetRouterMulticastflowlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRouterMulticastflowlistResult
	err := ctx.Invoke("fortios:index/getRouterMulticastflowlist:getRouterMulticastflowlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterMulticastflowlist.
type GetRouterMulticastflowlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterMulticastflowlist.
type GetRouterMulticastflowlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterMulticastflow`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterMulticastflowlistOutput(ctx *pulumi.Context, args GetRouterMulticastflowlistOutputArgs, opts ...pulumi.InvokeOption) GetRouterMulticastflowlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterMulticastflowlistResult, error) {
			args := v.(GetRouterMulticastflowlistArgs)
			r, err := GetRouterMulticastflowlist(ctx, &args, opts...)
			var s GetRouterMulticastflowlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterMulticastflowlistResultOutput)
}

// A collection of arguments for invoking getRouterMulticastflowlist.
type GetRouterMulticastflowlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterMulticastflowlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterMulticastflowlistArgs)(nil)).Elem()
}

// A collection of values returned by getRouterMulticastflowlist.
type GetRouterMulticastflowlistResultOutput struct{ *pulumi.OutputState }

func (GetRouterMulticastflowlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterMulticastflowlistResult)(nil)).Elem()
}

func (o GetRouterMulticastflowlistResultOutput) ToGetRouterMulticastflowlistResultOutput() GetRouterMulticastflowlistResultOutput {
	return o
}

func (o GetRouterMulticastflowlistResultOutput) ToGetRouterMulticastflowlistResultOutputWithContext(ctx context.Context) GetRouterMulticastflowlistResultOutput {
	return o
}

func (o GetRouterMulticastflowlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterMulticastflowlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterMulticastflowlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterMulticastflowlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `RouterMulticastflow`.
func (o GetRouterMulticastflowlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterMulticastflowlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterMulticastflowlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterMulticastflowlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterMulticastflowlistResultOutput{})
}