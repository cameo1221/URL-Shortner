<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/6295/6295417.png" width="100" />
</p>
<p align="center">
    <h1 align="center">URL-SHORTNER</h1>
</p>
<p align="center">
    <em><code>► Basic URL shortner for YAML & JSON files along with Postgres Database</code></em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/cameo1221/URL-Shortner?style=flat&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/cameo1221/URL-Shortner?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/cameo1221/URL-Shortner?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/cameo1221/URL-Shortner?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
		<em>Developed with the software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
  <img src="https://img.shields.io/badge/PostgreSQL-8A004A.svg?style=flat&logo=PostgreSQL&logoColor=white)](https://www.postgresql.org/" alt="PostgresSQL">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/docker-0db7ee.svg?style=flat&logo=docker&logoColor=white)](https://www.docker.com/" alt="Docker">
  <img src="https://img.shields.io/badge/JSON-EDD000.svg?style=flat&logo=JSON&logoColor=white)](https://www.json.org/)" alt="JSON">
  
</p>
<hr>

## 🔗 Quick Links

> - [📍 Overview](#-overview)
> - [📂 Repository Structure](#-repository-structure)
> - [🧩 Modules](#-modules)
> - [🚀 Getting Started](#-getting-started)
>   - [⚙️ Installation](#️-installation)
>   - [🤖 Running URL-Shortner](#-running-URL-Shortner)
> - [🤝 Contributing](#-contributing)
> - [📄 License](#-license)
> - [👏 Acknowledgments](#-acknowledgments)

---

## 📍 Overview

<code>► ## URL Shortener Service in Go with JSON/YAML Parsing and PostgreSQL Integration

This project implements a robust URL shortener service written in Go. It boasts the following functionalities:

##  Overview

**URL Shortener Service in Go with JSON/YAML Parsing and PostgreSQL Integration**

This project implements a robust URL shortener service written in Go. It boasts the following functionalities:

**Core Features:**

* **URL Shortening:** Efficiently shortens long URLs, providing a user-friendly and easy-to-share alternative.
* **Redirection:** Seamlessly redirects users to the original URL when they access the shortened version.

**Data Management:**

* **JSON/YAML Parsing:** Facilitates the creation of URL mappings in bulk through JSON and YAML configuration files.
* **PostgreSQL Integration:** Leverages PostgreSQL as a persistent storage solution, enabling reliable data storage and retrieval of shortened URLs.

**Benefits:**

* **Enhanced User Experience:** Simplifies URL sharing and improves readability, particularly for lengthy URLs.
* **Increased Efficiency:** Optimizes link management for applications and simplifies URL sharing within teams.
* **Flexibility:** Supports both JSON and YAML file formats for configuration, catering to different user preferences.
* **Scalability:** PostgreSQL offers a robust foundation for managing a large volume of URLs without compromising performance.

---

## 📦 Features

<code>► INSERT-TEXT-HERE</code>

---

## 📂 Repository Structure

```sh
Main Branch
└── URL-Shortner/
    ├── Handler
    │   └── handler.go
    ├── go.mod
    ├── go.sum
    └── main.go
Feature-YAML
└── URL-Shortner/
    ├── Handler
    │   └── handler.go
    ├── go.mod
    ├── go.sum
    ├── yamlfile.json
    └── main.go
Feature-JSON
└── URL-Shortner/
    ├── Handler
    │   └── handler.go
    ├── go.mod
    ├── go.sum
    ├── jsonfile.json
    └── main.go
Feature-Postgres
└── URL-Shortner/
    ├── Handler
    │   └── handler.go
    ├── go.mod
    ├── go.sum
    └── main.go
```

---

## 🧩 Modules

<details closed><summary>.</summary>

| File                                                                     | Summary                         |
| ---                                                                      | ---                             |
| [go.sum](https://github.com/cameo1221/URL-Shortner/blob/master/go.sum)   | <code>► This file lists all the external Go libraries your project depends on, along with their versions. It's automatically generated by the go mod download command.</code> |
| [go.mod](https://github.com/cameo1221/URL-Shortner/blob/master/go.mod)   | <code>►This file defines the Go module your project belongs to and specifies the dependencies required for that module. It can also be used to replace missing dependencies that weren't automatically downloaded.</code> |
| [main.go](https://github.com/cameo1221/URL-Shortner/blob/master/main.go) | <code>► This is the main entry point of your application. It typically contains the following: Import statements for the necessary Go packages (including the ones listed in go.sum). Function definitions related to service initialization (e.g., connecting to the database, starting the server) The main function that starts the service and listens for incoming requests.</code> |

</details>

<details closed><summary>Handler</summary>

| File                                                                                   | Summary                         |
| ---                                                                                    | ---                             |
| [handler.go](https://github.com/cameo1221/URL-Shortner/blob/master/Handler/handler.go) | <code>► Server Initialization: Creates a new http.ServeMux instance (multiplexer) named mux to handle incoming HTTP requests. Defines a simple "Hello" handler function (hello) that responds with "Hello" for the root path (/). Registers the hello handler with the mux for the root path. URL Mapping Configuration: Creates a static map named pathsToUrls to define pre-defined short URLs for specific paths. (This could be replaced with database interaction or another approach for dynamic mapping.) Utilizes the URLShortner.MapHandler function (presumably from your URLShortner package) to: Register the pre-defined mappings from pathsToUrls within the mux. YAML Configuration Processing: Defines a YAML string named yaml containing path-URL associations in a specific format. Calls the URLShortner.YAMLHandler function, passing: The YAML data as a byte slice ([]byte(yaml)) The previously created mapHandler (containing pre-defined mappings) This function likely parses the YAML data, creates additional URL mappings based on the parsed information, and combines those with the pre-defined mappings from mapHandler. Stores the combined handler in the yamlHandler variable. Server Start: Prints a message indicating the server will start on port 8080. Starts the HTTP server on port 8080, using the yamlHandler to process incoming requests. This means your server will handle both pre-defined URLs and those defined in the YAML configuration.</code> |

</details>

---

## 🚀 Getting Started

***Requirements***

Ensure you have the following dependencies installed on your system:

* **Go**: `version x.y.z`

### ⚙️ Installation

1. Clone the URL-Shortner repository:

```sh
git clone https://github.com/cameo1221/URL-Shortner
```

2. Change to the project directory:

```sh
cd URL-Shortner
```

3. Install the dependencies:

```sh
go build -o myapp
```

### 🤖 Running URL-Shortner

Use the following command to run URL-Shortner:

```sh
./myapp
```


## 🤝 Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Submit Pull Requests](https://github.com/cameo1221/URL-Shortner/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/cameo1221/URL-Shortner/discussions)**: Share your insights, provide feedback, or ask questions.
- **[Report Issues](https://github.com/cameo1221/URL-Shortner/issues)**: Submit bugs found or log feature requests for Url-shortner.

<details closed>
    <summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your GitHub account.
2. **Clone Locally**: Clone the forked repository to your local machine using a Git client.
   ```sh
   git clone https://github.com/cameo1221/URL-Shortner
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to GitHub**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.

Once your PR is reviewed and approved, it will be merged into the main branch.

</details>

---

[**Return**](#-quick-links)

---
