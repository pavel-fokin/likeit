# LikeIt!

LikeIt project allows to like a web page ¯\\\_(ツ)_/¯.

What can be easier than page with one button? The aim of this project to try different approaches that I know/learn and consider reasonable.

## List of approaches

- Decisions Log, aka Architectural Decision Records (ADRs). [Decisions Log](docs/decisions/0001-decisions-log.md)
- Use of a monorepo. [Monorepo](docs/decisions/0002-monorepo.md)
- Embed React app into a Go binary.

## Deployment

### AWS
    AWS Copilot is used to create resources and deploy service.

### GCP
    - `Cloud Build` - CI/CD.
    - `Artifacts Registry` as a repository with Docker images.
    - `Cloud Run` serverless containers to run service.
