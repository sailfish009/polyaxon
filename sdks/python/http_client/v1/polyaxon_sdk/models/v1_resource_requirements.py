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


class V1ResourceRequirements(object):
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
        'limits': 'object',
        'requests': 'object'
    }

    attribute_map = {
        'limits': 'limits',
        'requests': 'requests'
    }

    def __init__(self, limits=None, requests=None):  # noqa: E501
        """V1ResourceRequirements - a model defined in Swagger"""  # noqa: E501

        self._limits = None
        self._requests = None
        self.discriminator = None

        if limits is not None:
            self.limits = limits
        if requests is not None:
            self.requests = requests

    @property
    def limits(self):
        """Gets the limits of this V1ResourceRequirements.  # noqa: E501

        Limits describes the maximum amount of compute resources allowed.  # noqa: E501

        :return: The limits of this V1ResourceRequirements.  # noqa: E501
        :rtype: object
        """
        return self._limits

    @limits.setter
    def limits(self, limits):
        """Sets the limits of this V1ResourceRequirements.

        Limits describes the maximum amount of compute resources allowed.  # noqa: E501

        :param limits: The limits of this V1ResourceRequirements.  # noqa: E501
        :type: object
        """

        self._limits = limits

    @property
    def requests(self):
        """Gets the requests of this V1ResourceRequirements.  # noqa: E501

        Requests describes the minimum amount of compute resources required.  # noqa: E501

        :return: The requests of this V1ResourceRequirements.  # noqa: E501
        :rtype: object
        """
        return self._requests

    @requests.setter
    def requests(self, requests):
        """Sets the requests of this V1ResourceRequirements.

        Requests describes the minimum amount of compute resources required.  # noqa: E501

        :param requests: The requests of this V1ResourceRequirements.  # noqa: E501
        :type: object
        """

        self._requests = requests

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
        if issubclass(V1ResourceRequirements, dict):
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
        if not isinstance(other, V1ResourceRequirements):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
