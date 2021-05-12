// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as random from "@pulumi/random";

export const randomId = new random.RandomId("id", { byteLength: 4 }).hex;
export const randomShuffle = new random.RandomShuffle("shuffle", { inputs: [ "a", "b", "c" ] }).results;
export const randomString = new random.RandomString("string", { length: 32 }).result;
export const randomInteger = new random.RandomInteger("integer", { min: 128, max: 1024 }).result;
export const randomUuid = new random.RandomUuid("uuid").result;
export const randomPassword = new random.RandomPassword("password", { length: 32 }).result;
