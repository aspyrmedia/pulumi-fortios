// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure AntiVirus profiles.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.AntivirusProfile("trname", new()
    ///     {
    ///         AnalyticsBlFiletype = 0,
    ///         AnalyticsDb = "disable",
    ///         AnalyticsMaxUpload = 10,
    ///         AnalyticsWlFiletype = 0,
    ///         AvBlockLog = "enable",
    ///         AvVirusLog = "enable",
    ///         ExtendedLog = "disable",
    ///         FtgdAnalytics = "disable",
    ///         InspectionMode = "flow-based",
    ///         MobileMalwareDb = "enable",
    ///         ScanMode = "quick",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Antivirus Profile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/antivirusProfile:AntivirusProfile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/antivirusProfile:AntivirusProfile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/antivirusProfile:AntivirusProfile")]
    public partial class AntivirusProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Only submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Output("analyticsAcceptFiletype")]
        public Output<int> AnalyticsAcceptFiletype { get; private set; } = null!;

        /// <summary>
        /// Only submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Output("analyticsBlFiletype")]
        public Output<int> AnalyticsBlFiletype { get; private set; } = null!;

        /// <summary>
        /// Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("analyticsDb")]
        public Output<string> AnalyticsDb { get; private set; } = null!;

        /// <summary>
        /// Do not submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Output("analyticsIgnoreFiletype")]
        public Output<int> AnalyticsIgnoreFiletype { get; private set; } = null!;

        /// <summary>
        /// Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
        /// </summary>
        [Output("analyticsMaxUpload")]
        public Output<int> AnalyticsMaxUpload { get; private set; } = null!;

        /// <summary>
        /// Do not submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Output("analyticsWlFiletype")]
        public Output<int> AnalyticsWlFiletype { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("avBlockLog")]
        public Output<string> AvBlockLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("avVirusLog")]
        public Output<string> AvVirusLog { get; private set; } = null!;

        /// <summary>
        /// Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
        /// </summary>
        [Output("cifs")]
        public Output<Outputs.AntivirusProfileCifs> Cifs { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
        /// </summary>
        [Output("contentDisarm")]
        public Output<Outputs.AntivirusProfileContentDisarm> ContentDisarm { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("emsThreatFeed")]
        public Output<string> EmsThreatFeed { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable all external blocklists. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("externalBlocklistEnableAll")]
        public Output<string> ExternalBlocklistEnableAll { get; private set; } = null!;

        /// <summary>
        /// One or more external malware block lists. The structure of `external_blocklist` block is documented below.
        /// </summary>
        [Output("externalBlocklists")]
        public Output<ImmutableArray<Outputs.AntivirusProfileExternalBlocklist>> ExternalBlocklists { get; private set; } = null!;

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Output("featureSet")]
        public Output<string> FeatureSet { get; private set; } = null!;

        /// <summary>
        /// Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Output("fortiaiErrorAction")]
        public Output<string> FortiaiErrorAction { get; private set; } = null!;

        /// <summary>
        /// Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Output("fortiaiTimeoutAction")]
        public Output<string> FortiaiTimeoutAction { get; private set; } = null!;

        /// <summary>
        /// Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Output("fortindrErrorAction")]
        public Output<string> FortindrErrorAction { get; private set; } = null!;

        /// <summary>
        /// Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Output("fortindrTimeoutAction")]
        public Output<string> FortindrTimeoutAction { get; private set; } = null!;

        /// <summary>
        /// Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Output("fortisandboxErrorAction")]
        public Output<string> FortisandboxErrorAction { get; private set; } = null!;

        /// <summary>
        /// Maximum size of files that can be uploaded to FortiSandbox.
        /// </summary>
        [Output("fortisandboxMaxUpload")]
        public Output<int> FortisandboxMaxUpload { get; private set; } = null!;

        /// <summary>
        /// FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.
        /// </summary>
        [Output("fortisandboxMode")]
        public Output<string> FortisandboxMode { get; private set; } = null!;

        /// <summary>
        /// Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Output("fortisandboxTimeoutAction")]
        public Output<string> FortisandboxTimeoutAction { get; private set; } = null!;

        /// <summary>
        /// Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
        /// </summary>
        [Output("ftgdAnalytics")]
        public Output<string> FtgdAnalytics { get; private set; } = null!;

        /// <summary>
        /// Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
        /// </summary>
        [Output("ftp")]
        public Output<Outputs.AntivirusProfileFtp> Ftp { get; private set; } = null!;

        /// <summary>
        /// Configure HTTP AntiVirus options. The structure of `http` block is documented below.
        /// </summary>
        [Output("http")]
        public Output<Outputs.AntivirusProfileHttp> Http { get; private set; } = null!;

        /// <summary>
        /// Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
        /// </summary>
        [Output("imap")]
        public Output<Outputs.AntivirusProfileImap> Imap { get; private set; } = null!;

        /// <summary>
        /// Inspection mode. Valid values: `proxy`, `flow-based`.
        /// </summary>
        [Output("inspectionMode")]
        public Output<string> InspectionMode { get; private set; } = null!;

        /// <summary>
        /// Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
        /// </summary>
        [Output("mapi")]
        public Output<Outputs.AntivirusProfileMapi> Mapi { get; private set; } = null!;

        /// <summary>
        /// Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("mobileMalwareDb")]
        public Output<string> MobileMalwareDb { get; private set; } = null!;

        /// <summary>
        /// Configure AntiVirus quarantine settings. The structure of `nac_quar` block is documented below.
        /// </summary>
        [Output("nacQuar")]
        public Output<Outputs.AntivirusProfileNacQuar> NacQuar { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
        /// </summary>
        [Output("nntp")]
        public Output<Outputs.AntivirusProfileNntp> Nntp { get; private set; } = null!;

        /// <summary>
        /// Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
        /// </summary>
        [Output("outbreakPrevention")]
        public Output<Outputs.AntivirusProfileOutbreakPrevention> OutbreakPrevention { get; private set; } = null!;

        /// <summary>
        /// Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("outbreakPreventionArchiveScan")]
        public Output<string> OutbreakPreventionArchiveScan { get; private set; } = null!;

        /// <summary>
        /// Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
        /// </summary>
        [Output("pop3")]
        public Output<Outputs.AntivirusProfilePop3> Pop3 { get; private set; } = null!;

        /// <summary>
        /// Replacement message group customized for this profile.
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// Choose between full scan mode and quick scan mode.
        /// </summary>
        [Output("scanMode")]
        public Output<string> ScanMode { get; private set; } = null!;

        /// <summary>
        /// Configure SMB AntiVirus options. The structure of `smb` block is documented below.
        /// </summary>
        [Output("smb")]
        public Output<Outputs.AntivirusProfileSmb> Smb { get; private set; } = null!;

        /// <summary>
        /// Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
        /// </summary>
        [Output("smtp")]
        public Output<Outputs.AntivirusProfileSmtp> Smtp { get; private set; } = null!;

        /// <summary>
        /// Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
        /// </summary>
        [Output("ssh")]
        public Output<Outputs.AntivirusProfileSsh> Ssh { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a AntivirusProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AntivirusProfile(string name, AntivirusProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/antivirusProfile:AntivirusProfile", name, args ?? new AntivirusProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AntivirusProfile(string name, Input<string> id, AntivirusProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/antivirusProfile:AntivirusProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AntivirusProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AntivirusProfile Get(string name, Input<string> id, AntivirusProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new AntivirusProfile(name, id, state, options);
        }
    }

    public sealed class AntivirusProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Only submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsAcceptFiletype")]
        public Input<int>? AnalyticsAcceptFiletype { get; set; }

        /// <summary>
        /// Only submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsBlFiletype")]
        public Input<int>? AnalyticsBlFiletype { get; set; }

        /// <summary>
        /// Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("analyticsDb")]
        public Input<string>? AnalyticsDb { get; set; }

        /// <summary>
        /// Do not submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsIgnoreFiletype")]
        public Input<int>? AnalyticsIgnoreFiletype { get; set; }

        /// <summary>
        /// Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
        /// </summary>
        [Input("analyticsMaxUpload")]
        public Input<int>? AnalyticsMaxUpload { get; set; }

        /// <summary>
        /// Do not submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsWlFiletype")]
        public Input<int>? AnalyticsWlFiletype { get; set; }

        /// <summary>
        /// Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avBlockLog")]
        public Input<string>? AvBlockLog { get; set; }

        /// <summary>
        /// Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avVirusLog")]
        public Input<string>? AvVirusLog { get; set; }

        /// <summary>
        /// Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
        /// </summary>
        [Input("cifs")]
        public Input<Inputs.AntivirusProfileCifsArgs>? Cifs { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
        /// </summary>
        [Input("contentDisarm")]
        public Input<Inputs.AntivirusProfileContentDisarmArgs>? ContentDisarm { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("emsThreatFeed")]
        public Input<string>? EmsThreatFeed { get; set; }

        /// <summary>
        /// Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Enable/disable all external blocklists. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("externalBlocklistEnableAll")]
        public Input<string>? ExternalBlocklistEnableAll { get; set; }

        [Input("externalBlocklists")]
        private InputList<Inputs.AntivirusProfileExternalBlocklistArgs>? _externalBlocklists;

        /// <summary>
        /// One or more external malware block lists. The structure of `external_blocklist` block is documented below.
        /// </summary>
        public InputList<Inputs.AntivirusProfileExternalBlocklistArgs> ExternalBlocklists
        {
            get => _externalBlocklists ?? (_externalBlocklists = new InputList<Inputs.AntivirusProfileExternalBlocklistArgs>());
            set => _externalBlocklists = value;
        }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortiaiErrorAction")]
        public Input<string>? FortiaiErrorAction { get; set; }

        /// <summary>
        /// Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortiaiTimeoutAction")]
        public Input<string>? FortiaiTimeoutAction { get; set; }

        /// <summary>
        /// Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortindrErrorAction")]
        public Input<string>? FortindrErrorAction { get; set; }

        /// <summary>
        /// Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortindrTimeoutAction")]
        public Input<string>? FortindrTimeoutAction { get; set; }

        /// <summary>
        /// Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortisandboxErrorAction")]
        public Input<string>? FortisandboxErrorAction { get; set; }

        /// <summary>
        /// Maximum size of files that can be uploaded to FortiSandbox.
        /// </summary>
        [Input("fortisandboxMaxUpload")]
        public Input<int>? FortisandboxMaxUpload { get; set; }

        /// <summary>
        /// FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.
        /// </summary>
        [Input("fortisandboxMode")]
        public Input<string>? FortisandboxMode { get; set; }

        /// <summary>
        /// Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortisandboxTimeoutAction")]
        public Input<string>? FortisandboxTimeoutAction { get; set; }

        /// <summary>
        /// Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
        /// </summary>
        [Input("ftgdAnalytics")]
        public Input<string>? FtgdAnalytics { get; set; }

        /// <summary>
        /// Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
        /// </summary>
        [Input("ftp")]
        public Input<Inputs.AntivirusProfileFtpArgs>? Ftp { get; set; }

        /// <summary>
        /// Configure HTTP AntiVirus options. The structure of `http` block is documented below.
        /// </summary>
        [Input("http")]
        public Input<Inputs.AntivirusProfileHttpArgs>? Http { get; set; }

        /// <summary>
        /// Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
        /// </summary>
        [Input("imap")]
        public Input<Inputs.AntivirusProfileImapArgs>? Imap { get; set; }

        /// <summary>
        /// Inspection mode. Valid values: `proxy`, `flow-based`.
        /// </summary>
        [Input("inspectionMode")]
        public Input<string>? InspectionMode { get; set; }

        /// <summary>
        /// Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.AntivirusProfileMapiArgs>? Mapi { get; set; }

        /// <summary>
        /// Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("mobileMalwareDb")]
        public Input<string>? MobileMalwareDb { get; set; }

        /// <summary>
        /// Configure AntiVirus quarantine settings. The structure of `nac_quar` block is documented below.
        /// </summary>
        [Input("nacQuar")]
        public Input<Inputs.AntivirusProfileNacQuarArgs>? NacQuar { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
        /// </summary>
        [Input("nntp")]
        public Input<Inputs.AntivirusProfileNntpArgs>? Nntp { get; set; }

        /// <summary>
        /// Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
        /// </summary>
        [Input("outbreakPrevention")]
        public Input<Inputs.AntivirusProfileOutbreakPreventionArgs>? OutbreakPrevention { get; set; }

        /// <summary>
        /// Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("outbreakPreventionArchiveScan")]
        public Input<string>? OutbreakPreventionArchiveScan { get; set; }

        /// <summary>
        /// Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
        /// </summary>
        [Input("pop3")]
        public Input<Inputs.AntivirusProfilePop3Args>? Pop3 { get; set; }

        /// <summary>
        /// Replacement message group customized for this profile.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Choose between full scan mode and quick scan mode.
        /// </summary>
        [Input("scanMode")]
        public Input<string>? ScanMode { get; set; }

        /// <summary>
        /// Configure SMB AntiVirus options. The structure of `smb` block is documented below.
        /// </summary>
        [Input("smb")]
        public Input<Inputs.AntivirusProfileSmbArgs>? Smb { get; set; }

        /// <summary>
        /// Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
        /// </summary>
        [Input("smtp")]
        public Input<Inputs.AntivirusProfileSmtpArgs>? Smtp { get; set; }

        /// <summary>
        /// Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.AntivirusProfileSshArgs>? Ssh { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AntivirusProfileArgs()
        {
        }
        public static new AntivirusProfileArgs Empty => new AntivirusProfileArgs();
    }

    public sealed class AntivirusProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Only submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsAcceptFiletype")]
        public Input<int>? AnalyticsAcceptFiletype { get; set; }

        /// <summary>
        /// Only submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsBlFiletype")]
        public Input<int>? AnalyticsBlFiletype { get; set; }

        /// <summary>
        /// Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("analyticsDb")]
        public Input<string>? AnalyticsDb { get; set; }

        /// <summary>
        /// Do not submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsIgnoreFiletype")]
        public Input<int>? AnalyticsIgnoreFiletype { get; set; }

        /// <summary>
        /// Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
        /// </summary>
        [Input("analyticsMaxUpload")]
        public Input<int>? AnalyticsMaxUpload { get; set; }

        /// <summary>
        /// Do not submit files matching this DLP file-pattern to FortiSandbox.
        /// </summary>
        [Input("analyticsWlFiletype")]
        public Input<int>? AnalyticsWlFiletype { get; set; }

        /// <summary>
        /// Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avBlockLog")]
        public Input<string>? AvBlockLog { get; set; }

        /// <summary>
        /// Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avVirusLog")]
        public Input<string>? AvVirusLog { get; set; }

        /// <summary>
        /// Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
        /// </summary>
        [Input("cifs")]
        public Input<Inputs.AntivirusProfileCifsGetArgs>? Cifs { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
        /// </summary>
        [Input("contentDisarm")]
        public Input<Inputs.AntivirusProfileContentDisarmGetArgs>? ContentDisarm { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("emsThreatFeed")]
        public Input<string>? EmsThreatFeed { get; set; }

        /// <summary>
        /// Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Enable/disable all external blocklists. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("externalBlocklistEnableAll")]
        public Input<string>? ExternalBlocklistEnableAll { get; set; }

        [Input("externalBlocklists")]
        private InputList<Inputs.AntivirusProfileExternalBlocklistGetArgs>? _externalBlocklists;

        /// <summary>
        /// One or more external malware block lists. The structure of `external_blocklist` block is documented below.
        /// </summary>
        public InputList<Inputs.AntivirusProfileExternalBlocklistGetArgs> ExternalBlocklists
        {
            get => _externalBlocklists ?? (_externalBlocklists = new InputList<Inputs.AntivirusProfileExternalBlocklistGetArgs>());
            set => _externalBlocklists = value;
        }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortiaiErrorAction")]
        public Input<string>? FortiaiErrorAction { get; set; }

        /// <summary>
        /// Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortiaiTimeoutAction")]
        public Input<string>? FortiaiTimeoutAction { get; set; }

        /// <summary>
        /// Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortindrErrorAction")]
        public Input<string>? FortindrErrorAction { get; set; }

        /// <summary>
        /// Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortindrTimeoutAction")]
        public Input<string>? FortindrTimeoutAction { get; set; }

        /// <summary>
        /// Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortisandboxErrorAction")]
        public Input<string>? FortisandboxErrorAction { get; set; }

        /// <summary>
        /// Maximum size of files that can be uploaded to FortiSandbox.
        /// </summary>
        [Input("fortisandboxMaxUpload")]
        public Input<int>? FortisandboxMaxUpload { get; set; }

        /// <summary>
        /// FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.
        /// </summary>
        [Input("fortisandboxMode")]
        public Input<string>? FortisandboxMode { get; set; }

        /// <summary>
        /// Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
        /// </summary>
        [Input("fortisandboxTimeoutAction")]
        public Input<string>? FortisandboxTimeoutAction { get; set; }

        /// <summary>
        /// Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
        /// </summary>
        [Input("ftgdAnalytics")]
        public Input<string>? FtgdAnalytics { get; set; }

        /// <summary>
        /// Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
        /// </summary>
        [Input("ftp")]
        public Input<Inputs.AntivirusProfileFtpGetArgs>? Ftp { get; set; }

        /// <summary>
        /// Configure HTTP AntiVirus options. The structure of `http` block is documented below.
        /// </summary>
        [Input("http")]
        public Input<Inputs.AntivirusProfileHttpGetArgs>? Http { get; set; }

        /// <summary>
        /// Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
        /// </summary>
        [Input("imap")]
        public Input<Inputs.AntivirusProfileImapGetArgs>? Imap { get; set; }

        /// <summary>
        /// Inspection mode. Valid values: `proxy`, `flow-based`.
        /// </summary>
        [Input("inspectionMode")]
        public Input<string>? InspectionMode { get; set; }

        /// <summary>
        /// Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.AntivirusProfileMapiGetArgs>? Mapi { get; set; }

        /// <summary>
        /// Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("mobileMalwareDb")]
        public Input<string>? MobileMalwareDb { get; set; }

        /// <summary>
        /// Configure AntiVirus quarantine settings. The structure of `nac_quar` block is documented below.
        /// </summary>
        [Input("nacQuar")]
        public Input<Inputs.AntivirusProfileNacQuarGetArgs>? NacQuar { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
        /// </summary>
        [Input("nntp")]
        public Input<Inputs.AntivirusProfileNntpGetArgs>? Nntp { get; set; }

        /// <summary>
        /// Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
        /// </summary>
        [Input("outbreakPrevention")]
        public Input<Inputs.AntivirusProfileOutbreakPreventionGetArgs>? OutbreakPrevention { get; set; }

        /// <summary>
        /// Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("outbreakPreventionArchiveScan")]
        public Input<string>? OutbreakPreventionArchiveScan { get; set; }

        /// <summary>
        /// Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
        /// </summary>
        [Input("pop3")]
        public Input<Inputs.AntivirusProfilePop3GetArgs>? Pop3 { get; set; }

        /// <summary>
        /// Replacement message group customized for this profile.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Choose between full scan mode and quick scan mode.
        /// </summary>
        [Input("scanMode")]
        public Input<string>? ScanMode { get; set; }

        /// <summary>
        /// Configure SMB AntiVirus options. The structure of `smb` block is documented below.
        /// </summary>
        [Input("smb")]
        public Input<Inputs.AntivirusProfileSmbGetArgs>? Smb { get; set; }

        /// <summary>
        /// Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
        /// </summary>
        [Input("smtp")]
        public Input<Inputs.AntivirusProfileSmtpGetArgs>? Smtp { get; set; }

        /// <summary>
        /// Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.AntivirusProfileSshGetArgs>? Ssh { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AntivirusProfileState()
        {
        }
        public static new AntivirusProfileState Empty => new AntivirusProfileState();
    }
}
