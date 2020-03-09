#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
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

from tests.utils import BaseTestCase

from polyaxon.utils.string_utils import is_protected_type


class TestUtils(BaseTestCase):
    def test_is_protected_type(self):
        assert is_protected_type(None) is True
        assert is_protected_type(1) is True
        assert is_protected_type(1.1) is True
        assert is_protected_type("foo") is False
