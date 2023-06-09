// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallserviceGroup`.
func GetFirewallserviceGrouplist(ctx *pulumi.Context, args *GetFirewallserviceGrouplistArgs, opts ...pulumi.InvokeOption) (*GetFirewallserviceGrouplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallserviceGrouplistResult
	err := ctx.Invoke("fortios:index/getFirewallserviceGrouplist:getFirewallserviceGrouplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallserviceGrouplist.
type GetFirewallserviceGrouplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallserviceGrouplist.
type GetFirewallserviceGrouplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallserviceGroup`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallserviceGrouplistOutput(ctx *pulumi.Context, args GetFirewallserviceGrouplistOutputArgs, opts ...pulumi.InvokeOption) GetFirewallserviceGrouplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallserviceGrouplistResult, error) {
			args := v.(GetFirewallserviceGrouplistArgs)
			r, err := GetFirewallserviceGrouplist(ctx, &args, opts...)
			var s GetFirewallserviceGrouplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallserviceGrouplistResultOutput)
}

// A collection of arguments for invoking getFirewallserviceGrouplist.
type GetFirewallserviceGrouplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallserviceGrouplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallserviceGrouplistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallserviceGrouplist.
type GetFirewallserviceGrouplistResultOutput struct{ *pulumi.OutputState }

func (GetFirewallserviceGrouplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallserviceGrouplistResult)(nil)).Elem()
}

func (o GetFirewallserviceGrouplistResultOutput) ToGetFirewallserviceGrouplistResultOutput() GetFirewallserviceGrouplistResultOutput {
	return o
}

func (o GetFirewallserviceGrouplistResultOutput) ToGetFirewallserviceGrouplistResultOutputWithContext(ctx context.Context) GetFirewallserviceGrouplistResultOutput {
	return o
}

func (o GetFirewallserviceGrouplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallserviceGrouplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallserviceGrouplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallserviceGrouplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallserviceGroup`.
func (o GetFirewallserviceGrouplistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallserviceGrouplistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallserviceGrouplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallserviceGrouplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallserviceGrouplistResultOutput{})
}
