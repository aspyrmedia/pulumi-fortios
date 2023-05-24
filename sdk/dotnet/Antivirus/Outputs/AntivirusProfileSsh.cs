// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Aspyrmedia.Fortios.Antivirus.Outputs
{

    [OutputType]
    public sealed class AntivirusProfileSsh
    {
        public readonly string? ArchiveBlock;
        public readonly string? ArchiveLog;
        public readonly string? AvScan;
        public readonly string? Emulator;
        public readonly string? ExternalBlocklist;
        public readonly string? Fortiai;
        public readonly string? Fortindr;
        public readonly string? Fortisandbox;
        public readonly string? Options;
        public readonly string? OutbreakPrevention;
        public readonly string? Quarantine;

        [OutputConstructor]
        private AntivirusProfileSsh(
            string? archiveBlock,

            string? archiveLog,

            string? avScan,

            string? emulator,

            string? externalBlocklist,

            string? fortiai,

            string? fortindr,

            string? fortisandbox,

            string? options,

            string? outbreakPrevention,

            string? quarantine)
        {
            ArchiveBlock = archiveBlock;
            ArchiveLog = archiveLog;
            AvScan = avScan;
            Emulator = emulator;
            ExternalBlocklist = externalBlocklist;
            Fortiai = fortiai;
            Fortindr = fortindr;
            Fortisandbox = fortisandbox;
            Options = options;
            OutbreakPrevention = outbreakPrevention;
            Quarantine = quarantine;
        }
    }
}
