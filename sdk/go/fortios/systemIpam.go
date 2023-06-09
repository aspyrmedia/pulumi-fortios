// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP address management services. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// # System Ipam can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemIpam:SystemIpam labelname SystemIpam
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemIpam:SystemIpam labelname SystemIpam
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemIpam struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure IPAM pool subnet, Class A - Class B subnet.
	PoolSubnet pulumi.StringOutput `pulumi:"poolSubnet"`
	// Configure IPAM pools. The structure of `pools` block is documented below.
	Pools SystemIpamPoolArrayOutput `pulumi:"pools"`
	// Configure IPAM allocation rules. The structure of `rules` block is documented below.
	Rules SystemIpamRuleArrayOutput `pulumi:"rules"`
	// Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Enable/disable IP address management services. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemIpam registers a new resource with the given unique name, arguments, and options.
func NewSystemIpam(ctx *pulumi.Context,
	name string, args *SystemIpamArgs, opts ...pulumi.ResourceOption) (*SystemIpam, error) {
	if args == nil {
		args = &SystemIpamArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemIpam
	err := ctx.RegisterResource("fortios:index/systemIpam:SystemIpam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemIpam gets an existing SystemIpam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemIpam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemIpamState, opts ...pulumi.ResourceOption) (*SystemIpam, error) {
	var resource SystemIpam
	err := ctx.ReadResource("fortios:index/systemIpam:SystemIpam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemIpam resources.
type systemIpamState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure IPAM pool subnet, Class A - Class B subnet.
	PoolSubnet *string `pulumi:"poolSubnet"`
	// Configure IPAM pools. The structure of `pools` block is documented below.
	Pools []SystemIpamPool `pulumi:"pools"`
	// Configure IPAM allocation rules. The structure of `rules` block is documented below.
	Rules []SystemIpamRule `pulumi:"rules"`
	// Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
	ServerType *string `pulumi:"serverType"`
	// Enable/disable IP address management services. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemIpamState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure IPAM pool subnet, Class A - Class B subnet.
	PoolSubnet pulumi.StringPtrInput
	// Configure IPAM pools. The structure of `pools` block is documented below.
	Pools SystemIpamPoolArrayInput
	// Configure IPAM allocation rules. The structure of `rules` block is documented below.
	Rules SystemIpamRuleArrayInput
	// Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
	ServerType pulumi.StringPtrInput
	// Enable/disable IP address management services. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpamState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpamState)(nil)).Elem()
}

type systemIpamArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure IPAM pool subnet, Class A - Class B subnet.
	PoolSubnet *string `pulumi:"poolSubnet"`
	// Configure IPAM pools. The structure of `pools` block is documented below.
	Pools []SystemIpamPool `pulumi:"pools"`
	// Configure IPAM allocation rules. The structure of `rules` block is documented below.
	Rules []SystemIpamRule `pulumi:"rules"`
	// Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
	ServerType *string `pulumi:"serverType"`
	// Enable/disable IP address management services. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemIpam resource.
type SystemIpamArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure IPAM pool subnet, Class A - Class B subnet.
	PoolSubnet pulumi.StringPtrInput
	// Configure IPAM pools. The structure of `pools` block is documented below.
	Pools SystemIpamPoolArrayInput
	// Configure IPAM allocation rules. The structure of `rules` block is documented below.
	Rules SystemIpamRuleArrayInput
	// Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
	ServerType pulumi.StringPtrInput
	// Enable/disable IP address management services. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpamArgs)(nil)).Elem()
}

type SystemIpamInput interface {
	pulumi.Input

	ToSystemIpamOutput() SystemIpamOutput
	ToSystemIpamOutputWithContext(ctx context.Context) SystemIpamOutput
}

func (*SystemIpam) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpam)(nil)).Elem()
}

func (i *SystemIpam) ToSystemIpamOutput() SystemIpamOutput {
	return i.ToSystemIpamOutputWithContext(context.Background())
}

func (i *SystemIpam) ToSystemIpamOutputWithContext(ctx context.Context) SystemIpamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpamOutput)
}

// SystemIpamArrayInput is an input type that accepts SystemIpamArray and SystemIpamArrayOutput values.
// You can construct a concrete instance of `SystemIpamArrayInput` via:
//
//	SystemIpamArray{ SystemIpamArgs{...} }
type SystemIpamArrayInput interface {
	pulumi.Input

	ToSystemIpamArrayOutput() SystemIpamArrayOutput
	ToSystemIpamArrayOutputWithContext(context.Context) SystemIpamArrayOutput
}

type SystemIpamArray []SystemIpamInput

func (SystemIpamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpam)(nil)).Elem()
}

func (i SystemIpamArray) ToSystemIpamArrayOutput() SystemIpamArrayOutput {
	return i.ToSystemIpamArrayOutputWithContext(context.Background())
}

func (i SystemIpamArray) ToSystemIpamArrayOutputWithContext(ctx context.Context) SystemIpamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpamArrayOutput)
}

// SystemIpamMapInput is an input type that accepts SystemIpamMap and SystemIpamMapOutput values.
// You can construct a concrete instance of `SystemIpamMapInput` via:
//
//	SystemIpamMap{ "key": SystemIpamArgs{...} }
type SystemIpamMapInput interface {
	pulumi.Input

	ToSystemIpamMapOutput() SystemIpamMapOutput
	ToSystemIpamMapOutputWithContext(context.Context) SystemIpamMapOutput
}

type SystemIpamMap map[string]SystemIpamInput

func (SystemIpamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpam)(nil)).Elem()
}

func (i SystemIpamMap) ToSystemIpamMapOutput() SystemIpamMapOutput {
	return i.ToSystemIpamMapOutputWithContext(context.Background())
}

func (i SystemIpamMap) ToSystemIpamMapOutputWithContext(ctx context.Context) SystemIpamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpamMapOutput)
}

type SystemIpamOutput struct{ *pulumi.OutputState }

func (SystemIpamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpam)(nil)).Elem()
}

func (o SystemIpamOutput) ToSystemIpamOutput() SystemIpamOutput {
	return o
}

func (o SystemIpamOutput) ToSystemIpamOutputWithContext(ctx context.Context) SystemIpamOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemIpamOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemIpam) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Configure IPAM pool subnet, Class A - Class B subnet.
func (o SystemIpamOutput) PoolSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpam) pulumi.StringOutput { return v.PoolSubnet }).(pulumi.StringOutput)
}

// Configure IPAM pools. The structure of `pools` block is documented below.
func (o SystemIpamOutput) Pools() SystemIpamPoolArrayOutput {
	return o.ApplyT(func(v *SystemIpam) SystemIpamPoolArrayOutput { return v.Pools }).(SystemIpamPoolArrayOutput)
}

// Configure IPAM allocation rules. The structure of `rules` block is documented below.
func (o SystemIpamOutput) Rules() SystemIpamRuleArrayOutput {
	return o.ApplyT(func(v *SystemIpam) SystemIpamRuleArrayOutput { return v.Rules }).(SystemIpamRuleArrayOutput)
}

// Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
func (o SystemIpamOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpam) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

// Enable/disable IP address management services. Valid values: `enable`, `disable`.
func (o SystemIpamOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpam) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemIpamOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemIpam) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemIpamArrayOutput struct{ *pulumi.OutputState }

func (SystemIpamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpam)(nil)).Elem()
}

func (o SystemIpamArrayOutput) ToSystemIpamArrayOutput() SystemIpamArrayOutput {
	return o
}

func (o SystemIpamArrayOutput) ToSystemIpamArrayOutputWithContext(ctx context.Context) SystemIpamArrayOutput {
	return o
}

func (o SystemIpamArrayOutput) Index(i pulumi.IntInput) SystemIpamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemIpam {
		return vs[0].([]*SystemIpam)[vs[1].(int)]
	}).(SystemIpamOutput)
}

type SystemIpamMapOutput struct{ *pulumi.OutputState }

func (SystemIpamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpam)(nil)).Elem()
}

func (o SystemIpamMapOutput) ToSystemIpamMapOutput() SystemIpamMapOutput {
	return o
}

func (o SystemIpamMapOutput) ToSystemIpamMapOutputWithContext(ctx context.Context) SystemIpamMapOutput {
	return o
}

func (o SystemIpamMapOutput) MapIndex(k pulumi.StringInput) SystemIpamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemIpam {
		return vs[0].(map[string]*SystemIpam)[vs[1].(string)]
	}).(SystemIpamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpamInput)(nil)).Elem(), &SystemIpam{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpamArrayInput)(nil)).Elem(), SystemIpamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpamMapInput)(nil)).Elem(), SystemIpamMap{})
	pulumi.RegisterOutputType(SystemIpamOutput{})
	pulumi.RegisterOutputType(SystemIpamArrayOutput{})
	pulumi.RegisterOutputType(SystemIpamMapOutput{})
}
