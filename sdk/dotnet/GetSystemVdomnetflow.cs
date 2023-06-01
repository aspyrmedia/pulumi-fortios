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
    public static class GetSystemVdomnetflow
    {
        /// <summary>
        /// Use this data source to get information on fortios system vdomnetflow
        /// </summary>
        public static Task<GetSystemVdomnetflowResult> InvokeAsync(GetSystemVdomnetflowArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemVdomnetflowResult>("fortios:index/getSystemVdomnetflow:getSystemVdomnetflow", args ?? new GetSystemVdomnetflowArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system vdomnetflow
        /// </summary>
        public static Output<GetSystemVdomnetflowResult> Invoke(GetSystemVdomnetflowInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemVdomnetflowResult>("fortios:index/getSystemVdomnetflow:getSystemVdomnetflow", args ?? new GetSystemVdomnetflowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemVdomnetflowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemVdomnetflowArgs()
        {
        }
        public static new GetSystemVdomnetflowArgs Empty => new GetSystemVdomnetflowArgs();
    }

    public sealed class GetSystemVdomnetflowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemVdomnetflowInvokeArgs()
        {
        }
        public static new GetSystemVdomnetflowInvokeArgs Empty => new GetSystemVdomnetflowInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemVdomnetflowResult
    {
        /// <summary>
        /// NetFlow collector IP address.
        /// </summary>
        public readonly string CollectorIp;
        /// <summary>
        /// NetFlow collector port number.
        /// </summary>
        public readonly int CollectorPort;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server.
        /// </summary>
        public readonly string InterfaceSelectMethod;
        /// <summary>
        /// Source IP address for communication with the NetFlow agent.
        /// </summary>
        public readonly string SourceIp;
        /// <summary>
        /// Enable/disable NetFlow per VDOM.
        /// </summary>
        public readonly string VdomNetflow;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemVdomnetflowResult(
            string collectorIp,

            int collectorPort,

            string id,

            string @interface,

            string interfaceSelectMethod,

            string sourceIp,

            string vdomNetflow,

            string? vdomparam)
        {
            CollectorIp = collectorIp;
            CollectorPort = collectorPort;
            Id = id;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            SourceIp = sourceIp;
            VdomNetflow = vdomNetflow;
            Vdomparam = vdomparam;
        }
    }
}