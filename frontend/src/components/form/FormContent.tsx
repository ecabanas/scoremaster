import { useForm, SubmitHandler } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { formValidationSchema } from '../../validation/formValidation';
import { ParticipantFormData, FormContentProps } from '../../types/index';
import { FormInput } from './FormInput';
import { FormButton } from './FormButton';
import { useFormSubmit } from '../../hooks/useFormSubmit';



function FormContent({ onSubmit, onCancel }: FormContentProps) {
  const { submitForm, isSubmitting, error } = useFormSubmit('/survey/1');
  // Initialize React Hook Form with Yup validation
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError
  } = useForm<ParticipantFormData>({
    resolver: yupResolver(formValidationSchema),
    mode: 'onBlur'
  });

  // Submit handler
  const handleFormSubmit: SubmitHandler<ParticipantFormData> = async (data) => {
    try {
      const result = await submitForm(data);
      if (result) {
        onSubmit(data);
      }
    } catch (error: any) {
      // If the backend returns a duplicate email error that wasn't caught by validation
      if (error?.response?.data?.message?.includes('already exists')) {
        setError('email', {
          type: 'manual',
          message: 'This email is already registered'
        });
      }
    }
  };

  return (
    <form onSubmit={handleSubmit(handleFormSubmit)}>
      {error && (
        <div className="alert alert-error mb-4">
          <p>{error}</p>
        </div>
      )}

      <FormInput
        id="name"
        label="Name"
        type="text"
        register={register}
        error={errors.name}
        placeholder="Enter your name"
        required
      />

      <FormInput
        id="email"
        label="Email"
        type="email"
        register={register}
        error={errors.email}
        placeholder="Enter your email"
        required
      />

      {/* Action Buttons */}
      <div className="card-actions justify-end mt-6">
        <FormButton
          variant="secondary"
          onClick={onCancel}
          disabled={isSubmitting}
        >
          Cancel
        </FormButton>

        <FormButton
          type="submit"
          variant="primary"
          isLoading={isSubmitting}
        >
          Submit
        </FormButton>
      </div>
    </form>
  );
}

export default FormContent;