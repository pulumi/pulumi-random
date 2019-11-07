CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

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
