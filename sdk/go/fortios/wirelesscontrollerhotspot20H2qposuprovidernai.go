// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure online sign up (OSU) provider NAI list. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// # WirelessControllerHotspot20 H2QpOsuProviderNai can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20H2qposuprovidernai:Wirelesscontrollerhotspot20H2qposuprovidernai labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20H2qposuprovidernai:Wirelesscontrollerhotspot20H2qposuprovidernai labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Wirelesscontrollerhotspot20H2qposuprovidernai struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists Wirelesscontrollerhotspot20H2qposuprovidernaiNaiListArrayOutput `pulumi:"naiLists"`
	// OSU provider NAI ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerhotspot20H2qposuprovidernai registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerhotspot20H2qposuprovidernai(ctx *pulumi.Context,
	name string, args *Wirelesscontrollerhotspot20H2qposuprovidernaiArgs, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20H2qposuprovidernai, error) {
	if args == nil {
		args = &Wirelesscontrollerhotspot20H2qposuprovidernaiArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Wirelesscontrollerhotspot20H2qposuprovidernai
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerhotspot20H2qposuprovidernai:Wirelesscontrollerhotspot20H2qposuprovidernai", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerhotspot20H2qposuprovidernai gets an existing Wirelesscontrollerhotspot20H2qposuprovidernai resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerhotspot20H2qposuprovidernai(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Wirelesscontrollerhotspot20H2qposuprovidernaiState, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20H2qposuprovidernai, error) {
	var resource Wirelesscontrollerhotspot20H2qposuprovidernai
	err := ctx.ReadResource("fortios:index/wirelesscontrollerhotspot20H2qposuprovidernai:Wirelesscontrollerhotspot20H2qposuprovidernai", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wirelesscontrollerhotspot20H2qposuprovidernai resources.
type wirelesscontrollerhotspot20H2qposuprovidernaiState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists []Wirelesscontrollerhotspot20H2qposuprovidernaiNaiList `pulumi:"naiLists"`
	// OSU provider NAI ID.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists Wirelesscontrollerhotspot20H2qposuprovidernaiNaiListArrayInput
	// OSU provider NAI ID.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20H2qposuprovidernaiState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20H2qposuprovidernaiState)(nil)).Elem()
}

type wirelesscontrollerhotspot20H2qposuprovidernaiArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists []Wirelesscontrollerhotspot20H2qposuprovidernaiNaiList `pulumi:"naiLists"`
	// OSU provider NAI ID.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Wirelesscontrollerhotspot20H2qposuprovidernai resource.
type Wirelesscontrollerhotspot20H2qposuprovidernaiArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists Wirelesscontrollerhotspot20H2qposuprovidernaiNaiListArrayInput
	// OSU provider NAI ID.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20H2qposuprovidernaiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20H2qposuprovidernaiArgs)(nil)).Elem()
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20H2qposuprovidernaiOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiOutput
	ToWirelesscontrollerhotspot20H2qposuprovidernaiOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiOutput
}

func (*Wirelesscontrollerhotspot20H2qposuprovidernai) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20H2qposuprovidernai)(nil)).Elem()
}

func (i *Wirelesscontrollerhotspot20H2qposuprovidernai) ToWirelesscontrollerhotspot20H2qposuprovidernaiOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiOutput {
	return i.ToWirelesscontrollerhotspot20H2qposuprovidernaiOutputWithContext(context.Background())
}

func (i *Wirelesscontrollerhotspot20H2qposuprovidernai) ToWirelesscontrollerhotspot20H2qposuprovidernaiOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20H2qposuprovidernaiOutput)
}

// Wirelesscontrollerhotspot20H2qposuprovidernaiArrayInput is an input type that accepts Wirelesscontrollerhotspot20H2qposuprovidernaiArray and Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20H2qposuprovidernaiArrayInput` via:
//
//	Wirelesscontrollerhotspot20H2qposuprovidernaiArray{ Wirelesscontrollerhotspot20H2qposuprovidernaiArgs{...} }
type Wirelesscontrollerhotspot20H2qposuprovidernaiArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput
	ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutputWithContext(context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiArray []Wirelesscontrollerhotspot20H2qposuprovidernaiInput

func (Wirelesscontrollerhotspot20H2qposuprovidernaiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20H2qposuprovidernai)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20H2qposuprovidernaiArray) ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput {
	return i.ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20H2qposuprovidernaiArray) ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput)
}

// Wirelesscontrollerhotspot20H2qposuprovidernaiMapInput is an input type that accepts Wirelesscontrollerhotspot20H2qposuprovidernaiMap and Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20H2qposuprovidernaiMapInput` via:
//
//	Wirelesscontrollerhotspot20H2qposuprovidernaiMap{ "key": Wirelesscontrollerhotspot20H2qposuprovidernaiArgs{...} }
type Wirelesscontrollerhotspot20H2qposuprovidernaiMapInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput
	ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutputWithContext(context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiMap map[string]Wirelesscontrollerhotspot20H2qposuprovidernaiInput

func (Wirelesscontrollerhotspot20H2qposuprovidernaiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20H2qposuprovidernai)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20H2qposuprovidernaiMap) ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput {
	return i.ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20H2qposuprovidernaiMap) ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput)
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20H2qposuprovidernai)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) ToWirelesscontrollerhotspot20H2qposuprovidernaiOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) ToWirelesscontrollerhotspot20H2qposuprovidernaiOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qposuprovidernai) pulumi.StringPtrOutput {
		return v.DynamicSortSubtable
	}).(pulumi.StringPtrOutput)
}

// OSU NAI list. The structure of `naiList` block is documented below.
func (o Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) NaiLists() Wirelesscontrollerhotspot20H2qposuprovidernaiNaiListArrayOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qposuprovidernai) Wirelesscontrollerhotspot20H2qposuprovidernaiNaiListArrayOutput {
		return v.NaiLists
	}).(Wirelesscontrollerhotspot20H2qposuprovidernaiNaiListArrayOutput)
}

// OSU provider NAI ID.
func (o Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qposuprovidernai) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Wirelesscontrollerhotspot20H2qposuprovidernaiOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qposuprovidernai) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20H2qposuprovidernai)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput) ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput) ToWirelesscontrollerhotspot20H2qposuprovidernaiArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput) Index(i pulumi.IntInput) Wirelesscontrollerhotspot20H2qposuprovidernaiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20H2qposuprovidernai {
		return vs[0].([]*Wirelesscontrollerhotspot20H2qposuprovidernai)[vs[1].(int)]
	}).(Wirelesscontrollerhotspot20H2qposuprovidernaiOutput)
}

type Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20H2qposuprovidernai)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput) ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutput() Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput) ToWirelesscontrollerhotspot20H2qposuprovidernaiMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput) MapIndex(k pulumi.StringInput) Wirelesscontrollerhotspot20H2qposuprovidernaiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20H2qposuprovidernai {
		return vs[0].(map[string]*Wirelesscontrollerhotspot20H2qposuprovidernai)[vs[1].(string)]
	}).(Wirelesscontrollerhotspot20H2qposuprovidernaiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20H2qposuprovidernaiInput)(nil)).Elem(), &Wirelesscontrollerhotspot20H2qposuprovidernai{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20H2qposuprovidernaiArrayInput)(nil)).Elem(), Wirelesscontrollerhotspot20H2qposuprovidernaiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20H2qposuprovidernaiMapInput)(nil)).Elem(), Wirelesscontrollerhotspot20H2qposuprovidernaiMap{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20H2qposuprovidernaiOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20H2qposuprovidernaiArrayOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20H2qposuprovidernaiMapOutput{})
}