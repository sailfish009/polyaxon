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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('V1RunKind', function() {
      beforeEach(function() {
        instance = PolyaxonSdk.V1RunKind;
      });

      it('should create an instance of V1RunKind', function() {
        // TODO: update the code to test V1RunKind
        expect(instance).to.be.a('object');
      });

      it('should have the property job', function() {
        expect(instance).to.have.property('job');
        expect(instance.job).to.be("job");
      });

      it('should have the property service', function() {
        expect(instance).to.have.property('service');
        expect(instance.service).to.be("service");
      });

      it('should have the property dag', function() {
        expect(instance).to.have.property('dag');
        expect(instance.dag).to.be("dag");
      });

      it('should have the property spark', function() {
        expect(instance).to.have.property('spark');
        expect(instance.spark).to.be("spark");
      });

      it('should have the property dask', function() {
        expect(instance).to.have.property('dask');
        expect(instance.dask).to.be("dask");
      });

      it('should have the property flink', function() {
        expect(instance).to.have.property('flink');
        expect(instance.flink).to.be("flink");
      });

      it('should have the property ray', function() {
        expect(instance).to.have.property('ray');
        expect(instance.ray).to.be("ray");
      });

      it('should have the property mpijob', function() {
        expect(instance).to.have.property('mpijob');
        expect(instance.mpijob).to.be("mpijob");
      });

      it('should have the property tfjob', function() {
        expect(instance).to.have.property('tfjob');
        expect(instance.tfjob).to.be("tfjob");
      });

      it('should have the property pytorchjob', function() {
        expect(instance).to.have.property('pytorchjob');
        expect(instance.pytorchjob).to.be("pytorchjob");
      });

      it('should have the property parallel', function() {
        expect(instance).to.have.property('parallel');
        expect(instance.parallel).to.be("parallel");
      });

      it('should have the property scheduler', function() {
        expect(instance).to.have.property('scheduler');
        expect(instance.scheduler).to.be("scheduler");
      });

      it('should have the property tuner', function() {
        expect(instance).to.have.property('tuner');
        expect(instance.tuner).to.be("tuner");
      });

      it('should have the property watchdog', function() {
        expect(instance).to.have.property('watchdog');
        expect(instance.watchdog).to.be("watchdog");
      });

      it('should have the property notifier', function() {
        expect(instance).to.have.property('notifier');
        expect(instance.notifier).to.be("notifier");
      });

    });
  });

}));
