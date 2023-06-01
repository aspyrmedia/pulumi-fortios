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
    public sealed class Logsyslogd4OverridesettingCustomFieldName
    {
        /// <summary>
        /// Field custom name.
        /// </summary>
        public readonly string? Custom;
        /// <summary>
        /// Entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Field name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private Logsyslogd4OverridesettingCustomFieldName(
            string? custom,

            int? id,

            string? name)
        {
            Custom = custom;
            Id = id;
            Name = name;
        }
    }
}