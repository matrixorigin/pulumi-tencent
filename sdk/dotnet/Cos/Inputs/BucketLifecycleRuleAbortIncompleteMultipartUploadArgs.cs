// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Inputs
{

    public sealed class BucketLifecycleRuleAbortIncompleteMultipartUploadArgs : global::Pulumi.ResourceArgs
    {
        [Input("daysAfterInitiation", required: true)]
        public Input<int> DaysAfterInitiation { get; set; } = null!;

        public BucketLifecycleRuleAbortIncompleteMultipartUploadArgs()
        {
        }
        public static new BucketLifecycleRuleAbortIncompleteMultipartUploadArgs Empty => new BucketLifecycleRuleAbortIncompleteMultipartUploadArgs();
    }
}
