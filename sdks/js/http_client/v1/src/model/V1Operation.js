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
    define(['ApiClient', 'model/V1Cache', 'model/V1Component', 'model/V1Param', 'model/V1Plugins', 'model/V1Termination', 'model/V1TriggerPolicy'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1Cache'), require('./V1Component'), require('./V1Param'), require('./V1Plugins'), require('./V1Termination'), require('./V1TriggerPolicy'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1Operation = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1Cache, root.PolyaxonSdk.V1Component, root.PolyaxonSdk.V1Param, root.PolyaxonSdk.V1Plugins, root.PolyaxonSdk.V1Termination, root.PolyaxonSdk.V1TriggerPolicy);
  }
}(this, function(ApiClient, V1Cache, V1Component, V1Param, V1Plugins, V1Termination, V1TriggerPolicy) {
  'use strict';

  /**
   * The V1Operation model module.
   * @module model/V1Operation
   * @version 1.0.5
   */

  /**
   * Constructs a new <code>V1Operation</code>.
   * @alias module:model/V1Operation
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1Operation</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1Operation} obj Optional instance to populate.
   * @return {module:model/V1Operation} The populated <code>V1Operation</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('version'))
        obj.version = ApiClient.convertToType(data['version'], 'Number');
      if (data.hasOwnProperty('kind'))
        obj.kind = ApiClient.convertToType(data['kind'], 'String');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('tag'))
        obj.tag = ApiClient.convertToType(data['tag'], 'String');
      if (data.hasOwnProperty('description'))
        obj.description = ApiClient.convertToType(data['description'], 'String');
      if (data.hasOwnProperty('tags'))
        obj.tags = ApiClient.convertToType(data['tags'], ['String']);
      if (data.hasOwnProperty('profile'))
        obj.profile = ApiClient.convertToType(data['profile'], 'String');
      if (data.hasOwnProperty('queue'))
        obj.queue = ApiClient.convertToType(data['queue'], 'String');
      if (data.hasOwnProperty('cache'))
        obj.cache = V1Cache.constructFromObject(data['cache']);
      if (data.hasOwnProperty('schedule'))
        obj.schedule = ApiClient.convertToType(data['schedule'], Object);
      if (data.hasOwnProperty('parallel'))
        obj.parallel = ApiClient.convertToType(data['parallel'], Object);
      if (data.hasOwnProperty('dependencies'))
        obj.dependencies = ApiClient.convertToType(data['dependencies'], ['String']);
      if (data.hasOwnProperty('trigger'))
        obj.trigger = V1TriggerPolicy.constructFromObject(data['trigger']);
      if (data.hasOwnProperty('conditions'))
        obj.conditions = ApiClient.convertToType(data['conditions'], [Object]);
      if (data.hasOwnProperty('skip_on_upstream_skip'))
        obj.skip_on_upstream_skip = ApiClient.convertToType(data['skip_on_upstream_skip'], 'Boolean');
      if (data.hasOwnProperty('termination'))
        obj.termination = V1Termination.constructFromObject(data['termination']);
      if (data.hasOwnProperty('plugins'))
        obj.plugins = V1Plugins.constructFromObject(data['plugins']);
      if (data.hasOwnProperty('params'))
        obj.params = ApiClient.convertToType(data['params'], {'String': V1Param});
      if (data.hasOwnProperty('run_patch'))
        obj.run_patch = ApiClient.convertToType(data['run_patch'], Object);
      if (data.hasOwnProperty('dag_ref'))
        obj.dag_ref = ApiClient.convertToType(data['dag_ref'], 'String');
      if (data.hasOwnProperty('url_ref'))
        obj.url_ref = ApiClient.convertToType(data['url_ref'], 'String');
      if (data.hasOwnProperty('path_ref'))
        obj.path_ref = ApiClient.convertToType(data['path_ref'], 'String');
      if (data.hasOwnProperty('hub_ref'))
        obj.hub_ref = ApiClient.convertToType(data['hub_ref'], 'String');
      if (data.hasOwnProperty('component'))
        obj.component = V1Component.constructFromObject(data['component']);
    }
    return obj;
  }

  /**
   * @member {Number} version
   */
  exports.prototype.version = undefined;

  /**
   * @member {String} kind
   */
  exports.prototype.kind = undefined;

  /**
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * @member {String} tag
   */
  exports.prototype.tag = undefined;

  /**
   * @member {String} description
   */
  exports.prototype.description = undefined;

  /**
   * @member {Array.<String>} tags
   */
  exports.prototype.tags = undefined;

  /**
   * @member {String} profile
   */
  exports.prototype.profile = undefined;

  /**
   * @member {String} queue
   */
  exports.prototype.queue = undefined;

  /**
   * @member {module:model/V1Cache} cache
   */
  exports.prototype.cache = undefined;

  /**
   * @member {Object} schedule
   */
  exports.prototype.schedule = undefined;

  /**
   * @member {Object} parallel
   */
  exports.prototype.parallel = undefined;

  /**
   * @member {Array.<String>} dependencies
   */
  exports.prototype.dependencies = undefined;

  /**
   * @member {module:model/V1TriggerPolicy} trigger
   */
  exports.prototype.trigger = undefined;

  /**
   * @member {Array.<Object>} conditions
   */
  exports.prototype.conditions = undefined;

  /**
   * @member {Boolean} skip_on_upstream_skip
   */
  exports.prototype.skip_on_upstream_skip = undefined;

  /**
   * @member {module:model/V1Termination} termination
   */
  exports.prototype.termination = undefined;

  /**
   * @member {module:model/V1Plugins} plugins
   */
  exports.prototype.plugins = undefined;

  /**
   * @member {Object.<String, module:model/V1Param>} params
   */
  exports.prototype.params = undefined;

  /**
   * @member {Object} run_patch
   */
  exports.prototype.run_patch = undefined;

  /**
   * @member {String} dag_ref
   */
  exports.prototype.dag_ref = undefined;

  /**
   * @member {String} url_ref
   */
  exports.prototype.url_ref = undefined;

  /**
   * @member {String} path_ref
   */
  exports.prototype.path_ref = undefined;

  /**
   * @member {String} hub_ref
   */
  exports.prototype.hub_ref = undefined;

  /**
   * @member {module:model/V1Component} component
   */
  exports.prototype.component = undefined;

  return exports;

}));
