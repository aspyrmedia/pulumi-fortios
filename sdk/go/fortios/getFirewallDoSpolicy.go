// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall DoSpolicy
func LookupFirewallDoSpolicy(ctx *pulumi.Context, args *LookupFirewallDoSpolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallDoSpolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallDoSpolicyResult
	err := ctx.Invoke("fortios:index/getFirewallDoSpolicy:getFirewallDoSpolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallDoSpolicy.
type LookupFirewallDoSpolicyArgs struct {
	// Specify the policyid of the desired firewall DoSpolicy.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallDoSpolicy.
type LookupFirewallDoSpolicyResult struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []GetFirewallDoSpolicyAnomaly `pulumi:"anomalies"`
	// Comment.
	Comments string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetFirewallDoSpolicyDstaddr `pulumi:"dstaddrs"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Incoming interface name from available interfaces.
	Interface string `pulumi:"interface"`
	// Anomaly name.
	Name string `pulumi:"name"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []GetFirewallDoSpolicyService `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetFirewallDoSpolicySrcaddr `pulumi:"srcaddrs"`
	// Enable/disable this anomaly.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallDoSpolicyOutput(ctx *pulumi.Context, args LookupFirewallDoSpolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallDoSpolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallDoSpolicyResult, error) {
			args := v.(LookupFirewallDoSpolicyArgs)
			r, err := LookupFirewallDoSpolicy(ctx, &args, opts...)
			var s LookupFirewallDoSpolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallDoSpolicyResultOutput)
}

// A collection of arguments for invoking getFirewallDoSpolicy.
type LookupFirewallDoSpolicyOutputArgs struct {
	// Specify the policyid of the desired firewall DoSpolicy.
	Policyid pulumi.IntInput `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallDoSpolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallDoSpolicyArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallDoSpolicy.
type LookupFirewallDoSpolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallDoSpolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallDoSpolicyResult)(nil)).Elem()
}

func (o LookupFirewallDoSpolicyResultOutput) ToLookupFirewallDoSpolicyResultOutput() LookupFirewallDoSpolicyResultOutput {
	return o
}

func (o LookupFirewallDoSpolicyResultOutput) ToLookupFirewallDoSpolicyResultOutputWithContext(ctx context.Context) LookupFirewallDoSpolicyResultOutput {
	return o
}

// Anomaly name. The structure of `anomaly` block is documented below.
func (o LookupFirewallDoSpolicyResultOutput) Anomalies() GetFirewallDoSpolicyAnomalyArrayOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) []GetFirewallDoSpolicyAnomaly { return v.Anomalies }).(GetFirewallDoSpolicyAnomalyArrayOutput)
}

// Comment.
func (o LookupFirewallDoSpolicyResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) string { return v.Comments }).(pulumi.StringOutput)
}

// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
func (o LookupFirewallDoSpolicyResultOutput) Dstaddrs() GetFirewallDoSpolicyDstaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) []GetFirewallDoSpolicyDstaddr { return v.Dstaddrs }).(GetFirewallDoSpolicyDstaddrArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallDoSpolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Incoming interface name from available interfaces.
func (o LookupFirewallDoSpolicyResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Anomaly name.
func (o LookupFirewallDoSpolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy ID.
func (o LookupFirewallDoSpolicyResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) int { return v.Policyid }).(pulumi.IntOutput)
}

// Service object from available options. The structure of `service` block is documented below.
func (o LookupFirewallDoSpolicyResultOutput) Services() GetFirewallDoSpolicyServiceArrayOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) []GetFirewallDoSpolicyService { return v.Services }).(GetFirewallDoSpolicyServiceArrayOutput)
}

// Source address name from available addresses. The structure of `srcaddr` block is documented below.
func (o LookupFirewallDoSpolicyResultOutput) Srcaddrs() GetFirewallDoSpolicySrcaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) []GetFirewallDoSpolicySrcaddr { return v.Srcaddrs }).(GetFirewallDoSpolicySrcaddrArrayOutput)
}

// Enable/disable this anomaly.
func (o LookupFirewallDoSpolicyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFirewallDoSpolicyResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallDoSpolicyResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallDoSpolicyResultOutput{})
}