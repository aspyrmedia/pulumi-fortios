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
    public sealed class ApplicationListEntry
    {
        /// <summary>
        /// Pass or block traffic, or reset connection for traffic from this application. Valid values: `pass`, `block`, `reset`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// ID of allowed applications. The structure of `application` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntryApplication> Applications;
        /// <summary>
        /// Application behavior filter.
        /// </summary>
        public readonly string? Behavior;
        /// <summary>
        /// Category ID list. The structure of `category` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntryCategory> Categories;
        /// <summary>
        /// ID of excluded applications. The structure of `exclusion` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntryExclusion> Exclusions;
        /// <summary>
        /// Entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Enable/disable logging for this application list. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? LogPacket;
        /// <summary>
        /// Application parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntryParameter> Parameters;
        /// <summary>
        /// Per-IP traffic shaper.
        /// </summary>
        public readonly string? PerIpShaper;
        /// <summary>
        /// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        /// </summary>
        public readonly string? Popularity;
        /// <summary>
        /// Application protocol filter.
        /// </summary>
        public readonly string? Protocols;
        /// <summary>
        /// Quarantine method. Valid values: `none`, `attacker`.
        /// </summary>
        public readonly string? Quarantine;
        /// <summary>
        /// Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
        /// </summary>
        public readonly string? QuarantineExpiry;
        /// <summary>
        /// Enable/disable quarantine logging. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? QuarantineLog;
        /// <summary>
        /// Count of the rate.
        /// </summary>
        public readonly int? RateCount;
        /// <summary>
        /// Duration (sec) of the rate.
        /// </summary>
        public readonly int? RateDuration;
        /// <summary>
        /// Rate limit mode. Valid values: `periodical`, `continuous`.
        /// </summary>
        public readonly string? RateMode;
        /// <summary>
        /// Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`, `dhcp-client-mac`, `dns-domain`.
        /// </summary>
        public readonly string? RateTrack;
        /// <summary>
        /// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntryRisk> Risks;
        /// <summary>
        /// Session TTL (0 = default).
        /// </summary>
        public readonly int? SessionTtl;
        /// <summary>
        /// Traffic shaper.
        /// </summary>
        public readonly string? Shaper;
        /// <summary>
        /// Reverse traffic shaper.
        /// </summary>
        public readonly string? ShaperReverse;
        /// <summary>
        /// Application Sub-category ID list. The structure of `sub_category` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntrySubCategory> SubCategories;
        /// <summary>
        /// Application technology filter.
        /// </summary>
        public readonly string? Technology;
        /// <summary>
        /// Application vendor filter.
        /// </summary>
        public readonly string? Vendor;

        [OutputConstructor]
        private ApplicationListEntry(
            string? action,

            ImmutableArray<Outputs.ApplicationListEntryApplication> applications,

            string? behavior,

            ImmutableArray<Outputs.ApplicationListEntryCategory> categories,

            ImmutableArray<Outputs.ApplicationListEntryExclusion> exclusions,

            int? id,

            string? log,

            string? logPacket,

            ImmutableArray<Outputs.ApplicationListEntryParameter> parameters,

            string? perIpShaper,

            string? popularity,

            string? protocols,

            string? quarantine,

            string? quarantineExpiry,

            string? quarantineLog,

            int? rateCount,

            int? rateDuration,

            string? rateMode,

            string? rateTrack,

            ImmutableArray<Outputs.ApplicationListEntryRisk> risks,

            int? sessionTtl,

            string? shaper,

            string? shaperReverse,

            ImmutableArray<Outputs.ApplicationListEntrySubCategory> subCategories,

            string? technology,

            string? vendor)
        {
            Action = action;
            Applications = applications;
            Behavior = behavior;
            Categories = categories;
            Exclusions = exclusions;
            Id = id;
            Log = log;
            LogPacket = logPacket;
            Parameters = parameters;
            PerIpShaper = perIpShaper;
            Popularity = popularity;
            Protocols = protocols;
            Quarantine = quarantine;
            QuarantineExpiry = quarantineExpiry;
            QuarantineLog = quarantineLog;
            RateCount = rateCount;
            RateDuration = rateDuration;
            RateMode = rateMode;
            RateTrack = rateTrack;
            Risks = risks;
            SessionTtl = sessionTtl;
            Shaper = shaper;
            ShaperReverse = shaperReverse;
            SubCategories = subCategories;
            Technology = technology;
            Vendor = vendor;
        }
    }
}