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
    public static class GetSystemProxyarplist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemProxyarp`.
        /// </summary>
        public static Task<GetSystemProxyarplistResult> InvokeAsync(GetSystemProxyarplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemProxyarplistResult>("fortios:index/getSystemProxyarplist:getSystemProxyarplist", args ?? new GetSystemProxyarplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemProxyarp`.
        /// </summary>
        public static Output<GetSystemProxyarplistResult> Invoke(GetSystemProxyarplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemProxyarplistResult>("fortios:index/getSystemProxyarplist:getSystemProxyarplist", args ?? new GetSystemProxyarplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemProxyarplistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemProxyarplistArgs()
        {
        }
        public static new GetSystemProxyarplistArgs Empty => new GetSystemProxyarplistArgs();
    }

    public sealed class GetSystemProxyarplistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemProxyarplistInvokeArgs()
        {
        }
        public static new GetSystemProxyarplistInvokeArgs Empty => new GetSystemProxyarplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemProxyarplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// A list of the `fortios.SystemProxyarp`.
        /// </summary>
        public readonly ImmutableArray<int> Fosidlists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemProxyarplistResult(
            string? filter,

            ImmutableArray<int> fosidlists,

            string id,

            string? vdomparam)
        {
            Filter = filter;
            Fosidlists = fosidlists;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}
