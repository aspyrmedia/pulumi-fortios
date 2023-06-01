// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class ExtensioncontrollerExtenderprofileLanExtensionBackhaul
    {
        /// <summary>
        /// FortiExtender LAN extension backhaul name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// FortiExtender uplink port. Valid values: `primary`, `secondary`.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// WRR weight parameter.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private ExtensioncontrollerExtenderprofileLanExtensionBackhaul(
            string? name,

            string? port,

            string? role,

            int? weight)
        {
            Name = name;
            Port = port;
            Role = role;
            Weight = weight;
        }
    }
}