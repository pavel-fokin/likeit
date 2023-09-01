# LikeIt!

LikeIt service allows to like a web page ¯\_(ツ)_/¯. This is a reference project to try project structure and some scenarios.

- Semantic structuring.
- Embed React app into a Go binary.

## Deployment

### AWS
    AWS Copilot is used to create resources and deploy service.

### GCP
    - `Cloud Build` - CI/CD.
    - `Artifacts Registry` as a repository with Docker images.
    - `Cloud Run` serverless containers to run service.