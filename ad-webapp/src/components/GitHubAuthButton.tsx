import React from 'react';

const GitHubAuthButton: React.FC = () => {
	const handleLogin = () => {
		const githubAuthUrl = `https://github.com/login/oauth/authorize?client_id=${process.env.REACT_APP_GITHUB_CLIENT_ID}&redirect_uri=${process.env.REACT_APP_GITHUB_REDIRECT_URI}&scope=read:user`;
		window.location.href = githubAuthUrl;
	};

	return (
		<button onClick={handleLogin}>Login with GitHub</button>
	);
};

export default GitHubAuthButton;