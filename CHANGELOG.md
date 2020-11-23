CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to pulumi-terraform-bridge v2.13.2

---

## 2.4.2 (2020-11-23)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.4.1 (2020-11-05)
* Upgrade to pulumi-terraform-bridge v2.12.1

## 2.4.0 (2020-10-23)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.3.1 (2020-08-26)
* Upgrade to pulumi-terraform-bridge v2.7.3

## 2.3.0 (2020-08-24)
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python
* Upgrade to pulumi-terraform-bridge v2.7.2

## 2.2.0 (2020-07-07)
* Upgrade to v2.3.0 of the Random Terraform Provider

## 2.1.4 (2020-07-01)
* Publish Go examples in SDK

## 2.1.3 (2020-06-15)
* Switch to GitHub actions for build

## 2.1.2 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.1.1 (2020-05-11)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.0 (2020-04-28)
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-16)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.7.0 (2020-03-30)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4

## 1.6.0 (2020-03-13)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.5.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.4.0 (2019-12-12)
* Upgrade to pulumi-terraform-bridge v1.5.2

## 1.3.0 (2019-12-02)
* readme: fix example to create right passwd for RDS
[#64](https://github.com/pulumi/pulumi-random/pull/64)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform-bridge v1.4.3

## 1.2.0 (2019-11-07)
* Add a **preview** .NET SDK

## 1.1.0 (2019-10-02)
* Regenerate SDK against tf2pulumi 0.6.0
* Upgrade to v2.2.1 of the Random Terraform Provider

## 1.0.0 (2019-09-03)
* Use 1.0 version of Pulumi dependency

## 1.0.0-rc.1 (2019-08-28)
* Upgrade pulumi-terraform to 3f206601e7

## 1.0.0-beta.2 (2019-08-26)
___NULL___
* Upgrade to pulumi-terraform@58c7473d08

## 1.0.0-beta.1 (2019-08-13)
* Add `RandomPassword` which behaves like `RandomString` except the ID value stored in the checkpoint is always "none" instead of the underlying value

## 0.5.8 (2019-08-09)
* Upgrade to pulumi-terraform@9db2fc93cd
* Upgrade to v2.2.0 of the Random Terraform Provider

## 0.5.7 (2019-08-08)
* Update to pulumi-terraform@013b95b1c8

## 0.5.6 (2019-07-09)
* Fix detailed diffs with nested computed values.

## 0.5.5 (2019-07-08)
* Communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update

## 0.5.4 (2019-06-21)
* Add TypeScript type guards for each resource class ([7ace3e9b5f](https://github.com/pulumi/pulumi-terraform/commit/7ace3e9b5f2dcd4692b029ba4b80360582d7949a))
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.5.3 (2019-05-31)
* Update to v2.1.2 of the Terraform Random provider

## 0.5.2 (2019-04-22)
* Update to v2.1.1 of the Terraform Random provider

## 0.5.1 (2019-03-06)
* Depend on a more recent minimum version of the Pulumi Python SDK

## 0.5.0 (2019-03-05)
* Updated to the latest version of the `pulumi` SDK
* BREAKING: This version of the Random provider will not work side-by-side with previous versions

## 0.4.0 (2019-02-12)
* Add support for the `deleteBeforeReplace` resource option and improved delete-before-replace behaviour introduced in Pulumi v0.16.14

## 0.3.0 (2019-01-20)
* Add documentation comments to the Node.js SDK

## 0.1.0 (2018-02-02)
* Initial version of the Random provider

