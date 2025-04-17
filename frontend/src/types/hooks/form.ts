export type FormSubmitState = {
    isSubmitting: boolean;
    error: string | null;
    isSuccess: boolean;
};

export type FormSubmitAction =
    | { type: 'SUBMIT_START' }
    | { type: 'SUBMIT_SUCCESS' }
    | { type: 'SUBMIT_ERROR'; payload: string }
    | { type: 'RESET' };