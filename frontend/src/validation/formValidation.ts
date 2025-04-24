import * as yup from 'yup';
import apiClient from '../services/apiClient';

// Helper function to check if email exists
const checkEmailExists = async (email: string) => {
  try {
    const response = await apiClient.post('/participants/check-email', { email });
    return response.data.exists;
  } catch (error) {
    console.error('Error checking email:', error);
    return false; // Assume email doesn't exist if API fails
  }
};

// Define the validation schema using Yup
export const formValidationSchema = yup.object().shape({
  name: yup.string()
    .required('Name is required')
    .min(2, 'Name must be at least 2 characters'),

  email: yup.string()
    .required('Email is required')
    .email('Invalid email format')
    .test(
      'email-exists',
      'This email is already registered',
      async (email) => {
        if (!email) return true; // Skip validation if email is empty
        const exists = await checkEmailExists(email);
        return !exists; // Return false if email exists (validation fails)
      }
    )
});

// For non-async validation only
export const basicFormValidationSchema = yup.object().shape({
  name: yup.string()
    .required('Name is required')
    .min(2, 'Name must be at least 2 characters'),

  email: yup.string()
    .required('Email is required')
    .email('Invalid email format')
});