// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WiFi bridge access control list. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// # WirelessController AccessControlList can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerAccesscontrollist struct {
	pulumi.CustomResourceState

	// Description.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
	Layer3Ipv4Rules WirelesscontrollerAccesscontrollistLayer3Ipv4RuleArrayOutput `pulumi:"layer3Ipv4Rules"`
	// AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
	Layer3Ipv6Rules WirelesscontrollerAccesscontrollistLayer3Ipv6RuleArrayOutput `pulumi:"layer3Ipv6Rules"`
	// AP access control list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `layer3Ipv4Rules` block supports:
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerAccesscontrollist registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerAccesscontrollist(ctx *pulumi.Context,
	name string, args *WirelesscontrollerAccesscontrollistArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerAccesscontrollist, error) {
	if args == nil {
		args = &WirelesscontrollerAccesscontrollistArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerAccesscontrollist
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerAccesscontrollist gets an existing WirelesscontrollerAccesscontrollist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerAccesscontrollist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerAccesscontrollistState, opts ...pulumi.ResourceOption) (*WirelesscontrollerAccesscontrollist, error) {
	var resource WirelesscontrollerAccesscontrollist
	err := ctx.ReadResource("fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerAccesscontrollist resources.
type wirelesscontrollerAccesscontrollistState struct {
	// Description.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
	Layer3Ipv4Rules []WirelesscontrollerAccesscontrollistLayer3Ipv4Rule `pulumi:"layer3Ipv4Rules"`
	// AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
	Layer3Ipv6Rules []WirelesscontrollerAccesscontrollistLayer3Ipv6Rule `pulumi:"layer3Ipv6Rules"`
	// AP access control list name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `layer3Ipv4Rules` block supports:
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelesscontrollerAccesscontrollistState struct {
	// Description.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
	Layer3Ipv4Rules WirelesscontrollerAccesscontrollistLayer3Ipv4RuleArrayInput
	// AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
	Layer3Ipv6Rules WirelesscontrollerAccesscontrollistLayer3Ipv6RuleArrayInput
	// AP access control list name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `layer3Ipv4Rules` block supports:
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerAccesscontrollistState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerAccesscontrollistState)(nil)).Elem()
}

type wirelesscontrollerAccesscontrollistArgs struct {
	// Description.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
	Layer3Ipv4Rules []WirelesscontrollerAccesscontrollistLayer3Ipv4Rule `pulumi:"layer3Ipv4Rules"`
	// AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
	Layer3Ipv6Rules []WirelesscontrollerAccesscontrollistLayer3Ipv6Rule `pulumi:"layer3Ipv6Rules"`
	// AP access control list name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `layer3Ipv4Rules` block supports:
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelesscontrollerAccesscontrollist resource.
type WirelesscontrollerAccesscontrollistArgs struct {
	// Description.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
	Layer3Ipv4Rules WirelesscontrollerAccesscontrollistLayer3Ipv4RuleArrayInput
	// AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
	Layer3Ipv6Rules WirelesscontrollerAccesscontrollistLayer3Ipv6RuleArrayInput
	// AP access control list name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `layer3Ipv4Rules` block supports:
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerAccesscontrollistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerAccesscontrollistArgs)(nil)).Elem()
}

type WirelesscontrollerAccesscontrollistInput interface {
	pulumi.Input

	ToWirelesscontrollerAccesscontrollistOutput() WirelesscontrollerAccesscontrollistOutput
	ToWirelesscontrollerAccesscontrollistOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistOutput
}

func (*WirelesscontrollerAccesscontrollist) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerAccesscontrollist)(nil)).Elem()
}

func (i *WirelesscontrollerAccesscontrollist) ToWirelesscontrollerAccesscontrollistOutput() WirelesscontrollerAccesscontrollistOutput {
	return i.ToWirelesscontrollerAccesscontrollistOutputWithContext(context.Background())
}

func (i *WirelesscontrollerAccesscontrollist) ToWirelesscontrollerAccesscontrollistOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerAccesscontrollistOutput)
}

// WirelesscontrollerAccesscontrollistArrayInput is an input type that accepts WirelesscontrollerAccesscontrollistArray and WirelesscontrollerAccesscontrollistArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerAccesscontrollistArrayInput` via:
//
//	WirelesscontrollerAccesscontrollistArray{ WirelesscontrollerAccesscontrollistArgs{...} }
type WirelesscontrollerAccesscontrollistArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerAccesscontrollistArrayOutput() WirelesscontrollerAccesscontrollistArrayOutput
	ToWirelesscontrollerAccesscontrollistArrayOutputWithContext(context.Context) WirelesscontrollerAccesscontrollistArrayOutput
}

type WirelesscontrollerAccesscontrollistArray []WirelesscontrollerAccesscontrollistInput

func (WirelesscontrollerAccesscontrollistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerAccesscontrollist)(nil)).Elem()
}

func (i WirelesscontrollerAccesscontrollistArray) ToWirelesscontrollerAccesscontrollistArrayOutput() WirelesscontrollerAccesscontrollistArrayOutput {
	return i.ToWirelesscontrollerAccesscontrollistArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerAccesscontrollistArray) ToWirelesscontrollerAccesscontrollistArrayOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerAccesscontrollistArrayOutput)
}

// WirelesscontrollerAccesscontrollistMapInput is an input type that accepts WirelesscontrollerAccesscontrollistMap and WirelesscontrollerAccesscontrollistMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerAccesscontrollistMapInput` via:
//
//	WirelesscontrollerAccesscontrollistMap{ "key": WirelesscontrollerAccesscontrollistArgs{...} }
type WirelesscontrollerAccesscontrollistMapInput interface {
	pulumi.Input

	ToWirelesscontrollerAccesscontrollistMapOutput() WirelesscontrollerAccesscontrollistMapOutput
	ToWirelesscontrollerAccesscontrollistMapOutputWithContext(context.Context) WirelesscontrollerAccesscontrollistMapOutput
}

type WirelesscontrollerAccesscontrollistMap map[string]WirelesscontrollerAccesscontrollistInput

func (WirelesscontrollerAccesscontrollistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerAccesscontrollist)(nil)).Elem()
}

func (i WirelesscontrollerAccesscontrollistMap) ToWirelesscontrollerAccesscontrollistMapOutput() WirelesscontrollerAccesscontrollistMapOutput {
	return i.ToWirelesscontrollerAccesscontrollistMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerAccesscontrollistMap) ToWirelesscontrollerAccesscontrollistMapOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerAccesscontrollistMapOutput)
}

type WirelesscontrollerAccesscontrollistOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerAccesscontrollistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerAccesscontrollist)(nil)).Elem()
}

func (o WirelesscontrollerAccesscontrollistOutput) ToWirelesscontrollerAccesscontrollistOutput() WirelesscontrollerAccesscontrollistOutput {
	return o
}

func (o WirelesscontrollerAccesscontrollistOutput) ToWirelesscontrollerAccesscontrollistOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistOutput {
	return o
}

// Description.
func (o WirelesscontrollerAccesscontrollistOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerAccesscontrollist) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o WirelesscontrollerAccesscontrollistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerAccesscontrollist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
func (o WirelesscontrollerAccesscontrollistOutput) Layer3Ipv4Rules() WirelesscontrollerAccesscontrollistLayer3Ipv4RuleArrayOutput {
	return o.ApplyT(func(v *WirelesscontrollerAccesscontrollist) WirelesscontrollerAccesscontrollistLayer3Ipv4RuleArrayOutput {
		return v.Layer3Ipv4Rules
	}).(WirelesscontrollerAccesscontrollistLayer3Ipv4RuleArrayOutput)
}

// AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
func (o WirelesscontrollerAccesscontrollistOutput) Layer3Ipv6Rules() WirelesscontrollerAccesscontrollistLayer3Ipv6RuleArrayOutput {
	return o.ApplyT(func(v *WirelesscontrollerAccesscontrollist) WirelesscontrollerAccesscontrollistLayer3Ipv6RuleArrayOutput {
		return v.Layer3Ipv6Rules
	}).(WirelesscontrollerAccesscontrollistLayer3Ipv6RuleArrayOutput)
}

// AP access control list name.
func (o WirelesscontrollerAccesscontrollistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerAccesscontrollist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
//
// The `layer3Ipv4Rules` block supports:
func (o WirelesscontrollerAccesscontrollistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerAccesscontrollist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelesscontrollerAccesscontrollistArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerAccesscontrollistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerAccesscontrollist)(nil)).Elem()
}

func (o WirelesscontrollerAccesscontrollistArrayOutput) ToWirelesscontrollerAccesscontrollistArrayOutput() WirelesscontrollerAccesscontrollistArrayOutput {
	return o
}

func (o WirelesscontrollerAccesscontrollistArrayOutput) ToWirelesscontrollerAccesscontrollistArrayOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistArrayOutput {
	return o
}

func (o WirelesscontrollerAccesscontrollistArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerAccesscontrollistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerAccesscontrollist {
		return vs[0].([]*WirelesscontrollerAccesscontrollist)[vs[1].(int)]
	}).(WirelesscontrollerAccesscontrollistOutput)
}

type WirelesscontrollerAccesscontrollistMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerAccesscontrollistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerAccesscontrollist)(nil)).Elem()
}

func (o WirelesscontrollerAccesscontrollistMapOutput) ToWirelesscontrollerAccesscontrollistMapOutput() WirelesscontrollerAccesscontrollistMapOutput {
	return o
}

func (o WirelesscontrollerAccesscontrollistMapOutput) ToWirelesscontrollerAccesscontrollistMapOutputWithContext(ctx context.Context) WirelesscontrollerAccesscontrollistMapOutput {
	return o
}

func (o WirelesscontrollerAccesscontrollistMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerAccesscontrollistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerAccesscontrollist {
		return vs[0].(map[string]*WirelesscontrollerAccesscontrollist)[vs[1].(string)]
	}).(WirelesscontrollerAccesscontrollistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerAccesscontrollistInput)(nil)).Elem(), &WirelesscontrollerAccesscontrollist{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerAccesscontrollistArrayInput)(nil)).Elem(), WirelesscontrollerAccesscontrollistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerAccesscontrollistMapInput)(nil)).Elem(), WirelesscontrollerAccesscontrollistMap{})
	pulumi.RegisterOutputType(WirelesscontrollerAccesscontrollistOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerAccesscontrollistArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerAccesscontrollistMapOutput{})
}
