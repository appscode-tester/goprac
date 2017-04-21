podTemplate(label: 'pod3', containers: [
    containerTemplate(name: 'jnlp', image: 'appscode/jnlp-slave:canary', args: '${computer.jnlpmac} ${computer.name}'),
    containerTemplate(name: 'golang', image: 'golang:1.6.3', ttyEnabled: true, command: 'cat')
  ]) {

    node ('pod3') {
        stage 'Get a Golang project'
        git url: 'https://github.com/hashicorp/terraform.git'
        container('golang') {
            stage 'Build a Go project'
            sh """
            mkdir -p /go/src/github.com/hashicorp
            ln -s `pwd` /go/src/github.com/hashicorp/terraform
            cd /go/src/github.com/hashicorp/terraform && make core-dev
            """
        }

    }
}
