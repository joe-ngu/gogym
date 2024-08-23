<h1 align="center">
  <br>
  <a href="https://github.com/joe-ngu/gogym"><img src="https://raw.githubusercontent.com/joe-ngu/gogym/main/frontend/public/gogym.png" alt="Gogym" width="200"></a>
  <br>
  Gogym
  <br>
</h1>

<h4 align="center">A Golang REST API for Fitness Tracking</h4>

<p align="center">
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Golang"/>
  </a>
  <a href="https://www.postgresql.org/">
      <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL"/>
  </a>
  <a href="https://jwt.io">
    <img src="https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens" alt="JWT"/>
  </a>
  <a href="https://react.dev/">
    <img src="https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB/">
  </a>
  <a href="https://typescriptlang.org">
    <img src="https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white"/>
  </a>
  <a href="https://tailwindcss.com/">
    <img src="https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white"/>
  </a>
</p>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#quickstart">Quick Start</a> •
  <a href="#design">Design</a> •
  <a href="#credits">Credits</a> •
  <a href="#license">License</a>
</p>

<!-- TO DO: Add gif here -->
![screenshot]()

## Key Features
* **Exercise Library** - Create an extensive list of exercises and assign them to muscle groups
* **Workout Log** - Construct unique workouts and log the number of sets, reps, and load for each exercise of the workout
* **Secure Authentication** - JWT-based authentication for secure access and protected routes
    - users only have access to their own workouts 
    - only authenticated users can add to the exercise library
* **Dockerized Deployment** - Containerized using Docker for consistent and portable deployment across various environments
* **Error Handling** - Comprehensive error handling for smooth API interactions

## Quickstart

To clone and run this application, you'll need [Git](https://git-scm.com) and [Docker](https://docs.docker.com/get-docker/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/joe-ngu/gogym.git

# Go into the repository
$ cd gogym

# Run the quickstart script
$ ./quickstart

```

## Design


## Credits


## License

MIT