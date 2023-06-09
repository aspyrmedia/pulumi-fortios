// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure system-wide switch controller settings.
//
// ## Import
//
// # SwitchController System can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerSystem:SwitchcontrollerSystem labelname SwitchControllerSystem
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerSystem:SwitchcontrollerSystem labelname SwitchControllerSystem
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerSystem struct {
	pulumi.CustomResourceState

	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval pulumi.IntOutput `pulumi:"dataSyncInterval"`
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval pulumi.IntOutput `pulumi:"dynamicPeriodicInterval"`
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff pulumi.IntOutput `pulumi:"iotHoldoff"`
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle pulumi.IntOutput `pulumi:"iotMacIdle"`
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval pulumi.IntOutput `pulumi:"iotScanInterval"`
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold pulumi.IntOutput `pulumi:"iotWeightThreshold"`
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval pulumi.IntOutput `pulumi:"nacPeriodicInterval"`
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess pulumi.IntOutput `pulumi:"parallelProcess"`
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride pulumi.StringOutput `pulumi:"parallelProcessOverride"`
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode pulumi.StringOutput `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerSystem registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerSystem(ctx *pulumi.Context,
	name string, args *SwitchcontrollerSystemArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerSystem, error) {
	if args == nil {
		args = &SwitchcontrollerSystemArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerSystem
	err := ctx.RegisterResource("fortios:index/switchcontrollerSystem:SwitchcontrollerSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerSystem gets an existing SwitchcontrollerSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerSystemState, opts ...pulumi.ResourceOption) (*SwitchcontrollerSystem, error) {
	var resource SwitchcontrollerSystem
	err := ctx.ReadResource("fortios:index/switchcontrollerSystem:SwitchcontrollerSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerSystem resources.
type switchcontrollerSystemState struct {
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval *int `pulumi:"dataSyncInterval"`
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval *int `pulumi:"dynamicPeriodicInterval"`
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff *int `pulumi:"iotHoldoff"`
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle *int `pulumi:"iotMacIdle"`
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval *int `pulumi:"iotScanInterval"`
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold *int `pulumi:"iotWeightThreshold"`
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval *int `pulumi:"nacPeriodicInterval"`
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess *int `pulumi:"parallelProcess"`
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride *string `pulumi:"parallelProcessOverride"`
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerSystemState struct {
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval pulumi.IntPtrInput
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval pulumi.IntPtrInput
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff pulumi.IntPtrInput
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle pulumi.IntPtrInput
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval pulumi.IntPtrInput
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold pulumi.IntPtrInput
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval pulumi.IntPtrInput
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess pulumi.IntPtrInput
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride pulumi.StringPtrInput
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerSystemState)(nil)).Elem()
}

type switchcontrollerSystemArgs struct {
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval *int `pulumi:"dataSyncInterval"`
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval *int `pulumi:"dynamicPeriodicInterval"`
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff *int `pulumi:"iotHoldoff"`
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle *int `pulumi:"iotMacIdle"`
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval *int `pulumi:"iotScanInterval"`
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold *int `pulumi:"iotWeightThreshold"`
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval *int `pulumi:"nacPeriodicInterval"`
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess *int `pulumi:"parallelProcess"`
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride *string `pulumi:"parallelProcessOverride"`
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerSystem resource.
type SwitchcontrollerSystemArgs struct {
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval pulumi.IntPtrInput
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval pulumi.IntPtrInput
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff pulumi.IntPtrInput
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle pulumi.IntPtrInput
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval pulumi.IntPtrInput
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold pulumi.IntPtrInput
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval pulumi.IntPtrInput
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess pulumi.IntPtrInput
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride pulumi.StringPtrInput
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerSystemArgs)(nil)).Elem()
}

type SwitchcontrollerSystemInput interface {
	pulumi.Input

	ToSwitchcontrollerSystemOutput() SwitchcontrollerSystemOutput
	ToSwitchcontrollerSystemOutputWithContext(ctx context.Context) SwitchcontrollerSystemOutput
}

func (*SwitchcontrollerSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerSystem)(nil)).Elem()
}

func (i *SwitchcontrollerSystem) ToSwitchcontrollerSystemOutput() SwitchcontrollerSystemOutput {
	return i.ToSwitchcontrollerSystemOutputWithContext(context.Background())
}

func (i *SwitchcontrollerSystem) ToSwitchcontrollerSystemOutputWithContext(ctx context.Context) SwitchcontrollerSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerSystemOutput)
}

// SwitchcontrollerSystemArrayInput is an input type that accepts SwitchcontrollerSystemArray and SwitchcontrollerSystemArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerSystemArrayInput` via:
//
//	SwitchcontrollerSystemArray{ SwitchcontrollerSystemArgs{...} }
type SwitchcontrollerSystemArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerSystemArrayOutput() SwitchcontrollerSystemArrayOutput
	ToSwitchcontrollerSystemArrayOutputWithContext(context.Context) SwitchcontrollerSystemArrayOutput
}

type SwitchcontrollerSystemArray []SwitchcontrollerSystemInput

func (SwitchcontrollerSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerSystem)(nil)).Elem()
}

func (i SwitchcontrollerSystemArray) ToSwitchcontrollerSystemArrayOutput() SwitchcontrollerSystemArrayOutput {
	return i.ToSwitchcontrollerSystemArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerSystemArray) ToSwitchcontrollerSystemArrayOutputWithContext(ctx context.Context) SwitchcontrollerSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerSystemArrayOutput)
}

// SwitchcontrollerSystemMapInput is an input type that accepts SwitchcontrollerSystemMap and SwitchcontrollerSystemMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerSystemMapInput` via:
//
//	SwitchcontrollerSystemMap{ "key": SwitchcontrollerSystemArgs{...} }
type SwitchcontrollerSystemMapInput interface {
	pulumi.Input

	ToSwitchcontrollerSystemMapOutput() SwitchcontrollerSystemMapOutput
	ToSwitchcontrollerSystemMapOutputWithContext(context.Context) SwitchcontrollerSystemMapOutput
}

type SwitchcontrollerSystemMap map[string]SwitchcontrollerSystemInput

func (SwitchcontrollerSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerSystem)(nil)).Elem()
}

func (i SwitchcontrollerSystemMap) ToSwitchcontrollerSystemMapOutput() SwitchcontrollerSystemMapOutput {
	return i.ToSwitchcontrollerSystemMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerSystemMap) ToSwitchcontrollerSystemMapOutputWithContext(ctx context.Context) SwitchcontrollerSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerSystemMapOutput)
}

type SwitchcontrollerSystemOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerSystem)(nil)).Elem()
}

func (o SwitchcontrollerSystemOutput) ToSwitchcontrollerSystemOutput() SwitchcontrollerSystemOutput {
	return o
}

func (o SwitchcontrollerSystemOutput) ToSwitchcontrollerSystemOutputWithContext(ctx context.Context) SwitchcontrollerSystemOutput {
	return o
}

// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
func (o SwitchcontrollerSystemOutput) DataSyncInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.DataSyncInterval }).(pulumi.IntOutput)
}

// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
func (o SwitchcontrollerSystemOutput) DynamicPeriodicInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.DynamicPeriodicInterval }).(pulumi.IntOutput)
}

// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
func (o SwitchcontrollerSystemOutput) IotHoldoff() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.IotHoldoff }).(pulumi.IntOutput)
}

// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
func (o SwitchcontrollerSystemOutput) IotMacIdle() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.IotMacIdle }).(pulumi.IntOutput)
}

// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
func (o SwitchcontrollerSystemOutput) IotScanInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.IotScanInterval }).(pulumi.IntOutput)
}

// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
func (o SwitchcontrollerSystemOutput) IotWeightThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.IotWeightThreshold }).(pulumi.IntOutput)
}

// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
func (o SwitchcontrollerSystemOutput) NacPeriodicInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.NacPeriodicInterval }).(pulumi.IntOutput)
}

// Maximum number of parallel processes (1 - 300, default = 1).
func (o SwitchcontrollerSystemOutput) ParallelProcess() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.IntOutput { return v.ParallelProcess }).(pulumi.IntOutput)
}

// Enable/disable parallel process override. Valid values: `disable`, `enable`.
func (o SwitchcontrollerSystemOutput) ParallelProcessOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.StringOutput { return v.ParallelProcessOverride }).(pulumi.StringOutput)
}

// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
func (o SwitchcontrollerSystemOutput) TunnelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.StringOutput { return v.TunnelMode }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerSystemOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerSystem) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerSystemArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerSystem)(nil)).Elem()
}

func (o SwitchcontrollerSystemArrayOutput) ToSwitchcontrollerSystemArrayOutput() SwitchcontrollerSystemArrayOutput {
	return o
}

func (o SwitchcontrollerSystemArrayOutput) ToSwitchcontrollerSystemArrayOutputWithContext(ctx context.Context) SwitchcontrollerSystemArrayOutput {
	return o
}

func (o SwitchcontrollerSystemArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerSystem {
		return vs[0].([]*SwitchcontrollerSystem)[vs[1].(int)]
	}).(SwitchcontrollerSystemOutput)
}

type SwitchcontrollerSystemMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerSystem)(nil)).Elem()
}

func (o SwitchcontrollerSystemMapOutput) ToSwitchcontrollerSystemMapOutput() SwitchcontrollerSystemMapOutput {
	return o
}

func (o SwitchcontrollerSystemMapOutput) ToSwitchcontrollerSystemMapOutputWithContext(ctx context.Context) SwitchcontrollerSystemMapOutput {
	return o
}

func (o SwitchcontrollerSystemMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerSystem {
		return vs[0].(map[string]*SwitchcontrollerSystem)[vs[1].(string)]
	}).(SwitchcontrollerSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerSystemInput)(nil)).Elem(), &SwitchcontrollerSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerSystemArrayInput)(nil)).Elem(), SwitchcontrollerSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerSystemMapInput)(nil)).Elem(), SwitchcontrollerSystemMap{})
	pulumi.RegisterOutputType(SwitchcontrollerSystemOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerSystemArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerSystemMapOutput{})
}
