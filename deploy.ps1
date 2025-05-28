$projectRoot = Split-Path -Parent $MyInvocation.MyCommand.Definition
Write-Host "Projeto raiz: $projectRoot"

$dateTime = Get-Date -Format "yyyyMMdd-HHmm"
$imageBase = "authentication"
$imageTag = "${imageBase}:${dateTime}"

Write-Host "=== Buildando imagem Docker $imageTag ==="
docker build -t $imageTag $projectRoot
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao buildar imagem Docker"
    exit 1
}

$deploymentFile = "$projectRoot\k8s\deployment.yaml"
Write-Host "=== Atualizando tag da imagem no arquivo $deploymentFile ==="
(Get-Content $deploymentFile) -replace 'image: authentication:.*', "image: $imageTag" | Set-Content $deploymentFile

Write-Host "=== Criando namespace authentication ==="
kubectl apply -f "$projectRoot\k8s\namespace.yaml"
if ($LASTEXITCODE -ne 0) {
    Write-Warning "Falha ao criar namespace authentication, verifique se já existe e se está acessível"
}

if (-not $env:MY_POSTGRES_PASSWORD) {
    Write-Error "Variável de ambiente MY_POSTGRES_PASSWORD não está definida. Abortando."
    exit 1
}

kubectl delete secret postgres-secret -n authentication --ignore-not-found
kubectl create secret generic postgres-secret -n authentication --from-literal=password=$env:MY_POSTGRES_PASSWORD
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao criar a Secret postgres-secret"
    exit 1
}

Write-Host "=== Limpando imagens Docker antigas (mantendo 2 últimas) ==="
$images = docker images --format "{{.Repository}}:{{.Tag}}" | Where-Object { $_ -like "${imageBase}:*" } | Sort-Object
if ($images.Count -gt 2) {
    $imagesToDelete = $images | Select-Object -First ($images.Count - 2)
    foreach ($img in $imagesToDelete) {
        Write-Host "Removendo imagem antiga: $img"
        docker rmi $img
    }
} else {
    Write-Host "Nenhuma imagem antiga para remover."
}

kubectl apply -f "$projectRoot\k8s\deployment.yaml"
kubectl apply -f "$projectRoot\k8s\service.yaml"

kubectl get pods -n authentication
kubectl get svc -n authentication
