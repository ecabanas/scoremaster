import { useForm, SubmitHandler } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { formValidationSchema } from '../../validation/formValidation';
import { ParticipantFormData } from '../../types/ParticipantFormData';
import { createParticipant } from '../../services/apiClient';

interface FormContentProps {
  onSubmit: (data: ParticipantFormData) => void;
  onCancel: () => void;
}

function FormContent({ onSubmit, onCancel }: FormContentProps) {
  // Initialize React Hook Form with Yup validation
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<ParticipantFormData>({
    resolver: yupResolver(formValidationSchema), // Use the imported validation schema
  });

  // Submit handler
  const handleFormSubmit: SubmitHandler<ParticipantFormData> = async (data) => {
    try {
      const newParticipant = await createParticipant(data);
      onSubmit(newParticipant); // Pass validated data to the parent component
    } catch (error: any) {
      const errorMessage = error.response?.data?.message || 'An unexpected error occurred';
      alert(errorMessage); // Display error message
    }
  };

  return (
    <form onSubmit={handleSubmit(handleFormSubmit)}>
      {/* Name Input */}
      <div className="form-control mb-4">
        <label htmlFor="name" className="label">
          <span className="label-text">Name</span>
        </label>
        <input
          type="text"
          id="name"
          className={`input input-bordered ${errors.name ? 'input-error' : ''}`}
          {...register('name')} // Register the input with React Hook Form
        />
        {errors.name && <span className="text-error text-sm">{errors.name.message}</span>}
      </div>

      {/* Email Input */}
      <div className="form-control mb-4">
        <label htmlFor="email" className="label">
          <span className="label-text">Email</span>
        </label>
        <input
          type="email"
          id="email"
          className={`input input-bordered ${errors.email ? 'input-error' : ''}`}
          {...register('email')} // Register the input with React Hook Form
        />
        {errors.email && <span className="text-error text-sm">{errors.email.message}</span>}
      </div>

      {/* Action Buttons */}
      <div className="card-actions justify-end">
        <button type="button" className="btn btn-secondary" onClick={onCancel}>
          Cancel
        </button>
        <button type="submit" className="btn btn-primary">
          Submit
        </button>
      </div>
    </form>
  );
}

export default FormContent;