import os
import random
import subprocess
import openai
import uuid
from dotenv import load_dotenv

class AppGenerator:
    def __init__(self, model="gpt-4o"):
        load_dotenv()
        api_key = os.getenv('OPENAI_API_KEY')
        if not api_key:
            raise ValueError("API-key for Openai is not in .env")
        openai.api_key = api_key

        self.model = model
        self.languages = ["Go", "Python", "JavaScript"]
        self.history = [
            {"role": "system", "content": "You are a senior software developer proficient in multiple programming languages."}
        ]

        self.prompts = {
            'app_idea': self.load_prompt('scripts/app_idea_prompt.txt'),
            'file_structure': self.load_prompt('scripts/file_structure_prompt.txt'),
            'code_generation': self.load_prompt('scripts/code_generation_prompt.txt'),
            'commit_message': self.load_prompt('scripts/commit_message_prompt.txt'),
        }

    def load_prompt(self, filepath):
        with open(filepath, 'r', encoding='utf-8') as file:
            return file.read()

    def send_request(self, prompt, max_tokens):
        try:
            response = openai.chat.completions.create(
                model=self.model,
                messages=self.history + [{"role": "user", "content": prompt}],
                max_tokens=max_tokens,
                temperature=0.7,
            )
            answer = response.choices[0].message.content.strip()
            self.history.append({"role": "user", "content": prompt})
            self.history.append({"role": "assistant", "content": answer})
            return answer
        except Exception as e:
            print(f"Error to request OpenAI API: {e}")
            return ""

    def generate_unique_app_name(self, base_name):
        unique_id = str(uuid.uuid4())[:8]
        return f"{base_name}_{unique_id}"

    def generate_app_idea(self, language, n_files):
        prompt = self.prompts['app_idea'].format(language=language, n_files=n_files)
        return self.send_request(prompt, 150)

    def generate_file_structure(self, language, n_files):
        ext = self.get_file_extension(language)
        prompt = self.prompts['file_structure'].format(language=language, n_files=n_files, ext=ext)
        response = self.send_request(prompt, 50)
        return [fn.strip() for fn in response.split(',') if fn.strip()]

    def generate_code(self, file_name, language):
        prompt = self.prompts['code_generation'].format(file_name=file_name, language=language)
        responce = self.send_request(prompt, 1000)

        lines = responce.splitlines()
        answer = "\n".join(lines[1:-1])

        return answer

    def create_commit_message(self, file_name):
        prompt = self.prompts['commit_message'].format(file_name=file_name)
        return self.send_request(prompt, 50)

    def get_file_extension(self, language):
        extensions = {
            "Go": "go",
            "Python": "py",
            "JavaScript": "js"
        }
        return extensions.get(language, "txt")

    def create_repository(self, app_name):
        base_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'projects')
        app_dir = os.path.join(base_dir, app_name)
        os.makedirs(app_dir, exist_ok=True)
        os.chdir(app_dir)