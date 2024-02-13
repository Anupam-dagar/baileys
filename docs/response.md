# Response structure

Baileys simplifies API response management by offering a standardized response structure for all endpoints. This consistent format not only enhances clarity but also facilitates a seamless integration experience. Baileys' response handling methods empower developers to effortlessly send well-structured API responses, whether it's an error message, a successful request, or paginated data. This uniformity in response structure is a valuable feature that contributes to the efficiency and maintainability of your Baileys-powered application.

## Response structure
Following is a sample response for the baileys api
```json
{
  "status": {
    "code": 200,
    "message": "success",
    "type": "success",
    "totalCount": 1
  },
  "data": {
    // JSON structure of the provided struct or is a list of JSON structure of the provided struct basis the type of api called.
  },
}
```

## Convenient Response Handling
Baileys provides developers with methods for sending API responses in this structured format:
1. `ErrorResponse` - This method accepts the desired response status code and a Golang `error` object.
2. `SuccessResponse` - Use this method when your API request is successful. It requires the response status code and the data to be included under the `data` key.
3. `SuccessResponseWithCount` - Tailored for APIs with paginated responses, this function takes the response status code, the total count of the result set, and the relevant data to be sent.
