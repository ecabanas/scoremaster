import { ParticipantFormData } from "./participant";

export interface FormContentProps {
    onSubmit: (data: ParticipantFormData) => void;
    onCancel: () => void;
}