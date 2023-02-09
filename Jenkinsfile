pipeline {
    agent {
        node {
            'agent'
        }
    }

    options {
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scmGit(
                    branches: [[name: '*/develop']],
                    extensions: [],
                    userRemoteConfigs: [[url: 'https://github.com/bismarkanes/go-console-app']]
                )
            }
        }
    }
}
