// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS VDOM parameter.
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
//			_, err := fortios.NewIpsSettings(ctx, "trname", &fortios.IpsSettingsArgs{
//				IpsPacketQuota:      pulumi.Int(0),
//				PacketLogHistory:    pulumi.Int(1),
//				PacketLogMemory:     pulumi.Int(256),
//				PacketLogPostAttack: pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Ips Settings can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/ipsSettings:IpsSettings labelname IpsSettings
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/ipsSettings:IpsSettings labelname IpsSettings
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type IpsSettings struct {
	pulumi.CustomResourceState

	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota pulumi.IntOutput `pulumi:"ipsPacketQuota"`
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory pulumi.IntOutput `pulumi:"packetLogHistory"`
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory pulumi.IntOutput `pulumi:"packetLogMemory"`
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack pulumi.IntOutput `pulumi:"packetLogPostAttack"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsSettings registers a new resource with the given unique name, arguments, and options.
func NewIpsSettings(ctx *pulumi.Context,
	name string, args *IpsSettingsArgs, opts ...pulumi.ResourceOption) (*IpsSettings, error) {
	if args == nil {
		args = &IpsSettingsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsSettings
	err := ctx.RegisterResource("fortios:index/ipsSettings:IpsSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsSettings gets an existing IpsSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsSettingsState, opts ...pulumi.ResourceOption) (*IpsSettings, error) {
	var resource IpsSettings
	err := ctx.ReadResource("fortios:index/ipsSettings:IpsSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsSettings resources.
type ipsSettingsState struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota *int `pulumi:"ipsPacketQuota"`
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory *int `pulumi:"packetLogHistory"`
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory *int `pulumi:"packetLogMemory"`
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack *int `pulumi:"packetLogPostAttack"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsSettingsState struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota pulumi.IntPtrInput
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory pulumi.IntPtrInput
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory pulumi.IntPtrInput
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSettingsState)(nil)).Elem()
}

type ipsSettingsArgs struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota *int `pulumi:"ipsPacketQuota"`
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory *int `pulumi:"packetLogHistory"`
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory *int `pulumi:"packetLogMemory"`
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack *int `pulumi:"packetLogPostAttack"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsSettings resource.
type IpsSettingsArgs struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota pulumi.IntPtrInput
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory pulumi.IntPtrInput
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory pulumi.IntPtrInput
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSettingsArgs)(nil)).Elem()
}

type IpsSettingsInput interface {
	pulumi.Input

	ToIpsSettingsOutput() IpsSettingsOutput
	ToIpsSettingsOutputWithContext(ctx context.Context) IpsSettingsOutput
}

func (*IpsSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSettings)(nil)).Elem()
}

func (i *IpsSettings) ToIpsSettingsOutput() IpsSettingsOutput {
	return i.ToIpsSettingsOutputWithContext(context.Background())
}

func (i *IpsSettings) ToIpsSettingsOutputWithContext(ctx context.Context) IpsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsOutput)
}

// IpsSettingsArrayInput is an input type that accepts IpsSettingsArray and IpsSettingsArrayOutput values.
// You can construct a concrete instance of `IpsSettingsArrayInput` via:
//
//	IpsSettingsArray{ IpsSettingsArgs{...} }
type IpsSettingsArrayInput interface {
	pulumi.Input

	ToIpsSettingsArrayOutput() IpsSettingsArrayOutput
	ToIpsSettingsArrayOutputWithContext(context.Context) IpsSettingsArrayOutput
}

type IpsSettingsArray []IpsSettingsInput

func (IpsSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsSettings)(nil)).Elem()
}

func (i IpsSettingsArray) ToIpsSettingsArrayOutput() IpsSettingsArrayOutput {
	return i.ToIpsSettingsArrayOutputWithContext(context.Background())
}

func (i IpsSettingsArray) ToIpsSettingsArrayOutputWithContext(ctx context.Context) IpsSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsArrayOutput)
}

// IpsSettingsMapInput is an input type that accepts IpsSettingsMap and IpsSettingsMapOutput values.
// You can construct a concrete instance of `IpsSettingsMapInput` via:
//
//	IpsSettingsMap{ "key": IpsSettingsArgs{...} }
type IpsSettingsMapInput interface {
	pulumi.Input

	ToIpsSettingsMapOutput() IpsSettingsMapOutput
	ToIpsSettingsMapOutputWithContext(context.Context) IpsSettingsMapOutput
}

type IpsSettingsMap map[string]IpsSettingsInput

func (IpsSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsSettings)(nil)).Elem()
}

func (i IpsSettingsMap) ToIpsSettingsMapOutput() IpsSettingsMapOutput {
	return i.ToIpsSettingsMapOutputWithContext(context.Background())
}

func (i IpsSettingsMap) ToIpsSettingsMapOutputWithContext(ctx context.Context) IpsSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsMapOutput)
}

type IpsSettingsOutput struct{ *pulumi.OutputState }

func (IpsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSettings)(nil)).Elem()
}

func (o IpsSettingsOutput) ToIpsSettingsOutput() IpsSettingsOutput {
	return o
}

func (o IpsSettingsOutput) ToIpsSettingsOutputWithContext(ctx context.Context) IpsSettingsOutput {
	return o
}

// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
func (o IpsSettingsOutput) IpsPacketQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsSettings) pulumi.IntOutput { return v.IpsPacketQuota }).(pulumi.IntOutput)
}

// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
func (o IpsSettingsOutput) PacketLogHistory() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsSettings) pulumi.IntOutput { return v.PacketLogHistory }).(pulumi.IntOutput)
}

// Maximum memory can be used by packet log (64 - 8192 kB).
func (o IpsSettingsOutput) PacketLogMemory() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsSettings) pulumi.IntOutput { return v.PacketLogMemory }).(pulumi.IntOutput)
}

// Number of packets to log after the IPS signature is detected (0 - 255).
func (o IpsSettingsOutput) PacketLogPostAttack() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsSettings) pulumi.IntOutput { return v.PacketLogPostAttack }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IpsSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IpsSettingsArrayOutput struct{ *pulumi.OutputState }

func (IpsSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsSettings)(nil)).Elem()
}

func (o IpsSettingsArrayOutput) ToIpsSettingsArrayOutput() IpsSettingsArrayOutput {
	return o
}

func (o IpsSettingsArrayOutput) ToIpsSettingsArrayOutputWithContext(ctx context.Context) IpsSettingsArrayOutput {
	return o
}

func (o IpsSettingsArrayOutput) Index(i pulumi.IntInput) IpsSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsSettings {
		return vs[0].([]*IpsSettings)[vs[1].(int)]
	}).(IpsSettingsOutput)
}

type IpsSettingsMapOutput struct{ *pulumi.OutputState }

func (IpsSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsSettings)(nil)).Elem()
}

func (o IpsSettingsMapOutput) ToIpsSettingsMapOutput() IpsSettingsMapOutput {
	return o
}

func (o IpsSettingsMapOutput) ToIpsSettingsMapOutputWithContext(ctx context.Context) IpsSettingsMapOutput {
	return o
}

func (o IpsSettingsMapOutput) MapIndex(k pulumi.StringInput) IpsSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsSettings {
		return vs[0].(map[string]*IpsSettings)[vs[1].(string)]
	}).(IpsSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSettingsInput)(nil)).Elem(), &IpsSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSettingsArrayInput)(nil)).Elem(), IpsSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSettingsMapInput)(nil)).Elem(), IpsSettingsMap{})
	pulumi.RegisterOutputType(IpsSettingsOutput{})
	pulumi.RegisterOutputType(IpsSettingsArrayOutput{})
	pulumi.RegisterOutputType(IpsSettingsMapOutput{})
}
