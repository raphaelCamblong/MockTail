# MockTail: The Fast and Easy Mock API Server


```
███╗░░░███╗░█████╗░░█████╗░██╗░░██╗████████╗░█████╗░██╗██╗░░░░░
████╗░████║██╔══██╗██╔══██╗██║░██╔╝╚══██╔══╝██╔══██╗██║██║░░░░░
██╔████╔██║██║░░██║██║░░╚═╝█████═╝░░░░██║░░░███████║██║██║░░░░░
██║╚██╔╝██║██║░░██║██║░░██╗██╔═██╗░░░░██║░░░██╔══██║██║██║░░░░░
██║░╚═╝░██║╚█████╔╝╚█████╔╝██║░╚██╗░░░██║░░░██║░░██║██║███████╗
╚═╝░░░░░╚═╝░╚════╝░░╚════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═╝╚══════╝
```
MockTail is an ultra-fast, file-based mock server designed for developers who need to quickly test and simulate APIs. Its dynamic file system approach allows you to set up and serve mock endpoints almost instantly by simply creating files in a structured directory.

## No Complex Configuration Required

- **Drop in your mock data files**
- **Run the server**

And MockTail takes care of the rest.

This streamlined setup significantly accelerates your workflow, enabling you to focus on API development and debugging without the hassle of complex configurations. Enjoy seamless and efficient testing with MockTail, and get back to building faster than ever!

---

## Features

- **Fast and Lightweight**: Built with Go for high performance and low resource usage.
- **Dynamic Mocking**: Serve mock data dynamically based on request parameters and HTTP methods.
- **File-based Configuration**: Organize mock responses in a directory structure for easy management.
- **Customizable Responses**: Support for different HTTP status codes, query parameters, and response bodies.
- **CORS Support**: Middleware for Cross-Origin Resource Sharing (CORS).

## Coming Soon
 - **Docker Image**: for easy deployment and integration with CI/CD pipelines
- **Mock Data Generation**: for creating mock data files from OpenAPI specifications
- **lua scripting support**: for dynamic responses and more complex scenarios
- **Web UI**: for easy management of mock data files
- **Request Logging**: for tracking and analyzing API usage
- **Mock Data Validation**: for ensuring data integrity and consistency
- **Mock Data Versioning**: for managing different versions of mock data

- **Real simple backend**: for serving static files and APIs (v2)
---

### Directory Structure
Organize your mock data files in a structured directory:

```bash
api/
  └── v1/
      └── users/
          ├── get.200.json
          ├── get[slug].200.json
          ├── get.400.json
          ├── post.201.json
          └──(id)
              ├── get.200.json
              └── put.204.json
```

# File Naming Convention for MockTail

MockTail uses a structured file naming convention to dynamically serve mock data based on HTTP methods, URL slugs, query parameters, and status codes. The file name pattern follows this structure:

`<method>.[paramKey: value,...].<status-code>.json`

markdown
Copy code

### Breakdown:

- **Method**: Specifies the HTTP method for the request, such as `get`, `post`, `put`, `delete`.
- **ParamKey: Value**: Optional query parameters to match in the request URL.

- **Status Code**: The HTTP status code to return, placed at the end of the file name.
  - Example: `.200.json` for a successful response, `.404.json` for an error response.

- **File Extension**: `.json` denotes the response format.

### Examples:
- `get.404.json` → Matches a `GET` request and returns a `404 Not Found` response.
- `post.201.json` → Matches a `POST` request and returns a `201 Created` response.

### How It Works:

MockTail dynamically serves responses based on the file names in the mock data directory. The file structure is designed to match various request patterns, including different HTTP methods, slugs, query parameters, and status codes. By following this convention, you can easily simulate a wide variety of API behaviors with minimal configuration, just by organizing your mock data files.

This flexible system allows you to quickly set up mock endpoints and easily modify responses for different HTTP methods and request patterns. MockTail makes it simple to simulate APIs for testing and development purposes, saving time and improving your workflow.


## Installation

### Prerequisites
- Go 1.20+
- Git (for cloning the repository)

## CLI Options

| Flag              | Description                                   | Default  |
|-------------------|-----------------------------------------------|----------|
| `-e, --entry-point` | Path to the entry point file for API mocks   | `./api`  |
| `-m, --mock-dir`   | Directory for mock data storage              | `./api`  |
| `-l, --log-level`  | Log level (`debug`, `info`, `warn`, `error`) | `info`   |

### Clone the Repository
```bash
git clone https://github.com/yourusername/mockTail.git
cd mockTail
```

### Clone the Repository
Build and Run
```bash
go build -o mockTail
./mockTail --entry-point=./api --log-level=debug
```
