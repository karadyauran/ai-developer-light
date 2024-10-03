import os
import sys
import subprocess
from random import randint


def main():
    script_path = os.path.abspath('main.py')
    command = f'python {script_path}'

    hours = randint(10, 19)
    minutes = randint(0, 59)

    schedule = f'{minutes} {hours} * * *'

    cron_job = f'{schedule} {command}'

    try:
        existing_cron = subprocess.check_output(
            ['crontab', '-l'], stderr=subprocess.STDOUT).decode()
    except subprocess.CalledProcessError:
        existing_cron = ''

    cron_lines = existing_cron.strip().split('\n')
    cron_lines = [
        line for line in cron_lines if script_path not in line and line.strip() != '']

    cron_lines.append(cron_job)

    new_cron = existing_cron + '\n' + cron_job + '\n'

    process = subprocess.Popen(['crontab', '-'], stdin=subprocess.PIPE)
    process.communicate(new_cron.encode())


if __name__ == '__main__':
    main()
