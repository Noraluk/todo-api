pipeline {
    agent any
    stages {
        stage('Chekout') {
            steps {
                git url: 'https://github.com/Noraluk/todo-api.git', branch: 'develop'
                echo 'Checkout Completed'
            }
        }
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