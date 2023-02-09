pipeline {
    agent {
        node {
            label 'backend'
        }
    }

    triggers {
        pollSCM 'H/2 * * * *'
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
