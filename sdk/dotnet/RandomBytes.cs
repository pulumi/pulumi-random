// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Random
{
    /// <summary>
    /// The resource `random.RandomBytes` generates random bytes that are intended to be used as a secret, or key. Use this in preference to `random.RandomId` when the output is considered sensitive, and should not be displayed in the CLI.
    /// 
    /// This resource *does* use a cryptographic random number generator.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var jwtSecretRandomBytes = new Random.RandomBytes("jwtSecretRandomBytes", new()
    ///     {
    ///         Length = 64,
    ///     });
    /// 
    ///     var jwtSecretSecret = new Azure.KeyVault.Secret("jwtSecretSecret", new()
    ///     {
    ///         KeyVaultId = "some-azure-key-vault-id",
    ///         Value = jwtSecretRandomBytes.Base64,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Random bytes can be imported by specifying the value as base64 string.
    /// 
    /// ```sh
    /// $ pulumi import random:index/randomBytes:RandomBytes basic "8/fu3q+2DcgSJ19i0jZ5Cw=="
    /// ```
    /// </summary>
    [RandomResourceType("random:index/randomBytes:RandomBytes")]
    public partial class RandomBytes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The generated bytes presented in base64 string format.
        /// </summary>
        [Output("base64")]
        public Output<string> Base64 { get; private set; } = null!;

        /// <summary>
        /// The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the `length` parameter.
        /// </summary>
        [Output("hex")]
        public Output<string> Hex { get; private set; } = null!;

        /// <summary>
        /// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, string>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// The number of bytes requested. The minimum value for length is 1.
        /// </summary>
        [Output("length")]
        public Output<int> Length { get; private set; } = null!;


        /// <summary>
        /// Create a RandomBytes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomBytes(string name, RandomBytesArgs args, CustomResourceOptions? options = null)
            : base("random:index/randomBytes:RandomBytes", name, args ?? new RandomBytesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomBytes(string name, Input<string> id, RandomBytesState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomBytes:RandomBytes", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "base64",
                    "hex",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RandomBytes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomBytes Get(string name, Input<string> id, RandomBytesState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomBytes(name, id, state, options);
        }
    }

    public sealed class RandomBytesArgs : global::Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<string>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        /// </summary>
        public InputMap<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<string>());
            set => _keepers = value;
        }

        /// <summary>
        /// The number of bytes requested. The minimum value for length is 1.
        /// </summary>
        [Input("length", required: true)]
        public Input<int> Length { get; set; } = null!;

        public RandomBytesArgs()
        {
        }
        public static new RandomBytesArgs Empty => new RandomBytesArgs();
    }

    public sealed class RandomBytesState : global::Pulumi.ResourceArgs
    {
        [Input("base64")]
        private Input<string>? _base64;

        /// <summary>
        /// The generated bytes presented in base64 string format.
        /// </summary>
        public Input<string>? Base64
        {
            get => _base64;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _base64 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("hex")]
        private Input<string>? _hex;

        /// <summary>
        /// The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the `length` parameter.
        /// </summary>
        public Input<string>? Hex
        {
            get => _hex;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _hex = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("keepers")]
        private InputMap<string>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        /// </summary>
        public InputMap<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<string>());
            set => _keepers = value;
        }

        /// <summary>
        /// The number of bytes requested. The minimum value for length is 1.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        public RandomBytesState()
        {
        }
        public static new RandomBytesState Empty => new RandomBytesState();
    }
}
