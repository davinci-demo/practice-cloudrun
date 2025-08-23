locals {
  # Your GCP project id
  project_id = "<your project ID (not project number)"

  # The GitHub org (or username) where the actions will run
  github_org = "davinci-demo"

  # The individual repos that will be running actions that need to authenticate
  github_repos = [
    "practice-cloudrun"
  ]

  # IAM roles to assign to the Service Account
  roles = [
    "roles/artifactregistry.createOnPushWriter",
  ]
}

# Create the Service Account
resource "google_service_account" "gh_actions_sa" {
  account_id   = "gh-actions"
  display_name = "Service Account for GitHub Actions to push container images and Helm charts to Artifact Registry"
}

# Create a workload identity pool to associate the service account with
resource "google_iam_workload_identity_pool" "ghactions" {
  project                   = local.project_id
  workload_identity_pool_id = "ghactions-pool"
  display_name              = "ghactions-pool"
  description               = "For GitHub Actions authentication"
}

# Create a workload identity provider for GitHub
resource "google_iam_workload_identity_pool_provider" "ghoidc" {
  project                            = local.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.ghactions.workload_identity_pool_id
  workload_identity_pool_provider_id = "ghoidc-provider"
  display_name                       = "ghoidc-provider"
  description                        = "OIDC identity pool provider for execute GitHub Actions"
  # See. https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#understanding-the-oidc-token
  attribute_condition = "assertion.repository_owner == '${local.github_org}'"

  attribute_mapping = {
    "google.subject"             = "assertion.sub"
    "attribute.actor"            = "assertion.actor"
    "attribute.repository"       = "assertion.repository"
    "attribute.repoowner" = "assertion.repository_owner"
    "attribute.repoid"       = "assertion.repository_id"
  }

  oidc {
    issuer_uri        = "https://token.actions.githubusercontent.com"
    allowed_audiences = []
  }
}

# Allow the service account to authenticate with a GitHub action from each of the declared repos
resource "google_service_account_iam_member" "workload_identity_user" {
  for_each = toset(local.github_repos)

  service_account_id = google_service_account.gh_actions_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.ghactions.name}/attribute.repository/${local.github_org}/${each.value}"
}

# Attach the roles that the Service Account will need
resource "google_project_iam_member" "sa_roles" {
  for_each = toset(local.roles)

  project = local.project_id
  role    = each.value
  member  = "serviceAccount:${google_service_account.gh_actions_sa.email}"
}

# This is the value you'll need for the GitHub Action authenticate step
output "service_account_email" {
  description = "The Service Account email"
  value       = google_service_account.gh_actions_sa.email
}

# This is the value you'll need for the GitHub Action authenticate step
output "workload_identity_pool_name" {
  description = "Workload Identity Pood Provider ID"
  value       = google_iam_workload_identity_pool_provider.ghoidc.name
}
