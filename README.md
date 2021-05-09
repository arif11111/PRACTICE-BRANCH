# CI/CD  Kubernetes pipleine for GOLANG Application
go-calc is a simple 2 operand calculator over http written in go-lang. It has basic functionality of arithmetic operations like add, substract, divide, multiply and divide.


The go-lang application is packed as a image by building the dockerfile which is then deployed into Kubernetes cluster using Helm Charts.
The application is deployed into three different environments which are - Dev, Staging and Production.

We have make use of the following tools for this process - </br>
**[Terraform]** - For spinning up EKS cluster and corresponding the supporting VPC modules in AWS.</br>

**[Jenkins]** - We basically setup a master agent with the following of the tools involved (i.e. Docker, Kubectl, Helm) </br>

**[Docker]** - Which is used to build and publish image to the Docker Hub Repository. </br>

**[Kubernetes]** - The Kubeconfig file is loaded into the Jenkins pipeline which acts as a authentication for accesssing kubernetes cluster and through the Kubernetes CLI tool (i.e. kubectl) building the kubernetes resources.</br>

**[Helm Charts]** - This bascially bundles or package up the resource needed for building up the application. It also basically links the kubernetes objects which has been deployed for the kubernetes cluster.</br>


<h3> Creating AWS VPC and EKS cluster </h3>

**Provide AWS Credentials**

```
$ aws configure:
$............... AWS Access Key ID [None]:
$............... AWS Secret Access Key [None]:
$............... Default region name [None]:
$............... Default output format [None]:
```

Initialize and pull terraform cloud specific dependencies:
```
$ terraform init
```
View terraform plan:
```
$ terraform plan
```
Apply terraform plan:
```
$ terraform apply
```

Terraform modules will create
```
-   VPC
-   Subnets
-   Routes
-   IAM Roles for master and nodes
-   Security Groups "Firewall" to allow master and nodes to communicate
-   EKS cluster
-   Autoscaling Group will create nodes to be added to the cluster
```
A file `kubeconfig_my-cluster` is available in the directory which is the kubeconfig for the newly created cluster.

Export the KUBECONFIG variable to the above kubeconfig file to access the cluster.<br>
```
$ export KUBECONFIG="${PWD}/kubeconfig_my-cluster"`
````
<h2> Jenkins CI/CD for deploying golang application </h2> 
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

### Stag Namespace
* Kubernetes objects deployed 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117570078-81b4ab80-b0e6-11eb-9b34-8c83a3c885bb.PNG">

* Application endpoint IP address 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117570105-95601200-b0e6-11eb-84b0-996ddd1161e6.PNG">

* Sample inputs for the application 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117570122-9e50e380-b0e6-11eb-82b1-4320dcef559d.PNG">


### Prod Namespace
* Kubernetes objects deployed 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117570277-0c95a600-b0e7-11eb-8a64-a801c1c39416.PNG">

* Application endpoint IP address 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117570291-13bcb400-b0e7-11eb-8dae-c54238e605d9.PNG">

* Sample inputs for the application 
<img width="630" alt="Capture" src="https://user-images.githubusercontent.com/33144027/117570297-19b29500-b0e7-11eb-94cf-51b0a6be8f8c.PNG">




