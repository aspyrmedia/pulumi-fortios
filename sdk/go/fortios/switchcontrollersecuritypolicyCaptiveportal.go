// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Names of VLANs that use captive portal authentication. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// # SwitchControllerSecurityPolicy CaptivePortal can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollersecuritypolicyCaptiveportal:SwitchcontrollersecuritypolicyCaptiveportal labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollersecuritypolicyCaptiveportal:SwitchcontrollersecuritypolicyCaptiveportal labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollersecuritypolicyCaptiveportal struct {
	pulumi.CustomResourceState

	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy type. Valid values: `captive-portal`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Names of VLANs that use captive portal authentication.
	Vlan pulumi.StringOutput `pulumi:"vlan"`
}

// NewSwitchcontrollersecuritypolicyCaptiveportal registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollersecuritypolicyCaptiveportal(ctx *pulumi.Context,
	name string, args *SwitchcontrollersecuritypolicyCaptiveportalArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollersecuritypolicyCaptiveportal, error) {
	if args == nil {
		args = &SwitchcontrollersecuritypolicyCaptiveportalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollersecuritypolicyCaptiveportal
	err := ctx.RegisterResource("fortios:index/switchcontrollersecuritypolicyCaptiveportal:SwitchcontrollersecuritypolicyCaptiveportal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollersecuritypolicyCaptiveportal gets an existing SwitchcontrollersecuritypolicyCaptiveportal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollersecuritypolicyCaptiveportal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollersecuritypolicyCaptiveportalState, opts ...pulumi.ResourceOption) (*SwitchcontrollersecuritypolicyCaptiveportal, error) {
	var resource SwitchcontrollersecuritypolicyCaptiveportal
	err := ctx.ReadResource("fortios:index/switchcontrollersecuritypolicyCaptiveportal:SwitchcontrollersecuritypolicyCaptiveportal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollersecuritypolicyCaptiveportal resources.
type switchcontrollersecuritypolicyCaptiveportalState struct {
	// Policy name.
	Name *string `pulumi:"name"`
	// Policy type. Valid values: `captive-portal`.
	PolicyType *string `pulumi:"policyType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Names of VLANs that use captive portal authentication.
	Vlan *string `pulumi:"vlan"`
}

type SwitchcontrollersecuritypolicyCaptiveportalState struct {
	// Policy name.
	Name pulumi.StringPtrInput
	// Policy type. Valid values: `captive-portal`.
	PolicyType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Names of VLANs that use captive portal authentication.
	Vlan pulumi.StringPtrInput
}

func (SwitchcontrollersecuritypolicyCaptiveportalState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollersecuritypolicyCaptiveportalState)(nil)).Elem()
}

type switchcontrollersecuritypolicyCaptiveportalArgs struct {
	// Policy name.
	Name *string `pulumi:"name"`
	// Policy type. Valid values: `captive-portal`.
	PolicyType *string `pulumi:"policyType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Names of VLANs that use captive portal authentication.
	Vlan *string `pulumi:"vlan"`
}

// The set of arguments for constructing a SwitchcontrollersecuritypolicyCaptiveportal resource.
type SwitchcontrollersecuritypolicyCaptiveportalArgs struct {
	// Policy name.
	Name pulumi.StringPtrInput
	// Policy type. Valid values: `captive-portal`.
	PolicyType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Names of VLANs that use captive portal authentication.
	Vlan pulumi.StringPtrInput
}

func (SwitchcontrollersecuritypolicyCaptiveportalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollersecuritypolicyCaptiveportalArgs)(nil)).Elem()
}

type SwitchcontrollersecuritypolicyCaptiveportalInput interface {
	pulumi.Input

	ToSwitchcontrollersecuritypolicyCaptiveportalOutput() SwitchcontrollersecuritypolicyCaptiveportalOutput
	ToSwitchcontrollersecuritypolicyCaptiveportalOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalOutput
}

func (*SwitchcontrollersecuritypolicyCaptiveportal) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollersecuritypolicyCaptiveportal)(nil)).Elem()
}

func (i *SwitchcontrollersecuritypolicyCaptiveportal) ToSwitchcontrollersecuritypolicyCaptiveportalOutput() SwitchcontrollersecuritypolicyCaptiveportalOutput {
	return i.ToSwitchcontrollersecuritypolicyCaptiveportalOutputWithContext(context.Background())
}

func (i *SwitchcontrollersecuritypolicyCaptiveportal) ToSwitchcontrollersecuritypolicyCaptiveportalOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollersecuritypolicyCaptiveportalOutput)
}

// SwitchcontrollersecuritypolicyCaptiveportalArrayInput is an input type that accepts SwitchcontrollersecuritypolicyCaptiveportalArray and SwitchcontrollersecuritypolicyCaptiveportalArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollersecuritypolicyCaptiveportalArrayInput` via:
//
//	SwitchcontrollersecuritypolicyCaptiveportalArray{ SwitchcontrollersecuritypolicyCaptiveportalArgs{...} }
type SwitchcontrollersecuritypolicyCaptiveportalArrayInput interface {
	pulumi.Input

	ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutput() SwitchcontrollersecuritypolicyCaptiveportalArrayOutput
	ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutputWithContext(context.Context) SwitchcontrollersecuritypolicyCaptiveportalArrayOutput
}

type SwitchcontrollersecuritypolicyCaptiveportalArray []SwitchcontrollersecuritypolicyCaptiveportalInput

func (SwitchcontrollersecuritypolicyCaptiveportalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollersecuritypolicyCaptiveportal)(nil)).Elem()
}

func (i SwitchcontrollersecuritypolicyCaptiveportalArray) ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutput() SwitchcontrollersecuritypolicyCaptiveportalArrayOutput {
	return i.ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollersecuritypolicyCaptiveportalArray) ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollersecuritypolicyCaptiveportalArrayOutput)
}

// SwitchcontrollersecuritypolicyCaptiveportalMapInput is an input type that accepts SwitchcontrollersecuritypolicyCaptiveportalMap and SwitchcontrollersecuritypolicyCaptiveportalMapOutput values.
// You can construct a concrete instance of `SwitchcontrollersecuritypolicyCaptiveportalMapInput` via:
//
//	SwitchcontrollersecuritypolicyCaptiveportalMap{ "key": SwitchcontrollersecuritypolicyCaptiveportalArgs{...} }
type SwitchcontrollersecuritypolicyCaptiveportalMapInput interface {
	pulumi.Input

	ToSwitchcontrollersecuritypolicyCaptiveportalMapOutput() SwitchcontrollersecuritypolicyCaptiveportalMapOutput
	ToSwitchcontrollersecuritypolicyCaptiveportalMapOutputWithContext(context.Context) SwitchcontrollersecuritypolicyCaptiveportalMapOutput
}

type SwitchcontrollersecuritypolicyCaptiveportalMap map[string]SwitchcontrollersecuritypolicyCaptiveportalInput

func (SwitchcontrollersecuritypolicyCaptiveportalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollersecuritypolicyCaptiveportal)(nil)).Elem()
}

func (i SwitchcontrollersecuritypolicyCaptiveportalMap) ToSwitchcontrollersecuritypolicyCaptiveportalMapOutput() SwitchcontrollersecuritypolicyCaptiveportalMapOutput {
	return i.ToSwitchcontrollersecuritypolicyCaptiveportalMapOutputWithContext(context.Background())
}

func (i SwitchcontrollersecuritypolicyCaptiveportalMap) ToSwitchcontrollersecuritypolicyCaptiveportalMapOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollersecuritypolicyCaptiveportalMapOutput)
}

type SwitchcontrollersecuritypolicyCaptiveportalOutput struct{ *pulumi.OutputState }

func (SwitchcontrollersecuritypolicyCaptiveportalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollersecuritypolicyCaptiveportal)(nil)).Elem()
}

func (o SwitchcontrollersecuritypolicyCaptiveportalOutput) ToSwitchcontrollersecuritypolicyCaptiveportalOutput() SwitchcontrollersecuritypolicyCaptiveportalOutput {
	return o
}

func (o SwitchcontrollersecuritypolicyCaptiveportalOutput) ToSwitchcontrollersecuritypolicyCaptiveportalOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalOutput {
	return o
}

// Policy name.
func (o SwitchcontrollersecuritypolicyCaptiveportalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollersecuritypolicyCaptiveportal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy type. Valid values: `captive-portal`.
func (o SwitchcontrollersecuritypolicyCaptiveportalOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollersecuritypolicyCaptiveportal) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollersecuritypolicyCaptiveportalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollersecuritypolicyCaptiveportal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Names of VLANs that use captive portal authentication.
func (o SwitchcontrollersecuritypolicyCaptiveportalOutput) Vlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollersecuritypolicyCaptiveportal) pulumi.StringOutput { return v.Vlan }).(pulumi.StringOutput)
}

type SwitchcontrollersecuritypolicyCaptiveportalArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollersecuritypolicyCaptiveportalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollersecuritypolicyCaptiveportal)(nil)).Elem()
}

func (o SwitchcontrollersecuritypolicyCaptiveportalArrayOutput) ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutput() SwitchcontrollersecuritypolicyCaptiveportalArrayOutput {
	return o
}

func (o SwitchcontrollersecuritypolicyCaptiveportalArrayOutput) ToSwitchcontrollersecuritypolicyCaptiveportalArrayOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalArrayOutput {
	return o
}

func (o SwitchcontrollersecuritypolicyCaptiveportalArrayOutput) Index(i pulumi.IntInput) SwitchcontrollersecuritypolicyCaptiveportalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollersecuritypolicyCaptiveportal {
		return vs[0].([]*SwitchcontrollersecuritypolicyCaptiveportal)[vs[1].(int)]
	}).(SwitchcontrollersecuritypolicyCaptiveportalOutput)
}

type SwitchcontrollersecuritypolicyCaptiveportalMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollersecuritypolicyCaptiveportalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollersecuritypolicyCaptiveportal)(nil)).Elem()
}

func (o SwitchcontrollersecuritypolicyCaptiveportalMapOutput) ToSwitchcontrollersecuritypolicyCaptiveportalMapOutput() SwitchcontrollersecuritypolicyCaptiveportalMapOutput {
	return o
}

func (o SwitchcontrollersecuritypolicyCaptiveportalMapOutput) ToSwitchcontrollersecuritypolicyCaptiveportalMapOutputWithContext(ctx context.Context) SwitchcontrollersecuritypolicyCaptiveportalMapOutput {
	return o
}

func (o SwitchcontrollersecuritypolicyCaptiveportalMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollersecuritypolicyCaptiveportalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollersecuritypolicyCaptiveportal {
		return vs[0].(map[string]*SwitchcontrollersecuritypolicyCaptiveportal)[vs[1].(string)]
	}).(SwitchcontrollersecuritypolicyCaptiveportalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollersecuritypolicyCaptiveportalInput)(nil)).Elem(), &SwitchcontrollersecuritypolicyCaptiveportal{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollersecuritypolicyCaptiveportalArrayInput)(nil)).Elem(), SwitchcontrollersecuritypolicyCaptiveportalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollersecuritypolicyCaptiveportalMapInput)(nil)).Elem(), SwitchcontrollersecuritypolicyCaptiveportalMap{})
	pulumi.RegisterOutputType(SwitchcontrollersecuritypolicyCaptiveportalOutput{})
	pulumi.RegisterOutputType(SwitchcontrollersecuritypolicyCaptiveportalArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollersecuritypolicyCaptiveportalMapOutput{})
}
