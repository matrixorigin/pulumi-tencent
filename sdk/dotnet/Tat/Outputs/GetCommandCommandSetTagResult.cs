// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tat.Outputs
{

    [OutputType]
    public sealed class GetCommandCommandSetTagResult
    {
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private GetCommandCommandSetTagResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}