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

from polyaxon.polyflow.component.component import ComponentSchema, V1Component
from polyaxon.polyflow.references import (
    DagReferenceSchema,
    HubReferenceSchema,
    PathReferenceSchema,
    UrlReferenceSchema,
    V1DagReference,
    V1HubReference,
    V1PathReference,
    V1UrlReference,
)
from polyaxon.schemas.base import BaseOneOfSchema


class ComponentReferenceSchema(BaseOneOfSchema):
    TYPE_FIELD = "kind"
    TYPE_FIELD_REMOVE = False
    SCHEMAS = {
        V1Component.IDENTIFIER: ComponentSchema,
        V1DagReference.IDENTIFIER: DagReferenceSchema,
        V1HubReference.IDENTIFIER: HubReferenceSchema,
        V1PathReference.IDENTIFIER: PathReferenceSchema,
        V1UrlReference.IDENTIFIER: UrlReferenceSchema,
    }
