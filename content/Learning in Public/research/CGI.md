---
title: CGI
date created: Saturday, January 27th 2024, 11:55:48 am
date modified: Saturday, January 27th 2024, 3:14:25 pm
---

## What is It?

> No, not Computer Generated Images!

A Common Gateway Interface is a specification that enables web servers to execute an external program to respond to HTTP requests.

The TL;DR is that whenever the server gets a web request, we run the corresponding script and capture the output.

## Sequence Diagram of a Request

```mermaid
sequenceDiagram
	participant C as Client
	participant W as Web Server

	C->>W: GET / 
	W->>C: Send index.html
	C->>W: GET /api/todos
	create participant P as Process
	W->>P: Run todos.lua
    destroy P
    W->>P: Read stdout
	W->>C: Send stdout 
```
