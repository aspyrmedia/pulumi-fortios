// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure quarantine options.
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
//			_, err := fortios.NewAntivirusQuarantine(ctx, "trname", &fortios.AntivirusQuarantineArgs{
//				Agelimit:        pulumi.Int(0),
//				Destination:     pulumi.String("disk"),
//				Lowspace:        pulumi.String("ovrw-old"),
//				Maxfilesize:     pulumi.Int(0),
//				QuarantineQuota: pulumi.Int(0),
//				StoreBlocked:    pulumi.String("imap smtp pop3 http ftp nntp imaps smtps pop3s ftps mapi cifs"),
//				StoreHeuristic:  pulumi.String("imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"),
//				StoreInfected:   pulumi.String("imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"),
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
// # Antivirus Quarantine can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/antivirusQuarantine:AntivirusQuarantine labelname AntivirusQuarantine
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/antivirusQuarantine:AntivirusQuarantine labelname AntivirusQuarantine
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type AntivirusQuarantine struct {
	pulumi.CustomResourceState

	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit pulumi.IntOutput `pulumi:"agelimit"`
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked pulumi.StringOutput `pulumi:"dropBlocked"`
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic pulumi.StringOutput `pulumi:"dropHeuristic"`
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected pulumi.StringOutput `pulumi:"dropInfected"`
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning pulumi.StringOutput `pulumi:"dropMachineLearning"`
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace pulumi.StringOutput `pulumi:"lowspace"`
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize pulumi.IntOutput `pulumi:"maxfilesize"`
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota pulumi.IntOutput `pulumi:"quarantineQuota"`
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked pulumi.StringOutput `pulumi:"storeBlocked"`
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic pulumi.StringOutput `pulumi:"storeHeuristic"`
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected pulumi.StringOutput `pulumi:"storeInfected"`
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning pulumi.StringOutput `pulumi:"storeMachineLearning"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAntivirusQuarantine registers a new resource with the given unique name, arguments, and options.
func NewAntivirusQuarantine(ctx *pulumi.Context,
	name string, args *AntivirusQuarantineArgs, opts ...pulumi.ResourceOption) (*AntivirusQuarantine, error) {
	if args == nil {
		args = &AntivirusQuarantineArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AntivirusQuarantine
	err := ctx.RegisterResource("fortios:index/antivirusQuarantine:AntivirusQuarantine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAntivirusQuarantine gets an existing AntivirusQuarantine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAntivirusQuarantine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AntivirusQuarantineState, opts ...pulumi.ResourceOption) (*AntivirusQuarantine, error) {
	var resource AntivirusQuarantine
	err := ctx.ReadResource("fortios:index/antivirusQuarantine:AntivirusQuarantine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AntivirusQuarantine resources.
type antivirusQuarantineState struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit *int `pulumi:"agelimit"`
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination *string `pulumi:"destination"`
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked *string `pulumi:"dropBlocked"`
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic *string `pulumi:"dropHeuristic"`
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected *string `pulumi:"dropInfected"`
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning *string `pulumi:"dropMachineLearning"`
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace *string `pulumi:"lowspace"`
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize *int `pulumi:"maxfilesize"`
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota *int `pulumi:"quarantineQuota"`
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked *string `pulumi:"storeBlocked"`
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic *string `pulumi:"storeHeuristic"`
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected *string `pulumi:"storeInfected"`
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning *string `pulumi:"storeMachineLearning"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AntivirusQuarantineState struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit pulumi.IntPtrInput
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination pulumi.StringPtrInput
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked pulumi.StringPtrInput
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic pulumi.StringPtrInput
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected pulumi.StringPtrInput
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning pulumi.StringPtrInput
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace pulumi.StringPtrInput
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize pulumi.IntPtrInput
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota pulumi.IntPtrInput
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked pulumi.StringPtrInput
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic pulumi.StringPtrInput
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected pulumi.StringPtrInput
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AntivirusQuarantineState) ElementType() reflect.Type {
	return reflect.TypeOf((*antivirusQuarantineState)(nil)).Elem()
}

type antivirusQuarantineArgs struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit *int `pulumi:"agelimit"`
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination *string `pulumi:"destination"`
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked *string `pulumi:"dropBlocked"`
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic *string `pulumi:"dropHeuristic"`
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected *string `pulumi:"dropInfected"`
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning *string `pulumi:"dropMachineLearning"`
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace *string `pulumi:"lowspace"`
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize *int `pulumi:"maxfilesize"`
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota *int `pulumi:"quarantineQuota"`
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked *string `pulumi:"storeBlocked"`
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic *string `pulumi:"storeHeuristic"`
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected *string `pulumi:"storeInfected"`
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning *string `pulumi:"storeMachineLearning"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a AntivirusQuarantine resource.
type AntivirusQuarantineArgs struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit pulumi.IntPtrInput
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination pulumi.StringPtrInput
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked pulumi.StringPtrInput
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic pulumi.StringPtrInput
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected pulumi.StringPtrInput
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning pulumi.StringPtrInput
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace pulumi.StringPtrInput
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize pulumi.IntPtrInput
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota pulumi.IntPtrInput
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked pulumi.StringPtrInput
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic pulumi.StringPtrInput
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected pulumi.StringPtrInput
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AntivirusQuarantineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*antivirusQuarantineArgs)(nil)).Elem()
}

type AntivirusQuarantineInput interface {
	pulumi.Input

	ToAntivirusQuarantineOutput() AntivirusQuarantineOutput
	ToAntivirusQuarantineOutputWithContext(ctx context.Context) AntivirusQuarantineOutput
}

func (*AntivirusQuarantine) ElementType() reflect.Type {
	return reflect.TypeOf((**AntivirusQuarantine)(nil)).Elem()
}

func (i *AntivirusQuarantine) ToAntivirusQuarantineOutput() AntivirusQuarantineOutput {
	return i.ToAntivirusQuarantineOutputWithContext(context.Background())
}

func (i *AntivirusQuarantine) ToAntivirusQuarantineOutputWithContext(ctx context.Context) AntivirusQuarantineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntivirusQuarantineOutput)
}

// AntivirusQuarantineArrayInput is an input type that accepts AntivirusQuarantineArray and AntivirusQuarantineArrayOutput values.
// You can construct a concrete instance of `AntivirusQuarantineArrayInput` via:
//
//	AntivirusQuarantineArray{ AntivirusQuarantineArgs{...} }
type AntivirusQuarantineArrayInput interface {
	pulumi.Input

	ToAntivirusQuarantineArrayOutput() AntivirusQuarantineArrayOutput
	ToAntivirusQuarantineArrayOutputWithContext(context.Context) AntivirusQuarantineArrayOutput
}

type AntivirusQuarantineArray []AntivirusQuarantineInput

func (AntivirusQuarantineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AntivirusQuarantine)(nil)).Elem()
}

func (i AntivirusQuarantineArray) ToAntivirusQuarantineArrayOutput() AntivirusQuarantineArrayOutput {
	return i.ToAntivirusQuarantineArrayOutputWithContext(context.Background())
}

func (i AntivirusQuarantineArray) ToAntivirusQuarantineArrayOutputWithContext(ctx context.Context) AntivirusQuarantineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntivirusQuarantineArrayOutput)
}

// AntivirusQuarantineMapInput is an input type that accepts AntivirusQuarantineMap and AntivirusQuarantineMapOutput values.
// You can construct a concrete instance of `AntivirusQuarantineMapInput` via:
//
//	AntivirusQuarantineMap{ "key": AntivirusQuarantineArgs{...} }
type AntivirusQuarantineMapInput interface {
	pulumi.Input

	ToAntivirusQuarantineMapOutput() AntivirusQuarantineMapOutput
	ToAntivirusQuarantineMapOutputWithContext(context.Context) AntivirusQuarantineMapOutput
}

type AntivirusQuarantineMap map[string]AntivirusQuarantineInput

func (AntivirusQuarantineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AntivirusQuarantine)(nil)).Elem()
}

func (i AntivirusQuarantineMap) ToAntivirusQuarantineMapOutput() AntivirusQuarantineMapOutput {
	return i.ToAntivirusQuarantineMapOutputWithContext(context.Background())
}

func (i AntivirusQuarantineMap) ToAntivirusQuarantineMapOutputWithContext(ctx context.Context) AntivirusQuarantineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntivirusQuarantineMapOutput)
}

type AntivirusQuarantineOutput struct{ *pulumi.OutputState }

func (AntivirusQuarantineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AntivirusQuarantine)(nil)).Elem()
}

func (o AntivirusQuarantineOutput) ToAntivirusQuarantineOutput() AntivirusQuarantineOutput {
	return o
}

func (o AntivirusQuarantineOutput) ToAntivirusQuarantineOutputWithContext(ctx context.Context) AntivirusQuarantineOutput {
	return o
}

// Age limit for quarantined files (0 - 479 hours, 0 means forever).
func (o AntivirusQuarantineOutput) Agelimit() pulumi.IntOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.IntOutput { return v.Agelimit }).(pulumi.IntOutput)
}

// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
func (o AntivirusQuarantineOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
func (o AntivirusQuarantineOutput) DropBlocked() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.DropBlocked }).(pulumi.StringOutput)
}

// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
func (o AntivirusQuarantineOutput) DropHeuristic() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.DropHeuristic }).(pulumi.StringOutput)
}

// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
func (o AntivirusQuarantineOutput) DropInfected() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.DropInfected }).(pulumi.StringOutput)
}

// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
func (o AntivirusQuarantineOutput) DropMachineLearning() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.DropMachineLearning }).(pulumi.StringOutput)
}

// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
func (o AntivirusQuarantineOutput) Lowspace() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.Lowspace }).(pulumi.StringOutput)
}

// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
func (o AntivirusQuarantineOutput) Maxfilesize() pulumi.IntOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.IntOutput { return v.Maxfilesize }).(pulumi.IntOutput)
}

// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
func (o AntivirusQuarantineOutput) QuarantineQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.IntOutput { return v.QuarantineQuota }).(pulumi.IntOutput)
}

// Quarantine blocked files found in sessions using the selected protocols.
func (o AntivirusQuarantineOutput) StoreBlocked() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.StoreBlocked }).(pulumi.StringOutput)
}

// Quarantine files detected by heuristics found in sessions using the selected protocols.
func (o AntivirusQuarantineOutput) StoreHeuristic() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.StoreHeuristic }).(pulumi.StringOutput)
}

// Quarantine infected files found in sessions using the selected protocols.
func (o AntivirusQuarantineOutput) StoreInfected() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.StoreInfected }).(pulumi.StringOutput)
}

// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
func (o AntivirusQuarantineOutput) StoreMachineLearning() pulumi.StringOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringOutput { return v.StoreMachineLearning }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AntivirusQuarantineOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AntivirusQuarantine) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AntivirusQuarantineArrayOutput struct{ *pulumi.OutputState }

func (AntivirusQuarantineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AntivirusQuarantine)(nil)).Elem()
}

func (o AntivirusQuarantineArrayOutput) ToAntivirusQuarantineArrayOutput() AntivirusQuarantineArrayOutput {
	return o
}

func (o AntivirusQuarantineArrayOutput) ToAntivirusQuarantineArrayOutputWithContext(ctx context.Context) AntivirusQuarantineArrayOutput {
	return o
}

func (o AntivirusQuarantineArrayOutput) Index(i pulumi.IntInput) AntivirusQuarantineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AntivirusQuarantine {
		return vs[0].([]*AntivirusQuarantine)[vs[1].(int)]
	}).(AntivirusQuarantineOutput)
}

type AntivirusQuarantineMapOutput struct{ *pulumi.OutputState }

func (AntivirusQuarantineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AntivirusQuarantine)(nil)).Elem()
}

func (o AntivirusQuarantineMapOutput) ToAntivirusQuarantineMapOutput() AntivirusQuarantineMapOutput {
	return o
}

func (o AntivirusQuarantineMapOutput) ToAntivirusQuarantineMapOutputWithContext(ctx context.Context) AntivirusQuarantineMapOutput {
	return o
}

func (o AntivirusQuarantineMapOutput) MapIndex(k pulumi.StringInput) AntivirusQuarantineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AntivirusQuarantine {
		return vs[0].(map[string]*AntivirusQuarantine)[vs[1].(string)]
	}).(AntivirusQuarantineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AntivirusQuarantineInput)(nil)).Elem(), &AntivirusQuarantine{})
	pulumi.RegisterInputType(reflect.TypeOf((*AntivirusQuarantineArrayInput)(nil)).Elem(), AntivirusQuarantineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AntivirusQuarantineMapInput)(nil)).Elem(), AntivirusQuarantineMap{})
	pulumi.RegisterOutputType(AntivirusQuarantineOutput{})
	pulumi.RegisterOutputType(AntivirusQuarantineArrayOutput{})
	pulumi.RegisterOutputType(AntivirusQuarantineMapOutput{})
}
