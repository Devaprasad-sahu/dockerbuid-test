pipeline {
    agent { docker { image 'golang:1.14' } }
    environment {
        GOCACHE = '/tmp/gocache'
    }
    stages {
        stage('build') {
            steps {
                sh 'go build'
            }
        }
        stage('test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
