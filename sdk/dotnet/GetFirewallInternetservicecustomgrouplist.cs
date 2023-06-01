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
    public static class GetFirewallInternetservicecustomgrouplist
    {
        /// <summary>
        /// Provides a list of `fortios.FirewallInternetservicecustomgroup`.
        /// </summary>
        public static Task<GetFirewallInternetservicecustomgrouplistResult> InvokeAsync(GetFirewallInternetservicecustomgrouplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallInternetservicecustomgrouplistResult>("fortios:index/getFirewallInternetservicecustomgrouplist:getFirewallInternetservicecustomgrouplist", args ?? new GetFirewallInternetservicecustomgrouplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.FirewallInternetservicecustomgroup`.
        /// </summary>
        public static Output<GetFirewallInternetservicecustomgrouplistResult> Invoke(GetFirewallInternetservicecustomgrouplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallInternetservicecustomgrouplistResult>("fortios:index/getFirewallInternetservicecustomgrouplist:getFirewallInternetservicecustomgrouplist", args ?? new GetFirewallInternetservicecustomgrouplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallInternetservicecustomgrouplistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallInternetservicecustomgrouplistArgs()
        {
        }
        public static new GetFirewallInternetservicecustomgrouplistArgs Empty => new GetFirewallInternetservicecustomgrouplistArgs();
    }

    public sealed class GetFirewallInternetservicecustomgrouplistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallInternetservicecustomgrouplistInvokeArgs()
        {
        }
        public static new GetFirewallInternetservicecustomgrouplistInvokeArgs Empty => new GetFirewallInternetservicecustomgrouplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallInternetservicecustomgrouplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.FirewallInternetservicecustomgroup`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallInternetservicecustomgrouplistResult(
            string? filter,

            string id,

            ImmutableArray<string> namelists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Namelists = namelists;
            Vdomparam = vdomparam;
        }
    }
}