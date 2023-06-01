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
    public sealed class SpamfilterDnsblEntry
    {
        /// <summary>
        /// Reject connection or mark as spam email. Valid values: `reject`, `spam`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// DNSBL/ORBL entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// DNSBL or ORBL server name.
        /// </summary>
        public readonly string? Server;
        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SpamfilterDnsblEntry(
            string? action,

            int? id,

            string? server,

            string? status)
        {
            Action = action;
            Id = id;
            Server = server;
            Status = status;
        }
    }
}