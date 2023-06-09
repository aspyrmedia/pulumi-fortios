// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `<= 6.2.0`.
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
//			_, err := fortios.NewDlpFpsensitivity(ctx, "trname", nil)
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
// # Dlp FpSensitivity can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/dlpFpsensitivity:DlpFpsensitivity labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/dlpFpsensitivity:DlpFpsensitivity labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type DlpFpsensitivity struct {
	pulumi.CustomResourceState

	// DLP Sensitivity Levels.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpFpsensitivity registers a new resource with the given unique name, arguments, and options.
func NewDlpFpsensitivity(ctx *pulumi.Context,
	name string, args *DlpFpsensitivityArgs, opts ...pulumi.ResourceOption) (*DlpFpsensitivity, error) {
	if args == nil {
		args = &DlpFpsensitivityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource DlpFpsensitivity
	err := ctx.RegisterResource("fortios:index/dlpFpsensitivity:DlpFpsensitivity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpFpsensitivity gets an existing DlpFpsensitivity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpFpsensitivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpFpsensitivityState, opts ...pulumi.ResourceOption) (*DlpFpsensitivity, error) {
	var resource DlpFpsensitivity
	err := ctx.ReadResource("fortios:index/dlpFpsensitivity:DlpFpsensitivity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpFpsensitivity resources.
type dlpFpsensitivityState struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpFpsensitivityState struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpFpsensitivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFpsensitivityState)(nil)).Elem()
}

type dlpFpsensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpFpsensitivity resource.
type DlpFpsensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpFpsensitivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFpsensitivityArgs)(nil)).Elem()
}

type DlpFpsensitivityInput interface {
	pulumi.Input

	ToDlpFpsensitivityOutput() DlpFpsensitivityOutput
	ToDlpFpsensitivityOutputWithContext(ctx context.Context) DlpFpsensitivityOutput
}

func (*DlpFpsensitivity) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFpsensitivity)(nil)).Elem()
}

func (i *DlpFpsensitivity) ToDlpFpsensitivityOutput() DlpFpsensitivityOutput {
	return i.ToDlpFpsensitivityOutputWithContext(context.Background())
}

func (i *DlpFpsensitivity) ToDlpFpsensitivityOutputWithContext(ctx context.Context) DlpFpsensitivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpsensitivityOutput)
}

// DlpFpsensitivityArrayInput is an input type that accepts DlpFpsensitivityArray and DlpFpsensitivityArrayOutput values.
// You can construct a concrete instance of `DlpFpsensitivityArrayInput` via:
//
//	DlpFpsensitivityArray{ DlpFpsensitivityArgs{...} }
type DlpFpsensitivityArrayInput interface {
	pulumi.Input

	ToDlpFpsensitivityArrayOutput() DlpFpsensitivityArrayOutput
	ToDlpFpsensitivityArrayOutputWithContext(context.Context) DlpFpsensitivityArrayOutput
}

type DlpFpsensitivityArray []DlpFpsensitivityInput

func (DlpFpsensitivityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpFpsensitivity)(nil)).Elem()
}

func (i DlpFpsensitivityArray) ToDlpFpsensitivityArrayOutput() DlpFpsensitivityArrayOutput {
	return i.ToDlpFpsensitivityArrayOutputWithContext(context.Background())
}

func (i DlpFpsensitivityArray) ToDlpFpsensitivityArrayOutputWithContext(ctx context.Context) DlpFpsensitivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpsensitivityArrayOutput)
}

// DlpFpsensitivityMapInput is an input type that accepts DlpFpsensitivityMap and DlpFpsensitivityMapOutput values.
// You can construct a concrete instance of `DlpFpsensitivityMapInput` via:
//
//	DlpFpsensitivityMap{ "key": DlpFpsensitivityArgs{...} }
type DlpFpsensitivityMapInput interface {
	pulumi.Input

	ToDlpFpsensitivityMapOutput() DlpFpsensitivityMapOutput
	ToDlpFpsensitivityMapOutputWithContext(context.Context) DlpFpsensitivityMapOutput
}

type DlpFpsensitivityMap map[string]DlpFpsensitivityInput

func (DlpFpsensitivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpFpsensitivity)(nil)).Elem()
}

func (i DlpFpsensitivityMap) ToDlpFpsensitivityMapOutput() DlpFpsensitivityMapOutput {
	return i.ToDlpFpsensitivityMapOutputWithContext(context.Background())
}

func (i DlpFpsensitivityMap) ToDlpFpsensitivityMapOutputWithContext(ctx context.Context) DlpFpsensitivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpsensitivityMapOutput)
}

type DlpFpsensitivityOutput struct{ *pulumi.OutputState }

func (DlpFpsensitivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFpsensitivity)(nil)).Elem()
}

func (o DlpFpsensitivityOutput) ToDlpFpsensitivityOutput() DlpFpsensitivityOutput {
	return o
}

func (o DlpFpsensitivityOutput) ToDlpFpsensitivityOutputWithContext(ctx context.Context) DlpFpsensitivityOutput {
	return o
}

// DLP Sensitivity Levels.
func (o DlpFpsensitivityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpFpsensitivity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DlpFpsensitivityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpFpsensitivity) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DlpFpsensitivityArrayOutput struct{ *pulumi.OutputState }

func (DlpFpsensitivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpFpsensitivity)(nil)).Elem()
}

func (o DlpFpsensitivityArrayOutput) ToDlpFpsensitivityArrayOutput() DlpFpsensitivityArrayOutput {
	return o
}

func (o DlpFpsensitivityArrayOutput) ToDlpFpsensitivityArrayOutputWithContext(ctx context.Context) DlpFpsensitivityArrayOutput {
	return o
}

func (o DlpFpsensitivityArrayOutput) Index(i pulumi.IntInput) DlpFpsensitivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpFpsensitivity {
		return vs[0].([]*DlpFpsensitivity)[vs[1].(int)]
	}).(DlpFpsensitivityOutput)
}

type DlpFpsensitivityMapOutput struct{ *pulumi.OutputState }

func (DlpFpsensitivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpFpsensitivity)(nil)).Elem()
}

func (o DlpFpsensitivityMapOutput) ToDlpFpsensitivityMapOutput() DlpFpsensitivityMapOutput {
	return o
}

func (o DlpFpsensitivityMapOutput) ToDlpFpsensitivityMapOutputWithContext(ctx context.Context) DlpFpsensitivityMapOutput {
	return o
}

func (o DlpFpsensitivityMapOutput) MapIndex(k pulumi.StringInput) DlpFpsensitivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpFpsensitivity {
		return vs[0].(map[string]*DlpFpsensitivity)[vs[1].(string)]
	}).(DlpFpsensitivityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFpsensitivityInput)(nil)).Elem(), &DlpFpsensitivity{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFpsensitivityArrayInput)(nil)).Elem(), DlpFpsensitivityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFpsensitivityMapInput)(nil)).Elem(), DlpFpsensitivityMap{})
	pulumi.RegisterOutputType(DlpFpsensitivityOutput{})
	pulumi.RegisterOutputType(DlpFpsensitivityArrayOutput{})
	pulumi.RegisterOutputType(DlpFpsensitivityMapOutput{})
}
