// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.random.RandomIntegerArgs;
import com.pulumi.random.Utilities;
import com.pulumi.random.inputs.RandomIntegerState;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The resource `random.RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.
 * 
 * This resource can be used in conjunction with resources that have the `create_before_destroy` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.aws.albListenerRule;
 * import com.pulumi.aws.AlbListenerRuleArgs;
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
 *         // The following example shows how to generate a random priority
 *         // between 1 and 50000 for a aws_alb_listener_rule resource:
 *         var priority = new RandomInteger(&#34;priority&#34;, RandomIntegerArgs.builder()        
 *             .min(1)
 *             .max(50000)
 *             .keepers(Map.of(&#34;listener_arn&#34;, listenerArn))
 *             .build());
 * 
 *         var main = new AlbListenerRule(&#34;main&#34;, AlbListenerRuleArgs.builder()        
 *             .listenerArn(priority.keepers().applyValue(keepers -&gt; keepers.listenerArn()))
 *             .priority(priority.result())
 *             .action(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Random integers can be imported using the result, min, and max, with an
 * 
 * optional seed. This can be used to replace a config value with a value
 * 
 * interpolated from the random provider without experiencing diffs.
 * 
 * Example (values are separated by a ,):
 * 
 * ```sh
 * $ pulumi import random:index/randomInteger:RandomInteger priority 15390,1,50000
 * ```
 * 
 */
@ResourceType(type="random:index/randomInteger:RandomInteger")
public class RandomInteger extends com.pulumi.resources.CustomResource {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    @Export(name="keepers", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    public Output<Optional<Map<String,String>>> keepers() {
        return Codegen.optional(this.keepers);
    }
    /**
     * The maximum inclusive value of the range.
     * 
     */
    @Export(name="max", refs={Integer.class}, tree="[0]")
    private Output<Integer> max;

    /**
     * @return The maximum inclusive value of the range.
     * 
     */
    public Output<Integer> max() {
        return this.max;
    }
    /**
     * The minimum inclusive value of the range.
     * 
     */
    @Export(name="min", refs={Integer.class}, tree="[0]")
    private Output<Integer> min;

    /**
     * @return The minimum inclusive value of the range.
     * 
     */
    public Output<Integer> min() {
        return this.min;
    }
    /**
     * The random integer result.
     * 
     */
    @Export(name="result", refs={Integer.class}, tree="[0]")
    private Output<Integer> result;

    /**
     * @return The random integer result.
     * 
     */
    public Output<Integer> result() {
        return this.result;
    }
    /**
     * A custom seed to always produce the same value.
     * 
     */
    @Export(name="seed", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> seed;

    /**
     * @return A custom seed to always produce the same value.
     * 
     */
    public Output<Optional<String>> seed() {
        return Codegen.optional(this.seed);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RandomInteger(String name) {
        this(name, RandomIntegerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RandomInteger(String name, RandomIntegerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RandomInteger(String name, RandomIntegerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomInteger:RandomInteger", name, args == null ? RandomIntegerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RandomInteger(String name, Output<String> id, @Nullable RandomIntegerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomInteger:RandomInteger", name, state, makeResourceOptions(options, id));
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
    public static RandomInteger get(String name, Output<String> id, @Nullable RandomIntegerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RandomInteger(name, id, state, options);
    }
}
