pipeline {
    agent any
     tools {
        go 'go1.18'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}