import unittest
from unittest.mock import patch, MagicMock
from setup_cron import main
import os

class TestSetupCron(unittest.TestCase):

    @patch('setup_cron.subprocess.Popen')
    @patch('setup_cron.subprocess.check_output')
    @patch('setup_cron.os.path.abspath')
    def test_main(self, mock_abspath, mock_check_output, mock_popen):
        mock_abspath.return_value = '/path/to/main.py'

        mock_check_output.return_value = b'0 12 * * * python /path/to/main.py\n'

        mock_process = MagicMock()
        mock_popen.return_value = mock_process

        main()

        mock_check_output.assert_called_once_with(['crontab', '-l'], stderr=subprocess.STDOUT)
        mock_popen.assert_called_once_with(['crontab', '-'], stdin=subprocess.PIPE)

        args, _ = mock_process.communicate.call_args
        new_cron = args[0].decode()

        self.assertNotIn('0 12 * * * python /path/to/main.py', new_cron)
        self.assertIn('python /path/to/main.py', new_cron)

if __name__ == '__main__':
    unittest.main()