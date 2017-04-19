podTemplate(label: 'gopod', containers: [
    containerTemplate(name: 'golang', image: 'golang:1.6.3', ttyEnabled: true, command: 'cat')
  ]) {
      node ('gopod') {
        stage ("checkout and build") {
          checkout scm
          container('golang') {
            stage("build") {
              sh 'go version'
              sh 'pwd'
              sh 'ls -la'
            }
          }
        }
      }
    }
}
