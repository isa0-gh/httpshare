# Webhook Examples

## Creating a Webhook

### Using curl

```bash
curl -X POST http://localhost:8080/api/webhook \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://your-server.com/webhook",
    "events": ["upload", "delete", "rename", "move"],
    "active": true
  }'
```

### Using JavaScript

```javascript
fetch('http://localhost:8080/api/webhook', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    url: 'https://your-server.com/webhook',
    events: ['upload', 'delete', 'rename', 'move'],
    active: true
  })
})
.then(res => res.json())
.then(data => console.log('Webhook created:', data));
```

## Webhook Payload

When an event occurs, httpshare will send a POST request to your webhook URL with the following payload:

```json
{
  "event": "upload",
  "file_path": "documents/report.pdf",
  "timestamp": "2024-01-15T10:30:00Z",
  "details": {
    "filename": "report.pdf",
    "size": 1024000
  }
}
```

### Event Types

- **upload**: File uploaded
- **delete**: File deleted
- **rename**: File renamed
- **move**: File moved

### Event-Specific Details

#### Upload Event
```json
{
  "event": "upload",
  "file_path": "path/to/file.txt",
  "timestamp": "2024-01-15T10:30:00Z",
  "details": {
    "filename": "file.txt",
    "size": 1024
  }
}
```

#### Delete Event
```json
{
  "event": "delete",
  "file_path": "path/to/file.txt",
  "timestamp": "2024-01-15T10:30:00Z",
  "details": null
}
```

#### Rename Event
```json
{
  "event": "rename",
  "file_path": "path/to/oldname.txt",
  "timestamp": "2024-01-15T10:30:00Z",
  "details": {
    "newName": "newname.txt"
  }
}
```

#### Move Event
```json
{
  "event": "move",
  "file_path": "path/to/file.txt",
  "timestamp": "2024-01-15T10:30:00Z",
  "details": {
    "destination": "new/path/to/file.txt"
  }
}
```

## Example Webhook Server (Node.js)

```javascript
const express = require('express');
const app = express();

app.use(express.json());

app.post('/webhook', (req, res) => {
  const { event, file_path, timestamp, details } = req.body;
  
  console.log(`Event: ${event}`);
  console.log(`File: ${file_path}`);
  console.log(`Time: ${timestamp}`);
  console.log(`Details:`, details);
  
  // Your custom logic here
  // - Send notifications
  // - Update database
  // - Trigger other processes
  
  res.status(200).json({ received: true });
});

app.listen(3000, () => {
  console.log('Webhook server listening on port 3000');
});
```

## Example Webhook Server (Python)

```python
from flask import Flask, request, jsonify
from datetime import datetime

app = Flask(__name__)

@app.route('/webhook', methods=['POST'])
def webhook():
    data = request.json
    
    event = data.get('event')
    file_path = data.get('file_path')
    timestamp = data.get('timestamp')
    details = data.get('details')
    
    print(f"Event: {event}")
    print(f"File: {file_path}")
    print(f"Time: {timestamp}")
    print(f"Details: {details}")
    
    # Your custom logic here
    # - Send notifications
    # - Update database
    # - Trigger other processes
    
    return jsonify({'received': True}), 200

if __name__ == '__main__':
    app.run(port=3000)
```

## Testing Webhooks

You can use services like [webhook.site](https://webhook.site) or [requestbin.com](https://requestbin.com) to test your webhooks without setting up a server.

1. Go to webhook.site
2. Copy your unique URL
3. Create a webhook in httpshare with that URL
4. Perform file operations
5. See the webhook payloads in real-time

## Managing Webhooks

### List all webhooks
```bash
curl http://localhost:8080/api/webhooks
```

### Delete a webhook
```bash
curl -X DELETE http://localhost:8080/api/webhook/{webhook-id}
```

### Toggle webhook active status
```bash
curl -X POST http://localhost:8080/api/webhook/{webhook-id}/toggle
```

## Use Cases

1. **Backup Automation**: Trigger backups when files are uploaded
2. **Notifications**: Send Slack/Discord notifications on file changes
3. **Processing Pipeline**: Start processing jobs when files are uploaded
4. **Audit Logging**: Log all file operations to external system
5. **Sync**: Sync files to cloud storage on upload
6. **Security**: Alert on suspicious file operations
