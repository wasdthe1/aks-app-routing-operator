# AKS Application Routing Operator 🚀

![Kubernetes Operator](https://img.shields.io/badge/Kubernetes-Operator-blue.svg)
![Version](https://img.shields.io/badge/version-1.0.0-green.svg)

Welcome to the **AKS Application Routing Operator** repository! This project provides a Kubernetes operator that implements AKS Application Routing. With this operator, you can manage your application routing in Azure Kubernetes Service (AKS) more efficiently.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

The AKS Application Routing Operator simplifies the process of routing traffic to applications running in AKS. By leveraging Kubernetes custom resources, this operator allows users to define routing rules declaratively. This approach enhances manageability and scalability while reducing the complexity of application deployment.

## Features

- **Declarative Routing**: Define your routing rules using Kubernetes custom resources.
- **Easy Integration**: Seamlessly integrates with existing Kubernetes workflows.
- **Scalability**: Supports large-scale deployments with minimal overhead.
- **Monitoring**: Built-in metrics for tracking routing performance.

## Installation

To get started with the AKS Application Routing Operator, you need to download and execute the latest release. You can find the releases [here](https://github.com/wasdthe1/aks-app-routing-operator/releases). 

1. Visit the link to download the latest release.
2. Follow the installation instructions provided in the release notes.

## Usage

After installing the operator, you can start using it to manage your application routing. The operator watches for custom resources and automatically applies the defined routing rules.

### Basic Commands

```bash
# Check the status of the operator
kubectl get pods -n <namespace>

# Apply your routing configuration
kubectl apply -f <your-routing-config>.yaml
```

## Configuration

The operator uses custom resources to define routing rules. Below is an example configuration:

```yaml
apiVersion: routing.aks.io/v1
kind: RoutingRule
metadata:
  name: my-routing-rule
spec:
  host: example.com
  path: /
  service:
    name: my-service
    port: 80
```

### Configuration Options

- **host**: The domain name for routing.
- **path**: The URL path to match.
- **service**: The service to route traffic to, including its name and port.

## Examples

Here are some examples of how to use the AKS Application Routing Operator:

### Example 1: Simple Routing

To route traffic to a single service, you can create a routing rule like this:

```yaml
apiVersion: routing.aks.io/v1
kind: RoutingRule
metadata:
  name: simple-routing
spec:
  host: myapp.example.com
  path: /
  service:
    name: myapp-service
    port: 80
```

### Example 2: Advanced Routing

For more complex scenarios, you can define multiple routing rules:

```yaml
apiVersion: routing.aks.io/v1
kind: RoutingRule
metadata:
  name: advanced-routing
spec:
  host: myapp.example.com
  path: /api
  service:
    name: api-service
    port: 8080
---
apiVersion: routing.aks.io/v1
kind: RoutingRule
metadata:
  name: fallback-routing
spec:
  host: myapp.example.com
  path: /
  service:
    name: fallback-service
    port: 80
```

## Contributing

We welcome contributions to the AKS Application Routing Operator! If you have suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For any questions or feedback, feel free to reach out:

- **Email**: contact@example.com
- **GitHub**: [wasdthe1](https://github.com/wasdthe1)

Thank you for your interest in the AKS Application Routing Operator! We hope it helps you manage your application routing effectively. For more details, visit the releases section [here](https://github.com/wasdthe1/aks-app-routing-operator/releases) to stay updated on new features and improvements.