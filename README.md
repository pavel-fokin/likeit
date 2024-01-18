# LikeIt!

LikeIt service allows to like a web page ¯\_(ツ)_/¯.

What can be easier than page with one button? The aim of this project to try best practices that I know/learn and consider reasonable.

## List of best practices

- Project Decision Records, aka Architectural Decision Records (ADRs).
- Semantic structuring.
- Embed React app into a Go binary.

## Deployment

### AWS
    AWS Copilot is used to create resources and deploy service.

### GCP
    - `Cloud Build` - CI/CD.
    - `Artifacts Registry` as a repository with Docker images.
    - `Cloud Run` serverless containers to run service.
