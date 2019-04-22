## 0.5.3 (Unreleased)

## 0.5.2 (Released April 22nd, 2019)

### Improvements

- Updated to v2.1.1 of the upstream Terraform Random provider.

## 0.5.1 (Released March 6th, 2019)

### Improvements

- Fix an issue where the Python `pulumi_random` package was depending on an older `pulumi` package.

## 0.5.0 (Released March 5th, 2019)

### Important

Updating to v0.17.0 version of `@pulumi/pulumi`.  This is an update that will not play nicely
in side-by-side applications that pull in prior versions of this package.

See https://github.com/pulumi/pulumi/commit/7f5e089f043a70c02f7e03600d6404ff0e27cc9d for more details.

As such, we are rev'ing the minor version of the package from 0.4 to 0.5.  Recent version of `pulumi` will now detect, and warn, if different versions of `@pulumi/pulumi` are loaded into the same application.  If you encounter this warning, it is recommended you move to versions of the `@pulumi/...` packages that are compatible.  i.e. keep everything on 0.16.x until you are ready to move everything to 0.17.x.

## 0.4.0 (Released February 12th, 2019)

### Improvements

- Support for the `deleteBeforeReplace` resource option and improved
  delete-before-replace behaviour introduced in [Pulumi
  0.16.14](https://github.com/pulumi/pulumi/blob/master/CHANGELOG.md#01614-released-january-31st-2019).

## 0.3.0 (Released January 19th, 2019)

### Improvements

- Documentation comments for the Node.js SDK now include examples
