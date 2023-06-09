// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `>= 7.2.0`.
//
// ## Import
//
// # Icap ServerGroup can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/icapServergroup:IcapServergroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/icapServergroup:IcapServergroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type IcapServergroup struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringOutput `pulumi:"name"`
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists IcapServergroupServerListArrayOutput `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIcapServergroup registers a new resource with the given unique name, arguments, and options.
func NewIcapServergroup(ctx *pulumi.Context,
	name string, args *IcapServergroupArgs, opts ...pulumi.ResourceOption) (*IcapServergroup, error) {
	if args == nil {
		args = &IcapServergroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IcapServergroup
	err := ctx.RegisterResource("fortios:index/icapServergroup:IcapServergroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIcapServergroup gets an existing IcapServergroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIcapServergroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IcapServergroupState, opts ...pulumi.ResourceOption) (*IcapServergroup, error) {
	var resource IcapServergroup
	err := ctx.ReadResource("fortios:index/icapServergroup:IcapServergroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IcapServergroup resources.
type icapServergroupState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []IcapServergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IcapServergroupState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod pulumi.StringPtrInput
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists IcapServergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IcapServergroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*icapServergroupState)(nil)).Elem()
}

type icapServergroupArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []IcapServergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IcapServergroup resource.
type IcapServergroupArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod pulumi.StringPtrInput
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists IcapServergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IcapServergroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*icapServergroupArgs)(nil)).Elem()
}

type IcapServergroupInput interface {
	pulumi.Input

	ToIcapServergroupOutput() IcapServergroupOutput
	ToIcapServergroupOutputWithContext(ctx context.Context) IcapServergroupOutput
}

func (*IcapServergroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IcapServergroup)(nil)).Elem()
}

func (i *IcapServergroup) ToIcapServergroupOutput() IcapServergroupOutput {
	return i.ToIcapServergroupOutputWithContext(context.Background())
}

func (i *IcapServergroup) ToIcapServergroupOutputWithContext(ctx context.Context) IcapServergroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IcapServergroupOutput)
}

// IcapServergroupArrayInput is an input type that accepts IcapServergroupArray and IcapServergroupArrayOutput values.
// You can construct a concrete instance of `IcapServergroupArrayInput` via:
//
//	IcapServergroupArray{ IcapServergroupArgs{...} }
type IcapServergroupArrayInput interface {
	pulumi.Input

	ToIcapServergroupArrayOutput() IcapServergroupArrayOutput
	ToIcapServergroupArrayOutputWithContext(context.Context) IcapServergroupArrayOutput
}

type IcapServergroupArray []IcapServergroupInput

func (IcapServergroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IcapServergroup)(nil)).Elem()
}

func (i IcapServergroupArray) ToIcapServergroupArrayOutput() IcapServergroupArrayOutput {
	return i.ToIcapServergroupArrayOutputWithContext(context.Background())
}

func (i IcapServergroupArray) ToIcapServergroupArrayOutputWithContext(ctx context.Context) IcapServergroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IcapServergroupArrayOutput)
}

// IcapServergroupMapInput is an input type that accepts IcapServergroupMap and IcapServergroupMapOutput values.
// You can construct a concrete instance of `IcapServergroupMapInput` via:
//
//	IcapServergroupMap{ "key": IcapServergroupArgs{...} }
type IcapServergroupMapInput interface {
	pulumi.Input

	ToIcapServergroupMapOutput() IcapServergroupMapOutput
	ToIcapServergroupMapOutputWithContext(context.Context) IcapServergroupMapOutput
}

type IcapServergroupMap map[string]IcapServergroupInput

func (IcapServergroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IcapServergroup)(nil)).Elem()
}

func (i IcapServergroupMap) ToIcapServergroupMapOutput() IcapServergroupMapOutput {
	return i.ToIcapServergroupMapOutputWithContext(context.Background())
}

func (i IcapServergroupMap) ToIcapServergroupMapOutputWithContext(ctx context.Context) IcapServergroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IcapServergroupMapOutput)
}

type IcapServergroupOutput struct{ *pulumi.OutputState }

func (IcapServergroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IcapServergroup)(nil)).Elem()
}

func (o IcapServergroupOutput) ToIcapServergroupOutput() IcapServergroupOutput {
	return o
}

func (o IcapServergroupOutput) ToIcapServergroupOutputWithContext(ctx context.Context) IcapServergroupOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o IcapServergroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IcapServergroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
func (o IcapServergroupOutput) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *IcapServergroup) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
func (o IcapServergroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IcapServergroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
func (o IcapServergroupOutput) ServerLists() IcapServergroupServerListArrayOutput {
	return o.ApplyT(func(v *IcapServergroup) IcapServergroupServerListArrayOutput { return v.ServerLists }).(IcapServergroupServerListArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IcapServergroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IcapServergroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IcapServergroupArrayOutput struct{ *pulumi.OutputState }

func (IcapServergroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IcapServergroup)(nil)).Elem()
}

func (o IcapServergroupArrayOutput) ToIcapServergroupArrayOutput() IcapServergroupArrayOutput {
	return o
}

func (o IcapServergroupArrayOutput) ToIcapServergroupArrayOutputWithContext(ctx context.Context) IcapServergroupArrayOutput {
	return o
}

func (o IcapServergroupArrayOutput) Index(i pulumi.IntInput) IcapServergroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IcapServergroup {
		return vs[0].([]*IcapServergroup)[vs[1].(int)]
	}).(IcapServergroupOutput)
}

type IcapServergroupMapOutput struct{ *pulumi.OutputState }

func (IcapServergroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IcapServergroup)(nil)).Elem()
}

func (o IcapServergroupMapOutput) ToIcapServergroupMapOutput() IcapServergroupMapOutput {
	return o
}

func (o IcapServergroupMapOutput) ToIcapServergroupMapOutputWithContext(ctx context.Context) IcapServergroupMapOutput {
	return o
}

func (o IcapServergroupMapOutput) MapIndex(k pulumi.StringInput) IcapServergroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IcapServergroup {
		return vs[0].(map[string]*IcapServergroup)[vs[1].(string)]
	}).(IcapServergroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IcapServergroupInput)(nil)).Elem(), &IcapServergroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*IcapServergroupArrayInput)(nil)).Elem(), IcapServergroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IcapServergroupMapInput)(nil)).Elem(), IcapServergroupMap{})
	pulumi.RegisterOutputType(IcapServergroupOutput{})
	pulumi.RegisterOutputType(IcapServergroupArrayOutput{})
	pulumi.RegisterOutputType(IcapServergroupMapOutput{})
}
