// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Show Internet Service botnet. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// # Firewall InternetServiceBotnet can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetservicebotnet:FirewallInternetservicebotnet labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetservicebotnet:FirewallInternetservicebotnet labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetservicebotnet struct {
	pulumi.CustomResourceState

	// Internet Service Botnet ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Internet Service Botnet name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetservicebotnet registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetservicebotnet(ctx *pulumi.Context,
	name string, args *FirewallInternetservicebotnetArgs, opts ...pulumi.ResourceOption) (*FirewallInternetservicebotnet, error) {
	if args == nil {
		args = &FirewallInternetservicebotnetArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetservicebotnet
	err := ctx.RegisterResource("fortios:index/firewallInternetservicebotnet:FirewallInternetservicebotnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetservicebotnet gets an existing FirewallInternetservicebotnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetservicebotnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetservicebotnetState, opts ...pulumi.ResourceOption) (*FirewallInternetservicebotnet, error) {
	var resource FirewallInternetservicebotnet
	err := ctx.ReadResource("fortios:index/firewallInternetservicebotnet:FirewallInternetservicebotnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetservicebotnet resources.
type firewallInternetservicebotnetState struct {
	// Internet Service Botnet ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service Botnet name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetservicebotnetState struct {
	// Internet Service Botnet ID.
	Fosid pulumi.IntPtrInput
	// Internet Service Botnet name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetservicebotnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetservicebotnetState)(nil)).Elem()
}

type firewallInternetservicebotnetArgs struct {
	// Internet Service Botnet ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service Botnet name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetservicebotnet resource.
type FirewallInternetservicebotnetArgs struct {
	// Internet Service Botnet ID.
	Fosid pulumi.IntPtrInput
	// Internet Service Botnet name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetservicebotnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetservicebotnetArgs)(nil)).Elem()
}

type FirewallInternetservicebotnetInput interface {
	pulumi.Input

	ToFirewallInternetservicebotnetOutput() FirewallInternetservicebotnetOutput
	ToFirewallInternetservicebotnetOutputWithContext(ctx context.Context) FirewallInternetservicebotnetOutput
}

func (*FirewallInternetservicebotnet) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetservicebotnet)(nil)).Elem()
}

func (i *FirewallInternetservicebotnet) ToFirewallInternetservicebotnetOutput() FirewallInternetservicebotnetOutput {
	return i.ToFirewallInternetservicebotnetOutputWithContext(context.Background())
}

func (i *FirewallInternetservicebotnet) ToFirewallInternetservicebotnetOutputWithContext(ctx context.Context) FirewallInternetservicebotnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetservicebotnetOutput)
}

// FirewallInternetservicebotnetArrayInput is an input type that accepts FirewallInternetservicebotnetArray and FirewallInternetservicebotnetArrayOutput values.
// You can construct a concrete instance of `FirewallInternetservicebotnetArrayInput` via:
//
//	FirewallInternetservicebotnetArray{ FirewallInternetservicebotnetArgs{...} }
type FirewallInternetservicebotnetArrayInput interface {
	pulumi.Input

	ToFirewallInternetservicebotnetArrayOutput() FirewallInternetservicebotnetArrayOutput
	ToFirewallInternetservicebotnetArrayOutputWithContext(context.Context) FirewallInternetservicebotnetArrayOutput
}

type FirewallInternetservicebotnetArray []FirewallInternetservicebotnetInput

func (FirewallInternetservicebotnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetservicebotnet)(nil)).Elem()
}

func (i FirewallInternetservicebotnetArray) ToFirewallInternetservicebotnetArrayOutput() FirewallInternetservicebotnetArrayOutput {
	return i.ToFirewallInternetservicebotnetArrayOutputWithContext(context.Background())
}

func (i FirewallInternetservicebotnetArray) ToFirewallInternetservicebotnetArrayOutputWithContext(ctx context.Context) FirewallInternetservicebotnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetservicebotnetArrayOutput)
}

// FirewallInternetservicebotnetMapInput is an input type that accepts FirewallInternetservicebotnetMap and FirewallInternetservicebotnetMapOutput values.
// You can construct a concrete instance of `FirewallInternetservicebotnetMapInput` via:
//
//	FirewallInternetservicebotnetMap{ "key": FirewallInternetservicebotnetArgs{...} }
type FirewallInternetservicebotnetMapInput interface {
	pulumi.Input

	ToFirewallInternetservicebotnetMapOutput() FirewallInternetservicebotnetMapOutput
	ToFirewallInternetservicebotnetMapOutputWithContext(context.Context) FirewallInternetservicebotnetMapOutput
}

type FirewallInternetservicebotnetMap map[string]FirewallInternetservicebotnetInput

func (FirewallInternetservicebotnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetservicebotnet)(nil)).Elem()
}

func (i FirewallInternetservicebotnetMap) ToFirewallInternetservicebotnetMapOutput() FirewallInternetservicebotnetMapOutput {
	return i.ToFirewallInternetservicebotnetMapOutputWithContext(context.Background())
}

func (i FirewallInternetservicebotnetMap) ToFirewallInternetservicebotnetMapOutputWithContext(ctx context.Context) FirewallInternetservicebotnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetservicebotnetMapOutput)
}

type FirewallInternetservicebotnetOutput struct{ *pulumi.OutputState }

func (FirewallInternetservicebotnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetservicebotnet)(nil)).Elem()
}

func (o FirewallInternetservicebotnetOutput) ToFirewallInternetservicebotnetOutput() FirewallInternetservicebotnetOutput {
	return o
}

func (o FirewallInternetservicebotnetOutput) ToFirewallInternetservicebotnetOutputWithContext(ctx context.Context) FirewallInternetservicebotnetOutput {
	return o
}

// Internet Service Botnet ID.
func (o FirewallInternetservicebotnetOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallInternetservicebotnet) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Internet Service Botnet name.
func (o FirewallInternetservicebotnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetservicebotnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallInternetservicebotnetOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetservicebotnet) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetservicebotnetArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetservicebotnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetservicebotnet)(nil)).Elem()
}

func (o FirewallInternetservicebotnetArrayOutput) ToFirewallInternetservicebotnetArrayOutput() FirewallInternetservicebotnetArrayOutput {
	return o
}

func (o FirewallInternetservicebotnetArrayOutput) ToFirewallInternetservicebotnetArrayOutputWithContext(ctx context.Context) FirewallInternetservicebotnetArrayOutput {
	return o
}

func (o FirewallInternetservicebotnetArrayOutput) Index(i pulumi.IntInput) FirewallInternetservicebotnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetservicebotnet {
		return vs[0].([]*FirewallInternetservicebotnet)[vs[1].(int)]
	}).(FirewallInternetservicebotnetOutput)
}

type FirewallInternetservicebotnetMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetservicebotnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetservicebotnet)(nil)).Elem()
}

func (o FirewallInternetservicebotnetMapOutput) ToFirewallInternetservicebotnetMapOutput() FirewallInternetservicebotnetMapOutput {
	return o
}

func (o FirewallInternetservicebotnetMapOutput) ToFirewallInternetservicebotnetMapOutputWithContext(ctx context.Context) FirewallInternetservicebotnetMapOutput {
	return o
}

func (o FirewallInternetservicebotnetMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetservicebotnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetservicebotnet {
		return vs[0].(map[string]*FirewallInternetservicebotnet)[vs[1].(string)]
	}).(FirewallInternetservicebotnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetservicebotnetInput)(nil)).Elem(), &FirewallInternetservicebotnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetservicebotnetArrayInput)(nil)).Elem(), FirewallInternetservicebotnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetservicebotnetMapInput)(nil)).Elem(), FirewallInternetservicebotnetMap{})
	pulumi.RegisterOutputType(FirewallInternetservicebotnetOutput{})
	pulumi.RegisterOutputType(FirewallInternetservicebotnetArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetservicebotnetMapOutput{})
}