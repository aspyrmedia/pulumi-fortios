// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP blacklist reason. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Firewall InternetServiceIpblReason can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetserviceipblreason struct {
	pulumi.CustomResourceState

	// IP blacklist reason ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// IP blacklist reason name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetserviceipblreason registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetserviceipblreason(ctx *pulumi.Context,
	name string, args *FirewallInternetserviceipblreasonArgs, opts ...pulumi.ResourceOption) (*FirewallInternetserviceipblreason, error) {
	if args == nil {
		args = &FirewallInternetserviceipblreasonArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetserviceipblreason
	err := ctx.RegisterResource("fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetserviceipblreason gets an existing FirewallInternetserviceipblreason resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetserviceipblreason(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetserviceipblreasonState, opts ...pulumi.ResourceOption) (*FirewallInternetserviceipblreason, error) {
	var resource FirewallInternetserviceipblreason
	err := ctx.ReadResource("fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetserviceipblreason resources.
type firewallInternetserviceipblreasonState struct {
	// IP blacklist reason ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist reason name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetserviceipblreasonState struct {
	// IP blacklist reason ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist reason name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetserviceipblreasonState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetserviceipblreasonState)(nil)).Elem()
}

type firewallInternetserviceipblreasonArgs struct {
	// IP blacklist reason ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist reason name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetserviceipblreason resource.
type FirewallInternetserviceipblreasonArgs struct {
	// IP blacklist reason ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist reason name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetserviceipblreasonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetserviceipblreasonArgs)(nil)).Elem()
}

type FirewallInternetserviceipblreasonInput interface {
	pulumi.Input

	ToFirewallInternetserviceipblreasonOutput() FirewallInternetserviceipblreasonOutput
	ToFirewallInternetserviceipblreasonOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonOutput
}

func (*FirewallInternetserviceipblreason) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetserviceipblreason)(nil)).Elem()
}

func (i *FirewallInternetserviceipblreason) ToFirewallInternetserviceipblreasonOutput() FirewallInternetserviceipblreasonOutput {
	return i.ToFirewallInternetserviceipblreasonOutputWithContext(context.Background())
}

func (i *FirewallInternetserviceipblreason) ToFirewallInternetserviceipblreasonOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetserviceipblreasonOutput)
}

// FirewallInternetserviceipblreasonArrayInput is an input type that accepts FirewallInternetserviceipblreasonArray and FirewallInternetserviceipblreasonArrayOutput values.
// You can construct a concrete instance of `FirewallInternetserviceipblreasonArrayInput` via:
//
//	FirewallInternetserviceipblreasonArray{ FirewallInternetserviceipblreasonArgs{...} }
type FirewallInternetserviceipblreasonArrayInput interface {
	pulumi.Input

	ToFirewallInternetserviceipblreasonArrayOutput() FirewallInternetserviceipblreasonArrayOutput
	ToFirewallInternetserviceipblreasonArrayOutputWithContext(context.Context) FirewallInternetserviceipblreasonArrayOutput
}

type FirewallInternetserviceipblreasonArray []FirewallInternetserviceipblreasonInput

func (FirewallInternetserviceipblreasonArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetserviceipblreason)(nil)).Elem()
}

func (i FirewallInternetserviceipblreasonArray) ToFirewallInternetserviceipblreasonArrayOutput() FirewallInternetserviceipblreasonArrayOutput {
	return i.ToFirewallInternetserviceipblreasonArrayOutputWithContext(context.Background())
}

func (i FirewallInternetserviceipblreasonArray) ToFirewallInternetserviceipblreasonArrayOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetserviceipblreasonArrayOutput)
}

// FirewallInternetserviceipblreasonMapInput is an input type that accepts FirewallInternetserviceipblreasonMap and FirewallInternetserviceipblreasonMapOutput values.
// You can construct a concrete instance of `FirewallInternetserviceipblreasonMapInput` via:
//
//	FirewallInternetserviceipblreasonMap{ "key": FirewallInternetserviceipblreasonArgs{...} }
type FirewallInternetserviceipblreasonMapInput interface {
	pulumi.Input

	ToFirewallInternetserviceipblreasonMapOutput() FirewallInternetserviceipblreasonMapOutput
	ToFirewallInternetserviceipblreasonMapOutputWithContext(context.Context) FirewallInternetserviceipblreasonMapOutput
}

type FirewallInternetserviceipblreasonMap map[string]FirewallInternetserviceipblreasonInput

func (FirewallInternetserviceipblreasonMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetserviceipblreason)(nil)).Elem()
}

func (i FirewallInternetserviceipblreasonMap) ToFirewallInternetserviceipblreasonMapOutput() FirewallInternetserviceipblreasonMapOutput {
	return i.ToFirewallInternetserviceipblreasonMapOutputWithContext(context.Background())
}

func (i FirewallInternetserviceipblreasonMap) ToFirewallInternetserviceipblreasonMapOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetserviceipblreasonMapOutput)
}

type FirewallInternetserviceipblreasonOutput struct{ *pulumi.OutputState }

func (FirewallInternetserviceipblreasonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetserviceipblreason)(nil)).Elem()
}

func (o FirewallInternetserviceipblreasonOutput) ToFirewallInternetserviceipblreasonOutput() FirewallInternetserviceipblreasonOutput {
	return o
}

func (o FirewallInternetserviceipblreasonOutput) ToFirewallInternetserviceipblreasonOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonOutput {
	return o
}

// IP blacklist reason ID.
func (o FirewallInternetserviceipblreasonOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallInternetserviceipblreason) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// IP blacklist reason name.
func (o FirewallInternetserviceipblreasonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetserviceipblreason) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallInternetserviceipblreasonOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetserviceipblreason) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetserviceipblreasonArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetserviceipblreasonArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetserviceipblreason)(nil)).Elem()
}

func (o FirewallInternetserviceipblreasonArrayOutput) ToFirewallInternetserviceipblreasonArrayOutput() FirewallInternetserviceipblreasonArrayOutput {
	return o
}

func (o FirewallInternetserviceipblreasonArrayOutput) ToFirewallInternetserviceipblreasonArrayOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonArrayOutput {
	return o
}

func (o FirewallInternetserviceipblreasonArrayOutput) Index(i pulumi.IntInput) FirewallInternetserviceipblreasonOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetserviceipblreason {
		return vs[0].([]*FirewallInternetserviceipblreason)[vs[1].(int)]
	}).(FirewallInternetserviceipblreasonOutput)
}

type FirewallInternetserviceipblreasonMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetserviceipblreasonMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetserviceipblreason)(nil)).Elem()
}

func (o FirewallInternetserviceipblreasonMapOutput) ToFirewallInternetserviceipblreasonMapOutput() FirewallInternetserviceipblreasonMapOutput {
	return o
}

func (o FirewallInternetserviceipblreasonMapOutput) ToFirewallInternetserviceipblreasonMapOutputWithContext(ctx context.Context) FirewallInternetserviceipblreasonMapOutput {
	return o
}

func (o FirewallInternetserviceipblreasonMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetserviceipblreasonOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetserviceipblreason {
		return vs[0].(map[string]*FirewallInternetserviceipblreason)[vs[1].(string)]
	}).(FirewallInternetserviceipblreasonOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetserviceipblreasonInput)(nil)).Elem(), &FirewallInternetserviceipblreason{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetserviceipblreasonArrayInput)(nil)).Elem(), FirewallInternetserviceipblreasonArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetserviceipblreasonMapInput)(nil)).Elem(), FirewallInternetserviceipblreasonMap{})
	pulumi.RegisterOutputType(FirewallInternetserviceipblreasonOutput{})
	pulumi.RegisterOutputType(FirewallInternetserviceipblreasonArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetserviceipblreasonMapOutput{})
}
