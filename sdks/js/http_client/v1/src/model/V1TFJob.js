// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1CleanPodPolicy', 'model/V1KFReplica'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1CleanPodPolicy'), require('./V1KFReplica'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1TFJob = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1CleanPodPolicy, root.PolyaxonSdk.V1KFReplica);
  }
}(this, function(ApiClient, V1CleanPodPolicy, V1KFReplica) {
  'use strict';

  /**
   * The V1TFJob model module.
   * @module model/V1TFJob
   * @version 1.0.5
   */

  /**
   * Constructs a new <code>V1TFJob</code>.
   * @alias module:model/V1TFJob
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1TFJob</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1TFJob} obj Optional instance to populate.
   * @return {module:model/V1TFJob} The populated <code>V1TFJob</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('kind'))
        obj.kind = ApiClient.convertToType(data['kind'], 'String');
      if (data.hasOwnProperty('cleanPodPolicy'))
        obj.cleanPodPolicy = V1CleanPodPolicy.constructFromObject(data['cleanPodPolicy']);
      if (data.hasOwnProperty('chief'))
        obj.chief = V1KFReplica.constructFromObject(data['chief']);
      if (data.hasOwnProperty('ps'))
        obj.ps = V1KFReplica.constructFromObject(data['ps']);
      if (data.hasOwnProperty('worker'))
        obj.worker = V1KFReplica.constructFromObject(data['worker']);
      if (data.hasOwnProperty('evaluator'))
        obj.evaluator = V1KFReplica.constructFromObject(data['evaluator']);
    }
    return obj;
  }

  /**
   * @member {String} kind
   * @default 'tfjob'
   */
  exports.prototype.kind = 'tfjob';

  /**
   * @member {module:model/V1CleanPodPolicy} cleanPodPolicy
   */
  exports.prototype.cleanPodPolicy = undefined;

  /**
   * @member {module:model/V1KFReplica} chief
   */
  exports.prototype.chief = undefined;

  /**
   * @member {module:model/V1KFReplica} ps
   */
  exports.prototype.ps = undefined;

  /**
   * @member {module:model/V1KFReplica} worker
   */
  exports.prototype.worker = undefined;

  /**
   * @member {module:model/V1KFReplica} evaluator
   */
  exports.prototype.evaluator = undefined;

  return exports;

}));
