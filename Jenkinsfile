pipeline {
    agent any
    environment {
        gitcommit = "${gitcommit}"
    }
    tools {
        go 'go'
    }

    stages {
        stage ('SCM validation') {
            steps {
                checkout scm
                sh "git rev-parse --short HEAD > .git/commit-id"
                gitcommit = readFile('.git/commit-id').trim()
            }
        }
        stage ('Build') {
            steps {
                sh 'cd cmd'
                sh 'go build .'
            }
        }
    }
}