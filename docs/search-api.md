# Search Api Reference

The Search API, accessible for each entity in the framework, provides a RESTful interface for querying entities in the
database using the Baileys Query Language (BQL). The API supports filtering, sorting, pagination, and including related
nested entities in the response.

## Api Endpoint

`POST <base_path>/{entity}/search`

## Request Body

The request body for the search API is a JSON object containing the following keys:

- `filters`: (Optional) A BQL query string for filtering entities. Refer to [Baileys Query Language (BQL)](bql.md)for
  more details on BQL.
- `pageSize`: (Required) The maximum number of entities to return. Send -1 to retrieve all entities
  disabling pagination. 0 is not a valid value.
- `page`: (Optional) The page number of the entities to return. The default is 0 (page numbering starts from 0).
- `sort`: (Optional) The sorting order for the entities. The format is `Column:Order`, where `Column` is the column name
  and `Order` is either `asc` or `desc`. For example, `name:asc` sorts the entities by the `name` column in ascending
  order. Default is `id:asc`.
- `includes`: (Optional) A comma-separated list of related nested entities to include in the response. If not provided,
  data for related nested entities is set to null. Entity names must match those defined in the entity struct tags. (
  Refer to [Using bql to query related entities](bql.md#using-bql-to-query-related-entities) for details).

```json
{
  "filters": "title.like:sport;",
  "pageSize": 10,
  "page": 0,
  "sort": "name:asc",
  "includes": "poll_options"
}
```

## Response
The response from the Search API adheres to the standard response structure. For more details, refer to [Response Structure](response.md#response-structure).
