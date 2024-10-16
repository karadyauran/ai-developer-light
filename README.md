# AI Developer Light

AI Developer Light is an automated code generation tool that leverages OpenAI’s GPT models to generate random applications in various programming languages. The tool creates application ideas, structures, and code, organizing them into projects with Git repositories. It also includes functionality to schedule automatic code generation using cron jobs.

Features

- Random Application Generation: Generates random applications in languages like Python, Go, JavaScript, Ruby, and Java.
- Automated Code Creation: Uses OpenAI’s GPT models to generate code for each file in the application.
- Dynamic Project Structure: Randomly determines the number of files and how the application’s logic is divided among them.
- Git Integration: Initializes a Git repository for each project, adds files, commits changes, and can push to a remote repository.
- Cron Scheduling: Includes a script to set up a cron job that schedules the code generation process at random times.
- Configurable Prompts: Prompts used for code generation are stored in separate text files for easy customization.