import { usePopupForm } from '../../hooks/usePopupForm';
import FormContent from './FormContent';

function PopupForm() {
  const { isVisible, setIsVisible } = usePopupForm();

  const handleFormSubmit = (data: { name: string; email: string }) => {
    console.log('Form Data:', data);
    setIsVisible(false); // Close the pop-up after submission
  };

  return (
    <>
      {/* Button to open the form */}
      <button className="btn btn-primary" onClick={() => setIsVisible(true)}>
        Open Form
      </button>

      {/* Pop-up form */}
      {isVisible && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="card w-96 bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title text-center">Enter Your Details</h2>
              <FormContent
                onSubmit={handleFormSubmit} // Pass the submit handler
                onCancel={() => setIsVisible(false)} // Close the pop-up on cancel
              />
            </div>
          </div>
        </div>
      )}
    </>
  );
}

export default PopupForm;