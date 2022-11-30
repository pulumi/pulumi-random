// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.random.RandomPasswordArgs;
import com.pulumi.random.Utilities;
import com.pulumi.random.inputs.RandomPasswordState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **Note:** Requires random provider version &gt;= 2.2.0
 * 
 * Identical to random.RandomString with the exception that the
 * result is treated as sensitive and, thus, _not_ displayed in console output.
 * 
 * &gt; **Note:** All attributes including the generated password will be stored in
 * the raw state as plain-text. [Read more about sensitive data in
 * state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * This resource *does* use a cryptographic random number generator.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomPassword;
 * import com.pulumi.random.RandomPasswordArgs;
 * import com.pulumi.aws.rds.Instance;
 * import com.pulumi.aws.rds.InstanceArgs;
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
 *         var password = new RandomPassword(&#34;password&#34;, RandomPasswordArgs.builder()        
 *             .length(16)
 *             .special(true)
 *             .overrideSpecial(&#34;_%@&#34;)
 *             .build());
 * 
 *         var example = new Instance(&#34;example&#34;, InstanceArgs.builder()        
 *             .instanceClass(&#34;db.t3.micro&#34;)
 *             .allocatedStorage(64)
 *             .engine(&#34;mysql&#34;)
 *             .username(&#34;someone&#34;)
 *             .password(password.result())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Random Password can be imported by specifying the value of the string
 * 
 * ```sh
 *  $ pulumi import random:index/randomPassword:RandomPassword password securepassword
 * ```
 * 
 */
@ResourceType(type="random:index/randomPassword:RandomPassword")
public class RandomPassword extends com.pulumi.resources.CustomResource {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
     * documentation](../index.html) for more information.
     * 
     */
    @Export(name="keepers", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
     * documentation](../index.html) for more information.
     * 
     */
    public Output<Optional<Map<String,Object>>> keepers() {
        return Codegen.optional(this.keepers);
    }
    /**
     * The length of the string desired.
     * 
     */
    @Export(name="length", type=Integer.class, parameters={})
    private Output<Integer> length;

    /**
     * @return The length of the string desired.
     * 
     */
    public Output<Integer> length() {
        return this.length;
    }
    /**
     * Include lowercase alphabet characters in the result.
     * 
     */
    @Export(name="lower", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> lower;

    /**
     * @return Include lowercase alphabet characters in the result.
     * 
     */
    public Output<Optional<Boolean>> lower() {
        return Codegen.optional(this.lower);
    }
    /**
     * Minimum number of lowercase alphabet characters in the result.
     * 
     */
    @Export(name="minLower", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minLower;

    /**
     * @return Minimum number of lowercase alphabet characters in the result.
     * 
     */
    public Output<Optional<Integer>> minLower() {
        return Codegen.optional(this.minLower);
    }
    /**
     * Minimum number of numeric characters in the result.
     * 
     */
    @Export(name="minNumeric", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minNumeric;

    /**
     * @return Minimum number of numeric characters in the result.
     * 
     */
    public Output<Optional<Integer>> minNumeric() {
        return Codegen.optional(this.minNumeric);
    }
    /**
     * Minimum number of special characters in the result.
     * 
     */
    @Export(name="minSpecial", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minSpecial;

    /**
     * @return Minimum number of special characters in the result.
     * 
     */
    public Output<Optional<Integer>> minSpecial() {
        return Codegen.optional(this.minSpecial);
    }
    /**
     * Minimum number of uppercase alphabet characters in the result.
     * 
     */
    @Export(name="minUpper", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minUpper;

    /**
     * @return Minimum number of uppercase alphabet characters in the result.
     * 
     */
    public Output<Optional<Integer>> minUpper() {
        return Codegen.optional(this.minUpper);
    }
    /**
     * Include numeric characters in the result.
     * 
     */
    @Export(name="number", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> number;

    /**
     * @return Include numeric characters in the result.
     * 
     */
    public Output<Optional<Boolean>> number() {
        return Codegen.optional(this.number);
    }
    /**
     * Supply your own list of special characters to use for string generation. This overrides the default character list in
     * the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
     * generation.
     * 
     */
    @Export(name="overrideSpecial", type=String.class, parameters={})
    private Output</* @Nullable */ String> overrideSpecial;

    /**
     * @return Supply your own list of special characters to use for string generation. This overrides the default character list in
     * the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
     * generation.
     * 
     */
    public Output<Optional<String>> overrideSpecial() {
        return Codegen.optional(this.overrideSpecial);
    }
    /**
     * The generated random string.
     * 
     */
    @Export(name="result", type=String.class, parameters={})
    private Output<String> result;

    /**
     * @return The generated random string.
     * 
     */
    public Output<String> result() {
        return this.result;
    }
    /**
     * Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`
     * 
     */
    @Export(name="special", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> special;

    /**
     * @return Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`
     * 
     */
    public Output<Optional<Boolean>> special() {
        return Codegen.optional(this.special);
    }
    /**
     * Include uppercase alphabet characters in the result.
     * 
     */
    @Export(name="upper", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> upper;

    /**
     * @return Include uppercase alphabet characters in the result.
     * 
     */
    public Output<Optional<Boolean>> upper() {
        return Codegen.optional(this.upper);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RandomPassword(String name) {
        this(name, RandomPasswordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RandomPassword(String name, RandomPasswordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RandomPassword(String name, RandomPasswordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomPassword:RandomPassword", name, args == null ? RandomPasswordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RandomPassword(String name, Output<String> id, @Nullable RandomPasswordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomPassword:RandomPassword", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "result"
            ))
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
    public static RandomPassword get(String name, Output<String> id, @Nullable RandomPasswordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RandomPassword(name, id, state, options);
    }
}
