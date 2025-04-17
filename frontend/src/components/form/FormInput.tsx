import { FormInputProps } from '../../types/index';

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