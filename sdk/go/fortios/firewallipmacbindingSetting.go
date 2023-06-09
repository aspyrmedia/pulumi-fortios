// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP to MAC binding settings.
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
//			_, err := fortios.NewFirewallipmacbindingSetting(ctx, "trname", &fortios.FirewallipmacbindingSettingArgs{
//				Bindthroughfw: pulumi.String("disable"),
//				Bindtofw:      pulumi.String("disable"),
//				Undefinedhost: pulumi.String("block"),
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
// # FirewallIpmacbinding Setting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallipmacbindingSetting:FirewallipmacbindingSetting labelname FirewallIpmacbindingSetting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallipmacbindingSetting:FirewallipmacbindingSetting labelname FirewallIpmacbindingSetting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallipmacbindingSetting struct {
	pulumi.CustomResourceState

	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw pulumi.StringOutput `pulumi:"bindthroughfw"`
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw pulumi.StringOutput `pulumi:"bindtofw"`
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost pulumi.StringOutput `pulumi:"undefinedhost"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallipmacbindingSetting registers a new resource with the given unique name, arguments, and options.
func NewFirewallipmacbindingSetting(ctx *pulumi.Context,
	name string, args *FirewallipmacbindingSettingArgs, opts ...pulumi.ResourceOption) (*FirewallipmacbindingSetting, error) {
	if args == nil {
		args = &FirewallipmacbindingSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallipmacbindingSetting
	err := ctx.RegisterResource("fortios:index/firewallipmacbindingSetting:FirewallipmacbindingSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallipmacbindingSetting gets an existing FirewallipmacbindingSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallipmacbindingSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallipmacbindingSettingState, opts ...pulumi.ResourceOption) (*FirewallipmacbindingSetting, error) {
	var resource FirewallipmacbindingSetting
	err := ctx.ReadResource("fortios:index/firewallipmacbindingSetting:FirewallipmacbindingSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallipmacbindingSetting resources.
type firewallipmacbindingSettingState struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw *string `pulumi:"bindthroughfw"`
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw *string `pulumi:"bindtofw"`
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost *string `pulumi:"undefinedhost"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallipmacbindingSettingState struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw pulumi.StringPtrInput
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw pulumi.StringPtrInput
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallipmacbindingSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallipmacbindingSettingState)(nil)).Elem()
}

type firewallipmacbindingSettingArgs struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw *string `pulumi:"bindthroughfw"`
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw *string `pulumi:"bindtofw"`
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost *string `pulumi:"undefinedhost"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallipmacbindingSetting resource.
type FirewallipmacbindingSettingArgs struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw pulumi.StringPtrInput
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw pulumi.StringPtrInput
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallipmacbindingSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallipmacbindingSettingArgs)(nil)).Elem()
}

type FirewallipmacbindingSettingInput interface {
	pulumi.Input

	ToFirewallipmacbindingSettingOutput() FirewallipmacbindingSettingOutput
	ToFirewallipmacbindingSettingOutputWithContext(ctx context.Context) FirewallipmacbindingSettingOutput
}

func (*FirewallipmacbindingSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallipmacbindingSetting)(nil)).Elem()
}

func (i *FirewallipmacbindingSetting) ToFirewallipmacbindingSettingOutput() FirewallipmacbindingSettingOutput {
	return i.ToFirewallipmacbindingSettingOutputWithContext(context.Background())
}

func (i *FirewallipmacbindingSetting) ToFirewallipmacbindingSettingOutputWithContext(ctx context.Context) FirewallipmacbindingSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallipmacbindingSettingOutput)
}

// FirewallipmacbindingSettingArrayInput is an input type that accepts FirewallipmacbindingSettingArray and FirewallipmacbindingSettingArrayOutput values.
// You can construct a concrete instance of `FirewallipmacbindingSettingArrayInput` via:
//
//	FirewallipmacbindingSettingArray{ FirewallipmacbindingSettingArgs{...} }
type FirewallipmacbindingSettingArrayInput interface {
	pulumi.Input

	ToFirewallipmacbindingSettingArrayOutput() FirewallipmacbindingSettingArrayOutput
	ToFirewallipmacbindingSettingArrayOutputWithContext(context.Context) FirewallipmacbindingSettingArrayOutput
}

type FirewallipmacbindingSettingArray []FirewallipmacbindingSettingInput

func (FirewallipmacbindingSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallipmacbindingSetting)(nil)).Elem()
}

func (i FirewallipmacbindingSettingArray) ToFirewallipmacbindingSettingArrayOutput() FirewallipmacbindingSettingArrayOutput {
	return i.ToFirewallipmacbindingSettingArrayOutputWithContext(context.Background())
}

func (i FirewallipmacbindingSettingArray) ToFirewallipmacbindingSettingArrayOutputWithContext(ctx context.Context) FirewallipmacbindingSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallipmacbindingSettingArrayOutput)
}

// FirewallipmacbindingSettingMapInput is an input type that accepts FirewallipmacbindingSettingMap and FirewallipmacbindingSettingMapOutput values.
// You can construct a concrete instance of `FirewallipmacbindingSettingMapInput` via:
//
//	FirewallipmacbindingSettingMap{ "key": FirewallipmacbindingSettingArgs{...} }
type FirewallipmacbindingSettingMapInput interface {
	pulumi.Input

	ToFirewallipmacbindingSettingMapOutput() FirewallipmacbindingSettingMapOutput
	ToFirewallipmacbindingSettingMapOutputWithContext(context.Context) FirewallipmacbindingSettingMapOutput
}

type FirewallipmacbindingSettingMap map[string]FirewallipmacbindingSettingInput

func (FirewallipmacbindingSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallipmacbindingSetting)(nil)).Elem()
}

func (i FirewallipmacbindingSettingMap) ToFirewallipmacbindingSettingMapOutput() FirewallipmacbindingSettingMapOutput {
	return i.ToFirewallipmacbindingSettingMapOutputWithContext(context.Background())
}

func (i FirewallipmacbindingSettingMap) ToFirewallipmacbindingSettingMapOutputWithContext(ctx context.Context) FirewallipmacbindingSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallipmacbindingSettingMapOutput)
}

type FirewallipmacbindingSettingOutput struct{ *pulumi.OutputState }

func (FirewallipmacbindingSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallipmacbindingSetting)(nil)).Elem()
}

func (o FirewallipmacbindingSettingOutput) ToFirewallipmacbindingSettingOutput() FirewallipmacbindingSettingOutput {
	return o
}

func (o FirewallipmacbindingSettingOutput) ToFirewallipmacbindingSettingOutputWithContext(ctx context.Context) FirewallipmacbindingSettingOutput {
	return o
}

// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
func (o FirewallipmacbindingSettingOutput) Bindthroughfw() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingSetting) pulumi.StringOutput { return v.Bindthroughfw }).(pulumi.StringOutput)
}

// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
func (o FirewallipmacbindingSettingOutput) Bindtofw() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingSetting) pulumi.StringOutput { return v.Bindtofw }).(pulumi.StringOutput)
}

// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
func (o FirewallipmacbindingSettingOutput) Undefinedhost() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingSetting) pulumi.StringOutput { return v.Undefinedhost }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallipmacbindingSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallipmacbindingSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallipmacbindingSettingArrayOutput struct{ *pulumi.OutputState }

func (FirewallipmacbindingSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallipmacbindingSetting)(nil)).Elem()
}

func (o FirewallipmacbindingSettingArrayOutput) ToFirewallipmacbindingSettingArrayOutput() FirewallipmacbindingSettingArrayOutput {
	return o
}

func (o FirewallipmacbindingSettingArrayOutput) ToFirewallipmacbindingSettingArrayOutputWithContext(ctx context.Context) FirewallipmacbindingSettingArrayOutput {
	return o
}

func (o FirewallipmacbindingSettingArrayOutput) Index(i pulumi.IntInput) FirewallipmacbindingSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallipmacbindingSetting {
		return vs[0].([]*FirewallipmacbindingSetting)[vs[1].(int)]
	}).(FirewallipmacbindingSettingOutput)
}

type FirewallipmacbindingSettingMapOutput struct{ *pulumi.OutputState }

func (FirewallipmacbindingSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallipmacbindingSetting)(nil)).Elem()
}

func (o FirewallipmacbindingSettingMapOutput) ToFirewallipmacbindingSettingMapOutput() FirewallipmacbindingSettingMapOutput {
	return o
}

func (o FirewallipmacbindingSettingMapOutput) ToFirewallipmacbindingSettingMapOutputWithContext(ctx context.Context) FirewallipmacbindingSettingMapOutput {
	return o
}

func (o FirewallipmacbindingSettingMapOutput) MapIndex(k pulumi.StringInput) FirewallipmacbindingSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallipmacbindingSetting {
		return vs[0].(map[string]*FirewallipmacbindingSetting)[vs[1].(string)]
	}).(FirewallipmacbindingSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallipmacbindingSettingInput)(nil)).Elem(), &FirewallipmacbindingSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallipmacbindingSettingArrayInput)(nil)).Elem(), FirewallipmacbindingSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallipmacbindingSettingMapInput)(nil)).Elem(), FirewallipmacbindingSettingMap{})
	pulumi.RegisterOutputType(FirewallipmacbindingSettingOutput{})
	pulumi.RegisterOutputType(FirewallipmacbindingSettingArrayOutput{})
	pulumi.RegisterOutputType(FirewallipmacbindingSettingMapOutput{})
}
