// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskSourceResourceMySqlParam
    {
        public readonly string? DataSourceIncrementColumn;
        public readonly string? DataSourceIncrementMode;
        public readonly string? DataSourceMonitorMode;
        public readonly string? DataSourceMonitorResource;
        public readonly string? DataSourceStartFrom;
        public readonly string? DataTargetInsertMode;
        public readonly string? DataTargetPrimaryKeyField;
        public readonly ImmutableArray<Outputs.DatahubTaskSourceResourceMySqlParamDataTargetRecordMapping> DataTargetRecordMappings;
        public readonly string Database;
        public readonly string? DdlTopic;
        public readonly Outputs.DatahubTaskSourceResourceMySqlParamDropCls? DropCls;
        public readonly bool? DropInvalidMessage;
        public readonly string? IncludeContentChanges;
        public readonly bool? IncludeQuery;
        public readonly bool? IsTablePrefix;
        public readonly bool? IsTableRegular;
        public readonly string? KeyColumns;
        public readonly string? OutputFormat;
        public readonly bool? RecordWithSchema;
        public readonly string Resource;
        public readonly string? SignalDatabase;
        public readonly string? SnapshotMode;
        public readonly string Table;
        public readonly string? TopicRegex;
        public readonly string? TopicReplacement;

        [OutputConstructor]
        private DatahubTaskSourceResourceMySqlParam(
            string? dataSourceIncrementColumn,

            string? dataSourceIncrementMode,

            string? dataSourceMonitorMode,

            string? dataSourceMonitorResource,

            string? dataSourceStartFrom,

            string? dataTargetInsertMode,

            string? dataTargetPrimaryKeyField,

            ImmutableArray<Outputs.DatahubTaskSourceResourceMySqlParamDataTargetRecordMapping> dataTargetRecordMappings,

            string database,

            string? ddlTopic,

            Outputs.DatahubTaskSourceResourceMySqlParamDropCls? dropCls,

            bool? dropInvalidMessage,

            string? includeContentChanges,

            bool? includeQuery,

            bool? isTablePrefix,

            bool? isTableRegular,

            string? keyColumns,

            string? outputFormat,

            bool? recordWithSchema,

            string resource,

            string? signalDatabase,

            string? snapshotMode,

            string table,

            string? topicRegex,

            string? topicReplacement)
        {
            DataSourceIncrementColumn = dataSourceIncrementColumn;
            DataSourceIncrementMode = dataSourceIncrementMode;
            DataSourceMonitorMode = dataSourceMonitorMode;
            DataSourceMonitorResource = dataSourceMonitorResource;
            DataSourceStartFrom = dataSourceStartFrom;
            DataTargetInsertMode = dataTargetInsertMode;
            DataTargetPrimaryKeyField = dataTargetPrimaryKeyField;
            DataTargetRecordMappings = dataTargetRecordMappings;
            Database = database;
            DdlTopic = ddlTopic;
            DropCls = dropCls;
            DropInvalidMessage = dropInvalidMessage;
            IncludeContentChanges = includeContentChanges;
            IncludeQuery = includeQuery;
            IsTablePrefix = isTablePrefix;
            IsTableRegular = isTableRegular;
            KeyColumns = keyColumns;
            OutputFormat = outputFormat;
            RecordWithSchema = recordWithSchema;
            Resource = resource;
            SignalDatabase = signalDatabase;
            SnapshotMode = snapshotMode;
            Table = table;
            TopicRegex = topicRegex;
            TopicReplacement = topicReplacement;
        }
    }
}