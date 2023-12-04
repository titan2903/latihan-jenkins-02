pipeline {
    agent any

    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.21.4-alpine3.18'
                    label 'sandbox'
                }
            }
            steps {
                echo "Build Apps"
                
            }
        }

        stage('Test'){
            steps {
                echo "Test Apps"
                    
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploy Apps"
            }
        }
    }
}