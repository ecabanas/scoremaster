import { useState } from 'react';
import apiClient from '../services/apiClient';
import { ApiResponse, SubmitResponseData } from '../types/api/responses';
import { ParticipantFormData } from '../types/index';

export function useFormSubmit() {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isSuccess, setIsSuccess] = useState(false);

  const submitForm = async (data: ParticipantFormData): Promise<ApiResponse<SubmitResponseData> | null> => {
    setIsSubmitting(true);
    setError(null);
    setIsSuccess(false);

    try {
      const response = await apiClient.post<ApiResponse<SubmitResponseData>>(
        '/participants',
        data
      );
      setIsSuccess(true);
      return response.data;
    } catch (error: any) {
      const errorMessage = error.response?.data?.message || 'An unexpected error occurred';
      setError(errorMessage);
      return null;
    } finally {
      setIsSubmitting(false);
    }
  };

  const resetState = () => {
    setError(null);
    setIsSuccess(false);
    setIsSubmitting(false);
  };

  return {
    submitForm,
    isSubmitting,
    error,
    isSuccess,
    resetState
  };
}