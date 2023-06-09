// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router policy6
func LookupRouterPolicy6(ctx *pulumi.Context, args *LookupRouterPolicy6Args, opts ...pulumi.InvokeOption) (*LookupRouterPolicy6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterPolicy6Result
	err := ctx.Invoke("fortios:index/getRouterPolicy6:getRouterPolicy6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterPolicy6.
type LookupRouterPolicy6Args struct {
	// Specify the seqNum of the desired router policy6.
	SeqNum int `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterPolicy6.
type LookupRouterPolicy6Result struct {
	// Action of the policy route.
	Action string `pulumi:"action"`
	// Optional comments.
	Comments string `pulumi:"comments"`
	// Destination IPv6 prefix.
	Dst string `pulumi:"dst"`
	// Enable/disable negating destination address match.
	DstNegate string `pulumi:"dstNegate"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetRouterPolicy6Dstaddr `pulumi:"dstaddrs"`
	// End destination port number (1 - 65535).
	EndPort int `pulumi:"endPort"`
	// IPv6 address of the gateway.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Incoming interface name.
	InputDevice string `pulumi:"inputDevice"`
	// Enable/disable negation of input device match.
	InputDeviceNegate string `pulumi:"inputDeviceNegate"`
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []GetRouterPolicy6InternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []GetRouterPolicy6InternetServiceId `pulumi:"internetServiceIds"`
	// Outgoing interface name.
	OutputDevice string `pulumi:"outputDevice"`
	// Protocol number (0 - 255).
	Protocol int `pulumi:"protocol"`
	// Sequence number.
	SeqNum int `pulumi:"seqNum"`
	// Source IPv6 prefix.
	Src string `pulumi:"src"`
	// Enable/disable negating source address match.
	SrcNegate string `pulumi:"srcNegate"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetRouterPolicy6Srcaddr `pulumi:"srcaddrs"`
	// Start destination port number (1 - 65535).
	StartPort int `pulumi:"startPort"`
	// Enable/disable this policy route.
	Status string `pulumi:"status"`
	// Type of service bit pattern.
	Tos string `pulumi:"tos"`
	// Type of service evaluated bits.
	TosMask   string  `pulumi:"tosMask"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupRouterPolicy6Output(ctx *pulumi.Context, args LookupRouterPolicy6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterPolicy6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterPolicy6Result, error) {
			args := v.(LookupRouterPolicy6Args)
			r, err := LookupRouterPolicy6(ctx, &args, opts...)
			var s LookupRouterPolicy6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterPolicy6ResultOutput)
}

// A collection of arguments for invoking getRouterPolicy6.
type LookupRouterPolicy6OutputArgs struct {
	// Specify the seqNum of the desired router policy6.
	SeqNum pulumi.IntInput `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterPolicy6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterPolicy6Args)(nil)).Elem()
}

// A collection of values returned by getRouterPolicy6.
type LookupRouterPolicy6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterPolicy6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterPolicy6Result)(nil)).Elem()
}

func (o LookupRouterPolicy6ResultOutput) ToLookupRouterPolicy6ResultOutput() LookupRouterPolicy6ResultOutput {
	return o
}

func (o LookupRouterPolicy6ResultOutput) ToLookupRouterPolicy6ResultOutputWithContext(ctx context.Context) LookupRouterPolicy6ResultOutput {
	return o
}

// Action of the policy route.
func (o LookupRouterPolicy6ResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Action }).(pulumi.StringOutput)
}

// Optional comments.
func (o LookupRouterPolicy6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// Destination IPv6 prefix.
func (o LookupRouterPolicy6ResultOutput) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Dst }).(pulumi.StringOutput)
}

// Enable/disable negating destination address match.
func (o LookupRouterPolicy6ResultOutput) DstNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.DstNegate }).(pulumi.StringOutput)
}

// Destination address name. The structure of `dstaddr` block is documented below.
func (o LookupRouterPolicy6ResultOutput) Dstaddrs() GetRouterPolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) []GetRouterPolicy6Dstaddr { return v.Dstaddrs }).(GetRouterPolicy6DstaddrArrayOutput)
}

// End destination port number (1 - 65535).
func (o LookupRouterPolicy6ResultOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) int { return v.EndPort }).(pulumi.IntOutput)
}

// IPv6 address of the gateway.
func (o LookupRouterPolicy6ResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterPolicy6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Incoming interface name.
func (o LookupRouterPolicy6ResultOutput) InputDevice() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.InputDevice }).(pulumi.StringOutput)
}

// Enable/disable negation of input device match.
func (o LookupRouterPolicy6ResultOutput) InputDeviceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.InputDeviceNegate }).(pulumi.StringOutput)
}

// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
func (o LookupRouterPolicy6ResultOutput) InternetServiceCustoms() GetRouterPolicy6InternetServiceCustomArrayOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) []GetRouterPolicy6InternetServiceCustom {
		return v.InternetServiceCustoms
	}).(GetRouterPolicy6InternetServiceCustomArrayOutput)
}

// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
func (o LookupRouterPolicy6ResultOutput) InternetServiceIds() GetRouterPolicy6InternetServiceIdArrayOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) []GetRouterPolicy6InternetServiceId { return v.InternetServiceIds }).(GetRouterPolicy6InternetServiceIdArrayOutput)
}

// Outgoing interface name.
func (o LookupRouterPolicy6ResultOutput) OutputDevice() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.OutputDevice }).(pulumi.StringOutput)
}

// Protocol number (0 - 255).
func (o LookupRouterPolicy6ResultOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) int { return v.Protocol }).(pulumi.IntOutput)
}

// Sequence number.
func (o LookupRouterPolicy6ResultOutput) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) int { return v.SeqNum }).(pulumi.IntOutput)
}

// Source IPv6 prefix.
func (o LookupRouterPolicy6ResultOutput) Src() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Src }).(pulumi.StringOutput)
}

// Enable/disable negating source address match.
func (o LookupRouterPolicy6ResultOutput) SrcNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.SrcNegate }).(pulumi.StringOutput)
}

// Source address name. The structure of `srcaddr` block is documented below.
func (o LookupRouterPolicy6ResultOutput) Srcaddrs() GetRouterPolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) []GetRouterPolicy6Srcaddr { return v.Srcaddrs }).(GetRouterPolicy6SrcaddrArrayOutput)
}

// Start destination port number (1 - 65535).
func (o LookupRouterPolicy6ResultOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) int { return v.StartPort }).(pulumi.IntOutput)
}

// Enable/disable this policy route.
func (o LookupRouterPolicy6ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Status }).(pulumi.StringOutput)
}

// Type of service bit pattern.
func (o LookupRouterPolicy6ResultOutput) Tos() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.Tos }).(pulumi.StringOutput)
}

// Type of service evaluated bits.
func (o LookupRouterPolicy6ResultOutput) TosMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) string { return v.TosMask }).(pulumi.StringOutput)
}

func (o LookupRouterPolicy6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterPolicy6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterPolicy6ResultOutput{})
}
