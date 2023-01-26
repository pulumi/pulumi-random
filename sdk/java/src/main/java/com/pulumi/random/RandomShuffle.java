// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.random.RandomShuffleArgs;
import com.pulumi.random.Utilities;
import com.pulumi.random.inputs.RandomShuffleState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The resource `random.RandomShuffle` generates a random permutation of a list of strings given as an argument.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomShuffle;
 * import com.pulumi.random.RandomShuffleArgs;
 * import com.pulumi.aws.elb.LoadBalancer;
 * import com.pulumi.aws.elb.LoadBalancerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var az = new RandomShuffle(&#34;az&#34;, RandomShuffleArgs.builder()        
 *             .inputs(            
 *                 &#34;us-west-1a&#34;,
 *                 &#34;us-west-1c&#34;,
 *                 &#34;us-west-1d&#34;,
 *                 &#34;us-west-1e&#34;)
 *             .resultCount(2)
 *             .build());
 * 
 *         var example = new LoadBalancer(&#34;example&#34;, LoadBalancerArgs.builder()        
 *             .availabilityZones(az.results())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="random:index/randomShuffle:RandomShuffle")
public class RandomShuffle extends com.pulumi.resources.CustomResource {
    /**
     * The list of strings to shuffle.
     * 
     */
    @Export(name="inputs", type=List.class, parameters={String.class})
    private Output<List<String>> inputs;

    /**
     * @return The list of strings to shuffle.
     * 
     */
    public Output<List<String>> inputs() {
        return this.inputs;
    }
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    @Export(name="keepers", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    public Output<Optional<Map<String,String>>> keepers() {
        return Codegen.optional(this.keepers);
    }
    /**
     * The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     * 
     */
    @Export(name="resultCount", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> resultCount;

    /**
     * @return The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     * 
     */
    public Output<Optional<Integer>> resultCount() {
        return Codegen.optional(this.resultCount);
    }
    /**
     * Random permutation of the list of strings given in `input`.
     * 
     */
    @Export(name="results", type=List.class, parameters={String.class})
    private Output<List<String>> results;

    /**
     * @return Random permutation of the list of strings given in `input`.
     * 
     */
    public Output<List<String>> results() {
        return this.results;
    }
    /**
     * Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     * 
     */
    @Export(name="seed", type=String.class, parameters={})
    private Output</* @Nullable */ String> seed;

    /**
     * @return Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     * 
     */
    public Output<Optional<String>> seed() {
        return Codegen.optional(this.seed);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RandomShuffle(String name) {
        this(name, RandomShuffleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RandomShuffle(String name, RandomShuffleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RandomShuffle(String name, RandomShuffleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomShuffle:RandomShuffle", name, args == null ? RandomShuffleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RandomShuffle(String name, Output<String> id, @Nullable RandomShuffleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomShuffle:RandomShuffle", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RandomShuffle get(String name, Output<String> id, @Nullable RandomShuffleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RandomShuffle(name, id, state, options);
    }
}
