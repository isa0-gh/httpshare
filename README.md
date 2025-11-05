# httpshare

**httpshare** is a powerful web-based file manager written in Go. It allows you to quickly share, browse, and manage files over HTTP with minimal setup.


[![Releases](https://img.shields.io/github/v/release/isa0-gh/httpshare)](https://github.com/isa0-gh/httpshare/releases)

## Features

### File Management
- 📤 **File Upload**: Upload files directly from your browser
- 🗑️ **Delete**: Remove files and folders with confirmation
- ✏️ **Rename**: Rename files and directories
- 📁 **Create Folders**: Create new directories on the fly

### Browsing & Organization
- 🔍 **Search**: Real-time search for files and folders
- 🔤 **Sorting**: Sort by name, size, or modification date
- 📊 **File Information**: View file size, modification time, and permissions
- 🖼️ **Image Preview**: Automatic thumbnail display for images

### Content Preview
- 👁️ **File Preview**: Preview text files, PDFs, videos, and audio files directly in the browser
- 📝 **Markdown Rendering**: View `.md` files with full formatting, syntax highlighting, and styling
- 🖼️ **Gallery Mode**: Browse images with a beautiful slideshow interface
  - Navigate with arrow keys or on-screen buttons
  - Click any image to open it in gallery mode
  - View image counter and filename
- Support for: `.txt`, `.md`, `.log`, `.json`, `.xml`, `.csv`, `.pdf`, `.mp4`, `.webm`, `.mp3`, `.wav`

### Interface
- Lightweight and easy to use
- Modern, responsive web interface with dark theme
- Mobile-friendly design
- Configurable port
- Simple web interface for browsing and downloading files
- Configurable port and directory
- Cross-platform: works on Windows, Linux, and macOS

## Installation

Make sure you have [Go](https://golang.org/dl/) installed. Then, run:

```bash
go install github.com/isa0-gh/httpshare@latest
````

This installs the `httpshare` binary in your Go `bin` directory (usually `$GOPATH/bin`).

## Usage

**Start the file explorer with the following command:


```bash
httpshare [--port <port>] [--directory <path>]
```

### Options

* `--port`: Set the port to serve the files (default: `8080`)
* `--directory`: Specify the directory to serve (default: current directory)

Then, open your web browser and navigate to:

```
http://localhost:<port>
```

You will see a modern web interface to browse, manage, and preview files in the current directory.

## API Endpoints

The application provides RESTful API endpoints for programmatic access:

- **POST** `/api/upload` - Upload files
- **DELETE** `/api/delete?path=<path>` - Delete files/folders
- **POST** `/api/rename` - Rename files/folders
- **POST** `/api/mkdir` - Create new directory
- **GET** `/api/search?q=<query>&path=<path>` - Search files

### Example: Upload a file using curl
```bash
curl -F "file=@myfile.txt" -F "path=." http://localhost:8080/api/upload
```

### Example: Delete a file
```bash
curl -X DELETE "http://localhost:8080/api/delete?path=myfile.txt"
You will see a simple web interface to explore and download files in the specified directory.

### Example

Serve files from `/home/user/Documents` on port 3000:

```bash
httpshare --port 3000 --directory /home/user/Documents
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License.
