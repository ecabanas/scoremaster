import apiClient from '../services/apiClient';
import { ApiResponse, SubmitResponseData } from '../types/apiResponses';

export const useSubmitForm = () => {
  const sendDataToBackend = async (data: any) => {
    try {
      const response = await apiClient.post<ApiResponse<SubmitResponseData>>(
        `${import.meta.env.VITE_API_BASE_URL}/participants`,
        data
      );
      return response.data;
    } catch (error: any) {
      throw error.response?.data?.message || 'An unexpected error occurred';
    }
  };

  return { sendDataToBackend };
};