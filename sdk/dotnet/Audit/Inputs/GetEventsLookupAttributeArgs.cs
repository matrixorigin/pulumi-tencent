// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Audit.Inputs
{

    public sealed class GetEventsLookupAttributeInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributeKey", required: true)]
        public Input<string> AttributeKey { get; set; } = null!;

        [Input("attributeValue")]
        public Input<string>? AttributeValue { get; set; }

        public GetEventsLookupAttributeInputArgs()
        {
        }
        public static new GetEventsLookupAttributeInputArgs Empty => new GetEventsLookupAttributeInputArgs();
    }
}
