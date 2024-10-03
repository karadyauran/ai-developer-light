import os
import sys
import subprocess
from random import randint


def main():
    venv_python = os.path.join(os.getcwd(), '.venv', 'bin', 'python')
    script_path = os.path.abspath('main.py')
    command = f'{venv_python} {script_path}'
    logs = f'>> {os.path.abspath("logs")}/main.log 2>&1'

    hours = randint(10, 19)
    minutes = randint(0, 59)

    schedule = f'{minutes} {hours} * * *'

    cron_job = f'{schedule} {command} {logs}'

    try:
        existing_cron = subprocess.check_output(
            ['crontab', '-l'], stderr=subprocess.STDOUT).decode()
    except subprocess.CalledProcessError:
        existing_cron = ''

    cron_lines = [
        line for line in existing_cron.strip().split('\n')
        if script_path not in line and line.strip() != ''
    ]

    cron_lines.append(cron_job)

    new_cron = '\n'.join(cron_lines) + '\n'

    process = subprocess.Popen(['crontab', '-'], stdin=subprocess.PIPE)
    process.communicate(new_cron.encode())


if __name__ == '__main__':
    main()