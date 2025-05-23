// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.random.RandomIdArgs;
import com.pulumi.random.Utilities;
import com.pulumi.random.inputs.RandomIdState;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The resource `random.RandomId` generates random numbers that are intended to be
 * used as unique identifiers for other resources. If the output is considered
 * sensitive, and should not be displayed in the CLI, use `random.RandomBytes`
 * instead.
 * 
 * This resource *does* use a cryptographic random number generator in order
 * to minimize the chance of collisions, making the results of this resource
 * when a 16-byte identifier is requested of equivalent uniqueness to a
 * type-4 UUID.
 * 
 * This resource can be used in conjunction with resources that have
 * the `create_before_destroy` lifecycle flag set to avoid conflicts with
 * unique names during the brief period where both the old and new resources
 * exist concurrently.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomId;
 * import com.pulumi.random.RandomIdArgs;
 * import com.pulumi.aws.ec2.Instance;
 * import com.pulumi.aws.ec2.InstanceArgs;
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
 *         // The following example shows how to generate a unique name for an AWS EC2
 *         // instance that changes each time a new AMI id is selected.
 *         var server = new RandomId("server", RandomIdArgs.builder()
 *             .keepers(Map.of("ami_id", amiId))
 *             .byteLength(8)
 *             .build());
 * 
 *         var serverInstance = new Instance("serverInstance", InstanceArgs.builder()
 *             .tags(Map.of("Name", server.hex().applyValue(_hex -> String.format("web-server %s", _hex))))
 *             .ami(server.keepers().applyValue(_keepers -> _keepers.amiId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Random IDs can be imported using the b64_url with an optional prefix. This
 * 
 * can be used to replace a config value with a value interpolated from the
 * 
 * random provider without experiencing diffs.
 * 
 * Example with no prefix:
 * 
 * ```sh
 * $ pulumi import random:index/randomId:RandomId server p-9hUg
 * ```
 * 
 * Example with prefix (prefix is separated by a ,):
 * 
 * ```sh
 * $ pulumi import random:index/randomId:RandomId server my-prefix-,p-9hUg
 * ```
 * 
 */
@ResourceType(type="random:index/randomId:RandomId")
public class RandomId extends com.pulumi.resources.CustomResource {
    /**
     * The generated id presented in base64 without additional transformations.
     * 
     */
    @Export(name="b64Std", refs={String.class}, tree="[0]")
    private Output<String> b64Std;

    /**
     * @return The generated id presented in base64 without additional transformations.
     * 
     */
    public Output<String> b64Std() {
        return this.b64Std;
    }
    /**
     * The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
     * 
     */
    @Export(name="b64Url", refs={String.class}, tree="[0]")
    private Output<String> b64Url;

    /**
     * @return The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
     * 
     */
    public Output<String> b64Url() {
        return this.b64Url;
    }
    /**
     * The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
     * 
     */
    @Export(name="byteLength", refs={Integer.class}, tree="[0]")
    private Output<Integer> byteLength;

    /**
     * @return The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
     * 
     */
    public Output<Integer> byteLength() {
        return this.byteLength;
    }
    /**
     * The generated id presented in non-padded decimal digits.
     * 
     */
    @Export(name="dec", refs={String.class}, tree="[0]")
    private Output<String> dec;

    /**
     * @return The generated id presented in non-padded decimal digits.
     * 
     */
    public Output<String> dec() {
        return this.dec;
    }
    /**
     * The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
     * 
     */
    @Export(name="hex", refs={String.class}, tree="[0]")
    private Output<String> hex;

    /**
     * @return The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
     * 
     */
    public Output<String> hex() {
        return this.hex;
    }
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
     * Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be URL-safe or base64 encoded.
     * 
     */
    @Export(name="prefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> prefix;

    /**
     * @return Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be URL-safe or base64 encoded.
     * 
     */
    public Output<Optional<String>> prefix() {
        return Codegen.optional(this.prefix);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RandomId(java.lang.String name) {
        this(name, RandomIdArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RandomId(java.lang.String name, RandomIdArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RandomId(java.lang.String name, RandomIdArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomId:RandomId", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RandomId(java.lang.String name, Output<java.lang.String> id, @Nullable RandomIdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomId:RandomId", name, state, makeResourceOptions(options, id), false);
    }

    private static RandomIdArgs makeArgs(RandomIdArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RandomIdArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static RandomId get(java.lang.String name, Output<java.lang.String> id, @Nullable RandomIdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RandomId(name, id, state, options);
    }
}
