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
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import io.swagger.client.model.V1Cache;
import io.swagger.client.model.V1Component;
import io.swagger.client.model.V1Param;
import io.swagger.client.model.V1Plugins;
import io.swagger.client.model.V1Termination;
import io.swagger.client.model.V1TriggerPolicy;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * V1Operation
 */

public class V1Operation {
  @SerializedName("version")
  private Float version = null;

  @SerializedName("kind")
  private String kind = null;

  @SerializedName("name")
  private String name = null;

  @SerializedName("tag")
  private String tag = null;

  @SerializedName("description")
  private String description = null;

  @SerializedName("tags")
  private List<String> tags = null;

  @SerializedName("profile")
  private String profile = null;

  @SerializedName("queue")
  private String queue = null;

  @SerializedName("cache")
  private V1Cache cache = null;

  @SerializedName("schedule")
  private Object schedule = null;

  @SerializedName("parallel")
  private Object parallel = null;

  @SerializedName("dependencies")
  private List<String> dependencies = null;

  @SerializedName("trigger")
  private V1TriggerPolicy trigger = null;

  @SerializedName("conditions")
  private List<Object> conditions = null;

  @SerializedName("skip_on_upstream_skip")
  private Boolean skipOnUpstreamSkip = null;

  @SerializedName("termination")
  private V1Termination termination = null;

  @SerializedName("plugins")
  private V1Plugins plugins = null;

  @SerializedName("params")
  private Map<String, V1Param> params = null;

  @SerializedName("run_patch")
  private Object runPatch = null;

  @SerializedName("dag_ref")
  private String dagRef = null;

  @SerializedName("url_ref")
  private String urlRef = null;

  @SerializedName("path_ref")
  private String pathRef = null;

  @SerializedName("hub_ref")
  private String hubRef = null;

  @SerializedName("component")
  private V1Component component = null;

  public V1Operation version(Float version) {
    this.version = version;
    return this;
  }

   /**
   * Get version
   * @return version
  **/
  @ApiModelProperty(value = "")
  public Float getVersion() {
    return version;
  }

  public void setVersion(Float version) {
    this.version = version;
  }

  public V1Operation kind(String kind) {
    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @ApiModelProperty(value = "")
  public String getKind() {
    return kind;
  }

  public void setKind(String kind) {
    this.kind = kind;
  }

  public V1Operation name(String name) {
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @ApiModelProperty(value = "")
  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public V1Operation tag(String tag) {
    this.tag = tag;
    return this;
  }

   /**
   * Get tag
   * @return tag
  **/
  @ApiModelProperty(value = "")
  public String getTag() {
    return tag;
  }

  public void setTag(String tag) {
    this.tag = tag;
  }

  public V1Operation description(String description) {
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @ApiModelProperty(value = "")
  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public V1Operation tags(List<String> tags) {
    this.tags = tags;
    return this;
  }

  public V1Operation addTagsItem(String tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<String>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @ApiModelProperty(value = "")
  public List<String> getTags() {
    return tags;
  }

  public void setTags(List<String> tags) {
    this.tags = tags;
  }

  public V1Operation profile(String profile) {
    this.profile = profile;
    return this;
  }

   /**
   * Get profile
   * @return profile
  **/
  @ApiModelProperty(value = "")
  public String getProfile() {
    return profile;
  }

  public void setProfile(String profile) {
    this.profile = profile;
  }

  public V1Operation queue(String queue) {
    this.queue = queue;
    return this;
  }

   /**
   * Get queue
   * @return queue
  **/
  @ApiModelProperty(value = "")
  public String getQueue() {
    return queue;
  }

  public void setQueue(String queue) {
    this.queue = queue;
  }

  public V1Operation cache(V1Cache cache) {
    this.cache = cache;
    return this;
  }

   /**
   * Get cache
   * @return cache
  **/
  @ApiModelProperty(value = "")
  public V1Cache getCache() {
    return cache;
  }

  public void setCache(V1Cache cache) {
    this.cache = cache;
  }

  public V1Operation schedule(Object schedule) {
    this.schedule = schedule;
    return this;
  }

   /**
   * Get schedule
   * @return schedule
  **/
  @ApiModelProperty(value = "")
  public Object getSchedule() {
    return schedule;
  }

  public void setSchedule(Object schedule) {
    this.schedule = schedule;
  }

  public V1Operation parallel(Object parallel) {
    this.parallel = parallel;
    return this;
  }

   /**
   * Get parallel
   * @return parallel
  **/
  @ApiModelProperty(value = "")
  public Object getParallel() {
    return parallel;
  }

  public void setParallel(Object parallel) {
    this.parallel = parallel;
  }

  public V1Operation dependencies(List<String> dependencies) {
    this.dependencies = dependencies;
    return this;
  }

  public V1Operation addDependenciesItem(String dependenciesItem) {
    if (this.dependencies == null) {
      this.dependencies = new ArrayList<String>();
    }
    this.dependencies.add(dependenciesItem);
    return this;
  }

   /**
   * Get dependencies
   * @return dependencies
  **/
  @ApiModelProperty(value = "")
  public List<String> getDependencies() {
    return dependencies;
  }

  public void setDependencies(List<String> dependencies) {
    this.dependencies = dependencies;
  }

  public V1Operation trigger(V1TriggerPolicy trigger) {
    this.trigger = trigger;
    return this;
  }

   /**
   * Get trigger
   * @return trigger
  **/
  @ApiModelProperty(value = "")
  public V1TriggerPolicy getTrigger() {
    return trigger;
  }

  public void setTrigger(V1TriggerPolicy trigger) {
    this.trigger = trigger;
  }

  public V1Operation conditions(List<Object> conditions) {
    this.conditions = conditions;
    return this;
  }

  public V1Operation addConditionsItem(Object conditionsItem) {
    if (this.conditions == null) {
      this.conditions = new ArrayList<Object>();
    }
    this.conditions.add(conditionsItem);
    return this;
  }

   /**
   * Get conditions
   * @return conditions
  **/
  @ApiModelProperty(value = "")
  public List<Object> getConditions() {
    return conditions;
  }

  public void setConditions(List<Object> conditions) {
    this.conditions = conditions;
  }

  public V1Operation skipOnUpstreamSkip(Boolean skipOnUpstreamSkip) {
    this.skipOnUpstreamSkip = skipOnUpstreamSkip;
    return this;
  }

   /**
   * Get skipOnUpstreamSkip
   * @return skipOnUpstreamSkip
  **/
  @ApiModelProperty(value = "")
  public Boolean isSkipOnUpstreamSkip() {
    return skipOnUpstreamSkip;
  }

  public void setSkipOnUpstreamSkip(Boolean skipOnUpstreamSkip) {
    this.skipOnUpstreamSkip = skipOnUpstreamSkip;
  }

  public V1Operation termination(V1Termination termination) {
    this.termination = termination;
    return this;
  }

   /**
   * Get termination
   * @return termination
  **/
  @ApiModelProperty(value = "")
  public V1Termination getTermination() {
    return termination;
  }

  public void setTermination(V1Termination termination) {
    this.termination = termination;
  }

  public V1Operation plugins(V1Plugins plugins) {
    this.plugins = plugins;
    return this;
  }

   /**
   * Get plugins
   * @return plugins
  **/
  @ApiModelProperty(value = "")
  public V1Plugins getPlugins() {
    return plugins;
  }

  public void setPlugins(V1Plugins plugins) {
    this.plugins = plugins;
  }

  public V1Operation params(Map<String, V1Param> params) {
    this.params = params;
    return this;
  }

  public V1Operation putParamsItem(String key, V1Param paramsItem) {
    if (this.params == null) {
      this.params = new HashMap<String, V1Param>();
    }
    this.params.put(key, paramsItem);
    return this;
  }

   /**
   * Get params
   * @return params
  **/
  @ApiModelProperty(value = "")
  public Map<String, V1Param> getParams() {
    return params;
  }

  public void setParams(Map<String, V1Param> params) {
    this.params = params;
  }

  public V1Operation runPatch(Object runPatch) {
    this.runPatch = runPatch;
    return this;
  }

   /**
   * Get runPatch
   * @return runPatch
  **/
  @ApiModelProperty(value = "")
  public Object getRunPatch() {
    return runPatch;
  }

  public void setRunPatch(Object runPatch) {
    this.runPatch = runPatch;
  }

  public V1Operation dagRef(String dagRef) {
    this.dagRef = dagRef;
    return this;
  }

   /**
   * Get dagRef
   * @return dagRef
  **/
  @ApiModelProperty(value = "")
  public String getDagRef() {
    return dagRef;
  }

  public void setDagRef(String dagRef) {
    this.dagRef = dagRef;
  }

  public V1Operation urlRef(String urlRef) {
    this.urlRef = urlRef;
    return this;
  }

   /**
   * Get urlRef
   * @return urlRef
  **/
  @ApiModelProperty(value = "")
  public String getUrlRef() {
    return urlRef;
  }

  public void setUrlRef(String urlRef) {
    this.urlRef = urlRef;
  }

  public V1Operation pathRef(String pathRef) {
    this.pathRef = pathRef;
    return this;
  }

   /**
   * Get pathRef
   * @return pathRef
  **/
  @ApiModelProperty(value = "")
  public String getPathRef() {
    return pathRef;
  }

  public void setPathRef(String pathRef) {
    this.pathRef = pathRef;
  }

  public V1Operation hubRef(String hubRef) {
    this.hubRef = hubRef;
    return this;
  }

   /**
   * Get hubRef
   * @return hubRef
  **/
  @ApiModelProperty(value = "")
  public String getHubRef() {
    return hubRef;
  }

  public void setHubRef(String hubRef) {
    this.hubRef = hubRef;
  }

  public V1Operation component(V1Component component) {
    this.component = component;
    return this;
  }

   /**
   * Get component
   * @return component
  **/
  @ApiModelProperty(value = "")
  public V1Component getComponent() {
    return component;
  }

  public void setComponent(V1Component component) {
    this.component = component;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Operation v1Operation = (V1Operation) o;
    return Objects.equals(this.version, v1Operation.version) &&
        Objects.equals(this.kind, v1Operation.kind) &&
        Objects.equals(this.name, v1Operation.name) &&
        Objects.equals(this.tag, v1Operation.tag) &&
        Objects.equals(this.description, v1Operation.description) &&
        Objects.equals(this.tags, v1Operation.tags) &&
        Objects.equals(this.profile, v1Operation.profile) &&
        Objects.equals(this.queue, v1Operation.queue) &&
        Objects.equals(this.cache, v1Operation.cache) &&
        Objects.equals(this.schedule, v1Operation.schedule) &&
        Objects.equals(this.parallel, v1Operation.parallel) &&
        Objects.equals(this.dependencies, v1Operation.dependencies) &&
        Objects.equals(this.trigger, v1Operation.trigger) &&
        Objects.equals(this.conditions, v1Operation.conditions) &&
        Objects.equals(this.skipOnUpstreamSkip, v1Operation.skipOnUpstreamSkip) &&
        Objects.equals(this.termination, v1Operation.termination) &&
        Objects.equals(this.plugins, v1Operation.plugins) &&
        Objects.equals(this.params, v1Operation.params) &&
        Objects.equals(this.runPatch, v1Operation.runPatch) &&
        Objects.equals(this.dagRef, v1Operation.dagRef) &&
        Objects.equals(this.urlRef, v1Operation.urlRef) &&
        Objects.equals(this.pathRef, v1Operation.pathRef) &&
        Objects.equals(this.hubRef, v1Operation.hubRef) &&
        Objects.equals(this.component, v1Operation.component);
  }

  @Override
  public int hashCode() {
    return Objects.hash(version, kind, name, tag, description, tags, profile, queue, cache, schedule, parallel, dependencies, trigger, conditions, skipOnUpstreamSkip, termination, plugins, params, runPatch, dagRef, urlRef, pathRef, hubRef, component);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Operation {\n");
    
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    tag: ").append(toIndentedString(tag)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    profile: ").append(toIndentedString(profile)).append("\n");
    sb.append("    queue: ").append(toIndentedString(queue)).append("\n");
    sb.append("    cache: ").append(toIndentedString(cache)).append("\n");
    sb.append("    schedule: ").append(toIndentedString(schedule)).append("\n");
    sb.append("    parallel: ").append(toIndentedString(parallel)).append("\n");
    sb.append("    dependencies: ").append(toIndentedString(dependencies)).append("\n");
    sb.append("    trigger: ").append(toIndentedString(trigger)).append("\n");
    sb.append("    conditions: ").append(toIndentedString(conditions)).append("\n");
    sb.append("    skipOnUpstreamSkip: ").append(toIndentedString(skipOnUpstreamSkip)).append("\n");
    sb.append("    termination: ").append(toIndentedString(termination)).append("\n");
    sb.append("    plugins: ").append(toIndentedString(plugins)).append("\n");
    sb.append("    params: ").append(toIndentedString(params)).append("\n");
    sb.append("    runPatch: ").append(toIndentedString(runPatch)).append("\n");
    sb.append("    dagRef: ").append(toIndentedString(dagRef)).append("\n");
    sb.append("    urlRef: ").append(toIndentedString(urlRef)).append("\n");
    sb.append("    pathRef: ").append(toIndentedString(pathRef)).append("\n");
    sb.append("    hubRef: ").append(toIndentedString(hubRef)).append("\n");
    sb.append("    component: ").append(toIndentedString(component)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

