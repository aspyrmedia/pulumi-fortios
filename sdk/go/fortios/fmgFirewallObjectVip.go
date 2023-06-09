// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete firewall object virtual ip for FortiManager.
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
//			_, err := fortios.NewFmgFirewallObjectVip(ctx, "test1", &fortios.FmgFirewallObjectVipArgs{
//				ArpReply:      pulumi.String("enable"),
//				Comment:       pulumi.String("test obj vip"),
//				ConfigDefault: pulumi.String("enable"),
//				ExtIntf:       pulumi.String("any"),
//				ExtIp:         pulumi.String("2.2.2.2"),
//				MappedIp:      pulumi.String("1.1.1.1"),
//				Type:          pulumi.String("static-nat"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fortios.NewFmgFirewallObjectVip(ctx, "test2", &fortios.FmgFirewallObjectVipArgs{
//				ArpReply:      pulumi.String("disable"),
//				Comment:       pulumi.String("test obj vip"),
//				ConfigDefault: pulumi.String("enable"),
//				ExtIp:         pulumi.String("2.2.2.2-2.2.2.100"),
//				MappedAddr:    pulumi.String("update.microsoft.com"),
//				Type:          pulumi.String("fqdn"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FmgFirewallObjectVip struct {
	pulumi.CustomResourceState

	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply pulumi.StringPtrOutput `pulumi:"arpReply"`
	// Comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/Disable config default value. enabled by default.
	ConfigDefault pulumi.StringPtrOutput `pulumi:"configDefault"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf pulumi.StringPtrOutput `pulumi:"extIntf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp pulumi.StringPtrOutput `pulumi:"extIp"`
	// Mapped FQDN address name.
	MappedAddr pulumi.StringPtrOutput `pulumi:"mappedAddr"`
	// Mapped Ip address.
	MappedIp pulumi.StringPtrOutput `pulumi:"mappedIp"`
	// Virtual IP name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFmgFirewallObjectVip registers a new resource with the given unique name, arguments, and options.
func NewFmgFirewallObjectVip(ctx *pulumi.Context,
	name string, args *FmgFirewallObjectVipArgs, opts ...pulumi.ResourceOption) (*FmgFirewallObjectVip, error) {
	if args == nil {
		args = &FmgFirewallObjectVipArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FmgFirewallObjectVip
	err := ctx.RegisterResource("fortios:index/fmgFirewallObjectVip:FmgFirewallObjectVip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFmgFirewallObjectVip gets an existing FmgFirewallObjectVip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFmgFirewallObjectVip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FmgFirewallObjectVipState, opts ...pulumi.ResourceOption) (*FmgFirewallObjectVip, error) {
	var resource FmgFirewallObjectVip
	err := ctx.ReadResource("fortios:index/fmgFirewallObjectVip:FmgFirewallObjectVip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FmgFirewallObjectVip resources.
type fmgFirewallObjectVipState struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply *string `pulumi:"arpReply"`
	// Comments.
	Comment *string `pulumi:"comment"`
	// Enable/Disable config default value. enabled by default.
	ConfigDefault *string `pulumi:"configDefault"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf *string `pulumi:"extIntf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp *string `pulumi:"extIp"`
	// Mapped FQDN address name.
	MappedAddr *string `pulumi:"mappedAddr"`
	// Mapped Ip address.
	MappedIp *string `pulumi:"mappedIp"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type *string `pulumi:"type"`
}

type FmgFirewallObjectVipState struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply pulumi.StringPtrInput
	// Comments.
	Comment pulumi.StringPtrInput
	// Enable/Disable config default value. enabled by default.
	ConfigDefault pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp pulumi.StringPtrInput
	// Mapped FQDN address name.
	MappedAddr pulumi.StringPtrInput
	// Mapped Ip address.
	MappedIp pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type pulumi.StringPtrInput
}

func (FmgFirewallObjectVipState) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgFirewallObjectVipState)(nil)).Elem()
}

type fmgFirewallObjectVipArgs struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply *string `pulumi:"arpReply"`
	// Comments.
	Comment *string `pulumi:"comment"`
	// Enable/Disable config default value. enabled by default.
	ConfigDefault *string `pulumi:"configDefault"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf *string `pulumi:"extIntf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp *string `pulumi:"extIp"`
	// Mapped FQDN address name.
	MappedAddr *string `pulumi:"mappedAddr"`
	// Mapped Ip address.
	MappedIp *string `pulumi:"mappedIp"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FmgFirewallObjectVip resource.
type FmgFirewallObjectVipArgs struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply pulumi.StringPtrInput
	// Comments.
	Comment pulumi.StringPtrInput
	// Enable/Disable config default value. enabled by default.
	ConfigDefault pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp pulumi.StringPtrInput
	// Mapped FQDN address name.
	MappedAddr pulumi.StringPtrInput
	// Mapped Ip address.
	MappedIp pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type pulumi.StringPtrInput
}

func (FmgFirewallObjectVipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgFirewallObjectVipArgs)(nil)).Elem()
}

type FmgFirewallObjectVipInput interface {
	pulumi.Input

	ToFmgFirewallObjectVipOutput() FmgFirewallObjectVipOutput
	ToFmgFirewallObjectVipOutputWithContext(ctx context.Context) FmgFirewallObjectVipOutput
}

func (*FmgFirewallObjectVip) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgFirewallObjectVip)(nil)).Elem()
}

func (i *FmgFirewallObjectVip) ToFmgFirewallObjectVipOutput() FmgFirewallObjectVipOutput {
	return i.ToFmgFirewallObjectVipOutputWithContext(context.Background())
}

func (i *FmgFirewallObjectVip) ToFmgFirewallObjectVipOutputWithContext(ctx context.Context) FmgFirewallObjectVipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgFirewallObjectVipOutput)
}

// FmgFirewallObjectVipArrayInput is an input type that accepts FmgFirewallObjectVipArray and FmgFirewallObjectVipArrayOutput values.
// You can construct a concrete instance of `FmgFirewallObjectVipArrayInput` via:
//
//	FmgFirewallObjectVipArray{ FmgFirewallObjectVipArgs{...} }
type FmgFirewallObjectVipArrayInput interface {
	pulumi.Input

	ToFmgFirewallObjectVipArrayOutput() FmgFirewallObjectVipArrayOutput
	ToFmgFirewallObjectVipArrayOutputWithContext(context.Context) FmgFirewallObjectVipArrayOutput
}

type FmgFirewallObjectVipArray []FmgFirewallObjectVipInput

func (FmgFirewallObjectVipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgFirewallObjectVip)(nil)).Elem()
}

func (i FmgFirewallObjectVipArray) ToFmgFirewallObjectVipArrayOutput() FmgFirewallObjectVipArrayOutput {
	return i.ToFmgFirewallObjectVipArrayOutputWithContext(context.Background())
}

func (i FmgFirewallObjectVipArray) ToFmgFirewallObjectVipArrayOutputWithContext(ctx context.Context) FmgFirewallObjectVipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgFirewallObjectVipArrayOutput)
}

// FmgFirewallObjectVipMapInput is an input type that accepts FmgFirewallObjectVipMap and FmgFirewallObjectVipMapOutput values.
// You can construct a concrete instance of `FmgFirewallObjectVipMapInput` via:
//
//	FmgFirewallObjectVipMap{ "key": FmgFirewallObjectVipArgs{...} }
type FmgFirewallObjectVipMapInput interface {
	pulumi.Input

	ToFmgFirewallObjectVipMapOutput() FmgFirewallObjectVipMapOutput
	ToFmgFirewallObjectVipMapOutputWithContext(context.Context) FmgFirewallObjectVipMapOutput
}

type FmgFirewallObjectVipMap map[string]FmgFirewallObjectVipInput

func (FmgFirewallObjectVipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgFirewallObjectVip)(nil)).Elem()
}

func (i FmgFirewallObjectVipMap) ToFmgFirewallObjectVipMapOutput() FmgFirewallObjectVipMapOutput {
	return i.ToFmgFirewallObjectVipMapOutputWithContext(context.Background())
}

func (i FmgFirewallObjectVipMap) ToFmgFirewallObjectVipMapOutputWithContext(ctx context.Context) FmgFirewallObjectVipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgFirewallObjectVipMapOutput)
}

type FmgFirewallObjectVipOutput struct{ *pulumi.OutputState }

func (FmgFirewallObjectVipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgFirewallObjectVip)(nil)).Elem()
}

func (o FmgFirewallObjectVipOutput) ToFmgFirewallObjectVipOutput() FmgFirewallObjectVipOutput {
	return o
}

func (o FmgFirewallObjectVipOutput) ToFmgFirewallObjectVipOutputWithContext(ctx context.Context) FmgFirewallObjectVipOutput {
	return o
}

// ADOM name. default is 'root'.
func (o FmgFirewallObjectVipOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
func (o FmgFirewallObjectVipOutput) ArpReply() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.ArpReply }).(pulumi.StringPtrOutput)
}

// Comments.
func (o FmgFirewallObjectVipOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Enable/Disable config default value. enabled by default.
func (o FmgFirewallObjectVipOutput) ConfigDefault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.ConfigDefault }).(pulumi.StringPtrOutput)
}

// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
func (o FmgFirewallObjectVipOutput) ExtIntf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.ExtIntf }).(pulumi.StringPtrOutput)
}

// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
func (o FmgFirewallObjectVipOutput) ExtIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.ExtIp }).(pulumi.StringPtrOutput)
}

// Mapped FQDN address name.
func (o FmgFirewallObjectVipOutput) MappedAddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.MappedAddr }).(pulumi.StringPtrOutput)
}

// Mapped Ip address.
func (o FmgFirewallObjectVipOutput) MappedIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.MappedIp }).(pulumi.StringPtrOutput)
}

// Virtual IP name.
func (o FmgFirewallObjectVipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
func (o FmgFirewallObjectVipOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgFirewallObjectVip) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type FmgFirewallObjectVipArrayOutput struct{ *pulumi.OutputState }

func (FmgFirewallObjectVipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgFirewallObjectVip)(nil)).Elem()
}

func (o FmgFirewallObjectVipArrayOutput) ToFmgFirewallObjectVipArrayOutput() FmgFirewallObjectVipArrayOutput {
	return o
}

func (o FmgFirewallObjectVipArrayOutput) ToFmgFirewallObjectVipArrayOutputWithContext(ctx context.Context) FmgFirewallObjectVipArrayOutput {
	return o
}

func (o FmgFirewallObjectVipArrayOutput) Index(i pulumi.IntInput) FmgFirewallObjectVipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FmgFirewallObjectVip {
		return vs[0].([]*FmgFirewallObjectVip)[vs[1].(int)]
	}).(FmgFirewallObjectVipOutput)
}

type FmgFirewallObjectVipMapOutput struct{ *pulumi.OutputState }

func (FmgFirewallObjectVipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgFirewallObjectVip)(nil)).Elem()
}

func (o FmgFirewallObjectVipMapOutput) ToFmgFirewallObjectVipMapOutput() FmgFirewallObjectVipMapOutput {
	return o
}

func (o FmgFirewallObjectVipMapOutput) ToFmgFirewallObjectVipMapOutputWithContext(ctx context.Context) FmgFirewallObjectVipMapOutput {
	return o
}

func (o FmgFirewallObjectVipMapOutput) MapIndex(k pulumi.StringInput) FmgFirewallObjectVipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FmgFirewallObjectVip {
		return vs[0].(map[string]*FmgFirewallObjectVip)[vs[1].(string)]
	}).(FmgFirewallObjectVipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FmgFirewallObjectVipInput)(nil)).Elem(), &FmgFirewallObjectVip{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgFirewallObjectVipArrayInput)(nil)).Elem(), FmgFirewallObjectVipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgFirewallObjectVipMapInput)(nil)).Elem(), FmgFirewallObjectVipMap{})
	pulumi.RegisterOutputType(FmgFirewallObjectVipOutput{})
	pulumi.RegisterOutputType(FmgFirewallObjectVipArrayOutput{})
	pulumi.RegisterOutputType(FmgFirewallObjectVipMapOutput{})
}
