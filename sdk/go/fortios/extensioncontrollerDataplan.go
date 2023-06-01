// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FortiExtender dataplan configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// # ExtensionController Dataplan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/extensioncontrollerDataplan:ExtensioncontrollerDataplan labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/extensioncontrollerDataplan:ExtensioncontrollerDataplan labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type ExtensioncontrollerDataplan struct {
	pulumi.CustomResourceState

	// APN configuration.
	Apn pulumi.StringOutput `pulumi:"apn"`
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// Billing day of the month (1 - 31).
	BillingDate pulumi.IntOutput `pulumi:"billingDate"`
	// Capacity in MB (0 - 102400000).
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// Carrier configuration.
	Carrier pulumi.StringOutput `pulumi:"carrier"`
	// ICCID configuration.
	Iccid pulumi.StringOutput `pulumi:"iccid"`
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId pulumi.StringOutput `pulumi:"modemId"`
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee pulumi.IntOutput `pulumi:"monthlyFee"`
	// FortiExtender data plan name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage pulumi.StringOutput `pulumi:"overage"`
	// Password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn pulumi.StringOutput `pulumi:"pdn"`
	// Preferred subnet mask (0 - 32).
	PreferredSubnet pulumi.IntOutput `pulumi:"preferredSubnet"`
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork pulumi.StringOutput `pulumi:"privateNetwork"`
	// Signal period (600 to 18000 seconds).
	SignalPeriod pulumi.IntOutput `pulumi:"signalPeriod"`
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold pulumi.IntOutput `pulumi:"signalThreshold"`
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot pulumi.StringOutput `pulumi:"slot"`
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Username.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtensioncontrollerDataplan registers a new resource with the given unique name, arguments, and options.
func NewExtensioncontrollerDataplan(ctx *pulumi.Context,
	name string, args *ExtensioncontrollerDataplanArgs, opts ...pulumi.ResourceOption) (*ExtensioncontrollerDataplan, error) {
	if args == nil {
		args = &ExtensioncontrollerDataplanArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ExtensioncontrollerDataplan
	err := ctx.RegisterResource("fortios:index/extensioncontrollerDataplan:ExtensioncontrollerDataplan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtensioncontrollerDataplan gets an existing ExtensioncontrollerDataplan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtensioncontrollerDataplan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensioncontrollerDataplanState, opts ...pulumi.ResourceOption) (*ExtensioncontrollerDataplan, error) {
	var resource ExtensioncontrollerDataplan
	err := ctx.ReadResource("fortios:index/extensioncontrollerDataplan:ExtensioncontrollerDataplan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtensioncontrollerDataplan resources.
type extensioncontrollerDataplanState struct {
	// APN configuration.
	Apn *string `pulumi:"apn"`
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType *string `pulumi:"authType"`
	// Billing day of the month (1 - 31).
	BillingDate *int `pulumi:"billingDate"`
	// Capacity in MB (0 - 102400000).
	Capacity *int `pulumi:"capacity"`
	// Carrier configuration.
	Carrier *string `pulumi:"carrier"`
	// ICCID configuration.
	Iccid *string `pulumi:"iccid"`
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId *string `pulumi:"modemId"`
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee *int `pulumi:"monthlyFee"`
	// FortiExtender data plan name.
	Name *string `pulumi:"name"`
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage *string `pulumi:"overage"`
	// Password.
	Password *string `pulumi:"password"`
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn *string `pulumi:"pdn"`
	// Preferred subnet mask (0 - 32).
	PreferredSubnet *int `pulumi:"preferredSubnet"`
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork *string `pulumi:"privateNetwork"`
	// Signal period (600 to 18000 seconds).
	SignalPeriod *int `pulumi:"signalPeriod"`
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold *int `pulumi:"signalThreshold"`
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot *string `pulumi:"slot"`
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type *string `pulumi:"type"`
	// Username.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExtensioncontrollerDataplanState struct {
	// APN configuration.
	Apn pulumi.StringPtrInput
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType pulumi.StringPtrInput
	// Billing day of the month (1 - 31).
	BillingDate pulumi.IntPtrInput
	// Capacity in MB (0 - 102400000).
	Capacity pulumi.IntPtrInput
	// Carrier configuration.
	Carrier pulumi.StringPtrInput
	// ICCID configuration.
	Iccid pulumi.StringPtrInput
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId pulumi.StringPtrInput
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee pulumi.IntPtrInput
	// FortiExtender data plan name.
	Name pulumi.StringPtrInput
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage pulumi.StringPtrInput
	// Password.
	Password pulumi.StringPtrInput
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn pulumi.StringPtrInput
	// Preferred subnet mask (0 - 32).
	PreferredSubnet pulumi.IntPtrInput
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork pulumi.StringPtrInput
	// Signal period (600 to 18000 seconds).
	SignalPeriod pulumi.IntPtrInput
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold pulumi.IntPtrInput
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot pulumi.StringPtrInput
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type pulumi.StringPtrInput
	// Username.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtensioncontrollerDataplanState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensioncontrollerDataplanState)(nil)).Elem()
}

type extensioncontrollerDataplanArgs struct {
	// APN configuration.
	Apn *string `pulumi:"apn"`
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType *string `pulumi:"authType"`
	// Billing day of the month (1 - 31).
	BillingDate *int `pulumi:"billingDate"`
	// Capacity in MB (0 - 102400000).
	Capacity *int `pulumi:"capacity"`
	// Carrier configuration.
	Carrier *string `pulumi:"carrier"`
	// ICCID configuration.
	Iccid *string `pulumi:"iccid"`
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId *string `pulumi:"modemId"`
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee *int `pulumi:"monthlyFee"`
	// FortiExtender data plan name.
	Name *string `pulumi:"name"`
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage *string `pulumi:"overage"`
	// Password.
	Password *string `pulumi:"password"`
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn *string `pulumi:"pdn"`
	// Preferred subnet mask (0 - 32).
	PreferredSubnet *int `pulumi:"preferredSubnet"`
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork *string `pulumi:"privateNetwork"`
	// Signal period (600 to 18000 seconds).
	SignalPeriod *int `pulumi:"signalPeriod"`
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold *int `pulumi:"signalThreshold"`
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot *string `pulumi:"slot"`
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type *string `pulumi:"type"`
	// Username.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ExtensioncontrollerDataplan resource.
type ExtensioncontrollerDataplanArgs struct {
	// APN configuration.
	Apn pulumi.StringPtrInput
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType pulumi.StringPtrInput
	// Billing day of the month (1 - 31).
	BillingDate pulumi.IntPtrInput
	// Capacity in MB (0 - 102400000).
	Capacity pulumi.IntPtrInput
	// Carrier configuration.
	Carrier pulumi.StringPtrInput
	// ICCID configuration.
	Iccid pulumi.StringPtrInput
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId pulumi.StringPtrInput
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee pulumi.IntPtrInput
	// FortiExtender data plan name.
	Name pulumi.StringPtrInput
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage pulumi.StringPtrInput
	// Password.
	Password pulumi.StringPtrInput
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn pulumi.StringPtrInput
	// Preferred subnet mask (0 - 32).
	PreferredSubnet pulumi.IntPtrInput
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork pulumi.StringPtrInput
	// Signal period (600 to 18000 seconds).
	SignalPeriod pulumi.IntPtrInput
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold pulumi.IntPtrInput
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot pulumi.StringPtrInput
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type pulumi.StringPtrInput
	// Username.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtensioncontrollerDataplanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensioncontrollerDataplanArgs)(nil)).Elem()
}

type ExtensioncontrollerDataplanInput interface {
	pulumi.Input

	ToExtensioncontrollerDataplanOutput() ExtensioncontrollerDataplanOutput
	ToExtensioncontrollerDataplanOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanOutput
}

func (*ExtensioncontrollerDataplan) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensioncontrollerDataplan)(nil)).Elem()
}

func (i *ExtensioncontrollerDataplan) ToExtensioncontrollerDataplanOutput() ExtensioncontrollerDataplanOutput {
	return i.ToExtensioncontrollerDataplanOutputWithContext(context.Background())
}

func (i *ExtensioncontrollerDataplan) ToExtensioncontrollerDataplanOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerDataplanOutput)
}

// ExtensioncontrollerDataplanArrayInput is an input type that accepts ExtensioncontrollerDataplanArray and ExtensioncontrollerDataplanArrayOutput values.
// You can construct a concrete instance of `ExtensioncontrollerDataplanArrayInput` via:
//
//	ExtensioncontrollerDataplanArray{ ExtensioncontrollerDataplanArgs{...} }
type ExtensioncontrollerDataplanArrayInput interface {
	pulumi.Input

	ToExtensioncontrollerDataplanArrayOutput() ExtensioncontrollerDataplanArrayOutput
	ToExtensioncontrollerDataplanArrayOutputWithContext(context.Context) ExtensioncontrollerDataplanArrayOutput
}

type ExtensioncontrollerDataplanArray []ExtensioncontrollerDataplanInput

func (ExtensioncontrollerDataplanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensioncontrollerDataplan)(nil)).Elem()
}

func (i ExtensioncontrollerDataplanArray) ToExtensioncontrollerDataplanArrayOutput() ExtensioncontrollerDataplanArrayOutput {
	return i.ToExtensioncontrollerDataplanArrayOutputWithContext(context.Background())
}

func (i ExtensioncontrollerDataplanArray) ToExtensioncontrollerDataplanArrayOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerDataplanArrayOutput)
}

// ExtensioncontrollerDataplanMapInput is an input type that accepts ExtensioncontrollerDataplanMap and ExtensioncontrollerDataplanMapOutput values.
// You can construct a concrete instance of `ExtensioncontrollerDataplanMapInput` via:
//
//	ExtensioncontrollerDataplanMap{ "key": ExtensioncontrollerDataplanArgs{...} }
type ExtensioncontrollerDataplanMapInput interface {
	pulumi.Input

	ToExtensioncontrollerDataplanMapOutput() ExtensioncontrollerDataplanMapOutput
	ToExtensioncontrollerDataplanMapOutputWithContext(context.Context) ExtensioncontrollerDataplanMapOutput
}

type ExtensioncontrollerDataplanMap map[string]ExtensioncontrollerDataplanInput

func (ExtensioncontrollerDataplanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensioncontrollerDataplan)(nil)).Elem()
}

func (i ExtensioncontrollerDataplanMap) ToExtensioncontrollerDataplanMapOutput() ExtensioncontrollerDataplanMapOutput {
	return i.ToExtensioncontrollerDataplanMapOutputWithContext(context.Background())
}

func (i ExtensioncontrollerDataplanMap) ToExtensioncontrollerDataplanMapOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerDataplanMapOutput)
}

type ExtensioncontrollerDataplanOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerDataplanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensioncontrollerDataplan)(nil)).Elem()
}

func (o ExtensioncontrollerDataplanOutput) ToExtensioncontrollerDataplanOutput() ExtensioncontrollerDataplanOutput {
	return o
}

func (o ExtensioncontrollerDataplanOutput) ToExtensioncontrollerDataplanOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanOutput {
	return o
}

// APN configuration.
func (o ExtensioncontrollerDataplanOutput) Apn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Apn }).(pulumi.StringOutput)
}

// Authentication type. Valid values: `none`, `pap`, `chap`.
func (o ExtensioncontrollerDataplanOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// Billing day of the month (1 - 31).
func (o ExtensioncontrollerDataplanOutput) BillingDate() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.IntOutput { return v.BillingDate }).(pulumi.IntOutput)
}

// Capacity in MB (0 - 102400000).
func (o ExtensioncontrollerDataplanOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.IntOutput { return v.Capacity }).(pulumi.IntOutput)
}

// Carrier configuration.
func (o ExtensioncontrollerDataplanOutput) Carrier() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Carrier }).(pulumi.StringOutput)
}

// ICCID configuration.
func (o ExtensioncontrollerDataplanOutput) Iccid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Iccid }).(pulumi.StringOutput)
}

// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
func (o ExtensioncontrollerDataplanOutput) ModemId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.ModemId }).(pulumi.StringOutput)
}

// Monthly fee of dataplan (0 - 100000, in local currency).
func (o ExtensioncontrollerDataplanOutput) MonthlyFee() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.IntOutput { return v.MonthlyFee }).(pulumi.IntOutput)
}

// FortiExtender data plan name.
func (o ExtensioncontrollerDataplanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
func (o ExtensioncontrollerDataplanOutput) Overage() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Overage }).(pulumi.StringOutput)
}

// Password.
func (o ExtensioncontrollerDataplanOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
func (o ExtensioncontrollerDataplanOutput) Pdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Pdn }).(pulumi.StringOutput)
}

// Preferred subnet mask (0 - 32).
func (o ExtensioncontrollerDataplanOutput) PreferredSubnet() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.IntOutput { return v.PreferredSubnet }).(pulumi.IntOutput)
}

// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
func (o ExtensioncontrollerDataplanOutput) PrivateNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.PrivateNetwork }).(pulumi.StringOutput)
}

// Signal period (600 to 18000 seconds).
func (o ExtensioncontrollerDataplanOutput) SignalPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.IntOutput { return v.SignalPeriod }).(pulumi.IntOutput)
}

// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
func (o ExtensioncontrollerDataplanOutput) SignalThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.IntOutput { return v.SignalThreshold }).(pulumi.IntOutput)
}

// SIM slot configuration. Valid values: `sim1`, `sim2`.
func (o ExtensioncontrollerDataplanOutput) Slot() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Slot }).(pulumi.StringOutput)
}

// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
func (o ExtensioncontrollerDataplanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Username.
func (o ExtensioncontrollerDataplanOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExtensioncontrollerDataplanOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensioncontrollerDataplan) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExtensioncontrollerDataplanArrayOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerDataplanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensioncontrollerDataplan)(nil)).Elem()
}

func (o ExtensioncontrollerDataplanArrayOutput) ToExtensioncontrollerDataplanArrayOutput() ExtensioncontrollerDataplanArrayOutput {
	return o
}

func (o ExtensioncontrollerDataplanArrayOutput) ToExtensioncontrollerDataplanArrayOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanArrayOutput {
	return o
}

func (o ExtensioncontrollerDataplanArrayOutput) Index(i pulumi.IntInput) ExtensioncontrollerDataplanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtensioncontrollerDataplan {
		return vs[0].([]*ExtensioncontrollerDataplan)[vs[1].(int)]
	}).(ExtensioncontrollerDataplanOutput)
}

type ExtensioncontrollerDataplanMapOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerDataplanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensioncontrollerDataplan)(nil)).Elem()
}

func (o ExtensioncontrollerDataplanMapOutput) ToExtensioncontrollerDataplanMapOutput() ExtensioncontrollerDataplanMapOutput {
	return o
}

func (o ExtensioncontrollerDataplanMapOutput) ToExtensioncontrollerDataplanMapOutputWithContext(ctx context.Context) ExtensioncontrollerDataplanMapOutput {
	return o
}

func (o ExtensioncontrollerDataplanMapOutput) MapIndex(k pulumi.StringInput) ExtensioncontrollerDataplanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtensioncontrollerDataplan {
		return vs[0].(map[string]*ExtensioncontrollerDataplan)[vs[1].(string)]
	}).(ExtensioncontrollerDataplanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerDataplanInput)(nil)).Elem(), &ExtensioncontrollerDataplan{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerDataplanArrayInput)(nil)).Elem(), ExtensioncontrollerDataplanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerDataplanMapInput)(nil)).Elem(), ExtensioncontrollerDataplanMap{})
	pulumi.RegisterOutputType(ExtensioncontrollerDataplanOutput{})
	pulumi.RegisterOutputType(ExtensioncontrollerDataplanArrayOutput{})
	pulumi.RegisterOutputType(ExtensioncontrollerDataplanMapOutput{})
}