// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `FirewallVip`, we recommend that you use the new resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewFirewallObjectVip(ctx, "v11", &fortios.FirewallObjectVipArgs{
//				Comment: pulumi.String("fdsafdsafds"),
//				Extintf: pulumi.String("port3"),
//				Extip:   pulumi.String("11.1.1.1-21.1.1.1"),
//				Extport: pulumi.String("2-3"),
//				Mappedips: pulumi.StringArray{
//					pulumi.String("22.2.2.2-32.2.2.2"),
//				},
//				Mappedport:  pulumi.String("4-5"),
//				Portforward: pulumi.String("enable"),
//				Protocol:    pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FirewallObjectVip struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf pulumi.StringOutput `pulumi:"extintf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip pulumi.StringOutput `pulumi:"extip"`
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport pulumi.StringOutput `pulumi:"extport"`
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips pulumi.StringArrayOutput `pulumi:"mappedips"`
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport pulumi.StringOutput `pulumi:"mappedport"`
	// Virtual IP name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable port forwarding.
	Portforward pulumi.StringOutput `pulumi:"portforward"`
	// Protocol to use when forwarding packets.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
}

// NewFirewallObjectVip registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectVip(ctx *pulumi.Context,
	name string, args *FirewallObjectVipArgs, opts ...pulumi.ResourceOption) (*FirewallObjectVip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Extip == nil {
		return nil, errors.New("invalid value for required argument 'Extip'")
	}
	if args.Mappedips == nil {
		return nil, errors.New("invalid value for required argument 'Mappedips'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallObjectVip
	err := ctx.RegisterResource("fortios:index/firewallObjectVip:FirewallObjectVip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectVip gets an existing FirewallObjectVip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectVip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectVipState, opts ...pulumi.ResourceOption) (*FirewallObjectVip, error) {
	var resource FirewallObjectVip
	err := ctx.ReadResource("fortios:index/firewallObjectVip:FirewallObjectVip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectVip resources.
type firewallObjectVipState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf *string `pulumi:"extintf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip *string `pulumi:"extip"`
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport *string `pulumi:"extport"`
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips []string `pulumi:"mappedips"`
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport *string `pulumi:"mappedport"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Enable/disable port forwarding.
	Portforward *string `pulumi:"portforward"`
	// Protocol to use when forwarding packets.
	Protocol *string `pulumi:"protocol"`
}

type FirewallObjectVipState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip pulumi.StringPtrInput
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport pulumi.StringPtrInput
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips pulumi.StringArrayInput
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Enable/disable port forwarding.
	Portforward pulumi.StringPtrInput
	// Protocol to use when forwarding packets.
	Protocol pulumi.StringPtrInput
}

func (FirewallObjectVipState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectVipState)(nil)).Elem()
}

type firewallObjectVipArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf *string `pulumi:"extintf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip string `pulumi:"extip"`
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport *string `pulumi:"extport"`
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips []string `pulumi:"mappedips"`
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport *string `pulumi:"mappedport"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Enable/disable port forwarding.
	Portforward *string `pulumi:"portforward"`
	// Protocol to use when forwarding packets.
	Protocol *string `pulumi:"protocol"`
}

// The set of arguments for constructing a FirewallObjectVip resource.
type FirewallObjectVipArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip pulumi.StringInput
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport pulumi.StringPtrInput
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips pulumi.StringArrayInput
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Enable/disable port forwarding.
	Portforward pulumi.StringPtrInput
	// Protocol to use when forwarding packets.
	Protocol pulumi.StringPtrInput
}

func (FirewallObjectVipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectVipArgs)(nil)).Elem()
}

type FirewallObjectVipInput interface {
	pulumi.Input

	ToFirewallObjectVipOutput() FirewallObjectVipOutput
	ToFirewallObjectVipOutputWithContext(ctx context.Context) FirewallObjectVipOutput
}

func (*FirewallObjectVip) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectVip)(nil)).Elem()
}

func (i *FirewallObjectVip) ToFirewallObjectVipOutput() FirewallObjectVipOutput {
	return i.ToFirewallObjectVipOutputWithContext(context.Background())
}

func (i *FirewallObjectVip) ToFirewallObjectVipOutputWithContext(ctx context.Context) FirewallObjectVipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectVipOutput)
}

// FirewallObjectVipArrayInput is an input type that accepts FirewallObjectVipArray and FirewallObjectVipArrayOutput values.
// You can construct a concrete instance of `FirewallObjectVipArrayInput` via:
//
//	FirewallObjectVipArray{ FirewallObjectVipArgs{...} }
type FirewallObjectVipArrayInput interface {
	pulumi.Input

	ToFirewallObjectVipArrayOutput() FirewallObjectVipArrayOutput
	ToFirewallObjectVipArrayOutputWithContext(context.Context) FirewallObjectVipArrayOutput
}

type FirewallObjectVipArray []FirewallObjectVipInput

func (FirewallObjectVipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectVip)(nil)).Elem()
}

func (i FirewallObjectVipArray) ToFirewallObjectVipArrayOutput() FirewallObjectVipArrayOutput {
	return i.ToFirewallObjectVipArrayOutputWithContext(context.Background())
}

func (i FirewallObjectVipArray) ToFirewallObjectVipArrayOutputWithContext(ctx context.Context) FirewallObjectVipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectVipArrayOutput)
}

// FirewallObjectVipMapInput is an input type that accepts FirewallObjectVipMap and FirewallObjectVipMapOutput values.
// You can construct a concrete instance of `FirewallObjectVipMapInput` via:
//
//	FirewallObjectVipMap{ "key": FirewallObjectVipArgs{...} }
type FirewallObjectVipMapInput interface {
	pulumi.Input

	ToFirewallObjectVipMapOutput() FirewallObjectVipMapOutput
	ToFirewallObjectVipMapOutputWithContext(context.Context) FirewallObjectVipMapOutput
}

type FirewallObjectVipMap map[string]FirewallObjectVipInput

func (FirewallObjectVipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectVip)(nil)).Elem()
}

func (i FirewallObjectVipMap) ToFirewallObjectVipMapOutput() FirewallObjectVipMapOutput {
	return i.ToFirewallObjectVipMapOutputWithContext(context.Background())
}

func (i FirewallObjectVipMap) ToFirewallObjectVipMapOutputWithContext(ctx context.Context) FirewallObjectVipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectVipMapOutput)
}

type FirewallObjectVipOutput struct{ *pulumi.OutputState }

func (FirewallObjectVipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectVip)(nil)).Elem()
}

func (o FirewallObjectVipOutput) ToFirewallObjectVipOutput() FirewallObjectVipOutput {
	return o
}

func (o FirewallObjectVipOutput) ToFirewallObjectVipOutputWithContext(ctx context.Context) FirewallObjectVipOutput {
	return o
}

// Comment.
func (o FirewallObjectVipOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
func (o FirewallObjectVipOutput) Extintf() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Extintf }).(pulumi.StringOutput)
}

// IP address or address range on the external interface that you want to map to an address or address range on the
// destination network.
func (o FirewallObjectVipOutput) Extip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Extip }).(pulumi.StringOutput)
}

// Incoming port number range that you want to map to a port number range on the destination network.
func (o FirewallObjectVipOutput) Extport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Extport }).(pulumi.StringOutput)
}

// IP address or address range on the destination network to which the external IP address is mapped.
func (o FirewallObjectVipOutput) Mappedips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringArrayOutput { return v.Mappedips }).(pulumi.StringArrayOutput)
}

// Port number range on the destination network to which the external port number range is mapped.
func (o FirewallObjectVipOutput) Mappedport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Mappedport }).(pulumi.StringOutput)
}

// Virtual IP name.
func (o FirewallObjectVipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable port forwarding.
func (o FirewallObjectVipOutput) Portforward() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Portforward }).(pulumi.StringOutput)
}

// Protocol to use when forwarding packets.
func (o FirewallObjectVipOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVip) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

type FirewallObjectVipArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectVipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectVip)(nil)).Elem()
}

func (o FirewallObjectVipArrayOutput) ToFirewallObjectVipArrayOutput() FirewallObjectVipArrayOutput {
	return o
}

func (o FirewallObjectVipArrayOutput) ToFirewallObjectVipArrayOutputWithContext(ctx context.Context) FirewallObjectVipArrayOutput {
	return o
}

func (o FirewallObjectVipArrayOutput) Index(i pulumi.IntInput) FirewallObjectVipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallObjectVip {
		return vs[0].([]*FirewallObjectVip)[vs[1].(int)]
	}).(FirewallObjectVipOutput)
}

type FirewallObjectVipMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectVipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectVip)(nil)).Elem()
}

func (o FirewallObjectVipMapOutput) ToFirewallObjectVipMapOutput() FirewallObjectVipMapOutput {
	return o
}

func (o FirewallObjectVipMapOutput) ToFirewallObjectVipMapOutputWithContext(ctx context.Context) FirewallObjectVipMapOutput {
	return o
}

func (o FirewallObjectVipMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectVipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallObjectVip {
		return vs[0].(map[string]*FirewallObjectVip)[vs[1].(string)]
	}).(FirewallObjectVipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectVipInput)(nil)).Elem(), &FirewallObjectVip{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectVipArrayInput)(nil)).Elem(), FirewallObjectVipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectVipMapInput)(nil)).Elem(), FirewallObjectVipMap{})
	pulumi.RegisterOutputType(FirewallObjectVipOutput{})
	pulumi.RegisterOutputType(FirewallObjectVipArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectVipMapOutput{})
}