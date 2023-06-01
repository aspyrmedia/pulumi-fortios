// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiToken Mobile push services.
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
//			_, err := fortios.NewSystemFtmpush(ctx, "trname", &fortios.SystemFtmpushArgs{
//				ServerIp:   pulumi.String("0.0.0.0"),
//				ServerPort: pulumi.Int(4433),
//				Status:     pulumi.String("disable"),
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
// # System FtmPush can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemFtmpush:SystemFtmpush labelname SystemFtmPush
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemFtmpush:SystemFtmpush labelname SystemFtmPush
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemFtmpush struct {
	pulumi.CustomResourceState

	// IPv4 address or domain name of FortiToken Mobile push services server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
	ServerCert pulumi.StringOutput `pulumi:"serverCert"`
	// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
	ServerIp pulumi.StringOutput `pulumi:"serverIp"`
	// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFtmpush registers a new resource with the given unique name, arguments, and options.
func NewSystemFtmpush(ctx *pulumi.Context,
	name string, args *SystemFtmpushArgs, opts ...pulumi.ResourceOption) (*SystemFtmpush, error) {
	if args == nil {
		args = &SystemFtmpushArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemFtmpush
	err := ctx.RegisterResource("fortios:index/systemFtmpush:SystemFtmpush", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFtmpush gets an existing SystemFtmpush resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFtmpush(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFtmpushState, opts ...pulumi.ResourceOption) (*SystemFtmpush, error) {
	var resource SystemFtmpush
	err := ctx.ReadResource("fortios:index/systemFtmpush:SystemFtmpush", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFtmpush resources.
type systemFtmpushState struct {
	// IPv4 address or domain name of FortiToken Mobile push services server.
	Server *string `pulumi:"server"`
	// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
	ServerCert *string `pulumi:"serverCert"`
	// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
	ServerIp *string `pulumi:"serverIp"`
	// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
	ServerPort *int `pulumi:"serverPort"`
	// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemFtmpushState struct {
	// IPv4 address or domain name of FortiToken Mobile push services server.
	Server pulumi.StringPtrInput
	// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
	ServerCert pulumi.StringPtrInput
	// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
	ServerIp pulumi.StringPtrInput
	// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
	ServerPort pulumi.IntPtrInput
	// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFtmpushState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFtmpushState)(nil)).Elem()
}

type systemFtmpushArgs struct {
	// IPv4 address or domain name of FortiToken Mobile push services server.
	Server *string `pulumi:"server"`
	// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
	ServerCert *string `pulumi:"serverCert"`
	// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
	ServerIp *string `pulumi:"serverIp"`
	// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
	ServerPort *int `pulumi:"serverPort"`
	// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFtmpush resource.
type SystemFtmpushArgs struct {
	// IPv4 address or domain name of FortiToken Mobile push services server.
	Server pulumi.StringPtrInput
	// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
	ServerCert pulumi.StringPtrInput
	// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
	ServerIp pulumi.StringPtrInput
	// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
	ServerPort pulumi.IntPtrInput
	// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFtmpushArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFtmpushArgs)(nil)).Elem()
}

type SystemFtmpushInput interface {
	pulumi.Input

	ToSystemFtmpushOutput() SystemFtmpushOutput
	ToSystemFtmpushOutputWithContext(ctx context.Context) SystemFtmpushOutput
}

func (*SystemFtmpush) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFtmpush)(nil)).Elem()
}

func (i *SystemFtmpush) ToSystemFtmpushOutput() SystemFtmpushOutput {
	return i.ToSystemFtmpushOutputWithContext(context.Background())
}

func (i *SystemFtmpush) ToSystemFtmpushOutputWithContext(ctx context.Context) SystemFtmpushOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFtmpushOutput)
}

// SystemFtmpushArrayInput is an input type that accepts SystemFtmpushArray and SystemFtmpushArrayOutput values.
// You can construct a concrete instance of `SystemFtmpushArrayInput` via:
//
//	SystemFtmpushArray{ SystemFtmpushArgs{...} }
type SystemFtmpushArrayInput interface {
	pulumi.Input

	ToSystemFtmpushArrayOutput() SystemFtmpushArrayOutput
	ToSystemFtmpushArrayOutputWithContext(context.Context) SystemFtmpushArrayOutput
}

type SystemFtmpushArray []SystemFtmpushInput

func (SystemFtmpushArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFtmpush)(nil)).Elem()
}

func (i SystemFtmpushArray) ToSystemFtmpushArrayOutput() SystemFtmpushArrayOutput {
	return i.ToSystemFtmpushArrayOutputWithContext(context.Background())
}

func (i SystemFtmpushArray) ToSystemFtmpushArrayOutputWithContext(ctx context.Context) SystemFtmpushArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFtmpushArrayOutput)
}

// SystemFtmpushMapInput is an input type that accepts SystemFtmpushMap and SystemFtmpushMapOutput values.
// You can construct a concrete instance of `SystemFtmpushMapInput` via:
//
//	SystemFtmpushMap{ "key": SystemFtmpushArgs{...} }
type SystemFtmpushMapInput interface {
	pulumi.Input

	ToSystemFtmpushMapOutput() SystemFtmpushMapOutput
	ToSystemFtmpushMapOutputWithContext(context.Context) SystemFtmpushMapOutput
}

type SystemFtmpushMap map[string]SystemFtmpushInput

func (SystemFtmpushMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFtmpush)(nil)).Elem()
}

func (i SystemFtmpushMap) ToSystemFtmpushMapOutput() SystemFtmpushMapOutput {
	return i.ToSystemFtmpushMapOutputWithContext(context.Background())
}

func (i SystemFtmpushMap) ToSystemFtmpushMapOutputWithContext(ctx context.Context) SystemFtmpushMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFtmpushMapOutput)
}

type SystemFtmpushOutput struct{ *pulumi.OutputState }

func (SystemFtmpushOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFtmpush)(nil)).Elem()
}

func (o SystemFtmpushOutput) ToSystemFtmpushOutput() SystemFtmpushOutput {
	return o
}

func (o SystemFtmpushOutput) ToSystemFtmpushOutputWithContext(ctx context.Context) SystemFtmpushOutput {
	return o
}

// IPv4 address or domain name of FortiToken Mobile push services server.
func (o SystemFtmpushOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmpush) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
func (o SystemFtmpushOutput) ServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmpush) pulumi.StringOutput { return v.ServerCert }).(pulumi.StringOutput)
}

// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
func (o SystemFtmpushOutput) ServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmpush) pulumi.StringOutput { return v.ServerIp }).(pulumi.StringOutput)
}

// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
func (o SystemFtmpushOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFtmpush) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
func (o SystemFtmpushOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmpush) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemFtmpushOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFtmpush) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFtmpushArrayOutput struct{ *pulumi.OutputState }

func (SystemFtmpushArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFtmpush)(nil)).Elem()
}

func (o SystemFtmpushArrayOutput) ToSystemFtmpushArrayOutput() SystemFtmpushArrayOutput {
	return o
}

func (o SystemFtmpushArrayOutput) ToSystemFtmpushArrayOutputWithContext(ctx context.Context) SystemFtmpushArrayOutput {
	return o
}

func (o SystemFtmpushArrayOutput) Index(i pulumi.IntInput) SystemFtmpushOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFtmpush {
		return vs[0].([]*SystemFtmpush)[vs[1].(int)]
	}).(SystemFtmpushOutput)
}

type SystemFtmpushMapOutput struct{ *pulumi.OutputState }

func (SystemFtmpushMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFtmpush)(nil)).Elem()
}

func (o SystemFtmpushMapOutput) ToSystemFtmpushMapOutput() SystemFtmpushMapOutput {
	return o
}

func (o SystemFtmpushMapOutput) ToSystemFtmpushMapOutputWithContext(ctx context.Context) SystemFtmpushMapOutput {
	return o
}

func (o SystemFtmpushMapOutput) MapIndex(k pulumi.StringInput) SystemFtmpushOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFtmpush {
		return vs[0].(map[string]*SystemFtmpush)[vs[1].(string)]
	}).(SystemFtmpushOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFtmpushInput)(nil)).Elem(), &SystemFtmpush{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFtmpushArrayInput)(nil)).Elem(), SystemFtmpushArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFtmpushMapInput)(nil)).Elem(), SystemFtmpushMap{})
	pulumi.RegisterOutputType(SystemFtmpushOutput{})
	pulumi.RegisterOutputType(SystemFtmpushArrayOutput{})
	pulumi.RegisterOutputType(SystemFtmpushMapOutput{})
}