pipeline {
    agent {
        node {
            label 'agent'
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
                        extensions: scm.extensions,
                        userRemoteConfigs: scm.userRemoteConfigs
                    )
                }
            }
        }
    }
}
