node {
    def root = tool type: 'go', name: 'go'
    def gitcommit

    stage ('SCM validation') {
        checkout scm
        sh "git rev-parse --short HEAD > .git/commit-id"
        gitcommit = readFile('.git/commit-id').trim()
    }
    stage ('Build') {
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go build cmd/main.go'
        }
    }
}