# SqlV1StatementResultResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]interface{}** | A data property that contains an array of results. Each entry in the array is a separate result.  The value of &#x60;op&#x60; attribute (if present) represents the kind of change that a row can describe in a changelog:  &#x60;0&#x60;: represents &#x60;INSERT&#x60; (&#x60;+I&#x60;), i.e. insertion operation;  &#x60;1&#x60;: represents &#x60;UPDATE_BEFORE&#x60; (&#x60;-U&#x60;), i.e. update operation with the previous content of the updated row. This kind should occur together with &#x60;UPDATE_AFTER&#x60; for modelling an update that needs to retract the previous row first. It is useful in cases of a non-idempotent update, i.e., an update of a row that is not uniquely identifiable by a key;  &#x60;2&#x60;: represents &#x60;UPDATE_AFTER&#x60; (&#x60;+U&#x60;), i.e. update operation with new content of the updated row; This kind CAN occur together with &#x60;UPDATE_BEFORE&#x60; for modelling an update that needs to retract the previous row first or it describes an idempotent update, i.e., an update of a row that is uniquely identifiable by a key;  &#x60;3&#x60;: represents &#x60;DELETE&#x60; (&#x60;-D&#x60;), i.e. deletion operation;  Defaults to &#x60;0&#x60;.  ### Value encoding by column type  The schema gives each column&#39;s type, so the row carries only values. Scalars are JSON strings in SQL notation. A SQL &#x60;NULL&#x60; is a bare &#x60;null&#x60;.  | Schema &#x60;type&#x60;                              | Payload example                  | |--------------------------------------------|----------------------------------| | &#x60;CHAR&#x60;, &#x60;VARCHAR&#x60;                          | &#x60;\&quot;A string\&quot;&#x60;                     | | &#x60;BOOLEAN&#x60;                                  | &#x60;\&quot;TRUE\&quot;&#x60;                         | | &#x60;BINARY&#x60;, &#x60;VARBINARY&#x60;                      | &#x60;\&quot;x&#39;7f0203&#39;\&quot;&#x60;                    | | &#x60;TINYINT&#x60;, &#x60;SMALLINT&#x60;, &#x60;INTEGER&#x60;, &#x60;BIGINT&#x60; | &#x60;\&quot;42\&quot;&#x60;                           | | &#x60;FLOAT&#x60;, &#x60;DOUBLE&#x60;                          | &#x60;\&quot;1.1111112E7\&quot;&#x60;                  | | &#x60;DECIMAL&#x60;                                  | &#x60;\&quot;12.123\&quot;&#x60;                       | | &#x60;DATE&#x60;                                     | &#x60;\&quot;2023-04-06\&quot;&#x60;                   | | &#x60;TIME_WITHOUT_TIME_ZONE&#x60;                   | &#x60;\&quot;10:56:22.541\&quot;&#x60;                 | | &#x60;TIMESTAMP_WITHOUT_TIME_ZONE&#x60;              | &#x60;\&quot;2023-04-06 10:59:32.628\&quot;&#x60;      | | &#x60;TIMESTAMP_WITH_LOCAL_TIME_ZONE&#x60;           | &#x60;\&quot;2023-04-06 11:06:47.224\&quot;&#x60;      | | &#x60;INTERVAL_YEAR_MONTH&#x60;                      | &#x60;\&quot;+2000-02\&quot;&#x60;                     | | &#x60;INTERVAL_DAY_TIME&#x60;                        | &#x60;\&quot;+2 07:33:20.000\&quot;&#x60;              | | &#x60;ARRAY&#x60;                                    | &#x60;[\&quot;1\&quot;, \&quot;2\&quot;, \&quot;3\&quot;, null]&#x60;          | | &#x60;MULTISET&#x60;                                 | &#x60;[[\&quot;a\&quot;, \&quot;2\&quot;], [\&quot;b\&quot;, \&quot;1\&quot;]]&#x60;       | | &#x60;MAP&#x60;                                      | &#x60;[[\&quot;Bob\&quot;, \&quot;a\&quot;], [\&quot;Alice\&quot;, \&quot;b\&quot;]]&#x60; | | &#x60;ROW&#x60;, &#x60;STRUCTURED_TYPE&#x60;                   | &#x60;[\&quot;12\&quot;, \&quot;a\&quot;]&#x60;                    | | &#x60;VARIANT&#x60;                                  | &#x60;[11, \&quot;sensor-7\&quot;]&#x60;               |  Timestamps use a space separator, not &#x60;T&#x60;. &#x60;TIMESTAMP_WITH_LOCAL_TIME_ZONE&#x60; is rendered in the session time zone and carries no offset.  &#x60;ROW&#x60; is positional, since its field names are in the schema. &#x60;ARRAY&#x60; is a plain list. &#x60;MAP&#x60; is a list of &#x60;[key, value]&#x60; pairs and &#x60;MULTISET&#x60; a list of &#x60;[element, count]&#x60; pairs, so keys travel as data and may be &#x60;null&#x60;.  &#x60;&#x60;&#x60; schema:  {\&quot;type\&quot;: \&quot;ROW\&quot;, \&quot;fields\&quot;: [            {\&quot;name\&quot;: \&quot;id\&quot;,    \&quot;field_type\&quot;: {\&quot;type\&quot;: \&quot;INTEGER\&quot;}},            {\&quot;name\&quot;: \&quot;tags\&quot;,  \&quot;field_type\&quot;: {\&quot;type\&quot;: \&quot;ARRAY\&quot;,                                \&quot;element_type\&quot;: {\&quot;type\&quot;: \&quot;VARCHAR\&quot;}}},            {\&quot;name\&quot;: \&quot;props\&quot;, \&quot;field_type\&quot;: {\&quot;type\&quot;: \&quot;MAP\&quot;,                                \&quot;key_type\&quot;:   {\&quot;type\&quot;: \&quot;VARCHAR\&quot;},                                \&quot;value_type\&quot;: {\&quot;type\&quot;: \&quot;INTEGER\&quot;}}}]}  payload: [\&quot;12\&quot;, [\&quot;hot\&quot;, \&quot;cold\&quot;], [[\&quot;a\&quot;, \&quot;1\&quot;], [\&quot;b\&quot;, \&quot;2\&quot;]]] &#x60;&#x60;&#x60;  ### VARIANT values  A &#x60;VARIANT&#x60; schema node carries no element types, so the types live in the value and differ per row. Each node is a positional array starting with an integer type code, and the code&#39;s **sign** gives the node&#39;s shape:  | Code  | Node shape                      | Meaning                             | |-------|---------------------------------|-------------------------------------| | &#x60;&gt; 0&#x60; | &#x60;[code, value]&#x60;                 | a value                             | | &#x60;0&#x60;   | &#x60;[0]&#x60;                           | NULL, carries nothing               | | &#x60;&lt; 0&#x60; | &#x60;[code, metadataHex, valueHex]&#x60; | not decodable, raw bytes attached   |  | Code | Type    | Code | Type     | Code | Type            | |------|---------|------|----------|------|-----------------| | &#x60;-2&#x60; | INVALID | &#x60;4&#x60;  | TINYINT  | &#x60;10&#x60; | DECIMAL         | | &#x60;-1&#x60; | UNKNOWN | &#x60;5&#x60;  | SMALLINT | &#x60;11&#x60; | STRING          | | &#x60;0&#x60;  | NULL    | &#x60;6&#x60;  | INT      | &#x60;12&#x60; | DATE            | | &#x60;1&#x60;  | OBJECT  | &#x60;7&#x60;  | BIGINT   | &#x60;13&#x60; | TIMESTAMP       | | &#x60;2&#x60;  | ARRAY   | &#x60;8&#x60;  | FLOAT    | &#x60;14&#x60; | TIMESTAMP_LTZ   | | &#x60;3&#x60;  | BOOLEAN | &#x60;9&#x60;  | DOUBLE   | &#x60;15&#x60; | BYTES           |  Scalars use the same encoding as the column type of the same name above. &#x60;OBJECT&#x60; is an array of &#x60;[key, node]&#x60; pairs, &#x60;ARRAY&#x60; an array of nodes.  &#x60;&#x60;&#x60; [1, [   [\&quot;device\&quot;,  [11, \&quot;sensor-7\&quot;]],   [\&quot;meta\&quot;,    [1, [[\&quot;n\&quot;, [4, \&quot;7\&quot;]], [\&quot;ok\&quot;, [3, \&quot;TRUE\&quot;]]]]],   [\&quot;price\&quot;,   [10, \&quot;100.00\&quot;]],   [\&quot;seen_at\&quot;, [13, \&quot;2026-07-28 09:14:02.117000\&quot;]],   [\&quot;tags\&quot;,    [2, [[11, \&quot;hot\&quot;], [0]]]],   [\&quot;weird\&quot;,   [-1, \&quot;x&#39;0102&#39;\&quot;, \&quot;x&#39;7f2a&#39;\&quot;]] ]] &#x60;&#x60;&#x60;  | [optional] 

## Methods

### NewSqlV1StatementResultResults

`func NewSqlV1StatementResultResults() *SqlV1StatementResultResults`

NewSqlV1StatementResultResults instantiates a new SqlV1StatementResultResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementResultResultsWithDefaults

`func NewSqlV1StatementResultResultsWithDefaults() *SqlV1StatementResultResults`

NewSqlV1StatementResultResultsWithDefaults instantiates a new SqlV1StatementResultResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SqlV1StatementResultResults) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SqlV1StatementResultResults) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SqlV1StatementResultResults) SetData(v []interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SqlV1StatementResultResults) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


