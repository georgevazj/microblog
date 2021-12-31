pipeline {
    agent any
    tools {
        go 'go'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build $WORKSPACE/cmd/microblog/main.go'
            }
        }
    }
}