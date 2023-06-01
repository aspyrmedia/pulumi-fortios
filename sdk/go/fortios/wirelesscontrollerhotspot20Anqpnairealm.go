// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure network access identifier (NAI) realm.
//
// ## Import
//
// # WirelessControllerHotspot20 AnqpNaiRealm can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpnairealm:Wirelesscontrollerhotspot20Anqpnairealm labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpnairealm:Wirelesscontrollerhotspot20Anqpnairealm labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Wirelesscontrollerhotspot20Anqpnairealm struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// NAI list. The structure of `naiList` block is documented below.
	NaiLists Wirelesscontrollerhotspot20AnqpnairealmNaiListArrayOutput `pulumi:"naiLists"`
	// NAI realm list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerhotspot20Anqpnairealm registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerhotspot20Anqpnairealm(ctx *pulumi.Context,
	name string, args *Wirelesscontrollerhotspot20AnqpnairealmArgs, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20Anqpnairealm, error) {
	if args == nil {
		args = &Wirelesscontrollerhotspot20AnqpnairealmArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Wirelesscontrollerhotspot20Anqpnairealm
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerhotspot20Anqpnairealm:Wirelesscontrollerhotspot20Anqpnairealm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerhotspot20Anqpnairealm gets an existing Wirelesscontrollerhotspot20Anqpnairealm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerhotspot20Anqpnairealm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Wirelesscontrollerhotspot20AnqpnairealmState, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20Anqpnairealm, error) {
	var resource Wirelesscontrollerhotspot20Anqpnairealm
	err := ctx.ReadResource("fortios:index/wirelesscontrollerhotspot20Anqpnairealm:Wirelesscontrollerhotspot20Anqpnairealm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wirelesscontrollerhotspot20Anqpnairealm resources.
type wirelesscontrollerhotspot20AnqpnairealmState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// NAI list. The structure of `naiList` block is documented below.
	NaiLists []Wirelesscontrollerhotspot20AnqpnairealmNaiList `pulumi:"naiLists"`
	// NAI realm list name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Wirelesscontrollerhotspot20AnqpnairealmState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// NAI list. The structure of `naiList` block is documented below.
	NaiLists Wirelesscontrollerhotspot20AnqpnairealmNaiListArrayInput
	// NAI realm list name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20AnqpnairealmState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20AnqpnairealmState)(nil)).Elem()
}

type wirelesscontrollerhotspot20AnqpnairealmArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// NAI list. The structure of `naiList` block is documented below.
	NaiLists []Wirelesscontrollerhotspot20AnqpnairealmNaiList `pulumi:"naiLists"`
	// NAI realm list name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Wirelesscontrollerhotspot20Anqpnairealm resource.
type Wirelesscontrollerhotspot20AnqpnairealmArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// NAI list. The structure of `naiList` block is documented below.
	NaiLists Wirelesscontrollerhotspot20AnqpnairealmNaiListArrayInput
	// NAI realm list name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20AnqpnairealmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20AnqpnairealmArgs)(nil)).Elem()
}

type Wirelesscontrollerhotspot20AnqpnairealmInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpnairealmOutput() Wirelesscontrollerhotspot20AnqpnairealmOutput
	ToWirelesscontrollerhotspot20AnqpnairealmOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmOutput
}

func (*Wirelesscontrollerhotspot20Anqpnairealm) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20Anqpnairealm)(nil)).Elem()
}

func (i *Wirelesscontrollerhotspot20Anqpnairealm) ToWirelesscontrollerhotspot20AnqpnairealmOutput() Wirelesscontrollerhotspot20AnqpnairealmOutput {
	return i.ToWirelesscontrollerhotspot20AnqpnairealmOutputWithContext(context.Background())
}

func (i *Wirelesscontrollerhotspot20Anqpnairealm) ToWirelesscontrollerhotspot20AnqpnairealmOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpnairealmOutput)
}

// Wirelesscontrollerhotspot20AnqpnairealmArrayInput is an input type that accepts Wirelesscontrollerhotspot20AnqpnairealmArray and Wirelesscontrollerhotspot20AnqpnairealmArrayOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20AnqpnairealmArrayInput` via:
//
//	Wirelesscontrollerhotspot20AnqpnairealmArray{ Wirelesscontrollerhotspot20AnqpnairealmArgs{...} }
type Wirelesscontrollerhotspot20AnqpnairealmArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpnairealmArrayOutput() Wirelesscontrollerhotspot20AnqpnairealmArrayOutput
	ToWirelesscontrollerhotspot20AnqpnairealmArrayOutputWithContext(context.Context) Wirelesscontrollerhotspot20AnqpnairealmArrayOutput
}

type Wirelesscontrollerhotspot20AnqpnairealmArray []Wirelesscontrollerhotspot20AnqpnairealmInput

func (Wirelesscontrollerhotspot20AnqpnairealmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20Anqpnairealm)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20AnqpnairealmArray) ToWirelesscontrollerhotspot20AnqpnairealmArrayOutput() Wirelesscontrollerhotspot20AnqpnairealmArrayOutput {
	return i.ToWirelesscontrollerhotspot20AnqpnairealmArrayOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20AnqpnairealmArray) ToWirelesscontrollerhotspot20AnqpnairealmArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpnairealmArrayOutput)
}

// Wirelesscontrollerhotspot20AnqpnairealmMapInput is an input type that accepts Wirelesscontrollerhotspot20AnqpnairealmMap and Wirelesscontrollerhotspot20AnqpnairealmMapOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20AnqpnairealmMapInput` via:
//
//	Wirelesscontrollerhotspot20AnqpnairealmMap{ "key": Wirelesscontrollerhotspot20AnqpnairealmArgs{...} }
type Wirelesscontrollerhotspot20AnqpnairealmMapInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpnairealmMapOutput() Wirelesscontrollerhotspot20AnqpnairealmMapOutput
	ToWirelesscontrollerhotspot20AnqpnairealmMapOutputWithContext(context.Context) Wirelesscontrollerhotspot20AnqpnairealmMapOutput
}

type Wirelesscontrollerhotspot20AnqpnairealmMap map[string]Wirelesscontrollerhotspot20AnqpnairealmInput

func (Wirelesscontrollerhotspot20AnqpnairealmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20Anqpnairealm)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20AnqpnairealmMap) ToWirelesscontrollerhotspot20AnqpnairealmMapOutput() Wirelesscontrollerhotspot20AnqpnairealmMapOutput {
	return i.ToWirelesscontrollerhotspot20AnqpnairealmMapOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20AnqpnairealmMap) ToWirelesscontrollerhotspot20AnqpnairealmMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpnairealmMapOutput)
}

type Wirelesscontrollerhotspot20AnqpnairealmOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpnairealmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20Anqpnairealm)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpnairealmOutput) ToWirelesscontrollerhotspot20AnqpnairealmOutput() Wirelesscontrollerhotspot20AnqpnairealmOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnairealmOutput) ToWirelesscontrollerhotspot20AnqpnairealmOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Wirelesscontrollerhotspot20AnqpnairealmOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnairealm) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// NAI list. The structure of `naiList` block is documented below.
func (o Wirelesscontrollerhotspot20AnqpnairealmOutput) NaiLists() Wirelesscontrollerhotspot20AnqpnairealmNaiListArrayOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnairealm) Wirelesscontrollerhotspot20AnqpnairealmNaiListArrayOutput {
		return v.NaiLists
	}).(Wirelesscontrollerhotspot20AnqpnairealmNaiListArrayOutput)
}

// NAI realm list name.
func (o Wirelesscontrollerhotspot20AnqpnairealmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnairealm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Wirelesscontrollerhotspot20AnqpnairealmOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnairealm) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Wirelesscontrollerhotspot20AnqpnairealmArrayOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpnairealmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20Anqpnairealm)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpnairealmArrayOutput) ToWirelesscontrollerhotspot20AnqpnairealmArrayOutput() Wirelesscontrollerhotspot20AnqpnairealmArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnairealmArrayOutput) ToWirelesscontrollerhotspot20AnqpnairealmArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnairealmArrayOutput) Index(i pulumi.IntInput) Wirelesscontrollerhotspot20AnqpnairealmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20Anqpnairealm {
		return vs[0].([]*Wirelesscontrollerhotspot20Anqpnairealm)[vs[1].(int)]
	}).(Wirelesscontrollerhotspot20AnqpnairealmOutput)
}

type Wirelesscontrollerhotspot20AnqpnairealmMapOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpnairealmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20Anqpnairealm)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpnairealmMapOutput) ToWirelesscontrollerhotspot20AnqpnairealmMapOutput() Wirelesscontrollerhotspot20AnqpnairealmMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnairealmMapOutput) ToWirelesscontrollerhotspot20AnqpnairealmMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnairealmMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnairealmMapOutput) MapIndex(k pulumi.StringInput) Wirelesscontrollerhotspot20AnqpnairealmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20Anqpnairealm {
		return vs[0].(map[string]*Wirelesscontrollerhotspot20Anqpnairealm)[vs[1].(string)]
	}).(Wirelesscontrollerhotspot20AnqpnairealmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpnairealmInput)(nil)).Elem(), &Wirelesscontrollerhotspot20Anqpnairealm{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpnairealmArrayInput)(nil)).Elem(), Wirelesscontrollerhotspot20AnqpnairealmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpnairealmMapInput)(nil)).Elem(), Wirelesscontrollerhotspot20AnqpnairealmMap{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpnairealmOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpnairealmArrayOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpnairealmMapOutput{})
}