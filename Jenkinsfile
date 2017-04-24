podTemplate( inheritFrom: 'jenkins', label: 'mypod', containers: [
    containerTemplate(name: 'golang', image: 'golang:1.6.3', ttyEnabled: true, command: 'cat')
  ]) {

    node ('mypod') {
        container('golang') {
            stage 'Build a Go project'
            pwd
            sh 'ls -la'
            sh 'go version'
        }

    }
}
