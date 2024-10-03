# AI Developer Light

AI Developer Light is an automated code generation tool that leverages OpenAI’s GPT models to generate random applications in various programming languages. The tool creates application ideas, structures, and code, organizing them into projects with Git repositories. It also includes functionality to schedule automatic code generation using cron jobs.

Features

- Random Application Generation: Generates random applications in languages like Python, Go, JavaScript, Ruby, and Java.
- Automated Code Creation: Uses OpenAI’s GPT models to generate code for each file in the application.
- Dynamic Project Structure: Randomly determines the number of files and how the application’s logic is divided among them.
- Git Integration: Initializes a Git repository for each project, adds files, commits changes, and can push to a remote repository.
- Cron Scheduling: Includes a script to set up a cron job that schedules the code generation process at random times.
- Configurable Prompts: Prompts used for code generation are stored in separate text files for easy customization.
- Testing: Comes with test scripts to ensure the functionality of the main components.

## Project Structure
```
ai-developer-light/
├── app_generator.py
├── main.py
├── setup_cron.py
├── prompts/
│   ├── app_idea_prompt.txt
│   ├── code_generation_prompt.txt
│   ├── commit_message_prompt.txt
│   └── file_structure_prompt.txt
├── projects/
│   └── [Generated Projects]
├── tests/
│   ├── test_main.py
│   └── test_setup_cron.py
├── .env
└── requirements.txt
```

- app_generator.py: Contains the AppGenerator class responsible for interacting with the OpenAI API and generating application components.
- main.py: The main script that orchestrates the code generation process.
- setup_cron.py: Script to set up a cron job that schedules the execution of main.py.
- prompts/: Directory containing prompt templates used by AppGenerator.
- projects/: Directory where generated projects are stored.
- tests/: Contains test scripts for main.py and setup_cron.py.
- .env: Environment file containing the OpenAI API key.
- requirements.txt: Lists the Python dependencies required for the project.

Installation

1. Clone the Repository

```
git clone https://github.com/karadyauran/ai-developer-light.git
cd ai-developer-light
```


2. Create a Virtual Environment

```
python3 -m venv .venv
source .venv/bin/activate
```

3. Install Dependencies

```
pip install -r requirements.txt
```

4.	Set Up OpenAI API Key
 - Create a .env file in the root directory:

```
OPENAI_API_KEY=your_openai_api_key_here
```

- Replace your_openai_api_key_here with your actual OpenAI API key.

## Usage

### Generate a Project Manually

Run the ```main.py``` script to generate a new project:

```python main.py```

#### This will:

 - Generate a random application idea.
 - Create a unique project name and directory in projects/.
 - Generate code files based on the prompts.
 - Initialize a Git repository and commit the generated files.

### Schedule Automatic Generation

Use the ```setup_cron.py``` script to schedule automatic code generation:

```python setup_cron.py```

#### This script will:

- Set up a cron job that runs main.py at a random time between 10:00 and 19:59 every day.
- The scheduled times are randomized to provide variability.

Customize Prompts

Prompts used for code generation are stored in the prompts/ directory. You can edit these .txt files to customize the behavior and output of the code generator.

### Run Tests

To run the tests, execute the test scripts:

```
python tests/test_main.py
python tests/test_setup_cron.py
```

## Requirements

- Python 3.x
- An OpenAI API key
- Git (optional, for version control and pushing to remote repositories)
- Cron (for scheduling, on Unix-based systems)

## Important Notes

- API Key Security: Keep your OpenAI API key secure. Do not commit it to version control or share it publicly.
- .env File: Ensure that the .env file is included in .gitignore to prevent accidental commits.
- Permissions: The scripts assume you have the necessary permissions to create directories, modify crontab, and execute scripts on your system.
- OpenAI API Usage: Be aware of the costs associated with using the OpenAI API and monitor your usage accordingly.

## Contact

For questions or suggestions, please open an issue on the [GitHub repository](https://github.com/karadyauran/ai-developer-light).