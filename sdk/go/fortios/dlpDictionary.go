// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure dictionaries used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.
//
// ## Import
//
// # Dlp Dictionary can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/dlpDictionary:DlpDictionary labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/dlpDictionary:DlpDictionary labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type DlpDictionary struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// DLP dictionary entries. The structure of `entries` block is documented below.
	Entries DlpDictionaryEntryArrayOutput `pulumi:"entries"`
	// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
	MatchType pulumi.StringOutput `pulumi:"matchType"`
	// Name of table containing the dictionary.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpDictionary registers a new resource with the given unique name, arguments, and options.
func NewDlpDictionary(ctx *pulumi.Context,
	name string, args *DlpDictionaryArgs, opts ...pulumi.ResourceOption) (*DlpDictionary, error) {
	if args == nil {
		args = &DlpDictionaryArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource DlpDictionary
	err := ctx.RegisterResource("fortios:index/dlpDictionary:DlpDictionary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpDictionary gets an existing DlpDictionary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpDictionary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpDictionaryState, opts ...pulumi.ResourceOption) (*DlpDictionary, error) {
	var resource DlpDictionary
	err := ctx.ReadResource("fortios:index/dlpDictionary:DlpDictionary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpDictionary resources.
type dlpDictionaryState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// DLP dictionary entries. The structure of `entries` block is documented below.
	Entries []DlpDictionaryEntry `pulumi:"entries"`
	// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
	MatchType *string `pulumi:"matchType"`
	// Name of table containing the dictionary.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpDictionaryState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// DLP dictionary entries. The structure of `entries` block is documented below.
	Entries DlpDictionaryEntryArrayInput
	// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
	MatchType pulumi.StringPtrInput
	// Name of table containing the dictionary.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpDictionaryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpDictionaryState)(nil)).Elem()
}

type dlpDictionaryArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// DLP dictionary entries. The structure of `entries` block is documented below.
	Entries []DlpDictionaryEntry `pulumi:"entries"`
	// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
	MatchType *string `pulumi:"matchType"`
	// Name of table containing the dictionary.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpDictionary resource.
type DlpDictionaryArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// DLP dictionary entries. The structure of `entries` block is documented below.
	Entries DlpDictionaryEntryArrayInput
	// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
	MatchType pulumi.StringPtrInput
	// Name of table containing the dictionary.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpDictionaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpDictionaryArgs)(nil)).Elem()
}

type DlpDictionaryInput interface {
	pulumi.Input

	ToDlpDictionaryOutput() DlpDictionaryOutput
	ToDlpDictionaryOutputWithContext(ctx context.Context) DlpDictionaryOutput
}

func (*DlpDictionary) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpDictionary)(nil)).Elem()
}

func (i *DlpDictionary) ToDlpDictionaryOutput() DlpDictionaryOutput {
	return i.ToDlpDictionaryOutputWithContext(context.Background())
}

func (i *DlpDictionary) ToDlpDictionaryOutputWithContext(ctx context.Context) DlpDictionaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpDictionaryOutput)
}

// DlpDictionaryArrayInput is an input type that accepts DlpDictionaryArray and DlpDictionaryArrayOutput values.
// You can construct a concrete instance of `DlpDictionaryArrayInput` via:
//
//	DlpDictionaryArray{ DlpDictionaryArgs{...} }
type DlpDictionaryArrayInput interface {
	pulumi.Input

	ToDlpDictionaryArrayOutput() DlpDictionaryArrayOutput
	ToDlpDictionaryArrayOutputWithContext(context.Context) DlpDictionaryArrayOutput
}

type DlpDictionaryArray []DlpDictionaryInput

func (DlpDictionaryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpDictionary)(nil)).Elem()
}

func (i DlpDictionaryArray) ToDlpDictionaryArrayOutput() DlpDictionaryArrayOutput {
	return i.ToDlpDictionaryArrayOutputWithContext(context.Background())
}

func (i DlpDictionaryArray) ToDlpDictionaryArrayOutputWithContext(ctx context.Context) DlpDictionaryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpDictionaryArrayOutput)
}

// DlpDictionaryMapInput is an input type that accepts DlpDictionaryMap and DlpDictionaryMapOutput values.
// You can construct a concrete instance of `DlpDictionaryMapInput` via:
//
//	DlpDictionaryMap{ "key": DlpDictionaryArgs{...} }
type DlpDictionaryMapInput interface {
	pulumi.Input

	ToDlpDictionaryMapOutput() DlpDictionaryMapOutput
	ToDlpDictionaryMapOutputWithContext(context.Context) DlpDictionaryMapOutput
}

type DlpDictionaryMap map[string]DlpDictionaryInput

func (DlpDictionaryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpDictionary)(nil)).Elem()
}

func (i DlpDictionaryMap) ToDlpDictionaryMapOutput() DlpDictionaryMapOutput {
	return i.ToDlpDictionaryMapOutputWithContext(context.Background())
}

func (i DlpDictionaryMap) ToDlpDictionaryMapOutputWithContext(ctx context.Context) DlpDictionaryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpDictionaryMapOutput)
}

type DlpDictionaryOutput struct{ *pulumi.OutputState }

func (DlpDictionaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpDictionary)(nil)).Elem()
}

func (o DlpDictionaryOutput) ToDlpDictionaryOutput() DlpDictionaryOutput {
	return o
}

func (o DlpDictionaryOutput) ToDlpDictionaryOutputWithContext(ctx context.Context) DlpDictionaryOutput {
	return o
}

// Optional comments.
func (o DlpDictionaryOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpDictionary) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DlpDictionaryOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpDictionary) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// DLP dictionary entries. The structure of `entries` block is documented below.
func (o DlpDictionaryOutput) Entries() DlpDictionaryEntryArrayOutput {
	return o.ApplyT(func(v *DlpDictionary) DlpDictionaryEntryArrayOutput { return v.Entries }).(DlpDictionaryEntryArrayOutput)
}

// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
func (o DlpDictionaryOutput) MatchType() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpDictionary) pulumi.StringOutput { return v.MatchType }).(pulumi.StringOutput)
}

// Name of table containing the dictionary.
func (o DlpDictionaryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpDictionary) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o DlpDictionaryOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpDictionary) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DlpDictionaryOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpDictionary) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DlpDictionaryArrayOutput struct{ *pulumi.OutputState }

func (DlpDictionaryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpDictionary)(nil)).Elem()
}

func (o DlpDictionaryArrayOutput) ToDlpDictionaryArrayOutput() DlpDictionaryArrayOutput {
	return o
}

func (o DlpDictionaryArrayOutput) ToDlpDictionaryArrayOutputWithContext(ctx context.Context) DlpDictionaryArrayOutput {
	return o
}

func (o DlpDictionaryArrayOutput) Index(i pulumi.IntInput) DlpDictionaryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpDictionary {
		return vs[0].([]*DlpDictionary)[vs[1].(int)]
	}).(DlpDictionaryOutput)
}

type DlpDictionaryMapOutput struct{ *pulumi.OutputState }

func (DlpDictionaryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpDictionary)(nil)).Elem()
}

func (o DlpDictionaryMapOutput) ToDlpDictionaryMapOutput() DlpDictionaryMapOutput {
	return o
}

func (o DlpDictionaryMapOutput) ToDlpDictionaryMapOutputWithContext(ctx context.Context) DlpDictionaryMapOutput {
	return o
}

func (o DlpDictionaryMapOutput) MapIndex(k pulumi.StringInput) DlpDictionaryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpDictionary {
		return vs[0].(map[string]*DlpDictionary)[vs[1].(string)]
	}).(DlpDictionaryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpDictionaryInput)(nil)).Elem(), &DlpDictionary{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpDictionaryArrayInput)(nil)).Elem(), DlpDictionaryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpDictionaryMapInput)(nil)).Elem(), DlpDictionaryMap{})
	pulumi.RegisterOutputType(DlpDictionaryOutput{})
	pulumi.RegisterOutputType(DlpDictionaryArrayOutput{})
	pulumi.RegisterOutputType(DlpDictionaryMapOutput{})
}
