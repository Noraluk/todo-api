pipeline {
    agent any

    tools {
        go 'go1.14'
    }

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                go test ./...
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}