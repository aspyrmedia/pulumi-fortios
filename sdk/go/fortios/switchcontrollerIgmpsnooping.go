// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch IGMP snooping global settings.
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
//			_, err := fortios.NewSwitchcontrollerIgmpsnooping(ctx, "trname", &fortios.SwitchcontrollerIgmpsnoopingArgs{
//				AgingTime:             pulumi.Int(300),
//				FloodUnknownMulticast: pulumi.String("disable"),
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
// # SwitchController IgmpSnooping can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerIgmpsnooping:SwitchcontrollerIgmpsnooping labelname SwitchControllerIgmpSnooping
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerIgmpsnooping:SwitchcontrollerIgmpsnooping labelname SwitchControllerIgmpSnooping
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerIgmpsnooping struct {
	pulumi.CustomResourceState

	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime pulumi.IntOutput `pulumi:"agingTime"`
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast pulumi.StringOutput `pulumi:"floodUnknownMulticast"`
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval pulumi.IntOutput `pulumi:"queryInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerIgmpsnooping registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerIgmpsnooping(ctx *pulumi.Context,
	name string, args *SwitchcontrollerIgmpsnoopingArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerIgmpsnooping, error) {
	if args == nil {
		args = &SwitchcontrollerIgmpsnoopingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerIgmpsnooping
	err := ctx.RegisterResource("fortios:index/switchcontrollerIgmpsnooping:SwitchcontrollerIgmpsnooping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerIgmpsnooping gets an existing SwitchcontrollerIgmpsnooping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerIgmpsnooping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerIgmpsnoopingState, opts ...pulumi.ResourceOption) (*SwitchcontrollerIgmpsnooping, error) {
	var resource SwitchcontrollerIgmpsnooping
	err := ctx.ReadResource("fortios:index/switchcontrollerIgmpsnooping:SwitchcontrollerIgmpsnooping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerIgmpsnooping resources.
type switchcontrollerIgmpsnoopingState struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime *int `pulumi:"agingTime"`
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast *string `pulumi:"floodUnknownMulticast"`
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval *int `pulumi:"queryInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerIgmpsnoopingState struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime pulumi.IntPtrInput
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast pulumi.StringPtrInput
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerIgmpsnoopingState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerIgmpsnoopingState)(nil)).Elem()
}

type switchcontrollerIgmpsnoopingArgs struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime *int `pulumi:"agingTime"`
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast *string `pulumi:"floodUnknownMulticast"`
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval *int `pulumi:"queryInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerIgmpsnooping resource.
type SwitchcontrollerIgmpsnoopingArgs struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime pulumi.IntPtrInput
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast pulumi.StringPtrInput
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerIgmpsnoopingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerIgmpsnoopingArgs)(nil)).Elem()
}

type SwitchcontrollerIgmpsnoopingInput interface {
	pulumi.Input

	ToSwitchcontrollerIgmpsnoopingOutput() SwitchcontrollerIgmpsnoopingOutput
	ToSwitchcontrollerIgmpsnoopingOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingOutput
}

func (*SwitchcontrollerIgmpsnooping) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerIgmpsnooping)(nil)).Elem()
}

func (i *SwitchcontrollerIgmpsnooping) ToSwitchcontrollerIgmpsnoopingOutput() SwitchcontrollerIgmpsnoopingOutput {
	return i.ToSwitchcontrollerIgmpsnoopingOutputWithContext(context.Background())
}

func (i *SwitchcontrollerIgmpsnooping) ToSwitchcontrollerIgmpsnoopingOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerIgmpsnoopingOutput)
}

// SwitchcontrollerIgmpsnoopingArrayInput is an input type that accepts SwitchcontrollerIgmpsnoopingArray and SwitchcontrollerIgmpsnoopingArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerIgmpsnoopingArrayInput` via:
//
//	SwitchcontrollerIgmpsnoopingArray{ SwitchcontrollerIgmpsnoopingArgs{...} }
type SwitchcontrollerIgmpsnoopingArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerIgmpsnoopingArrayOutput() SwitchcontrollerIgmpsnoopingArrayOutput
	ToSwitchcontrollerIgmpsnoopingArrayOutputWithContext(context.Context) SwitchcontrollerIgmpsnoopingArrayOutput
}

type SwitchcontrollerIgmpsnoopingArray []SwitchcontrollerIgmpsnoopingInput

func (SwitchcontrollerIgmpsnoopingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerIgmpsnooping)(nil)).Elem()
}

func (i SwitchcontrollerIgmpsnoopingArray) ToSwitchcontrollerIgmpsnoopingArrayOutput() SwitchcontrollerIgmpsnoopingArrayOutput {
	return i.ToSwitchcontrollerIgmpsnoopingArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerIgmpsnoopingArray) ToSwitchcontrollerIgmpsnoopingArrayOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerIgmpsnoopingArrayOutput)
}

// SwitchcontrollerIgmpsnoopingMapInput is an input type that accepts SwitchcontrollerIgmpsnoopingMap and SwitchcontrollerIgmpsnoopingMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerIgmpsnoopingMapInput` via:
//
//	SwitchcontrollerIgmpsnoopingMap{ "key": SwitchcontrollerIgmpsnoopingArgs{...} }
type SwitchcontrollerIgmpsnoopingMapInput interface {
	pulumi.Input

	ToSwitchcontrollerIgmpsnoopingMapOutput() SwitchcontrollerIgmpsnoopingMapOutput
	ToSwitchcontrollerIgmpsnoopingMapOutputWithContext(context.Context) SwitchcontrollerIgmpsnoopingMapOutput
}

type SwitchcontrollerIgmpsnoopingMap map[string]SwitchcontrollerIgmpsnoopingInput

func (SwitchcontrollerIgmpsnoopingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerIgmpsnooping)(nil)).Elem()
}

func (i SwitchcontrollerIgmpsnoopingMap) ToSwitchcontrollerIgmpsnoopingMapOutput() SwitchcontrollerIgmpsnoopingMapOutput {
	return i.ToSwitchcontrollerIgmpsnoopingMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerIgmpsnoopingMap) ToSwitchcontrollerIgmpsnoopingMapOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerIgmpsnoopingMapOutput)
}

type SwitchcontrollerIgmpsnoopingOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerIgmpsnoopingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerIgmpsnooping)(nil)).Elem()
}

func (o SwitchcontrollerIgmpsnoopingOutput) ToSwitchcontrollerIgmpsnoopingOutput() SwitchcontrollerIgmpsnoopingOutput {
	return o
}

func (o SwitchcontrollerIgmpsnoopingOutput) ToSwitchcontrollerIgmpsnoopingOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingOutput {
	return o
}

// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
func (o SwitchcontrollerIgmpsnoopingOutput) AgingTime() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerIgmpsnooping) pulumi.IntOutput { return v.AgingTime }).(pulumi.IntOutput)
}

// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
func (o SwitchcontrollerIgmpsnoopingOutput) FloodUnknownMulticast() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerIgmpsnooping) pulumi.StringOutput { return v.FloodUnknownMulticast }).(pulumi.StringOutput)
}

// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
func (o SwitchcontrollerIgmpsnoopingOutput) QueryInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerIgmpsnooping) pulumi.IntOutput { return v.QueryInterval }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerIgmpsnoopingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerIgmpsnooping) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerIgmpsnoopingArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerIgmpsnoopingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerIgmpsnooping)(nil)).Elem()
}

func (o SwitchcontrollerIgmpsnoopingArrayOutput) ToSwitchcontrollerIgmpsnoopingArrayOutput() SwitchcontrollerIgmpsnoopingArrayOutput {
	return o
}

func (o SwitchcontrollerIgmpsnoopingArrayOutput) ToSwitchcontrollerIgmpsnoopingArrayOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingArrayOutput {
	return o
}

func (o SwitchcontrollerIgmpsnoopingArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerIgmpsnoopingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerIgmpsnooping {
		return vs[0].([]*SwitchcontrollerIgmpsnooping)[vs[1].(int)]
	}).(SwitchcontrollerIgmpsnoopingOutput)
}

type SwitchcontrollerIgmpsnoopingMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerIgmpsnoopingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerIgmpsnooping)(nil)).Elem()
}

func (o SwitchcontrollerIgmpsnoopingMapOutput) ToSwitchcontrollerIgmpsnoopingMapOutput() SwitchcontrollerIgmpsnoopingMapOutput {
	return o
}

func (o SwitchcontrollerIgmpsnoopingMapOutput) ToSwitchcontrollerIgmpsnoopingMapOutputWithContext(ctx context.Context) SwitchcontrollerIgmpsnoopingMapOutput {
	return o
}

func (o SwitchcontrollerIgmpsnoopingMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerIgmpsnoopingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerIgmpsnooping {
		return vs[0].(map[string]*SwitchcontrollerIgmpsnooping)[vs[1].(string)]
	}).(SwitchcontrollerIgmpsnoopingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerIgmpsnoopingInput)(nil)).Elem(), &SwitchcontrollerIgmpsnooping{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerIgmpsnoopingArrayInput)(nil)).Elem(), SwitchcontrollerIgmpsnoopingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerIgmpsnoopingMapInput)(nil)).Elem(), SwitchcontrollerIgmpsnoopingMap{})
	pulumi.RegisterOutputType(SwitchcontrollerIgmpsnoopingOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerIgmpsnoopingArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerIgmpsnoopingMapOutput{})
}
