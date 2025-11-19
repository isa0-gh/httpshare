# httpshare
> This is mirror of the original repo https://gitlab.com/isa0/httpshare
![httpshare](https://gitlab.com/uploads/-/system/project/avatar/75901585/httpshare.png)
**httpshare** is a powerful web-based file manager written in Go. It allows you to quickly share, browse, and manage files over HTTP with minimal setup.

## Features

### File Management
- 📤 **File Upload**: Upload files directly from your browser with drag & drop support
- 🗑️ **Delete**: Remove files and folders with confirmation
- ✏️ **Rename**: Rename files and directories
- 📁 **Create Folders**: Create new directories on the fly
- 📦 **ZIP Download**: Download files and folders as ZIP archives
- ✂️ **Copy/Move**: Copy and move files with keyboard shortcuts (Ctrl+C/Ctrl+V)
- 🔄 **Bulk Operations**: Select multiple files for batch operations

### Browsing & Organization
- 🔍 **Search**: Real-time search for files and folders
- 🔤 **Sorting**: Sort by name, size, or modification date
- 📊 **File Information**: View file size, modification time, and permissions
- 🖼️ **Image Preview**: Automatic thumbnail display for images

### Content Preview & Media
- 👁️ **File Preview**: Preview text files, PDFs, videos, and audio files directly in the browser
- 📝 **Markdown Rendering**: View `.md` files with full formatting, syntax highlighting, and styling
- 🖼️ **Gallery Mode**: Browse images with a beautiful slideshow interface
  - Navigate with arrow keys or on-screen buttons
  - Click any image to open it in gallery mode
  - View image counter and filename
- 🎬 **Video Player**: Enhanced video player with subtitle support (.vtt files)
- 🎵 **Music Playlist**: Create and play audio playlists with auto-play
- 📄 **Office Documents**: Preview support for .docx, .xlsx, .pptx files
- Support for: `.txt`, `.md`, `.log`, `.json`, `.xml`, `.csv`, `.pdf`, `.mp4`, `.webm`, `.avi`, `.mov`, `.mkv`, `.mp3`, `.wav`, `.ogg`, `.flac`

### Sharing & Collaboration
- 🔗 **Share Links**: Generate shareable links for files
- ⏰ **Expiring Links**: Set expiration time for share links
- 📊 **Download Statistics**: Track download counts for shared files
- 💬 **File Comments**: Add comments to files for collaboration
- 🔔 **Webhooks**: Configure webhooks for file events (upload, delete, rename, move)

### Interface & Customization
- 🎨 **Theme Support**: Light, dark, and auto themes
- ⌨️ **Keyboard Shortcuts**: 
  - Ctrl/Cmd + U: Upload
  - Ctrl/Cmd + N: New folder
  - Ctrl/Cmd + A: Select all
  - Ctrl/Cmd + C: Copy selected
  - Ctrl/Cmd + V: Paste
  - Escape: Close modals
  - Arrow keys: Navigate gallery
- Modern, responsive web interface
- Mobile-friendly design
- Configurable port and directory
- Cross-platform: works on Windows, Linux, and macOS

## Quick Start

👉 **New to httpshare?** Check out the [Quick Start Guide](QUICKSTART.md) for a fast introduction!

## Installation

Make sure you have [Go](https://golang.org/dl/) installed. Then, run:

```bash
go install gitlab.com/isa0/httpshare@latest
````

This installs the `httpshare` binary in your Go `bin` directory (usually `$GOPATH/bin`).

## Usage

Start the file explorer with the following command:


```bash
httpshare [--port <port>] [--directory <path>]
```

### Options

* `--port`: Set the port to serve the files (default: `8080`)
* `--directory`: Specify the directory to serve (default: current directory)
* `--log`: Write logs into a file (default: disabled)

Then, open your web browser and navigate to:

```
http://localhost:<port>
```

You will see a modern web interface to browse, manage, and preview files in the current directory.

## API Endpoints

The application provides RESTful API endpoints for programmatic access:

### File Operations
- **POST** `/api/upload` - Upload files
- **DELETE** `/api/delete?path=<path>` - Delete files/folders
- **POST** `/api/rename` - Rename files/folders
- **POST** `/api/mkdir` - Create new directory
- **GET** `/api/search?q=<query>&path=<path>` - Search files
- **GET** `/api/download-zip?path=<path>` - Download file/folder as ZIP
- **POST** `/api/copy` - Copy files/folders
- **POST** `/api/move` - Move files/folders

### Sharing
- **POST** `/api/share` - Create a share link
- **GET** `/api/shares` - List all share links
- **DELETE** `/api/share/:id` - Delete a share link
- **GET** `/share/:id` - Access file via share link

### Comments
- **POST** `/api/comment` - Add a comment to a file
- **GET** `/api/comments?path=<path>` - Get comments for a file

### Webhooks
- **POST** `/api/webhook` - Create a webhook
- **GET** `/api/webhooks` - List all webhooks
- **DELETE** `/api/webhook/:id` - Delete a webhook
- **POST** `/api/webhook/:id/toggle` - Toggle webhook active status

### Examples

#### Upload a file using curl
```bash
curl -F "file=@myfile.txt" -F "path=." http://localhost:8080/api/upload
```

#### Delete a file
```bash
curl -X DELETE "http://localhost:8080/api/delete?path=myfile.txt"
```

#### Create a share link
```bash
curl -X POST -F "path=myfile.txt" -F "expiresIn=24" http://localhost:8080/api/share
```

#### Add a comment
```bash
curl -X POST -F "path=myfile.txt" -F "author=John" -F "content=Great file!" http://localhost:8080/api/comment
```

#### Create a webhook
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/webhook","events":["upload","delete"]}' \
  http://localhost:8080/api/webhook
```
You will see a simple web interface to explore and download files in the specified directory.

### Example

Serve files from `/home/user/Documents` on port 3000:

```bash
httpshare --port 3000 --directory /home/user/Documents
```

## Documentation

- 📖 [Quick Start Guide](QUICKSTART.md) - Get started in minutes
- ⌨️ [Keyboard Shortcuts](examples/keyboard-shortcuts.md) - Speed up your workflow
- 🔗 [Share Links Guide](examples/share-links-guide.md) - Learn about file sharing
- 🔔 [Webhook Examples](examples/webhook-example.md) - Integrate with other services
- 📝 [Changelog](CHANGELOG.md) - See what's new

## Contributing

Contributions are welcome! Feel free to submit issues or merge requests.

## License

This project is licensed under the MIT License.
