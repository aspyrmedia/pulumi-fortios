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

    public sealed class SystemVirtualwanlinkServiceDst6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Physical interface name.
        /// 
        /// 
        /// 
        /// 
        /// 
        /// The `dst6` block supports:
        /// 
        /// 
        /// The `src6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemVirtualwanlinkServiceDst6Args()
        {
        }
        public static new SystemVirtualwanlinkServiceDst6Args Empty => new SystemVirtualwanlinkServiceDst6Args();
    }
}