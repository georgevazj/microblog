node {
    def root = tool type: 'go', name: 'Go 1.15'
    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
    }
}