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

from __future__ import absolute_import

# flake8: noqa

# import apis into api package
from polyaxon_sdk.api.auth_service_api import AuthServiceApi
from polyaxon_sdk.api.project_service_api import ProjectServiceApi
from polyaxon_sdk.api.run_service_api import RunServiceApi
from polyaxon_sdk.api.user_service_api import UserServiceApi
from polyaxon_sdk.api.version_service_api import VersionServiceApi
