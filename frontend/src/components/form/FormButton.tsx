import React, { forwardRef } from 'react';

type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'success';
type ButtonSize = 'sm' | 'md' | 'lg';
type ButtonType = 'button' | 'submit' | 'reset';

interface FormButtonProps {
    type?: ButtonType;
    variant: ButtonVariant;
    size?: ButtonSize;
    onClick?: () => void;
    disabled?: boolean;
    isLoading?: boolean;
    children: React.ReactNode;
    className?: string;
}

export const FormButton = forwardRef<HTMLButtonElement, FormButtonProps>(({
    type = 'button',
    variant = 'primary',
    size = 'md',
    onClick,
    disabled = false,
    isLoading = false,
    children,
    className = '',
}, ref) => {
    const variantClasses = {
        primary: 'btn-primary',
        secondary: 'btn-secondary',
        danger: 'btn-error',
        success: 'btn-success',
    };

    const sizeClasses = {
        sm: 'btn-sm',
        md: '',
        lg: 'btn-lg',
    };

    return (
        <button
            type={type}
            className={`btn ${variantClasses[variant]} ${sizeClasses[size]} ${className}`}
            onClick={onClick}
            disabled={disabled || isLoading}
            aria-busy={isLoading}
            ref={ref}
        >
            {isLoading && <span className="loading loading-spinner loading-xs mr-2"></span>}
            {children}
        </button>
    );
});

FormButton.displayName = 'FormButton';