# Share Links Guide

Share links allow you to create temporary or permanent URLs for sharing files with others without giving them access to your entire file system.

## Creating Share Links

### Via Web Interface

1. Navigate to the file you want to share
2. Click the 🔗 icon next to the file
3. (Optional) Set expiration time in hours
4. Click "Create Link"
5. Copy the generated link

### Via API

```bash
# Create a permanent share link
curl -X POST http://localhost:8080/api/share \
  -F "path=documents/report.pdf"

# Create a link that expires in 24 hours
curl -X POST http://localhost:8080/api/share \
  -F "path=documents/report.pdf" \
  -F "expiresIn=24"
```

Response:
```json
{
  "id": "a1b2c3d4e5f6g7h8",
  "url": "/share/a1b2c3d4e5f6g7h8",
  "link": {
    "id": "a1b2c3d4e5f6g7h8",
    "file_path": "documents/report.pdf",
    "expires_at": "2024-01-16T10:30:00Z",
    "downloads": 0,
    "created_at": "2024-01-15T10:30:00Z"
  }
}
```

## Using Share Links

Share the full URL with others:
```
http://your-server.com:8080/share/a1b2c3d4e5f6g7h8
```

When someone visits this link:
- The file will be downloaded automatically
- Download counter will increment
- If expired, they'll see an error message

## Managing Share Links

### List All Share Links

```bash
curl http://localhost:8080/api/shares
```

Response:
```json
[
  {
    "id": "a1b2c3d4e5f6g7h8",
    "file_path": "documents/report.pdf",
    "expires_at": "2024-01-16T10:30:00Z",
    "downloads": 5,
    "created_at": "2024-01-15T10:30:00Z"
  }
]
```

### Delete a Share Link

```bash
curl -X DELETE http://localhost:8080/api/share/a1b2c3d4e5f6g7h8
```

## Features

### Expiration

- Set expiration time in hours when creating the link
- Expired links are automatically deleted when accessed
- Leave empty for permanent links

### Download Tracking

- Each access increments the download counter
- View statistics via API
- Useful for monitoring file distribution

### Security

- Links use random 32-character IDs (hard to guess)
- No directory listing through share links
- Only the specific file is accessible
- Expired links are immediately inaccessible

## Use Cases

### 1. Temporary File Sharing

Share a file with a client that expires after 48 hours:

```bash
curl -X POST http://localhost:8080/api/share \
  -F "path=client-files/proposal.pdf" \
  -F "expiresIn=48"
```

### 2. Public Downloads

Create a permanent link for a public file:

```bash
curl -X POST http://localhost:8080/api/share \
  -F "path=public/software-v1.0.zip"
```

### 3. One-Time Downloads

Create a link with max downloads (future feature):

```bash
# This will be available in future versions
curl -X POST http://localhost:8080/api/share \
  -F "path=secret/document.pdf" \
  -F "maxDownloads=1"
```

### 4. Analytics

Track how many times a file has been downloaded:

```bash
curl http://localhost:8080/api/shares | jq '.[] | select(.file_path == "documents/report.pdf")'
```

## Best Practices

1. **Use Expiration**: Always set expiration for sensitive files
2. **Monitor Downloads**: Regularly check download counts
3. **Clean Up**: Delete old share links you no longer need
4. **Secure Your Server**: Use HTTPS in production
5. **Access Control**: Consider adding authentication for share link creation

## Integration Examples

### JavaScript

```javascript
async function createShareLink(filePath, expiresIn = null) {
  const formData = new FormData();
  formData.append('path', filePath);
  if (expiresIn) {
    formData.append('expiresIn', expiresIn);
  }
  
  const response = await fetch('/api/share', {
    method: 'POST',
    body: formData
  });
  
  const data = await response.json();
  return `${window.location.origin}${data.url}`;
}

// Usage
const shareUrl = await createShareLink('documents/report.pdf', 24);
console.log('Share URL:', shareUrl);
```

### Python

```python
import requests

def create_share_link(file_path, expires_in=None):
    data = {'path': file_path}
    if expires_in:
        data['expiresIn'] = expires_in
    
    response = requests.post(
        'http://localhost:8080/api/share',
        data=data
    )
    
    result = response.json()
    return f"http://localhost:8080{result['url']}"

# Usage
share_url = create_share_link('documents/report.pdf', 24)
print(f'Share URL: {share_url}')
```

## Troubleshooting

### Link Not Working

- Check if the link has expired
- Verify the file still exists
- Check server logs for errors

### Cannot Create Link

- Verify the file path is correct
- Check file permissions
- Ensure the server is running

### Downloads Not Counting

- Each unique access counts as one download
- Browser caching may affect counts
- Check if the link is being accessed correctly
