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

    public sealed class VpnipsecPhase1CertificateGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate name.
        /// 
        /// The `ipv4_exclude_range` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VpnipsecPhase1CertificateGetArgs()
        {
        }
        public static new VpnipsecPhase1CertificateGetArgs Empty => new VpnipsecPhase1CertificateGetArgs();
    }
}