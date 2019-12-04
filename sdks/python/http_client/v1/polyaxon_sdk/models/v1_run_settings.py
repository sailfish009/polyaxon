#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    OpenAPI spec version: 1.0.0
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1RunSettings(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'namespace': 'str',
        'agent': 'V1RunSettingsCatalog',
        'queue': 'V1RunSettingsCatalog',
        'logs_store': 'V1RunSettingsCatalog',
        'outputs_store': 'V1RunSettingsCatalog',
        'init_artifacts_stores': 'list[V1RunSettingsCatalog]',
        'artifacts_stores': 'list[V1RunSettingsCatalog]',
        'git_accesses': 'list[V1RunSettingsCatalog]',
        'registry_access': 'V1RunSettingsCatalog',
        'k8s_secrets': 'list[V1RunSettingsCatalog]',
        'k8s_config_maps': 'list[V1RunSettingsCatalog]'
    }

    attribute_map = {
        'namespace': 'namespace',
        'agent': 'agent',
        'queue': 'queue',
        'logs_store': 'logs_store',
        'outputs_store': 'outputs_store',
        'init_artifacts_stores': 'init_artifacts_stores',
        'artifacts_stores': 'artifacts_stores',
        'git_accesses': 'git_accesses',
        'registry_access': 'registry_access',
        'k8s_secrets': 'k8s_secrets',
        'k8s_config_maps': 'k8s_config_maps'
    }

    def __init__(self, namespace=None, agent=None, queue=None, logs_store=None, outputs_store=None, init_artifacts_stores=None, artifacts_stores=None, git_accesses=None, registry_access=None, k8s_secrets=None, k8s_config_maps=None):  # noqa: E501
        """V1RunSettings - a model defined in Swagger"""  # noqa: E501

        self._namespace = None
        self._agent = None
        self._queue = None
        self._logs_store = None
        self._outputs_store = None
        self._init_artifacts_stores = None
        self._artifacts_stores = None
        self._git_accesses = None
        self._registry_access = None
        self._k8s_secrets = None
        self._k8s_config_maps = None
        self.discriminator = None

        if namespace is not None:
            self.namespace = namespace
        if agent is not None:
            self.agent = agent
        if queue is not None:
            self.queue = queue
        if logs_store is not None:
            self.logs_store = logs_store
        if outputs_store is not None:
            self.outputs_store = outputs_store
        if init_artifacts_stores is not None:
            self.init_artifacts_stores = init_artifacts_stores
        if artifacts_stores is not None:
            self.artifacts_stores = artifacts_stores
        if git_accesses is not None:
            self.git_accesses = git_accesses
        if registry_access is not None:
            self.registry_access = registry_access
        if k8s_secrets is not None:
            self.k8s_secrets = k8s_secrets
        if k8s_config_maps is not None:
            self.k8s_config_maps = k8s_config_maps

    @property
    def namespace(self):
        """Gets the namespace of this V1RunSettings.  # noqa: E501


        :return: The namespace of this V1RunSettings.  # noqa: E501
        :rtype: str
        """
        return self._namespace

    @namespace.setter
    def namespace(self, namespace):
        """Sets the namespace of this V1RunSettings.


        :param namespace: The namespace of this V1RunSettings.  # noqa: E501
        :type: str
        """

        self._namespace = namespace

    @property
    def agent(self):
        """Gets the agent of this V1RunSettings.  # noqa: E501


        :return: The agent of this V1RunSettings.  # noqa: E501
        :rtype: V1RunSettingsCatalog
        """
        return self._agent

    @agent.setter
    def agent(self, agent):
        """Sets the agent of this V1RunSettings.


        :param agent: The agent of this V1RunSettings.  # noqa: E501
        :type: V1RunSettingsCatalog
        """

        self._agent = agent

    @property
    def queue(self):
        """Gets the queue of this V1RunSettings.  # noqa: E501


        :return: The queue of this V1RunSettings.  # noqa: E501
        :rtype: V1RunSettingsCatalog
        """
        return self._queue

    @queue.setter
    def queue(self, queue):
        """Sets the queue of this V1RunSettings.


        :param queue: The queue of this V1RunSettings.  # noqa: E501
        :type: V1RunSettingsCatalog
        """

        self._queue = queue

    @property
    def logs_store(self):
        """Gets the logs_store of this V1RunSettings.  # noqa: E501


        :return: The logs_store of this V1RunSettings.  # noqa: E501
        :rtype: V1RunSettingsCatalog
        """
        return self._logs_store

    @logs_store.setter
    def logs_store(self, logs_store):
        """Sets the logs_store of this V1RunSettings.


        :param logs_store: The logs_store of this V1RunSettings.  # noqa: E501
        :type: V1RunSettingsCatalog
        """

        self._logs_store = logs_store

    @property
    def outputs_store(self):
        """Gets the outputs_store of this V1RunSettings.  # noqa: E501


        :return: The outputs_store of this V1RunSettings.  # noqa: E501
        :rtype: V1RunSettingsCatalog
        """
        return self._outputs_store

    @outputs_store.setter
    def outputs_store(self, outputs_store):
        """Sets the outputs_store of this V1RunSettings.


        :param outputs_store: The outputs_store of this V1RunSettings.  # noqa: E501
        :type: V1RunSettingsCatalog
        """

        self._outputs_store = outputs_store

    @property
    def init_artifacts_stores(self):
        """Gets the init_artifacts_stores of this V1RunSettings.  # noqa: E501


        :return: The init_artifacts_stores of this V1RunSettings.  # noqa: E501
        :rtype: list[V1RunSettingsCatalog]
        """
        return self._init_artifacts_stores

    @init_artifacts_stores.setter
    def init_artifacts_stores(self, init_artifacts_stores):
        """Sets the init_artifacts_stores of this V1RunSettings.


        :param init_artifacts_stores: The init_artifacts_stores of this V1RunSettings.  # noqa: E501
        :type: list[V1RunSettingsCatalog]
        """

        self._init_artifacts_stores = init_artifacts_stores

    @property
    def artifacts_stores(self):
        """Gets the artifacts_stores of this V1RunSettings.  # noqa: E501


        :return: The artifacts_stores of this V1RunSettings.  # noqa: E501
        :rtype: list[V1RunSettingsCatalog]
        """
        return self._artifacts_stores

    @artifacts_stores.setter
    def artifacts_stores(self, artifacts_stores):
        """Sets the artifacts_stores of this V1RunSettings.


        :param artifacts_stores: The artifacts_stores of this V1RunSettings.  # noqa: E501
        :type: list[V1RunSettingsCatalog]
        """

        self._artifacts_stores = artifacts_stores

    @property
    def git_accesses(self):
        """Gets the git_accesses of this V1RunSettings.  # noqa: E501


        :return: The git_accesses of this V1RunSettings.  # noqa: E501
        :rtype: list[V1RunSettingsCatalog]
        """
        return self._git_accesses

    @git_accesses.setter
    def git_accesses(self, git_accesses):
        """Sets the git_accesses of this V1RunSettings.


        :param git_accesses: The git_accesses of this V1RunSettings.  # noqa: E501
        :type: list[V1RunSettingsCatalog]
        """

        self._git_accesses = git_accesses

    @property
    def registry_access(self):
        """Gets the registry_access of this V1RunSettings.  # noqa: E501


        :return: The registry_access of this V1RunSettings.  # noqa: E501
        :rtype: V1RunSettingsCatalog
        """
        return self._registry_access

    @registry_access.setter
    def registry_access(self, registry_access):
        """Sets the registry_access of this V1RunSettings.


        :param registry_access: The registry_access of this V1RunSettings.  # noqa: E501
        :type: V1RunSettingsCatalog
        """

        self._registry_access = registry_access

    @property
    def k8s_secrets(self):
        """Gets the k8s_secrets of this V1RunSettings.  # noqa: E501


        :return: The k8s_secrets of this V1RunSettings.  # noqa: E501
        :rtype: list[V1RunSettingsCatalog]
        """
        return self._k8s_secrets

    @k8s_secrets.setter
    def k8s_secrets(self, k8s_secrets):
        """Sets the k8s_secrets of this V1RunSettings.


        :param k8s_secrets: The k8s_secrets of this V1RunSettings.  # noqa: E501
        :type: list[V1RunSettingsCatalog]
        """

        self._k8s_secrets = k8s_secrets

    @property
    def k8s_config_maps(self):
        """Gets the k8s_config_maps of this V1RunSettings.  # noqa: E501


        :return: The k8s_config_maps of this V1RunSettings.  # noqa: E501
        :rtype: list[V1RunSettingsCatalog]
        """
        return self._k8s_config_maps

    @k8s_config_maps.setter
    def k8s_config_maps(self, k8s_config_maps):
        """Sets the k8s_config_maps of this V1RunSettings.


        :param k8s_config_maps: The k8s_config_maps of this V1RunSettings.  # noqa: E501
        :type: list[V1RunSettingsCatalog]
        """

        self._k8s_config_maps = k8s_config_maps

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(V1RunSettings, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1RunSettings):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
