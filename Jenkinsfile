pipeline {
    agent {
        node {
            label 'agent'
        }
    }

    triggers {
        pollSCM 'H/5 * * * *'
    }

    options {
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    checkout scmGit(
                        branches: [[name: '*/develop']],
                        extensions: [],
                        userRemoteConfigs: [[url: 'https://github.com/bismarkanes/go-console-app']]
                    )
                }
            }
        }
    }
}
