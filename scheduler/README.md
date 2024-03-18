
## How it works

- Operates at configured intervals.
- The automatic message sending feature is enabled by default, but it can be toggled using the API. When set to OFF, it will neither receive nor send any messages.
- Retrieves 2 messages in each interval.
- Concurrently sends these two messages to the configured Webhook API endpoint.
- Upon success, updates the relevant database entries and concurrently caches the sent messageId and sending time.