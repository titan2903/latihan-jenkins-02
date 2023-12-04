pipeline {
    agent {
        label 'sandbox'
    }

    environment {
        WEBHOOK = credentials('WEBHOOK_URL_DISCORD')
    }

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
                label 'sandbox'
            }

            steps {
                echo "Build Apps"
                sh 'docker build -t goapps:1.0 .'
            }
        }

        stage('Push to GCR') {
            steps {
                echo "push to google cloud registry"
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploy Apps"
            }
        }
    }

    post {
        success {
            echo "Post Success"
            discordSend description: "Jenkins Pipeline Deploy", footer: "Deploy Success", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK"
        }
        failure {
            echo "Post Failure"
            discordSend description: "Jenkins Pipeline Deploy", footer: "Deploy Failure", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK"
        }
    }
}