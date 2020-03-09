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
    define(['ApiClient', 'model/RuntimeError', 'model/V1HubModel', 'model/V1ListHubModelsResponse'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/RuntimeError'), require('../model/V1HubModel'), require('../model/V1ListHubModelsResponse'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.HubModelsV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.RuntimeError, root.PolyaxonSdk.V1HubModel, root.PolyaxonSdk.V1ListHubModelsResponse);
  }
}(this, function(ApiClient, RuntimeError, V1HubModel, V1ListHubModelsResponse) {
  'use strict';

  /**
   * HubModelsV1 service.
   * @module api/HubModelsV1Api
   * @version 1.0.5
   */

  /**
   * Constructs a new HubModelsV1Api. 
   * @alias module:api/HubModelsV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the createHubModel operation.
     * @callback module:api/HubModelsV1Api~createHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create hub model
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1HubModel} body Model body
     * @param {module:api/HubModelsV1Api~createHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    this.createHubModel = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createHubModel");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createHubModel");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1HubModel;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteHubModel operation.
     * @callback module:api/HubModelsV1Api~deleteHubModelCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete hub model
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/HubModelsV1Api~deleteHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteHubModel = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteHubModel");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteHubModel");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getHubModel operation.
     * @callback module:api/HubModelsV1Api~getHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get hub model
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/HubModelsV1Api~getHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    this.getHubModel = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getHubModel");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getHubModel");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1HubModel;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listHubModelNames operation.
     * @callback module:api/HubModelsV1Api~listHubModelNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListHubModelsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub model names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/HubModelsV1Api~listHubModelNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListHubModelsResponse}
     */
    this.listHubModelNames = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listHubModelNames");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListHubModelsResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listHubModels operation.
     * @callback module:api/HubModelsV1Api~listHubModelsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListHubModelsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub models
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/HubModelsV1Api~listHubModelsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListHubModelsResponse}
     */
    this.listHubModels = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listHubModels");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListHubModelsResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the patchHubModel operation.
     * @callback module:api/HubModelsV1Api~patchHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch hub model
     * @param {String} owner Owner of the namespace
     * @param {String} model_uuid UUID
     * @param {module:model/V1HubModel} body Model body
     * @param {module:api/HubModelsV1Api~patchHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    this.patchHubModel = function(owner, model_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchHubModel");
      }

      // verify the required parameter 'model_uuid' is set
      if (model_uuid === undefined || model_uuid === null) {
        throw new Error("Missing the required parameter 'model_uuid' when calling patchHubModel");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchHubModel");
      }


      var pathParams = {
        'owner': owner,
        'model.uuid': model_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1HubModel;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{model.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the updateHubModel operation.
     * @callback module:api/HubModelsV1Api~updateHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update hub model
     * @param {String} owner Owner of the namespace
     * @param {String} model_uuid UUID
     * @param {module:model/V1HubModel} body Model body
     * @param {module:api/HubModelsV1Api~updateHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    this.updateHubModel = function(owner, model_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateHubModel");
      }

      // verify the required parameter 'model_uuid' is set
      if (model_uuid === undefined || model_uuid === null) {
        throw new Error("Missing the required parameter 'model_uuid' when calling updateHubModel");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateHubModel");
      }


      var pathParams = {
        'owner': owner,
        'model.uuid': model_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1HubModel;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{model.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
