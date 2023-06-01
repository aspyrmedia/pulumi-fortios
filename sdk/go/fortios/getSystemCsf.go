// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system csf
func LookupSystemCsf(ctx *pulumi.Context, args *LookupSystemCsfArgs, opts ...pulumi.InvokeOption) (*LookupSystemCsfResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemCsfResult
	err := ctx.Invoke("fortios:index/getSystemCsf:getSystemCsf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemCsf.
type LookupSystemCsfArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemCsf.
type LookupSystemCsfResult struct {
	// Accept connections with unknown certificates and ask admin for approval.
	AcceptAuthByCert string `pulumi:"acceptAuthByCert"`
	// Authorization request type.
	AuthorizationRequestType string `pulumi:"authorizationRequestType"`
	// Certificate.
	Certificate string `pulumi:"certificate"`
	// Configuration sync mode.
	ConfigurationSync string `pulumi:"configurationSync"`
	// Enable/disable downstream device access to this device's configuration and data.
	DownstreamAccess string `pulumi:"downstreamAccess"`
	// Default access profile for requests from downstream devices.
	DownstreamAccprofile string `pulumi:"downstreamAccprofile"`
	// Fabric connector configuration. The structure of `fabricConnector` block is documented below.
	FabricConnectors []GetSystemCsfFabricConnector `pulumi:"fabricConnectors"`
	// Fabric device configuration. The structure of `fabricDevice` block is documented below.
	FabricDevices []GetSystemCsfFabricDevice `pulumi:"fabricDevices"`
	// Fabric CMDB Object Unification
	FabricObjectUnification string `pulumi:"fabricObjectUnification"`
	// Number of worker processes for Security Fabric daemon.
	FabricWorkers int `pulumi:"fabricWorkers"`
	// Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
	FixedKey string `pulumi:"fixedKey"`
	// Fabric FortiCloud account unification.
	ForticloudAccountEnforcement string `pulumi:"forticloudAccountEnforcement"`
	// Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
	GroupName string `pulumi:"groupName"`
	// Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
	GroupPassword string `pulumi:"groupPassword"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable broadcast of discovery messages for log unification.
	LogUnification string `pulumi:"logUnification"`
	// Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
	ManagementIp string `pulumi:"managementIp"`
	// Overriding port for management connection (Overrides admin port).
	ManagementPort int `pulumi:"managementPort"`
	// SAML setting configuration synchronization.
	SamlConfigurationSync string `pulumi:"samlConfigurationSync"`
	// Enable/disable Security Fabric.
	Status string `pulumi:"status"`
	// Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
	TrustedLists []GetSystemCsfTrustedList `pulumi:"trustedLists"`
	// IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
	Upstream string `pulumi:"upstream"`
	// IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
	UpstreamIp string `pulumi:"upstreamIp"`
	// The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
	UpstreamPort int     `pulumi:"upstreamPort"`
	Vdomparam    *string `pulumi:"vdomparam"`
}

func LookupSystemCsfOutput(ctx *pulumi.Context, args LookupSystemCsfOutputArgs, opts ...pulumi.InvokeOption) LookupSystemCsfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemCsfResult, error) {
			args := v.(LookupSystemCsfArgs)
			r, err := LookupSystemCsf(ctx, &args, opts...)
			var s LookupSystemCsfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemCsfResultOutput)
}

// A collection of arguments for invoking getSystemCsf.
type LookupSystemCsfOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemCsfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemCsfArgs)(nil)).Elem()
}

// A collection of values returned by getSystemCsf.
type LookupSystemCsfResultOutput struct{ *pulumi.OutputState }

func (LookupSystemCsfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemCsfResult)(nil)).Elem()
}

func (o LookupSystemCsfResultOutput) ToLookupSystemCsfResultOutput() LookupSystemCsfResultOutput {
	return o
}

func (o LookupSystemCsfResultOutput) ToLookupSystemCsfResultOutputWithContext(ctx context.Context) LookupSystemCsfResultOutput {
	return o
}

// Accept connections with unknown certificates and ask admin for approval.
func (o LookupSystemCsfResultOutput) AcceptAuthByCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.AcceptAuthByCert }).(pulumi.StringOutput)
}

// Authorization request type.
func (o LookupSystemCsfResultOutput) AuthorizationRequestType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.AuthorizationRequestType }).(pulumi.StringOutput)
}

// Certificate.
func (o LookupSystemCsfResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.Certificate }).(pulumi.StringOutput)
}

// Configuration sync mode.
func (o LookupSystemCsfResultOutput) ConfigurationSync() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.ConfigurationSync }).(pulumi.StringOutput)
}

// Enable/disable downstream device access to this device's configuration and data.
func (o LookupSystemCsfResultOutput) DownstreamAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.DownstreamAccess }).(pulumi.StringOutput)
}

// Default access profile for requests from downstream devices.
func (o LookupSystemCsfResultOutput) DownstreamAccprofile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.DownstreamAccprofile }).(pulumi.StringOutput)
}

// Fabric connector configuration. The structure of `fabricConnector` block is documented below.
func (o LookupSystemCsfResultOutput) FabricConnectors() GetSystemCsfFabricConnectorArrayOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) []GetSystemCsfFabricConnector { return v.FabricConnectors }).(GetSystemCsfFabricConnectorArrayOutput)
}

// Fabric device configuration. The structure of `fabricDevice` block is documented below.
func (o LookupSystemCsfResultOutput) FabricDevices() GetSystemCsfFabricDeviceArrayOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) []GetSystemCsfFabricDevice { return v.FabricDevices }).(GetSystemCsfFabricDeviceArrayOutput)
}

// Fabric CMDB Object Unification
func (o LookupSystemCsfResultOutput) FabricObjectUnification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.FabricObjectUnification }).(pulumi.StringOutput)
}

// Number of worker processes for Security Fabric daemon.
func (o LookupSystemCsfResultOutput) FabricWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) int { return v.FabricWorkers }).(pulumi.IntOutput)
}

// Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
func (o LookupSystemCsfResultOutput) FixedKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.FixedKey }).(pulumi.StringOutput)
}

// Fabric FortiCloud account unification.
func (o LookupSystemCsfResultOutput) ForticloudAccountEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.ForticloudAccountEnforcement }).(pulumi.StringOutput)
}

// Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
func (o LookupSystemCsfResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
func (o LookupSystemCsfResultOutput) GroupPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.GroupPassword }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemCsfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable broadcast of discovery messages for log unification.
func (o LookupSystemCsfResultOutput) LogUnification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.LogUnification }).(pulumi.StringOutput)
}

// Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
func (o LookupSystemCsfResultOutput) ManagementIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.ManagementIp }).(pulumi.StringOutput)
}

// Overriding port for management connection (Overrides admin port).
func (o LookupSystemCsfResultOutput) ManagementPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) int { return v.ManagementPort }).(pulumi.IntOutput)
}

// SAML setting configuration synchronization.
func (o LookupSystemCsfResultOutput) SamlConfigurationSync() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.SamlConfigurationSync }).(pulumi.StringOutput)
}

// Enable/disable Security Fabric.
func (o LookupSystemCsfResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.Status }).(pulumi.StringOutput)
}

// Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
func (o LookupSystemCsfResultOutput) TrustedLists() GetSystemCsfTrustedListArrayOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) []GetSystemCsfTrustedList { return v.TrustedLists }).(GetSystemCsfTrustedListArrayOutput)
}

// IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
func (o LookupSystemCsfResultOutput) Upstream() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.Upstream }).(pulumi.StringOutput)
}

// IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
func (o LookupSystemCsfResultOutput) UpstreamIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) string { return v.UpstreamIp }).(pulumi.StringOutput)
}

// The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
func (o LookupSystemCsfResultOutput) UpstreamPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) int { return v.UpstreamPort }).(pulumi.IntOutput)
}

func (o LookupSystemCsfResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemCsfResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemCsfResultOutput{})
}