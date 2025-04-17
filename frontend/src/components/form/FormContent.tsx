import { useForm, SubmitHandler } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { formValidationSchema } from '../../validation/formValidation';
import { ParticipantFormData, FormContentProps } from '../../types/index';
import { FormInput } from './FormInput';
import { FormButton } from './FormButton';
import { useFormSubmit } from '../../hooks/useFormSubmit';



function FormContent({ onSubmit, onCancel }: FormContentProps) {
  const { submitForm, isSubmitting, error } = useFormSubmit();

  // Initialize React Hook Form with Yup validation
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<ParticipantFormData>({
    resolver: yupResolver(formValidationSchema),
  });

  // Submit handler
  const handleFormSubmit: SubmitHandler<ParticipantFormData> = async (data) => {
    const result = await submitForm(data);
    if (result) {
      onSubmit(data);
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