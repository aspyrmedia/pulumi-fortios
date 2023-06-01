// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DLP profiles. Applies to FortiOS Version `>= 7.2.0`.
//
// ## Import
//
// # Dlp Profile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/dlpProfile:DlpProfile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/dlpProfile:DlpProfile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type DlpProfile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog pulumi.StringOutput `pulumi:"dlpLog"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`
	// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	FullArchiveProto pulumi.StringOutput `pulumi:"fullArchiveProto"`
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog pulumi.StringOutput `pulumi:"nacQuarLog"`
	// Name of the DLP profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// Replacement message group used by this DLP profile.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Set up DLP rules for this profile. The structure of `rule` block is documented below.
	Rules DlpProfileRuleArrayOutput `pulumi:"rules"`
	// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	SummaryProto pulumi.StringOutput `pulumi:"summaryProto"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpProfile registers a new resource with the given unique name, arguments, and options.
func NewDlpProfile(ctx *pulumi.Context,
	name string, args *DlpProfileArgs, opts ...pulumi.ResourceOption) (*DlpProfile, error) {
	if args == nil {
		args = &DlpProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource DlpProfile
	err := ctx.RegisterResource("fortios:index/dlpProfile:DlpProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpProfile gets an existing DlpProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpProfileState, opts ...pulumi.ResourceOption) (*DlpProfile, error) {
	var resource DlpProfile
	err := ctx.ReadResource("fortios:index/dlpProfile:DlpProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpProfile resources.
type dlpProfileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog *string `pulumi:"dlpLog"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	FullArchiveProto *string `pulumi:"fullArchiveProto"`
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog *string `pulumi:"nacQuarLog"`
	// Name of the DLP profile.
	Name *string `pulumi:"name"`
	// Replacement message group used by this DLP profile.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Set up DLP rules for this profile. The structure of `rule` block is documented below.
	Rules []DlpProfileRule `pulumi:"rules"`
	// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	SummaryProto *string `pulumi:"summaryProto"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	FullArchiveProto pulumi.StringPtrInput
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog pulumi.StringPtrInput
	// Name of the DLP profile.
	Name pulumi.StringPtrInput
	// Replacement message group used by this DLP profile.
	ReplacemsgGroup pulumi.StringPtrInput
	// Set up DLP rules for this profile. The structure of `rule` block is documented below.
	Rules DlpProfileRuleArrayInput
	// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	SummaryProto pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpProfileState)(nil)).Elem()
}

type dlpProfileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog *string `pulumi:"dlpLog"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	FullArchiveProto *string `pulumi:"fullArchiveProto"`
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog *string `pulumi:"nacQuarLog"`
	// Name of the DLP profile.
	Name *string `pulumi:"name"`
	// Replacement message group used by this DLP profile.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Set up DLP rules for this profile. The structure of `rule` block is documented below.
	Rules []DlpProfileRule `pulumi:"rules"`
	// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	SummaryProto *string `pulumi:"summaryProto"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpProfile resource.
type DlpProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	FullArchiveProto pulumi.StringPtrInput
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog pulumi.StringPtrInput
	// Name of the DLP profile.
	Name pulumi.StringPtrInput
	// Replacement message group used by this DLP profile.
	ReplacemsgGroup pulumi.StringPtrInput
	// Set up DLP rules for this profile. The structure of `rule` block is documented below.
	Rules DlpProfileRuleArrayInput
	// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
	SummaryProto pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpProfileArgs)(nil)).Elem()
}

type DlpProfileInput interface {
	pulumi.Input

	ToDlpProfileOutput() DlpProfileOutput
	ToDlpProfileOutputWithContext(ctx context.Context) DlpProfileOutput
}

func (*DlpProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpProfile)(nil)).Elem()
}

func (i *DlpProfile) ToDlpProfileOutput() DlpProfileOutput {
	return i.ToDlpProfileOutputWithContext(context.Background())
}

func (i *DlpProfile) ToDlpProfileOutputWithContext(ctx context.Context) DlpProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpProfileOutput)
}

// DlpProfileArrayInput is an input type that accepts DlpProfileArray and DlpProfileArrayOutput values.
// You can construct a concrete instance of `DlpProfileArrayInput` via:
//
//	DlpProfileArray{ DlpProfileArgs{...} }
type DlpProfileArrayInput interface {
	pulumi.Input

	ToDlpProfileArrayOutput() DlpProfileArrayOutput
	ToDlpProfileArrayOutputWithContext(context.Context) DlpProfileArrayOutput
}

type DlpProfileArray []DlpProfileInput

func (DlpProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpProfile)(nil)).Elem()
}

func (i DlpProfileArray) ToDlpProfileArrayOutput() DlpProfileArrayOutput {
	return i.ToDlpProfileArrayOutputWithContext(context.Background())
}

func (i DlpProfileArray) ToDlpProfileArrayOutputWithContext(ctx context.Context) DlpProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpProfileArrayOutput)
}

// DlpProfileMapInput is an input type that accepts DlpProfileMap and DlpProfileMapOutput values.
// You can construct a concrete instance of `DlpProfileMapInput` via:
//
//	DlpProfileMap{ "key": DlpProfileArgs{...} }
type DlpProfileMapInput interface {
	pulumi.Input

	ToDlpProfileMapOutput() DlpProfileMapOutput
	ToDlpProfileMapOutputWithContext(context.Context) DlpProfileMapOutput
}

type DlpProfileMap map[string]DlpProfileInput

func (DlpProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpProfile)(nil)).Elem()
}

func (i DlpProfileMap) ToDlpProfileMapOutput() DlpProfileMapOutput {
	return i.ToDlpProfileMapOutputWithContext(context.Background())
}

func (i DlpProfileMap) ToDlpProfileMapOutputWithContext(ctx context.Context) DlpProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpProfileMapOutput)
}

type DlpProfileOutput struct{ *pulumi.OutputState }

func (DlpProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpProfile)(nil)).Elem()
}

func (o DlpProfileOutput) ToDlpProfileOutput() DlpProfileOutput {
	return o
}

func (o DlpProfileOutput) ToDlpProfileOutputWithContext(ctx context.Context) DlpProfileOutput {
	return o
}

// Comment.
func (o DlpProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Enable/disable DLP logging. Valid values: `enable`, `disable`.
func (o DlpProfileOutput) DlpLog() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.DlpLog }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DlpProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
func (o DlpProfileOutput) ExtendedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.ExtendedLog }).(pulumi.StringOutput)
}

// Flow/proxy feature set. Valid values: `flow`, `proxy`.
func (o DlpProfileOutput) FeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.FeatureSet }).(pulumi.StringOutput)
}

// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
func (o DlpProfileOutput) FullArchiveProto() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.FullArchiveProto }).(pulumi.StringOutput)
}

// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
func (o DlpProfileOutput) NacQuarLog() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.NacQuarLog }).(pulumi.StringOutput)
}

// Name of the DLP profile.
func (o DlpProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Replacement message group used by this DLP profile.
func (o DlpProfileOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// Set up DLP rules for this profile. The structure of `rule` block is documented below.
func (o DlpProfileOutput) Rules() DlpProfileRuleArrayOutput {
	return o.ApplyT(func(v *DlpProfile) DlpProfileRuleArrayOutput { return v.Rules }).(DlpProfileRuleArrayOutput)
}

// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
func (o DlpProfileOutput) SummaryProto() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringOutput { return v.SummaryProto }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DlpProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DlpProfileArrayOutput struct{ *pulumi.OutputState }

func (DlpProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpProfile)(nil)).Elem()
}

func (o DlpProfileArrayOutput) ToDlpProfileArrayOutput() DlpProfileArrayOutput {
	return o
}

func (o DlpProfileArrayOutput) ToDlpProfileArrayOutputWithContext(ctx context.Context) DlpProfileArrayOutput {
	return o
}

func (o DlpProfileArrayOutput) Index(i pulumi.IntInput) DlpProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpProfile {
		return vs[0].([]*DlpProfile)[vs[1].(int)]
	}).(DlpProfileOutput)
}

type DlpProfileMapOutput struct{ *pulumi.OutputState }

func (DlpProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpProfile)(nil)).Elem()
}

func (o DlpProfileMapOutput) ToDlpProfileMapOutput() DlpProfileMapOutput {
	return o
}

func (o DlpProfileMapOutput) ToDlpProfileMapOutputWithContext(ctx context.Context) DlpProfileMapOutput {
	return o
}

func (o DlpProfileMapOutput) MapIndex(k pulumi.StringInput) DlpProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpProfile {
		return vs[0].(map[string]*DlpProfile)[vs[1].(string)]
	}).(DlpProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpProfileInput)(nil)).Elem(), &DlpProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpProfileArrayInput)(nil)).Elem(), DlpProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpProfileMapInput)(nil)).Elem(), DlpProfileMap{})
	pulumi.RegisterOutputType(DlpProfileOutput{})
	pulumi.RegisterOutputType(DlpProfileArrayOutput{})
	pulumi.RegisterOutputType(DlpProfileMapOutput{})
}