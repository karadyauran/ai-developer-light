// src/components/Profile.tsx
import React, { useEffect, useState } from 'react';

const Profile: React.FC = () => {
	const [userData, setUserData] = useState<any>(null);

	useEffect(() => {
		const storedUserData = localStorage.getItem('userData');
		if (storedUserData) {
			setUserData(JSON.parse(storedUserData));
		}
	}, []);

	if (!userData) {
		return <div>Failed to load user data. Try logging in again.</div>;
	}

	return (
		<div>
			<h1>User profile</h1>
			<img
					src={userData.avatar.value}
					alt="GitHub Avatar"
					style={{ width: '100px', height: '100px', borderRadius: '50%' }}
				/>
			<p><strong>Name:</strong> {userData.username}</p>
			<p><strong>Email:</strong> {userData.email.value}</p>
		</div>
	);
};

export default Profile;