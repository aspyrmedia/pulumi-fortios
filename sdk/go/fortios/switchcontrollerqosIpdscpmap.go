// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch QoS IP precedence/DSCP.
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
//			_, err := fortios.NewSwitchcontrollerqosIpdscpmap(ctx, "trname", &fortios.SwitchcontrollerqosIpdscpmapArgs{
//				Description: pulumi.String("DEIW"),
//				Maps: fortios.SwitchcontrollerqosIpdscpmapMapTypeArray{
//					&fortios.SwitchcontrollerqosIpdscpmapMapTypeArgs{
//						CosQueue: pulumi.Int(3),
//						Diffserv: pulumi.String("CS0 CS1 AF11"),
//						Name:     pulumi.String("1"),
//					},
//				},
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
// # SwitchControllerQos IpDscpMap can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerqosIpdscpmap struct {
	pulumi.CustomResourceState

	// Description of the ip-dscp map name.
	Description pulumi.StringOutput `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
	Maps SwitchcontrollerqosIpdscpmapMapTypeArrayOutput `pulumi:"maps"`
	// Dscp map name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerqosIpdscpmap registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerqosIpdscpmap(ctx *pulumi.Context,
	name string, args *SwitchcontrollerqosIpdscpmapArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerqosIpdscpmap, error) {
	if args == nil {
		args = &SwitchcontrollerqosIpdscpmapArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerqosIpdscpmap
	err := ctx.RegisterResource("fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerqosIpdscpmap gets an existing SwitchcontrollerqosIpdscpmap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerqosIpdscpmap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerqosIpdscpmapState, opts ...pulumi.ResourceOption) (*SwitchcontrollerqosIpdscpmap, error) {
	var resource SwitchcontrollerqosIpdscpmap
	err := ctx.ReadResource("fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerqosIpdscpmap resources.
type switchcontrollerqosIpdscpmapState struct {
	// Description of the ip-dscp map name.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
	Maps []SwitchcontrollerqosIpdscpmapMapType `pulumi:"maps"`
	// Dscp map name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerqosIpdscpmapState struct {
	// Description of the ip-dscp map name.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
	Maps SwitchcontrollerqosIpdscpmapMapTypeArrayInput
	// Dscp map name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerqosIpdscpmapState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerqosIpdscpmapState)(nil)).Elem()
}

type switchcontrollerqosIpdscpmapArgs struct {
	// Description of the ip-dscp map name.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
	Maps []SwitchcontrollerqosIpdscpmapMapType `pulumi:"maps"`
	// Dscp map name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerqosIpdscpmap resource.
type SwitchcontrollerqosIpdscpmapArgs struct {
	// Description of the ip-dscp map name.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
	Maps SwitchcontrollerqosIpdscpmapMapTypeArrayInput
	// Dscp map name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerqosIpdscpmapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerqosIpdscpmapArgs)(nil)).Elem()
}

type SwitchcontrollerqosIpdscpmapInput interface {
	pulumi.Input

	ToSwitchcontrollerqosIpdscpmapOutput() SwitchcontrollerqosIpdscpmapOutput
	ToSwitchcontrollerqosIpdscpmapOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapOutput
}

func (*SwitchcontrollerqosIpdscpmap) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerqosIpdscpmap)(nil)).Elem()
}

func (i *SwitchcontrollerqosIpdscpmap) ToSwitchcontrollerqosIpdscpmapOutput() SwitchcontrollerqosIpdscpmapOutput {
	return i.ToSwitchcontrollerqosIpdscpmapOutputWithContext(context.Background())
}

func (i *SwitchcontrollerqosIpdscpmap) ToSwitchcontrollerqosIpdscpmapOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerqosIpdscpmapOutput)
}

// SwitchcontrollerqosIpdscpmapArrayInput is an input type that accepts SwitchcontrollerqosIpdscpmapArray and SwitchcontrollerqosIpdscpmapArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerqosIpdscpmapArrayInput` via:
//
//	SwitchcontrollerqosIpdscpmapArray{ SwitchcontrollerqosIpdscpmapArgs{...} }
type SwitchcontrollerqosIpdscpmapArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerqosIpdscpmapArrayOutput() SwitchcontrollerqosIpdscpmapArrayOutput
	ToSwitchcontrollerqosIpdscpmapArrayOutputWithContext(context.Context) SwitchcontrollerqosIpdscpmapArrayOutput
}

type SwitchcontrollerqosIpdscpmapArray []SwitchcontrollerqosIpdscpmapInput

func (SwitchcontrollerqosIpdscpmapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerqosIpdscpmap)(nil)).Elem()
}

func (i SwitchcontrollerqosIpdscpmapArray) ToSwitchcontrollerqosIpdscpmapArrayOutput() SwitchcontrollerqosIpdscpmapArrayOutput {
	return i.ToSwitchcontrollerqosIpdscpmapArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerqosIpdscpmapArray) ToSwitchcontrollerqosIpdscpmapArrayOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerqosIpdscpmapArrayOutput)
}

// SwitchcontrollerqosIpdscpmapMapInput is an input type that accepts SwitchcontrollerqosIpdscpmapMap and SwitchcontrollerqosIpdscpmapMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerqosIpdscpmapMapInput` via:
//
//	SwitchcontrollerqosIpdscpmapMap{ "key": SwitchcontrollerqosIpdscpmapArgs{...} }
type SwitchcontrollerqosIpdscpmapMapInput interface {
	pulumi.Input

	ToSwitchcontrollerqosIpdscpmapMapOutput() SwitchcontrollerqosIpdscpmapMapOutput
	ToSwitchcontrollerqosIpdscpmapMapOutputWithContext(context.Context) SwitchcontrollerqosIpdscpmapMapOutput
}

type SwitchcontrollerqosIpdscpmapMap map[string]SwitchcontrollerqosIpdscpmapInput

func (SwitchcontrollerqosIpdscpmapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerqosIpdscpmap)(nil)).Elem()
}

func (i SwitchcontrollerqosIpdscpmapMap) ToSwitchcontrollerqosIpdscpmapMapOutput() SwitchcontrollerqosIpdscpmapMapOutput {
	return i.ToSwitchcontrollerqosIpdscpmapMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerqosIpdscpmapMap) ToSwitchcontrollerqosIpdscpmapMapOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerqosIpdscpmapMapOutput)
}

type SwitchcontrollerqosIpdscpmapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerqosIpdscpmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerqosIpdscpmap)(nil)).Elem()
}

func (o SwitchcontrollerqosIpdscpmapOutput) ToSwitchcontrollerqosIpdscpmapOutput() SwitchcontrollerqosIpdscpmapOutput {
	return o
}

func (o SwitchcontrollerqosIpdscpmapOutput) ToSwitchcontrollerqosIpdscpmapOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapOutput {
	return o
}

// Description of the ip-dscp map name.
func (o SwitchcontrollerqosIpdscpmapOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerqosIpdscpmap) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchcontrollerqosIpdscpmapOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerqosIpdscpmap) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
func (o SwitchcontrollerqosIpdscpmapOutput) Maps() SwitchcontrollerqosIpdscpmapMapTypeArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerqosIpdscpmap) SwitchcontrollerqosIpdscpmapMapTypeArrayOutput { return v.Maps }).(SwitchcontrollerqosIpdscpmapMapTypeArrayOutput)
}

// Dscp map name.
func (o SwitchcontrollerqosIpdscpmapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerqosIpdscpmap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerqosIpdscpmapOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerqosIpdscpmap) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerqosIpdscpmapArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerqosIpdscpmapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerqosIpdscpmap)(nil)).Elem()
}

func (o SwitchcontrollerqosIpdscpmapArrayOutput) ToSwitchcontrollerqosIpdscpmapArrayOutput() SwitchcontrollerqosIpdscpmapArrayOutput {
	return o
}

func (o SwitchcontrollerqosIpdscpmapArrayOutput) ToSwitchcontrollerqosIpdscpmapArrayOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapArrayOutput {
	return o
}

func (o SwitchcontrollerqosIpdscpmapArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerqosIpdscpmapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerqosIpdscpmap {
		return vs[0].([]*SwitchcontrollerqosIpdscpmap)[vs[1].(int)]
	}).(SwitchcontrollerqosIpdscpmapOutput)
}

type SwitchcontrollerqosIpdscpmapMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerqosIpdscpmapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerqosIpdscpmap)(nil)).Elem()
}

func (o SwitchcontrollerqosIpdscpmapMapOutput) ToSwitchcontrollerqosIpdscpmapMapOutput() SwitchcontrollerqosIpdscpmapMapOutput {
	return o
}

func (o SwitchcontrollerqosIpdscpmapMapOutput) ToSwitchcontrollerqosIpdscpmapMapOutputWithContext(ctx context.Context) SwitchcontrollerqosIpdscpmapMapOutput {
	return o
}

func (o SwitchcontrollerqosIpdscpmapMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerqosIpdscpmapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerqosIpdscpmap {
		return vs[0].(map[string]*SwitchcontrollerqosIpdscpmap)[vs[1].(string)]
	}).(SwitchcontrollerqosIpdscpmapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerqosIpdscpmapInput)(nil)).Elem(), &SwitchcontrollerqosIpdscpmap{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerqosIpdscpmapArrayInput)(nil)).Elem(), SwitchcontrollerqosIpdscpmapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerqosIpdscpmapMapInput)(nil)).Elem(), SwitchcontrollerqosIpdscpmapMap{})
	pulumi.RegisterOutputType(SwitchcontrollerqosIpdscpmapOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerqosIpdscpmapArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerqosIpdscpmapMapOutput{})
}
