pipeline {
    agent any

    environment {
        GO111MODULE = 'on'
        DEPLOY_SERVER = 'http://10.12.17.45:3897'
        DEPLOY_USER = 'stella'
        DEPLOY_PATH = '/path/to/your/app'
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/newprintgit/main-admin-api.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myapp'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                sshagent(['your-jenkins-ssh-credentials-id']) {
                    sh '''
                        ssh ${DEPLOY_USER}@${DEPLOY_SERVER} "mkdir -p ${DEPLOY_PATH}"
                        scp myapp ${DEPLOY_USER}@${DEPLOY_SERVER}:${DEPLOY_PATH}/
                        scp myapp.service ${DEPLOY_USER}@${DEPLOY_SERVER}:/etc/systemd/system/
                        ssh ${DEPLOY_USER}@${DEPLOY_SERVER} "sudo systemctl daemon-reload && sudo systemctl restart myapp"
                    '''
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}