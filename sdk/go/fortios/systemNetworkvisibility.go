// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure network visibility settings.
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
//			_, err := fortios.NewSystemNetworkvisibility(ctx, "trname", &fortios.SystemNetworkvisibilityArgs{
//				DestinationHostnameVisibility: pulumi.String("enable"),
//				DestinationLocation:           pulumi.String("enable"),
//				DestinationVisibility:         pulumi.String("enable"),
//				HostnameLimit:                 pulumi.Int(5000),
//				HostnameTtl:                   pulumi.Int(86400),
//				SourceLocation:                pulumi.String("enable"),
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
// # System NetworkVisibility can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemNetworkvisibility:SystemNetworkvisibility labelname SystemNetworkVisibility
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemNetworkvisibility:SystemNetworkvisibility labelname SystemNetworkVisibility
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemNetworkvisibility struct {
	pulumi.CustomResourceState

	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility pulumi.StringOutput `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation pulumi.StringOutput `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility pulumi.StringOutput `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit pulumi.IntOutput `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl pulumi.IntOutput `pulumi:"hostnameTtl"`
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation pulumi.StringOutput `pulumi:"sourceLocation"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemNetworkvisibility registers a new resource with the given unique name, arguments, and options.
func NewSystemNetworkvisibility(ctx *pulumi.Context,
	name string, args *SystemNetworkvisibilityArgs, opts ...pulumi.ResourceOption) (*SystemNetworkvisibility, error) {
	if args == nil {
		args = &SystemNetworkvisibilityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemNetworkvisibility
	err := ctx.RegisterResource("fortios:index/systemNetworkvisibility:SystemNetworkvisibility", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemNetworkvisibility gets an existing SystemNetworkvisibility resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemNetworkvisibility(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemNetworkvisibilityState, opts ...pulumi.ResourceOption) (*SystemNetworkvisibility, error) {
	var resource SystemNetworkvisibility
	err := ctx.ReadResource("fortios:index/systemNetworkvisibility:SystemNetworkvisibility", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemNetworkvisibility resources.
type systemNetworkvisibilityState struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility *string `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation *string `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility *string `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit *int `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl *int `pulumi:"hostnameTtl"`
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation *string `pulumi:"sourceLocation"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemNetworkvisibilityState struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility pulumi.StringPtrInput
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation pulumi.StringPtrInput
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility pulumi.StringPtrInput
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit pulumi.IntPtrInput
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl pulumi.IntPtrInput
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemNetworkvisibilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNetworkvisibilityState)(nil)).Elem()
}

type systemNetworkvisibilityArgs struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility *string `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation *string `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility *string `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit *int `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl *int `pulumi:"hostnameTtl"`
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation *string `pulumi:"sourceLocation"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemNetworkvisibility resource.
type SystemNetworkvisibilityArgs struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility pulumi.StringPtrInput
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation pulumi.StringPtrInput
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility pulumi.StringPtrInput
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit pulumi.IntPtrInput
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl pulumi.IntPtrInput
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemNetworkvisibilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNetworkvisibilityArgs)(nil)).Elem()
}

type SystemNetworkvisibilityInput interface {
	pulumi.Input

	ToSystemNetworkvisibilityOutput() SystemNetworkvisibilityOutput
	ToSystemNetworkvisibilityOutputWithContext(ctx context.Context) SystemNetworkvisibilityOutput
}

func (*SystemNetworkvisibility) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNetworkvisibility)(nil)).Elem()
}

func (i *SystemNetworkvisibility) ToSystemNetworkvisibilityOutput() SystemNetworkvisibilityOutput {
	return i.ToSystemNetworkvisibilityOutputWithContext(context.Background())
}

func (i *SystemNetworkvisibility) ToSystemNetworkvisibilityOutputWithContext(ctx context.Context) SystemNetworkvisibilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkvisibilityOutput)
}

// SystemNetworkvisibilityArrayInput is an input type that accepts SystemNetworkvisibilityArray and SystemNetworkvisibilityArrayOutput values.
// You can construct a concrete instance of `SystemNetworkvisibilityArrayInput` via:
//
//	SystemNetworkvisibilityArray{ SystemNetworkvisibilityArgs{...} }
type SystemNetworkvisibilityArrayInput interface {
	pulumi.Input

	ToSystemNetworkvisibilityArrayOutput() SystemNetworkvisibilityArrayOutput
	ToSystemNetworkvisibilityArrayOutputWithContext(context.Context) SystemNetworkvisibilityArrayOutput
}

type SystemNetworkvisibilityArray []SystemNetworkvisibilityInput

func (SystemNetworkvisibilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNetworkvisibility)(nil)).Elem()
}

func (i SystemNetworkvisibilityArray) ToSystemNetworkvisibilityArrayOutput() SystemNetworkvisibilityArrayOutput {
	return i.ToSystemNetworkvisibilityArrayOutputWithContext(context.Background())
}

func (i SystemNetworkvisibilityArray) ToSystemNetworkvisibilityArrayOutputWithContext(ctx context.Context) SystemNetworkvisibilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkvisibilityArrayOutput)
}

// SystemNetworkvisibilityMapInput is an input type that accepts SystemNetworkvisibilityMap and SystemNetworkvisibilityMapOutput values.
// You can construct a concrete instance of `SystemNetworkvisibilityMapInput` via:
//
//	SystemNetworkvisibilityMap{ "key": SystemNetworkvisibilityArgs{...} }
type SystemNetworkvisibilityMapInput interface {
	pulumi.Input

	ToSystemNetworkvisibilityMapOutput() SystemNetworkvisibilityMapOutput
	ToSystemNetworkvisibilityMapOutputWithContext(context.Context) SystemNetworkvisibilityMapOutput
}

type SystemNetworkvisibilityMap map[string]SystemNetworkvisibilityInput

func (SystemNetworkvisibilityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNetworkvisibility)(nil)).Elem()
}

func (i SystemNetworkvisibilityMap) ToSystemNetworkvisibilityMapOutput() SystemNetworkvisibilityMapOutput {
	return i.ToSystemNetworkvisibilityMapOutputWithContext(context.Background())
}

func (i SystemNetworkvisibilityMap) ToSystemNetworkvisibilityMapOutputWithContext(ctx context.Context) SystemNetworkvisibilityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkvisibilityMapOutput)
}

type SystemNetworkvisibilityOutput struct{ *pulumi.OutputState }

func (SystemNetworkvisibilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNetworkvisibility)(nil)).Elem()
}

func (o SystemNetworkvisibilityOutput) ToSystemNetworkvisibilityOutput() SystemNetworkvisibilityOutput {
	return o
}

func (o SystemNetworkvisibilityOutput) ToSystemNetworkvisibilityOutputWithContext(ctx context.Context) SystemNetworkvisibilityOutput {
	return o
}

// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
func (o SystemNetworkvisibilityOutput) DestinationHostnameVisibility() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.StringOutput { return v.DestinationHostnameVisibility }).(pulumi.StringOutput)
}

// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
func (o SystemNetworkvisibilityOutput) DestinationLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.StringOutput { return v.DestinationLocation }).(pulumi.StringOutput)
}

// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
func (o SystemNetworkvisibilityOutput) DestinationVisibility() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.StringOutput { return v.DestinationVisibility }).(pulumi.StringOutput)
}

// Limit of the number of hostname table entries (0 - 50000).
func (o SystemNetworkvisibilityOutput) HostnameLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.IntOutput { return v.HostnameLimit }).(pulumi.IntOutput)
}

// TTL of hostname table entries (60 - 86400).
func (o SystemNetworkvisibilityOutput) HostnameTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.IntOutput { return v.HostnameTtl }).(pulumi.IntOutput)
}

// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
func (o SystemNetworkvisibilityOutput) SourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.StringOutput { return v.SourceLocation }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemNetworkvisibilityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNetworkvisibility) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemNetworkvisibilityArrayOutput struct{ *pulumi.OutputState }

func (SystemNetworkvisibilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNetworkvisibility)(nil)).Elem()
}

func (o SystemNetworkvisibilityArrayOutput) ToSystemNetworkvisibilityArrayOutput() SystemNetworkvisibilityArrayOutput {
	return o
}

func (o SystemNetworkvisibilityArrayOutput) ToSystemNetworkvisibilityArrayOutputWithContext(ctx context.Context) SystemNetworkvisibilityArrayOutput {
	return o
}

func (o SystemNetworkvisibilityArrayOutput) Index(i pulumi.IntInput) SystemNetworkvisibilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemNetworkvisibility {
		return vs[0].([]*SystemNetworkvisibility)[vs[1].(int)]
	}).(SystemNetworkvisibilityOutput)
}

type SystemNetworkvisibilityMapOutput struct{ *pulumi.OutputState }

func (SystemNetworkvisibilityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNetworkvisibility)(nil)).Elem()
}

func (o SystemNetworkvisibilityMapOutput) ToSystemNetworkvisibilityMapOutput() SystemNetworkvisibilityMapOutput {
	return o
}

func (o SystemNetworkvisibilityMapOutput) ToSystemNetworkvisibilityMapOutputWithContext(ctx context.Context) SystemNetworkvisibilityMapOutput {
	return o
}

func (o SystemNetworkvisibilityMapOutput) MapIndex(k pulumi.StringInput) SystemNetworkvisibilityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemNetworkvisibility {
		return vs[0].(map[string]*SystemNetworkvisibility)[vs[1].(string)]
	}).(SystemNetworkvisibilityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkvisibilityInput)(nil)).Elem(), &SystemNetworkvisibility{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkvisibilityArrayInput)(nil)).Elem(), SystemNetworkvisibilityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkvisibilityMapInput)(nil)).Elem(), SystemNetworkvisibilityMap{})
	pulumi.RegisterOutputType(SystemNetworkvisibilityOutput{})
	pulumi.RegisterOutputType(SystemNetworkvisibilityArrayOutput{})
	pulumi.RegisterOutputType(SystemNetworkvisibilityMapOutput{})
}
