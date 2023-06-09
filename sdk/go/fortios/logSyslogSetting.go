// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure logging to remote Syslog logging servers.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `LogsyslogdSetting`, we recommend that you use the new resource.
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
//			_, err := fortios.NewLogSyslogSetting(ctx, "test2", &fortios.LogSyslogSettingArgs{
//				Facility: pulumi.String("local7"),
//				Format:   pulumi.String("csv"),
//				Mode:     pulumi.String("udp"),
//				Port:     pulumi.String("514"),
//				Server:   pulumi.String("2.2.2.2"),
//				SourceIp: pulumi.String("10.2.2.199"),
//				Status:   pulumi.String("enable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type LogSyslogSetting struct {
	pulumi.CustomResourceState

	// Remote syslog facility.
	Facility pulumi.StringOutput `pulumi:"facility"`
	// Log format.
	Format pulumi.StringOutput `pulumi:"format"`
	// Remote syslog logging over UDP/Reliable TCP.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Server listen port.
	Port pulumi.StringOutput `pulumi:"port"`
	// Address of remote syslog server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable remote syslog logging.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewLogSyslogSetting registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogSetting(ctx *pulumi.Context,
	name string, args *LogSyslogSettingArgs, opts ...pulumi.ResourceOption) (*LogSyslogSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogSetting
	err := ctx.RegisterResource("fortios:index/logSyslogSetting:LogSyslogSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogSetting gets an existing LogSyslogSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogSettingState, opts ...pulumi.ResourceOption) (*LogSyslogSetting, error) {
	var resource LogSyslogSetting
	err := ctx.ReadResource("fortios:index/logSyslogSetting:LogSyslogSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogSetting resources.
type logSyslogSettingState struct {
	// Remote syslog facility.
	Facility *string `pulumi:"facility"`
	// Log format.
	Format *string `pulumi:"format"`
	// Remote syslog logging over UDP/Reliable TCP.
	Mode *string `pulumi:"mode"`
	// Server listen port.
	Port *string `pulumi:"port"`
	// Address of remote syslog server.
	Server *string `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable remote syslog logging.
	Status *string `pulumi:"status"`
}

type LogSyslogSettingState struct {
	// Remote syslog facility.
	Facility pulumi.StringPtrInput
	// Log format.
	Format pulumi.StringPtrInput
	// Remote syslog logging over UDP/Reliable TCP.
	Mode pulumi.StringPtrInput
	// Server listen port.
	Port pulumi.StringPtrInput
	// Address of remote syslog server.
	Server pulumi.StringPtrInput
	// Source IP address of syslog.
	SourceIp pulumi.StringPtrInput
	// Enable/disable remote syslog logging.
	Status pulumi.StringPtrInput
}

func (LogSyslogSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogSettingState)(nil)).Elem()
}

type logSyslogSettingArgs struct {
	// Remote syslog facility.
	Facility *string `pulumi:"facility"`
	// Log format.
	Format *string `pulumi:"format"`
	// Remote syslog logging over UDP/Reliable TCP.
	Mode *string `pulumi:"mode"`
	// Server listen port.
	Port *string `pulumi:"port"`
	// Address of remote syslog server.
	Server *string `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable remote syslog logging.
	Status string `pulumi:"status"`
}

// The set of arguments for constructing a LogSyslogSetting resource.
type LogSyslogSettingArgs struct {
	// Remote syslog facility.
	Facility pulumi.StringPtrInput
	// Log format.
	Format pulumi.StringPtrInput
	// Remote syslog logging over UDP/Reliable TCP.
	Mode pulumi.StringPtrInput
	// Server listen port.
	Port pulumi.StringPtrInput
	// Address of remote syslog server.
	Server pulumi.StringPtrInput
	// Source IP address of syslog.
	SourceIp pulumi.StringPtrInput
	// Enable/disable remote syslog logging.
	Status pulumi.StringInput
}

func (LogSyslogSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogSettingArgs)(nil)).Elem()
}

type LogSyslogSettingInput interface {
	pulumi.Input

	ToLogSyslogSettingOutput() LogSyslogSettingOutput
	ToLogSyslogSettingOutputWithContext(ctx context.Context) LogSyslogSettingOutput
}

func (*LogSyslogSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogSetting)(nil)).Elem()
}

func (i *LogSyslogSetting) ToLogSyslogSettingOutput() LogSyslogSettingOutput {
	return i.ToLogSyslogSettingOutputWithContext(context.Background())
}

func (i *LogSyslogSetting) ToLogSyslogSettingOutputWithContext(ctx context.Context) LogSyslogSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogSettingOutput)
}

// LogSyslogSettingArrayInput is an input type that accepts LogSyslogSettingArray and LogSyslogSettingArrayOutput values.
// You can construct a concrete instance of `LogSyslogSettingArrayInput` via:
//
//	LogSyslogSettingArray{ LogSyslogSettingArgs{...} }
type LogSyslogSettingArrayInput interface {
	pulumi.Input

	ToLogSyslogSettingArrayOutput() LogSyslogSettingArrayOutput
	ToLogSyslogSettingArrayOutputWithContext(context.Context) LogSyslogSettingArrayOutput
}

type LogSyslogSettingArray []LogSyslogSettingInput

func (LogSyslogSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogSetting)(nil)).Elem()
}

func (i LogSyslogSettingArray) ToLogSyslogSettingArrayOutput() LogSyslogSettingArrayOutput {
	return i.ToLogSyslogSettingArrayOutputWithContext(context.Background())
}

func (i LogSyslogSettingArray) ToLogSyslogSettingArrayOutputWithContext(ctx context.Context) LogSyslogSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogSettingArrayOutput)
}

// LogSyslogSettingMapInput is an input type that accepts LogSyslogSettingMap and LogSyslogSettingMapOutput values.
// You can construct a concrete instance of `LogSyslogSettingMapInput` via:
//
//	LogSyslogSettingMap{ "key": LogSyslogSettingArgs{...} }
type LogSyslogSettingMapInput interface {
	pulumi.Input

	ToLogSyslogSettingMapOutput() LogSyslogSettingMapOutput
	ToLogSyslogSettingMapOutputWithContext(context.Context) LogSyslogSettingMapOutput
}

type LogSyslogSettingMap map[string]LogSyslogSettingInput

func (LogSyslogSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogSetting)(nil)).Elem()
}

func (i LogSyslogSettingMap) ToLogSyslogSettingMapOutput() LogSyslogSettingMapOutput {
	return i.ToLogSyslogSettingMapOutputWithContext(context.Background())
}

func (i LogSyslogSettingMap) ToLogSyslogSettingMapOutputWithContext(ctx context.Context) LogSyslogSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogSettingMapOutput)
}

type LogSyslogSettingOutput struct{ *pulumi.OutputState }

func (LogSyslogSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogSetting)(nil)).Elem()
}

func (o LogSyslogSettingOutput) ToLogSyslogSettingOutput() LogSyslogSettingOutput {
	return o
}

func (o LogSyslogSettingOutput) ToLogSyslogSettingOutputWithContext(ctx context.Context) LogSyslogSettingOutput {
	return o
}

// Remote syslog facility.
func (o LogSyslogSettingOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.Facility }).(pulumi.StringOutput)
}

// Log format.
func (o LogSyslogSettingOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Remote syslog logging over UDP/Reliable TCP.
func (o LogSyslogSettingOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Server listen port.
func (o LogSyslogSettingOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Address of remote syslog server.
func (o LogSyslogSettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Source IP address of syslog.
func (o LogSyslogSettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable remote syslog logging.
func (o LogSyslogSettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogSetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type LogSyslogSettingArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogSetting)(nil)).Elem()
}

func (o LogSyslogSettingArrayOutput) ToLogSyslogSettingArrayOutput() LogSyslogSettingArrayOutput {
	return o
}

func (o LogSyslogSettingArrayOutput) ToLogSyslogSettingArrayOutputWithContext(ctx context.Context) LogSyslogSettingArrayOutput {
	return o
}

func (o LogSyslogSettingArrayOutput) Index(i pulumi.IntInput) LogSyslogSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogSetting {
		return vs[0].([]*LogSyslogSetting)[vs[1].(int)]
	}).(LogSyslogSettingOutput)
}

type LogSyslogSettingMapOutput struct{ *pulumi.OutputState }

func (LogSyslogSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogSetting)(nil)).Elem()
}

func (o LogSyslogSettingMapOutput) ToLogSyslogSettingMapOutput() LogSyslogSettingMapOutput {
	return o
}

func (o LogSyslogSettingMapOutput) ToLogSyslogSettingMapOutputWithContext(ctx context.Context) LogSyslogSettingMapOutput {
	return o
}

func (o LogSyslogSettingMapOutput) MapIndex(k pulumi.StringInput) LogSyslogSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogSetting {
		return vs[0].(map[string]*LogSyslogSetting)[vs[1].(string)]
	}).(LogSyslogSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogSettingInput)(nil)).Elem(), &LogSyslogSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogSettingArrayInput)(nil)).Elem(), LogSyslogSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogSettingMapInput)(nil)).Elem(), LogSyslogSettingMap{})
	pulumi.RegisterOutputType(LogSyslogSettingOutput{})
	pulumi.RegisterOutputType(LogSyslogSettingArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogSettingMapOutput{})
}
