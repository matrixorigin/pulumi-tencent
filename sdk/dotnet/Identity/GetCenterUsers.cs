// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Identity
{
    public static class GetCenterUsers
    {
        public static Task<GetCenterUsersResult> InvokeAsync(GetCenterUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCenterUsersResult>("tencentcloud:Identity/getCenterUsers:getCenterUsers", args ?? new GetCenterUsersArgs(), options.WithDefaults());

        public static Output<GetCenterUsersResult> Invoke(GetCenterUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCenterUsersResult>("tencentcloud:Identity/getCenterUsers:getCenterUsers", args ?? new GetCenterUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCenterUsersArgs : global::Pulumi.InvokeArgs
    {
        [Input("filter")]
        public string? Filter { get; set; }

        [Input("filterGroups")]
        private List<string>? _filterGroups;
        public List<string> FilterGroups
        {
            get => _filterGroups ?? (_filterGroups = new List<string>());
            set => _filterGroups = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("sortField")]
        public string? SortField { get; set; }

        [Input("sortType")]
        public string? SortType { get; set; }

        [Input("userStatus")]
        public string? UserStatus { get; set; }

        [Input("userType")]
        public string? UserType { get; set; }

        [Input("zoneId", required: true)]
        public string ZoneId { get; set; } = null!;

        public GetCenterUsersArgs()
        {
        }
        public static new GetCenterUsersArgs Empty => new GetCenterUsersArgs();
    }

    public sealed class GetCenterUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("filterGroups")]
        private InputList<string>? _filterGroups;
        public InputList<string> FilterGroups
        {
            get => _filterGroups ?? (_filterGroups = new InputList<string>());
            set => _filterGroups = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("sortField")]
        public Input<string>? SortField { get; set; }

        [Input("sortType")]
        public Input<string>? SortType { get; set; }

        [Input("userStatus")]
        public Input<string>? UserStatus { get; set; }

        [Input("userType")]
        public Input<string>? UserType { get; set; }

        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public GetCenterUsersInvokeArgs()
        {
        }
        public static new GetCenterUsersInvokeArgs Empty => new GetCenterUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetCenterUsersResult
    {
        public readonly string? Filter;
        public readonly ImmutableArray<string> FilterGroups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string? SortField;
        public readonly string? SortType;
        public readonly string? UserStatus;
        public readonly string? UserType;
        public readonly ImmutableArray<Outputs.GetCenterUsersUserResult> Users;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetCenterUsersResult(
            string? filter,

            ImmutableArray<string> filterGroups,

            string id,

            string? resultOutputFile,

            string? sortField,

            string? sortType,

            string? userStatus,

            string? userType,

            ImmutableArray<Outputs.GetCenterUsersUserResult> users,

            string zoneId)
        {
            Filter = filter;
            FilterGroups = filterGroups;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SortField = sortField;
            SortType = sortType;
            UserStatus = userStatus;
            UserType = userType;
            Users = users;
            ZoneId = zoneId;
        }
    }
}
