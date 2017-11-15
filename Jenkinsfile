node('golang') {
	def root = tool name: 'Go 1.9', type: 'go'
	def version = '0.1.0'
    def gitUrl = 'git@github.com:stanchan/supervisord.git'
    def httpGitUrl = 'https://github.com/stanchan/supervisord.git'
    def projectPath = 'github.com/stanchan/supervisord'

	git([url: gitUrl, branch: 'master', credentialsId: 'git-key'])

	stage('prep') {
		withEnv(["GOROOT=${root}", "GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin"]) {
			sh 'go version'
			sh "go get ${projectPath}"
		}
	}

	stage('build') {
		withEnv(["GOROOT=${root}", "GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin"]) {
			sh 'cd $WORKSPACE && go build .'
			sh 'mv $WORKSPACE/bin/supervisord $WORKSPACE/'
		}
	}

	stage('publish') {
		nexusArtifactUploader artifacts: [[ artifactId: 'supervisord',
											classifier: '',
											file: 'supervisord',
											type: 'bin' ]],
							  credentialsId: 'nexus-deploy',
							  groupId: 'com.github.uberdeploy.supervisord',
							  nexusUrl: 'nexus',
							  nexusVersion: 'nexus3',
							  protocol: 'https',
							  repository: 'maven-releases',
							  version: version
	}
}
