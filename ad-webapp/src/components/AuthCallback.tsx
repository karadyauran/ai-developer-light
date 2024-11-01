// src/components/AuthCallback.tsx
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import apiRoutes from '../routes/apiRoutes';

const AuthCallback: React.FC = () => {
	const navigate = useNavigate();
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		const urlParams = new URLSearchParams(window.location.search);
		const code = urlParams.get('code');

		if (code) {
			axios.post(apiRoutes.githubAuth, { code })
				.then(response => {
					localStorage.setItem('userData', JSON.stringify(response.data));
					setLoading(false);
					navigate('/profile');
				})
				.catch(error => {
					console.error('Authorization error:', error);
					setError('Authorization error. Try again.');
					setLoading(false);
				});
		} else {
			setError('Authorization code is missing from URL');
			setLoading(false);
		}
	}, [navigate]);

	if (loading) {
		return <div>Authorization via GitHub...</div>;
	}

	if (error) {
		return <div>{error}</div>;
	}

	return null;
};

export default AuthCallback;