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


class V1RepoInit(object):
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
        'name': 'str',
        'commit': 'str',
        'branch': 'str'
    }

    attribute_map = {
        'name': 'name',
        'commit': 'commit',
        'branch': 'branch'
    }

    def __init__(self, name=None, commit=None, branch=None):  # noqa: E501
        """V1RepoInit - a model defined in Swagger"""  # noqa: E501

        self._name = None
        self._commit = None
        self._branch = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if commit is not None:
            self.commit = commit
        if branch is not None:
            self.branch = branch

    @property
    def name(self):
        """Gets the name of this V1RepoInit.  # noqa: E501


        :return: The name of this V1RepoInit.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1RepoInit.


        :param name: The name of this V1RepoInit.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def commit(self):
        """Gets the commit of this V1RepoInit.  # noqa: E501


        :return: The commit of this V1RepoInit.  # noqa: E501
        :rtype: str
        """
        return self._commit

    @commit.setter
    def commit(self, commit):
        """Sets the commit of this V1RepoInit.


        :param commit: The commit of this V1RepoInit.  # noqa: E501
        :type: str
        """

        self._commit = commit

    @property
    def branch(self):
        """Gets the branch of this V1RepoInit.  # noqa: E501


        :return: The branch of this V1RepoInit.  # noqa: E501
        :rtype: str
        """
        return self._branch

    @branch.setter
    def branch(self, branch):
        """Sets the branch of this V1RepoInit.


        :param branch: The branch of this V1RepoInit.  # noqa: E501
        :type: str
        """

        self._branch = branch

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
        if issubclass(V1RepoInit, dict):
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
        if not isinstance(other, V1RepoInit):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
