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
                sh 'docker build -t gcr.io/ancient-alloy-406700/goapps:${BUILD_NUMBER} .'
            }
        }

        stage('Push to GCR') {
            environment {
                GCR_SERVICE_ACCOUNT = credentials('gcp-service-account-gcr')
            }

            steps {
                echo "push to google cloud registry"
                sh 'cat $GCR_SERVICE_ACCOUNT | docker login -u _json_key --password-stdin https://gcr.io'
                sh 'docker push gcr.io/ancient-alloy-406700/goapps:${BUILD_NUMBER}'
            }

            post {
                success {
                    echo "Post Success"
                    discordSend description: "Jenkins Pipeline Push", footer: "Push Success image goapps:${BUILD_NUMBER}", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK"
                }

                failure {
                    echo "Post Failure"
                    discordSend description: "Jenkins Pipeline Push", footer: "Push Failure image goapps:${BUILD_NUMBER}", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "$WEBHOOK"
                }
            }
        }

        stage ('Connect to GCP') {
            environment {
                JENKINS_AGENT_1 = credentials('jenkins-agent-1')
                GCR_SERVICE_ACCOUNT = credentials('gcp-service-account-gcr')
            }

            steps {
                echo 'Connect to access GCP'
                sh 'ssh -o StrictHostKeyChecking=no -i $JENKINS_AGENT_1 titan@192.168.1.131 "rm -rf ~/gcp-service-account.json"'
                sh 'scp -o StrictHostKeyChecking=no -i $JENKINS_AGENT_1 $GCR_SERVICE_ACCOUNT titan@192.168.1.131:~/gcp-service-account.json'
                sh 'ssh -o StrictHostKeyChecking=no -i $JENKINS_AGENT_1 titan@192.168.1.131 "gcloud auth activate-service-account $(cat gcp-service-account.json | jq -r .client_email) --key-file=gcp-service-account.json"'
                sh 'ssh -o StrictHostKeyChecking=no -i $JENKINS_AGENT_1 titan@192.168.1.131 "gcloud auth list"'
                sh 'ssh -o StrictHostKeyChecking=no -i $JENKINS_AGENT_1 titan@192.168.1.131 "gcloud container clusters get-credentials cluster-jenkins-2 --zone asia-southeast2-a --project ancient-alloy-406700"'
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploy apps with kubernetes"
                sh 'helm repo add goapps-charts https://adhithia21.github.io/helm-charts/charts'
                sh 'helm upgrade --install goapps goapps-charts/application'
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