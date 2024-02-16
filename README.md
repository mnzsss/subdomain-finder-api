# Subdomain Finder API with Subfinder and Go

## Overview

This project demonstrates how to use [Subfinder](https://github.com/projectdiscovery/subfinder), a subdomain discovery tool written in Go, to find subdomains of a given domain. Additionally, it includes instructions on how to deploy the project on [Vercel](https://vercel.com/).

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your machine.
- A Vercel account. You can sign up [here](https://vercel.com/signup).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mnzsss/subdomain-finder-api.git
   cd subdomain-finder-api
    ```

2. Install the Go dependencies:
   ```bash
   go mod tidy
   ```

3. Run the project locally:
   ```bash
    go run main.go
    ```

4. Open your browser and navigate to `http://localhost:3333/api/subdomains?domain=example.com`. You should see a JSON response with the subdomains of `example.com`.

## Deployment

1. Install the Vercel CLI:

   ```bash
   npm install -g vercel
   ```

2. Deploy the project:

   ```bash
    vercel init
    ```

3. Follow the instructions to deploy the project to Vercel.
4. Once deployed, open your browser and navigate to `https://your-vercel-deployment-url/api/subdomains?domain=example.com`. You should see a JSON response with the subdomains of `example.com`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


