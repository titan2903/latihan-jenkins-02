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
            steps {
                echo "Build Apps"
                sh 'docker build -t gcr.io/ancient-alloy-406700/goapps:${BUILD_NUMBER}.0 .'
            }
        }

        stage('Push to GCR') {
            environment {
                GCR_SERVICE_ACCOUNT = credentials('gcp-service-account-gcr')
            }

            steps {
                echo "push to google cloud registry"
                sh 'cat $GCR_SERVICE_ACCOUNT | docker login -u _json_key --password-stdin https://gcr.io'
                sh 'docker push gcr.io/ancient-alloy-406700/goapps:${BUILD_NUMBER}.0'
            }

            post {
                success {
                    echo "Post Success"
                    discordSend description: "Jenkins Pipeline Push", footer: "Push Success image goapps:${BUILD_NUMBER}.0", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK"
                }
                failure {
                    echo "Post Failure"
                    discordSend description: "Jenkins Pipeline Push", footer: "Push Failure image goapps:${BUILD_NUMBER}.0", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK"
                }
            }
        }

        stage('Deploy') {
            environment {
                KUBE_CONFIG = credentials('kube-config')
            }

            steps {
                echo "Deploy apps with kubernetes"
                sh 'helm repo add goapps-charts https://adhithia21.github.io/helm-charts/charts'
                sh 'helm upgrade --install app goapps-charts/application'
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