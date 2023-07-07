```powershell
$env:PROJECT_ID="cloud-run-poc-392105"
$env:LOCATION="asia-northeast3"
$env:REPO_NAME="kktae-tmp"
$env:FORWARDING_SERVICE_URL="https://google.com"
```


### if repo not exist, run next gcloud cli command
```powershell
gcloud artifacts repositories create $env:REPO_NAME `
    --repository-format=docker `
    --location=$env:LOCATION `
    --description="poc repo" `
    --async
```

```powershell

docker build -t asia-northeast3-docker.pkg.dev/$env:PROJECT_ID/$env:REPO_NAME/go-forward-service .

docker push asia-northeast3-docker.pkg.dev/$env:PROJECT_ID/$env:REPO_NAME/go-forward-service

```