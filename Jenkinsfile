pipeline {
    agent any

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/Suhelbashakalaigiri/go-task-api.git'
            }
        }

        stage('Test') {
            steps {
                bat 'go test -v'
            }
        }

        stage('Vet') {
            steps {
                bat 'go vet ./...'
            }
        }

        stage('Build') {
            steps {
                bat 'go build -o go-task-api.exe'
            }
        }

        stage('Docker Build') {
            steps {
                bat 'docker build -t suhelbasha7324/go-task-api:%BUILD_NUMBER% .'
            }
        }

        stage('Docker Push') {
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: 'dockerhub-creds',
                        usernameVariable: 'DOCKER_USERNAME',
                        passwordVariable: 'DOCKER_PASSWORD'
                    )
                ]) {
                    bat '''
                        powershell -NoProfile -Command "$env:DOCKER_PASSWORD | docker login -u $env:DOCKER_USERNAME --password-stdin"
                        if errorlevel 1 exit /b 1

                        docker push %DOCKER_USERNAME%/go-task-api:%BUILD_NUMBER%
                        if errorlevel 1 exit /b 1

                        docker logout
                    '''
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([
                    file(
                        credentialsId: 'kubeconfig-docker-desktop',
                        variable: 'KUBECONFIG'
                    )
                ]) {
                    bat '''
                        kubectl config current-context
                        kubectl get nodes

                        kubectl set image deployment/go-task-api go-task-api=suhelbasha7324/go-task-api:%BUILD_NUMBER%

                        kubectl rollout status deployment/go-task-api
                    '''
                }
            }
        }
    }
}