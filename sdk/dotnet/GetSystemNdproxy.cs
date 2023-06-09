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
    public static class GetSystemNdproxy
    {
        /// <summary>
        /// Use this data source to get information on fortios system ndproxy
        /// </summary>
        public static Task<GetSystemNdproxyResult> InvokeAsync(GetSystemNdproxyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemNdproxyResult>("fortios:index/getSystemNdproxy:getSystemNdproxy", args ?? new GetSystemNdproxyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system ndproxy
        /// </summary>
        public static Output<GetSystemNdproxyResult> Invoke(GetSystemNdproxyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemNdproxyResult>("fortios:index/getSystemNdproxy:getSystemNdproxy", args ?? new GetSystemNdproxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemNdproxyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemNdproxyArgs()
        {
        }
        public static new GetSystemNdproxyArgs Empty => new GetSystemNdproxyArgs();
    }

    public sealed class GetSystemNdproxyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemNdproxyInvokeArgs()
        {
        }
        public static new GetSystemNdproxyInvokeArgs Empty => new GetSystemNdproxyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemNdproxyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemNdproxyMemberResult> Members;
        /// <summary>
        /// Enable/disable neighbor discovery proxy.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemNdproxyResult(
            string id,

            ImmutableArray<Outputs.GetSystemNdproxyMemberResult> members,

            string status,

            string? vdomparam)
        {
            Id = id;
            Members = members;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
