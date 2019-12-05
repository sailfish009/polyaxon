# OrganizationsV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createOrganization**](OrganizationsV1Api.md#createOrganization) | **POST** /api/v1/organizations/create | List runs
[**createOrganizationMember**](OrganizationsV1Api.md#createOrganizationMember) | **POST** /api/v1/organizations/{owner}/members | Delete runs
[**deleteOrganization**](OrganizationsV1Api.md#deleteOrganization) | **DELETE** /api/v1/organizations/{owner} | Patch run
[**deleteOrganizationMember**](OrganizationsV1Api.md#deleteOrganizationMember) | **DELETE** /api/v1/organizations/{owner}/members/{member.user} | Invalidate runs
[**getOrganization**](OrganizationsV1Api.md#getOrganization) | **GET** /api/v1/organizations/{owner} | Create new run
[**getOrganizationMember**](OrganizationsV1Api.md#getOrganizationMember) | **GET** /api/v1/organizations/{owner}/members/{member.user} | Stop run
[**listOrganizationMembers**](OrganizationsV1Api.md#listOrganizationMembers) | **GET** /api/v1/organizations/{owner}/members | Delete run
[**listOrganizationNames**](OrganizationsV1Api.md#listOrganizationNames) | **GET** /api/v1/organizations/names | List bookmarked runs for user
[**listOrganizations**](OrganizationsV1Api.md#listOrganizations) | **GET** /api/v1/organizations/list | List archived runs for user
[**patchOrganization**](OrganizationsV1Api.md#patchOrganization) | **PATCH** /api/v1/organizations/{owner} | Update run
[**patchOrganizationMember**](OrganizationsV1Api.md#patchOrganizationMember) | **PATCH** /api/v1/organizations/{owner}/members/{member.user} | Invalidate run
[**updateOrganization**](OrganizationsV1Api.md#updateOrganization) | **PUT** /api/v1/organizations/{owner} | Get run
[**updateOrganizationMember**](OrganizationsV1Api.md#updateOrganizationMember) | **PUT** /api/v1/organizations/{owner}/members/{member.user} | Stop runs


<a name="createOrganization"></a>
# **createOrganization**
> V1Organization createOrganization(body)

List runs

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
V1Organization body = new V1Organization(); // V1Organization | 
try {
    V1Organization result = apiInstance.createOrganization(body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#createOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1Organization**](V1Organization.md)|  |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="createOrganizationMember"></a>
# **createOrganizationMember**
> V1OrganizationMember createOrganizationMember(owner, body)

Delete runs

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1OrganizationMember body = new V1OrganizationMember(); // V1OrganizationMember | Organization body
try {
    V1OrganizationMember result = apiInstance.createOrganizationMember(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#createOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteOrganization"></a>
# **deleteOrganization**
> deleteOrganization(owner)

Patch run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
try {
    apiInstance.deleteOrganization(owner);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#deleteOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteOrganizationMember"></a>
# **deleteOrganizationMember**
> deleteOrganizationMember(owner, memberUser, memberRole, memberCreatedAt, memberUpdatedAt)

Invalidate runs

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String memberUser = "memberUser_example"; // String | User
String memberRole = "memberRole_example"; // String | Role.
OffsetDateTime memberCreatedAt = OffsetDateTime.now(); // OffsetDateTime | Optional time when the entityt was created.
OffsetDateTime memberUpdatedAt = OffsetDateTime.now(); // OffsetDateTime | Optional last time the entity was updated.
try {
    apiInstance.deleteOrganizationMember(owner, memberUser, memberRole, memberCreatedAt, memberUpdatedAt);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#deleteOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **memberUser** | **String**| User |
 **memberRole** | **String**| Role. | [optional]
 **memberCreatedAt** | **OffsetDateTime**| Optional time when the entityt was created. | [optional]
 **memberUpdatedAt** | **OffsetDateTime**| Optional last time the entity was updated. | [optional]

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getOrganization"></a>
# **getOrganization**
> V1Organization getOrganization(owner)

Create new run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
try {
    V1Organization result = apiInstance.getOrganization(owner);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#getOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getOrganizationMember"></a>
# **getOrganizationMember**
> V1OrganizationMember getOrganizationMember(owner, memberUser, memberRole, memberCreatedAt, memberUpdatedAt)

Stop run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String memberUser = "memberUser_example"; // String | User
String memberRole = "memberRole_example"; // String | Role.
OffsetDateTime memberCreatedAt = OffsetDateTime.now(); // OffsetDateTime | Optional time when the entityt was created.
OffsetDateTime memberUpdatedAt = OffsetDateTime.now(); // OffsetDateTime | Optional last time the entity was updated.
try {
    V1OrganizationMember result = apiInstance.getOrganizationMember(owner, memberUser, memberRole, memberCreatedAt, memberUpdatedAt);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#getOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **memberUser** | **String**| User |
 **memberRole** | **String**| Role. | [optional]
 **memberCreatedAt** | **OffsetDateTime**| Optional time when the entityt was created. | [optional]
 **memberUpdatedAt** | **OffsetDateTime**| Optional last time the entity was updated. | [optional]

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOrganizationMembers"></a>
# **listOrganizationMembers**
> V1ListOrganizationMembersResponse listOrganizationMembers(owner)

Delete run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
try {
    V1ListOrganizationMembersResponse result = apiInstance.listOrganizationMembers(owner);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#listOrganizationMembers");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |

### Return type

[**V1ListOrganizationMembersResponse**](V1ListOrganizationMembersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOrganizationNames"></a>
# **listOrganizationNames**
> V1ListOrganizationsResponse listOrganizationNames()

List bookmarked runs for user

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
try {
    V1ListOrganizationsResponse result = apiInstance.listOrganizationNames();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#listOrganizationNames");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListOrganizationsResponse**](V1ListOrganizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOrganizations"></a>
# **listOrganizations**
> V1ListOrganizationsResponse listOrganizations()

List archived runs for user

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
try {
    V1ListOrganizationsResponse result = apiInstance.listOrganizations();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#listOrganizations");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListOrganizationsResponse**](V1ListOrganizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="patchOrganization"></a>
# **patchOrganization**
> V1Organization patchOrganization(owner, body)

Update run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1Organization body = new V1Organization(); // V1Organization | Organization body
try {
    V1Organization result = apiInstance.patchOrganization(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#patchOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1Organization**](V1Organization.md)| Organization body |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="patchOrganizationMember"></a>
# **patchOrganizationMember**
> V1OrganizationMember patchOrganizationMember(owner, memberUser, body)

Invalidate run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String memberUser = "memberUser_example"; // String | User
V1OrganizationMember body = new V1OrganizationMember(); // V1OrganizationMember | Organization body
try {
    V1OrganizationMember result = apiInstance.patchOrganizationMember(owner, memberUser, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#patchOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **memberUser** | **String**| User |
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateOrganization"></a>
# **updateOrganization**
> V1Organization updateOrganization(owner, body)

Get run

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1Organization body = new V1Organization(); // V1Organization | Organization body
try {
    V1Organization result = apiInstance.updateOrganization(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#updateOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1Organization**](V1Organization.md)| Organization body |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateOrganizationMember"></a>
# **updateOrganizationMember**
> V1OrganizationMember updateOrganizationMember(owner, memberUser, body)

Stop runs

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String memberUser = "memberUser_example"; // String | User
V1OrganizationMember body = new V1OrganizationMember(); // V1OrganizationMember | Organization body
try {
    V1OrganizationMember result = apiInstance.updateOrganizationMember(owner, memberUser, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#updateOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **memberUser** | **String**| User |
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json
