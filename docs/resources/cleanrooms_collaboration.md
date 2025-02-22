---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cleanrooms_collaboration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a collaboration between AWS accounts that allows for secure data collaboration
---

# awscc_cleanrooms_collaboration (Resource)

Represents a collaboration between AWS accounts that allows for secure data collaboration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `creator_display_name` (String)
- `creator_member_abilities` (Set of String)
- `description` (String)
- `members` (Attributes List) (see [below for nested schema](#nestedatt--members))
- `name` (String)
- `query_log_status` (String)

### Optional

- `data_encryption_metadata` (Attributes) (see [below for nested schema](#nestedatt--data_encryption_metadata))
- `tags` (Attributes Set) An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `collaboration_identifier` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Required:

- `account_id` (String)
- `display_name` (String)
- `member_abilities` (Set of String)


<a id="nestedatt--data_encryption_metadata"></a>
### Nested Schema for `data_encryption_metadata`

Required:

- `allow_cleartext` (Boolean)
- `allow_duplicates` (Boolean)
- `allow_joins_on_columns_with_different_names` (Boolean)
- `preserve_nulls` (Boolean)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cleanrooms_collaboration.example <resource ID>
```
