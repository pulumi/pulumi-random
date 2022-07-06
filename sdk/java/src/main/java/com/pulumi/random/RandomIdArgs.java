// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RandomIdArgs extends com.pulumi.resources.ResourceArgs {

    public static final RandomIdArgs Empty = new RandomIdArgs();

    /**
     * The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
     * 
     */
    @Import(name="byteLength", required=true)
    private Output<Integer> byteLength;

    /**
     * @return The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
     * 
     */
    public Output<Integer> byteLength() {
        return this.byteLength;
    }

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
     * documentation](../index.html) for more information.
     * 
     */
    @Import(name="keepers")
    private @Nullable Output<Map<String,Object>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
     * documentation](../index.html) for more information.
     * 
     */
    public Optional<Output<Map<String,Object>>> keepers() {
        return Optional.ofNullable(this.keepers);
    }

    /**
     * Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be
     * URL-safe or base64 encoded.
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be
     * URL-safe or base64 encoded.
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    private RandomIdArgs() {}

    private RandomIdArgs(RandomIdArgs $) {
        this.byteLength = $.byteLength;
        this.keepers = $.keepers;
        this.prefix = $.prefix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RandomIdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RandomIdArgs $;

        public Builder() {
            $ = new RandomIdArgs();
        }

        public Builder(RandomIdArgs defaults) {
            $ = new RandomIdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param byteLength The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
         * 
         * @return builder
         * 
         */
        public Builder byteLength(Output<Integer> byteLength) {
            $.byteLength = byteLength;
            return this;
        }

        /**
         * @param byteLength The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
         * 
         * @return builder
         * 
         */
        public Builder byteLength(Integer byteLength) {
            return byteLength(Output.of(byteLength));
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
         * documentation](../index.html) for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(@Nullable Output<Map<String,Object>> keepers) {
            $.keepers = keepers;
            return this;
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
         * documentation](../index.html) for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(Map<String,Object> keepers) {
            return keepers(Output.of(keepers));
        }

        /**
         * @param prefix Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be
         * URL-safe or base64 encoded.
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be
         * URL-safe or base64 encoded.
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        public RandomIdArgs build() {
            $.byteLength = Objects.requireNonNull($.byteLength, "expected parameter 'byteLength' to be non-null");
            return $;
        }
    }

}
