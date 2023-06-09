// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// # SystemReplacemsg Ec can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgEc:SystemreplacemsgEc labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgEc:SystemreplacemsgEc labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgEc struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag. Valid values: `none`, `text`, `html`, `wml`.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemreplacemsgEc registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgEc(ctx *pulumi.Context,
	name string, args *SystemreplacemsgEcArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgEc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgEc
	err := ctx.RegisterResource("fortios:index/systemreplacemsgEc:SystemreplacemsgEc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgEc gets an existing SystemreplacemsgEc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgEc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgEcState, opts ...pulumi.ResourceOption) (*SystemreplacemsgEc, error) {
	var resource SystemreplacemsgEc
	err := ctx.ReadResource("fortios:index/systemreplacemsgEc:SystemreplacemsgEc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgEc resources.
type systemreplacemsgEcState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag. Valid values: `none`, `text`, `html`, `wml`.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemreplacemsgEcState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag. Valid values: `none`, `text`, `html`, `wml`.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgEcState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgEcState)(nil)).Elem()
}

type systemreplacemsgEcArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag. Valid values: `none`, `text`, `html`, `wml`.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemreplacemsgEc resource.
type SystemreplacemsgEcArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag. Valid values: `none`, `text`, `html`, `wml`.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgEcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgEcArgs)(nil)).Elem()
}

type SystemreplacemsgEcInput interface {
	pulumi.Input

	ToSystemreplacemsgEcOutput() SystemreplacemsgEcOutput
	ToSystemreplacemsgEcOutputWithContext(ctx context.Context) SystemreplacemsgEcOutput
}

func (*SystemreplacemsgEc) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgEc)(nil)).Elem()
}

func (i *SystemreplacemsgEc) ToSystemreplacemsgEcOutput() SystemreplacemsgEcOutput {
	return i.ToSystemreplacemsgEcOutputWithContext(context.Background())
}

func (i *SystemreplacemsgEc) ToSystemreplacemsgEcOutputWithContext(ctx context.Context) SystemreplacemsgEcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgEcOutput)
}

// SystemreplacemsgEcArrayInput is an input type that accepts SystemreplacemsgEcArray and SystemreplacemsgEcArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgEcArrayInput` via:
//
//	SystemreplacemsgEcArray{ SystemreplacemsgEcArgs{...} }
type SystemreplacemsgEcArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgEcArrayOutput() SystemreplacemsgEcArrayOutput
	ToSystemreplacemsgEcArrayOutputWithContext(context.Context) SystemreplacemsgEcArrayOutput
}

type SystemreplacemsgEcArray []SystemreplacemsgEcInput

func (SystemreplacemsgEcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgEc)(nil)).Elem()
}

func (i SystemreplacemsgEcArray) ToSystemreplacemsgEcArrayOutput() SystemreplacemsgEcArrayOutput {
	return i.ToSystemreplacemsgEcArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgEcArray) ToSystemreplacemsgEcArrayOutputWithContext(ctx context.Context) SystemreplacemsgEcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgEcArrayOutput)
}

// SystemreplacemsgEcMapInput is an input type that accepts SystemreplacemsgEcMap and SystemreplacemsgEcMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgEcMapInput` via:
//
//	SystemreplacemsgEcMap{ "key": SystemreplacemsgEcArgs{...} }
type SystemreplacemsgEcMapInput interface {
	pulumi.Input

	ToSystemreplacemsgEcMapOutput() SystemreplacemsgEcMapOutput
	ToSystemreplacemsgEcMapOutputWithContext(context.Context) SystemreplacemsgEcMapOutput
}

type SystemreplacemsgEcMap map[string]SystemreplacemsgEcInput

func (SystemreplacemsgEcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgEc)(nil)).Elem()
}

func (i SystemreplacemsgEcMap) ToSystemreplacemsgEcMapOutput() SystemreplacemsgEcMapOutput {
	return i.ToSystemreplacemsgEcMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgEcMap) ToSystemreplacemsgEcMapOutputWithContext(ctx context.Context) SystemreplacemsgEcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgEcMapOutput)
}

type SystemreplacemsgEcOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgEcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgEc)(nil)).Elem()
}

func (o SystemreplacemsgEcOutput) ToSystemreplacemsgEcOutput() SystemreplacemsgEcOutput {
	return o
}

func (o SystemreplacemsgEcOutput) ToSystemreplacemsgEcOutputWithContext(ctx context.Context) SystemreplacemsgEcOutput {
	return o
}

// Message string.
func (o SystemreplacemsgEcOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgEc) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag. Valid values: `none`, `text`, `html`, `wml`.
func (o SystemreplacemsgEcOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgEc) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgEcOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgEc) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgEcOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgEc) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgEcOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgEc) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgEcArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgEcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgEc)(nil)).Elem()
}

func (o SystemreplacemsgEcArrayOutput) ToSystemreplacemsgEcArrayOutput() SystemreplacemsgEcArrayOutput {
	return o
}

func (o SystemreplacemsgEcArrayOutput) ToSystemreplacemsgEcArrayOutputWithContext(ctx context.Context) SystemreplacemsgEcArrayOutput {
	return o
}

func (o SystemreplacemsgEcArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgEcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgEc {
		return vs[0].([]*SystemreplacemsgEc)[vs[1].(int)]
	}).(SystemreplacemsgEcOutput)
}

type SystemreplacemsgEcMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgEcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgEc)(nil)).Elem()
}

func (o SystemreplacemsgEcMapOutput) ToSystemreplacemsgEcMapOutput() SystemreplacemsgEcMapOutput {
	return o
}

func (o SystemreplacemsgEcMapOutput) ToSystemreplacemsgEcMapOutputWithContext(ctx context.Context) SystemreplacemsgEcMapOutput {
	return o
}

func (o SystemreplacemsgEcMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgEcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgEc {
		return vs[0].(map[string]*SystemreplacemsgEc)[vs[1].(string)]
	}).(SystemreplacemsgEcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgEcInput)(nil)).Elem(), &SystemreplacemsgEc{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgEcArrayInput)(nil)).Elem(), SystemreplacemsgEcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgEcMapInput)(nil)).Elem(), SystemreplacemsgEcMap{})
	pulumi.RegisterOutputType(SystemreplacemsgEcOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgEcArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgEcMapOutput{})
}
