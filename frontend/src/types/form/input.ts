import { FieldError, UseFormRegister } from 'react-hook-form';

export type InputType = 'text' | 'email' | 'password' | 'number';

export interface FormInputProps {
    id: string;
    label: string;
    type: InputType;
    register: UseFormRegister<any>;
    error?: FieldError;
    placeholder?: string;
    required?: boolean;
}