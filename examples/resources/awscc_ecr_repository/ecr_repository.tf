resource "awscc_ecr_repository" "this" {
  repository_name      = "example-ecr-repository"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration = {
    scan_on_push = true
  }

}
