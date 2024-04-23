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
    /// The resource `random.RandomString` generates a random permutation of alphanumeric characters and optionally special characters.
    /// 
    /// This resource *does* use a cryptographic random number generator.
    /// 
    /// Historically this resource's intended usage has been ambiguous as the original example used it in a password. For backwards compatibility it will continue to exist. For unique ids please use random_id, for sensitive random values please use random_password.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var random = new Random.RandomString("random", new()
    ///     {
    ///         Length = 16,
    ///         OverrideSpecial = "/@£$",
    ///         Special = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import external strings into your Pulumi programs as RandomString resources as follows:
    /// </summary>
    [RandomResourceType("random:index/randomString:RandomString")]
    public partial class RandomString : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, string>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        /// </summary>
        [Output("length")]
        public Output<int> Length { get; private set; } = null!;

        /// <summary>
        /// Include lowercase alphabet characters in the result. Default value is `true`.
        /// </summary>
        [Output("lower")]
        public Output<bool> Lower { get; private set; } = null!;

        /// <summary>
        /// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        /// </summary>
        [Output("minLower")]
        public Output<int> MinLower { get; private set; } = null!;

        /// <summary>
        /// Minimum number of numeric characters in the result. Default value is `0`.
        /// </summary>
        [Output("minNumeric")]
        public Output<int> MinNumeric { get; private set; } = null!;

        /// <summary>
        /// Minimum number of special characters in the result. Default value is `0`.
        /// </summary>
        [Output("minSpecial")]
        public Output<int> MinSpecial { get; private set; } = null!;

        /// <summary>
        /// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        /// </summary>
        [Output("minUpper")]
        public Output<int> MinUpper { get; private set; } = null!;

        /// <summary>
        /// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        /// </summary>
        [Output("number")]
        public Output<bool> Number { get; private set; } = null!;

        /// <summary>
        /// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        /// </summary>
        [Output("numeric")]
        public Output<bool> Numeric { get; private set; } = null!;

        /// <summary>
        /// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        /// </summary>
        [Output("overrideSpecial")]
        public Output<string?> OverrideSpecial { get; private set; } = null!;

        /// <summary>
        /// The generated random string.
        /// </summary>
        [Output("result")]
        public Output<string> Result { get; private set; } = null!;

        /// <summary>
        /// Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
        /// </summary>
        [Output("special")]
        public Output<bool> Special { get; private set; } = null!;

        /// <summary>
        /// Include uppercase alphabet characters in the result. Default value is `true`.
        /// </summary>
        [Output("upper")]
        public Output<bool> Upper { get; private set; } = null!;


        /// <summary>
        /// Create a RandomString resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomString(string name, RandomStringArgs args, CustomResourceOptions? options = null)
            : base("random:index/randomString:RandomString", name, args ?? new RandomStringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomString(string name, Input<string> id, RandomStringState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomString:RandomString", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RandomString resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomString Get(string name, Input<string> id, RandomStringState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomString(name, id, state, options);
        }
    }

    public sealed class RandomStringArgs : global::Pulumi.ResourceArgs
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
        /// The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        /// </summary>
        [Input("length", required: true)]
        public Input<int> Length { get; set; } = null!;

        /// <summary>
        /// Include lowercase alphabet characters in the result. Default value is `true`.
        /// </summary>
        [Input("lower")]
        public Input<bool>? Lower { get; set; }

        /// <summary>
        /// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        /// </summary>
        [Input("minLower")]
        public Input<int>? MinLower { get; set; }

        /// <summary>
        /// Minimum number of numeric characters in the result. Default value is `0`.
        /// </summary>
        [Input("minNumeric")]
        public Input<int>? MinNumeric { get; set; }

        /// <summary>
        /// Minimum number of special characters in the result. Default value is `0`.
        /// </summary>
        [Input("minSpecial")]
        public Input<int>? MinSpecial { get; set; }

        /// <summary>
        /// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        /// </summary>
        [Input("minUpper")]
        public Input<int>? MinUpper { get; set; }

        /// <summary>
        /// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        /// </summary>
        [Input("number")]
        public Input<bool>? Number { get; set; }

        /// <summary>
        /// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        /// </summary>
        [Input("numeric")]
        public Input<bool>? Numeric { get; set; }

        /// <summary>
        /// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        /// </summary>
        [Input("overrideSpecial")]
        public Input<string>? OverrideSpecial { get; set; }

        /// <summary>
        /// Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
        /// </summary>
        [Input("special")]
        public Input<bool>? Special { get; set; }

        /// <summary>
        /// Include uppercase alphabet characters in the result. Default value is `true`.
        /// </summary>
        [Input("upper")]
        public Input<bool>? Upper { get; set; }

        public RandomStringArgs()
        {
        }
        public static new RandomStringArgs Empty => new RandomStringArgs();
    }

    public sealed class RandomStringState : global::Pulumi.ResourceArgs
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
        /// The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// Include lowercase alphabet characters in the result. Default value is `true`.
        /// </summary>
        [Input("lower")]
        public Input<bool>? Lower { get; set; }

        /// <summary>
        /// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        /// </summary>
        [Input("minLower")]
        public Input<int>? MinLower { get; set; }

        /// <summary>
        /// Minimum number of numeric characters in the result. Default value is `0`.
        /// </summary>
        [Input("minNumeric")]
        public Input<int>? MinNumeric { get; set; }

        /// <summary>
        /// Minimum number of special characters in the result. Default value is `0`.
        /// </summary>
        [Input("minSpecial")]
        public Input<int>? MinSpecial { get; set; }

        /// <summary>
        /// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        /// </summary>
        [Input("minUpper")]
        public Input<int>? MinUpper { get; set; }

        /// <summary>
        /// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        /// </summary>
        [Input("number")]
        public Input<bool>? Number { get; set; }

        /// <summary>
        /// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        /// </summary>
        [Input("numeric")]
        public Input<bool>? Numeric { get; set; }

        /// <summary>
        /// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        /// </summary>
        [Input("overrideSpecial")]
        public Input<string>? OverrideSpecial { get; set; }

        /// <summary>
        /// The generated random string.
        /// </summary>
        [Input("result")]
        public Input<string>? Result { get; set; }

        /// <summary>
        /// Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
        /// </summary>
        [Input("special")]
        public Input<bool>? Special { get; set; }

        /// <summary>
        /// Include uppercase alphabet characters in the result. Default value is `true`.
        /// </summary>
        [Input("upper")]
        public Input<bool>? Upper { get; set; }

        public RandomStringState()
        {
        }
        public static new RandomStringState Empty => new RandomStringState();
    }
}
