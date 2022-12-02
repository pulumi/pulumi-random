// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.random.RandomUuidArgs;
import com.pulumi.random.Utilities;
import com.pulumi.random.inputs.RandomUuidState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The resource `random.RandomUuid` generates random uuid string that is intended to be used as unique identifiers for other resources.
 * 
 * This resource uses [hashicorp/go-uuid](https://github.com/hashicorp/go-uuid) to generate a UUID-formatted string for use with services needed a unique string identifier.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomUuid;
 * import com.pulumi.azure.core.ResourceGroup;
 * import com.pulumi.azure.core.ResourceGroupArgs;
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
 *         var testRandomUuid = new RandomUuid(&#34;testRandomUuid&#34;);
 * 
 *         var testResourceGroup = new ResourceGroup(&#34;testResourceGroup&#34;, ResourceGroupArgs.builder()        
 *             .location(&#34;Central US&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Random UUID&#39;s can be imported. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs.
 * 
 * ```sh
 *  $ pulumi import random:index/randomUuid:RandomUuid main aabbccdd-eeff-0011-2233-445566778899
 * ```
 * 
 */
@ResourceType(type="random:index/randomUuid:RandomUuid")
public class RandomUuid extends com.pulumi.resources.CustomResource {
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
     * The generated uuid presented in string format.
     * 
     */
    @Export(name="result", type=String.class, parameters={})
    private Output<String> result;

    /**
     * @return The generated uuid presented in string format.
     * 
     */
    public Output<String> result() {
        return this.result;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RandomUuid(String name) {
        this(name, RandomUuidArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RandomUuid(String name, @Nullable RandomUuidArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RandomUuid(String name, @Nullable RandomUuidArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomUuid:RandomUuid", name, args == null ? RandomUuidArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RandomUuid(String name, Output<String> id, @Nullable RandomUuidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomUuid:RandomUuid", name, state, makeResourceOptions(options, id));
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
    public static RandomUuid get(String name, Output<String> id, @Nullable RandomUuidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RandomUuid(name, id, state, options);
    }
}
