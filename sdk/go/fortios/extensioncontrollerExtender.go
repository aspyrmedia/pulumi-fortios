// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Extender controller configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// # ExtensionController Extender can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type ExtensioncontrollerExtender struct {
	pulumi.CustomResourceState

	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess pulumi.StringOutput `pulumi:"allowaccess"`
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringOutput `pulumi:"authorized"`
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit pulumi.IntOutput `pulumi:"bandwidthLimit"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Device ID.
	DeviceId pulumi.IntOutput `pulumi:"deviceId"`
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth pulumi.StringOutput `pulumi:"enforceBandwidth"`
	// FortiExtender name.
	ExtName pulumi.StringOutput `pulumi:"extName"`
	// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
	ExtensionType pulumi.StringOutput `pulumi:"extensionType"`
	// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
	FirmwareProvisionLatest pulumi.StringOutput `pulumi:"firmwareProvisionLatest"`
	// FortiExtender serial number.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Set the managed extender's administrator password.
	LoginPassword pulumi.StringPtrOutput `pulumi:"loginPassword"`
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange pulumi.StringOutput `pulumi:"loginPasswordChange"`
	// FortiExtender entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
	OverrideAllowaccess pulumi.StringOutput `pulumi:"overrideAllowaccess"`
	// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
	OverrideEnforceBandwidth pulumi.StringOutput `pulumi:"overrideEnforceBandwidth"`
	// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
	OverrideLoginPasswordChange pulumi.StringOutput `pulumi:"overrideLoginPasswordChange"`
	// FortiExtender profile configuration.
	Profile pulumi.StringOutput `pulumi:"profile"`
	// VDOM.
	Vdom pulumi.IntOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
	WanExtension ExtensioncontrollerExtenderWanExtensionOutput `pulumi:"wanExtension"`
}

// NewExtensioncontrollerExtender registers a new resource with the given unique name, arguments, and options.
func NewExtensioncontrollerExtender(ctx *pulumi.Context,
	name string, args *ExtensioncontrollerExtenderArgs, opts ...pulumi.ResourceOption) (*ExtensioncontrollerExtender, error) {
	if args == nil {
		args = &ExtensioncontrollerExtenderArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ExtensioncontrollerExtender
	err := ctx.RegisterResource("fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtensioncontrollerExtender gets an existing ExtensioncontrollerExtender resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtensioncontrollerExtender(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensioncontrollerExtenderState, opts ...pulumi.ResourceOption) (*ExtensioncontrollerExtender, error) {
	var resource ExtensioncontrollerExtender
	err := ctx.ReadResource("fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtensioncontrollerExtender resources.
type extensioncontrollerExtenderState struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess *string `pulumi:"allowaccess"`
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit *int `pulumi:"bandwidthLimit"`
	// Description.
	Description *string `pulumi:"description"`
	// Device ID.
	DeviceId *int `pulumi:"deviceId"`
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth *string `pulumi:"enforceBandwidth"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
	ExtensionType *string `pulumi:"extensionType"`
	// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
	FirmwareProvisionLatest *string `pulumi:"firmwareProvisionLatest"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// Set the managed extender's administrator password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange *string `pulumi:"loginPasswordChange"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
	OverrideAllowaccess *string `pulumi:"overrideAllowaccess"`
	// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
	OverrideEnforceBandwidth *string `pulumi:"overrideEnforceBandwidth"`
	// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
	OverrideLoginPasswordChange *string `pulumi:"overrideLoginPasswordChange"`
	// FortiExtender profile configuration.
	Profile *string `pulumi:"profile"`
	// VDOM.
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
	WanExtension *ExtensioncontrollerExtenderWanExtension `pulumi:"wanExtension"`
}

type ExtensioncontrollerExtenderState struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess pulumi.StringPtrInput
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit pulumi.IntPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Device ID.
	DeviceId pulumi.IntPtrInput
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
	ExtensionType pulumi.StringPtrInput
	// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
	FirmwareProvisionLatest pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// Set the managed extender's administrator password.
	LoginPassword pulumi.StringPtrInput
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange pulumi.StringPtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
	OverrideAllowaccess pulumi.StringPtrInput
	// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
	OverrideEnforceBandwidth pulumi.StringPtrInput
	// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
	OverrideLoginPasswordChange pulumi.StringPtrInput
	// FortiExtender profile configuration.
	Profile pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
	WanExtension ExtensioncontrollerExtenderWanExtensionPtrInput
}

func (ExtensioncontrollerExtenderState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensioncontrollerExtenderState)(nil)).Elem()
}

type extensioncontrollerExtenderArgs struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess *string `pulumi:"allowaccess"`
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit *int `pulumi:"bandwidthLimit"`
	// Description.
	Description *string `pulumi:"description"`
	// Device ID.
	DeviceId *int `pulumi:"deviceId"`
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth *string `pulumi:"enforceBandwidth"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
	ExtensionType *string `pulumi:"extensionType"`
	// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
	FirmwareProvisionLatest *string `pulumi:"firmwareProvisionLatest"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// Set the managed extender's administrator password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange *string `pulumi:"loginPasswordChange"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
	OverrideAllowaccess *string `pulumi:"overrideAllowaccess"`
	// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
	OverrideEnforceBandwidth *string `pulumi:"overrideEnforceBandwidth"`
	// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
	OverrideLoginPasswordChange *string `pulumi:"overrideLoginPasswordChange"`
	// FortiExtender profile configuration.
	Profile *string `pulumi:"profile"`
	// VDOM.
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
	WanExtension *ExtensioncontrollerExtenderWanExtension `pulumi:"wanExtension"`
}

// The set of arguments for constructing a ExtensioncontrollerExtender resource.
type ExtensioncontrollerExtenderArgs struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess pulumi.StringPtrInput
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit pulumi.IntPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Device ID.
	DeviceId pulumi.IntPtrInput
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
	ExtensionType pulumi.StringPtrInput
	// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
	FirmwareProvisionLatest pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// Set the managed extender's administrator password.
	LoginPassword pulumi.StringPtrInput
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange pulumi.StringPtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
	OverrideAllowaccess pulumi.StringPtrInput
	// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
	OverrideEnforceBandwidth pulumi.StringPtrInput
	// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
	OverrideLoginPasswordChange pulumi.StringPtrInput
	// FortiExtender profile configuration.
	Profile pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
	WanExtension ExtensioncontrollerExtenderWanExtensionPtrInput
}

func (ExtensioncontrollerExtenderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensioncontrollerExtenderArgs)(nil)).Elem()
}

type ExtensioncontrollerExtenderInput interface {
	pulumi.Input

	ToExtensioncontrollerExtenderOutput() ExtensioncontrollerExtenderOutput
	ToExtensioncontrollerExtenderOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderOutput
}

func (*ExtensioncontrollerExtender) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensioncontrollerExtender)(nil)).Elem()
}

func (i *ExtensioncontrollerExtender) ToExtensioncontrollerExtenderOutput() ExtensioncontrollerExtenderOutput {
	return i.ToExtensioncontrollerExtenderOutputWithContext(context.Background())
}

func (i *ExtensioncontrollerExtender) ToExtensioncontrollerExtenderOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerExtenderOutput)
}

// ExtensioncontrollerExtenderArrayInput is an input type that accepts ExtensioncontrollerExtenderArray and ExtensioncontrollerExtenderArrayOutput values.
// You can construct a concrete instance of `ExtensioncontrollerExtenderArrayInput` via:
//
//	ExtensioncontrollerExtenderArray{ ExtensioncontrollerExtenderArgs{...} }
type ExtensioncontrollerExtenderArrayInput interface {
	pulumi.Input

	ToExtensioncontrollerExtenderArrayOutput() ExtensioncontrollerExtenderArrayOutput
	ToExtensioncontrollerExtenderArrayOutputWithContext(context.Context) ExtensioncontrollerExtenderArrayOutput
}

type ExtensioncontrollerExtenderArray []ExtensioncontrollerExtenderInput

func (ExtensioncontrollerExtenderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensioncontrollerExtender)(nil)).Elem()
}

func (i ExtensioncontrollerExtenderArray) ToExtensioncontrollerExtenderArrayOutput() ExtensioncontrollerExtenderArrayOutput {
	return i.ToExtensioncontrollerExtenderArrayOutputWithContext(context.Background())
}

func (i ExtensioncontrollerExtenderArray) ToExtensioncontrollerExtenderArrayOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerExtenderArrayOutput)
}

// ExtensioncontrollerExtenderMapInput is an input type that accepts ExtensioncontrollerExtenderMap and ExtensioncontrollerExtenderMapOutput values.
// You can construct a concrete instance of `ExtensioncontrollerExtenderMapInput` via:
//
//	ExtensioncontrollerExtenderMap{ "key": ExtensioncontrollerExtenderArgs{...} }
type ExtensioncontrollerExtenderMapInput interface {
	pulumi.Input

	ToExtensioncontrollerExtenderMapOutput() ExtensioncontrollerExtenderMapOutput
	ToExtensioncontrollerExtenderMapOutputWithContext(context.Context) ExtensioncontrollerExtenderMapOutput
}

type ExtensioncontrollerExtenderMap map[string]ExtensioncontrollerExtenderInput

func (ExtensioncontrollerExtenderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensioncontrollerExtender)(nil)).Elem()
}

func (i ExtensioncontrollerExtenderMap) ToExtensioncontrollerExtenderMapOutput() ExtensioncontrollerExtenderMapOutput {
	return i.ToExtensioncontrollerExtenderMapOutputWithContext(context.Background())
}

func (i ExtensioncontrollerExtenderMap) ToExtensioncontrollerExtenderMapOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerExtenderMapOutput)
}

type ExtensioncontrollerExtenderOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerExtenderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensioncontrollerExtender)(nil)).Elem()
}

func (o ExtensioncontrollerExtenderOutput) ToExtensioncontrollerExtenderOutput() ExtensioncontrollerExtenderOutput {
	return o
}

func (o ExtensioncontrollerExtenderOutput) ToExtensioncontrollerExtenderOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderOutput {
	return o
}

// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
func (o ExtensioncontrollerExtenderOutput) Allowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.Allowaccess }).(pulumi.StringOutput)
}

// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
func (o ExtensioncontrollerExtenderOutput) Authorized() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.Authorized }).(pulumi.StringOutput)
}

// FortiExtender LAN extension bandwidth limit (Mbps).
func (o ExtensioncontrollerExtenderOutput) BandwidthLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.IntOutput { return v.BandwidthLimit }).(pulumi.IntOutput)
}

// Description.
func (o ExtensioncontrollerExtenderOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Device ID.
func (o ExtensioncontrollerExtenderOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
func (o ExtensioncontrollerExtenderOutput) EnforceBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.EnforceBandwidth }).(pulumi.StringOutput)
}

// FortiExtender name.
func (o ExtensioncontrollerExtenderOutput) ExtName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.ExtName }).(pulumi.StringOutput)
}

// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
func (o ExtensioncontrollerExtenderOutput) ExtensionType() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.ExtensionType }).(pulumi.StringOutput)
}

// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
func (o ExtensioncontrollerExtenderOutput) FirmwareProvisionLatest() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.FirmwareProvisionLatest }).(pulumi.StringOutput)
}

// FortiExtender serial number.
func (o ExtensioncontrollerExtenderOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// Set the managed extender's administrator password.
func (o ExtensioncontrollerExtenderOutput) LoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringPtrOutput { return v.LoginPassword }).(pulumi.StringPtrOutput)
}

// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
func (o ExtensioncontrollerExtenderOutput) LoginPasswordChange() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.LoginPasswordChange }).(pulumi.StringOutput)
}

// FortiExtender entry name.
func (o ExtensioncontrollerExtenderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
func (o ExtensioncontrollerExtenderOutput) OverrideAllowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.OverrideAllowaccess }).(pulumi.StringOutput)
}

// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
func (o ExtensioncontrollerExtenderOutput) OverrideEnforceBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.OverrideEnforceBandwidth }).(pulumi.StringOutput)
}

// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
func (o ExtensioncontrollerExtenderOutput) OverrideLoginPasswordChange() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.OverrideLoginPasswordChange }).(pulumi.StringOutput)
}

// FortiExtender profile configuration.
func (o ExtensioncontrollerExtenderOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringOutput { return v.Profile }).(pulumi.StringOutput)
}

// VDOM.
func (o ExtensioncontrollerExtenderOutput) Vdom() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.IntOutput { return v.Vdom }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExtensioncontrollerExtenderOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
func (o ExtensioncontrollerExtenderOutput) WanExtension() ExtensioncontrollerExtenderWanExtensionOutput {
	return o.ApplyT(func(v *ExtensioncontrollerExtender) ExtensioncontrollerExtenderWanExtensionOutput {
		return v.WanExtension
	}).(ExtensioncontrollerExtenderWanExtensionOutput)
}

type ExtensioncontrollerExtenderArrayOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerExtenderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensioncontrollerExtender)(nil)).Elem()
}

func (o ExtensioncontrollerExtenderArrayOutput) ToExtensioncontrollerExtenderArrayOutput() ExtensioncontrollerExtenderArrayOutput {
	return o
}

func (o ExtensioncontrollerExtenderArrayOutput) ToExtensioncontrollerExtenderArrayOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderArrayOutput {
	return o
}

func (o ExtensioncontrollerExtenderArrayOutput) Index(i pulumi.IntInput) ExtensioncontrollerExtenderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtensioncontrollerExtender {
		return vs[0].([]*ExtensioncontrollerExtender)[vs[1].(int)]
	}).(ExtensioncontrollerExtenderOutput)
}

type ExtensioncontrollerExtenderMapOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerExtenderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensioncontrollerExtender)(nil)).Elem()
}

func (o ExtensioncontrollerExtenderMapOutput) ToExtensioncontrollerExtenderMapOutput() ExtensioncontrollerExtenderMapOutput {
	return o
}

func (o ExtensioncontrollerExtenderMapOutput) ToExtensioncontrollerExtenderMapOutputWithContext(ctx context.Context) ExtensioncontrollerExtenderMapOutput {
	return o
}

func (o ExtensioncontrollerExtenderMapOutput) MapIndex(k pulumi.StringInput) ExtensioncontrollerExtenderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtensioncontrollerExtender {
		return vs[0].(map[string]*ExtensioncontrollerExtender)[vs[1].(string)]
	}).(ExtensioncontrollerExtenderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerExtenderInput)(nil)).Elem(), &ExtensioncontrollerExtender{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerExtenderArrayInput)(nil)).Elem(), ExtensioncontrollerExtenderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerExtenderMapInput)(nil)).Elem(), ExtensioncontrollerExtenderMap{})
	pulumi.RegisterOutputType(ExtensioncontrollerExtenderOutput{})
	pulumi.RegisterOutputType(ExtensioncontrollerExtenderArrayOutput{})
	pulumi.RegisterOutputType(ExtensioncontrollerExtenderMapOutput{})
}