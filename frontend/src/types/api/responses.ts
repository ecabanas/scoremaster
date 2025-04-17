export interface ApiResponse<T> {
    success: boolean;
    message: string;
    data?: T;
}

export interface SubmitResponseData {
    id: string;
    name: string;
    email: string;
}