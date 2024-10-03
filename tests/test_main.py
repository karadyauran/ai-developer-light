import unittest
from unittest.mock import patch, MagicMock
from main import main

class TestMain(unittest.TestCase):

    @patch('main.subprocess.run')
    @patch('main.open', create=True)
    @patch('main.AppGenerator')
    def test_main(self, mock_app_generator, mock_open, mock_subprocess_run):
        mock_generator = MagicMock()
        mock_app_generator.return_value = mock_generator

        mock_generator.languages = ['Python']
        mock_generator.generate_app_idea.return_value = 'TestApp\nThis is a test application.'
        mock_generator.generate_unique_app_name.return_value = 'TestApp_12345678'
        mock_generator.generate_file_structure.return_value = ['file1.py', 'file2.py']
        mock_generator.generate_code.return_value = 'print("Hello, World!")'
        mock_generator.create_commit_message.return_value = 'Add file1.py'

        with patch('main.random.choice', return_value='Python'), \
             patch('main.random.randint', side_effect=[3, 2]):

            main()

            mock_generator.generate_app_idea.assert_called_once_with('Python', 3)
            mock_generator.generate_unique_app_name.assert_called_once_with('TestApp')
            mock_generator.create_repository.assert_called_once_with('TestApp_12345678')
            mock_generator.generate_file_structure.assert_called_once_with('Python', 2)
            self.assertEqual(mock_generator.generate_code.call_count, 2)
            self.assertEqual(mock_generator.create_commit_message.call_count, 2)

            self.assertEqual(mock_subprocess_run.call_count, 4)

            self.assertEqual(mock_open.call_count, 2)

if __name__ == '__main__':
    unittest.main()