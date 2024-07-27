
<h1 align="center">
  Todo List using Go + React
</h1>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a>
</p>

![screenshot](https://github.com/felipeganho/to-do-list/blob/main/images/home.png)

## Key Features

-   ⚙️ Tech Stack: Go, React, TypeScript, MongoDB, TanStack Query, ChakraUI
-   ✅ Create, Read, Update, and Delete (CRUD) functionality for todos
-   🌓 Light and Dark mode for user interface
-   📱 Responsive design for various screen sizes
-   🌐 Deployment
-   🔄 Real-time data fetching, caching, and updates with TanStack Query
-   🎨 Stylish UI components with ChakraUI

## How To Use

To clone and run this application, you'll need [Git](https://git-scm.com), [Node.js](https://nodejs.org/en/download/) (which comes with [npm](http://npmjs.com)) and [Go](https://go.dev/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/felipeganho/to-do-list

# Go into the repository
$ cd to-do-list

# Copy .env.example
$ cp .env.example .env

# Install dependencies
$ npm install

# Need a Mongo DB string connection
$ Set MONGODB_URI on .env

# Add alias on your .bashrc
$ alias air='~/go/bin/air'

# Start backend with Go
$ air

# Run the app
$ npm run dev
```
