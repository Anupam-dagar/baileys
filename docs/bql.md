# Baileys Query Language (bql)
Baileys Query Language (BQL) is an expressive query language utilized in the search API provided by the framework. It empowers you to extensively search for entities using a formatted query string, which is incorporated into the `filters` key of the search API.

## Format of a bql query
A bql query follows the format `Column.Op:value`, separated by `;`, where:
- `Column` is the name of the column in the table
- `Op` is the operator to be used for comparison. See [Operators](#Operators) section for more details on available operators.
- `value` is the value for comparison

## Sample bql query
A bql query on the entity `user` could be `name.like:john;age.gt:20;`. This effectively generates the SQL query `SELECT * FROM user WHERE name ILIKE 'john' AND age > 20`.

## Operators
BQL supports a variety of search operators for flexible data filtering. Here is a list of the allowed search operators:
1. `eq`: Equal to 
2. `ne`: Not equal to 
3. `isNull`: Is null 
4. `nn`: Is not null 
5. `gt`: Greater than 
6. `ge`: Greater than or equal to 
7. `lt`: Less than 
8. `le`: Less than or equal to 
9. `like`: Like (substring match)
10. `nl`: Not like (substring does not match)
11. `in`: In a list of values 
12. `nin`: Not in a list of values

## Usage on sample entity `user`
Assuming a user entity with columns like `id`, `name`, `age`, `email`, and `deleted_at`, here are sample BQL queries for each operator:
1. `id.eq:ee87a79b-2838-46fb-95e8-10f5ccd80f88;` - Get user with id `ee87a79b-2838-46fb-95e8-10f5ccd80f88`
2. `name.ne:john;` - Get all users whose name is not `john`.
3. `deleted_at.isNull;` - Get all users whose `deleted_at` is null.
4. `email.nn;` - Get all users whose `email` is not null.
5. `age.gt:20;` - Get all users whose `age` is greater than 20.
6. `age.ge:20;` - Get all users whose `age` is greater than or equal to 20.
7. `age.lt:20;` - Get all users whose `age` is less than 20.
8. `age.le:20;` - Get all users whose `age` is less than or equal to 20.
9. `name.like:john;` - Get all users whose `name` contains `john` (case-insensitive).
10. `name.nl:john;` - Get all users whose `name` does not contain `john` (case-insensitive).
11. `id.in:ee87a79b-2838-46fb-95e8-10f5ccd80f88,68506a38-55d4-4510-b9b7-2d96543f1852;` - Get all users whose `id` is in the list of ids.
12. `id.nin:ee87a79b-2838-46fb-95e8-10f5ccd80f88,68506a38-55d4-4510-b9b7-2d96543f1852;` - Get all users whose `id` is not in the list of ids. 

## Combining multiple bql queries
Multiple BQL queries can be combined using `;` to form a complex query using the `AND` operator. For example, `name.like:john;age.gt:20;` retrieves users whose name contains `john` and age is greater than `20`.

## Querying related entities (Joins)
For entities with relationships, such as the polls and poll_options entities, defining the relationship and querying related entities involves specific steps:
### Defining the relation in entity
Baileys requires custom tags (`join` - `sql join "ON" condition` and `tableName` - `table name of the related entity`) to define relations between entities. For example:
```go
type Poll struct {
	entity.BaseModel
	Title       string       `gorm:"column:title" json:"title"`
	PollOptions []PollOption `join:"join poll_options on poll_options.poll_id = polls.id" tableName:"poll_options"`
}
```

### Using bql to query related entities
To filter over nested entities, the `Column` part of the bql query is written as `NestedEntityFieldName-NestedEntityColumnName` where
- `NestedEntityFieldName`: is the golang name of the nested entity field in the parent entity (here `PollOptions`)
- `NestedEntityColumnName`: is the name of the column present in the nested entity (here `title`)

For example, to get all polls where one of the poll_option title is "Golang", the bql query will be `PollOptions-title.eq:Golang;`.

## Querying date time fields
When querying date-time columns, the date-time value should be in the ISO format `YYYY-MM-DDTHH:MM:SS.sssZ` (milliseconds and time zone are optional).

### Sample Query
To get all polls created after 6 PM on January 5, 2024, the bql query will be `created_at.gt:2024-01-05 18:00:00;`.
