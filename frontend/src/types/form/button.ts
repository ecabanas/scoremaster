import { ReactNode } from 'react';

export type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'success';
export type ButtonSize = 'sm' | 'md' | 'lg';
export type ButtonType = 'button' | 'submit' | 'reset';

export interface FormButtonProps {
    type?: ButtonType;
    variant: ButtonVariant;
    size?: ButtonSize;
    onClick?: () => void;
    disabled?: boolean;
    isLoading?: boolean;
    children: ReactNode;
    className?: string;
}