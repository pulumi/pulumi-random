// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Random
{
    /// <summary>
    /// The resource `random..RandomPet` generates random pet names that are intended to be
    /// used as unique identifiers for other resources.
    /// 
    /// This resource can be used in conjunction with resources that have
    /// the `create_before_destroy` lifecycle flag set, to avoid conflicts with
    /// unique names during the brief period where both the old and new resources
    /// exist concurrently.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-random/blob/master/website/docs/r/pet.html.markdown.
    /// </summary>
    public partial class RandomPet : Pulumi.CustomResource
    {
        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, object>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// The length (in words) of the pet name.
        /// </summary>
        [Output("length")]
        public Output<int?> Length { get; private set; } = null!;

        /// <summary>
        /// A string to prefix the name with.
        /// </summary>
        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;

        /// <summary>
        /// The character to separate words in the pet name.
        /// </summary>
        [Output("separator")]
        public Output<string?> Separator { get; private set; } = null!;


        /// <summary>
        /// Create a RandomPet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomPet(string name, RandomPetArgs? args = null, CustomResourceOptions? options = null)
            : base("random:index/randomPet:RandomPet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private RandomPet(string name, Input<string> id, RandomPetState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomPet:RandomPet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RandomPet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomPet Get(string name, Input<string> id, RandomPetState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomPet(name, id, state, options);
        }
    }

    public sealed class RandomPetArgs : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<object>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        /// <summary>
        /// The length (in words) of the pet name.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// A string to prefix the name with.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The character to separate words in the pet name.
        /// </summary>
        [Input("separator")]
        public Input<string>? Separator { get; set; }

        public RandomPetArgs()
        {
        }
    }

    public sealed class RandomPetState : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<object>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        /// <summary>
        /// The length (in words) of the pet name.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// A string to prefix the name with.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The character to separate words in the pet name.
        /// </summary>
        [Input("separator")]
        public Input<string>? Separator { get; set; }

        public RandomPetState()
        {
        }
    }
}
