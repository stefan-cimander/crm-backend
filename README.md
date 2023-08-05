# CRM Backend

Simple CRM backend written in Go for the capstone project of the Udacity Golang course.

## Getting Started

This CRM backend project provides a basic HTTP server implementation for the following API endpoints:

HTTP Method | URL Path        | Description
----------- | --------------- | ----------------------------
GET         | /customers      | Gets a list of all customers
GET         | /customers/{id} | Gets a single customer
POST        | /customers      | Adds a new customer
PUT         | /customers/{id} | Updates an existing customer
DELETE      | /customers/{id} | Deletes a customer

## Local Setup

When setting up the project on your computer, first of all make sure you have **Go** installed. You can download the latest version 
[here](https://go.dev/dl/) and verify the installation by running the `go version` command in a terminal window.

In order to run the CRM backend on your local machine, navigate to the project root directory and run the `go run main.go` command.
The CRM backend HTTP server can then be reached at http://localhost:3000.

The project's unit tests can be executed by running `go test`.

## Build Status

[![Go](https://github.com/stefan-cimander/crm-backend/actions/workflows/go.yml/badge.svg)](https://github.com/stefan-cimander/crm-backend/actions/workflows/go.yml)