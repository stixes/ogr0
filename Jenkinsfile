pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'docker build --pull -t stixes/ogr0:latest .'
      }
    }
    stage('Test') {
      steps {
        sh 'docker-compose -f docker-compose.test.yml up --abort-on-container-exit --exit-code-from sut -V'
        sh 'docker-compose -f docker-compose.test.yml down -v'
      }
    }
    stage('Push') {
      environment {
        HUB = credentials('hub_login')
      }
      steps {
        sh 'echo $HUB_PSW|docker login -u $HUB_USR --password-stdin'
        sh 'docker push stixes/ogr0:latest'
        sh 'rm -f ~/.docker/config.json'
      }
    }
  }
}
