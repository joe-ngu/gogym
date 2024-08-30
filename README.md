<h1 align="center">
  <br>
  <a href="https://github.com/joe-ngu/gogym"><img src="https://raw.githubusercontent.com/joe-ngu/gogym/main/assets/logo.png" alt="Gogym" width="200"></a>
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

## Key Features
- **Extensive Exercise Library**  
  Build a comprehensive library of exercises, categorized by muscle groups, for easy management and selection.

- **Custom Workout Logging**  
  Design personalized workouts by selecting exercises and logging specific details such as sets, reps, and weight used for each session.

- **Robust Authentication and Security**  
  - Secure user access through JWT-based authentication, ensuring protected routes.
  - Users can only view and manage their own workouts, maintaining privacy.
  - Exercise library management is restricted to authenticated users only.

- **Dockerized Deployment**  
  Deploy the application seamlessly across different environments with Docker, ensuring consistency and portability.

- **Comprehensive Error Handling**  
  Implemented thorough error management to enhance user experience by handling API errors smoothly and efficiently.

## Quickstart

To clone and run this application, you'll need [Git](https://git-scm.com) and [Docker](https://docs.docker.com/get-docker/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/joe-ngu/gogym.git

# Go into the repository
$ cd gogym

# Run the quickstart script
$ make quickstart

```

## Design

### Application Architecture

This application is built with a modern, scalable architecture designed to handle both present requirements and future expansions:

- **Backend**: The backend is developed using Go, chosen for its performance and concurrency support. It handles API requests, business logic, and database interactions.

- **Frontend**: The frontend is built with React, providing a dynamic and responsive user interface. React was selected for its component-based architecture, which enables the creation of reusable UI elements and efficient state management.

- **Database**: PostgreSQL serves as the relational database, known for its robustness, reliability, and support for complex queries. The database schema is designed to efficiently manage the application’s data, ensuring data integrity and scalability.

### Database Structure

The following diagram illustrates the PostgreSQL table structure used in this application, detailing the relationships and key fields:

![PostgreSQL Table Structure](https://raw.githubusercontent.com/joe-ngu/gogym/main/assets/db_diagram.png)


## Credits
This project was my first attempt at building an API in Golang and I learned a lot from various resources, documentation, and tutorials provided by the community of Golang developers. Special thanks to the following individuals:

- **[AnthonyGG](https://github.com/anthdm)** 
- **[Tiago](https://github.com/sikozonpc)** 



## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.