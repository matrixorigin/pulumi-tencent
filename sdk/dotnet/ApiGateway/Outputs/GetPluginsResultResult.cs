// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetPluginsResultResult
    {
        public readonly string ApiId;
        public readonly string ApiName;
        public readonly string ApiType;
        public readonly bool AttachedOtherPlugin;
        public readonly bool IsAttached;
        public readonly string Method;
        public readonly string Path;

        [OutputConstructor]
        private GetPluginsResultResult(
            string apiId,

            string apiName,

            string apiType,

            bool attachedOtherPlugin,

            bool isAttached,

            string method,

            string path)
        {
            ApiId = apiId;
            ApiName = apiName;
            ApiType = apiType;
            AttachedOtherPlugin = attachedOtherPlugin;
            IsAttached = isAttached;
            Method = method;
            Path = path;
        }
    }
}