// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Tencentcloud
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("tencentcloud");

        private static readonly __Value<Pulumi.Tencentcloud.Config.Types.AssumeRole?> _assumeRole = new __Value<Pulumi.Tencentcloud.Config.Types.AssumeRole?>(() => __config.GetObject<Pulumi.Tencentcloud.Config.Types.AssumeRole>("assumeRole"));
        /// <summary>
        /// The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.
        /// </summary>
        public static Pulumi.Tencentcloud.Config.Types.AssumeRole? AssumeRole
        {
            get => _assumeRole.Get();
            set => _assumeRole.Set(value);
        }

        private static readonly __Value<Pulumi.Tencentcloud.Config.Types.AssumeRoleWithSaml?> _assumeRoleWithSaml = new __Value<Pulumi.Tencentcloud.Config.Types.AssumeRoleWithSaml?>(() => __config.GetObject<Pulumi.Tencentcloud.Config.Types.AssumeRoleWithSaml>("assumeRoleWithSaml"));
        /// <summary>
        /// The `assume_role_with_saml` block. If provided, terraform will attempt to assume this role using the supplied
        /// credentials.
        /// </summary>
        public static Pulumi.Tencentcloud.Config.Types.AssumeRoleWithSaml? AssumeRoleWithSaml
        {
            get => _assumeRoleWithSaml.Get();
            set => _assumeRoleWithSaml.Set(value);
        }

        private static readonly __Value<Pulumi.Tencentcloud.Config.Types.AssumeRoleWithWebIdentity?> _assumeRoleWithWebIdentity = new __Value<Pulumi.Tencentcloud.Config.Types.AssumeRoleWithWebIdentity?>(() => __config.GetObject<Pulumi.Tencentcloud.Config.Types.AssumeRoleWithWebIdentity>("assumeRoleWithWebIdentity"));
        /// <summary>
        /// The `assume_role_with_web_identity` block. If provided, terraform will attempt to assume this role using the supplied
        /// credentials.
        /// </summary>
        public static Pulumi.Tencentcloud.Config.Types.AssumeRoleWithWebIdentity? AssumeRoleWithWebIdentity
        {
            get => _assumeRoleWithWebIdentity.Get();
            set => _assumeRoleWithWebIdentity.Set(value);
        }

        private static readonly __Value<string?> _camRoleName = new __Value<string?>(() => __config.Get("camRoleName"));
        /// <summary>
        /// The name of the CVM instance CAM role. It can be sourced from the `TENCENTCLOUD_CAM_ROLE_NAME` environment variable.
        /// </summary>
        public static string? CamRoleName
        {
            get => _camRoleName.Get();
            set => _camRoleName.Set(value);
        }

        private static readonly __Value<string?> _cosDomain = new __Value<string?>(() => __config.Get("cosDomain"));
        /// <summary>
        /// The cos domain of the API request, Default is `https://cos.{region}.myqcloud.com`, Other Examples:
        /// `https://cluster-123456.cos-cdc.ap-guangzhou.myqcloud.com`.
        /// </summary>
        public static string? CosDomain
        {
            get => _cosDomain.Get();
            set => _cosDomain.Set(value);
        }

        private static readonly __Value<string?> _domain = new __Value<string?>(() => __config.Get("domain"));
        /// <summary>
        /// The root domain of the API request, Default is `tencentcloudapi.com`.
        /// </summary>
        public static string? Domain
        {
            get => _domain.Get();
            set => _domain.Set(value);
        }

        private static readonly __Value<bool?> _enablePodOidc = new __Value<bool?>(() => __config.GetBoolean("enablePodOidc"));
        /// <summary>
        /// Whether to enable pod oidc.
        /// </summary>
        public static bool? EnablePodOidc
        {
            get => _enablePodOidc.Get();
            set => _enablePodOidc.Set(value);
        }

        private static readonly __Value<string?> _profile = new __Value<string?>(() => __config.Get("profile"));
        /// <summary>
        /// The profile name as set in the shared credentials. It can also be sourced from the `TENCENTCLOUD_PROFILE` environment
        /// variable. If not set, the default profile created with `tccli configure` will be used.
        /// </summary>
        public static string? Profile
        {
            get => _profile.Get();
            set => _profile.Set(value);
        }

        private static readonly __Value<string?> _protocol = new __Value<string?>(() => __config.Get("protocol"));
        /// <summary>
        /// The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.
        /// </summary>
        public static string? Protocol
        {
            get => _protocol.Get();
            set => _protocol.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region") ?? Utilities.GetEnv("TENCENTCLOUD_REGION"));
        /// <summary>
        /// This is the TencentCloud region. It can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The
        /// default input value is ap-guangzhou.
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _secretId = new __Value<string?>(() => __config.Get("secretId") ?? Utilities.GetEnv("TENCENTCLOUD_SECRET_ID"));
        /// <summary>
        /// This is the TencentCloud access key. It can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.
        /// </summary>
        public static string? SecretId
        {
            get => _secretId.Get();
            set => _secretId.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey") ?? Utilities.GetEnv("TENCENTCLOUD_SECRET_KEY"));
        /// <summary>
        /// This is the TencentCloud secret key. It can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

        private static readonly __Value<string?> _securityToken = new __Value<string?>(() => __config.Get("securityToken") ?? Utilities.GetEnv("TENCENTCLOUD_SECURITY_TOKEN"));
        /// <summary>
        /// TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUD_SECURITY_TOKEN`
        /// environment variable. Notice: for supported products, please refer to: [temporary key supported
        /// products](https://intl.cloud.tencent.com/document/product/598/10588).
        /// </summary>
        public static string? SecurityToken
        {
            get => _securityToken.Get();
            set => _securityToken.Set(value);
        }

        private static readonly __Value<string?> _sharedCredentialsDir = new __Value<string?>(() => __config.Get("sharedCredentialsDir"));
        /// <summary>
        /// The directory of the shared credentials. It can also be sourced from the `TENCENTCLOUD_SHARED_CREDENTIALS_DIR`
        /// environment variable. If not set this defaults to ~/.tccli.
        /// </summary>
        public static string? SharedCredentialsDir
        {
            get => _sharedCredentialsDir.Get();
            set => _sharedCredentialsDir.Set(value);
        }

        public static class Types
        {

             public class AssumeRole
             {
                public string? ExternalId { get; set; } = null!;
                public string? Policy { get; set; } = null!;
                public string RoleArn { get; set; }
                public int SessionDuration { get; set; }
                public string SessionName { get; set; }
            }

             public class AssumeRoleWithSaml
             {
                public string PrincipalArn { get; set; }
                public string RoleArn { get; set; }
                public string SamlAssertion { get; set; }
                public int SessionDuration { get; set; }
                public string SessionName { get; set; }
            }

             public class AssumeRoleWithWebIdentity
             {
                public string RoleArn { get; set; }
                public int SessionDuration { get; set; }
                public string SessionName { get; set; }
                public string WebIdentityToken { get; set; }
            }
        }
    }
}