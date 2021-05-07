    def createNamespace (namespace) {
        echo "Creating namespace ${namespace} if needed"

        sh "[ ! -z \"\$(kubectl get ns ${namespace} -o name 2>/dev/null)\" ] || kubectl create ns ${namespace}"
    }
    
   def helmInstall (namespace, release) {
    

    script {
        release = "${release}-${namespace}"
        echo "Installing ${release} in ${namespace}"
        sh """
            helm upgrade --install --namespace ${namespace} ${release} \
             -f helm/gocalc/values-${namespace}.yaml                   \
                 /helm/gocalc
        """
        sh "sleep 5"
    }
}


pipeline {
    
    parameters {
        string (name: 'GIT_BRANCH',           defaultValue: 'main',  description: 'Git branch to build')
        booleanParam (name: 'DEPLOY_TO_PROD', defaultValue: false,     description: 'deploy to production without manual approval')
    
    environment {
        IMAGE_NAME = 'gocalc'
        TEST_LOCAL_PORT = 8081
        DEPLOY_PROD = false
        ID = "${DOCKER_REG}-${IMAGE_NAME}"
        branch = GIT_BRANCH    
        PARAMETERS_FILE = "${JENKINS_HOME}/parameters.groovy"
    }
    
    agent { node { label 'master' } }

    stages {
        
        stage('Git clone and setup'){
            steps{
                echo "checkout code"
                git branch: 'main', credentialsId: 'GitHubID', url: 'https://github.com/arif11111/Gocalc-Helm-.git'
                }
               
            }
        }
        
        stage('Build'){
            steps{
                echo "Building Image"
                
                withCredentials([usernameColonPassword(credentialsId: 'dockerID', variable: 'docker_credn')]) {
                    sh "docker login"
                }    
                    sh " docker build -t ${DOCKER_REG}/${IMAGE_NAME} .  "
                    
                    echo "Running tests"

                    // Kill container in case there is a leftover
                    sh "[ -z \"\$(docker ps -a | grep ${ID} 2>/dev/null)\" ] || docker rm -f ${ID}"

                    echo "Starting ${IMAGE_NAME} container"
                    sh "docker run --detach --name ${ID} --rm --p ${TEST_LOCAL_PORT}:8080 ${DOCKER_REG}/${IMAGE_NAME}"
                    
                    script {
                       
                    host_ip = localhost:"${TEST_LOCAL_PORT}""
                    }
                
            }
        }
        
        stage('Local Tests') {
            
            steps {
                script{
                    if curl -s --head  --request GET http://${host_ip}/status | grep "200 OK" > /dev/null; then 
                    echo "Application is running"
                    else
                    echo "Error: Application is not reachable"
                    fi
                }
            }
        }
        
        stage('Publish Docker Image')
        {
            steps {

            echo "Stop and remove container"
            sh "docker stop ${ID}"
            
            echo "Pushing ${DOCKER_REG}/${IMAGE_NAME} image to registry"
            sh "docker push ${DOCKER_REG}/${IMAGE_NAME}"
            
            }
        }
        
        stage('Deploy to Dev') {
            steps {

                // Validate kubectl
            withCredentials([file(credentialsId: 'afc31d60-b3f2-422d-9efe-deea4591a4bc', variable: 'KUBECONFIG')]) {
                bat "minikube kubectl --  get pods -A"
		        }
            }
	    }
	    
	    stage('Deploy to dev') {
            steps {
                script {
                    namespace = 'dev'

                    echo "Deploying application ${IMAGE_NAME} to ${namespace} namespace"
                    createNamespace (namespace)


                    // Deploy with helm
                    echo "Deploying"
                    helmInstall(namespace, "${IMAGE_NAME}")
                }
            }
        }
        
        stage('Deploy to UAT') {
            steps {
                script {
                    namespace = 'uat'

                    echo "Deploying application ${IMAGE_NAME} to ${namespace} namespace"
                    createNamespace (namespace)


                    // Deploy with helm
                    echo "Deploying"
                    helmInstall(namespace, "${IMAGE_NAME}")
                }
            }
        }
        
        
        stage('Deploy to Prod') {
        input 'Proceed and deploy to Production?'    

            steps {
                // Prevent any older builds from deploying to production
                namespace = 'production'
                echo "Deploying application ${IMAGE_NAME}:${DOCKER_TAG} to ${namespace} namespace"
                createNamespace (namespace)
                
                echo "Deploying"
                helmInstall(namespace, "${IMAGE_NAME}")
                
            }
        }
        
    }    

}

