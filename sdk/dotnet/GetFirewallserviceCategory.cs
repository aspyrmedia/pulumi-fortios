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
    public static class GetFirewallserviceCategory
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewallservice category
        /// </summary>
        public static Task<GetFirewallserviceCategoryResult> InvokeAsync(GetFirewallserviceCategoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallserviceCategoryResult>("fortios:index/getFirewallserviceCategory:getFirewallserviceCategory", args ?? new GetFirewallserviceCategoryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewallservice category
        /// </summary>
        public static Output<GetFirewallserviceCategoryResult> Invoke(GetFirewallserviceCategoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallserviceCategoryResult>("fortios:index/getFirewallserviceCategory:getFirewallserviceCategory", args ?? new GetFirewallserviceCategoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallserviceCategoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallservice category.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallserviceCategoryArgs()
        {
        }
        public static new GetFirewallserviceCategoryArgs Empty => new GetFirewallserviceCategoryArgs();
    }

    public sealed class GetFirewallserviceCategoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallservice category.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallserviceCategoryInvokeArgs()
        {
        }
        public static new GetFirewallserviceCategoryInvokeArgs Empty => new GetFirewallserviceCategoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallserviceCategoryResult
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Service category name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallserviceCategoryResult(
            string comment,

            string fabricObject,

            string id,

            string name,

            string? vdomparam)
        {
            Comment = comment;
            FabricObject = fabricObject;
            Id = id;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
