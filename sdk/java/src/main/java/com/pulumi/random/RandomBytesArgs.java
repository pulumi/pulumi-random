// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RandomBytesArgs extends com.pulumi.resources.ResourceArgs {

    public static final RandomBytesArgs Empty = new RandomBytesArgs();

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    @Import(name="keepers")
    private @Nullable Output<Map<String,String>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    public Optional<Output<Map<String,String>>> keepers() {
        return Optional.ofNullable(this.keepers);
    }

    /**
     * The number of bytes requested. The minimum value for length is 1.
     * 
     */
    @Import(name="length", required=true)
    private Output<Integer> length;

    /**
     * @return The number of bytes requested. The minimum value for length is 1.
     * 
     */
    public Output<Integer> length() {
        return this.length;
    }

    private RandomBytesArgs() {}

    private RandomBytesArgs(RandomBytesArgs $) {
        this.keepers = $.keepers;
        this.length = $.length;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RandomBytesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RandomBytesArgs $;

        public Builder() {
            $ = new RandomBytesArgs();
        }

        public Builder(RandomBytesArgs defaults) {
            $ = new RandomBytesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(@Nullable Output<Map<String,String>> keepers) {
            $.keepers = keepers;
            return this;
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(Map<String,String> keepers) {
            return keepers(Output.of(keepers));
        }

        /**
         * @param length The number of bytes requested. The minimum value for length is 1.
         * 
         * @return builder
         * 
         */
        public Builder length(Output<Integer> length) {
            $.length = length;
            return this;
        }

        /**
         * @param length The number of bytes requested. The minimum value for length is 1.
         * 
         * @return builder
         * 
         */
        public Builder length(Integer length) {
            return length(Output.of(length));
        }

        public RandomBytesArgs build() {
            $.length = Objects.requireNonNull($.length, "expected parameter 'length' to be non-null");
            return $;
        }
    }

}
