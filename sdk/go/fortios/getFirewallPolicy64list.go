// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallPolicy64`.
func GetFirewallPolicy64list(ctx *pulumi.Context, args *GetFirewallPolicy64listArgs, opts ...pulumi.InvokeOption) (*GetFirewallPolicy64listResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallPolicy64listResult
	err := ctx.Invoke("fortios:index/getFirewallPolicy64list:getFirewallPolicy64list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallPolicy64list.
type GetFirewallPolicy64listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallPolicy64list.
type GetFirewallPolicy64listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallPolicy64`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetFirewallPolicy64listOutput(ctx *pulumi.Context, args GetFirewallPolicy64listOutputArgs, opts ...pulumi.InvokeOption) GetFirewallPolicy64listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallPolicy64listResult, error) {
			args := v.(GetFirewallPolicy64listArgs)
			r, err := GetFirewallPolicy64list(ctx, &args, opts...)
			var s GetFirewallPolicy64listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallPolicy64listResultOutput)
}

// A collection of arguments for invoking getFirewallPolicy64list.
type GetFirewallPolicy64listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallPolicy64listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallPolicy64listArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallPolicy64list.
type GetFirewallPolicy64listResultOutput struct{ *pulumi.OutputState }

func (GetFirewallPolicy64listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallPolicy64listResult)(nil)).Elem()
}

func (o GetFirewallPolicy64listResultOutput) ToGetFirewallPolicy64listResultOutput() GetFirewallPolicy64listResultOutput {
	return o
}

func (o GetFirewallPolicy64listResultOutput) ToGetFirewallPolicy64listResultOutputWithContext(ctx context.Context) GetFirewallPolicy64listResultOutput {
	return o
}

func (o GetFirewallPolicy64listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallPolicy64listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallPolicy64listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallPolicy64listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallPolicy64`.
func (o GetFirewallPolicy64listResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetFirewallPolicy64listResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetFirewallPolicy64listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallPolicy64listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallPolicy64listResultOutput{})
}