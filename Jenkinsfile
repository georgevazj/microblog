pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('Fix workspace permissions') {
            steps {
                sh 'chmod -R 777 $WORKSPACE'
            }
        }
        stage('build') {
            steps {
                sh 'go build $WORKSPACE/cmd/microblog/main.go'
            }
        }
    }
}