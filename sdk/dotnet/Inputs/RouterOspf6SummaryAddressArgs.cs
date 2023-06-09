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

    public sealed class RouterOspf6SummaryAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable advertise status. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("advertise")]
        public Input<string>? Advertise { get; set; }

        /// <summary>
        /// Summary address entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv6 prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        /// <summary>
        /// Tag value.
        /// </summary>
        [Input("tag")]
        public Input<int>? Tag { get; set; }

        public RouterOspf6SummaryAddressArgs()
        {
        }
        public static new RouterOspf6SummaryAddressArgs Empty => new RouterOspf6SummaryAddressArgs();
    }
}
