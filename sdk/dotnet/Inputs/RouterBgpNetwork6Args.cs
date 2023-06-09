// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class RouterBgpNetwork6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("backdoor")]
        public Input<string>? Backdoor { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable ensure BGP network route exists in IGP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("networkImportCheck")]
        public Input<string>? NetworkImportCheck { get; set; }

        /// <summary>
        /// Aggregate IPv6 prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        /// <summary>
        /// Route map to modify generated route.
        /// 
        /// The `network6` block supports:
        /// 
        /// 
        /// 
        /// The `redistribute6` block supports:
        /// </summary>
        [Input("routeMap")]
        public Input<string>? RouteMap { get; set; }

        public RouterBgpNetwork6Args()
        {
        }
        public static new RouterBgpNetwork6Args Empty => new RouterBgpNetwork6Args();
    }
}
