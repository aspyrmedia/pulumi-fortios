// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DSCP based priority table.
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
//			_, err := fortios.NewSystemDscpbasedpriority(ctx, "trname", &fortios.SystemDscpbasedpriorityArgs{
//				Ds:       pulumi.Int(1),
//				Fosid:    pulumi.Int(1),
//				Priority: pulumi.String("low"),
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
// # System DscpBasedPriority can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemDscpbasedpriority struct {
	pulumi.CustomResourceState

	// DSCP(DiffServ) DS value (0 - 63).
	Ds pulumi.IntOutput `pulumi:"ds"`
	// Item ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// DSCP based priority level. Valid values: `low`, `medium`, `high`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemDscpbasedpriority registers a new resource with the given unique name, arguments, and options.
func NewSystemDscpbasedpriority(ctx *pulumi.Context,
	name string, args *SystemDscpbasedpriorityArgs, opts ...pulumi.ResourceOption) (*SystemDscpbasedpriority, error) {
	if args == nil {
		args = &SystemDscpbasedpriorityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemDscpbasedpriority
	err := ctx.RegisterResource("fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDscpbasedpriority gets an existing SystemDscpbasedpriority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDscpbasedpriority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDscpbasedpriorityState, opts ...pulumi.ResourceOption) (*SystemDscpbasedpriority, error) {
	var resource SystemDscpbasedpriority
	err := ctx.ReadResource("fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDscpbasedpriority resources.
type systemDscpbasedpriorityState struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds *int `pulumi:"ds"`
	// Item ID.
	Fosid *int `pulumi:"fosid"`
	// DSCP based priority level. Valid values: `low`, `medium`, `high`.
	Priority *string `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemDscpbasedpriorityState struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds pulumi.IntPtrInput
	// Item ID.
	Fosid pulumi.IntPtrInput
	// DSCP based priority level. Valid values: `low`, `medium`, `high`.
	Priority pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDscpbasedpriorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDscpbasedpriorityState)(nil)).Elem()
}

type systemDscpbasedpriorityArgs struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds *int `pulumi:"ds"`
	// Item ID.
	Fosid *int `pulumi:"fosid"`
	// DSCP based priority level. Valid values: `low`, `medium`, `high`.
	Priority *string `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemDscpbasedpriority resource.
type SystemDscpbasedpriorityArgs struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds pulumi.IntPtrInput
	// Item ID.
	Fosid pulumi.IntPtrInput
	// DSCP based priority level. Valid values: `low`, `medium`, `high`.
	Priority pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDscpbasedpriorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDscpbasedpriorityArgs)(nil)).Elem()
}

type SystemDscpbasedpriorityInput interface {
	pulumi.Input

	ToSystemDscpbasedpriorityOutput() SystemDscpbasedpriorityOutput
	ToSystemDscpbasedpriorityOutputWithContext(ctx context.Context) SystemDscpbasedpriorityOutput
}

func (*SystemDscpbasedpriority) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDscpbasedpriority)(nil)).Elem()
}

func (i *SystemDscpbasedpriority) ToSystemDscpbasedpriorityOutput() SystemDscpbasedpriorityOutput {
	return i.ToSystemDscpbasedpriorityOutputWithContext(context.Background())
}

func (i *SystemDscpbasedpriority) ToSystemDscpbasedpriorityOutputWithContext(ctx context.Context) SystemDscpbasedpriorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDscpbasedpriorityOutput)
}

// SystemDscpbasedpriorityArrayInput is an input type that accepts SystemDscpbasedpriorityArray and SystemDscpbasedpriorityArrayOutput values.
// You can construct a concrete instance of `SystemDscpbasedpriorityArrayInput` via:
//
//	SystemDscpbasedpriorityArray{ SystemDscpbasedpriorityArgs{...} }
type SystemDscpbasedpriorityArrayInput interface {
	pulumi.Input

	ToSystemDscpbasedpriorityArrayOutput() SystemDscpbasedpriorityArrayOutput
	ToSystemDscpbasedpriorityArrayOutputWithContext(context.Context) SystemDscpbasedpriorityArrayOutput
}

type SystemDscpbasedpriorityArray []SystemDscpbasedpriorityInput

func (SystemDscpbasedpriorityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDscpbasedpriority)(nil)).Elem()
}

func (i SystemDscpbasedpriorityArray) ToSystemDscpbasedpriorityArrayOutput() SystemDscpbasedpriorityArrayOutput {
	return i.ToSystemDscpbasedpriorityArrayOutputWithContext(context.Background())
}

func (i SystemDscpbasedpriorityArray) ToSystemDscpbasedpriorityArrayOutputWithContext(ctx context.Context) SystemDscpbasedpriorityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDscpbasedpriorityArrayOutput)
}

// SystemDscpbasedpriorityMapInput is an input type that accepts SystemDscpbasedpriorityMap and SystemDscpbasedpriorityMapOutput values.
// You can construct a concrete instance of `SystemDscpbasedpriorityMapInput` via:
//
//	SystemDscpbasedpriorityMap{ "key": SystemDscpbasedpriorityArgs{...} }
type SystemDscpbasedpriorityMapInput interface {
	pulumi.Input

	ToSystemDscpbasedpriorityMapOutput() SystemDscpbasedpriorityMapOutput
	ToSystemDscpbasedpriorityMapOutputWithContext(context.Context) SystemDscpbasedpriorityMapOutput
}

type SystemDscpbasedpriorityMap map[string]SystemDscpbasedpriorityInput

func (SystemDscpbasedpriorityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDscpbasedpriority)(nil)).Elem()
}

func (i SystemDscpbasedpriorityMap) ToSystemDscpbasedpriorityMapOutput() SystemDscpbasedpriorityMapOutput {
	return i.ToSystemDscpbasedpriorityMapOutputWithContext(context.Background())
}

func (i SystemDscpbasedpriorityMap) ToSystemDscpbasedpriorityMapOutputWithContext(ctx context.Context) SystemDscpbasedpriorityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDscpbasedpriorityMapOutput)
}

type SystemDscpbasedpriorityOutput struct{ *pulumi.OutputState }

func (SystemDscpbasedpriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDscpbasedpriority)(nil)).Elem()
}

func (o SystemDscpbasedpriorityOutput) ToSystemDscpbasedpriorityOutput() SystemDscpbasedpriorityOutput {
	return o
}

func (o SystemDscpbasedpriorityOutput) ToSystemDscpbasedpriorityOutputWithContext(ctx context.Context) SystemDscpbasedpriorityOutput {
	return o
}

// DSCP(DiffServ) DS value (0 - 63).
func (o SystemDscpbasedpriorityOutput) Ds() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDscpbasedpriority) pulumi.IntOutput { return v.Ds }).(pulumi.IntOutput)
}

// Item ID.
func (o SystemDscpbasedpriorityOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDscpbasedpriority) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// DSCP based priority level. Valid values: `low`, `medium`, `high`.
func (o SystemDscpbasedpriorityOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDscpbasedpriority) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemDscpbasedpriorityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDscpbasedpriority) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemDscpbasedpriorityArrayOutput struct{ *pulumi.OutputState }

func (SystemDscpbasedpriorityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDscpbasedpriority)(nil)).Elem()
}

func (o SystemDscpbasedpriorityArrayOutput) ToSystemDscpbasedpriorityArrayOutput() SystemDscpbasedpriorityArrayOutput {
	return o
}

func (o SystemDscpbasedpriorityArrayOutput) ToSystemDscpbasedpriorityArrayOutputWithContext(ctx context.Context) SystemDscpbasedpriorityArrayOutput {
	return o
}

func (o SystemDscpbasedpriorityArrayOutput) Index(i pulumi.IntInput) SystemDscpbasedpriorityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDscpbasedpriority {
		return vs[0].([]*SystemDscpbasedpriority)[vs[1].(int)]
	}).(SystemDscpbasedpriorityOutput)
}

type SystemDscpbasedpriorityMapOutput struct{ *pulumi.OutputState }

func (SystemDscpbasedpriorityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDscpbasedpriority)(nil)).Elem()
}

func (o SystemDscpbasedpriorityMapOutput) ToSystemDscpbasedpriorityMapOutput() SystemDscpbasedpriorityMapOutput {
	return o
}

func (o SystemDscpbasedpriorityMapOutput) ToSystemDscpbasedpriorityMapOutputWithContext(ctx context.Context) SystemDscpbasedpriorityMapOutput {
	return o
}

func (o SystemDscpbasedpriorityMapOutput) MapIndex(k pulumi.StringInput) SystemDscpbasedpriorityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDscpbasedpriority {
		return vs[0].(map[string]*SystemDscpbasedpriority)[vs[1].(string)]
	}).(SystemDscpbasedpriorityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDscpbasedpriorityInput)(nil)).Elem(), &SystemDscpbasedpriority{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDscpbasedpriorityArrayInput)(nil)).Elem(), SystemDscpbasedpriorityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDscpbasedpriorityMapInput)(nil)).Elem(), SystemDscpbasedpriorityMap{})
	pulumi.RegisterOutputType(SystemDscpbasedpriorityOutput{})
	pulumi.RegisterOutputType(SystemDscpbasedpriorityArrayOutput{})
	pulumi.RegisterOutputType(SystemDscpbasedpriorityMapOutput{})
}