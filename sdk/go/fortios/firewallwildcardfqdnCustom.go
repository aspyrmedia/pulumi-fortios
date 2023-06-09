// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Config global/VDOM Wildcard FQDN address.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewFirewallwildcardfqdnCustom(ctx, "trname", &fortios.FirewallwildcardfqdnCustomArgs{
//				Color:        pulumi.Int(0),
//				Visibility:   pulumi.String("enable"),
//				WildcardFqdn: pulumi.String("*.go.google.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # FirewallWildcardFqdn Custom can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallwildcardfqdnCustom:FirewallwildcardfqdnCustom labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallwildcardfqdnCustom:FirewallwildcardfqdnCustom labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallwildcardfqdnCustom struct {
	pulumi.CustomResourceState

	// GUI icon color.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
	// Wildcard FQDN.
	WildcardFqdn pulumi.StringOutput `pulumi:"wildcardFqdn"`
}

// NewFirewallwildcardfqdnCustom registers a new resource with the given unique name, arguments, and options.
func NewFirewallwildcardfqdnCustom(ctx *pulumi.Context,
	name string, args *FirewallwildcardfqdnCustomArgs, opts ...pulumi.ResourceOption) (*FirewallwildcardfqdnCustom, error) {
	if args == nil {
		args = &FirewallwildcardfqdnCustomArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallwildcardfqdnCustom
	err := ctx.RegisterResource("fortios:index/firewallwildcardfqdnCustom:FirewallwildcardfqdnCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallwildcardfqdnCustom gets an existing FirewallwildcardfqdnCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallwildcardfqdnCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallwildcardfqdnCustomState, opts ...pulumi.ResourceOption) (*FirewallwildcardfqdnCustom, error) {
	var resource FirewallwildcardfqdnCustom
	err := ctx.ReadResource("fortios:index/firewallwildcardfqdnCustom:FirewallwildcardfqdnCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallwildcardfqdnCustom resources.
type firewallwildcardfqdnCustomState struct {
	// GUI icon color.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Address name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
	// Wildcard FQDN.
	WildcardFqdn *string `pulumi:"wildcardFqdn"`
}

type FirewallwildcardfqdnCustomState struct {
	// GUI icon color.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
	// Wildcard FQDN.
	WildcardFqdn pulumi.StringPtrInput
}

func (FirewallwildcardfqdnCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallwildcardfqdnCustomState)(nil)).Elem()
}

type firewallwildcardfqdnCustomArgs struct {
	// GUI icon color.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Address name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
	// Wildcard FQDN.
	WildcardFqdn *string `pulumi:"wildcardFqdn"`
}

// The set of arguments for constructing a FirewallwildcardfqdnCustom resource.
type FirewallwildcardfqdnCustomArgs struct {
	// GUI icon color.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
	// Wildcard FQDN.
	WildcardFqdn pulumi.StringPtrInput
}

func (FirewallwildcardfqdnCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallwildcardfqdnCustomArgs)(nil)).Elem()
}

type FirewallwildcardfqdnCustomInput interface {
	pulumi.Input

	ToFirewallwildcardfqdnCustomOutput() FirewallwildcardfqdnCustomOutput
	ToFirewallwildcardfqdnCustomOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomOutput
}

func (*FirewallwildcardfqdnCustom) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallwildcardfqdnCustom)(nil)).Elem()
}

func (i *FirewallwildcardfqdnCustom) ToFirewallwildcardfqdnCustomOutput() FirewallwildcardfqdnCustomOutput {
	return i.ToFirewallwildcardfqdnCustomOutputWithContext(context.Background())
}

func (i *FirewallwildcardfqdnCustom) ToFirewallwildcardfqdnCustomOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallwildcardfqdnCustomOutput)
}

// FirewallwildcardfqdnCustomArrayInput is an input type that accepts FirewallwildcardfqdnCustomArray and FirewallwildcardfqdnCustomArrayOutput values.
// You can construct a concrete instance of `FirewallwildcardfqdnCustomArrayInput` via:
//
//	FirewallwildcardfqdnCustomArray{ FirewallwildcardfqdnCustomArgs{...} }
type FirewallwildcardfqdnCustomArrayInput interface {
	pulumi.Input

	ToFirewallwildcardfqdnCustomArrayOutput() FirewallwildcardfqdnCustomArrayOutput
	ToFirewallwildcardfqdnCustomArrayOutputWithContext(context.Context) FirewallwildcardfqdnCustomArrayOutput
}

type FirewallwildcardfqdnCustomArray []FirewallwildcardfqdnCustomInput

func (FirewallwildcardfqdnCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallwildcardfqdnCustom)(nil)).Elem()
}

func (i FirewallwildcardfqdnCustomArray) ToFirewallwildcardfqdnCustomArrayOutput() FirewallwildcardfqdnCustomArrayOutput {
	return i.ToFirewallwildcardfqdnCustomArrayOutputWithContext(context.Background())
}

func (i FirewallwildcardfqdnCustomArray) ToFirewallwildcardfqdnCustomArrayOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallwildcardfqdnCustomArrayOutput)
}

// FirewallwildcardfqdnCustomMapInput is an input type that accepts FirewallwildcardfqdnCustomMap and FirewallwildcardfqdnCustomMapOutput values.
// You can construct a concrete instance of `FirewallwildcardfqdnCustomMapInput` via:
//
//	FirewallwildcardfqdnCustomMap{ "key": FirewallwildcardfqdnCustomArgs{...} }
type FirewallwildcardfqdnCustomMapInput interface {
	pulumi.Input

	ToFirewallwildcardfqdnCustomMapOutput() FirewallwildcardfqdnCustomMapOutput
	ToFirewallwildcardfqdnCustomMapOutputWithContext(context.Context) FirewallwildcardfqdnCustomMapOutput
}

type FirewallwildcardfqdnCustomMap map[string]FirewallwildcardfqdnCustomInput

func (FirewallwildcardfqdnCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallwildcardfqdnCustom)(nil)).Elem()
}

func (i FirewallwildcardfqdnCustomMap) ToFirewallwildcardfqdnCustomMapOutput() FirewallwildcardfqdnCustomMapOutput {
	return i.ToFirewallwildcardfqdnCustomMapOutputWithContext(context.Background())
}

func (i FirewallwildcardfqdnCustomMap) ToFirewallwildcardfqdnCustomMapOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallwildcardfqdnCustomMapOutput)
}

type FirewallwildcardfqdnCustomOutput struct{ *pulumi.OutputState }

func (FirewallwildcardfqdnCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallwildcardfqdnCustom)(nil)).Elem()
}

func (o FirewallwildcardfqdnCustomOutput) ToFirewallwildcardfqdnCustomOutput() FirewallwildcardfqdnCustomOutput {
	return o
}

func (o FirewallwildcardfqdnCustomOutput) ToFirewallwildcardfqdnCustomOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomOutput {
	return o
}

// GUI icon color.
func (o FirewallwildcardfqdnCustomOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o FirewallwildcardfqdnCustomOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Address name.
func (o FirewallwildcardfqdnCustomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallwildcardfqdnCustomOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallwildcardfqdnCustomOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable address visibility. Valid values: `enable`, `disable`.
func (o FirewallwildcardfqdnCustomOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

// Wildcard FQDN.
func (o FirewallwildcardfqdnCustomOutput) WildcardFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnCustom) pulumi.StringOutput { return v.WildcardFqdn }).(pulumi.StringOutput)
}

type FirewallwildcardfqdnCustomArrayOutput struct{ *pulumi.OutputState }

func (FirewallwildcardfqdnCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallwildcardfqdnCustom)(nil)).Elem()
}

func (o FirewallwildcardfqdnCustomArrayOutput) ToFirewallwildcardfqdnCustomArrayOutput() FirewallwildcardfqdnCustomArrayOutput {
	return o
}

func (o FirewallwildcardfqdnCustomArrayOutput) ToFirewallwildcardfqdnCustomArrayOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomArrayOutput {
	return o
}

func (o FirewallwildcardfqdnCustomArrayOutput) Index(i pulumi.IntInput) FirewallwildcardfqdnCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallwildcardfqdnCustom {
		return vs[0].([]*FirewallwildcardfqdnCustom)[vs[1].(int)]
	}).(FirewallwildcardfqdnCustomOutput)
}

type FirewallwildcardfqdnCustomMapOutput struct{ *pulumi.OutputState }

func (FirewallwildcardfqdnCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallwildcardfqdnCustom)(nil)).Elem()
}

func (o FirewallwildcardfqdnCustomMapOutput) ToFirewallwildcardfqdnCustomMapOutput() FirewallwildcardfqdnCustomMapOutput {
	return o
}

func (o FirewallwildcardfqdnCustomMapOutput) ToFirewallwildcardfqdnCustomMapOutputWithContext(ctx context.Context) FirewallwildcardfqdnCustomMapOutput {
	return o
}

func (o FirewallwildcardfqdnCustomMapOutput) MapIndex(k pulumi.StringInput) FirewallwildcardfqdnCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallwildcardfqdnCustom {
		return vs[0].(map[string]*FirewallwildcardfqdnCustom)[vs[1].(string)]
	}).(FirewallwildcardfqdnCustomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallwildcardfqdnCustomInput)(nil)).Elem(), &FirewallwildcardfqdnCustom{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallwildcardfqdnCustomArrayInput)(nil)).Elem(), FirewallwildcardfqdnCustomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallwildcardfqdnCustomMapInput)(nil)).Elem(), FirewallwildcardfqdnCustomMap{})
	pulumi.RegisterOutputType(FirewallwildcardfqdnCustomOutput{})
	pulumi.RegisterOutputType(FirewallwildcardfqdnCustomArrayOutput{})
	pulumi.RegisterOutputType(FirewallwildcardfqdnCustomMapOutput{})
}
