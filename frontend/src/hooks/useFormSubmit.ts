import { useReducer } from 'react';
import { useNavigate } from 'react-router-dom';
import apiClient from '../services/apiClient';
import {
  ApiResponse,
  SubmitResponseData,
  ParticipantFormData,
  FormSubmitState,
  FormSubmitAction
} from '../types';


/**
 * Custom hook for handling form submission state and API interactions
 * @param redirectPath Optional path to redirect to after successful submission
 * @returns Object with form submission state and functions
 */
export function useFormSubmit(redirectPath?: string) {
  const navigate = useNavigate();
  const initialState: FormSubmitState = {
    isSubmitting: false,
    error: null,
    isSuccess: false
  };

  const reducer = (state: FormSubmitState, action: FormSubmitAction): FormSubmitState => {
    switch (action.type) {
      case 'SUBMIT_START':
        return { ...state, isSubmitting: true, error: null, isSuccess: false };
      case 'SUBMIT_SUCCESS':
        return { ...state, isSubmitting: false, isSuccess: true };
      case 'SUBMIT_ERROR':
        return { ...state, isSubmitting: false, error: action.payload };
      case 'RESET':
        return initialState;
      default:
        return state;
    }
  };

  const [state, dispatch] = useReducer(reducer, initialState);

  const submitForm = async (data: ParticipantFormData): Promise<ApiResponse<SubmitResponseData> | null> => {
    dispatch({ type: 'SUBMIT_START' });

    try {
      const response = await apiClient.post<ApiResponse<SubmitResponseData>>(
        '/participants',
        data
      );
      dispatch({ type: 'SUBMIT_SUCCESS' });

      // If redirect path is provided, navigate to it
      if (redirectPath) {
        navigate(redirectPath);
      }
      return response.data;
    } catch (error: unknown) {
      const errorMessage =
        error instanceof Error ? error.message :
          typeof error === 'object' && error && 'response' in error ?
            (error.response as any)?.data?.message : 'An unexpected error occurred';

      dispatch({ type: 'SUBMIT_ERROR', payload: errorMessage });
      return null;
    }
  };

  const resetState = () => {
    dispatch({ type: 'RESET' });
  };

  return {
    ...state,
    submitForm,
    resetState
  };
}