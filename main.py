import os
from app_generator import AppGenerator
import random
import subprocess
from dotenv import load_dotenv
from time import sleep


def main():
    load_dotenv()
    mode = os.getenv('MODE')

    generator = AppGenerator()
    language = random.choice(generator.languages)
    n_files = random.randint(4, 7)

    idea_response = generator.generate_app_idea(language, n_files)
    idea_lines = idea_response.strip().splitlines()
    if not idea_lines:
        print("Cannot get app idea.")
        return
    base_app_name = idea_lines[0].strip()
    app_description = "\n".join(idea_lines[1:]).strip()

    app_name = generator.generate_unique_app_name(base_app_name)

    generator.create_repository(app_name)

    n_files_logic = random.randint(4, n_files)
    file_names = generator.generate_file_structure(language, n_files_logic)

    for file_name in file_names:
        code = generator.generate_code(file_name, language)
        if code:
            with open(file_name, 'w') as f:
                f.write(code)

            if mode == 'prod':

                subprocess.run(['git', 'add', file_name])
                commit_message = "ai: " + generator.create_commit_message(file_name)
                subprocess.run(['git', 'commit', '-m', commit_message])

                delay_in_seconds = random.randint(30, 120)
                sleep(delay_in_seconds)

    if mode == 'prod':
        subprocess.run(['git', 'push', '-u', 'origin', 'main'])


if __name__ == "__main__":
    main()
