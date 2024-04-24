# Monorepo

## Context

This project consists of two components (backend and frontend), which are
co-developed. As a the sole developer, I need a simple setup for this project
that helps to oversee the whole solution.

## Decision
I decided to use a monorepo approach to manage the project's source code.
The monorepo will be following the next structure.

```sh
/likeit   # Project name.
    /apps   # All software components of the project.
    /deploy # Configurations for the different deployment environments (AWS, GCP).
    /docs   # Project documentation.
```

## Consequences

### Pros:

    - **Holistic view:** Keeps all the code and information under one repo.
    That makes it simpler to oversee the whole project.

### Cons:

    - **Tooling:** It's harder to setup CI/CD pipelines and configure building of images.
    - **Layout:** It can be harder to manage project layout
    when using different programming languages.