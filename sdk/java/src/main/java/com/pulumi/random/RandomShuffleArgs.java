// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RandomShuffleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RandomShuffleArgs Empty = new RandomShuffleArgs();

    /**
     * The list of strings to shuffle.
     * 
     */
    @Import(name="inputs", required=true)
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
     * The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     * 
     */
    @Import(name="resultCount")
    private @Nullable Output<Integer> resultCount;

    /**
     * @return The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     * 
     */
    public Optional<Output<Integer>> resultCount() {
        return Optional.ofNullable(this.resultCount);
    }

    /**
     * Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     * 
     */
    @Import(name="seed")
    private @Nullable Output<String> seed;

    /**
     * @return Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     * 
     */
    public Optional<Output<String>> seed() {
        return Optional.ofNullable(this.seed);
    }

    private RandomShuffleArgs() {}

    private RandomShuffleArgs(RandomShuffleArgs $) {
        this.inputs = $.inputs;
        this.keepers = $.keepers;
        this.resultCount = $.resultCount;
        this.seed = $.seed;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RandomShuffleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RandomShuffleArgs $;

        public Builder() {
            $ = new RandomShuffleArgs();
        }

        public Builder(RandomShuffleArgs defaults) {
            $ = new RandomShuffleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param inputs The list of strings to shuffle.
         * 
         * @return builder
         * 
         */
        public Builder inputs(Output<List<String>> inputs) {
            $.inputs = inputs;
            return this;
        }

        /**
         * @param inputs The list of strings to shuffle.
         * 
         * @return builder
         * 
         */
        public Builder inputs(List<String> inputs) {
            return inputs(Output.of(inputs));
        }

        /**
         * @param inputs The list of strings to shuffle.
         * 
         * @return builder
         * 
         */
        public Builder inputs(String... inputs) {
            return inputs(List.of(inputs));
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
         * @param resultCount The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
         * 
         * @return builder
         * 
         */
        public Builder resultCount(@Nullable Output<Integer> resultCount) {
            $.resultCount = resultCount;
            return this;
        }

        /**
         * @param resultCount The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
         * 
         * @return builder
         * 
         */
        public Builder resultCount(Integer resultCount) {
            return resultCount(Output.of(resultCount));
        }

        /**
         * @param seed Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
         * 
         * @return builder
         * 
         */
        public Builder seed(@Nullable Output<String> seed) {
            $.seed = seed;
            return this;
        }

        /**
         * @param seed Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
         * 
         * @return builder
         * 
         */
        public Builder seed(String seed) {
            return seed(Output.of(seed));
        }

        public RandomShuffleArgs build() {
            $.inputs = Objects.requireNonNull($.inputs, "expected parameter 'inputs' to be non-null");
            return $;
        }
    }

}
