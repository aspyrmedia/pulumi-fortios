// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define geoip country name-ID table. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # System GeoipCountry can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemGeoipcountry:SystemGeoipcountry labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemGeoipcountry:SystemGeoipcountry labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemGeoipcountry struct {
	pulumi.CustomResourceState

	// Country ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Country name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemGeoipcountry registers a new resource with the given unique name, arguments, and options.
func NewSystemGeoipcountry(ctx *pulumi.Context,
	name string, args *SystemGeoipcountryArgs, opts ...pulumi.ResourceOption) (*SystemGeoipcountry, error) {
	if args == nil {
		args = &SystemGeoipcountryArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemGeoipcountry
	err := ctx.RegisterResource("fortios:index/systemGeoipcountry:SystemGeoipcountry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGeoipcountry gets an existing SystemGeoipcountry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGeoipcountry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGeoipcountryState, opts ...pulumi.ResourceOption) (*SystemGeoipcountry, error) {
	var resource SystemGeoipcountry
	err := ctx.ReadResource("fortios:index/systemGeoipcountry:SystemGeoipcountry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGeoipcountry resources.
type systemGeoipcountryState struct {
	// Country ID.
	Fosid *string `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemGeoipcountryState struct {
	// Country ID.
	Fosid pulumi.StringPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGeoipcountryState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeoipcountryState)(nil)).Elem()
}

type systemGeoipcountryArgs struct {
	// Country ID.
	Fosid *string `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemGeoipcountry resource.
type SystemGeoipcountryArgs struct {
	// Country ID.
	Fosid pulumi.StringPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGeoipcountryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGeoipcountryArgs)(nil)).Elem()
}

type SystemGeoipcountryInput interface {
	pulumi.Input

	ToSystemGeoipcountryOutput() SystemGeoipcountryOutput
	ToSystemGeoipcountryOutputWithContext(ctx context.Context) SystemGeoipcountryOutput
}

func (*SystemGeoipcountry) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeoipcountry)(nil)).Elem()
}

func (i *SystemGeoipcountry) ToSystemGeoipcountryOutput() SystemGeoipcountryOutput {
	return i.ToSystemGeoipcountryOutputWithContext(context.Background())
}

func (i *SystemGeoipcountry) ToSystemGeoipcountryOutputWithContext(ctx context.Context) SystemGeoipcountryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipcountryOutput)
}

// SystemGeoipcountryArrayInput is an input type that accepts SystemGeoipcountryArray and SystemGeoipcountryArrayOutput values.
// You can construct a concrete instance of `SystemGeoipcountryArrayInput` via:
//
//	SystemGeoipcountryArray{ SystemGeoipcountryArgs{...} }
type SystemGeoipcountryArrayInput interface {
	pulumi.Input

	ToSystemGeoipcountryArrayOutput() SystemGeoipcountryArrayOutput
	ToSystemGeoipcountryArrayOutputWithContext(context.Context) SystemGeoipcountryArrayOutput
}

type SystemGeoipcountryArray []SystemGeoipcountryInput

func (SystemGeoipcountryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGeoipcountry)(nil)).Elem()
}

func (i SystemGeoipcountryArray) ToSystemGeoipcountryArrayOutput() SystemGeoipcountryArrayOutput {
	return i.ToSystemGeoipcountryArrayOutputWithContext(context.Background())
}

func (i SystemGeoipcountryArray) ToSystemGeoipcountryArrayOutputWithContext(ctx context.Context) SystemGeoipcountryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipcountryArrayOutput)
}

// SystemGeoipcountryMapInput is an input type that accepts SystemGeoipcountryMap and SystemGeoipcountryMapOutput values.
// You can construct a concrete instance of `SystemGeoipcountryMapInput` via:
//
//	SystemGeoipcountryMap{ "key": SystemGeoipcountryArgs{...} }
type SystemGeoipcountryMapInput interface {
	pulumi.Input

	ToSystemGeoipcountryMapOutput() SystemGeoipcountryMapOutput
	ToSystemGeoipcountryMapOutputWithContext(context.Context) SystemGeoipcountryMapOutput
}

type SystemGeoipcountryMap map[string]SystemGeoipcountryInput

func (SystemGeoipcountryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGeoipcountry)(nil)).Elem()
}

func (i SystemGeoipcountryMap) ToSystemGeoipcountryMapOutput() SystemGeoipcountryMapOutput {
	return i.ToSystemGeoipcountryMapOutputWithContext(context.Background())
}

func (i SystemGeoipcountryMap) ToSystemGeoipcountryMapOutputWithContext(ctx context.Context) SystemGeoipcountryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGeoipcountryMapOutput)
}

type SystemGeoipcountryOutput struct{ *pulumi.OutputState }

func (SystemGeoipcountryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGeoipcountry)(nil)).Elem()
}

func (o SystemGeoipcountryOutput) ToSystemGeoipcountryOutput() SystemGeoipcountryOutput {
	return o
}

func (o SystemGeoipcountryOutput) ToSystemGeoipcountryOutputWithContext(ctx context.Context) SystemGeoipcountryOutput {
	return o
}

// Country ID.
func (o SystemGeoipcountryOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeoipcountry) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// Country name.
func (o SystemGeoipcountryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGeoipcountry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemGeoipcountryOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemGeoipcountry) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemGeoipcountryArrayOutput struct{ *pulumi.OutputState }

func (SystemGeoipcountryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGeoipcountry)(nil)).Elem()
}

func (o SystemGeoipcountryArrayOutput) ToSystemGeoipcountryArrayOutput() SystemGeoipcountryArrayOutput {
	return o
}

func (o SystemGeoipcountryArrayOutput) ToSystemGeoipcountryArrayOutputWithContext(ctx context.Context) SystemGeoipcountryArrayOutput {
	return o
}

func (o SystemGeoipcountryArrayOutput) Index(i pulumi.IntInput) SystemGeoipcountryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemGeoipcountry {
		return vs[0].([]*SystemGeoipcountry)[vs[1].(int)]
	}).(SystemGeoipcountryOutput)
}

type SystemGeoipcountryMapOutput struct{ *pulumi.OutputState }

func (SystemGeoipcountryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGeoipcountry)(nil)).Elem()
}

func (o SystemGeoipcountryMapOutput) ToSystemGeoipcountryMapOutput() SystemGeoipcountryMapOutput {
	return o
}

func (o SystemGeoipcountryMapOutput) ToSystemGeoipcountryMapOutputWithContext(ctx context.Context) SystemGeoipcountryMapOutput {
	return o
}

func (o SystemGeoipcountryMapOutput) MapIndex(k pulumi.StringInput) SystemGeoipcountryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemGeoipcountry {
		return vs[0].(map[string]*SystemGeoipcountry)[vs[1].(string)]
	}).(SystemGeoipcountryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeoipcountryInput)(nil)).Elem(), &SystemGeoipcountry{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeoipcountryArrayInput)(nil)).Elem(), SystemGeoipcountryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGeoipcountryMapInput)(nil)).Elem(), SystemGeoipcountryMap{})
	pulumi.RegisterOutputType(SystemGeoipcountryOutput{})
	pulumi.RegisterOutputType(SystemGeoipcountryArrayOutput{})
	pulumi.RegisterOutputType(SystemGeoipcountryMapOutput{})
}
