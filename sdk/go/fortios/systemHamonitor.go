// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure HA monitor.
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
//			_, err := fortios.NewSystemHamonitor(ctx, "trname", &fortios.SystemHamonitorArgs{
//				MonitorVlan:         pulumi.String("disable"),
//				VlanHbInterval:      pulumi.Int(5),
//				VlanHbLostThreshold: pulumi.Int(3),
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
// # System HaMonitor can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemHamonitor:SystemHamonitor labelname SystemHaMonitor
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemHamonitor:SystemHamonitor labelname SystemHaMonitor
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemHamonitor struct {
	pulumi.CustomResourceState

	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan pulumi.StringOutput `pulumi:"monitorVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure heartbeat interval (seconds).
	VlanHbInterval pulumi.IntOutput `pulumi:"vlanHbInterval"`
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold pulumi.IntOutput `pulumi:"vlanHbLostThreshold"`
}

// NewSystemHamonitor registers a new resource with the given unique name, arguments, and options.
func NewSystemHamonitor(ctx *pulumi.Context,
	name string, args *SystemHamonitorArgs, opts ...pulumi.ResourceOption) (*SystemHamonitor, error) {
	if args == nil {
		args = &SystemHamonitorArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemHamonitor
	err := ctx.RegisterResource("fortios:index/systemHamonitor:SystemHamonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemHamonitor gets an existing SystemHamonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemHamonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemHamonitorState, opts ...pulumi.ResourceOption) (*SystemHamonitor, error) {
	var resource SystemHamonitor
	err := ctx.ReadResource("fortios:index/systemHamonitor:SystemHamonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemHamonitor resources.
type systemHamonitorState struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan *string `pulumi:"monitorVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure heartbeat interval (seconds).
	VlanHbInterval *int `pulumi:"vlanHbInterval"`
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold *int `pulumi:"vlanHbLostThreshold"`
}

type SystemHamonitorState struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure heartbeat interval (seconds).
	VlanHbInterval pulumi.IntPtrInput
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold pulumi.IntPtrInput
}

func (SystemHamonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemHamonitorState)(nil)).Elem()
}

type systemHamonitorArgs struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan *string `pulumi:"monitorVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure heartbeat interval (seconds).
	VlanHbInterval *int `pulumi:"vlanHbInterval"`
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold *int `pulumi:"vlanHbLostThreshold"`
}

// The set of arguments for constructing a SystemHamonitor resource.
type SystemHamonitorArgs struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure heartbeat interval (seconds).
	VlanHbInterval pulumi.IntPtrInput
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold pulumi.IntPtrInput
}

func (SystemHamonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemHamonitorArgs)(nil)).Elem()
}

type SystemHamonitorInput interface {
	pulumi.Input

	ToSystemHamonitorOutput() SystemHamonitorOutput
	ToSystemHamonitorOutputWithContext(ctx context.Context) SystemHamonitorOutput
}

func (*SystemHamonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemHamonitor)(nil)).Elem()
}

func (i *SystemHamonitor) ToSystemHamonitorOutput() SystemHamonitorOutput {
	return i.ToSystemHamonitorOutputWithContext(context.Background())
}

func (i *SystemHamonitor) ToSystemHamonitorOutputWithContext(ctx context.Context) SystemHamonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemHamonitorOutput)
}

// SystemHamonitorArrayInput is an input type that accepts SystemHamonitorArray and SystemHamonitorArrayOutput values.
// You can construct a concrete instance of `SystemHamonitorArrayInput` via:
//
//	SystemHamonitorArray{ SystemHamonitorArgs{...} }
type SystemHamonitorArrayInput interface {
	pulumi.Input

	ToSystemHamonitorArrayOutput() SystemHamonitorArrayOutput
	ToSystemHamonitorArrayOutputWithContext(context.Context) SystemHamonitorArrayOutput
}

type SystemHamonitorArray []SystemHamonitorInput

func (SystemHamonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemHamonitor)(nil)).Elem()
}

func (i SystemHamonitorArray) ToSystemHamonitorArrayOutput() SystemHamonitorArrayOutput {
	return i.ToSystemHamonitorArrayOutputWithContext(context.Background())
}

func (i SystemHamonitorArray) ToSystemHamonitorArrayOutputWithContext(ctx context.Context) SystemHamonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemHamonitorArrayOutput)
}

// SystemHamonitorMapInput is an input type that accepts SystemHamonitorMap and SystemHamonitorMapOutput values.
// You can construct a concrete instance of `SystemHamonitorMapInput` via:
//
//	SystemHamonitorMap{ "key": SystemHamonitorArgs{...} }
type SystemHamonitorMapInput interface {
	pulumi.Input

	ToSystemHamonitorMapOutput() SystemHamonitorMapOutput
	ToSystemHamonitorMapOutputWithContext(context.Context) SystemHamonitorMapOutput
}

type SystemHamonitorMap map[string]SystemHamonitorInput

func (SystemHamonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemHamonitor)(nil)).Elem()
}

func (i SystemHamonitorMap) ToSystemHamonitorMapOutput() SystemHamonitorMapOutput {
	return i.ToSystemHamonitorMapOutputWithContext(context.Background())
}

func (i SystemHamonitorMap) ToSystemHamonitorMapOutputWithContext(ctx context.Context) SystemHamonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemHamonitorMapOutput)
}

type SystemHamonitorOutput struct{ *pulumi.OutputState }

func (SystemHamonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemHamonitor)(nil)).Elem()
}

func (o SystemHamonitorOutput) ToSystemHamonitorOutput() SystemHamonitorOutput {
	return o
}

func (o SystemHamonitorOutput) ToSystemHamonitorOutputWithContext(ctx context.Context) SystemHamonitorOutput {
	return o
}

// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
func (o SystemHamonitorOutput) MonitorVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemHamonitor) pulumi.StringOutput { return v.MonitorVlan }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemHamonitorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemHamonitor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Configure heartbeat interval (seconds).
func (o SystemHamonitorOutput) VlanHbInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemHamonitor) pulumi.IntOutput { return v.VlanHbInterval }).(pulumi.IntOutput)
}

// VLAN lost heartbeat threshold (1 - 60).
func (o SystemHamonitorOutput) VlanHbLostThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemHamonitor) pulumi.IntOutput { return v.VlanHbLostThreshold }).(pulumi.IntOutput)
}

type SystemHamonitorArrayOutput struct{ *pulumi.OutputState }

func (SystemHamonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemHamonitor)(nil)).Elem()
}

func (o SystemHamonitorArrayOutput) ToSystemHamonitorArrayOutput() SystemHamonitorArrayOutput {
	return o
}

func (o SystemHamonitorArrayOutput) ToSystemHamonitorArrayOutputWithContext(ctx context.Context) SystemHamonitorArrayOutput {
	return o
}

func (o SystemHamonitorArrayOutput) Index(i pulumi.IntInput) SystemHamonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemHamonitor {
		return vs[0].([]*SystemHamonitor)[vs[1].(int)]
	}).(SystemHamonitorOutput)
}

type SystemHamonitorMapOutput struct{ *pulumi.OutputState }

func (SystemHamonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemHamonitor)(nil)).Elem()
}

func (o SystemHamonitorMapOutput) ToSystemHamonitorMapOutput() SystemHamonitorMapOutput {
	return o
}

func (o SystemHamonitorMapOutput) ToSystemHamonitorMapOutputWithContext(ctx context.Context) SystemHamonitorMapOutput {
	return o
}

func (o SystemHamonitorMapOutput) MapIndex(k pulumi.StringInput) SystemHamonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemHamonitor {
		return vs[0].(map[string]*SystemHamonitor)[vs[1].(string)]
	}).(SystemHamonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemHamonitorInput)(nil)).Elem(), &SystemHamonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemHamonitorArrayInput)(nil)).Elem(), SystemHamonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemHamonitorMapInput)(nil)).Elem(), SystemHamonitorMap{})
	pulumi.RegisterOutputType(SystemHamonitorOutput{})
	pulumi.RegisterOutputType(SystemHamonitorArrayOutput{})
	pulumi.RegisterOutputType(SystemHamonitorMapOutput{})
}