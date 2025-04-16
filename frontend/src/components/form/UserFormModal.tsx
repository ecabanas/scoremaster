import { useRef, useEffect } from 'react';
import { useModalState } from '../../hooks/useFormState';
import FormContent from './FormContent';
import { ParticipantFormData } from '../../types/ParticipantFormData';
import { FormButton } from './FormButton';

const UserFormModal = () => {
  const { isVisible, open, close } = useModalState();
  const modalRef = useRef<HTMLDivElement>(null);
  const openButtonRef = useRef<HTMLButtonElement>(null);

  // Handle form submission
  const handleFormSubmit = (data: ParticipantFormData) => {
    console.log('Form Data:', data);
    close(); // Close the modal after submission
  };

  // Close when clicking outside of modal
  const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
    if (modalRef.current && !modalRef.current.contains(e.target as Node)) {
      close();
    }
  };

  // Manage focus for accessibility
  useEffect(() => {
    if (isVisible) {
      // Focus the first input when modal opens
      setTimeout(() => {
        modalRef.current?.querySelector('input')?.focus();
      }, 100);
    } else {
      // Return focus to the button when modal closes
      openButtonRef.current?.focus();
    }
  }, [isVisible]);

  return (
    <>
      {/* Button to open the modal */}
      <FormButton
        variant="primary"
        onClick={open}
        ref={openButtonRef}
        aria-haspopup="dialog"
        aria-expanded={isVisible}
      >
        Enter Your Details
      </FormButton>

      {/* Modal overlay */}
      {isVisible && (
        <div
          className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
          onClick={handleOverlayClick}
          role="dialog"
          aria-modal="true"
          aria-labelledby="popup-title"
        >
          {/* Modal content */}
          <div
            className="card w-96 bg-base-100 shadow-xl"
            ref={modalRef}
          >
            <div className="card-body">
              <h2 id="popup-title" className="card-title text-center">
                Enter Your Details
              </h2>
              <FormContent
                onSubmit={handleFormSubmit}
                onCancel={close}
              />
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default UserFormModal;