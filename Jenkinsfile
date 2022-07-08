pipeline {
    agent any
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