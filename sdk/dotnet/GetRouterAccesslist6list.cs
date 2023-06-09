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
    public static class GetRouterAccesslist6list
    {
        /// <summary>
        /// Provides a list of `fortios.RouterAccesslist6`.
        /// </summary>
        public static Task<GetRouterAccesslist6listResult> InvokeAsync(GetRouterAccesslist6listArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterAccesslist6listResult>("fortios:index/getRouterAccesslist6list:getRouterAccesslist6list", args ?? new GetRouterAccesslist6listArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.RouterAccesslist6`.
        /// </summary>
        public static Output<GetRouterAccesslist6listResult> Invoke(GetRouterAccesslist6listInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterAccesslist6listResult>("fortios:index/getRouterAccesslist6list:getRouterAccesslist6list", args ?? new GetRouterAccesslist6listInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterAccesslist6listArgs : global::Pulumi.InvokeArgs
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

        public GetRouterAccesslist6listArgs()
        {
        }
        public static new GetRouterAccesslist6listArgs Empty => new GetRouterAccesslist6listArgs();
    }

    public sealed class GetRouterAccesslist6listInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetRouterAccesslist6listInvokeArgs()
        {
        }
        public static new GetRouterAccesslist6listInvokeArgs Empty => new GetRouterAccesslist6listInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterAccesslist6listResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.RouterAccesslist6`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterAccesslist6listResult(
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
