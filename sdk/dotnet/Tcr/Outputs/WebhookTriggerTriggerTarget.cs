// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcr.Outputs
{

    [OutputType]
    public sealed class WebhookTriggerTriggerTarget
    {
        public readonly string Address;
        public readonly ImmutableArray<Outputs.WebhookTriggerTriggerTargetHeader> Headers;

        [OutputConstructor]
        private WebhookTriggerTriggerTarget(
            string address,

            ImmutableArray<Outputs.WebhookTriggerTriggerTargetHeader> headers)
        {
            Address = address;
            Headers = headers;
        }
    }
}