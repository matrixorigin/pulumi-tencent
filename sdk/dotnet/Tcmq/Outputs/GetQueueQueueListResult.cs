// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcmq.Outputs
{

    [OutputType]
    public sealed class GetQueueQueueListResult
    {
        public readonly int ActiveMsgNum;
        public readonly int Bps;
        public readonly int CreateTime;
        public readonly int CreateUin;
        public readonly ImmutableArray<Outputs.GetQueueQueueListDeadLetterPolicyResult> DeadLetterPolicies;
        public readonly ImmutableArray<Outputs.GetQueueQueueListDeadLetterSourceResult> DeadLetterSources;
        public readonly int DelayMsgNum;
        public readonly int InactiveMsgNum;
        public readonly int LastModifyTime;
        public readonly int MaxDelaySeconds;
        public readonly int MaxMsgBacklogSize;
        public readonly int MaxMsgHeapNum;
        public readonly int MaxMsgSize;
        public readonly int MaxUnackedMsgNum;
        public readonly int MinMsgTime;
        public readonly int MsgRetentionSeconds;
        public readonly string NamespaceName;
        public readonly int PollingWaitSeconds;
        public readonly int Qps;
        public readonly string QueueId;
        public readonly string QueueName;
        public readonly int RetentionSizeInMb;
        public readonly int RewindMsgNum;
        public readonly int RewindSeconds;
        public readonly int Status;
        public readonly ImmutableArray<Outputs.GetQueueQueueListTagResult> Tags;
        public readonly string TenantId;
        public readonly bool Trace;
        public readonly bool Transaction;
        public readonly ImmutableArray<Outputs.GetQueueQueueListTransactionPolicyResult> TransactionPolicies;
        public readonly int VisibilityTimeout;

        [OutputConstructor]
        private GetQueueQueueListResult(
            int activeMsgNum,

            int bps,

            int createTime,

            int createUin,

            ImmutableArray<Outputs.GetQueueQueueListDeadLetterPolicyResult> deadLetterPolicies,

            ImmutableArray<Outputs.GetQueueQueueListDeadLetterSourceResult> deadLetterSources,

            int delayMsgNum,

            int inactiveMsgNum,

            int lastModifyTime,

            int maxDelaySeconds,

            int maxMsgBacklogSize,

            int maxMsgHeapNum,

            int maxMsgSize,

            int maxUnackedMsgNum,

            int minMsgTime,

            int msgRetentionSeconds,

            string namespaceName,

            int pollingWaitSeconds,

            int qps,

            string queueId,

            string queueName,

            int retentionSizeInMb,

            int rewindMsgNum,

            int rewindSeconds,

            int status,

            ImmutableArray<Outputs.GetQueueQueueListTagResult> tags,

            string tenantId,

            bool trace,

            bool transaction,

            ImmutableArray<Outputs.GetQueueQueueListTransactionPolicyResult> transactionPolicies,

            int visibilityTimeout)
        {
            ActiveMsgNum = activeMsgNum;
            Bps = bps;
            CreateTime = createTime;
            CreateUin = createUin;
            DeadLetterPolicies = deadLetterPolicies;
            DeadLetterSources = deadLetterSources;
            DelayMsgNum = delayMsgNum;
            InactiveMsgNum = inactiveMsgNum;
            LastModifyTime = lastModifyTime;
            MaxDelaySeconds = maxDelaySeconds;
            MaxMsgBacklogSize = maxMsgBacklogSize;
            MaxMsgHeapNum = maxMsgHeapNum;
            MaxMsgSize = maxMsgSize;
            MaxUnackedMsgNum = maxUnackedMsgNum;
            MinMsgTime = minMsgTime;
            MsgRetentionSeconds = msgRetentionSeconds;
            NamespaceName = namespaceName;
            PollingWaitSeconds = pollingWaitSeconds;
            Qps = qps;
            QueueId = queueId;
            QueueName = queueName;
            RetentionSizeInMb = retentionSizeInMb;
            RewindMsgNum = rewindMsgNum;
            RewindSeconds = rewindSeconds;
            Status = status;
            Tags = tags;
            TenantId = tenantId;
            Trace = trace;
            Transaction = transaction;
            TransactionPolicies = transactionPolicies;
            VisibilityTimeout = visibilityTimeout;
        }
    }
}