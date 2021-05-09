    //function to retrieve ip of the application deployed
def getnodehost (namespace, servicename) {

	script {
		sh """
            export NODE_PORT="\$(kubectl get --namespace ${namespace} -o jsonpath="{.spec.ports[0].nodePort}" services ${servicename})"
        """
        sh """ 
            export NODE_IP="\$(kubectl get nodes --namespace ${namespace} -o jsonpath="{.items[0].status.addresses[0].address}")"
        """
        host_ip = "http://$NODE_IP:$NODE_PORT/status"		    
	    }
	return "$host_ip"
    }

// function to test the particular application status
    def runcurl (ip_host) {

	    script{
            repsone=sh(script:"curl -k -s -X GET --url http://$ip_host")
            
            if ($response == '200')
                echo "Application is running"
            else
                echo "Application is unreachable"
                currentBuild.result = 'ABORTED'
            fi
	    }   
    }


// function to create the namespace
  def createNamespace (namespace) {
        
        echo "Creating namespace ${namespace} if needed"

        sh "[ ! -z \"\$(kubectl get ns ${namespace} -o name 2>/dev/null)\" ] || kubectl create ns ${namespace}"
   }
    

//function to install or upgrade the application according to the namespace and helm release
    def helmInstall (namespace, release) {
        
	script {
        release = "${release}-${namespace}"
        echo "Installing ${release} in ${namespace}"
        sh """
            helm upgrade --install --namespace ${namespace} ${release} \
             -f ./helm/gocalc/values-${namespace}.yaml                   \
                 ./helm/gocalc
        """
        sh "sleep 5"
	    }
    }
    


    
pipeline {
    
    parameters {
        string (name: 'GIT_BRANCH', defaultValue: 'main',  description: 'Git branch to build')

    }
    
    environment {
	DOCKER_REG = 'a5edevopstuts' 
        IMAGE_NAME = 'gocalc'  
        TEST_LOCAL_PORT = '8081' //local port to test docker image locally 
        ID = "${DOCKER_REG}-${IMAGE_NAME}"  // container ID for running the docker image locally    
    }
    
    agent { node { label 'master' } }

    stages {
        
        stage('Git checkout'){
            steps{
                echo "checkout code"
                git branch: 'main', credentialsId: 'GitHubID', url: 'https://github.com/arif11111/Gocalc-Helm-.git'
                }
               
            }
        
        stage('Build'){
            steps{
                
                echo "Building Image"
                sh " docker build -t ${DOCKER_REG}/${IMAGE_NAME} .  "
                    
                echo "Running tests"

                // Kill container in case there is a leftover
                sh "[ -z \"\$(docker ps -a | grep ${ID} 2>/dev/null)\" ] || docker rm -f ${ID}"

                echo "Starting ${IMAGE_NAME} container"
                sh "docker run --detach --name ${ID} --rm --p ${TEST_LOCAL_PORT}:8080 ${DOCKER_REG}/${IMAGE_NAME}"
                        
            }
        }
        
        stage('Local Tests') {
            
            steps {
	        echo "Running Docker image locally"
		
		//accessing the status of the application
                script{
                    host_ip = "localhost:${TEST_LOCAL_PORT}/status"
                    runcurl (host_ip)
                }
            }
        }
        
        stage('Publish Docker Image')
        {
            steps {

            echo "Stop and remove container"
            sh "docker stop ${ID}"
            
            echo "Pushing ${DOCKER_REG}/${IMAGE_NAME} image to registry"
             withCredentials([usernameColonPassword(credentialsId: 'dockerID', variable: 'docker_credn')]) {
                    sh "docker login"
                    sh "docker push ${DOCKER_REG}/${IMAGE_NAME}"
                }
            }
        }
        
       
	    
	    stage('Deploy to dev') {
            steps {
                script {
                    namespace = 'dev'

                    echo "Creating namespace ${namespace}"
                    createNamespace (namespace)

                    echo "Deploying application ${IMAGE_NAME} to ${namespace} namespace"
                    helmInstall (namespace, "${IMAGE_NAME}")
                }
            }
        }
        
        
	stage('Development Env Test')
        {
            steps{
                script{
		    echo "Accessing the status of application in ${namespace} namespace"
		            
                    servicename = "${IMAGE_NAME}-${namespace}-service"
                    
                    host_ip = getnodehost (namespace, servicename)

                    runcurl(host_ip)
                    }
            }
        }

        stage('Deploy to Staging') {
            steps {
                script {
                    namespace = 'staging'

                    echo "Creating namespace ${namespace}"
                    createNamespace(namespace)

                    echo "Deploying application ${IMAGE_NAME} to ${namespace} namespace"
                    helmInstall(namespace, "${IMAGE_NAME}")
                }
            }
        }
        
        
        
        stage('Staging Env test')
        {
            steps{
                script{
		    echo "Accessing the status of application in ${namespace} namespace"
		            
                    servicename = "${IMAGE_NAME}-${namespace}-service"
                    
                    host_ip = getnodehost (namespace, servicename)

                    runcurl(host_ip)
                    }
            }
        }
        
        
        
        stage('Deploy to Prod') {
            steps {
                input 'Proceed and deploy to Production?' 
                
                script{
                    namespace = 'production'

                    echo "Creating namespace ${namespace}"
                    createNamespace (namespace)
                
                    echo "Deploying application ${IMAGE_NAME} to ${namespace} namespace"
                    helmInstall(namespace, "${IMAGE_NAME}")
                }
            }
        }
        
        stage('Prod Env test')
        {
            steps{
                script{
		    echo "Accessing the status of application in ${namespace} namespace"
		            
                    servicename = "${IMAGE_NAME}-${namespace}-service"
                    
                    host_ip = getnodehost ("${namespace}", servicename)

                    runcurl(host_ip)
                    }
            }
        }
        
    }    

}
