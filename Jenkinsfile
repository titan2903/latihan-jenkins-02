pipeline {
    agent any

    stages {
        stage('Test Goapps'){
            agent {
                docker {
                    image 'golang:1.21.4-alpine3.18'
                    label 'sandbox'
                }
            }
            steps {
                echo "Test Golang Apps"
                sh 'GOCACHE=/tmp/ go test -v ./...'
            }   
        }
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.21.4-alpine3.18'
                    label 'sandbox'
                }
            }
            steps {
                echo "Build Apps"
                sh 'docker build -t goapps:1.0 .'
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploy Apps"
            }
        }
    }
}