
# 🚀 URL Shortener – A Scalable, Production-Ready DevOps Project  

## Overview  
This project is a fully functional **URL Shortener** designed with **scalability, automation, and best DevOps practices** in mind. It follows a **modern cloud-native architecture** and incorporates **CI/CD, containerization, Kubernetes orchestration, and Infrastructure-as-Code (IaC)** to provide a real-world DevOps experience.  

## 🎯 **Key Features**  
- **RESTful APIs** built with **Golang** for URL shortening and redirection  
- **MySQL** database for persistent storage of shortened URLs  
- **Dockerized** application for portability and easy deployment  
- **Kubernetes** manifests for container orchestration and auto-scaling  
- **GitHub Actions CI/CD** pipeline for automated build, testing, and deployment  
- **Terraform** for Infrastructure-as-Code (**IaC**) to manage cloud infrastructure  
- **Structured logging** using Zap for better observability  
- **Monitoring & Alerting** integration using Prometheus & Grafana (Upcoming Improvement)  

## 🏗 **Architecture Overview**  
The application follows a **microservices-style architecture** with **modular components**:  
```
.
├── api-services
│   └── main.go               # Core application logic
├── handlers
│   └── handlers.go           # API request handlers
├── models
│   └── url_mappings.go       # Database model definitions
├── dbutils
│   └── db.go                 # Database connection and queries
├── Dockerfile                # Docker build configuration
├── docker-compose.yml        # Local multi-service setup
├── kubernetes
│   ├── app-deployment.yaml   # Kubernetes Deployment for the app
│   ├── service.yaml          # Kubernetes Service for load balancing
│   ├── db-deployment.yaml    # Kubernetes Deployment for MySQL
│   ├── db-service.yaml       # Kubernetes Service for MySQL
├── terraform
│   ├── modules
│   │   ├── eks               # Kubernetes cluster setup
│   │   ├── rds               # Database infrastructure
│   │   └── vpc               # Networking setup
│   ├── main.tf               # Terraform entry point
│   ├── outputs.tf
│   └── variables.tf
├── .github
│   ├── workflows
│   │   ├── ci-cd.yaml        # GitHub Actions CI/CD pipeline
└── README.md                 # Project documentation
```

## 🔧 **Tech Stack & Tools Used**  
| Category           | Technologies Used |
|--------------------|------------------|
| **Programming**   | Golang |
| **Database**      | MySQL |
| **Containerization** | Docker |
| **Orchestration** | Kubernetes |
| **CI/CD** | GitHub Actions |
| **Infrastructure as Code (IaC)** | Terraform |
| **API Framework** | Gin (Golang) |

---

## 🚀 **Deployment Process**  

### **1️⃣ Running Locally (Docker Compose)**  
To test the app locally with **Docker Compose**, run:  
```sh
docker-compose up --build
```
The API will be available at `http://localhost:8080`.  

### **2️⃣ Kubernetes Deployment**  
Ensure you have `kubectl` and `kubeconfig` set up. Deploy using:  
```sh
kubectl apply -f kubernetes/
```
Verify deployment:  
```sh
kubectl get pods
kubectl get services
```

### **3️⃣ CI/CD Pipeline (GitHub Actions)**  
The **GitHub Actions workflow** automates:  
✅ **Building Docker Images**  
✅ **Pushing Images to Container Registry**  
✅ **Deploying to Kubernetes**  

To trigger the CI/CD pipeline, simply **push changes to GitHub**!  

---

## 📌 **Key Challenges & How They Were Solved**  

### **✅ Scalable & Stateless Microservice Architecture**  
The application was **containerized** using Docker and orchestrated with Kubernetes to allow **scalability, failover handling, and ease of deployment**.  

### **✅ Database Layer Optimization**  
Implemented efficient **MySQL indexing & query optimization** to improve lookup speeds and prevent performance bottlenecks.  

### **✅ Automating Deployments with CI/CD**  
Leveraged **GitHub Actions** to automate testing, image building, and Kubernetes deployments, ensuring a **fast and reliable DevOps workflow**.  





---

## 🤝 **Contributing & Feedback**  
Contributions & feedback are welcome! Feel free to **open issues** or **submit pull requests** to improve the project.  

---


Note: Certain changes and modifications still required and this is a continuos onging project. 
