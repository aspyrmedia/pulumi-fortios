// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Web proxy address group configuration.
//
// ## Import
//
// # Firewall ProxyAddrgrp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallProxyaddrgrp:FirewallProxyaddrgrp labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallProxyaddrgrp:FirewallProxyaddrgrp labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallProxyaddrgrp struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Members of address group. The structure of `member` block is documented below.
	Members FirewallProxyaddrgrpMemberArrayOutput `pulumi:"members"`
	// Address group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyaddrgrpTaggingArrayOutput `pulumi:"taggings"`
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallProxyaddrgrp registers a new resource with the given unique name, arguments, and options.
func NewFirewallProxyaddrgrp(ctx *pulumi.Context,
	name string, args *FirewallProxyaddrgrpArgs, opts ...pulumi.ResourceOption) (*FirewallProxyaddrgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallProxyaddrgrp
	err := ctx.RegisterResource("fortios:index/firewallProxyaddrgrp:FirewallProxyaddrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProxyaddrgrp gets an existing FirewallProxyaddrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProxyaddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProxyaddrgrpState, opts ...pulumi.ResourceOption) (*FirewallProxyaddrgrp, error) {
	var resource FirewallProxyaddrgrp
	err := ctx.ReadResource("fortios:index/firewallProxyaddrgrp:FirewallProxyaddrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProxyaddrgrp resources.
type firewallProxyaddrgrpState struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Members of address group. The structure of `member` block is documented below.
	Members []FirewallProxyaddrgrpMember `pulumi:"members"`
	// Address group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyaddrgrpTagging `pulumi:"taggings"`
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallProxyaddrgrpState struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Members of address group. The structure of `member` block is documented below.
	Members FirewallProxyaddrgrpMemberArrayInput
	// Address group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyaddrgrpTaggingArrayInput
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyaddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyaddrgrpState)(nil)).Elem()
}

type firewallProxyaddrgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Members of address group. The structure of `member` block is documented below.
	Members []FirewallProxyaddrgrpMember `pulumi:"members"`
	// Address group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyaddrgrpTagging `pulumi:"taggings"`
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallProxyaddrgrp resource.
type FirewallProxyaddrgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Members of address group. The structure of `member` block is documented below.
	Members FirewallProxyaddrgrpMemberArrayInput
	// Address group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyaddrgrpTaggingArrayInput
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyaddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyaddrgrpArgs)(nil)).Elem()
}

type FirewallProxyaddrgrpInput interface {
	pulumi.Input

	ToFirewallProxyaddrgrpOutput() FirewallProxyaddrgrpOutput
	ToFirewallProxyaddrgrpOutputWithContext(ctx context.Context) FirewallProxyaddrgrpOutput
}

func (*FirewallProxyaddrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyaddrgrp)(nil)).Elem()
}

func (i *FirewallProxyaddrgrp) ToFirewallProxyaddrgrpOutput() FirewallProxyaddrgrpOutput {
	return i.ToFirewallProxyaddrgrpOutputWithContext(context.Background())
}

func (i *FirewallProxyaddrgrp) ToFirewallProxyaddrgrpOutputWithContext(ctx context.Context) FirewallProxyaddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyaddrgrpOutput)
}

// FirewallProxyaddrgrpArrayInput is an input type that accepts FirewallProxyaddrgrpArray and FirewallProxyaddrgrpArrayOutput values.
// You can construct a concrete instance of `FirewallProxyaddrgrpArrayInput` via:
//
//	FirewallProxyaddrgrpArray{ FirewallProxyaddrgrpArgs{...} }
type FirewallProxyaddrgrpArrayInput interface {
	pulumi.Input

	ToFirewallProxyaddrgrpArrayOutput() FirewallProxyaddrgrpArrayOutput
	ToFirewallProxyaddrgrpArrayOutputWithContext(context.Context) FirewallProxyaddrgrpArrayOutput
}

type FirewallProxyaddrgrpArray []FirewallProxyaddrgrpInput

func (FirewallProxyaddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxyaddrgrp)(nil)).Elem()
}

func (i FirewallProxyaddrgrpArray) ToFirewallProxyaddrgrpArrayOutput() FirewallProxyaddrgrpArrayOutput {
	return i.ToFirewallProxyaddrgrpArrayOutputWithContext(context.Background())
}

func (i FirewallProxyaddrgrpArray) ToFirewallProxyaddrgrpArrayOutputWithContext(ctx context.Context) FirewallProxyaddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyaddrgrpArrayOutput)
}

// FirewallProxyaddrgrpMapInput is an input type that accepts FirewallProxyaddrgrpMap and FirewallProxyaddrgrpMapOutput values.
// You can construct a concrete instance of `FirewallProxyaddrgrpMapInput` via:
//
//	FirewallProxyaddrgrpMap{ "key": FirewallProxyaddrgrpArgs{...} }
type FirewallProxyaddrgrpMapInput interface {
	pulumi.Input

	ToFirewallProxyaddrgrpMapOutput() FirewallProxyaddrgrpMapOutput
	ToFirewallProxyaddrgrpMapOutputWithContext(context.Context) FirewallProxyaddrgrpMapOutput
}

type FirewallProxyaddrgrpMap map[string]FirewallProxyaddrgrpInput

func (FirewallProxyaddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxyaddrgrp)(nil)).Elem()
}

func (i FirewallProxyaddrgrpMap) ToFirewallProxyaddrgrpMapOutput() FirewallProxyaddrgrpMapOutput {
	return i.ToFirewallProxyaddrgrpMapOutputWithContext(context.Background())
}

func (i FirewallProxyaddrgrpMap) ToFirewallProxyaddrgrpMapOutputWithContext(ctx context.Context) FirewallProxyaddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyaddrgrpMapOutput)
}

type FirewallProxyaddrgrpOutput struct{ *pulumi.OutputState }

func (FirewallProxyaddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyaddrgrp)(nil)).Elem()
}

func (o FirewallProxyaddrgrpOutput) ToFirewallProxyaddrgrpOutput() FirewallProxyaddrgrpOutput {
	return o
}

func (o FirewallProxyaddrgrpOutput) ToFirewallProxyaddrgrpOutputWithContext(ctx context.Context) FirewallProxyaddrgrpOutput {
	return o
}

// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
func (o FirewallProxyaddrgrpOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Optional comments.
func (o FirewallProxyaddrgrpOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallProxyaddrgrpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Members of address group. The structure of `member` block is documented below.
func (o FirewallProxyaddrgrpOutput) Members() FirewallProxyaddrgrpMemberArrayOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) FirewallProxyaddrgrpMemberArrayOutput { return v.Members }).(FirewallProxyaddrgrpMemberArrayOutput)
}

// Address group name.
func (o FirewallProxyaddrgrpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o FirewallProxyaddrgrpOutput) Taggings() FirewallProxyaddrgrpTaggingArrayOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) FirewallProxyaddrgrpTaggingArrayOutput { return v.Taggings }).(FirewallProxyaddrgrpTaggingArrayOutput)
}

// Source or destination address group type. Valid values: `src`, `dst`.
func (o FirewallProxyaddrgrpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallProxyaddrgrpOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallProxyaddrgrpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
func (o FirewallProxyaddrgrpOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddrgrp) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type FirewallProxyaddrgrpArrayOutput struct{ *pulumi.OutputState }

func (FirewallProxyaddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxyaddrgrp)(nil)).Elem()
}

func (o FirewallProxyaddrgrpArrayOutput) ToFirewallProxyaddrgrpArrayOutput() FirewallProxyaddrgrpArrayOutput {
	return o
}

func (o FirewallProxyaddrgrpArrayOutput) ToFirewallProxyaddrgrpArrayOutputWithContext(ctx context.Context) FirewallProxyaddrgrpArrayOutput {
	return o
}

func (o FirewallProxyaddrgrpArrayOutput) Index(i pulumi.IntInput) FirewallProxyaddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallProxyaddrgrp {
		return vs[0].([]*FirewallProxyaddrgrp)[vs[1].(int)]
	}).(FirewallProxyaddrgrpOutput)
}

type FirewallProxyaddrgrpMapOutput struct{ *pulumi.OutputState }

func (FirewallProxyaddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxyaddrgrp)(nil)).Elem()
}

func (o FirewallProxyaddrgrpMapOutput) ToFirewallProxyaddrgrpMapOutput() FirewallProxyaddrgrpMapOutput {
	return o
}

func (o FirewallProxyaddrgrpMapOutput) ToFirewallProxyaddrgrpMapOutputWithContext(ctx context.Context) FirewallProxyaddrgrpMapOutput {
	return o
}

func (o FirewallProxyaddrgrpMapOutput) MapIndex(k pulumi.StringInput) FirewallProxyaddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallProxyaddrgrp {
		return vs[0].(map[string]*FirewallProxyaddrgrp)[vs[1].(string)]
	}).(FirewallProxyaddrgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyaddrgrpInput)(nil)).Elem(), &FirewallProxyaddrgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyaddrgrpArrayInput)(nil)).Elem(), FirewallProxyaddrgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyaddrgrpMapInput)(nil)).Elem(), FirewallProxyaddrgrpMap{})
	pulumi.RegisterOutputType(FirewallProxyaddrgrpOutput{})
	pulumi.RegisterOutputType(FirewallProxyaddrgrpArrayOutput{})
	pulumi.RegisterOutputType(FirewallProxyaddrgrpMapOutput{})
}
