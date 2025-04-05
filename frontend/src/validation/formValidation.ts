import * as yup from 'yup';

// Define the validation schema using Yup
export const formValidationSchema = yup.object().shape({
  name: yup.string().required('Name is required').min(2, 'Name must be at least 2 characters'),
  email: yup.string().required('Email is required').email('Invalid email format'),
});