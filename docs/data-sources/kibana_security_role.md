---
subcategory: "Kibana"
layout: ""
page_title: "Elasticstack: elasticstack_kibana_security_role Data Source"
description: |-
  Retrieve a specific Kibana role. See https://www.elastic.co/guide/en/kibana/master/role-management-specific-api-get.html
---

# Data Source: elasticstack_kibana_security_role

Use this data source to get information about an existing Kibana role. 

## Example Usage

```terraform
provider "elasticstack" {
  elasticsearch {}
  kibana {}
}

data "elasticstack_kibana_security_role" "example" {
  name = "sample_role"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name for the role.

### Optional

- `description` (String) Description for the role
- `metadata` (String) Optional meta-data.

### Read-Only

- `elasticsearch` (Set of Object) Elasticsearch cluster and index privileges. (see [below for nested schema](#nestedatt--elasticsearch))
- `id` (String) The ID of this resource.
- `kibana` (Set of Object) The list of objects that specify the Kibana privileges for the role. (see [below for nested schema](#nestedatt--kibana))

<a id="nestedatt--elasticsearch"></a>
### Nested Schema for `elasticsearch`

Read-Only:

- `cluster` (Set of String)
- `indices` (Set of Object) (see [below for nested schema](#nestedobjatt--elasticsearch--indices))
- `remote_indices` (Set of Object) (see [below for nested schema](#nestedobjatt--elasticsearch--remote_indices))
- `run_as` (Set of String)

<a id="nestedobjatt--elasticsearch--indices"></a>
### Nested Schema for `elasticsearch.indices`

Read-Only:

- `field_security` (List of Object) (see [below for nested schema](#nestedobjatt--elasticsearch--indices--field_security))
- `names` (Set of String)
- `privileges` (Set of String)
- `query` (String)

<a id="nestedobjatt--elasticsearch--indices--field_security"></a>
### Nested Schema for `elasticsearch.indices.field_security`

Read-Only:

- `except` (Set of String)
- `grant` (Set of String)



<a id="nestedobjatt--elasticsearch--remote_indices"></a>
### Nested Schema for `elasticsearch.remote_indices`

Read-Only:

- `clusters` (Set of String)
- `field_security` (List of Object) (see [below for nested schema](#nestedobjatt--elasticsearch--remote_indices--field_security))
- `names` (Set of String)
- `privileges` (Set of String)
- `query` (String)

<a id="nestedobjatt--elasticsearch--remote_indices--field_security"></a>
### Nested Schema for `elasticsearch.remote_indices.field_security`

Read-Only:

- `except` (Set of String)
- `grant` (Set of String)




<a id="nestedatt--kibana"></a>
### Nested Schema for `kibana`

Read-Only:

- `base` (Set of String)
- `feature` (Set of Object) (see [below for nested schema](#nestedobjatt--kibana--feature))
- `spaces` (Set of String)

<a id="nestedobjatt--kibana--feature"></a>
### Nested Schema for `kibana.feature`

Read-Only:

- `name` (String)
- `privileges` (Set of String)
