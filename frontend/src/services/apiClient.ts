import axios from 'axios';
// Create an Axios instance
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Create a new participant
export const createParticipant = async (participant: { name: string; email: string }) => {
    const response = await apiClient.post('/participants', participant);
    return response.data;
  };

export default apiClient;