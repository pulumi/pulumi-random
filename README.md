[![Build Status](https://travis-ci.com/pulumi/pulumi-random.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-random)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Frandom.svg)](https://npmjs.com/package/@pulumi/random)
[![Python version](https://badge.fury.io/py/pulumi-random.svg)](https://pypi.org/project/pulumi-random)
[![GoDoc](https://godoc.org/github.com/pulumi/pulumi-random?status.svg)](https://godoc.org/github.com/pulumi/pulumi-random)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-random/blob/master/LICENSE)

# Random Provider

The random provider allows the safe use of randomness in a Pulumi program. This allows you to generate resource
properties, such as names, that contain randomness in a way that works with Pulumi's goal state oriented approach.
Using randomness as usual would not work well with Pulumi, because by definition, each time the program is evaluated,
a new random state would be produced, necessitating re-convergence on the goal state. This provider understands
how to work with the Pulumi resource lifecycle to accomplish randomness safely and in a way that works as desired.

## Example

For example, to generate a random password, allocate a `RandomPassword` resource
and then use its `result` output property (of type `Output<string>`) to pass
to another resource.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as random from "@pulumi/random";

const password = new random.RandomPassword("password", {
    length: 16,
    overrideSpecial: "_%@",
    special: true,
});
const example = new aws.rds.Instance("example", {
    password: password.result,
});
```

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/random

or `yarn`:

    $ yarn add @pulumi/random

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_random

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-random/sdk/go/...

## Reference

For detailed reference documentation, please visit [the API docs](
https://pulumi.io/reference/pkg/nodejs/@pulumi/random/index.html).
