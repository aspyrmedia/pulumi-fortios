// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetFirewallAddress6template
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall address6template
        /// </summary>
        public static Task<GetFirewallAddress6templateResult> InvokeAsync(GetFirewallAddress6templateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallAddress6templateResult>("fortios:index/getFirewallAddress6template:getFirewallAddress6template", args ?? new GetFirewallAddress6templateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall address6template
        /// </summary>
        public static Output<GetFirewallAddress6templateResult> Invoke(GetFirewallAddress6templateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallAddress6templateResult>("fortios:index/getFirewallAddress6template:getFirewallAddress6template", args ?? new GetFirewallAddress6templateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallAddress6templateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall address6template.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallAddress6templateArgs()
        {
        }
        public static new GetFirewallAddress6templateArgs Empty => new GetFirewallAddress6templateArgs();
    }

    public sealed class GetFirewallAddress6templateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall address6template.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallAddress6templateInvokeArgs()
        {
        }
        public static new GetFirewallAddress6templateInvokeArgs Empty => new GetFirewallAddress6templateInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallAddress6templateResult
    {
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IPv6 address prefix.
        /// </summary>
        public readonly string Ip6;
        /// <summary>
        /// Subnet segment value name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of IPv6 subnet segments.
        /// </summary>
        public readonly int SubnetSegmentCount;
        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallAddress6templateSubnetSegmentResult> SubnetSegments;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallAddress6templateResult(
            string fabricObject,

            string id,

            string ip6,

            string name,

            int subnetSegmentCount,

            ImmutableArray<Outputs.GetFirewallAddress6templateSubnetSegmentResult> subnetSegments,

            string? vdomparam)
        {
            FabricObject = fabricObject;
            Id = id;
            Ip6 = ip6;
            Name = name;
            SubnetSegmentCount = subnetSegmentCount;
            SubnetSegments = subnetSegments;
            Vdomparam = vdomparam;
        }
    }
}