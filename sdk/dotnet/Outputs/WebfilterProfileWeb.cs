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
    public sealed class WebfilterProfileWeb
    {
        /// <summary>
        /// FortiGuard allowlist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.
        /// </summary>
        public readonly string? Allowlist;
        /// <summary>
        /// Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Blacklist;
        /// <summary>
        /// Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Blocklist;
        /// <summary>
        /// Banned word table ID.
        /// </summary>
        public readonly int? BwordTable;
        /// <summary>
        /// Banned word score threshold.
        /// </summary>
        public readonly int? BwordThreshold;
        /// <summary>
        /// Content header list.
        /// </summary>
        public readonly int? ContentHeaderList;
        /// <summary>
        /// Search keywords to log when match is found. The structure of `keyword_match` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebfilterProfileWebKeywordMatch> KeywordMatches;
        /// <summary>
        /// Enable/disable logging all search phrases. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? LogSearch;
        /// <summary>
        /// Safe search type. Valid values: `url`, `header`.
        /// </summary>
        public readonly string? SafeSearch;
        /// <summary>
        /// URL filter table ID.
        /// </summary>
        public readonly int? UrlfilterTable;
        /// <summary>
        /// Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
        /// </summary>
        public readonly string? VimeoRestrict;
        /// <summary>
        /// FortiGuard whitelist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.
        /// </summary>
        public readonly string? Whitelist;
        /// <summary>
        /// YouTube EDU filter level. Valid values: `none`, `strict`, `moderate`.
        /// </summary>
        public readonly string? YoutubeRestrict;

        [OutputConstructor]
        private WebfilterProfileWeb(
            string? allowlist,

            string? blacklist,

            string? blocklist,

            int? bwordTable,

            int? bwordThreshold,

            int? contentHeaderList,

            ImmutableArray<Outputs.WebfilterProfileWebKeywordMatch> keywordMatches,

            string? logSearch,

            string? safeSearch,

            int? urlfilterTable,

            string? vimeoRestrict,

            string? whitelist,

            string? youtubeRestrict)
        {
            Allowlist = allowlist;
            Blacklist = blacklist;
            Blocklist = blocklist;
            BwordTable = bwordTable;
            BwordThreshold = bwordThreshold;
            ContentHeaderList = contentHeaderList;
            KeywordMatches = keywordMatches;
            LogSearch = logSearch;
            SafeSearch = safeSearch;
            UrlfilterTable = urlfilterTable;
            VimeoRestrict = vimeoRestrict;
            Whitelist = whitelist;
            YoutubeRestrict = youtubeRestrict;
        }
    }
}