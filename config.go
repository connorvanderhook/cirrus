package cirrus

// AppConfig is where the secrets are
type AppConfig struct {
	Address string `hcl:"address" envconfig:"ADDRESS"`
	BaseURL string `hcl:"base_url" envconfig:"BASE_URL"`

	FileStorage struct {
		Type string `hcl:"type" envconfig:"FILE_STORAGE_TYPE"`

		Local struct {
			Dir string `hcl:"dir" envconfig:"FILE_STORAGE_LOCAL_DIR"`
		} `hcl:"local"`

		GoogleCloudStorage struct {
			ServiceAccountFile string `hcl:"service_account_file" envconfig:"FILE_STORAGE_GCS_SERVICE_ACCOUNT_FILE"`
			Bucket             string `hcl:"bucket" envconfig:"FILE_STORAGE_GCS_BUCKET"`
		} `hcl:"google_cloud_storage"`

		AmazonS3 struct {
			AccessKey string `hcl:"access_key" envconfig:"FILE_STORAGE_S3_ACCESS_KEY"`
			SecretKey string `hcl:"secret_key" envconfig:"FILE_STORAGE_S3_SECRET_KEY"`
			Region    string `hcl:"region" envconfig:"FILE_STORAGE_S3_REGION"`
			Bucket    string `hcl:"bucket" envconfig:"FILE_STORAGE_S3_BUCKET"`
		} `hcl:"amazon_s3"`
	} `hcl:"file_storage"`
}
