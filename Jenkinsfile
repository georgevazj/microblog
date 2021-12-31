node {
    def root = tool type: 'go', name: 'go'
    def gitcommit

    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
    }

    stage ('SCM validation') {
        checkout scm
        sh "git rev-parse --short HEAD > .git/commit-id"
        gitcommit = readFile('.git/commit-id').trim()
    }
    stage ('Build') {
        sh 'go build cmd/main.go'
    }
}