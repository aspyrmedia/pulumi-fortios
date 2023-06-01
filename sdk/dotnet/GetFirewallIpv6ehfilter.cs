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
    public static class GetFirewallIpv6ehfilter
    {
        /// <summary>
        /// Use this data source to get information on fortios firewall ipv6ehfilter
        /// </summary>
        public static Task<GetFirewallIpv6ehfilterResult> InvokeAsync(GetFirewallIpv6ehfilterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallIpv6ehfilterResult>("fortios:index/getFirewallIpv6ehfilter:getFirewallIpv6ehfilter", args ?? new GetFirewallIpv6ehfilterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios firewall ipv6ehfilter
        /// </summary>
        public static Output<GetFirewallIpv6ehfilterResult> Invoke(GetFirewallIpv6ehfilterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallIpv6ehfilterResult>("fortios:index/getFirewallIpv6ehfilter:getFirewallIpv6ehfilter", args ?? new GetFirewallIpv6ehfilterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallIpv6ehfilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallIpv6ehfilterArgs()
        {
        }
        public static new GetFirewallIpv6ehfilterArgs Empty => new GetFirewallIpv6ehfilterArgs();
    }

    public sealed class GetFirewallIpv6ehfilterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallIpv6ehfilterInvokeArgs()
        {
        }
        public static new GetFirewallIpv6ehfilterInvokeArgs Empty => new GetFirewallIpv6ehfilterInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallIpv6ehfilterResult
    {
        /// <summary>
        /// Enable/disable blocking packets with the Authentication header (default = disable).
        /// </summary>
        public readonly string Auth;
        /// <summary>
        /// Enable/disable blocking packets with Destination Options headers (default = disable).
        /// </summary>
        public readonly string DestOpt;
        /// <summary>
        /// Enable/disable blocking packets with the Fragment header (default = disable).
        /// </summary>
        public readonly string Fragment;
        /// <summary>
        /// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        /// </summary>
        public readonly int HdoptType;
        /// <summary>
        /// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
        /// </summary>
        public readonly string HopOpt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable blocking packets with the No Next header (default = disable)
        /// </summary>
        public readonly string NoNext;
        /// <summary>
        /// Enable/disable blocking packets with Routing headers (default = enable).
        /// </summary>
        public readonly string Routing;
        /// <summary>
        /// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        /// </summary>
        public readonly int RoutingType;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallIpv6ehfilterResult(
            string auth,

            string destOpt,

            string fragment,

            int hdoptType,

            string hopOpt,

            string id,

            string noNext,

            string routing,

            int routingType,

            string? vdomparam)
        {
            Auth = auth;
            DestOpt = destOpt;
            Fragment = fragment;
            HdoptType = hdoptType;
            HopOpt = hopOpt;
            Id = id;
            NoNext = noNext;
            Routing = routing;
            RoutingType = routingType;
            Vdomparam = vdomparam;
        }
    }
}