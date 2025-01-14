// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo
{
    [TencentcloudResourceType("tencentcloud:Teo/zoneSetting:ZoneSetting")]
    public partial class ZoneSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Acceleration area of the zone. Valid values: `mainland`, `overseas`.
        /// </summary>
        [Output("area")]
        public Output<string> Area { get; private set; } = null!;

        /// <summary>
        /// Cache expiration time configuration.
        /// </summary>
        [Output("cache")]
        public Output<Outputs.ZoneSettingCache> Cache { get; private set; } = null!;

        /// <summary>
        /// Node cache key configuration.
        /// </summary>
        [Output("cacheKey")]
        public Output<Outputs.ZoneSettingCacheKey> CacheKey { get; private set; } = null!;

        /// <summary>
        /// Cache pre-refresh configuration.
        /// </summary>
        [Output("cachePrefresh")]
        public Output<Outputs.ZoneSettingCachePrefresh> CachePrefresh { get; private set; } = null!;

        /// <summary>
        /// Origin-pull client IP header configuration.
        /// </summary>
        [Output("clientIpHeader")]
        public Output<Outputs.ZoneSettingClientIpHeader> ClientIpHeader { get; private set; } = null!;

        /// <summary>
        /// Smart compression configuration.
        /// </summary>
        [Output("compression")]
        public Output<Outputs.ZoneSettingCompression> Compression { get; private set; } = null!;

        /// <summary>
        /// Force HTTPS redirect configuration.
        /// </summary>
        [Output("forceRedirect")]
        public Output<Outputs.ZoneSettingForceRedirect> ForceRedirect { get; private set; } = null!;

        /// <summary>
        /// HTTPS acceleration configuration.
        /// </summary>
        [Output("https")]
        public Output<Outputs.ZoneSettingHttps> Https { get; private set; } = null!;

        /// <summary>
        /// IPv6 access configuration.
        /// </summary>
        [Output("ipv6")]
        public Output<Outputs.ZoneSettingIpv6> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// Browser cache configuration.
        /// </summary>
        [Output("maxAge")]
        public Output<Outputs.ZoneSettingMaxAge> MaxAge { get; private set; } = null!;

        /// <summary>
        /// Offline cache configuration.
        /// </summary>
        [Output("offlineCache")]
        public Output<Outputs.ZoneSettingOfflineCache> OfflineCache { get; private set; } = null!;

        /// <summary>
        /// Origin server configuration.
        /// </summary>
        [Output("origin")]
        public Output<Outputs.ZoneSettingOrigin> Origin { get; private set; } = null!;

        /// <summary>
        /// Maximum size of files transferred over POST request.
        /// </summary>
        [Output("postMaxSize")]
        public Output<Outputs.ZoneSettingPostMaxSize> PostMaxSize { get; private set; } = null!;

        /// <summary>
        /// QUIC access configuration.
        /// </summary>
        [Output("quic")]
        public Output<Outputs.ZoneSettingQuic> Quic { get; private set; } = null!;

        /// <summary>
        /// Smart acceleration configuration.
        /// </summary>
        [Output("smartRouting")]
        public Output<Outputs.ZoneSettingSmartRouting> SmartRouting { get; private set; } = null!;

        /// <summary>
        /// HTTP2 origin-pull configuration.
        /// </summary>
        [Output("upstreamHttp2")]
        public Output<Outputs.ZoneSettingUpstreamHttp2> UpstreamHttp2 { get; private set; } = null!;

        /// <summary>
        /// WebSocket configuration.
        /// </summary>
        [Output("webSocket")]
        public Output<Outputs.ZoneSettingWebSocket> WebSocket { get; private set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a ZoneSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZoneSetting(string name, ZoneSettingArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/zoneSetting:ZoneSetting", name, args ?? new ZoneSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZoneSetting(string name, Input<string> id, ZoneSettingState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/zoneSetting:ZoneSetting", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ZoneSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZoneSetting Get(string name, Input<string> id, ZoneSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new ZoneSetting(name, id, state, options);
        }
    }

    public sealed class ZoneSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cache expiration time configuration.
        /// </summary>
        [Input("cache")]
        public Input<Inputs.ZoneSettingCacheArgs>? Cache { get; set; }

        /// <summary>
        /// Node cache key configuration.
        /// </summary>
        [Input("cacheKey")]
        public Input<Inputs.ZoneSettingCacheKeyArgs>? CacheKey { get; set; }

        /// <summary>
        /// Cache pre-refresh configuration.
        /// </summary>
        [Input("cachePrefresh")]
        public Input<Inputs.ZoneSettingCachePrefreshArgs>? CachePrefresh { get; set; }

        /// <summary>
        /// Origin-pull client IP header configuration.
        /// </summary>
        [Input("clientIpHeader")]
        public Input<Inputs.ZoneSettingClientIpHeaderArgs>? ClientIpHeader { get; set; }

        /// <summary>
        /// Smart compression configuration.
        /// </summary>
        [Input("compression")]
        public Input<Inputs.ZoneSettingCompressionArgs>? Compression { get; set; }

        /// <summary>
        /// Force HTTPS redirect configuration.
        /// </summary>
        [Input("forceRedirect")]
        public Input<Inputs.ZoneSettingForceRedirectArgs>? ForceRedirect { get; set; }

        /// <summary>
        /// HTTPS acceleration configuration.
        /// </summary>
        [Input("https")]
        public Input<Inputs.ZoneSettingHttpsArgs>? Https { get; set; }

        /// <summary>
        /// IPv6 access configuration.
        /// </summary>
        [Input("ipv6")]
        public Input<Inputs.ZoneSettingIpv6Args>? Ipv6 { get; set; }

        /// <summary>
        /// Browser cache configuration.
        /// </summary>
        [Input("maxAge")]
        public Input<Inputs.ZoneSettingMaxAgeArgs>? MaxAge { get; set; }

        /// <summary>
        /// Offline cache configuration.
        /// </summary>
        [Input("offlineCache")]
        public Input<Inputs.ZoneSettingOfflineCacheArgs>? OfflineCache { get; set; }

        /// <summary>
        /// Origin server configuration.
        /// </summary>
        [Input("origin")]
        public Input<Inputs.ZoneSettingOriginArgs>? Origin { get; set; }

        /// <summary>
        /// Maximum size of files transferred over POST request.
        /// </summary>
        [Input("postMaxSize")]
        public Input<Inputs.ZoneSettingPostMaxSizeArgs>? PostMaxSize { get; set; }

        /// <summary>
        /// QUIC access configuration.
        /// </summary>
        [Input("quic")]
        public Input<Inputs.ZoneSettingQuicArgs>? Quic { get; set; }

        /// <summary>
        /// Smart acceleration configuration.
        /// </summary>
        [Input("smartRouting")]
        public Input<Inputs.ZoneSettingSmartRoutingArgs>? SmartRouting { get; set; }

        /// <summary>
        /// HTTP2 origin-pull configuration.
        /// </summary>
        [Input("upstreamHttp2")]
        public Input<Inputs.ZoneSettingUpstreamHttp2Args>? UpstreamHttp2 { get; set; }

        /// <summary>
        /// WebSocket configuration.
        /// </summary>
        [Input("webSocket")]
        public Input<Inputs.ZoneSettingWebSocketArgs>? WebSocket { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public ZoneSettingArgs()
        {
        }
        public static new ZoneSettingArgs Empty => new ZoneSettingArgs();
    }

    public sealed class ZoneSettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceleration area of the zone. Valid values: `mainland`, `overseas`.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// Cache expiration time configuration.
        /// </summary>
        [Input("cache")]
        public Input<Inputs.ZoneSettingCacheGetArgs>? Cache { get; set; }

        /// <summary>
        /// Node cache key configuration.
        /// </summary>
        [Input("cacheKey")]
        public Input<Inputs.ZoneSettingCacheKeyGetArgs>? CacheKey { get; set; }

        /// <summary>
        /// Cache pre-refresh configuration.
        /// </summary>
        [Input("cachePrefresh")]
        public Input<Inputs.ZoneSettingCachePrefreshGetArgs>? CachePrefresh { get; set; }

        /// <summary>
        /// Origin-pull client IP header configuration.
        /// </summary>
        [Input("clientIpHeader")]
        public Input<Inputs.ZoneSettingClientIpHeaderGetArgs>? ClientIpHeader { get; set; }

        /// <summary>
        /// Smart compression configuration.
        /// </summary>
        [Input("compression")]
        public Input<Inputs.ZoneSettingCompressionGetArgs>? Compression { get; set; }

        /// <summary>
        /// Force HTTPS redirect configuration.
        /// </summary>
        [Input("forceRedirect")]
        public Input<Inputs.ZoneSettingForceRedirectGetArgs>? ForceRedirect { get; set; }

        /// <summary>
        /// HTTPS acceleration configuration.
        /// </summary>
        [Input("https")]
        public Input<Inputs.ZoneSettingHttpsGetArgs>? Https { get; set; }

        /// <summary>
        /// IPv6 access configuration.
        /// </summary>
        [Input("ipv6")]
        public Input<Inputs.ZoneSettingIpv6GetArgs>? Ipv6 { get; set; }

        /// <summary>
        /// Browser cache configuration.
        /// </summary>
        [Input("maxAge")]
        public Input<Inputs.ZoneSettingMaxAgeGetArgs>? MaxAge { get; set; }

        /// <summary>
        /// Offline cache configuration.
        /// </summary>
        [Input("offlineCache")]
        public Input<Inputs.ZoneSettingOfflineCacheGetArgs>? OfflineCache { get; set; }

        /// <summary>
        /// Origin server configuration.
        /// </summary>
        [Input("origin")]
        public Input<Inputs.ZoneSettingOriginGetArgs>? Origin { get; set; }

        /// <summary>
        /// Maximum size of files transferred over POST request.
        /// </summary>
        [Input("postMaxSize")]
        public Input<Inputs.ZoneSettingPostMaxSizeGetArgs>? PostMaxSize { get; set; }

        /// <summary>
        /// QUIC access configuration.
        /// </summary>
        [Input("quic")]
        public Input<Inputs.ZoneSettingQuicGetArgs>? Quic { get; set; }

        /// <summary>
        /// Smart acceleration configuration.
        /// </summary>
        [Input("smartRouting")]
        public Input<Inputs.ZoneSettingSmartRoutingGetArgs>? SmartRouting { get; set; }

        /// <summary>
        /// HTTP2 origin-pull configuration.
        /// </summary>
        [Input("upstreamHttp2")]
        public Input<Inputs.ZoneSettingUpstreamHttp2GetArgs>? UpstreamHttp2 { get; set; }

        /// <summary>
        /// WebSocket configuration.
        /// </summary>
        [Input("webSocket")]
        public Input<Inputs.ZoneSettingWebSocketGetArgs>? WebSocket { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ZoneSettingState()
        {
        }
        public static new ZoneSettingState Empty => new ZoneSettingState();
    }
}
