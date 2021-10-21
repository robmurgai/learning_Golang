pipeline {
    agent any
    // agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
