// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Fortinet Single Sign On (FSSO) server.
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
//			_, err := fortios.NewSystemFssopolling(ctx, "trname", &fortios.SystemFssopollingArgs{
//				Authentication: pulumi.String("disable"),
//				ListeningPort:  pulumi.Int(8000),
//				Status:         pulumi.String("enable"),
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
// # System FssoPolling can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemFssopolling:SystemFssopolling labelname SystemFssoPolling
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemFssopolling:SystemFssopolling labelname SystemFssoPolling
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemFssopolling struct {
	pulumi.CustomResourceState

	// Password to connect to FSSO Agent.
	AuthPassword pulumi.StringPtrOutput `pulumi:"authPassword"`
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`
	// Listening port to accept clients (1 - 65535).
	ListeningPort pulumi.IntOutput `pulumi:"listeningPort"`
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFssopolling registers a new resource with the given unique name, arguments, and options.
func NewSystemFssopolling(ctx *pulumi.Context,
	name string, args *SystemFssopollingArgs, opts ...pulumi.ResourceOption) (*SystemFssopolling, error) {
	if args == nil {
		args = &SystemFssopollingArgs{}
	}

	if args.AuthPassword != nil {
		args.AuthPassword = pulumi.ToSecret(args.AuthPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authPassword",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemFssopolling
	err := ctx.RegisterResource("fortios:index/systemFssopolling:SystemFssopolling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFssopolling gets an existing SystemFssopolling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFssopolling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFssopollingState, opts ...pulumi.ResourceOption) (*SystemFssopolling, error) {
	var resource SystemFssopolling
	err := ctx.ReadResource("fortios:index/systemFssopolling:SystemFssopolling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFssopolling resources.
type systemFssopollingState struct {
	// Password to connect to FSSO Agent.
	AuthPassword *string `pulumi:"authPassword"`
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication *string `pulumi:"authentication"`
	// Listening port to accept clients (1 - 65535).
	ListeningPort *int `pulumi:"listeningPort"`
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemFssopollingState struct {
	// Password to connect to FSSO Agent.
	AuthPassword pulumi.StringPtrInput
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringPtrInput
	// Listening port to accept clients (1 - 65535).
	ListeningPort pulumi.IntPtrInput
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFssopollingState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFssopollingState)(nil)).Elem()
}

type systemFssopollingArgs struct {
	// Password to connect to FSSO Agent.
	AuthPassword *string `pulumi:"authPassword"`
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication *string `pulumi:"authentication"`
	// Listening port to accept clients (1 - 65535).
	ListeningPort *int `pulumi:"listeningPort"`
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFssopolling resource.
type SystemFssopollingArgs struct {
	// Password to connect to FSSO Agent.
	AuthPassword pulumi.StringPtrInput
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringPtrInput
	// Listening port to accept clients (1 - 65535).
	ListeningPort pulumi.IntPtrInput
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFssopollingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFssopollingArgs)(nil)).Elem()
}

type SystemFssopollingInput interface {
	pulumi.Input

	ToSystemFssopollingOutput() SystemFssopollingOutput
	ToSystemFssopollingOutputWithContext(ctx context.Context) SystemFssopollingOutput
}

func (*SystemFssopolling) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFssopolling)(nil)).Elem()
}

func (i *SystemFssopolling) ToSystemFssopollingOutput() SystemFssopollingOutput {
	return i.ToSystemFssopollingOutputWithContext(context.Background())
}

func (i *SystemFssopolling) ToSystemFssopollingOutputWithContext(ctx context.Context) SystemFssopollingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFssopollingOutput)
}

// SystemFssopollingArrayInput is an input type that accepts SystemFssopollingArray and SystemFssopollingArrayOutput values.
// You can construct a concrete instance of `SystemFssopollingArrayInput` via:
//
//	SystemFssopollingArray{ SystemFssopollingArgs{...} }
type SystemFssopollingArrayInput interface {
	pulumi.Input

	ToSystemFssopollingArrayOutput() SystemFssopollingArrayOutput
	ToSystemFssopollingArrayOutputWithContext(context.Context) SystemFssopollingArrayOutput
}

type SystemFssopollingArray []SystemFssopollingInput

func (SystemFssopollingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFssopolling)(nil)).Elem()
}

func (i SystemFssopollingArray) ToSystemFssopollingArrayOutput() SystemFssopollingArrayOutput {
	return i.ToSystemFssopollingArrayOutputWithContext(context.Background())
}

func (i SystemFssopollingArray) ToSystemFssopollingArrayOutputWithContext(ctx context.Context) SystemFssopollingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFssopollingArrayOutput)
}

// SystemFssopollingMapInput is an input type that accepts SystemFssopollingMap and SystemFssopollingMapOutput values.
// You can construct a concrete instance of `SystemFssopollingMapInput` via:
//
//	SystemFssopollingMap{ "key": SystemFssopollingArgs{...} }
type SystemFssopollingMapInput interface {
	pulumi.Input

	ToSystemFssopollingMapOutput() SystemFssopollingMapOutput
	ToSystemFssopollingMapOutputWithContext(context.Context) SystemFssopollingMapOutput
}

type SystemFssopollingMap map[string]SystemFssopollingInput

func (SystemFssopollingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFssopolling)(nil)).Elem()
}

func (i SystemFssopollingMap) ToSystemFssopollingMapOutput() SystemFssopollingMapOutput {
	return i.ToSystemFssopollingMapOutputWithContext(context.Background())
}

func (i SystemFssopollingMap) ToSystemFssopollingMapOutputWithContext(ctx context.Context) SystemFssopollingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFssopollingMapOutput)
}

type SystemFssopollingOutput struct{ *pulumi.OutputState }

func (SystemFssopollingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFssopolling)(nil)).Elem()
}

func (o SystemFssopollingOutput) ToSystemFssopollingOutput() SystemFssopollingOutput {
	return o
}

func (o SystemFssopollingOutput) ToSystemFssopollingOutputWithContext(ctx context.Context) SystemFssopollingOutput {
	return o
}

// Password to connect to FSSO Agent.
func (o SystemFssopollingOutput) AuthPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFssopolling) pulumi.StringPtrOutput { return v.AuthPassword }).(pulumi.StringPtrOutput)
}

// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
func (o SystemFssopollingOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFssopolling) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

// Listening port to accept clients (1 - 65535).
func (o SystemFssopollingOutput) ListeningPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFssopolling) pulumi.IntOutput { return v.ListeningPort }).(pulumi.IntOutput)
}

// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
func (o SystemFssopollingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFssopolling) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemFssopollingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFssopolling) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFssopollingArrayOutput struct{ *pulumi.OutputState }

func (SystemFssopollingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFssopolling)(nil)).Elem()
}

func (o SystemFssopollingArrayOutput) ToSystemFssopollingArrayOutput() SystemFssopollingArrayOutput {
	return o
}

func (o SystemFssopollingArrayOutput) ToSystemFssopollingArrayOutputWithContext(ctx context.Context) SystemFssopollingArrayOutput {
	return o
}

func (o SystemFssopollingArrayOutput) Index(i pulumi.IntInput) SystemFssopollingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFssopolling {
		return vs[0].([]*SystemFssopolling)[vs[1].(int)]
	}).(SystemFssopollingOutput)
}

type SystemFssopollingMapOutput struct{ *pulumi.OutputState }

func (SystemFssopollingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFssopolling)(nil)).Elem()
}

func (o SystemFssopollingMapOutput) ToSystemFssopollingMapOutput() SystemFssopollingMapOutput {
	return o
}

func (o SystemFssopollingMapOutput) ToSystemFssopollingMapOutputWithContext(ctx context.Context) SystemFssopollingMapOutput {
	return o
}

func (o SystemFssopollingMapOutput) MapIndex(k pulumi.StringInput) SystemFssopollingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFssopolling {
		return vs[0].(map[string]*SystemFssopolling)[vs[1].(string)]
	}).(SystemFssopollingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFssopollingInput)(nil)).Elem(), &SystemFssopolling{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFssopollingArrayInput)(nil)).Elem(), SystemFssopollingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFssopollingMapInput)(nil)).Elem(), SystemFssopollingMap{})
	pulumi.RegisterOutputType(SystemFssopollingOutput{})
	pulumi.RegisterOutputType(SystemFssopollingArrayOutput{})
	pulumi.RegisterOutputType(SystemFssopollingMapOutput{})
}
