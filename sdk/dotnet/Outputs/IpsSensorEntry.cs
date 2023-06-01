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
    public sealed class IpsSensorEntry
    {
        /// <summary>
        /// Action taken with traffic in which signatures are detected. Valid values: `pass`, `block`, `reset`, `default`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Applications to be protected. set application ? lists available applications. all includes all applications. other includes all unlisted applications.
        /// </summary>
        public readonly string? Application;
        /// <summary>
        /// List of CVE IDs of the signatures to add to the sensor The structure of `cve` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpsSensorEntryCfe> Cves;
        /// <summary>
        /// Signature default action filter. Valid values: `all`, `pass`, `block`.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// Signature default status filter. Valid values: `all`, `enable`, `disable`.
        /// </summary>
        public readonly string? DefaultStatus;
        /// <summary>
        /// Traffic from selected source or destination IP addresses is exempt from this signature. The structure of `exempt_ip` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpsSensorEntryExemptIp> ExemptIps;
        /// <summary>
        /// Rule ID in IPS database (0 - 4294967295).
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Filter by signature last modified date. Formats: before &lt;date&gt;, after &lt;date&gt;, between &lt;start-date&gt; &lt;end-date&gt;.
        /// </summary>
        public readonly string? LastModified;
        /// <summary>
        /// Protect client or server traffic.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Enable/disable logging of signatures included in filter. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Enable/disable logging of attack context: URL buffer, header buffer, body buffer, packet buffer. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? LogAttackContext;
        /// <summary>
        /// Enable/disable packet logging. Enable to save the packet that triggers the filter. You can download the packets in pcap format for diagnostic use. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? LogPacket;
        /// <summary>
        /// Operating systems to be protected.  all includes all operating systems. other includes all unlisted operating systems.
        /// </summary>
        public readonly string? Os;
        /// <summary>
        /// Protocols to be examined. set protocol ? lists available protocols. all includes all protocols. other includes all unlisted protocols.
        /// </summary>
        public readonly string? Protocol;
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
        /// Identifies the predefined or custom IPS signatures to add to the sensor. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpsSensorEntryRule> Rules;
        /// <summary>
        /// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        /// </summary>
        public readonly string? Severity;
        /// <summary>
        /// Status of the signatures included in filter. default enables the filter and only use filters with default status of enable. Filters with default status of disable will not be used. Valid values: `disable`, `enable`, `default`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// List of signature vulnerability types to filter by. The structure of `vuln_type` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpsSensorEntryVulnType> VulnTypes;

        [OutputConstructor]
        private IpsSensorEntry(
            string? action,

            string? application,

            ImmutableArray<Outputs.IpsSensorEntryCfe> cves,

            string? defaultAction,

            string? defaultStatus,

            ImmutableArray<Outputs.IpsSensorEntryExemptIp> exemptIps,

            int? id,

            string? lastModified,

            string? location,

            string? log,

            string? logAttackContext,

            string? logPacket,

            string? os,

            string? protocol,

            string? quarantine,

            string? quarantineExpiry,

            string? quarantineLog,

            int? rateCount,

            int? rateDuration,

            string? rateMode,

            string? rateTrack,

            ImmutableArray<Outputs.IpsSensorEntryRule> rules,

            string? severity,

            string? status,

            ImmutableArray<Outputs.IpsSensorEntryVulnType> vulnTypes)
        {
            Action = action;
            Application = application;
            Cves = cves;
            DefaultAction = defaultAction;
            DefaultStatus = defaultStatus;
            ExemptIps = exemptIps;
            Id = id;
            LastModified = lastModified;
            Location = location;
            Log = log;
            LogAttackContext = logAttackContext;
            LogPacket = logPacket;
            Os = os;
            Protocol = protocol;
            Quarantine = quarantine;
            QuarantineExpiry = quarantineExpiry;
            QuarantineLog = quarantineLog;
            RateCount = rateCount;
            RateDuration = rateDuration;
            RateMode = rateMode;
            RateTrack = rateTrack;
            Rules = rules;
            Severity = severity;
            Status = status;
            VulnTypes = vulnTypes;
        }
    }
}