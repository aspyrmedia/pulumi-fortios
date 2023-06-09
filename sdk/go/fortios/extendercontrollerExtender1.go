// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Extender controller configuration.
//
// > The resource applies to FortiOS Version >= 6.4.2. For FortiOS Version < 6.4.2, see `ExtendercontrollerExtender`.
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
//			_, err := fortios.NewExtendercontrollerExtender1(ctx, "trname", &fortios.ExtendercontrollerExtender1Args{
//				Authorized: pulumi.String("disable"),
//				ControllerReport: &fortios.ExtendercontrollerExtender1ControllerReportArgs{
//					Interval:        pulumi.Int(300),
//					SignalThreshold: pulumi.Int(10),
//					Status:          pulumi.String("disable"),
//				},
//				ExtName: pulumi.String("2932"),
//				Fosid:   pulumi.String("FX201E5919004031"),
//				Modem1: &fortios.ExtendercontrollerExtender1Modem1Args{
//					AutoSwitch: &fortios.ExtendercontrollerExtender1Modem1AutoSwitchArgs{
//						Dataplan:            pulumi.String("disable"),
//						Disconnect:          pulumi.String("disable"),
//						DisconnectPeriod:    pulumi.Int(600),
//						DisconnectThreshold: pulumi.Int(3),
//						Signal:              pulumi.String("disable"),
//						SwitchBack:          pulumi.String("timer"),
//						SwitchBackTime:      pulumi.String("00:01"),
//						SwitchBackTimer:     pulumi.Int(86400),
//					},
//					ConnStatus:    pulumi.Int(0),
//					DefaultSim:    pulumi.String("sim2"),
//					Gps:           pulumi.String("enable"),
//					RedundantIntf: pulumi.String("s1"),
//					RedundantMode: pulumi.String("enable"),
//					Sim1Pin:       pulumi.String("disable"),
//					Sim1PinCode:   pulumi.String("testpincode"),
//					Sim2Pin:       pulumi.String("disable"),
//				},
//				Modem2: &fortios.ExtendercontrollerExtender1Modem2Args{
//					AutoSwitch: &fortios.ExtendercontrollerExtender1Modem2AutoSwitchArgs{
//						Dataplan:            pulumi.String("disable"),
//						Disconnect:          pulumi.String("disable"),
//						DisconnectPeriod:    pulumi.Int(600),
//						DisconnectThreshold: pulumi.Int(3),
//						Signal:              pulumi.String("disable"),
//						SwitchBackTime:      pulumi.String("00:01"),
//						SwitchBackTimer:     pulumi.Int(86400),
//					},
//					ConnStatus:    pulumi.Int(0),
//					DefaultSim:    pulumi.String("sim1"),
//					Gps:           pulumi.String("enable"),
//					RedundantMode: pulumi.String("disable"),
//					Sim1Pin:       pulumi.String("disable"),
//					Sim2Pin:       pulumi.String("disable"),
//				},
//				Vdom: pulumi.Int(0),
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
// # ExtenderController Extender1 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/extendercontrollerExtender1:ExtendercontrollerExtender1 labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/extendercontrollerExtender1:ExtendercontrollerExtender1 labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type ExtendercontrollerExtender1 struct {
	pulumi.CustomResourceState

	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringOutput `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport ExtendercontrollerExtender1ControllerReportOutput `pulumi:"controllerReport"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// FortiExtender name.
	ExtName pulumi.StringOutput `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrOutput `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 ExtendercontrollerExtender1Modem1Output `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 ExtendercontrollerExtender1Modem2Output `pulumi:"modem2"`
	// FortiExtender entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VDOM
	Vdom pulumi.IntOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtendercontrollerExtender1 registers a new resource with the given unique name, arguments, and options.
func NewExtendercontrollerExtender1(ctx *pulumi.Context,
	name string, args *ExtendercontrollerExtender1Args, opts ...pulumi.ResourceOption) (*ExtendercontrollerExtender1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorized == nil {
		return nil, errors.New("invalid value for required argument 'Authorized'")
	}
	if args.LoginPassword != nil {
		args.LoginPassword = pulumi.ToSecret(args.LoginPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"loginPassword",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource ExtendercontrollerExtender1
	err := ctx.RegisterResource("fortios:index/extendercontrollerExtender1:ExtendercontrollerExtender1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtendercontrollerExtender1 gets an existing ExtendercontrollerExtender1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtendercontrollerExtender1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtendercontrollerExtender1State, opts ...pulumi.ResourceOption) (*ExtendercontrollerExtender1, error) {
	var resource ExtendercontrollerExtender1
	err := ctx.ReadResource("fortios:index/extendercontrollerExtender1:ExtendercontrollerExtender1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtendercontrollerExtender1 resources.
type extendercontrollerExtender1State struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport *ExtendercontrollerExtender1ControllerReport `pulumi:"controllerReport"`
	// Description.
	Description *string `pulumi:"description"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiExtender login password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 *ExtendercontrollerExtender1Modem1 `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 *ExtendercontrollerExtender1Modem2 `pulumi:"modem2"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// VDOM
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExtendercontrollerExtender1State struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport ExtendercontrollerExtender1ControllerReportPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrInput
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 ExtendercontrollerExtender1Modem1PtrInput
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 ExtendercontrollerExtender1Modem2PtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// VDOM
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtendercontrollerExtender1State) ElementType() reflect.Type {
	return reflect.TypeOf((*extendercontrollerExtender1State)(nil)).Elem()
}

type extendercontrollerExtender1Args struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized string `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport *ExtendercontrollerExtender1ControllerReport `pulumi:"controllerReport"`
	// Description.
	Description *string `pulumi:"description"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiExtender login password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 *ExtendercontrollerExtender1Modem1 `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 *ExtendercontrollerExtender1Modem2 `pulumi:"modem2"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// VDOM
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ExtendercontrollerExtender1 resource.
type ExtendercontrollerExtender1Args struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringInput
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport ExtendercontrollerExtender1ControllerReportPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrInput
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 ExtendercontrollerExtender1Modem1PtrInput
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 ExtendercontrollerExtender1Modem2PtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// VDOM
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtendercontrollerExtender1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*extendercontrollerExtender1Args)(nil)).Elem()
}

type ExtendercontrollerExtender1Input interface {
	pulumi.Input

	ToExtendercontrollerExtender1Output() ExtendercontrollerExtender1Output
	ToExtendercontrollerExtender1OutputWithContext(ctx context.Context) ExtendercontrollerExtender1Output
}

func (*ExtendercontrollerExtender1) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendercontrollerExtender1)(nil)).Elem()
}

func (i *ExtendercontrollerExtender1) ToExtendercontrollerExtender1Output() ExtendercontrollerExtender1Output {
	return i.ToExtendercontrollerExtender1OutputWithContext(context.Background())
}

func (i *ExtendercontrollerExtender1) ToExtendercontrollerExtender1OutputWithContext(ctx context.Context) ExtendercontrollerExtender1Output {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendercontrollerExtender1Output)
}

// ExtendercontrollerExtender1ArrayInput is an input type that accepts ExtendercontrollerExtender1Array and ExtendercontrollerExtender1ArrayOutput values.
// You can construct a concrete instance of `ExtendercontrollerExtender1ArrayInput` via:
//
//	ExtendercontrollerExtender1Array{ ExtendercontrollerExtender1Args{...} }
type ExtendercontrollerExtender1ArrayInput interface {
	pulumi.Input

	ToExtendercontrollerExtender1ArrayOutput() ExtendercontrollerExtender1ArrayOutput
	ToExtendercontrollerExtender1ArrayOutputWithContext(context.Context) ExtendercontrollerExtender1ArrayOutput
}

type ExtendercontrollerExtender1Array []ExtendercontrollerExtender1Input

func (ExtendercontrollerExtender1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtendercontrollerExtender1)(nil)).Elem()
}

func (i ExtendercontrollerExtender1Array) ToExtendercontrollerExtender1ArrayOutput() ExtendercontrollerExtender1ArrayOutput {
	return i.ToExtendercontrollerExtender1ArrayOutputWithContext(context.Background())
}

func (i ExtendercontrollerExtender1Array) ToExtendercontrollerExtender1ArrayOutputWithContext(ctx context.Context) ExtendercontrollerExtender1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendercontrollerExtender1ArrayOutput)
}

// ExtendercontrollerExtender1MapInput is an input type that accepts ExtendercontrollerExtender1Map and ExtendercontrollerExtender1MapOutput values.
// You can construct a concrete instance of `ExtendercontrollerExtender1MapInput` via:
//
//	ExtendercontrollerExtender1Map{ "key": ExtendercontrollerExtender1Args{...} }
type ExtendercontrollerExtender1MapInput interface {
	pulumi.Input

	ToExtendercontrollerExtender1MapOutput() ExtendercontrollerExtender1MapOutput
	ToExtendercontrollerExtender1MapOutputWithContext(context.Context) ExtendercontrollerExtender1MapOutput
}

type ExtendercontrollerExtender1Map map[string]ExtendercontrollerExtender1Input

func (ExtendercontrollerExtender1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtendercontrollerExtender1)(nil)).Elem()
}

func (i ExtendercontrollerExtender1Map) ToExtendercontrollerExtender1MapOutput() ExtendercontrollerExtender1MapOutput {
	return i.ToExtendercontrollerExtender1MapOutputWithContext(context.Background())
}

func (i ExtendercontrollerExtender1Map) ToExtendercontrollerExtender1MapOutputWithContext(ctx context.Context) ExtendercontrollerExtender1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendercontrollerExtender1MapOutput)
}

type ExtendercontrollerExtender1Output struct{ *pulumi.OutputState }

func (ExtendercontrollerExtender1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendercontrollerExtender1)(nil)).Elem()
}

func (o ExtendercontrollerExtender1Output) ToExtendercontrollerExtender1Output() ExtendercontrollerExtender1Output {
	return o
}

func (o ExtendercontrollerExtender1Output) ToExtendercontrollerExtender1OutputWithContext(ctx context.Context) ExtendercontrollerExtender1Output {
	return o
}

// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
func (o ExtendercontrollerExtender1Output) Authorized() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringOutput { return v.Authorized }).(pulumi.StringOutput)
}

// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
func (o ExtendercontrollerExtender1Output) ControllerReport() ExtendercontrollerExtender1ControllerReportOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) ExtendercontrollerExtender1ControllerReportOutput {
		return v.ControllerReport
	}).(ExtendercontrollerExtender1ControllerReportOutput)
}

// Description.
func (o ExtendercontrollerExtender1Output) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// FortiExtender name.
func (o ExtendercontrollerExtender1Output) ExtName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringOutput { return v.ExtName }).(pulumi.StringOutput)
}

// FortiExtender serial number.
func (o ExtendercontrollerExtender1Output) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// FortiExtender login password.
func (o ExtendercontrollerExtender1Output) LoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringPtrOutput { return v.LoginPassword }).(pulumi.StringPtrOutput)
}

// Configuration options for modem 1. The structure of `modem1` block is documented below.
func (o ExtendercontrollerExtender1Output) Modem1() ExtendercontrollerExtender1Modem1Output {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) ExtendercontrollerExtender1Modem1Output { return v.Modem1 }).(ExtendercontrollerExtender1Modem1Output)
}

// Configuration options for modem 2. The structure of `modem2` block is documented below.
func (o ExtendercontrollerExtender1Output) Modem2() ExtendercontrollerExtender1Modem2Output {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) ExtendercontrollerExtender1Modem2Output { return v.Modem2 }).(ExtendercontrollerExtender1Modem2Output)
}

// FortiExtender entry name.
func (o ExtendercontrollerExtender1Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// VDOM
func (o ExtendercontrollerExtender1Output) Vdom() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.IntOutput { return v.Vdom }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExtendercontrollerExtender1Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendercontrollerExtender1) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExtendercontrollerExtender1ArrayOutput struct{ *pulumi.OutputState }

func (ExtendercontrollerExtender1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtendercontrollerExtender1)(nil)).Elem()
}

func (o ExtendercontrollerExtender1ArrayOutput) ToExtendercontrollerExtender1ArrayOutput() ExtendercontrollerExtender1ArrayOutput {
	return o
}

func (o ExtendercontrollerExtender1ArrayOutput) ToExtendercontrollerExtender1ArrayOutputWithContext(ctx context.Context) ExtendercontrollerExtender1ArrayOutput {
	return o
}

func (o ExtendercontrollerExtender1ArrayOutput) Index(i pulumi.IntInput) ExtendercontrollerExtender1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtendercontrollerExtender1 {
		return vs[0].([]*ExtendercontrollerExtender1)[vs[1].(int)]
	}).(ExtendercontrollerExtender1Output)
}

type ExtendercontrollerExtender1MapOutput struct{ *pulumi.OutputState }

func (ExtendercontrollerExtender1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtendercontrollerExtender1)(nil)).Elem()
}

func (o ExtendercontrollerExtender1MapOutput) ToExtendercontrollerExtender1MapOutput() ExtendercontrollerExtender1MapOutput {
	return o
}

func (o ExtendercontrollerExtender1MapOutput) ToExtendercontrollerExtender1MapOutputWithContext(ctx context.Context) ExtendercontrollerExtender1MapOutput {
	return o
}

func (o ExtendercontrollerExtender1MapOutput) MapIndex(k pulumi.StringInput) ExtendercontrollerExtender1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtendercontrollerExtender1 {
		return vs[0].(map[string]*ExtendercontrollerExtender1)[vs[1].(string)]
	}).(ExtendercontrollerExtender1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendercontrollerExtender1Input)(nil)).Elem(), &ExtendercontrollerExtender1{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendercontrollerExtender1ArrayInput)(nil)).Elem(), ExtendercontrollerExtender1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendercontrollerExtender1MapInput)(nil)).Elem(), ExtendercontrollerExtender1Map{})
	pulumi.RegisterOutputType(ExtendercontrollerExtender1Output{})
	pulumi.RegisterOutputType(ExtendercontrollerExtender1ArrayOutput{})
	pulumi.RegisterOutputType(ExtendercontrollerExtender1MapOutput{})
}
