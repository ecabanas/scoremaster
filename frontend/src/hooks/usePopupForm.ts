import { useState } from 'react';

export function usePopupForm() {
  const [isVisible, setIsVisible] = useState(false);

  return {
    isVisible,
    setIsVisible,
  };
}