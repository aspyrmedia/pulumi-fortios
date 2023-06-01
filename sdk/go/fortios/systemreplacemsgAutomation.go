// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// # SystemReplacemsg Automation can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgAutomation:SystemreplacemsgAutomation labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgAutomation:SystemreplacemsgAutomation labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgAutomation struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag. Valid values: `none`, `text`, `html`.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemreplacemsgAutomation registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgAutomation(ctx *pulumi.Context,
	name string, args *SystemreplacemsgAutomationArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgAutomation, error) {
	if args == nil {
		args = &SystemreplacemsgAutomationArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgAutomation
	err := ctx.RegisterResource("fortios:index/systemreplacemsgAutomation:SystemreplacemsgAutomation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgAutomation gets an existing SystemreplacemsgAutomation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgAutomation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgAutomationState, opts ...pulumi.ResourceOption) (*SystemreplacemsgAutomation, error) {
	var resource SystemreplacemsgAutomation
	err := ctx.ReadResource("fortios:index/systemreplacemsgAutomation:SystemreplacemsgAutomation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgAutomation resources.
type systemreplacemsgAutomationState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag. Valid values: `none`, `text`, `html`.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemreplacemsgAutomationState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag. Valid values: `none`, `text`, `html`.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgAutomationState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgAutomationState)(nil)).Elem()
}

type systemreplacemsgAutomationArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag. Valid values: `none`, `text`, `html`.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemreplacemsgAutomation resource.
type SystemreplacemsgAutomationArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag. Valid values: `none`, `text`, `html`.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgAutomationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgAutomationArgs)(nil)).Elem()
}

type SystemreplacemsgAutomationInput interface {
	pulumi.Input

	ToSystemreplacemsgAutomationOutput() SystemreplacemsgAutomationOutput
	ToSystemreplacemsgAutomationOutputWithContext(ctx context.Context) SystemreplacemsgAutomationOutput
}

func (*SystemreplacemsgAutomation) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgAutomation)(nil)).Elem()
}

func (i *SystemreplacemsgAutomation) ToSystemreplacemsgAutomationOutput() SystemreplacemsgAutomationOutput {
	return i.ToSystemreplacemsgAutomationOutputWithContext(context.Background())
}

func (i *SystemreplacemsgAutomation) ToSystemreplacemsgAutomationOutputWithContext(ctx context.Context) SystemreplacemsgAutomationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgAutomationOutput)
}

// SystemreplacemsgAutomationArrayInput is an input type that accepts SystemreplacemsgAutomationArray and SystemreplacemsgAutomationArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgAutomationArrayInput` via:
//
//	SystemreplacemsgAutomationArray{ SystemreplacemsgAutomationArgs{...} }
type SystemreplacemsgAutomationArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgAutomationArrayOutput() SystemreplacemsgAutomationArrayOutput
	ToSystemreplacemsgAutomationArrayOutputWithContext(context.Context) SystemreplacemsgAutomationArrayOutput
}

type SystemreplacemsgAutomationArray []SystemreplacemsgAutomationInput

func (SystemreplacemsgAutomationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgAutomation)(nil)).Elem()
}

func (i SystemreplacemsgAutomationArray) ToSystemreplacemsgAutomationArrayOutput() SystemreplacemsgAutomationArrayOutput {
	return i.ToSystemreplacemsgAutomationArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgAutomationArray) ToSystemreplacemsgAutomationArrayOutputWithContext(ctx context.Context) SystemreplacemsgAutomationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgAutomationArrayOutput)
}

// SystemreplacemsgAutomationMapInput is an input type that accepts SystemreplacemsgAutomationMap and SystemreplacemsgAutomationMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgAutomationMapInput` via:
//
//	SystemreplacemsgAutomationMap{ "key": SystemreplacemsgAutomationArgs{...} }
type SystemreplacemsgAutomationMapInput interface {
	pulumi.Input

	ToSystemreplacemsgAutomationMapOutput() SystemreplacemsgAutomationMapOutput
	ToSystemreplacemsgAutomationMapOutputWithContext(context.Context) SystemreplacemsgAutomationMapOutput
}

type SystemreplacemsgAutomationMap map[string]SystemreplacemsgAutomationInput

func (SystemreplacemsgAutomationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgAutomation)(nil)).Elem()
}

func (i SystemreplacemsgAutomationMap) ToSystemreplacemsgAutomationMapOutput() SystemreplacemsgAutomationMapOutput {
	return i.ToSystemreplacemsgAutomationMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgAutomationMap) ToSystemreplacemsgAutomationMapOutputWithContext(ctx context.Context) SystemreplacemsgAutomationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgAutomationMapOutput)
}

type SystemreplacemsgAutomationOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgAutomationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgAutomation)(nil)).Elem()
}

func (o SystemreplacemsgAutomationOutput) ToSystemreplacemsgAutomationOutput() SystemreplacemsgAutomationOutput {
	return o
}

func (o SystemreplacemsgAutomationOutput) ToSystemreplacemsgAutomationOutputWithContext(ctx context.Context) SystemreplacemsgAutomationOutput {
	return o
}

// Message string.
func (o SystemreplacemsgAutomationOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgAutomation) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag. Valid values: `none`, `text`, `html`.
func (o SystemreplacemsgAutomationOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgAutomation) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgAutomationOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgAutomation) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgAutomationOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgAutomation) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgAutomationOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgAutomation) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgAutomationArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgAutomationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgAutomation)(nil)).Elem()
}

func (o SystemreplacemsgAutomationArrayOutput) ToSystemreplacemsgAutomationArrayOutput() SystemreplacemsgAutomationArrayOutput {
	return o
}

func (o SystemreplacemsgAutomationArrayOutput) ToSystemreplacemsgAutomationArrayOutputWithContext(ctx context.Context) SystemreplacemsgAutomationArrayOutput {
	return o
}

func (o SystemreplacemsgAutomationArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgAutomationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgAutomation {
		return vs[0].([]*SystemreplacemsgAutomation)[vs[1].(int)]
	}).(SystemreplacemsgAutomationOutput)
}

type SystemreplacemsgAutomationMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgAutomationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgAutomation)(nil)).Elem()
}

func (o SystemreplacemsgAutomationMapOutput) ToSystemreplacemsgAutomationMapOutput() SystemreplacemsgAutomationMapOutput {
	return o
}

func (o SystemreplacemsgAutomationMapOutput) ToSystemreplacemsgAutomationMapOutputWithContext(ctx context.Context) SystemreplacemsgAutomationMapOutput {
	return o
}

func (o SystemreplacemsgAutomationMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgAutomationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgAutomation {
		return vs[0].(map[string]*SystemreplacemsgAutomation)[vs[1].(string)]
	}).(SystemreplacemsgAutomationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgAutomationInput)(nil)).Elem(), &SystemreplacemsgAutomation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgAutomationArrayInput)(nil)).Elem(), SystemreplacemsgAutomationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgAutomationMapInput)(nil)).Elem(), SystemreplacemsgAutomationMap{})
	pulumi.RegisterOutputType(SystemreplacemsgAutomationOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgAutomationArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgAutomationMapOutput{})
}