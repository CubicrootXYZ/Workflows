# Workflows

Collection of commonly used GitHub workflows.

## Building images

This workflow will build and push images to Dockerhub. For any change an image with the commit SHA as tag will be published. For releases the release tag will be used to tag the image.

Provide a `static_tag` so any commit on `main` branch will result in an image pushed with that tag. Can be used to provide a `beta` or `main` tagged image.

```yaml
jobs:
  build_image:
    uses: CubicrootXYZ/Workflows/.github/workflows/build_image.yaml@v1.0.0
    with:
      docker_build_args: "--no-cache"
      docker_file_path: "./"
      image_name: "example/image"
      vuln_scan: false # Uses grype to scan for vulnerabilities. 
    secrets:
      dockerhub_user: ${{ secrets.DOCKERHUB_USERNAME }}
      dockerhub_token: ${{ secrets.DOCKERHUB_TOKEN }}
```

## Golang

### Code Quality

To ensure code quality use the provided `golang_quality` workflow. It runs multiple analysis tools.

```yaml
permissions:
  # Required by golangci job to write annotations to the merge request.
  contents: read
  checks: write
  
jobs:
  golang_quality:
    uses: CubicrootXYZ/Workflows/.github/workflows/golang_quality.yaml@v1.0.0
    with:
      workdir: "golang/application/subfolder/"
```

### Execute Tests

To execute golang tests use the provided `golang_test` workflow.

```yaml
jobs:
  golang_test:
    uses: CubicrootXYZ/Workflows/.github/workflows/golang_test.yaml@v1.0.0
    with:
      workdir: "tests/golang_test/"
      services: '{"database": {"image": "mysql:8.0", "ports": ["3306:3306"]}}'
      env: '{"TESTENV": "value"}'
```

## Jinja

### Validate generates are up to date

Add the following job to your CI to ensure `jinja` generated files are always up to date. Uses `makejinja` for templating.

```yaml
jobs:
  jinja:
    uses: CubicrootXYZ/Workflows/.github/workflows/jinja.yaml@v1.0.0
    with:
      workdir: "tests/jinja/"
```

## OpenAPI

### Render OpenAPI 2 specs with redoc

Render an OpenAPI 2 spec with `redoc`. The spec file needs to be present in the given artifact. The rendered `index.html` will be added to the artifact `rendered-api-docu`.

```yaml
jobs:
  render:
    uses: CubicrootXYZ/Workflows/.github/workflows/openapi2_render.yaml@v1.0.0
    with:
      spec_artifact_name: spec
      spec_artifact_path: tests/openapi2/
      spec_filename: petstore.yaml
```

### Generate OpenAPI 2 spec from golang source code

Generate OpenAPI 2 spec file based on `swag` compatible annotiations in the source code.

```yaml
jobs:
  build:
    uses: CubicrootXYZ/Workflows/.github/workflows/openapi2_golang_build.yaml@v1.0.0
    with:
      entrypoint: cmd/main.go
      workdir: ./
```

## GitHub Pages

Make sure GitHub pages are activated and set to "GitHub Action" in your repository settings.

Add the following permissions to the `workflows` file:

```yaml
permissions:
  contents: read
  pages: write
  id-token: write
```

To deploy an artifact to GitHub pages use the following job:

```yaml
jobs:
  pages_test:
    uses: CubicrootXYZ/Workflows/.github/workflows/pages.yaml@v1.0.0
    with:
      artifact_name: static-html
      artifact_path: index.html
```

## Renovate

Create a new access token and set it as `RENOVATE_TOKEN` repo secret. 

```yaml
jobs:
  renovate_test:
    uses: CubicrootXYZ/Workflows/.github/workflows/workflows/renovate.yaml@v1.0.0
    with:
      author: "Max <max@users.noreply.github.com>"
    secrets:
      token: "${{ secrets.RENOVATE_TOKEN }}"
```