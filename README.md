# SWE-Sumerians
A repository for creating a simple DMS software

## The general work flow idea
```mermaid
flowchart TD;
    A[User Boots Program] -->|Inserts username| B(Web App Interface Opens);
    B --> C{User Selects Action};
    C --> D(Upload File);
    C --> E(Retrieve File);
    C --> F(Edit File);
    D -->|File is stored in DB| G[Returns to Web App Interface];
    E -->|File is fetched from DB| G;
    F -->|File Interface is opened| G;
    G -->|Web Interface Updated| C;
```

## File Structure for the Project

go.main is the root of the project which coordinates the UI (a react app) with the assorted types of RESTful API calls needed (all of which are put into a folder rather than having individual folders for handlers, models, etc. 

```mermaid
flowchart TD
    A[Sumerians SWE 'root'] -->B(go.main)
    A-->C(api)
    A-->D(UI/Frontend)
    C-->E(models)
    C-->F(repositories)
    C-->G(utils)
    C-->H(handlers)
```

## Tools That We are Using

### Backend
Make sure if you are working on the backend you have these installed. 

- Go version 1.24
- Fiber: https://github.com/gofiber/fiber

