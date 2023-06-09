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
    public static class GetRouterAspathlist
    {
        /// <summary>
        /// Use this data source to get information on an fortios router aspathlist
        /// </summary>
        public static Task<GetRouterAspathlistResult> InvokeAsync(GetRouterAspathlistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterAspathlistResult>("fortios:index/getRouterAspathlist:getRouterAspathlist", args ?? new GetRouterAspathlistArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router aspathlist
        /// </summary>
        public static Output<GetRouterAspathlistResult> Invoke(GetRouterAspathlistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterAspathlistResult>("fortios:index/getRouterAspathlist:getRouterAspathlist", args ?? new GetRouterAspathlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterAspathlistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router aspathlist.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterAspathlistArgs()
        {
        }
        public static new GetRouterAspathlistArgs Empty => new GetRouterAspathlistArgs();
    }

    public sealed class GetRouterAspathlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router aspathlist.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterAspathlistInvokeArgs()
        {
        }
        public static new GetRouterAspathlistInvokeArgs Empty => new GetRouterAspathlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterAspathlistResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// AS path list name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// AS path list rule. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterAspathlistRuleResult> Rules;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterAspathlistResult(
            string id,

            string name,

            ImmutableArray<Outputs.GetRouterAspathlistRuleResult> rules,

            string? vdomparam)
        {
            Id = id;
            Name = name;
            Rules = rules;
            Vdomparam = vdomparam;
        }
    }
}
