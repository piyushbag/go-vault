### ultimate go software design with kubernetes

1.0 - Introduction
Introduction to all the engineering that you will learn.
1.1: Design Philosophy, Guidelines, What to Expect
- make things easy to understand
- KISS
- precision
1.2: Tooling to Install
1.3: Images to Install

2.0 - Modules
A walkthrough of how the module ecosystem works. We talk about the engineering decisions you need to make as they relate to modules.
- related to dependencies
- GOMODCACHE - all dependencies live in the disk location
- go clean -modcache : to clear mod folder / all GOMODCACHE / clear all dependencies
- name of the repo + unique name : create namespace to access code 

2.1: Adding Dependencies
2.2: Module Mirrors
2.3: Checksum Database
2.4: Vendoring
2.5: MVS Algorithm
2.6: Upgrading Dependencies

3.0 - Deploy First Mentality
We begin to build the service with a focus on the ability to deploy the service in Kuberenetes.
3.1: Project Layers, Policies, and Guidelines
3.2: Prepare Project
3.3: Logging

4.0 - Kubernetes
We introduce Kubernetes and get a K8s environment up and running. At this point, everything we do runs in the K8s environment.
4.1: Clusters, Nodes and Pods
4.2: Start the Kuberenetes Cluster
4.3: Create/Build Dockerfile for the Service
4.4: Create/Apply K8s Deployment for the Service

5.0 - Kubernetes Quotas
We introduce applying Quotas to the deployment and discuss the problems that can result.
5.1: Setting CPU Quotas
5.2: Adjust GOMAXPROCS to match the CPU Quota

6.0 - Finish Initial Service Startup/Shutdown
We finish the inital startup and shutdown of the service.
6.1: Configuration
6.2: Debugging / Metrics
6.3: Telepresence for Cluster Access
6.4: Shutdown Signaling and Load Shedding

7.0 - Web Framework
We build out our own router by extending an existing one. This gives us a framework for injecting business logic into the processing of requests. It also allows for more consistency in the handling of requests.
7.1: Basic Structure of an HTTP Router
7.2: Liveness and Readiness Handlers
7.3: Customize the Router
7.4: Middleware Support
7.5: Sending Responses

8.0 - Middleware
We add middleware functions for business-level logic that needs to be injected into the processing of requests.
8.1: Logging
8.2: Error Handling
8.2.1: Understanding what Error Handling Means
8.2.2: Declaring Custom Error Types
8.2.3: Consistent Handling and Response
8.3: Panic Handling
8.4: Metrics

9.0 - JSON Web Tokens (JWT)
We gain an understanding of how JWT’s work and their shortcomings.
9.1: Understanding JWT
9.2: Private/Public Key Generation
9.3: Token Generation
9.4: Token Signature Validation

10 - Secrets Support
We add Hashicorp’s Vault to our K8s environment. Then we add support to our admin tooling to add our private key to Vault and configure the tool to run inside a K8s init container. Then we write a package for retrieving the private key for application use.
10.1: Kubernetes Support for Vault
10.2: Admin Support to Load Keys
10.3: Add Init Containers to Load Keys
10.4: Create Vault Package
10.5: Vault Unit Test

11 - Authentication / Authorization
We integrate authentication and authorization support into the project by developing a packages to generate and validate tokens. Then we integrate the packages into the application and test things are working.
11.1: Auth Package
11.2: Auth Unit Test
11.3: Add Middleware

12 - Database Support
We add a Postgres database to our K8s environment. Then we write a small database package that provides support for using the SQLx package more effectively. Finally, integrate the database package on application startup.
12.1: Kubernetes Support for Postgres
12.2: Create Database Package
12.3: Update Readiness Handler to Perform DB Checks

13.0 - Database Migrations and Seeding
We define our schema and provide support for migration schema changes over time. We also provide support for seeding the database. FInally, we add support in Kubernetes to run the migration and seeding on POD startup.
13.1: Maintaining Database Schemas
13.2: Seeding Data
13.3: Add Init Containers to Automate Migrations

14.0 - Business Packages
We talk about the business packages that will exist in the core layer. Then we add the core user and storage packages for managing users in the database. We provide support for adding, updating, deleting, retrieving, and caching users plus talk about all the engineering decisions that go into these types of CRUD based APIs.
14.1: Design Philosophies, Policies, and Guidelines
14.2: Core Business Package Design
14.3: Store Database Package Design
14.4: Store Cache Package Design

15.0 - Testing Data Business Packages
We add docker and unit testing support for writing unit tests against a real database and write the actual user package tests.
15.1: Support for Starting and Stopping Containers
15.2: Support for Starting and Stopping a Unit Test
15.3: Write User CRUD Data Unit Tests

16.0 - REST API
We add the web handlers for the new user CRUD support. We also write integration tests to validate everything is working.
16.1: Writing User Web Handlers
16.2: Support for Starting and Stopping an Integration Test
16.3: Write Integration Tests for Users

17.0 - Open Telemetry
We add tracing to the project by integrating Open Telemetry and Zipkin.
17.1: Integrate OTEL Web Handler into the Framework
17.2: Integrate OTEL into Service Startup
17.3: Kubernetes Support for Zipkin
17.4: Add Tracing Calls Inside Functions to Trace

18.0 - Review Service Project
Review service project and get it running
18.1: Check For Dependcy Upgrades
18.2: Rebuild and Run the Project