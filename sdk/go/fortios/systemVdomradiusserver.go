// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.
//
// ## Import
//
// # System VdomRadiusServer can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemVdomradiusserver:SystemVdomradiusserver labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemVdomradiusserver:SystemVdomradiusserver labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemVdomradiusserver struct {
	pulumi.CustomResourceState

	// Name of the VDOM that you are adding the RADIUS server to.
	Name pulumi.StringOutput `pulumi:"name"`
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom pulumi.StringOutput `pulumi:"radiusServerVdom"`
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomradiusserver registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomradiusserver(ctx *pulumi.Context,
	name string, args *SystemVdomradiusserverArgs, opts ...pulumi.ResourceOption) (*SystemVdomradiusserver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RadiusServerVdom == nil {
		return nil, errors.New("invalid value for required argument 'RadiusServerVdom'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVdomradiusserver
	err := ctx.RegisterResource("fortios:index/systemVdomradiusserver:SystemVdomradiusserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomradiusserver gets an existing SystemVdomradiusserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomradiusserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomradiusserverState, opts ...pulumi.ResourceOption) (*SystemVdomradiusserver, error) {
	var resource SystemVdomradiusserver
	err := ctx.ReadResource("fortios:index/systemVdomradiusserver:SystemVdomradiusserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomradiusserver resources.
type systemVdomradiusserverState struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name *string `pulumi:"name"`
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom *string `pulumi:"radiusServerVdom"`
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdomradiusserverState struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name pulumi.StringPtrInput
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom pulumi.StringPtrInput
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomradiusserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomradiusserverState)(nil)).Elem()
}

type systemVdomradiusserverArgs struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name *string `pulumi:"name"`
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom string `pulumi:"radiusServerVdom"`
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomradiusserver resource.
type SystemVdomradiusserverArgs struct {
	// Name of the VDOM that you are adding the RADIUS server to.
	Name pulumi.StringPtrInput
	// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
	RadiusServerVdom pulumi.StringInput
	// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomradiusserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomradiusserverArgs)(nil)).Elem()
}

type SystemVdomradiusserverInput interface {
	pulumi.Input

	ToSystemVdomradiusserverOutput() SystemVdomradiusserverOutput
	ToSystemVdomradiusserverOutputWithContext(ctx context.Context) SystemVdomradiusserverOutput
}

func (*SystemVdomradiusserver) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomradiusserver)(nil)).Elem()
}

func (i *SystemVdomradiusserver) ToSystemVdomradiusserverOutput() SystemVdomradiusserverOutput {
	return i.ToSystemVdomradiusserverOutputWithContext(context.Background())
}

func (i *SystemVdomradiusserver) ToSystemVdomradiusserverOutputWithContext(ctx context.Context) SystemVdomradiusserverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomradiusserverOutput)
}

// SystemVdomradiusserverArrayInput is an input type that accepts SystemVdomradiusserverArray and SystemVdomradiusserverArrayOutput values.
// You can construct a concrete instance of `SystemVdomradiusserverArrayInput` via:
//
//	SystemVdomradiusserverArray{ SystemVdomradiusserverArgs{...} }
type SystemVdomradiusserverArrayInput interface {
	pulumi.Input

	ToSystemVdomradiusserverArrayOutput() SystemVdomradiusserverArrayOutput
	ToSystemVdomradiusserverArrayOutputWithContext(context.Context) SystemVdomradiusserverArrayOutput
}

type SystemVdomradiusserverArray []SystemVdomradiusserverInput

func (SystemVdomradiusserverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomradiusserver)(nil)).Elem()
}

func (i SystemVdomradiusserverArray) ToSystemVdomradiusserverArrayOutput() SystemVdomradiusserverArrayOutput {
	return i.ToSystemVdomradiusserverArrayOutputWithContext(context.Background())
}

func (i SystemVdomradiusserverArray) ToSystemVdomradiusserverArrayOutputWithContext(ctx context.Context) SystemVdomradiusserverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomradiusserverArrayOutput)
}

// SystemVdomradiusserverMapInput is an input type that accepts SystemVdomradiusserverMap and SystemVdomradiusserverMapOutput values.
// You can construct a concrete instance of `SystemVdomradiusserverMapInput` via:
//
//	SystemVdomradiusserverMap{ "key": SystemVdomradiusserverArgs{...} }
type SystemVdomradiusserverMapInput interface {
	pulumi.Input

	ToSystemVdomradiusserverMapOutput() SystemVdomradiusserverMapOutput
	ToSystemVdomradiusserverMapOutputWithContext(context.Context) SystemVdomradiusserverMapOutput
}

type SystemVdomradiusserverMap map[string]SystemVdomradiusserverInput

func (SystemVdomradiusserverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomradiusserver)(nil)).Elem()
}

func (i SystemVdomradiusserverMap) ToSystemVdomradiusserverMapOutput() SystemVdomradiusserverMapOutput {
	return i.ToSystemVdomradiusserverMapOutputWithContext(context.Background())
}

func (i SystemVdomradiusserverMap) ToSystemVdomradiusserverMapOutputWithContext(ctx context.Context) SystemVdomradiusserverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomradiusserverMapOutput)
}

type SystemVdomradiusserverOutput struct{ *pulumi.OutputState }

func (SystemVdomradiusserverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomradiusserver)(nil)).Elem()
}

func (o SystemVdomradiusserverOutput) ToSystemVdomradiusserverOutput() SystemVdomradiusserverOutput {
	return o
}

func (o SystemVdomradiusserverOutput) ToSystemVdomradiusserverOutputWithContext(ctx context.Context) SystemVdomradiusserverOutput {
	return o
}

// Name of the VDOM that you are adding the RADIUS server to.
func (o SystemVdomradiusserverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomradiusserver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
func (o SystemVdomradiusserverOutput) RadiusServerVdom() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomradiusserver) pulumi.StringOutput { return v.RadiusServerVdom }).(pulumi.StringOutput)
}

// Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
func (o SystemVdomradiusserverOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomradiusserver) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemVdomradiusserverOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVdomradiusserver) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemVdomradiusserverArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomradiusserverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomradiusserver)(nil)).Elem()
}

func (o SystemVdomradiusserverArrayOutput) ToSystemVdomradiusserverArrayOutput() SystemVdomradiusserverArrayOutput {
	return o
}

func (o SystemVdomradiusserverArrayOutput) ToSystemVdomradiusserverArrayOutputWithContext(ctx context.Context) SystemVdomradiusserverArrayOutput {
	return o
}

func (o SystemVdomradiusserverArrayOutput) Index(i pulumi.IntInput) SystemVdomradiusserverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomradiusserver {
		return vs[0].([]*SystemVdomradiusserver)[vs[1].(int)]
	}).(SystemVdomradiusserverOutput)
}

type SystemVdomradiusserverMapOutput struct{ *pulumi.OutputState }

func (SystemVdomradiusserverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomradiusserver)(nil)).Elem()
}

func (o SystemVdomradiusserverMapOutput) ToSystemVdomradiusserverMapOutput() SystemVdomradiusserverMapOutput {
	return o
}

func (o SystemVdomradiusserverMapOutput) ToSystemVdomradiusserverMapOutputWithContext(ctx context.Context) SystemVdomradiusserverMapOutput {
	return o
}

func (o SystemVdomradiusserverMapOutput) MapIndex(k pulumi.StringInput) SystemVdomradiusserverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomradiusserver {
		return vs[0].(map[string]*SystemVdomradiusserver)[vs[1].(string)]
	}).(SystemVdomradiusserverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomradiusserverInput)(nil)).Elem(), &SystemVdomradiusserver{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomradiusserverArrayInput)(nil)).Elem(), SystemVdomradiusserverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomradiusserverMapInput)(nil)).Elem(), SystemVdomradiusserverMap{})
	pulumi.RegisterOutputType(SystemVdomradiusserverOutput{})
	pulumi.RegisterOutputType(SystemVdomradiusserverArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomradiusserverMapOutput{})
}
