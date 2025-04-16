import { FieldError, UseFormRegister } from 'react-hook-form';

type InputType = 'text' | 'email' | 'password' | 'number';

export interface FormInputProps {
    id: string;
    label: string;
    type: InputType;
    register: UseFormRegister<any>;
    error?: FieldError;
    placeholder?: string;
    required?: boolean;
}

export const FormInput = ({
    id,
    label,
    type,
    register,
    error,
    placeholder,
    required = false
}: FormInputProps) => {
    return (
        <div className="form-control mb-4">
            <label htmlFor={id} className="label">
                <span className="label-text">
                    {label} {required && <span className="text-error">*</span>}
                </span>
            </label>
            <input
                type={type}
                id={id}
                placeholder={placeholder}
                className={`input input-bordered ${error ? 'input-error' : ''}`}
                {...register(id)}
                aria-invalid={error ? 'true' : 'false'}
                aria-describedby={error ? `${id}-error` : undefined}
            />
            {error && (
                <span id={`${id}-error`} className="text-error text-sm mt-1">
                    {error.message}
                </span>
            )}
        </div>
    );
};