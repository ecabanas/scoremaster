import apiClient from '../services/apiClient';
import { ApiResponse, SubmitResponseData } from '../types/apiResponses';
import { ParticipantFormData } from '../types/ParticipantFormData';

export const useSubmitForm = () => {
  const sendDataToBackend = async (data: ParticipantFormData) => {
    try {
      const response = await apiClient.post<ApiResponse<SubmitResponseData>>(
        '/participants',
        data
      );
      return response.data;
    } catch (error: any) {
      throw error.response?.data?.message || 'An unexpected error occurred';
    }
  };

  return { sendDataToBackend };
};