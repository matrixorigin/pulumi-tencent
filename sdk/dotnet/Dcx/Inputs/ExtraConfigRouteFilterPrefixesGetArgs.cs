// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dcx.Inputs
{

    public sealed class ExtraConfigRouteFilterPrefixesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        public ExtraConfigRouteFilterPrefixesGetArgs()
        {
        }
        public static new ExtraConfigRouteFilterPrefixesGetArgs Empty => new ExtraConfigRouteFilterPrefixesGetArgs();
    }
}