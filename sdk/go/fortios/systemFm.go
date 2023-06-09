// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FM. Applies to FortiOS Version `<= 7.0.1`.
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
//			_, err := fortios.NewSystemFm(ctx, "trname", &fortios.SystemFmArgs{
//				AutoBackup:             pulumi.String("disable"),
//				Ip:                     pulumi.String("0.0.0.0"),
//				Ipsec:                  pulumi.String("disable"),
//				ScheduledConfigRestore: pulumi.String("disable"),
//				Status:                 pulumi.String("disable"),
//				Vdom:                   pulumi.String("root"),
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
// # System Fm can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemFm:SystemFm labelname SystemFm
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemFm:SystemFm labelname SystemFm
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemFm struct {
	pulumi.CustomResourceState

	// Enable/disable automatic backup. Valid values: `enable`, `disable`.
	AutoBackup pulumi.StringOutput `pulumi:"autoBackup"`
	// ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// IP address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Enable/disable IPsec. Valid values: `enable`, `disable`.
	Ipsec pulumi.StringOutput `pulumi:"ipsec"`
	// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
	ScheduledConfigRestore pulumi.StringOutput `pulumi:"scheduledConfigRestore"`
	// Enable/disable FM. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// VDOM.
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFm registers a new resource with the given unique name, arguments, and options.
func NewSystemFm(ctx *pulumi.Context,
	name string, args *SystemFmArgs, opts ...pulumi.ResourceOption) (*SystemFm, error) {
	if args == nil {
		args = &SystemFmArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemFm
	err := ctx.RegisterResource("fortios:index/systemFm:SystemFm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFm gets an existing SystemFm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFmState, opts ...pulumi.ResourceOption) (*SystemFm, error) {
	var resource SystemFm
	err := ctx.ReadResource("fortios:index/systemFm:SystemFm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFm resources.
type systemFmState struct {
	// Enable/disable automatic backup. Valid values: `enable`, `disable`.
	AutoBackup *string `pulumi:"autoBackup"`
	// ID.
	Fosid *string `pulumi:"fosid"`
	// IP address.
	Ip *string `pulumi:"ip"`
	// Enable/disable IPsec. Valid values: `enable`, `disable`.
	Ipsec *string `pulumi:"ipsec"`
	// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
	ScheduledConfigRestore *string `pulumi:"scheduledConfigRestore"`
	// Enable/disable FM. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// VDOM.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemFmState struct {
	// Enable/disable automatic backup. Valid values: `enable`, `disable`.
	AutoBackup pulumi.StringPtrInput
	// ID.
	Fosid pulumi.StringPtrInput
	// IP address.
	Ip pulumi.StringPtrInput
	// Enable/disable IPsec. Valid values: `enable`, `disable`.
	Ipsec pulumi.StringPtrInput
	// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
	ScheduledConfigRestore pulumi.StringPtrInput
	// Enable/disable FM. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFmState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFmState)(nil)).Elem()
}

type systemFmArgs struct {
	// Enable/disable automatic backup. Valid values: `enable`, `disable`.
	AutoBackup *string `pulumi:"autoBackup"`
	// ID.
	Fosid *string `pulumi:"fosid"`
	// IP address.
	Ip *string `pulumi:"ip"`
	// Enable/disable IPsec. Valid values: `enable`, `disable`.
	Ipsec *string `pulumi:"ipsec"`
	// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
	ScheduledConfigRestore *string `pulumi:"scheduledConfigRestore"`
	// Enable/disable FM. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// VDOM.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFm resource.
type SystemFmArgs struct {
	// Enable/disable automatic backup. Valid values: `enable`, `disable`.
	AutoBackup pulumi.StringPtrInput
	// ID.
	Fosid pulumi.StringPtrInput
	// IP address.
	Ip pulumi.StringPtrInput
	// Enable/disable IPsec. Valid values: `enable`, `disable`.
	Ipsec pulumi.StringPtrInput
	// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
	ScheduledConfigRestore pulumi.StringPtrInput
	// Enable/disable FM. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFmArgs)(nil)).Elem()
}

type SystemFmInput interface {
	pulumi.Input

	ToSystemFmOutput() SystemFmOutput
	ToSystemFmOutputWithContext(ctx context.Context) SystemFmOutput
}

func (*SystemFm) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFm)(nil)).Elem()
}

func (i *SystemFm) ToSystemFmOutput() SystemFmOutput {
	return i.ToSystemFmOutputWithContext(context.Background())
}

func (i *SystemFm) ToSystemFmOutputWithContext(ctx context.Context) SystemFmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFmOutput)
}

// SystemFmArrayInput is an input type that accepts SystemFmArray and SystemFmArrayOutput values.
// You can construct a concrete instance of `SystemFmArrayInput` via:
//
//	SystemFmArray{ SystemFmArgs{...} }
type SystemFmArrayInput interface {
	pulumi.Input

	ToSystemFmArrayOutput() SystemFmArrayOutput
	ToSystemFmArrayOutputWithContext(context.Context) SystemFmArrayOutput
}

type SystemFmArray []SystemFmInput

func (SystemFmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFm)(nil)).Elem()
}

func (i SystemFmArray) ToSystemFmArrayOutput() SystemFmArrayOutput {
	return i.ToSystemFmArrayOutputWithContext(context.Background())
}

func (i SystemFmArray) ToSystemFmArrayOutputWithContext(ctx context.Context) SystemFmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFmArrayOutput)
}

// SystemFmMapInput is an input type that accepts SystemFmMap and SystemFmMapOutput values.
// You can construct a concrete instance of `SystemFmMapInput` via:
//
//	SystemFmMap{ "key": SystemFmArgs{...} }
type SystemFmMapInput interface {
	pulumi.Input

	ToSystemFmMapOutput() SystemFmMapOutput
	ToSystemFmMapOutputWithContext(context.Context) SystemFmMapOutput
}

type SystemFmMap map[string]SystemFmInput

func (SystemFmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFm)(nil)).Elem()
}

func (i SystemFmMap) ToSystemFmMapOutput() SystemFmMapOutput {
	return i.ToSystemFmMapOutputWithContext(context.Background())
}

func (i SystemFmMap) ToSystemFmMapOutputWithContext(ctx context.Context) SystemFmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFmMapOutput)
}

type SystemFmOutput struct{ *pulumi.OutputState }

func (SystemFmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFm)(nil)).Elem()
}

func (o SystemFmOutput) ToSystemFmOutput() SystemFmOutput {
	return o
}

func (o SystemFmOutput) ToSystemFmOutputWithContext(ctx context.Context) SystemFmOutput {
	return o
}

// Enable/disable automatic backup. Valid values: `enable`, `disable`.
func (o SystemFmOutput) AutoBackup() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.AutoBackup }).(pulumi.StringOutput)
}

// ID.
func (o SystemFmOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// IP address.
func (o SystemFmOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Enable/disable IPsec. Valid values: `enable`, `disable`.
func (o SystemFmOutput) Ipsec() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.Ipsec }).(pulumi.StringOutput)
}

// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
func (o SystemFmOutput) ScheduledConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.ScheduledConfigRestore }).(pulumi.StringOutput)
}

// Enable/disable FM. Valid values: `enable`, `disable`.
func (o SystemFmOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// VDOM.
func (o SystemFmOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemFmOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFm) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFmArrayOutput struct{ *pulumi.OutputState }

func (SystemFmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFm)(nil)).Elem()
}

func (o SystemFmArrayOutput) ToSystemFmArrayOutput() SystemFmArrayOutput {
	return o
}

func (o SystemFmArrayOutput) ToSystemFmArrayOutputWithContext(ctx context.Context) SystemFmArrayOutput {
	return o
}

func (o SystemFmArrayOutput) Index(i pulumi.IntInput) SystemFmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFm {
		return vs[0].([]*SystemFm)[vs[1].(int)]
	}).(SystemFmOutput)
}

type SystemFmMapOutput struct{ *pulumi.OutputState }

func (SystemFmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFm)(nil)).Elem()
}

func (o SystemFmMapOutput) ToSystemFmMapOutput() SystemFmMapOutput {
	return o
}

func (o SystemFmMapOutput) ToSystemFmMapOutputWithContext(ctx context.Context) SystemFmMapOutput {
	return o
}

func (o SystemFmMapOutput) MapIndex(k pulumi.StringInput) SystemFmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFm {
		return vs[0].(map[string]*SystemFm)[vs[1].(string)]
	}).(SystemFmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFmInput)(nil)).Elem(), &SystemFm{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFmArrayInput)(nil)).Elem(), SystemFmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFmMapInput)(nil)).Elem(), SystemFmMap{})
	pulumi.RegisterOutputType(SystemFmOutput{})
	pulumi.RegisterOutputType(SystemFmArrayOutput{})
	pulumi.RegisterOutputType(SystemFmMapOutput{})
}
