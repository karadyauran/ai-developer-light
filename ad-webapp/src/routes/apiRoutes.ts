const API_BASE_URL = process.env.REACT_APP_API_BASE_URL;

const apiRoutes = {
	githubAuth: `${API_BASE_URL}/api/v1/oauth/github/authenticate`,
};

export default apiRoutes;