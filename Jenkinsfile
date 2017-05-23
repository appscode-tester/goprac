node {
  def PWD = pwd()
  def project_dir = "${PWD}/src/github.com/ashiquzzaman33/goprac"
  def go_version = "1.8.1"
  stage("set env") {
      env.GOROOT= "${PWD}/go"
      env.GOPATH= "${PWD}"
      env.GOBIN = "${GOPATH}/bin"
      env.PATH = "${env.GOROOT}/bin:${env.PATH}:$GOPATH:$GOBIN"
  }
  stage('builddeps') {
    sh 'sudo apt update &&\
        sudo apt install -y python-dev libyaml-dev python-pip build-essential libsqlite3-dev &&\
        sudo pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage &&\
        sudo apt install curl'
  }
  Another update :D
  stage("go setup") {
      try {
        sh "go version"
      } catch (e) {
          sh "curl -OL https://storage.googleapis.com/golang/go${go_version}.linux-amd64.tar.gz &&\
          tar -xzf go${go_version}.linux-amd64.tar.gz &&\
          rm -rf go${go_version}.linux-amd64.tar.gz"
      }

  }
  stage("checkout") {
      dir("${project_dir}") {
         checkout scm
      }
  }
  stage("build") {
      dir("${project_dir}") {
        sh "go install ./..."
      }
  }
  stage("build") {
      dir("${project_dir}") {
        sh "goprac"
      }
  }
}
