// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Outputs
{

    [OutputType]
    public sealed class GetDataSourceListOrderFieldResult
    {
        public readonly string Direction;
        public readonly string Name;

        [OutputConstructor]
        private GetDataSourceListOrderFieldResult(
            string direction,

            string name)
        {
            Direction = direction;
            Name = name;
        }
    }
}