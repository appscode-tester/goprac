parallel 'docker-test' :{
  node ("master") {
    stage ("Docker Build") {
      docker.image('appscode/golang-agent:1.5').inside {
        stage ("Checkout") {
          checkout scm
        }
        stage ("Build"){
          sh 'go version'
          sh 'go fmt ./cmd/...'
        }
        stage ("Test and validate"){
          sh "go test ./pkg/..."
        }
      }
    }
  }
}, "direct-test" :{
  node {
    stage ("Without Docker Build") {
      checkout scm
      stage("Prepare and Test") {
        sh '''
          cd ../
          sudo chmod -R 777 .
            if [ -d "bin" ]; then
                rm -rf bin
            fi

            sudo apt update
            sudo apt install -y python-dev libyaml-dev python-pip build-essential libsqlite3-dev
            sudo pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage
            sudo apt install wget

            if [ -d "go" ]; then
              echo "go already installed"
            else
              wget -c https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz
              tar -xzf go1.8.1.linux-amd64.tar.gz
              rm -rf go1.8.1.linux-amd64.tar.gz
            fi

            export GOROOT=$PWD/go
            export GOPATH=$PWD
            export GOBIN=$GOPATH/bin
            cd $JOB_NAME
            $GOROOT/bin/go version
            $GOROOT/bin/go test ./cmd/... ./pkg/...
            '''
      }
    }
  }
}

