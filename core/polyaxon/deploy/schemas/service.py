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

from marshmallow import EXCLUDE, fields

from polyaxon.deploy.schemas.celery import CelerySchema
from polyaxon.k8s import k8s_schemas
from polyaxon.schemas.base import BaseCamelSchema, BaseConfig
from polyaxon.schemas.fields.swagger import SwaggerField


class ServiceSchema(BaseCamelSchema):
    enabled = fields.Bool(allow_none=True)
    image = fields.Str(allow_none=True)
    image_tag = fields.Str(allow_none=True)
    image_pull_policy = fields.Str(allow_none=True)
    replicas = fields.Int(allow_none=True)
    concurrency = fields.Int(allow_none=True)
    resources = SwaggerField(cls=k8s_schemas.V1ResourceRequirements, allow_none=True)

    class Meta:
        unknown = EXCLUDE

    @staticmethod
    def schema_config():
        return V1Service


class V1Service(BaseConfig):
    SCHEMA = ServiceSchema
    REDUCED_ATTRIBUTES = [
        "enabled",
        "image",
        "imageTag",
        "imagePullPolicy",
        "replicas",
        "concurrency",
        "resources",
    ]

    def __init__(
        self,
        enabled=None,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        concurrency=None,
        resources=None,
    ):
        self.enabled = enabled
        self.image = image
        self.image_tag = image_tag
        self.image_pull_policy = image_pull_policy
        self.replicas = replicas
        self.concurrency = concurrency
        self.resources = resources


class WorkerSchema(ServiceSchema):
    celery = fields.Nested(CelerySchema, allow_none=True)

    @staticmethod
    def schema_config():
        return WorkerConfig


class WorkerConfig(V1Service):
    SCHEMA = WorkerSchema
    REDUCED_ATTRIBUTES = V1Service.REDUCED_ATTRIBUTES + ["celery"]

    def __init__(
        self,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        concurrency=None,
        resources=None,
        celery=None,
    ):
        super().__init__(
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            concurrency=concurrency,
            resources=resources,
        )
        self.celery = celery


class HelperSchema(ServiceSchema):
    sleepInterval = fields.Int(allow_none=True)
    syncInterval = fields.Int(allow_none=True)

    @staticmethod
    def schema_config():
        return HelperConfig


class HelperConfig(V1Service):
    SCHEMA = HelperSchema
    REDUCED_ATTRIBUTES = V1Service.REDUCED_ATTRIBUTES + [
        "sleepInterval",
        "syncInterval",
    ]

    def __init__(
        self,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        concurrency=None,
        resources=None,
        sleep_interval=None,
        sync_interval=None,
    ):
        super().__init__(
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            concurrency=concurrency,
            resources=resources,
        )
        self.sleep_interval = sleep_interval
        self.sync_interval = sync_interval


class AgentSchema(ServiceSchema):
    instance = fields.String(allow_none=True)
    token = fields.String(allow_none=True)

    @staticmethod
    def schema_config():
        return AgentConfig


class AgentConfig(V1Service):
    SCHEMA = AgentSchema
    REDUCED_ATTRIBUTES = V1Service.REDUCED_ATTRIBUTES + ["instance", "token"]

    def __init__(
        self,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        concurrency=None,
        resources=None,
        instance=None,
        token=None,
    ):
        super().__init__(
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            concurrency=concurrency,
            resources=resources,
        )
        self.instance = instance
        self.token = token


class ApiSchema(ServiceSchema):
    service = fields.Dict(allow_none=True)

    @staticmethod
    def schema_config():
        return ApiConfig


class ApiConfig(V1Service):
    SCHEMA = ApiSchema

    def __init__(
        self,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        concurrency=None,
        resources=None,
        service=None,
    ):
        super().__init__(
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            concurrency=concurrency,
            resources=resources,
        )
        self.service = service


class HooksSchema(ServiceSchema):
    load_fixtures = fields.Bool(allow_none=True)

    @staticmethod
    def schema_config():
        return HooksConfig


class HooksConfig(V1Service):
    SCHEMA = HooksSchema
    REDUCED_ATTRIBUTES = V1Service.REDUCED_ATTRIBUTES + ["loadFixtures"]

    def __init__(
        self,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        concurrency=None,
        resources=None,
        load_fixtures=None,
    ):
        super().__init__(
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            concurrency=concurrency,
            resources=resources,
        )
        self.load_fixtures = load_fixtures


class ThirdPartyServiceSchema(ServiceSchema):
    enabled = fields.Bool(allow_none=True)
    persistence = fields.Dict(allow_none=True)
    node_selector = fields.Dict(allow_none=True)
    affinity = fields.Dict(allow_none=True)
    tolerations = fields.List(fields.Dict(allow_none=True), allow_none=True)

    @staticmethod
    def schema_config():
        return ThirdPartyV1Service


class ThirdPartyV1Service(V1Service):
    SCHEMA = ThirdPartyServiceSchema
    REDUCED_ATTRIBUTES = [
        "enabled",
        "image",
        "imageTag",
        "imagePullPolicy",
        "replicas",
        "concurrency",
        "resources",
        "persistence",
        "nodeSelector",
        "affinity",
        "tolerations",
    ]

    def __init__(
        self,
        enabled=None,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        resources=None,
        persistence=None,
        node_selector=None,
        affinity=None,
        tolerations=None,
    ):
        super().__init__(
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            resources=resources,
        )
        self.enabled = enabled
        self.persistence = persistence
        self.node_selector = node_selector
        self.affinity = affinity
        self.tolerations = tolerations


class PostgresqlSchema(ThirdPartyServiceSchema):
    postgres_user = fields.Str(allow_none=True)
    postgres_password = fields.Str(allow_none=True)
    postgres_database = fields.Str(allow_none=True)
    conn_max_age = fields.Int(allow_none=True)

    @staticmethod
    def schema_config():
        return PostgresqlConfig


class PostgresqlConfig(ThirdPartyV1Service):
    SCHEMA = PostgresqlSchema
    REDUCED_ATTRIBUTES = ThirdPartyV1Service.REDUCED_ATTRIBUTES + [
        "postgresUser",
        "postgresPassword",
        "postgresDatabase",
        "connMaxAge",
    ]

    def __init__(
        self,
        enabled=None,
        postgres_user=None,
        postgres_password=None,
        postgres_database=None,
        conn_max_age=None,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        resources=None,
        persistence=None,
        node_selector=None,
        affinity=None,
        tolerations=None,
    ):
        super().__init__(
            enabled=enabled,
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            resources=resources,
            persistence=persistence,
            node_selector=node_selector,
            affinity=affinity,
            tolerations=tolerations,
        )
        self.postgres_user = postgres_user
        self.postgres_password = postgres_password
        self.postgres_database = postgres_database
        self.conn_max_age = conn_max_age


class RedisSchema(ThirdPartyServiceSchema):
    image = fields.Raw(allow_none=True)
    use_password = fields.Bool(allow_none=True)
    password = fields.Str(allow_none=True)

    @staticmethod
    def schema_config():
        return RedisConfig


class RedisConfig(ThirdPartyV1Service):
    SCHEMA = RedisSchema
    REDUCED_ATTRIBUTES = ThirdPartyV1Service.REDUCED_ATTRIBUTES + [
        "usePassword",
        "password",
    ]

    def __init__(
        self,
        enabled=None,
        use_password=None,
        password=None,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        resources=None,
        persistence=None,
        node_selector=None,
        affinity=None,
        tolerations=None,
    ):
        super().__init__(
            enabled=enabled,
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            resources=resources,
            persistence=persistence,
            node_selector=node_selector,
            affinity=affinity,
            tolerations=tolerations,
        )
        self.use_password = use_password
        self.password = password


class RabbitmqSchema(ThirdPartyServiceSchema):
    rabbitmq_username = fields.Str(allow_none=True)
    rabbitmq_password = fields.Str(allow_none=True)

    @staticmethod
    def schema_config():
        return RabbitmqConfig


class RabbitmqConfig(ThirdPartyV1Service):
    SCHEMA = RabbitmqSchema
    REDUCED_ATTRIBUTES = ThirdPartyV1Service.REDUCED_ATTRIBUTES + [
        "rabbitmqUsername",
        "rabbitmqPassword",
    ]

    def __init__(
        self,
        enabled=None,
        rabbitmq_username=None,
        rabbitmq_password=None,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        resources=None,
        persistence=None,
        node_selector=None,
        affinity=None,
        tolerations=None,
    ):
        super().__init__(
            enabled=enabled,
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            resources=resources,
            persistence=persistence,
            node_selector=node_selector,
            affinity=affinity,
            tolerations=tolerations,
        )
        self.rabbitmq_username = rabbitmq_username
        self.rabbitmq_password = rabbitmq_password


class DockerRegistrySchema(ThirdPartyServiceSchema):
    registry_user = fields.Str(allow_none=True)
    registry_password = fields.Str(allow_none=True)

    @staticmethod
    def schema_config():
        return DockerRegistryConfig


class DockerRegistryConfig(ThirdPartyV1Service):
    SCHEMA = DockerRegistrySchema
    REDUCED_ATTRIBUTES = ThirdPartyV1Service.REDUCED_ATTRIBUTES + [
        "registryUser",
        "registryPassword",
    ]

    def __init__(
        self,
        enabled=None,
        registry_user=None,
        registry_password=None,
        image=None,
        image_tag=None,
        image_pull_policy=None,
        replicas=None,
        resources=None,
        persistence=None,
        node_selector=None,
        affinity=None,
        tolerations=None,
    ):
        super().__init__(
            enabled=enabled,
            image=image,
            image_tag=image_tag,
            image_pull_policy=image_pull_policy,
            replicas=replicas,
            resources=resources,
            persistence=persistence,
            node_selector=node_selector,
            affinity=affinity,
            tolerations=tolerations,
        )
        self.registry_user = registry_user
        self.registry_password = registry_password


class ExternalServiceSchema(BaseCamelSchema):
    user = fields.Str(allow_none=True)
    password = fields.Str(allow_none=True)
    host = fields.Str(allow_none=True)
    port = fields.Int(allow_none=True)
    database = fields.Str(allow_none=True)
    use_password = fields.Bool(allow_none=True)
    conn_max_age = fields.Int(allow_none=True)

    @staticmethod
    def schema_config():
        return ExternalV1Service


class ExternalV1Service(BaseConfig):
    SCHEMA = ExternalServiceSchema
    REDUCED_ATTRIBUTES = [
        "user",
        "password",
        "host",
        "port",
        "database",
        "usePassword",
        "connMaxAge",
    ]

    def __init__(
        self,
        user=None,
        password=None,
        host=None,
        port=None,
        database=None,
        use_password=None,
        conn_max_age=None,
    ):
        self.user = user
        self.password = password
        self.host = host
        self.port = port
        self.database = database
        self.use_password = use_password
        self.conn_max_age = conn_max_age


class ExternalServicesSchema(BaseCamelSchema):
    redis = fields.Nested(ExternalServiceSchema, allow_none=True)
    rabbitmq = fields.Nested(ExternalServiceSchema, allow_none=True)
    postgresql = fields.Nested(ExternalServiceSchema, allow_none=True)

    @staticmethod
    def schema_config():
        return ExternalServicesConfig


class ExternalServicesConfig(BaseConfig):
    SCHEMA = ExternalServicesSchema
    REDUCED_ATTRIBUTES = ["redis", "rabbitmq", "postgresql"]

    def __init__(self, redis=None, rabbitmq=None, postgresql=None):
        self.redis = redis
        self.rabbitmq = rabbitmq
        self.postgresql = postgresql
