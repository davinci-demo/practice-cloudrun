
locals {
  # Your GCP project id
  project_id = "<your project ID (not project number)"
  # What to name the service
  cloudrun_name = "practice-davinci"

  cloudrun_location =    "us-central1"
  image_url =    "gcr.io/curious-athlete-469521-c8/practice-cloudrun@sha256:bb7001b77700ce2a6e6025fa682c18f10721abf22bb028beb72936c9010ecadd"

  # TODO move to Secrets Manager
  db_server_url = "postgresql://<your Postgres database>"
}

  provider "google" {
    project = local.project_id
  }

  resource "google_cloud_run_v2_service" "default" {
    name     = local.cloudrun_name
    location = local.cloudrun_location
    deletion_protection   = false

    template {
      containers {
        image = local.image_url

      # Environment variables
      env {
        name  = "DB_SERVER_URL"
        value = local.db_server_url
      }
      env {
        name  = "DB_MAX_CONNECTIONS"
        value = 100
      }
      env {
        name  = "DB_MAX_IDLE_CONNECTIONS"
        value = 2
      }
      env {
        name  = "DB_MAX_LIFETIME_CONNECTIONS"
        value = 300
      }
      env {
        name  = "SERVER_READ_TIMEOUT"
        value = 60
      }
      env {
        name  = "SERVER_URL"
        value = "0.0.0.0:8080"
      }

      }
    }
  }

  resource "google_cloud_run_v2_service_iam_member" "noauth" {
    location = google_cloud_run_v2_service.default.location
    name     = google_cloud_run_v2_service.default.name
    role     = "roles/run.invoker"
    member   = "allUsers"
  }
  
  
