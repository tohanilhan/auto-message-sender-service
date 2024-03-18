## Auto Message Sender Service API

This is a simpe API that controls the Auto Message Sending behavior and lists the sent messages.

## API USAGE

#### Get Sent Messages

```http
  GET /api/v1/messages
```

| Header | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `x-ins-api-auth-key` | `string` | **Mandatory**. API Key. |

##### Returns list of sent messages

#### Example Response:
```json
{
    "data": {
        "count": 3,
        "messages": [
            {
                "messageId": "05c36ad2-8c64-42ae-9339-18046da2943a",
                "sent_time": "2024-03-17 20:35:08"
            },
            {
                "messageId": "9780d953-f103-4190-b833-7c0ce3b10f85",
                "sent_time": "2024-03-17 20:35:13"
            },
            {
                "messageId": "04dd725d-94df-49ae-a6d0-52bbf2f676ee",
                "sent_time": "2024-03-17 20:35:18"
            },
        ]
    },
    "error": false,
    "message": "Messages retrieved successfully."
}
```

#### ON/OFF Automatic Message Sending

```http
  POST /api/v1/change-auto-sending/${option}
```

| Parameter | Type    | Description                       |
| :-------- | :------- | :-------------------------------- |
| `option`      | `string` | **Mandatory**. ON/OFF |

##### Changes the behavior of the automatic message sending functionality

#### Example Response
```json
{
  "error": false,
  "message": "Auto sending feature is now OFF."
}
```

## Documentation

Swagger Endpoint for API Documentation can be found on here after deploying the project: [Swagger Documentation](http://127.0.0.1:8787/api/v1/swagger/index.html)




  