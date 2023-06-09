// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system sdnconnector
func LookupSystemSdnconnector(ctx *pulumi.Context, args *LookupSystemSdnconnectorArgs, opts ...pulumi.InvokeOption) (*LookupSystemSdnconnectorResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemSdnconnectorResult
	err := ctx.Invoke("fortios:index/getSystemSdnconnector:getSystemSdnconnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemSdnconnector.
type LookupSystemSdnconnectorArgs struct {
	// Specify the name of the desired system sdnconnector.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemSdnconnector.
type LookupSystemSdnconnectorResult struct {
	// AWS access key ID.
	AccessKey string `pulumi:"accessKey"`
	// IBM cloud API key or service ID API key.
	ApiKey string `pulumi:"apiKey"`
	// Azure server region.
	AzureRegion string `pulumi:"azureRegion"`
	// Azure client ID (application ID).
	ClientId string `pulumi:"clientId"`
	// Azure client secret (application key).
	ClientSecret string `pulumi:"clientSecret"`
	// Compartment ID.
	CompartmentId string `pulumi:"compartmentId"`
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration int `pulumi:"computeGeneration"`
	// Domain name.
	Domain string `pulumi:"domain"`
	// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
	ExternalAccountLists []GetSystemSdnconnectorExternalAccountList `pulumi:"externalAccountLists"`
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps []GetSystemSdnconnectorExternalIp `pulumi:"externalIps"`
	// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
	ForwardingRules []GetSystemSdnconnectorForwardingRule `pulumi:"forwardingRules"`
	// GCP project name.
	GcpProject string `pulumi:"gcpProject"`
	// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
	GcpProjectLists []GetSystemSdnconnectorGcpProjectList `pulumi:"gcpProjectLists"`
	// Group name of computers.
	GroupName string `pulumi:"groupName"`
	// Enable/disable use for FortiGate HA service.
	HaStatus string `pulumi:"haStatus"`
	// IBM cloud region name.
	IbmRegion string `pulumi:"ibmRegion"`
	// IBM cloud compute generation 1 region name.
	IbmRegionGen1 string `pulumi:"ibmRegionGen1"`
	// IBM cloud compute generation 2 region name.
	IbmRegionGen2 string `pulumi:"ibmRegionGen2"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Private key password.
	KeyPasswd string `pulumi:"keyPasswd"`
	// Azure Stack login endpoint.
	LoginEndpoint string `pulumi:"loginEndpoint"`
	// GCP zone name.
	Name string `pulumi:"name"`
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics []GetSystemSdnconnectorNic `pulumi:"nics"`
	// OCI certificate.
	OciCert string `pulumi:"ociCert"`
	// OCI pubkey fingerprint.
	OciFingerprint string `pulumi:"ociFingerprint"`
	// OCI server region.
	OciRegion string `pulumi:"ociRegion"`
	// OCI region type.
	OciRegionType string `pulumi:"ociRegionType"`
	// Password of the remote SDN connector as login credentials.
	Password string `pulumi:"password"`
	// Private key of GCP service account.
	PrivateKey string `pulumi:"privateKey"`
	// AWS region name.
	Region string `pulumi:"region"`
	// Resource group of Azure route table.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Azure Stack resource URL.
	ResourceUrl string `pulumi:"resourceUrl"`
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables []GetSystemSdnconnectorRouteTable `pulumi:"routeTables"`
	// Configure Azure route. The structure of `route` block is documented below.
	Routes []GetSystemSdnconnectorRoute `pulumi:"routes"`
	// AWS secret access key.
	SecretKey string `pulumi:"secretKey"`
	// Secret token of Kubernetes service account.
	SecretToken string `pulumi:"secretToken"`
	// Server address of the remote SDN connector.
	Server string `pulumi:"server"`
	// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
	ServerLists []GetSystemSdnconnectorServerList `pulumi:"serverLists"`
	// Port number of the remote SDN connector.
	ServerPort int `pulumi:"serverPort"`
	// GCP service account email.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Enable/disable connection to the remote SDN connector.
	Status string `pulumi:"status"`
	// Subscription ID of Azure route table.
	SubscriptionId string `pulumi:"subscriptionId"`
	// Tenant ID (directory ID).
	TenantId string `pulumi:"tenantId"`
	// Type of SDN connector.
	Type string `pulumi:"type"`
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval int `pulumi:"updateInterval"`
	// Enable/disable using IAM role from metadata to call API.
	UseMetadataIam string `pulumi:"useMetadataIam"`
	// User ID.
	UserId string `pulumi:"userId"`
	// Username of the remote SDN connector as login credentials.
	Username string `pulumi:"username"`
	// vCenter server password for NSX quarantine.
	VcenterPassword string `pulumi:"vcenterPassword"`
	// vCenter server address for NSX quarantine.
	VcenterServer string `pulumi:"vcenterServer"`
	// vCenter server username for NSX quarantine.
	VcenterUsername string  `pulumi:"vcenterUsername"`
	Vdomparam       *string `pulumi:"vdomparam"`
	// Enable/disable server certificate verification.
	VerifyCertificate string `pulumi:"verifyCertificate"`
	// AWS VPC ID.
	VpcId string `pulumi:"vpcId"`
}

func LookupSystemSdnconnectorOutput(ctx *pulumi.Context, args LookupSystemSdnconnectorOutputArgs, opts ...pulumi.InvokeOption) LookupSystemSdnconnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemSdnconnectorResult, error) {
			args := v.(LookupSystemSdnconnectorArgs)
			r, err := LookupSystemSdnconnector(ctx, &args, opts...)
			var s LookupSystemSdnconnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemSdnconnectorResultOutput)
}

// A collection of arguments for invoking getSystemSdnconnector.
type LookupSystemSdnconnectorOutputArgs struct {
	// Specify the name of the desired system sdnconnector.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemSdnconnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemSdnconnectorArgs)(nil)).Elem()
}

// A collection of values returned by getSystemSdnconnector.
type LookupSystemSdnconnectorResultOutput struct{ *pulumi.OutputState }

func (LookupSystemSdnconnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemSdnconnectorResult)(nil)).Elem()
}

func (o LookupSystemSdnconnectorResultOutput) ToLookupSystemSdnconnectorResultOutput() LookupSystemSdnconnectorResultOutput {
	return o
}

func (o LookupSystemSdnconnectorResultOutput) ToLookupSystemSdnconnectorResultOutputWithContext(ctx context.Context) LookupSystemSdnconnectorResultOutput {
	return o
}

// AWS access key ID.
func (o LookupSystemSdnconnectorResultOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.AccessKey }).(pulumi.StringOutput)
}

// IBM cloud API key or service ID API key.
func (o LookupSystemSdnconnectorResultOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.ApiKey }).(pulumi.StringOutput)
}

// Azure server region.
func (o LookupSystemSdnconnectorResultOutput) AzureRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.AzureRegion }).(pulumi.StringOutput)
}

// Azure client ID (application ID).
func (o LookupSystemSdnconnectorResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Azure client secret (application key).
func (o LookupSystemSdnconnectorResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// Compartment ID.
func (o LookupSystemSdnconnectorResultOutput) CompartmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.CompartmentId }).(pulumi.StringOutput)
}

// Compute generation for IBM cloud infrastructure.
func (o LookupSystemSdnconnectorResultOutput) ComputeGeneration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) int { return v.ComputeGeneration }).(pulumi.IntOutput)
}

// Domain name.
func (o LookupSystemSdnconnectorResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Domain }).(pulumi.StringOutput)
}

// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) ExternalAccountLists() GetSystemSdnconnectorExternalAccountListArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorExternalAccountList {
		return v.ExternalAccountLists
	}).(GetSystemSdnconnectorExternalAccountListArrayOutput)
}

// Configure GCP external IP. The structure of `externalIp` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) ExternalIps() GetSystemSdnconnectorExternalIpArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorExternalIp { return v.ExternalIps }).(GetSystemSdnconnectorExternalIpArrayOutput)
}

// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) ForwardingRules() GetSystemSdnconnectorForwardingRuleArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorForwardingRule { return v.ForwardingRules }).(GetSystemSdnconnectorForwardingRuleArrayOutput)
}

// GCP project name.
func (o LookupSystemSdnconnectorResultOutput) GcpProject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.GcpProject }).(pulumi.StringOutput)
}

// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) GcpProjectLists() GetSystemSdnconnectorGcpProjectListArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorGcpProjectList { return v.GcpProjectLists }).(GetSystemSdnconnectorGcpProjectListArrayOutput)
}

// Group name of computers.
func (o LookupSystemSdnconnectorResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// Enable/disable use for FortiGate HA service.
func (o LookupSystemSdnconnectorResultOutput) HaStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.HaStatus }).(pulumi.StringOutput)
}

// IBM cloud region name.
func (o LookupSystemSdnconnectorResultOutput) IbmRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.IbmRegion }).(pulumi.StringOutput)
}

// IBM cloud compute generation 1 region name.
func (o LookupSystemSdnconnectorResultOutput) IbmRegionGen1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.IbmRegionGen1 }).(pulumi.StringOutput)
}

// IBM cloud compute generation 2 region name.
func (o LookupSystemSdnconnectorResultOutput) IbmRegionGen2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.IbmRegionGen2 }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemSdnconnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Private key password.
func (o LookupSystemSdnconnectorResultOutput) KeyPasswd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.KeyPasswd }).(pulumi.StringOutput)
}

// Azure Stack login endpoint.
func (o LookupSystemSdnconnectorResultOutput) LoginEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.LoginEndpoint }).(pulumi.StringOutput)
}

// GCP zone name.
func (o LookupSystemSdnconnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configure Azure network interface. The structure of `nic` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) Nics() GetSystemSdnconnectorNicArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorNic { return v.Nics }).(GetSystemSdnconnectorNicArrayOutput)
}

// OCI certificate.
func (o LookupSystemSdnconnectorResultOutput) OciCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.OciCert }).(pulumi.StringOutput)
}

// OCI pubkey fingerprint.
func (o LookupSystemSdnconnectorResultOutput) OciFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.OciFingerprint }).(pulumi.StringOutput)
}

// OCI server region.
func (o LookupSystemSdnconnectorResultOutput) OciRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.OciRegion }).(pulumi.StringOutput)
}

// OCI region type.
func (o LookupSystemSdnconnectorResultOutput) OciRegionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.OciRegionType }).(pulumi.StringOutput)
}

// Password of the remote SDN connector as login credentials.
func (o LookupSystemSdnconnectorResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Password }).(pulumi.StringOutput)
}

// Private key of GCP service account.
func (o LookupSystemSdnconnectorResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

// AWS region name.
func (o LookupSystemSdnconnectorResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Region }).(pulumi.StringOutput)
}

// Resource group of Azure route table.
func (o LookupSystemSdnconnectorResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Azure Stack resource URL.
func (o LookupSystemSdnconnectorResultOutput) ResourceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.ResourceUrl }).(pulumi.StringOutput)
}

// Configure Azure route table. The structure of `routeTable` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) RouteTables() GetSystemSdnconnectorRouteTableArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorRouteTable { return v.RouteTables }).(GetSystemSdnconnectorRouteTableArrayOutput)
}

// Configure Azure route. The structure of `route` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) Routes() GetSystemSdnconnectorRouteArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorRoute { return v.Routes }).(GetSystemSdnconnectorRouteArrayOutput)
}

// AWS secret access key.
func (o LookupSystemSdnconnectorResultOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.SecretKey }).(pulumi.StringOutput)
}

// Secret token of Kubernetes service account.
func (o LookupSystemSdnconnectorResultOutput) SecretToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.SecretToken }).(pulumi.StringOutput)
}

// Server address of the remote SDN connector.
func (o LookupSystemSdnconnectorResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Server }).(pulumi.StringOutput)
}

// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
func (o LookupSystemSdnconnectorResultOutput) ServerLists() GetSystemSdnconnectorServerListArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) []GetSystemSdnconnectorServerList { return v.ServerLists }).(GetSystemSdnconnectorServerListArrayOutput)
}

// Port number of the remote SDN connector.
func (o LookupSystemSdnconnectorResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

// GCP service account email.
func (o LookupSystemSdnconnectorResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Enable/disable connection to the remote SDN connector.
func (o LookupSystemSdnconnectorResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Status }).(pulumi.StringOutput)
}

// Subscription ID of Azure route table.
func (o LookupSystemSdnconnectorResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Tenant ID (directory ID).
func (o LookupSystemSdnconnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Type of SDN connector.
func (o LookupSystemSdnconnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
func (o LookupSystemSdnconnectorResultOutput) UpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) int { return v.UpdateInterval }).(pulumi.IntOutput)
}

// Enable/disable using IAM role from metadata to call API.
func (o LookupSystemSdnconnectorResultOutput) UseMetadataIam() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.UseMetadataIam }).(pulumi.StringOutput)
}

// User ID.
func (o LookupSystemSdnconnectorResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.UserId }).(pulumi.StringOutput)
}

// Username of the remote SDN connector as login credentials.
func (o LookupSystemSdnconnectorResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.Username }).(pulumi.StringOutput)
}

// vCenter server password for NSX quarantine.
func (o LookupSystemSdnconnectorResultOutput) VcenterPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.VcenterPassword }).(pulumi.StringOutput)
}

// vCenter server address for NSX quarantine.
func (o LookupSystemSdnconnectorResultOutput) VcenterServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.VcenterServer }).(pulumi.StringOutput)
}

// vCenter server username for NSX quarantine.
func (o LookupSystemSdnconnectorResultOutput) VcenterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.VcenterUsername }).(pulumi.StringOutput)
}

func (o LookupSystemSdnconnectorResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable server certificate verification.
func (o LookupSystemSdnconnectorResultOutput) VerifyCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.VerifyCertificate }).(pulumi.StringOutput)
}

// AWS VPC ID.
func (o LookupSystemSdnconnectorResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnconnectorResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemSdnconnectorResultOutput{})
}
