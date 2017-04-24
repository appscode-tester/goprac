podTemplate( name: 'gotest', namespace: 'jenkins-automation', inheritFrom: 'jenkins', label: 'gotest', containers: [
    containerTemplate(name: 'golang', image: 'golang:1.6.3', ttyEnabled: true, command: 'cat')
  ]) {
    node ('gotest') {
        container('golang') {
            stage ('Build a Go project'){
              checkout scm
              pwd
              sh 'ls -la'
              sh 'go version'
              sh 'go run cmd/main.go'
            }
        }
    }
}
