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
