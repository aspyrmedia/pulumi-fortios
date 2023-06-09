// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallschedule group
func LookupFirewallscheduleGroup(ctx *pulumi.Context, args *LookupFirewallscheduleGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallscheduleGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallscheduleGroupResult
	err := ctx.Invoke("fortios:index/getFirewallscheduleGroup:getFirewallscheduleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallscheduleGroup.
type LookupFirewallscheduleGroupArgs struct {
	// Specify the name of the desired firewallschedule group.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallscheduleGroup.
type LookupFirewallscheduleGroupResult struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Schedules added to the schedule group. The structure of `member` block is documented below.
	Members []GetFirewallscheduleGroupMember `pulumi:"members"`
	// Schedule name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallscheduleGroupOutput(ctx *pulumi.Context, args LookupFirewallscheduleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallscheduleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallscheduleGroupResult, error) {
			args := v.(LookupFirewallscheduleGroupArgs)
			r, err := LookupFirewallscheduleGroup(ctx, &args, opts...)
			var s LookupFirewallscheduleGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallscheduleGroupResultOutput)
}

// A collection of arguments for invoking getFirewallscheduleGroup.
type LookupFirewallscheduleGroupOutputArgs struct {
	// Specify the name of the desired firewallschedule group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallscheduleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallscheduleGroupArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallscheduleGroup.
type LookupFirewallscheduleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallscheduleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallscheduleGroupResult)(nil)).Elem()
}

func (o LookupFirewallscheduleGroupResultOutput) ToLookupFirewallscheduleGroupResultOutput() LookupFirewallscheduleGroupResultOutput {
	return o
}

func (o LookupFirewallscheduleGroupResultOutput) ToLookupFirewallscheduleGroupResultOutputWithContext(ctx context.Context) LookupFirewallscheduleGroupResultOutput {
	return o
}

// Color of icon on the GUI.
func (o LookupFirewallscheduleGroupResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallscheduleGroupResult) int { return v.Color }).(pulumi.IntOutput)
}

// Security Fabric global object setting.
func (o LookupFirewallscheduleGroupResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallscheduleGroupResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallscheduleGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallscheduleGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Schedules added to the schedule group. The structure of `member` block is documented below.
func (o LookupFirewallscheduleGroupResultOutput) Members() GetFirewallscheduleGroupMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallscheduleGroupResult) []GetFirewallscheduleGroupMember { return v.Members }).(GetFirewallscheduleGroupMemberArrayOutput)
}

// Schedule name.
func (o LookupFirewallscheduleGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallscheduleGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallscheduleGroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallscheduleGroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallscheduleGroupResultOutput{})
}
