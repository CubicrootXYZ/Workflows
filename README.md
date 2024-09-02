# Workflows

Collection of commonly used GitHub workflows.

## Building images

This workflow will build and push images to Dockerhub. For any change an image with the commit SHA as tag will be published. For releases the release tag will be used to tag the image.

```yaml
jobs:
  build_image:
    uses: CubicrootXYZ/Workflows/.github/workflows/build_image.yaml@v1.0.0
    with:
      docker_build_args: "--no-cache"
      docker_file_path: "./"
      image_name: "example/image"
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