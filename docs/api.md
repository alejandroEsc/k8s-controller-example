# Protocol Documentation
<a name="top"/>

## Table of Contents

- [any.proto](#any.proto)
    - [Any](#google.protobuf.Any)
  
  
  
  

- [openapiv2.proto](#openapiv2.proto)
    - [Contact](#grpc.gateway.protoc_gen_swagger.options.Contact)
    - [ExternalDocumentation](#grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation)
    - [Info](#grpc.gateway.protoc_gen_swagger.options.Info)
    - [JSONSchema](#grpc.gateway.protoc_gen_swagger.options.JSONSchema)
    - [Operation](#grpc.gateway.protoc_gen_swagger.options.Operation)
    - [Schema](#grpc.gateway.protoc_gen_swagger.options.Schema)
    - [Scopes](#grpc.gateway.protoc_gen_swagger.options.Scopes)
    - [Scopes.ScopeEntry](#grpc.gateway.protoc_gen_swagger.options.Scopes.ScopeEntry)
    - [SecurityDefinitions](#grpc.gateway.protoc_gen_swagger.options.SecurityDefinitions)
    - [SecurityDefinitions.SecurityEntry](#grpc.gateway.protoc_gen_swagger.options.SecurityDefinitions.SecurityEntry)
    - [SecurityRequirement](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement)
    - [SecurityRequirement.SecurityRequirementEntry](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement.SecurityRequirementEntry)
    - [SecurityRequirement.SecurityRequirementValue](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement.SecurityRequirementValue)
    - [SecurityScheme](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme)
    - [Swagger](#grpc.gateway.protoc_gen_swagger.options.Swagger)
    - [Tag](#grpc.gateway.protoc_gen_swagger.options.Tag)
  
    - [JSONSchema.JSONSchemaSimpleTypes](#grpc.gateway.protoc_gen_swagger.options.JSONSchema.JSONSchemaSimpleTypes)
    - [SecurityScheme.Flow](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme.Flow)
    - [SecurityScheme.In](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme.In)
    - [SecurityScheme.Type](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme.Type)
    - [Swagger.SwaggerScheme](#grpc.gateway.protoc_gen_swagger.options.Swagger.SwaggerScheme)
  
  
  

- [descriptor.proto](#descriptor.proto)
    - [DescriptorProto](#google.protobuf.DescriptorProto)
    - [DescriptorProto.ExtensionRange](#google.protobuf.DescriptorProto.ExtensionRange)
    - [DescriptorProto.ReservedRange](#google.protobuf.DescriptorProto.ReservedRange)
    - [EnumDescriptorProto](#google.protobuf.EnumDescriptorProto)
    - [EnumDescriptorProto.EnumReservedRange](#google.protobuf.EnumDescriptorProto.EnumReservedRange)
    - [EnumOptions](#google.protobuf.EnumOptions)
    - [EnumValueDescriptorProto](#google.protobuf.EnumValueDescriptorProto)
    - [EnumValueOptions](#google.protobuf.EnumValueOptions)
    - [ExtensionRangeOptions](#google.protobuf.ExtensionRangeOptions)
    - [FieldDescriptorProto](#google.protobuf.FieldDescriptorProto)
    - [FieldOptions](#google.protobuf.FieldOptions)
    - [FileDescriptorProto](#google.protobuf.FileDescriptorProto)
    - [FileDescriptorSet](#google.protobuf.FileDescriptorSet)
    - [FileOptions](#google.protobuf.FileOptions)
    - [GeneratedCodeInfo](#google.protobuf.GeneratedCodeInfo)
    - [GeneratedCodeInfo.Annotation](#google.protobuf.GeneratedCodeInfo.Annotation)
    - [MessageOptions](#google.protobuf.MessageOptions)
    - [MethodDescriptorProto](#google.protobuf.MethodDescriptorProto)
    - [MethodOptions](#google.protobuf.MethodOptions)
    - [OneofDescriptorProto](#google.protobuf.OneofDescriptorProto)
    - [OneofOptions](#google.protobuf.OneofOptions)
    - [ServiceDescriptorProto](#google.protobuf.ServiceDescriptorProto)
    - [ServiceOptions](#google.protobuf.ServiceOptions)
    - [SourceCodeInfo](#google.protobuf.SourceCodeInfo)
    - [SourceCodeInfo.Location](#google.protobuf.SourceCodeInfo.Location)
    - [UninterpretedOption](#google.protobuf.UninterpretedOption)
    - [UninterpretedOption.NamePart](#google.protobuf.UninterpretedOption.NamePart)
  
    - [FieldDescriptorProto.Label](#google.protobuf.FieldDescriptorProto.Label)
    - [FieldDescriptorProto.Type](#google.protobuf.FieldDescriptorProto.Type)
    - [FieldOptions.CType](#google.protobuf.FieldOptions.CType)
    - [FieldOptions.JSType](#google.protobuf.FieldOptions.JSType)
    - [FileOptions.OptimizeMode](#google.protobuf.FileOptions.OptimizeMode)
    - [MethodOptions.IdempotencyLevel](#google.protobuf.MethodOptions.IdempotencyLevel)
  
  
  

- [annotations.proto](#annotations.proto)
  
  
    - [File-level Extensions](#annotations.proto-extensions)
    - [File-level Extensions](#annotations.proto-extensions)
    - [File-level Extensions](#annotations.proto-extensions)
    - [File-level Extensions](#annotations.proto-extensions)
  
  

- [pod.proto](#pod.proto)
    - [HealthStatus](#pod.HealthStatus)
    - [HealthStatusRequest](#pod.HealthStatusRequest)
    - [PodInfo](#pod.PodInfo)
    - [PodInfoRequest](#pod.PodInfoRequest)
  
  
  
    - [PodMessagingService](#pod.PodMessagingService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="any.proto"/>
<p align="right"><a href="#top">Top</a></p>

## any.proto



<a name="google.protobuf.Any"/>

### Any
`Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C&#43;&#43;.

Foo foo = ...;
Any any;
any.PackFrom(foo);
...
if (any.UnpackTo(&amp;foo)) {
...
}

Example 2: Pack and unpack a message in Java.

Foo foo = ...;
Any any = Any.pack(foo);
...
if (any.is(Foo.class)) {
foo = any.unpack(Foo.class);
}

Example 3: Pack and unpack a message in Python.

foo = Foo(...)
any = Any()
any.Pack(foo)
...
if any.Is(Foo.DESCRIPTOR):
any.Unpack(foo)
...

Example 4: Pack and unpack a message in Go

foo := &amp;pb.Foo{...}
any, err := ptypes.MarshalAny(foo)
...
foo := &amp;pb.Foo{}
if err := ptypes.UnmarshalAny(any, foo); err != nil {
...
}

The pack methods provided by protobuf library will by default use
&#39;type.googleapis.com/full.type.name&#39; as the type URL and the unpack
methods only use the fully qualified type name after the last &#39;/&#39;
in the type URL, for example &#34;foo.bar.com/x/y.z&#34; will yield type
name &#34;y.z&#34;.


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

package google.profile;
message Person {
string first_name = 1;
string last_name = 2;
}

{
&#34;@type&#34;: &#34;type.googleapis.com/google.profile.Person&#34;,
&#34;firstName&#34;: &lt;string&gt;,
&#34;lastName&#34;: &lt;string&gt;
}

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

{
&#34;@type&#34;: &#34;type.googleapis.com/google.protobuf.Duration&#34;,
&#34;value&#34;: &#34;1.212s&#34;
}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type_url | [string](#string) |  | A URL/resource name whose content describes the type of the serialized protocol buffer message.

For URLs which use the scheme `http`, `https`, or no scheme, the following restrictions and interpretations apply:

If no scheme is provided, `https` is assumed. The last segment of the URL&#39;s path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading &#34;.&#34; is not accepted). An HTTP GET on the URL must yield a [google.protobuf.Type][] value in binary format, or produce an error. Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.)

Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics. |
| value | [bytes](#bytes) |  | Must be a valid serialized protocol buffer of the above specified type. |





 

 

 

 



<a name="openapiv2.proto"/>
<p align="right"><a href="#top">Top</a></p>

## openapiv2.proto



<a name="grpc.gateway.protoc_gen_swagger.options.Contact"/>

### Contact
`Contact` is a representation of OpenAPI v2 specification&#39;s Contact object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#contactObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| url | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation"/>

### ExternalDocumentation
`ExternalDocumentation` is a representation of OpenAPI v2 specification&#39;s
ExternalDocumentation object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#externalDocumentationObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  |  |
| url | [string](#string) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.Info"/>

### Info
`Info` is a representation of OpenAPI v2 specification&#39;s Info object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#infoObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| terms_of_service | [string](#string) |  |  |
| contact | [Contact](#grpc.gateway.protoc_gen_swagger.options.Contact) |  |  |
| version | [string](#string) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.JSONSchema"/>

### JSONSchema
`JSONSchema` represents properties from JSON Schema taken, and as used, in
the OpenAPI v2 spec.

This includes changes made by OpenAPI v2.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject

See also: https://cswr.github.io/JsonSchema/spec/basic_types/,
https://github.com/json-schema-org/json-schema-spec/blob/master/schema.json

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| default | [string](#string) |  |  |
| multiple_of | [double](#double) |  |  |
| maximum | [double](#double) |  |  |
| exclusive_maximum | [bool](#bool) |  |  |
| minimum | [double](#double) |  |  |
| exclusive_minimum | [bool](#bool) |  |  |
| max_length | [uint64](#uint64) |  |  |
| min_length | [uint64](#uint64) |  |  |
| pattern | [string](#string) |  |  |
| max_items | [uint64](#uint64) |  |  |
| min_items | [uint64](#uint64) |  |  |
| unique_items | [bool](#bool) |  |  |
| max_properties | [uint64](#uint64) |  |  |
| min_properties | [uint64](#uint64) |  |  |
| required | [string](#string) | repeated |  |
| array | [string](#string) | repeated | Items in &#39;array&#39; must be unique. |
| type | [JSONSchema.JSONSchemaSimpleTypes](#grpc.gateway.protoc_gen_swagger.options.JSONSchema.JSONSchemaSimpleTypes) | repeated |  |






<a name="grpc.gateway.protoc_gen_swagger.options.Operation"/>

### Operation
`Operation` is a representation of OpenAPI v2 specification&#39;s Operation object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#operationObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [string](#string) | repeated |  |
| summary | [string](#string) |  |  |
| description | [string](#string) |  |  |
| external_docs | [ExternalDocumentation](#grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation) |  |  |
| operation_id | [string](#string) |  |  |
| consumes | [string](#string) | repeated |  |
| produces | [string](#string) | repeated |  |
| schemes | [string](#string) | repeated |  |
| deprecated | [bool](#bool) |  |  |
| security | [SecurityRequirement](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement) | repeated |  |






<a name="grpc.gateway.protoc_gen_swagger.options.Schema"/>

### Schema
`Schema` is a representation of OpenAPI v2 specification&#39;s Schema object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json_schema | [JSONSchema](#grpc.gateway.protoc_gen_swagger.options.JSONSchema) |  |  |
| discriminator | [string](#string) |  |  |
| read_only | [bool](#bool) |  |  |
| external_docs | [ExternalDocumentation](#grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation) |  |  |
| example | [.google.protobuf.Any](#grpc.gateway.protoc_gen_swagger.options..google.protobuf.Any) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.Scopes"/>

### Scopes
`Scopes` is a representation of OpenAPI v2 specification&#39;s Scopes object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#scopesObject

Lists the available scopes for an OAuth2 security scheme.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scope | [Scopes.ScopeEntry](#grpc.gateway.protoc_gen_swagger.options.Scopes.ScopeEntry) | repeated | Maps between a name of a scope to a short description of it (as the value of the property). |






<a name="grpc.gateway.protoc_gen_swagger.options.Scopes.ScopeEntry"/>

### Scopes.ScopeEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.SecurityDefinitions"/>

### SecurityDefinitions
`SecurityDefinitions` is a representation of OpenAPI v2 specification&#39;s
Security Definitions object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityDefinitionsObject

A declaration of the security schemes available to be used in the
specification. This does not enforce the security schemes on the operations
and only serves to provide the relevant details for each scheme.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security | [SecurityDefinitions.SecurityEntry](#grpc.gateway.protoc_gen_swagger.options.SecurityDefinitions.SecurityEntry) | repeated | A single security scheme definition, mapping a &#34;name&#34; to the scheme it defines. |






<a name="grpc.gateway.protoc_gen_swagger.options.SecurityDefinitions.SecurityEntry"/>

### SecurityDefinitions.SecurityEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [SecurityScheme](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.SecurityRequirement"/>

### SecurityRequirement
`SecurityRequirement` is a representation of OpenAPI v2 specification&#39;s
Security Requirement object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityRequirementObject

Lists the required security schemes to execute this operation. The object can
have multiple security schemes declared in it which are all required (that
is, there is a logical AND between the schemes).

The name used for each property MUST correspond to a security scheme
declared in the Security Definitions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security_requirement | [SecurityRequirement.SecurityRequirementEntry](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement.SecurityRequirementEntry) | repeated | Each name must correspond to a security scheme which is declared in the Security Definitions. If the security scheme is of type &#34;oauth2&#34;, then the value is a list of scope names required for the execution. For other security scheme types, the array MUST be empty. |






<a name="grpc.gateway.protoc_gen_swagger.options.SecurityRequirement.SecurityRequirementEntry"/>

### SecurityRequirement.SecurityRequirementEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [SecurityRequirement.SecurityRequirementValue](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement.SecurityRequirementValue) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.SecurityRequirement.SecurityRequirementValue"/>

### SecurityRequirement.SecurityRequirementValue
If the security scheme is of type &#34;oauth2&#34;, then the value is a list of
scope names required for the execution. For other security scheme types,
the array MUST be empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scope | [string](#string) | repeated |  |






<a name="grpc.gateway.protoc_gen_swagger.options.SecurityScheme"/>

### SecurityScheme
`SecurityScheme` is a representation of OpenAPI v2 specification&#39;s
Security Scheme object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securitySchemeObject

Allows the definition of a security scheme that can be used by the
operations. Supported schemes are basic authentication, an API key (either as
a header or as a query parameter) and OAuth2&#39;s common flows (implicit,
password, application and access code).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [SecurityScheme.Type](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme.Type) |  | Required. The type of the security scheme. Valid values are &#34;basic&#34;, &#34;apiKey&#34; or &#34;oauth2&#34;. |
| description | [string](#string) |  | A short description for security scheme. |
| name | [string](#string) |  | Required. The name of the header or query parameter to be used.

Valid for apiKey. |
| in | [SecurityScheme.In](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme.In) |  | Required. The location of the API key. Valid values are &#34;query&#34; or &#34;header&#34;.

Valid for apiKey. |
| flow | [SecurityScheme.Flow](#grpc.gateway.protoc_gen_swagger.options.SecurityScheme.Flow) |  | Required. The flow used by the OAuth2 security scheme. Valid values are &#34;implicit&#34;, &#34;password&#34;, &#34;application&#34; or &#34;accessCode&#34;.

Valid for oauth2. |
| authorization_url | [string](#string) |  | Required. The authorization URL to be used for this flow. This SHOULD be in the form of a URL.

Valid for oauth2/implicit and oauth2/accessCode. |
| token_url | [string](#string) |  | Required. The token URL to be used for this flow. This SHOULD be in the form of a URL.

Valid for oauth2/password, oauth2/application and oauth2/accessCode. |
| scopes | [Scopes](#grpc.gateway.protoc_gen_swagger.options.Scopes) |  | Required. The available scopes for the OAuth2 security scheme.

Valid for oauth2. |






<a name="grpc.gateway.protoc_gen_swagger.options.Swagger"/>

### Swagger
`Swagger` is a representation of OpenAPI v2 specification&#39;s Swagger object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#swaggerObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| swagger | [string](#string) |  |  |
| info | [Info](#grpc.gateway.protoc_gen_swagger.options.Info) |  |  |
| host | [string](#string) |  |  |
| base_path | [string](#string) |  |  |
| schemes | [Swagger.SwaggerScheme](#grpc.gateway.protoc_gen_swagger.options.Swagger.SwaggerScheme) | repeated |  |
| consumes | [string](#string) | repeated |  |
| produces | [string](#string) | repeated |  |
| security_definitions | [SecurityDefinitions](#grpc.gateway.protoc_gen_swagger.options.SecurityDefinitions) |  |  |
| security | [SecurityRequirement](#grpc.gateway.protoc_gen_swagger.options.SecurityRequirement) | repeated |  |
| external_docs | [ExternalDocumentation](#grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation) |  |  |






<a name="grpc.gateway.protoc_gen_swagger.options.Tag"/>

### Tag
`Tag` is a representation of OpenAPI v2 specification&#39;s Tag object.

See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#tagObject

TODO(ivucica): document fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | TODO(ivucica): Description should be extracted from comments on the proto service object. |
| external_docs | [ExternalDocumentation](#grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation) |  |  |





 


<a name="grpc.gateway.protoc_gen_swagger.options.JSONSchema.JSONSchemaSimpleTypes"/>

### JSONSchema.JSONSchemaSimpleTypes


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| ARRAY | 1 |  |
| BOOLEAN | 2 |  |
| INTEGER | 3 |  |
| NULL | 4 |  |
| NUMBER | 5 |  |
| OBJECT | 6 |  |
| STRING | 7 |  |



<a name="grpc.gateway.protoc_gen_swagger.options.SecurityScheme.Flow"/>

### SecurityScheme.Flow
Required. The flow used by the OAuth2 security scheme. Valid values are
&#34;implicit&#34;, &#34;password&#34;, &#34;application&#34; or &#34;accessCode&#34;.

| Name | Number | Description |
| ---- | ------ | ----------- |
| FLOW_INVALID | 0 |  |
| FLOW_IMPLICIT | 1 |  |
| FLOW_PASSWORD | 2 |  |
| FLOW_APPLICATION | 3 |  |
| FLOW_ACCESS_CODE | 4 |  |



<a name="grpc.gateway.protoc_gen_swagger.options.SecurityScheme.In"/>

### SecurityScheme.In
Required. The location of the API key. Valid values are &#34;query&#34; or &#34;header&#34;.

| Name | Number | Description |
| ---- | ------ | ----------- |
| IN_INVALID | 0 |  |
| IN_QUERY | 1 |  |
| IN_HEADER | 2 |  |



<a name="grpc.gateway.protoc_gen_swagger.options.SecurityScheme.Type"/>

### SecurityScheme.Type
Required. The type of the security scheme. Valid values are &#34;basic&#34;,
&#34;apiKey&#34; or &#34;oauth2&#34;.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_INVALID | 0 |  |
| TYPE_BASIC | 1 |  |
| TYPE_API_KEY | 2 |  |
| TYPE_OAUTH2 | 3 |  |



<a name="grpc.gateway.protoc_gen_swagger.options.Swagger.SwaggerScheme"/>

### Swagger.SwaggerScheme


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| HTTP | 1 |  |
| HTTPS | 2 |  |
| WS | 3 |  |
| WSS | 4 |  |


 

 

 



<a name="descriptor.proto"/>
<p align="right"><a href="#top">Top</a></p>

## descriptor.proto



<a name="google.protobuf.DescriptorProto"/>

### DescriptorProto
Describes a message type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| field | [FieldDescriptorProto](#google.protobuf.FieldDescriptorProto) | repeated |  |
| extension | [FieldDescriptorProto](#google.protobuf.FieldDescriptorProto) | repeated |  |
| nested_type | [DescriptorProto](#google.protobuf.DescriptorProto) | repeated |  |
| enum_type | [EnumDescriptorProto](#google.protobuf.EnumDescriptorProto) | repeated |  |
| extension_range | [DescriptorProto.ExtensionRange](#google.protobuf.DescriptorProto.ExtensionRange) | repeated |  |
| oneof_decl | [OneofDescriptorProto](#google.protobuf.OneofDescriptorProto) | repeated |  |
| options | [MessageOptions](#google.protobuf.MessageOptions) | optional |  |
| reserved_range | [DescriptorProto.ReservedRange](#google.protobuf.DescriptorProto.ReservedRange) | repeated |  |
| reserved_name | [string](#string) | repeated | Reserved field names, which may not be used by fields in the same message. A given name may only be reserved once. |






<a name="google.protobuf.DescriptorProto.ExtensionRange"/>

### DescriptorProto.ExtensionRange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int32](#int32) | optional |  |
| end | [int32](#int32) | optional |  |
| options | [ExtensionRangeOptions](#google.protobuf.ExtensionRangeOptions) | optional |  |






<a name="google.protobuf.DescriptorProto.ReservedRange"/>

### DescriptorProto.ReservedRange
Range of reserved tag numbers. Reserved tag numbers may not be used by
fields or extension ranges in the same message. Reserved ranges may
not overlap.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int32](#int32) | optional | Inclusive. |
| end | [int32](#int32) | optional | Exclusive. |






<a name="google.protobuf.EnumDescriptorProto"/>

### EnumDescriptorProto
Describes an enum type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| value | [EnumValueDescriptorProto](#google.protobuf.EnumValueDescriptorProto) | repeated |  |
| options | [EnumOptions](#google.protobuf.EnumOptions) | optional |  |
| reserved_range | [EnumDescriptorProto.EnumReservedRange](#google.protobuf.EnumDescriptorProto.EnumReservedRange) | repeated | Range of reserved numeric values. Reserved numeric values may not be used by enum values in the same enum declaration. Reserved ranges may not overlap. |
| reserved_name | [string](#string) | repeated | Reserved enum value names, which may not be reused. A given name may only be reserved once. |






<a name="google.protobuf.EnumDescriptorProto.EnumReservedRange"/>

### EnumDescriptorProto.EnumReservedRange
Range of reserved numeric values. Reserved values may not be used by
entries in the same enum. Reserved ranges may not overlap.

Note that this is distinct from DescriptorProto.ReservedRange in that it
is inclusive such that it can appropriately represent the entire int32
domain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int32](#int32) | optional | Inclusive. |
| end | [int32](#int32) | optional | Inclusive. |






<a name="google.protobuf.EnumOptions"/>

### EnumOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allow_alias | [bool](#bool) | optional | Set this option to true to allow mapping different tag names to the same value. |
| deprecated | [bool](#bool) | optional | Is this enum deprecated? Depending on the target platform, this can emit Deprecated annotations for the enum, or it will be completely ignored; in the very least, this is a formalization for deprecating enums. |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.EnumValueDescriptorProto"/>

### EnumValueDescriptorProto
Describes a value within an enum.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| number | [int32](#int32) | optional |  |
| options | [EnumValueOptions](#google.protobuf.EnumValueOptions) | optional |  |






<a name="google.protobuf.EnumValueOptions"/>

### EnumValueOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deprecated | [bool](#bool) | optional | Is this enum value deprecated? Depending on the target platform, this can emit Deprecated annotations for the enum value, or it will be completely ignored; in the very least, this is a formalization for deprecating enum values. |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.ExtensionRangeOptions"/>

### ExtensionRangeOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.FieldDescriptorProto"/>

### FieldDescriptorProto
Describes a field within a message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| number | [int32](#int32) | optional |  |
| label | [FieldDescriptorProto.Label](#google.protobuf.FieldDescriptorProto.Label) | optional |  |
| type | [FieldDescriptorProto.Type](#google.protobuf.FieldDescriptorProto.Type) | optional | If type_name is set, this need not be set. If both this and type_name are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP. |
| type_name | [string](#string) | optional | For message and enum types, this is the name of the type. If the name starts with a &#39;.&#39;, it is fully-qualified. Otherwise, C&#43;&#43;-like scoping rules are used to find the type (i.e. first the nested types within this message are searched, then within the parent, on up to the root namespace). |
| extendee | [string](#string) | optional | For extensions, this is the name of the type being extended. It is resolved in the same manner as type_name. |
| default_value | [string](#string) | optional | For numeric types, contains the original text representation of the value. For booleans, &#34;true&#34; or &#34;false&#34;. For strings, contains the default text contents (not escaped in any way). For bytes, contains the C escaped value. All bytes &gt;= 128 are escaped. TODO(kenton): Base-64 encode? |
| oneof_index | [int32](#int32) | optional | If set, gives the index of a oneof in the containing type&#39;s oneof_decl list. This field is a member of that oneof. |
| json_name | [string](#string) | optional | JSON name of this field. The value is set by protocol compiler. If the user has set a &#34;json_name&#34; option on this field, that option&#39;s value will be used. Otherwise, it&#39;s deduced from the field&#39;s name by converting it to camelCase. |
| options | [FieldOptions](#google.protobuf.FieldOptions) | optional |  |






<a name="google.protobuf.FieldOptions"/>

### FieldOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ctype | [FieldOptions.CType](#google.protobuf.FieldOptions.CType) | optional | The ctype option instructs the C&#43;&#43; code generator to use a different representation of the field than it normally would. See the specific options below. This option is not yet implemented in the open source release -- sorry, we&#39;ll try to include it in a future version! |
| packed | [bool](#bool) | optional | The packed option can be enabled for repeated primitive fields to enable a more efficient representation on the wire. Rather than repeatedly writing the tag and type for each element, the entire array is encoded as a single length-delimited blob. In proto3, only explicit setting it to false will avoid using packed encoding. |
| jstype | [FieldOptions.JSType](#google.protobuf.FieldOptions.JSType) | optional | The jstype option determines the JavaScript type used for values of the field. The option is permitted only for 64 bit integral and fixed types (int64, uint64, sint64, fixed64, sfixed64). A field with jstype JS_STRING is represented as JavaScript string, which avoids loss of precision that can happen when a large value is converted to a floating point JavaScript. Specifying JS_NUMBER for the jstype causes the generated JavaScript code to use the JavaScript &#34;number&#34; type. The behavior of the default option JS_NORMAL is implementation dependent.

This option is an enum to permit additional types to be added, e.g. goog.math.Integer. |
| lazy | [bool](#bool) | optional | Should this field be parsed lazily? Lazy applies only to message-type fields. It means that when the outer message is initially parsed, the inner message&#39;s contents will not be parsed but instead stored in encoded form. The inner message will actually be parsed when it is first accessed.

This is only a hint. Implementations are free to choose whether to use eager or lazy parsing regardless of the value of this option. However, setting this option true suggests that the protocol author believes that using lazy parsing on this field is worth the additional bookkeeping overhead typically needed to implement it.

This option does not affect the public interface of any generated code; all method signatures remain the same. Furthermore, thread-safety of the interface is not affected by this option; const methods remain safe to call from multiple threads concurrently, while non-const methods continue to require exclusive access.

Note that implementations may choose not to check required fields within a lazy sub-message. That is, calling IsInitialized() on the outer message may return true even if the inner message has missing required fields. This is necessary because otherwise the inner message would have to be parsed in order to perform the check, defeating the purpose of lazy parsing. An implementation which chooses not to check required fields must be consistent about it. That is, for any particular sub-message, the implementation must either *always* check its required fields, or *never check its required fields, regardless of whether or not the message has been parsed. |
| deprecated | [bool](#bool) | optional | Is this field deprecated? Depending on the target platform, this can emit Deprecated annotations for accessors, or it will be completely ignored; in the very least, this is a formalization for deprecating fields. |
| weak | [bool](#bool) | optional | For Google-internal migration only. Do not use. |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.FileDescriptorProto"/>

### FileDescriptorProto
Describes a complete .proto file.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional | file name, relative to root of source tree |
| package | [string](#string) | optional | e.g. &#34;foo&#34;, &#34;foo.bar&#34;, etc. |
| dependency | [string](#string) | repeated | Names of files imported by this file. |
| public_dependency | [int32](#int32) | repeated | Indexes of the public imported files in the dependency list above. |
| weak_dependency | [int32](#int32) | repeated | Indexes of the weak imported files in the dependency list. For Google-internal migration only. Do not use. |
| message_type | [DescriptorProto](#google.protobuf.DescriptorProto) | repeated | All top-level definitions in this file. |
| enum_type | [EnumDescriptorProto](#google.protobuf.EnumDescriptorProto) | repeated |  |
| service | [ServiceDescriptorProto](#google.protobuf.ServiceDescriptorProto) | repeated |  |
| extension | [FieldDescriptorProto](#google.protobuf.FieldDescriptorProto) | repeated |  |
| options | [FileOptions](#google.protobuf.FileOptions) | optional |  |
| source_code_info | [SourceCodeInfo](#google.protobuf.SourceCodeInfo) | optional | This field contains optional information about the original source code. You may safely remove this entire field without harming runtime functionality of the descriptors -- the information is needed only by development tools. |
| syntax | [string](#string) | optional | The syntax of the proto file. The supported values are &#34;proto2&#34; and &#34;proto3&#34;. |






<a name="google.protobuf.FileDescriptorSet"/>

### FileDescriptorSet
The protocol compiler can output a FileDescriptorSet containing the .proto
files it parses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file | [FileDescriptorProto](#google.protobuf.FileDescriptorProto) | repeated |  |






<a name="google.protobuf.FileOptions"/>

### FileOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| java_package | [string](#string) | optional | Sets the Java package where classes generated from this .proto will be placed. By default, the proto package is used, but this is often inappropriate because proto packages do not normally start with backwards domain names. |
| java_outer_classname | [string](#string) | optional | If set, all the classes from the .proto file are wrapped in a single outer class with the given name. This applies to both Proto1 (equivalent to the old &#34;--one_java_file&#34; option) and Proto2 (where a .proto always translates to a single class, but you may want to explicitly choose the class name). |
| java_multiple_files | [bool](#bool) | optional | If set true, then the Java code generator will generate a separate .java file for each top-level message, enum, and service defined in the .proto file. Thus, these types will *not* be nested inside the outer class named by java_outer_classname. However, the outer class will still be generated to contain the file&#39;s getDescriptor() method as well as any top-level extensions defined in the file. |
| java_generate_equals_and_hash | [bool](#bool) | optional | This option does nothing. |
| java_string_check_utf8 | [bool](#bool) | optional | If set true, then the Java2 code generator will generate code that throws an exception whenever an attempt is made to assign a non-UTF-8 byte sequence to a string field. Message reflection will do the same. However, an extension field still accepts non-UTF-8 byte sequences. This option has no effect on when used with the lite runtime. |
| optimize_for | [FileOptions.OptimizeMode](#google.protobuf.FileOptions.OptimizeMode) | optional |  |
| go_package | [string](#string) | optional | Sets the Go package where structs generated from this .proto will be placed. If omitted, the Go package will be derived from the following: - The basename of the package import path, if provided. - Otherwise, the package statement in the .proto file, if present. - Otherwise, the basename of the .proto file, without extension. |
| cc_generic_services | [bool](#bool) | optional | Should generic services be generated in each language? &#34;Generic&#34; services are not specific to any particular RPC system. They are generated by the main code generators in each language (without additional plugins). Generic services were the only kind of service generation supported by early versions of google.protobuf.

Generic services are now considered deprecated in favor of using plugins that generate code specific to your particular RPC system. Therefore, these default to false. Old code which depends on generic services should explicitly set them to true. |
| java_generic_services | [bool](#bool) | optional |  |
| py_generic_services | [bool](#bool) | optional |  |
| php_generic_services | [bool](#bool) | optional |  |
| deprecated | [bool](#bool) | optional | Is this file deprecated? Depending on the target platform, this can emit Deprecated annotations for everything in the file, or it will be completely ignored; in the very least, this is a formalization for deprecating files. |
| cc_enable_arenas | [bool](#bool) | optional | Enables the use of arenas for the proto messages in this file. This applies only to generated classes for C&#43;&#43;. |
| objc_class_prefix | [string](#string) | optional | Sets the objective c class prefix which is prepended to all objective c generated classes from this .proto. There is no default. |
| csharp_namespace | [string](#string) | optional | Namespace for generated classes; defaults to the package. |
| swift_prefix | [string](#string) | optional | By default Swift generators will take the proto package and CamelCase it replacing &#39;.&#39; with underscore and use that to prefix the types/symbols defined. When this options is provided, they will use this value instead to prefix the types/symbols defined. |
| php_class_prefix | [string](#string) | optional | Sets the php class prefix which is prepended to all php generated classes from this .proto. Default is empty. |
| php_namespace | [string](#string) | optional | Use this option to change the namespace of php generated classes. Default is empty. When this option is empty, the package name will be used for determining the namespace. |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See the documentation for the &#34;Options&#34; section above. |






<a name="google.protobuf.GeneratedCodeInfo"/>

### GeneratedCodeInfo
Describes the relationship between generated code and its original source
file. A GeneratedCodeInfo message is associated with only one generated
source file, but may contain references to different source .proto files.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| annotation | [GeneratedCodeInfo.Annotation](#google.protobuf.GeneratedCodeInfo.Annotation) | repeated | An Annotation connects some span of text in generated code to an element of its generating .proto file. |






<a name="google.protobuf.GeneratedCodeInfo.Annotation"/>

### GeneratedCodeInfo.Annotation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [int32](#int32) | repeated | Identifies the element in the original source .proto file. This field is formatted the same as SourceCodeInfo.Location.path. |
| source_file | [string](#string) | optional | Identifies the filesystem path to the original source .proto. |
| begin | [int32](#int32) | optional | Identifies the starting offset in bytes in the generated code that relates to the identified object. |
| end | [int32](#int32) | optional | Identifies the ending offset in bytes in the generated code that relates to the identified offset. The end offset should be one past the last relevant byte (so the length of the text = end - begin). |






<a name="google.protobuf.MessageOptions"/>

### MessageOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message_set_wire_format | [bool](#bool) | optional | Set true to use the old proto1 MessageSet wire format for extensions. This is provided for backwards-compatibility with the MessageSet wire format. You should not use this for any other reason: It&#39;s less efficient, has fewer features, and is more complicated.

The message must be defined exactly as follows: message Foo { option message_set_wire_format = true; extensions 4 to max; } Note that the message cannot have any defined fields; MessageSets only have extensions.

All extensions of your type must be singular messages; e.g. they cannot be int32s, enums, or repeated messages.

Because this is an option, the above two restrictions are not enforced by the protocol compiler. |
| no_standard_descriptor_accessor | [bool](#bool) | optional | Disables the generation of the standard &#34;descriptor()&#34; accessor, which can conflict with a field of the same name. This is meant to make migration from proto1 easier; new code should avoid fields named &#34;descriptor&#34;. |
| deprecated | [bool](#bool) | optional | Is this message deprecated? Depending on the target platform, this can emit Deprecated annotations for the message, or it will be completely ignored; in the very least, this is a formalization for deprecating messages. |
| map_entry | [bool](#bool) | optional | Whether the message is an automatically generated map entry type for the maps field.

For maps fields: map&lt;KeyType, ValueType&gt; map_field = 1; The parsed descriptor looks like: message MapFieldEntry { option map_entry = true; optional KeyType key = 1; optional ValueType value = 2; } repeated MapFieldEntry map_field = 1;

Implementations may choose not to generate the map_entry=true message, but use a native map in the target language to hold the keys and values. The reflection APIs in such implementions still need to work as if the field is a repeated message field.

NOTE: Do not set the option in .proto files. Always use the maps syntax instead. The option should only be implicitly set by the proto compiler parser. |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.MethodDescriptorProto"/>

### MethodDescriptorProto
Describes a method of a service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| input_type | [string](#string) | optional | Input and output type names. These are resolved in the same way as FieldDescriptorProto.type_name, but must refer to a message type. |
| output_type | [string](#string) | optional |  |
| options | [MethodOptions](#google.protobuf.MethodOptions) | optional |  |
| client_streaming | [bool](#bool) | optional | Identifies if client streams multiple client messages |
| server_streaming | [bool](#bool) | optional | Identifies if server streams multiple server messages |






<a name="google.protobuf.MethodOptions"/>

### MethodOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deprecated | [bool](#bool) | optional | Is this method deprecated? Depending on the target platform, this can emit Deprecated annotations for the method, or it will be completely ignored; in the very least, this is a formalization for deprecating methods. |
| idempotency_level | [MethodOptions.IdempotencyLevel](#google.protobuf.MethodOptions.IdempotencyLevel) | optional |  |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.OneofDescriptorProto"/>

### OneofDescriptorProto
Describes a oneof.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| options | [OneofOptions](#google.protobuf.OneofOptions) | optional |  |






<a name="google.protobuf.OneofOptions"/>

### OneofOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.ServiceDescriptorProto"/>

### ServiceDescriptorProto
Describes a service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional |  |
| method | [MethodDescriptorProto](#google.protobuf.MethodDescriptorProto) | repeated |  |
| options | [ServiceOptions](#google.protobuf.ServiceOptions) | optional |  |






<a name="google.protobuf.ServiceOptions"/>

### ServiceOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deprecated | [bool](#bool) | optional | Is this service deprecated? Depending on the target platform, this can emit Deprecated annotations for the service, or it will be completely ignored; in the very least, this is a formalization for deprecating services. |
| uninterpreted_option | [UninterpretedOption](#google.protobuf.UninterpretedOption) | repeated | The parser stores options it doesn&#39;t recognize here. See above. |






<a name="google.protobuf.SourceCodeInfo"/>

### SourceCodeInfo
Encapsulates information about the original source file from which a
FileDescriptorProto was generated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [SourceCodeInfo.Location](#google.protobuf.SourceCodeInfo.Location) | repeated | A Location identifies a piece of source code in a .proto file which corresponds to a particular definition. This information is intended to be useful to IDEs, code indexers, documentation generators, and similar tools.

For example, say we have a file like: message Foo { optional string foo = 1; } Let&#39;s look at just the field definition: optional string foo = 1; ^ ^^ ^^ ^ ^^^ a bc de f ghi We have the following locations: span path represents [a,i) [ 4, 0, 2, 0 ] The whole field definition. [a,b) [ 4, 0, 2, 0, 4 ] The label (optional). [c,d) [ 4, 0, 2, 0, 5 ] The type (string). [e,f) [ 4, 0, 2, 0, 1 ] The name (foo). [g,h) [ 4, 0, 2, 0, 3 ] The number (1).

Notes: - A location may refer to a repeated field itself (i.e. not to any particular index within it). This is used whenever a set of elements are logically enclosed in a single code segment. For example, an entire extend block (possibly containing multiple extension definitions) will have an outer location whose path refers to the &#34;extensions&#34; repeated field without an index. - Multiple locations may have the same path. This happens when a single logical declaration is spread out across multiple places. The most obvious example is the &#34;extend&#34; block again -- there may be multiple extend blocks in the same scope, each of which will have the same path. - A location&#39;s span is not always a subset of its parent&#39;s span. For example, the &#34;extendee&#34; of an extension declaration appears at the beginning of the &#34;extend&#34; block and is shared by all extensions within the block. - Just because a location&#39;s span is a subset of some other location&#39;s span does not mean that it is a descendent. For example, a &#34;group&#34; defines both a type and a field in a single declaration. Thus, the locations corresponding to the type and field and their components will overlap. - Code which tries to interpret locations should probably be designed to ignore those that it doesn&#39;t understand, as more types of locations could be recorded in the future. |






<a name="google.protobuf.SourceCodeInfo.Location"/>

### SourceCodeInfo.Location



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [int32](#int32) | repeated | Identifies which part of the FileDescriptorProto was defined at this location.

Each element is a field number or an index. They form a path from the root FileDescriptorProto to the place where the definition. For example, this path: [ 4, 3, 2, 7, 1 ] refers to: file.message_type(3) // 4, 3 .field(7) // 2, 7 .name() // 1 This is because FileDescriptorProto.message_type has field number 4: repeated DescriptorProto message_type = 4; and DescriptorProto.field has field number 2: repeated FieldDescriptorProto field = 2; and FieldDescriptorProto.name has field number 1: optional string name = 1;

Thus, the above path gives the location of a field name. If we removed the last element: [ 4, 3, 2, 7 ] this path refers to the whole field declaration (from the beginning of the label to the terminating semicolon). |
| span | [int32](#int32) | repeated | Always has exactly three or four elements: start line, start column, end line (optional, otherwise assumed same as start line), end column. These are packed into a single field for efficiency. Note that line and column numbers are zero-based -- typically you will want to add 1 to each before displaying to a user. |
| leading_comments | [string](#string) | optional | If this SourceCodeInfo represents a complete declaration, these are any comments appearing before and after the declaration which appear to be attached to the declaration.

A series of line comments appearing on consecutive lines, with no other tokens appearing on those lines, will be treated as a single comment.

leading_detached_comments will keep paragraphs of comments that appear before (but not connected to) the current element. Each paragraph, separated by empty lines, will be one comment element in the repeated field.

Only the comment content is provided; comment markers (e.g. //) are stripped out. For block comments, leading whitespace and an asterisk will be stripped from the beginning of each line other than the first. Newlines are included in the output.

Examples:

optional int32 foo = 1; // Comment attached to foo. Comment attached to bar. optional int32 bar = 2;

optional string baz = 3; Comment attached to baz. Another line attached to baz.

Comment attached to qux.

Another line attached to qux. optional double qux = 4;

Detached comment for corge. This is not leading or trailing comments to qux or corge because there are blank lines separating it from both.

Detached comment for corge paragraph 2.

optional string corge = 5; Block comment attached to corge. Leading asterisks will be removed. Block comment attached to grault. optional int32 grault = 6;

ignored detached comments. |
| trailing_comments | [string](#string) | optional |  |
| leading_detached_comments | [string](#string) | repeated |  |






<a name="google.protobuf.UninterpretedOption"/>

### UninterpretedOption
A message representing a option the parser does not recognize. This only
appears in options protos created by the compiler::Parser class.
DescriptorPool resolves these when building Descriptor objects. Therefore,
options protos in descriptor objects (e.g. returned by Descriptor::options(),
or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
in them.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [UninterpretedOption.NamePart](#google.protobuf.UninterpretedOption.NamePart) | repeated |  |
| identifier_value | [string](#string) | optional | The value of the uninterpreted option, in whatever type the tokenizer identified it as during parsing. Exactly one of these should be set. |
| positive_int_value | [uint64](#uint64) | optional |  |
| negative_int_value | [int64](#int64) | optional |  |
| double_value | [double](#double) | optional |  |
| string_value | [bytes](#bytes) | optional |  |
| aggregate_value | [string](#string) | optional |  |






<a name="google.protobuf.UninterpretedOption.NamePart"/>

### UninterpretedOption.NamePart
The name of the uninterpreted option.  Each string represents a segment in
a dot-separated name.  is_extension is true iff a segment represents an
extension (denoted with parentheses in options specs in .proto files).
E.g.,{ [&#34;foo&#34;, false], [&#34;bar.baz&#34;, true], [&#34;qux&#34;, false] } represents
&#34;foo.(bar.baz).qux&#34;.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name_part | [string](#string) | required |  |
| is_extension | [bool](#bool) | required |  |





 


<a name="google.protobuf.FieldDescriptorProto.Label"/>

### FieldDescriptorProto.Label


| Name | Number | Description |
| ---- | ------ | ----------- |
| LABEL_OPTIONAL | 1 | 0 is reserved for errors |
| LABEL_REQUIRED | 2 |  |
| LABEL_REPEATED | 3 |  |



<a name="google.protobuf.FieldDescriptorProto.Type"/>

### FieldDescriptorProto.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_DOUBLE | 1 | 0 is reserved for errors. Order is weird for historical reasons. |
| TYPE_FLOAT | 2 |  |
| TYPE_INT64 | 3 | Not ZigZag encoded. Negative numbers take 10 bytes. Use TYPE_SINT64 if negative values are likely. |
| TYPE_UINT64 | 4 |  |
| TYPE_INT32 | 5 | Not ZigZag encoded. Negative numbers take 10 bytes. Use TYPE_SINT32 if negative values are likely. |
| TYPE_FIXED64 | 6 |  |
| TYPE_FIXED32 | 7 |  |
| TYPE_BOOL | 8 |  |
| TYPE_STRING | 9 |  |
| TYPE_GROUP | 10 | Tag-delimited aggregate. Group type is deprecated and not supported in proto3. However, Proto3 implementations should still be able to parse the group wire format and treat group fields as unknown fields. |
| TYPE_MESSAGE | 11 | Length-delimited aggregate. |
| TYPE_BYTES | 12 | New in version 2. |
| TYPE_UINT32 | 13 |  |
| TYPE_ENUM | 14 |  |
| TYPE_SFIXED32 | 15 |  |
| TYPE_SFIXED64 | 16 |  |
| TYPE_SINT32 | 17 | Uses ZigZag encoding. |
| TYPE_SINT64 | 18 | Uses ZigZag encoding. |



<a name="google.protobuf.FieldOptions.CType"/>

### FieldOptions.CType


| Name | Number | Description |
| ---- | ------ | ----------- |
| STRING | 0 | Default mode. |
| CORD | 1 |  |
| STRING_PIECE | 2 |  |



<a name="google.protobuf.FieldOptions.JSType"/>

### FieldOptions.JSType


| Name | Number | Description |
| ---- | ------ | ----------- |
| JS_NORMAL | 0 | Use the default type. |
| JS_STRING | 1 | Use JavaScript strings. |
| JS_NUMBER | 2 | Use JavaScript numbers. |



<a name="google.protobuf.FileOptions.OptimizeMode"/>

### FileOptions.OptimizeMode
Generated classes can be optimized for speed or code size.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SPEED | 1 | Generate complete code for parsing, serialization, |
| CODE_SIZE | 2 | etc.

Use ReflectionOps to implement these methods. |
| LITE_RUNTIME | 3 | Generate code using MessageLite and the lite runtime. |



<a name="google.protobuf.MethodOptions.IdempotencyLevel"/>

### MethodOptions.IdempotencyLevel
Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
or neither? HTTP based RPC implementation may choose GET verb for safe
methods, and PUT verb for idempotent methods instead of the default POST.

| Name | Number | Description |
| ---- | ------ | ----------- |
| IDEMPOTENCY_UNKNOWN | 0 |  |
| NO_SIDE_EFFECTS | 1 | implies idempotent |
| IDEMPOTENT | 2 | idempotent, but may have side effects |


 

 

 



<a name="annotations.proto"/>
<p align="right"><a href="#top">Top</a></p>

## annotations.proto


 

 


<a name="annotations.proto-extensions"/>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| openapiv2_swagger | Swagger | google.protobuf.FileOptions | 1042 | ID assigned by protobuf-global-extension-registry@google.com for grpc-gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they extend different descriptor messages. |
| openapiv2_schema | Schema | google.protobuf.MessageOptions | 1042 | ID assigned by protobuf-global-extension-registry@google.com for grpc-gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they extend different descriptor messages. |
| openapiv2_operation | Operation | google.protobuf.MethodOptions | 1042 | ID assigned by protobuf-global-extension-registry@google.com for grpc-gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they extend different descriptor messages. |
| openapiv2_tag | Tag | google.protobuf.ServiceOptions | 1042 | ID assigned by protobuf-global-extension-registry@google.com for grpc-gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they extend different descriptor messages. |

 

 



<a name="pod.proto"/>
<p align="right"><a href="#top">Top</a></p>

## pod.proto



<a name="pod.HealthStatus"/>

### HealthStatus
Replies on health


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isOK | [bool](#bool) |  | isOK is true if healthy |






<a name="pod.HealthStatusRequest"/>

### HealthStatusRequest
Request for health status






<a name="pod.PodInfo"/>

### PodInfo
Return for pod info request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="pod.PodInfoRequest"/>

### PodInfoRequest
Request for pod info





 

 

 


<a name="pod.PodMessagingService"/>

### PodMessagingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Health | [HealthStatusRequest](#pod.HealthStatusRequest) | [HealthStatus](#pod.HealthStatusRequest) | health status |
| podInfo | [PodInfoRequest](#pod.PodInfoRequest) | [PodInfo](#pod.PodInfoRequest) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

