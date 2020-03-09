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
    define(['ApiClient', 'model/RuntimeError', 'model/V1Dashboard', 'model/V1ListDashboardsResponse'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/RuntimeError'), require('../model/V1Dashboard'), require('../model/V1ListDashboardsResponse'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.ProjectDashboardsV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.RuntimeError, root.PolyaxonSdk.V1Dashboard, root.PolyaxonSdk.V1ListDashboardsResponse);
  }
}(this, function(ApiClient, RuntimeError, V1Dashboard, V1ListDashboardsResponse) {
  'use strict';

  /**
   * ProjectDashboardsV1 service.
   * @module api/ProjectDashboardsV1Api
   * @version 1.0.5
   */

  /**
   * Constructs a new ProjectDashboardsV1Api. 
   * @alias module:api/ProjectDashboardsV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the createProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~createProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {module:model/V1Dashboard} body Dashboard body
     * @param {module:api/ProjectDashboardsV1Api~createProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.createProjectDashboard = function(owner, project, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling createProjectDashboard");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~deleteProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project where the notification will be assigned
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/ProjectDashboardsV1Api~deleteProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteProjectDashboard = function(owner, project, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling deleteProjectDashboard");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
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
        '/api/v1/{owner}/{project}/dashboards/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~getProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project where the notification will be assigned
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/ProjectDashboardsV1Api~getProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.getProjectDashboard = function(owner, project, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling getProjectDashboard");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listProjectDashboardNames operation.
     * @callback module:api/ProjectDashboardsV1Api~listProjectDashboardNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListDashboardsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/ProjectDashboardsV1Api~listProjectDashboardNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListDashboardsResponse}
     */
    this.listProjectDashboardNames = function(owner, project, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listProjectDashboardNames");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling listProjectDashboardNames");
      }


      var pathParams = {
        'owner': owner,
        'project': project
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
      var returnType = V1ListDashboardsResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listProjectDashboards operation.
     * @callback module:api/ProjectDashboardsV1Api~listProjectDashboardsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListDashboardsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List project dashboards
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/ProjectDashboardsV1Api~listProjectDashboardsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListDashboardsResponse}
     */
    this.listProjectDashboards = function(owner, project, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listProjectDashboards");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling listProjectDashboards");
      }


      var pathParams = {
        'owner': owner,
        'project': project
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
      var returnType = V1ListDashboardsResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the patchProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~patchProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} dashboard_uuid UUID
     * @param {module:model/V1Dashboard} body Dashboard body
     * @param {module:api/ProjectDashboardsV1Api~patchProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.patchProjectDashboard = function(owner, project, dashboard_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling patchProjectDashboard");
      }

      // verify the required parameter 'dashboard_uuid' is set
      if (dashboard_uuid === undefined || dashboard_uuid === null) {
        throw new Error("Missing the required parameter 'dashboard_uuid' when calling patchProjectDashboard");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'dashboard.uuid': dashboard_uuid
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the promoteProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~promoteProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Promote project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} dashboard_uuid UUID
     * @param {module:api/ProjectDashboardsV1Api~promoteProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.promoteProjectDashboard = function(owner, project, dashboard_uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling promoteProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling promoteProjectDashboard");
      }

      // verify the required parameter 'dashboard_uuid' is set
      if (dashboard_uuid === undefined || dashboard_uuid === null) {
        throw new Error("Missing the required parameter 'dashboard_uuid' when calling promoteProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'dashboard.uuid': dashboard_uuid
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}/promote', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the updateProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~updateProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} dashboard_uuid UUID
     * @param {module:model/V1Dashboard} body Dashboard body
     * @param {module:api/ProjectDashboardsV1Api~updateProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.updateProjectDashboard = function(owner, project, dashboard_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling updateProjectDashboard");
      }

      // verify the required parameter 'dashboard_uuid' is set
      if (dashboard_uuid === undefined || dashboard_uuid === null) {
        throw new Error("Missing the required parameter 'dashboard_uuid' when calling updateProjectDashboard");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'dashboard.uuid': dashboard_uuid
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
