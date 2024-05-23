# November-2023-hiring-Sneh16Shah
november-2023-hiring-Sneh16Shah created by GitHub Classroom

<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#system-design">System Design</a></li>
    <li><a href="#implemented-features">Implemented Features</a></li>
    <li><a href="#structure">Structure</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

The Log Ingestor is a centralized logging system designed to efficiently ingest logs and offer a user-friendly web interface for querying log data. Key functionalities include:

- **Log Ingestion**:
  - Logs are ingested through a script accessible at `[http://localhost:3000/]`.
  - Logs are structured with details such as level, message, resourceId, timestamp, traceId, spanId, commit, and metadata.parentResourceId.

- **User Interface**:
  - Web UI offers search capabilities using various query types:
    - Full-text search for text-based queries.
    - Filter-based queries for:
      - Level
      - Message
      - ResourceId
      - Timestamp
      - TraceId
      - SpanId
      - Commit
      - Metadata.ParentResourceId
    - Date range filtering for time-specific log retrieval.

The system aims to provide a streamlined experience for users to efficiently search and retrieve logs based on specific criteria through an intuitive web interface.



### Built With
Language/Frameworks used in this project

* [![React][React.js]][React-url]
* [![Go][Go-icon]][Go-url]
* [![MongoDB][Mongodb-icon]][Mongodb-url]
<p align="right">(<a href="#readme-top">back to top</a>)</p>

## System Design

The project is structured as a client-server architecture. Here's a breakdown:

- Client-Side (Frontend)

  - **Folder Structure**:
    - **`client/`**: Contains the React.js frontend application.
      - **`src/`**: Holds React components and application logic.
      - **`public/`**: Contains the HTML entry point (`index.html`).

- Server-Side (Backend)

  - **Folder Structure**:
    - **`server/`**: Contains the Golang backend application.
      - **`databaseHandler/`**: Handles MongoDB interactions (`mongodb.go`).
      - **`queryHandler/`**: Manages different types of log queries.
      - **`routeHandler/`**: Defines API endpoints for log ingestion and querying.
      - **`utils/`**: Includes utility functions (`utils.go`).
<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Implemented Features
- Develop a mechanism to ingest logs in the provided format via an HTTP server, which runs on port `3000` by default.(I have used injest.js script to insert logs)
- Ensure Scalability if new key-value comes in logs.
- A User interface UI for full-text search across logs.
- Include filters based on:
    - level
    - message
    - resourceId
    - timestamp
    - traceId
    - spanId
    - commit
    - metadata.parentResourceId
- Advanced Features
    - Implement search within specific date ranges.
    - Utilize regular expressions for search.
    - Allow combining multiple filters.
- The code passing all the given validation requirements and it does error-handling quiet well
<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Structure
```
Repository
├───README.md
│   
├───client
│   │   .env
│   │   .gitignore
│   │   ingest.js
│   │   package-lock.json
│   │   package.json
│   │   README.md
│   │   
│   ├───public
│   │       index.html
│   │       
│   └───src
│       │   App.js
│       │   index.js
│       │   
│       └───components
│               DateList.js
│               Filters.js
│               LogList.js
│               SearchBar.js
│               
└───server
    │   go.mod
    │   go.sum
    │   main.go
    │   
    ├───databaseHandler
    │       mongodb.go
    │       
    ├───queryHandler
    │       typeDateRange.go
    │       typeFilters.go
    │       typeFullText.go
    │       
    ├───routeHandler
    │       ingestLog.go
    │       queryLog.go
    │       
    └───utils
            utils.go
```         


<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites 
 - For Go (Golang):

1. **Download and Install Go:**
   - Download the Go installation package suitable for your operating system from the [official Go website](https://golang.org/dl/).
   - Follow the installation instructions for your OS.

2. **Set Up Go Environment:**
   - Set the `GOPATH` environment variable to point to your Go workspace.
   - Ensure that the Go `bin` directory is added to your system's `PATH`.

  - For React.js:

1. **Node.js and npm:**
   - Install Node.js, which includes npm, from the [official Node.js website](https://nodejs.org/).
   - Follow the installation instructions for your OS.


### Installation

1. Clone the repo
   ```
   git clone https://github.com/dyte-submissions/november-2023-hiring-Sneh16Shah.git
   ```
2. For the Frontend part
  - Navigate to the Client Directory:
   ```
   cd client
   ```
  - Install Dependencies:
    ```
    npm install
    ```
  - In case of any error:
    ```
    npm audit fix
    ```
  - Run the Frontend Application:
    ```
    npm start
    ```
3. For the Backend part
   - Navigate to the Server Directory:
   ```
   cd client
   ```
   - Install Dependencies
     ```
     go mod tidy
     ```
   - Run the Backend Application
     ```
     go run main.go
     ```
4. For Ingesting log into Mongodb
   - Navigate to the Client Directory:
   ```
   cd client
   ```
  - Run the ingest.js File:
    ```
    node ingest.js
    ```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage
When you completely run Frontend part this web design will be shown
![image](https://github.com/dyte-submissions/november-2023-hiring-Sneh16Shah/assets/75805672/1b7083fa-34d2-4d52-a52e-2d703b8f1674)

- Now User can type text to text box and hit button to see their desired logs. For example text = "error"
![image](https://github.com/dyte-submissions/november-2023-hiring-Sneh16Shah/assets/75805672/b83656f6-9ffc-4e6f-8d05-366c1db6a060)

- User can also apply single filter and multiple filter to get their desired logs.
  - single filter (traceId = "abc-xyz-123")
![image](https://github.com/dyte-submissions/november-2023-hiring-Sneh16Shah/assets/75805672/b56d42c8-090f-4f12-b356-c95c23596d48)

  - multiple filter (traceId = "abc-xyz-123" and timestamp = "2023-09-16"
![image](https://github.com/dyte-submissions/november-2023-hiring-Sneh16Shah/assets/75805672/f859d0de-6a38-4dfd-93f1-2d07aefdc9fc)

- User can also see the logs in date ranges. For user's simplicity they only needs to provide the date in YYYY-MM-DD instead of iso-format.
    startDate = "2023-09-15", endDate = "2023-09-16"
![image](https://github.com/dyte-submissions/november-2023-hiring-Sneh16Shah/assets/75805672/be2fbe54-46ff-43b5-aa85-648d529db0ff)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Sneh Shah
* [![Github][Github-icon]][Github-url]
* [![Linkedln][Linkedln-icon]][Linkedln-url]
* [![Twitter][Twitter-icon]][Twitter-url]


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [GitHub Pages](https://pages.github.com)
* [Icons](https://github.com/alexandresanlim/Badges4-README.md-Profile)

<p align="right">(<a href="#readme-top">back to top</a>)</p>




[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Go-icon]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://golang.org/
[Mongodb-icon]: https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white
[Mongodb-url]: https://www.mongodb.com/
[Github-icon]: https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white
[Github-url]: https://github.com/Sneh16Shah
[Linkedln-icon]: https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white
[Linkedln-url]: https://www.linkedin.com/in/snehshah16/
[Twitter-icon]: https://img.shields.io/badge/X-000000?style=for-the-badge&logo=x&logoColor=white
[Twitter-url]: https://twitter.com/SnehSha04698488
