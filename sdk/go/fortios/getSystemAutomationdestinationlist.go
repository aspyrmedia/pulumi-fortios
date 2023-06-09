// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemAutomationdestination`.
func GetSystemAutomationdestinationlist(ctx *pulumi.Context, args *GetSystemAutomationdestinationlistArgs, opts ...pulumi.InvokeOption) (*GetSystemAutomationdestinationlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemAutomationdestinationlistResult
	err := ctx.Invoke("fortios:index/getSystemAutomationdestinationlist:getSystemAutomationdestinationlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAutomationdestinationlist.
type GetSystemAutomationdestinationlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAutomationdestinationlist.
type GetSystemAutomationdestinationlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemAutomationdestination`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAutomationdestinationlistOutput(ctx *pulumi.Context, args GetSystemAutomationdestinationlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemAutomationdestinationlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAutomationdestinationlistResult, error) {
			args := v.(GetSystemAutomationdestinationlistArgs)
			r, err := GetSystemAutomationdestinationlist(ctx, &args, opts...)
			var s GetSystemAutomationdestinationlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAutomationdestinationlistResultOutput)
}

// A collection of arguments for invoking getSystemAutomationdestinationlist.
type GetSystemAutomationdestinationlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAutomationdestinationlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationdestinationlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAutomationdestinationlist.
type GetSystemAutomationdestinationlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemAutomationdestinationlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationdestinationlistResult)(nil)).Elem()
}

func (o GetSystemAutomationdestinationlistResultOutput) ToGetSystemAutomationdestinationlistResultOutput() GetSystemAutomationdestinationlistResultOutput {
	return o
}

func (o GetSystemAutomationdestinationlistResultOutput) ToGetSystemAutomationdestinationlistResultOutputWithContext(ctx context.Context) GetSystemAutomationdestinationlistResultOutput {
	return o
}

func (o GetSystemAutomationdestinationlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationdestinationlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAutomationdestinationlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAutomationdestinationlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemAutomationdestination`.
func (o GetSystemAutomationdestinationlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAutomationdestinationlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAutomationdestinationlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationdestinationlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAutomationdestinationlistResultOutput{})
}
