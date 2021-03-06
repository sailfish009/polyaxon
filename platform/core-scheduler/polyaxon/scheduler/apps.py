#!/usr/bin/python
#
# Copyright 2018-2021 Polyaxon, Inc.
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

from django.apps import AppConfig


class SchedulerConfig(AppConfig):
    name = "scheduler"
    verbose_name = "Scheduler"

    def ready(self):
        import coredb.signals.runs  # noqa

        from coredb import executor, operations
        from polycommon import conf
        from polycommon import auditor

        conf.validate_and_setup()
        operations.validate_and_setup()
        executor.validate_and_setup()
        auditor.validate_and_setup()

        import polycommon.options.conf_subscriptions  # noqa
        from polycommon.events import auditor_subscriptions  # noqa
