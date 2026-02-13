# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/note.proto](#proto_note-proto)
    - [CreateNoteRequest](#db_service-CreateNoteRequest)
    - [DeleteNoteRequest](#db_service-DeleteNoteRequest)
    - [GetAllNotesRequest](#db_service-GetAllNotesRequest)
    - [GetAllNotesResponse](#db_service-GetAllNotesResponse)
    - [GetNoteRequest](#db_service-GetNoteRequest)
    - [NoteResponse](#db_service-NoteResponse)
    - [UpdateNoteRequest](#db_service-UpdateNoteRequest)
  
    - [NotesService](#db_service-NotesService)
  
- [proto/user.proto](#proto_user-proto)
    - [CreateUserRequest](#db_service-CreateUserRequest)
    - [DeleteUserRequest](#db_service-DeleteUserRequest)
    - [GetAllUsersRequest](#db_service-GetAllUsersRequest)
    - [GetAllUsersResponse](#db_service-GetAllUsersResponse)
    - [GetUserRequest](#db_service-GetUserRequest)
    - [IsValidResponse](#db_service-IsValidResponse)
    - [UpdatePasswordRequest](#db_service-UpdatePasswordRequest)
    - [UpdateUserRequest](#db_service-UpdateUserRequest)
    - [User](#db_service-User)
    - [VerifyPasswordRequest](#db_service-VerifyPasswordRequest)
  
    - [UserService](#db_service-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_note-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/note.proto



<a name="db_service-CreateNoteRequest"></a>

### CreateNoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| content | [string](#string) |  |  |






<a name="db_service-DeleteNoteRequest"></a>

### DeleteNoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="db_service-GetAllNotesRequest"></a>

### GetAllNotesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="db_service-GetAllNotesResponse"></a>

### GetAllNotesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notes | [NoteResponse](#db_service-NoteResponse) | repeated |  |
| total_count | [int32](#int32) |  |  |






<a name="db_service-GetNoteRequest"></a>

### GetNoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="db_service-NoteResponse"></a>

### NoteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="db_service-UpdateNoteRequest"></a>

### UpdateNoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| new_label | [string](#string) | optional |  |
| new_content | [string](#string) | optional |  |





 

 

 


<a name="db_service-NotesService"></a>

### NotesService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateNote | [CreateNoteRequest](#db_service-CreateNoteRequest) | [NoteResponse](#db_service-NoteResponse) |  |
| GetNote | [GetNoteRequest](#db_service-GetNoteRequest) | [NoteResponse](#db_service-NoteResponse) |  |
| GetAllNotes | [GetAllNotesRequest](#db_service-GetAllNotesRequest) | [GetAllNotesResponse](#db_service-GetAllNotesResponse) |  |
| UpdateNote | [UpdateNoteRequest](#db_service-UpdateNoteRequest) | [NoteResponse](#db_service-NoteResponse) |  |
| DeleteNote | [DeleteNoteRequest](#db_service-DeleteNoteRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |

 



<a name="proto_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/user.proto



<a name="db_service-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="db_service-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="db_service-GetAllUsersRequest"></a>

### GetAllUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="db_service-GetAllUsersResponse"></a>

### GetAllUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#db_service-User) | repeated |  |
| total_count | [int32](#int32) |  |  |






<a name="db_service-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="db_service-IsValidResponse"></a>

### IsValidResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  |  |






<a name="db_service-UpdatePasswordRequest"></a>

### UpdatePasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| oldPassword | [string](#string) |  |  |
| newPassword | [string](#string) |  |  |






<a name="db_service-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |






<a name="db_service-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |






<a name="db_service-VerifyPasswordRequest"></a>

### VerifyPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| password | [string](#string) |  |  |





 

 

 


<a name="db_service-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#db_service-CreateUserRequest) | [User](#db_service-User) |  |
| GetAllUsers | [GetAllUsersRequest](#db_service-GetAllUsersRequest) | [GetAllUsersResponse](#db_service-GetAllUsersResponse) |  |
| GetUser | [GetUserRequest](#db_service-GetUserRequest) | [User](#db_service-User) |  |
| UpdateUser | [UpdateUserRequest](#db_service-UpdateUserRequest) | [User](#db_service-User) |  |
| UpdatePassword | [UpdatePasswordRequest](#db_service-UpdatePasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteUser | [DeleteUserRequest](#db_service-DeleteUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| VerifyPassword | [VerifyPasswordRequest](#db_service-VerifyPasswordRequest) | [IsValidResponse](#db_service-IsValidResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

