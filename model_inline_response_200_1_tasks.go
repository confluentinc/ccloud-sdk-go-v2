/*
 * Confluent Cloud APIs
 *
 * <div class=\"status-info\"> <p class=\"status-info-title\">Note</p> This documents the collection of Confluent Cloud APIs. Each API documents its <a href=\"#section/Versioning/API-Lifecycle-Policy\">lifecycle phase</a>. APIs marked as Early Access or Preview are not ready for production usage. We're currently working with a select group of customers to get feedback and iterate on these APIs. </div>  # Introduction  Confluent Cloud APIs are a core building block of Confluent Cloud. You can use the APIs to manage your own account or to integrate Confluent into your product.  Most of the APIs are organized around <a href=\"http://en.wikipedia.org/wiki/Representational_State_Transfer\" target=\"_blank\">REST</a> and the resources which make up Confluent Cloud. The APIs have predictable resource-oriented URLs, transport data using JSON, and use standard HTTP verbs, response codes, authentication, and design principles.  # Object Model  <div class=\"status-info\"> <p class=\"status-info-title\">Note</p> This section describes the object model for many Confluent Cloud APIs, but not all. The Connect v1 API group has a different object model. Please review the example request and response bodies in the Connect v1 API documentation <a href=\"#tag/Connectors-(v1)\">below</a> to see its object model. </div>  Confluent Cloud APIs are primarily designed to be declarative and intent-oriented. In other words,  tell the API what you want (e.g., throughput or SLOs) and it will figure out how to make it happen  (e.g., cluster sizing). A Confluent object acts as a \"record of intent\" — after you create the object, Confluent Cloud will work tirelessly in the background to ensure that the object exists as specified.  Confluent APIs represent objects in JSON with media-type `application/json`.  Many objects follow a model consisting of `spec` and `status`. An object's `spec` tells Confluent the _desired state_ (specification) of the resource. The object may not be immediately available or changes may not be immediately applied. For this reason, many objects also have a `status` property that provides info about the current _actual state_ of the resource. Confluent Cloud is continuously and actively managing each resource's actual state to match it's desired state.  All Confluent objects share a set of common properties:  * **api_version** – API objects have an `api_version` field indicating their API version. * **kind** – API objects have a `kind` field indicating the kind of object it is. * **id** – Each object in the API will have an identifier, indicated via its `id` field,   and should be treated as an opaque string unless otherwise specified.  There are a number of other [standard properties](#standard-properties) and that you'll encounter used by many API objects. And of course, objects have plenty of non-standard fields that are specific to each object _kind_... this is what makes them interesting!  # Authentication  Confluent uses API keys for integrating with Confluent Cloud. Integrations must be authorized and authenticated before they can access or manage resources in Confluent Cloud. You can manage your API keys in the Confluent Cloud Dashboard or Confluent Cloud CLI. See [the docs for details on creating and managing API keys](https://docs.confluent.io/cloud/current/client-apps/api-keys.html).  Today, you can divide API keys into two classes:  * **Cloud API Keys** – These grant access to the Confluent Cloud Control Plane APIs,   such as for Provisioning and Metrics integrations. * **Cluster API Keys** – These grant access to a single Confluent cluster, such as a specific   Kafka or Schema Registry cluster.  An API key is owned by a User or Service Account and inherits the permissions granted to the owner.  * For example, if service account `Armageddon` is granted ACLs on Kafka Cluster `neptune`,   then a Cluster API Key for `neptune` owned by `Armageddon` will have these ACLs. * **Note 1:** ACLs aren't currently enforced for users, only service accounts. * **Note 2:** API Keys are automatically deleted when the owning user or service account is   deleted (e.g., an employee leaves the company, or moves to a new department and an SSO   integration removes their user in Confluent Cloud as they no longer require access) * Confluent **strongly recommend** that you use service accounts for all production-critical access.  Your API keys carry many privileges, so **be sure to keep them secure**! Do not share your API keys in publicly accessible areas such as GitHub, client-side code, and so forth.  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  To use your API key, send it as an `Authorization: Basic {key}` header. Remember that HTTP Basic authorization requires you to colon-separate and base64 encode your key. For example, if your API Key ID is `ABCDEFGH123456789` and the corresponding API Key Secret is `XNCIW93I2L1SQPJSJ823K1LS902KLDFMCZPWEO`, then the authorization header will be      Authorization: Basic QUJDREVGR0gxMjM0NTY3ODk6WE5DSVc5M0kyTDFTUVBKU0o4MjNLMUxTOTAyS0xERk1DWlBXRU8=  This example header can be generated from the API key (on Mac OS X) with      $ echo -n \"ABCDEFGH123456789:XNCIW93I2L1SQPJSJ823K1LS902KLDFMCZPWEO\" | base64  <!-- TODO: port this back to the Confluent API Design Guide -->  <SecurityDefinitions />  # Errors  <div class=\"status-info\"> <p class=\"status-info-title\">Note</p> This section describes the structure of error responses for many Confluent Cloud APIs, but not all. The Connect v1 API group has a different set of structures for error responses. Please review the example request and response bodies in the Connect v1 API documentation <a href=\"#tag/Connectors-(v1)\">below</a> to see its error behaviour. </div>  Confluent uses conventional [HTTP status codes](#section/HTTP-Guidelines/Status-Codes) to indicate the success or failure of an API request.  Failures follow a standard model to tell you about what went wrong. They may include one or more error objects with the following fields:  | Field      | Type    | Description |------------|---------|-------------------------------------- | id*        | UUID    | A unique identifier for this particular occurrence of the problem. | status     | String  | The HTTP status code applicable to this problem. | code       | String  | An application-specific error code. | title      | String  | A short, human-readable summary of the problem that **should not** change from occurrence to occurrence of the problem, except for purposes of localization. | detail*    | String  | A human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized. | source     | Object  | An object that references the source of the error, and optionally includes any of the following members: | &nbsp;&nbsp;pointer   | String  | A <a href=\"https://tools.ietf.org/html/rfc6901\" target=\"_blank\">JSON Pointer</a> to the associated entity in the request document (e.g. `\"/spec/title\"` for a specific attribute). | &nbsp;&nbsp;parameter | String  | A string indicating which URI query parameter caused the error. | meta       | Object  | A meta object that contains non-standard meta-information about the error. | resolution | String  | Instructions for the end-user for correcting the error.  \\* indicates a required field  All errors include an `id` and some `detail` message. The `id` is a unique identifier — use it when you're working with Confluent support to debug a problem with a specific API call. The `detail` describes what went wrong.  Some errors that could be handled programmatically (e.g., a Kafka cluster config is invalid) may include an error `code` that briefly explains the error reported.  Validation issues and similar errors include a `source` which tells you exactly what in the request was responsible for the error.  For example, a failure may look like      {       \"errors\": [{         \"status\": \"422\",         \"code\": \"invalid_configuration\",         \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\",         \"title\": \"The Kafka cluster configuration is invalid\",         \"detail\": \"The property '/cluster/storage_size' of type string did not match the following type: integer\",         \"source\": {           \"pointer\": \"/cluster/storage_size\"         }       }]     }  If a request fails validation, it will return an HTTP `422 Unprocessable Entity` with a list of fields that failed validation.  # Pagination  <div class=\"status-info\"> <p class=\"status-info-title\">Note</p> This section describes the pagination behavior of “list” operations for many Confluent Cloud APIs, but not all. The Connect v1 API list operations do not support pagination. </div>  All API resources have support for bulk reads via \"list\" API operations. For example, you can \"list Kafka clusters\", \"list api keys\", and \"list environments\". These \"list\" operations require pagination; by requesting smaller subsets of data, API clients receive a response much faster than requesting the entire, potentially large, data set.  All \"list\" operations follow the same pattern with the following parameters:  * `page_size` –  client-provided max number of items per page, only valid on the first request. * `page_token` –  server-generated token used for traversing through the result set.  A paginated response may include any of the following pagination links. API clients may follow the respective link to page forward or backward through the result set as desired.  | [Link Relation](https://www.iana.org/assignments/link-relations/link-relations.xml) | Description |---------|--------------------------------------- | `next`  | A link to the next page of results. A response that does not contain a next link does not have further data to fetch. | `prev`  | A link to the previous page of results. A response that does not contain a prev link has no previous data. This link is **optional** for collections that cannot be traversed backward. | `first` | A link to the first page of results. This link is **optional** for collections that cannot be indexed directly to a given page. | `last`  | A link to the last page of results. This link is **optional** for collections that cannot be indexed directly to a given page.  API clients must treat pagination links and the `page_token` parameter in particular as an opaque string.   An example paginated list response may look like ``` {     \"api_version\": \"v2\",     \"kind\": \"KafkaClusterList\",     \"metadata\": {         \"next\": \"https://api.confluent.cloud/kafka-clusters?page_token=ABCDEFGHIJKLMNOP1234567890\"     }     \"data\": [         {             \"metadata\": {                 \"id\": \"lkc-abc123\",                 \"self\": \"https://api.confluent.cloud/kafka-clusters/lkc-abc123\",                 \"resource_name\": \"crn://confluent.cloud/kafka=lkc-abc123\",             }             \"spec\": {                 \"display_name\": \"My Kafka Cluster\",                 <snip>             },             \"status\": {                 \"phase\": \"RUNNING\",                 <snip>             }         },         <snip>     ] } ```  # Rate Limiting  To protect the stability of the API and keep it available to all users, Confluent employs multiple safeguards. If you send too many requests in quick succession or perform too many concurrent operations, you may be throttled or have your request rejected with an error.  When a rate limit is breached, an HTTP `429 Too Many Requests` error is returned. <!-- with the following header:  | Header                  | Description |-------------------------|---------------------------------------- | `Retry-After`           | The number of seconds to wait until the rate limit window resets. Only sent when the rate limit is reached.  --> <!-- TODO make this true  Confluent enforces multiple kinds of limits, including request rate and concurrency limits, both per user and organization wide. Unauthenticated requests are associated with the originating IP address, and not the user making requests.  -->  Integrations should gracefully handle these limits by watching for `429` error responses and building in a retry mechanism. This mechanism should follow a capped exponential backoff policy to prevent [retry amplification](https://landing.google.com/sre/sre-book/chapters/addressing-cascading-failures/) (\"retry storms\") and also introduce some randomness (\"jitter\") to avoid the [thundering herd effect](https://en.wikipedia.org/wiki/Thundering_herd_problem).   If you’re running into this error and think you need a higher rate limit, contact Confluent at [support@confluent.io](mailto:support@confluent.io).  # Identifiers and URLs  Most resources have multiple identifiers: * `id` is the \"natural identifier\" for an object. It is only unique within its parent resource.   The `id` is unique across time: the ID will not be reclaimed and reused after an object is deleted. * `resource_name` is a Uniform Resource Identifier (URI) that is globally unique across all resources.   This encompasses all parent resource `kind`s and `id`s necessary to uniquely identify a particular   instance of this object `kind`. Because it uses object `id`s, the CRN will not be reclaimed and   reused after an object is deleted. It is represented as a Confluent Resource Name (see below).  * `self` is a Uniform Resource Locator (URL) at which an object can be addressed.   This URL encodes the service location, API version, and other particulars necessary to   locate the resource at a point in time.  To see how these relate to each other, consider `KafkaBroker` with `broker.id=2` in a `KafkaCluster` in Confluent Cloud identified as `lkc-xsi8201`. In such an example, the `KafkaBroker` has `id=2`, the `resource_name` is `crn://confluent.cloud/kafka=lkc-xsi8201/broker=2` and the `self` URL may be something like `https://pkc-8wlk2n.us-west-2.aws.confluent.cloud`. Note that different identifiers carry different information for different purposes, but the `resource_name` is the most complete and canonical identifier.  ## Confluent Resource Names (CRNs)  *Confluent Resource Names* (CRNs) are used to uniquely identify all Confluent resources.  A CRN is a valid URI having an \"authority\" of `confluent.cloud` or a self-managed <a href=\"https://docs.confluent.io/current/security/rbac/configure-mds/index.html\" target=\"_blank\"> metadata service URL</a>, followed by the minimal hierarchical set of key-value pairs necessary to uniquely identify a resource.  Here are some examples for basic resources in Confluent Cloud:  | Resource            | Example CRN                                                 | |---------------------|-------------------------------------------------------------| | Environment         | crn://confluent.cloud/environment=e-xyz1234                 | | User                | crn://confluent.cloud/user=u-rst9876                        | | API Key             | crn://confluent.cloud/api-key=ABCDEFG9876543210             | | Service Account     | crn://confluent.cloud/service-account=sa-lmn2468            | | Service Account Key | crn://confluent.cloud/service-account-key=HIJKLMN0123456789 | | Kafka Cluster       | crn://confluent.cloud/kafka=lkc-abc123                      |  # Data Types  ### Primitive Types  | Data Type  | Representation |------------|--------------------- | Integers   | Each API may specify the type as `int32` or `int64`. Note that many languages, including JavaScript, are limited to a max size of approx `2**53` and don't correctly handle large `int64` values with their default JSON parser. | Dates      | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. UTC timezones are assumed, unless otherwise given. | Times      | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. UTC timezones are assumed, unless otherwise given. | Durations  | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. | Periods    | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. UTC timezones are assumed, unless otherwise given. | Ranges     | All ranges are represented using half-open intervals with naming conventions like `[start_XXX, end_XXX)` such as `[start_time, end_time)`. | Enums      | Most APIs use <a href=\"https://opensource.zalando.com/restful-api-guidelines/#112\" target=\"_blank\">`x-extensible-enum`</a> as an open-ended list of values. This improves compatibility compared with a standard `enum` which by definition represents a closed set. All enums have a `0`-valued entry which either serves as the default for common cases, or represents `UNSPECIFIED` when no default exists and results in an error.  <!-- TODO ### Standard Objects  | Money Object | https://schema.org/MonetaryAmount or https://opensource.zalando.com/restful-api-guidelines/#173 | Price Specification | https://schema.org/PriceSpecification -> https://schema.org/UnitPriceSpecification and https://schema.org/PaymentChargeSpecification -->  ### Standard Properties  Confluent uses this set of standard properties to ensure common concepts use the same name and semantics across different APIs.  | Name             | Description |------------------|------------------------------------------ | **api_version**  | Many API objects have an `api_version` field indicating their API version. See the [Object Model](#section/Object-Model). | **kind**         | Many API objects have a `kind` field indicating the kind of object it is. See the [Object Model](#section/Object-Model). | **id**           | Many objects in the API will have an identifier, indicated via its `id` field, and should be treated as an opaque string unless otherwise specified. See the [Object Model](#section/Object-Model). | **name**         | Objects which support a client-provided unique identifier instead of a generated `id` will indicate this identifier via its `name` field. | **display_name** | The human-readable display name of an API object. | **title**        | The official name of an API object, such as a company name. It should be treated as the formal version of `display_name`. | **description**  | One or more paragraphs of text description of an entity. | **created_at**   | The date and time the object was created, represented as a string in <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> format. | **updated_at**   | The date and time the object was last modified, represented as a string in <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> format. | **deleted_at**   | If present, the date and time after which the object was/will be deleted, represented as a string in <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> format. | **page_token**   | The pagination token in the List request. See [Pagination](#section/Pagination). | **page_size**    | The pagination size in the List request. See [Pagination](#section/Pagination). | **total_size**   | The total count of items in the list irrespective of pagination. See [Pagination](#section/Pagination). | **spec**         | The _desired state_ specification of the resource, as observed by Confluent Cloud. | **status**       | The current _actual state_ of the resource, as observed by Confluent Cloud.  # Versioning  Confluent APIs ensure stability for your integrations by avoiding the introduction of breaking changes to customers unexpectedly. Confluent will make non-breaking API changes without advance notice. Thus, API clients **must**  follow the [Compatibility Policy](#section/Versioning/Compatibility-Policy) below to ensure your ingtegration remains stable. All APIs follow the API Lifecycle Policy described below, which describes the guarantees API clients can rely on.  Breaking changes will be [widely communicated](#communication) in advance in accordance with the Confluent [Deprecation Policy](#section/Versioning/Deprecation-Policy). Confluent will provide  timelines and a migration path for all API changes, where available. Be sure to subscribe to one or more [communication channels](#communication) so you don't miss any updates!  One exception to these guidelines is for critical security issues. Confluent will take any necessary actions to mitigate any critical security issue as soon as possible, which may include disabling the vulnerable functionality until a proper solution is available.  Do not consume any Confluent API unless it is documented in the API Reference. All undocumented endpoints should be considered private, subject to change without notice, and not covered by any agreements.  > Note: The version in the URL (e.g. \"v1\" or \"v2\") is not a \"major version\" in the [Semantic Versioning](https://semver.org/) sense. It is a \"generational version\" or \"meta version\", as seen in APIs like <a href=\"https://developer.github.com/v3/versions/\" target=\"_blank\">Github API</a> or the <a href=\"https://stripe.com/docs/api/versioning\" target=\"_blank\">Stripe API</a>.  ## Known Issues  During the Early Access and Preview periods, we have a few known issues.  | Issue          | Description                                                                   | Proposed Resolution |----------------|-------------------------------------------------------------------------------|----------------------------------------------------- | Quota Exceeded | Some \"Quota Exceeded\" errors will be returned as HTTP 400 instead of HTTP 402 | Return 402 consistently for \"Quota Exceeded\" errors   ## API Lifecycle Policy  The following status labels are applicable to APIs, features, and SDK versions, based on the current support status of each:  * **Early Access** – May change at any time. Not recommended for production usage. Not officially supported by   Confluent. Intended for user feedback only. Users must be granted explicit access to the API by Confluent. * **Preview** – Unlikely to change between Preview and General Availability. Not recommended for production usage.   Officially supported by Confluent for non-production usage. For Closed Previews, users must be granted explicit   access to the API by Confluent. * **Generally Available (GA)** – Will not change at short notice. Recommended for production usage.   Officially supported by Confluent for non-production and production usage. * **Deprecated** – No longer supported. Will be removed in the future at the announced date.   Use is discouraged and migration following the upgrade guide is recommended. * **Sunset** – Removed, and no longer supported or available.  Resources, operations, and individual fields in the <a href=\"./openapi/v3\" target=\"_blank\">OpenAPI spec</a> will be annotated with `x-lifecycle-stage`, `x-deprecated-at`, and `x-sunset-at`. These annotations will appear in the corresponding API Reference Documentation. An API is \"Generally Available\" unless explicitly marked otherwise.  ## Compatibility Policy  Confluent Cloud APIs are governed by <a href=\"https://docs.confluent.io/current/cloud/limits.html#upgrade-policy\" target=\"_blank\"> Confluent Cloud Upgrade Policy</a>, which means that backward incompatible changes and deprecations will be made approximately once per year, and 180 days notice will be provided via email to all registered Confluent Cloud users.  ### Backward Compatibility  > *An API version is backward compatible if a program written against the previous version of the > API will continue to work the same way, without modification, against this version of the API.*  Confluent considers the following changes to be backward compatible:  * Adding new API resources. * Adding new optional parameters to existing API requests (e.g., query string). * Adding new properties to existing API resources (e.g., request body). * Changing the order of properties in existing API responses. * Changing the length or format of object IDs or other opaque strings.   * Unless otherwise documented, you can safely assume object IDs generated by Confluent will never exceed 255     characters, but you should be able to handle IDs of up to that length. If you're using MySQL,     for example, you should store IDs in a `VARCHAR(255) COLLATE utf8_bin` column.   * This includes adding or removing fixed prefixes (such as `lkc-` on Kafka cluster IDs).   * This includes API keys, API tokens, and similar authentication mechanisms.   * This includes all strings described as \"opaque\" in the docs, such as pagination cursors. * Adding new API event types. * Adding new properties to existing API event types. * Omitting properties with null values from existing API responses.  ### Forward Compatibility  > *An API version is forward compatible if a program written against the next version of the API > will continue to work the same way, without modification, against this version of the API.*  In other words, a forward compatible API will accept input intended for a later version of itself.  Confluent does not guarantee the forward compatibility of the APIs, but Confluent does generally follow the guidelines given by the [Robustness principle](https://en.wikipedia.org/wiki/Robustness_principle). This means that the API determines what to do with a request based only on the parts that it recognizes.  This is often referred to as the MUST IGNORE rule.  * Request parameters that are not recognized will be ignored (e.g., query string). * Request properties that are not recognized will be ignored (e.g., request body). * Request metadata that are not recognized will be ignored (e.g., request headers).  API clients must also follow the MUST IGNORE rule.  * Response properties that are not recognized must be ignored (e.g., response body). * Response metadata that are not recognized must be ignored (e.g., response headers).  Additionally, there is a more subtle related rule called the MUST FORWARD rule. Any parts of a request that an API doesn't recognize must be forwarded unchanged.  * Response properties that are not recognized must be included in any input subsequent updates (e.g., request body)   * This includes future `PUT` requests in a read/modify/write operation.     (This isn't required for `PATCH` partial updates, which is why Confluent APIs use `PATCH`.) * Event processors must not strip unknown properties before forwarding messages.  ### Client Responsibilities  * Resource and rate limits, and the default and maximum sizes of paginated data **are not**   considered part of the API contract and may change (possibly dynamically). It is the client's   responsibility to read the road signs and obey the speed limit. * If a property has a primitive type and the API documentation does not explicitly limit its   possible values, clients **must not** assume the values are constrained to a particular set   of possible responses. * If a property of an object is not explicitly declared as mandatory in the API, clients   **must not** assume it will be present. * A resource **may** be modified to return a \"redirection\" response (e.g. `301`, `307`) instead of   directly returning the resource. Clients **must** handle HTTP-level redirects, and respect HTTP   headers (e.g. `Location`).  ## Deprecation Policy  Confluent will announce deprecations at least 180 days in advance of a breaking change and will continue to maintain the deprecated APIs in their original form during this time.  Exceptions to this policy apply in case of critical security vulnerabilities or functional defects.  ### Communication  When a deprecation is announced, the details and any relevant migration information will be available on one or more of the following channels:  * Lifecycle, deprecation and \"since\" annotations in the <a href=\"./openapi/v3\" target=\"_blank\">OpenAPI spec</a> * Announcements on the <a href=\"https://www.confluent.io/blog/\" target=\"_blank\">Developer Blog</a>,   <a href=\"https://confluentcommunity.slack.com\" target=\"_blank\">Community Slack</a>   (<a href=\"https://slackpass.io/confluentcommunity\" target=\"_blank\">join!</a>),   <a href=\"https://groups.google.com/forum/#!forum/confluent-platform\" target=\"_blank\">Google Group</a>,   the <a href=\"https://twitter.com/ConfluentInc\" target=\"_blank\">@ConfluentInc twitter</a>   account, and similar channels * Enterprise customers may receive information by email to their specified Confluent contact, if applicable.  <!-- TODO: ### Discoverability -->  # HTTP Guidelines  ## Status Codes  Confluent respects the meanings and behavior of HTTP status codes as defined in <a href=\"https://tools.ietf.org/html/rfc2616\">RFC2616</a> and elsewhere.  * Codes in the `2xx` range indicate success * Codes in the `3xx` range indicate redirection * Codes in the `4xx` range indicate an error caused by the client request   (e.g., a required parameter was omitted, an invalid cluster configuration was provided, etc.) * Codes in the `5xx` range indicate an error with Confluent's servers (these are rare)  The various HTTP status codes that might be returned are listed below.  | Code | Title                  | Description |------|------------------------|-------------------------------- | 200  | OK                     | Everything worked as expected. | 201  | Created                | The resource was created. Follow the `Location` header. | 204  | No Content             | Everything worked and there is no content to return. | 400  | Bad Request         | The request was unacceptable, often due to malformed syntax, or a missing or malformed parameter. | 401  | Unauthorized           | No valid credentials provided. or the credentials are unsuitable, invalid, or unauthorized. | 402  | Over Quota             | The request was valid, but you've exceeded your plan quota or limits. | 404  | Not Found              | The requested resource doesn't exist or you're unauthorized to know it exists. | 409  | Conflict               | The request conflicts with another request (perhaps it already exists or was based on a stale version of data). | 422  | Validation Failed      | The request was parsed correctly but failed some sort of validation. | 429  | Too Many Requests      | Too many requests hit the API too quickly. Confluent recommends an exponential backoff of your requests. | 500, 502, 503, 504 | Server Errors | Something went wrong on Confluent's end. (These are rare.)  This list is not exhaustive; other standard HTTP error codes may be used, including `304`, `307`, `308`, `405`, `406`, `408`, `410`, and `415`.  For more details, see https://httpstatuses.com.  <!--  ## Method Overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, Confluent API operations that require `PUT`, `PATCH`, and `DELETE` verbs will be inaccessible.  To avoid this issue, Confluent APIs support the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `PUT`, `PATCH`, and `DELETE` requests via a `POST` request.  For example, to call a Confluent `PATCH` resource via a `POST` request, you can include `X-HTTP-Method-Override: PATCH` as a header.  ## User Agent Required  Confluent API requests **should** include a valid `User-Agent` header. Requests with no `User-Agent` header may be rejected. You should use the name of your integration for the `User-Agent` header value and include contact information so that Confluent can contact you if there are problems.  > If your integration is acting as a proxy or gateway, you **should** forward the User-Agent > of the originating client with your API requests.  Here's a complete example:      User-Agent: CoolToolName/1.2.3 (https://example.org/CoolTool/; CoolTool@example.org) UsedBaseLibrary/2.1.0  The minimum user agent string is the integration name and version: `name/version`. You can string together multiple values in a space-separated list. The full syntax is:      name/version [(comments)] [name/version [(comments)]] [...]​  For the integration name, use a string (without whitespace) that clearly and meaningfully identifies your integration.  * Avoid ambiguous names: `Confluent-Integration`, `Kafka-Sink` * Use clear and meaningful names: `ABCTools-ToolName`, `StackStorm-Confluent-Plugin`  For the version, use a semantic version, build ID, commit hash, or other identifier that is updated automatically when you release a new version.  Wrap comments in parentheses `()` as a semi-colon separated list. Helpful comments to include:  * A public URL for your integration, such as a GitHub link or a page in your   docs site that describes the integration. * Contact information so that Confluent can easily reach the integration publisher. This   information from the user agent string will never be shared nor used by Confluent for   any purpose other than discussing the integration with its publisher.  If you provide an invalid `User-Agent` header, you may receive a `403 Forbidden` response.  --> 
 *
 * API version: 1.0.0
 * Contact: support@confluent.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	"encoding/json"
)

// InlineResponse2001Tasks struct for InlineResponse2001Tasks
type InlineResponse2001Tasks struct {
	// The ID of task.
	Id int32 `json:"id"`
	// The state of the connector.
	State string `json:"state"`
	// The worker ID of the connector.
	WorkerId string `json:"worker_id"`
	Msg *string `json:"msg,omitempty"`
}

// NewInlineResponse2001Tasks instantiates a new InlineResponse2001Tasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001Tasks(id int32, state string, workerId string, ) *InlineResponse2001Tasks {
	this := InlineResponse2001Tasks{}
	this.Id = id
	this.State = state
	this.WorkerId = workerId
	return &this
}

// NewInlineResponse2001TasksWithDefaults instantiates a new InlineResponse2001Tasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001TasksWithDefaults() *InlineResponse2001Tasks {
	this := InlineResponse2001Tasks{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse2001Tasks) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Tasks) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse2001Tasks) SetId(v int32) {
	o.Id = v
}

// GetState returns the State field value
func (o *InlineResponse2001Tasks) GetState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Tasks) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InlineResponse2001Tasks) SetState(v string) {
	o.State = v
}

// GetWorkerId returns the WorkerId field value
func (o *InlineResponse2001Tasks) GetWorkerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Tasks) GetWorkerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *InlineResponse2001Tasks) SetWorkerId(v string) {
	o.WorkerId = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *InlineResponse2001Tasks) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Tasks) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *InlineResponse2001Tasks) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *InlineResponse2001Tasks) SetMsg(v string) {
	o.Msg = &v
}

func (o InlineResponse2001Tasks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["worker_id"] = o.WorkerId
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001Tasks struct {
	value *InlineResponse2001Tasks
	isSet bool
}

func (v NullableInlineResponse2001Tasks) Get() *InlineResponse2001Tasks {
	return v.value
}

func (v *NullableInlineResponse2001Tasks) Set(val *InlineResponse2001Tasks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001Tasks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001Tasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001Tasks(val *InlineResponse2001Tasks) *NullableInlineResponse2001Tasks {
	return &NullableInlineResponse2001Tasks{value: val, isSet: true}
}

func (v NullableInlineResponse2001Tasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001Tasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
