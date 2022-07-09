pipeline {
    agent any
     tools {
        go 'go-1.18.3'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                // sh 'go test ./...'
            }
        }
        stage("Build") {
            steps {
                echo 'Building...'
                sh 'docker build -t todo-api .'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                withCredentials([usernamePassword(credentialsId: 'docker', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                echo '${env.dockerhubUser}'
                // sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                // sh 'docker push shadowshotx/product-go-micro'
                }
            }
        }
    }
}