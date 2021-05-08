# CI/CD  Kubernetes pipleine for GOLANG Application
go-calc is a simple 2 operand calculator over http written in go-lang. It has basic functionality of arithmetic operations like add, substract, divide, multiply and divide.


The go-lang application is packed as a image by building the dockerfile which is then deployed into Kubernetes cluster using Helm Charts.
The application is deployed into three different environments which are - Dev, Staging and Production.

We have make use of the following tools for this process - </br>
**[Jenkins]** - We basically setup a master agent with the following of the tools involved (i.e. Docker, Kubectl, Helm) </br>

**[Docker]** - Which is used to build and publish image to the Docker Hub Repository. </br>

**[Kubernetes]** - The Kubeconfig file is loaded into the Jenkins pipeline which acts as a authentication for accesssing kubernetes cluster and through the Kubernetes CLI tool (i.e. kubectl) building the kubernetes resources.</br>

**[Helm Charts]** - This bascially bundles or package up the resource needed for building up the application. It also basically links the kubernetes objects which has been deployed for the kubernetes cluster.</br>

* The Jenkins pipeline basically builds the docker image through the Dockerfile provided in the Source Code and run a container locally to test the image wether it is accessible or not.</br>
* Then, the Jenkins pushes the image `a5edevopstuts/gocalc` to the docker hub  repository. </br>
* Jenkins then through the `KUBECONFIG` file access the kubernetes cluster and creates three namespaces - dev, stag and prod. </br>
* Simultaneously, Jenkins install or upgrade the Helm chart into the particular namespaces. </br>
* Helm charts contains three values.yaml file - `values-dev.yaml`, `values-stag.yaml` and `values-prod.yaml` which differs according to their particular namespaces.

### DEV Namespace
* Kubernetes objects deployed 

<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117549944-f6da9f00-b05a-11eb-82ee-200c2deefd48.PNG">

* Application endpoint IP address 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117550135-2ccc5300-b05c-11eb-9868-514d9b3057f0.PNG">

* Sample inputs for the application 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117550138-335aca80-b05c-11eb-93c7-621019e28276.PNG">


