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
    public sealed class SwitchcontrollerManagedswitchIgmpSnoopingVlan
    {
        /// <summary>
        /// IGMP snooping proxy for the VLAN interface. Valid values: `disable`, `enable`, `global`.
        /// </summary>
        public readonly string? Proxy;
        /// <summary>
        /// Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Querier;
        /// <summary>
        /// IGMP snooping querier address.
        /// </summary>
        public readonly string? QuerierAddr;
        /// <summary>
        /// IGMP snooping querier version.
        /// 
        /// The `n802_1x_settings` block supports:
        /// </summary>
        public readonly int? Version;
        /// <summary>
        /// List of FortiSwitch VLANs.
        /// </summary>
        public readonly string? VlanName;

        [OutputConstructor]
        private SwitchcontrollerManagedswitchIgmpSnoopingVlan(
            string? proxy,

            string? querier,

            string? querierAddr,

            int? version,

            string? vlanName)
        {
            Proxy = proxy;
            Querier = querier;
            QuerierAddr = querierAddr;
            Version = version;
            VlanName = vlanName;
        }
    }
}