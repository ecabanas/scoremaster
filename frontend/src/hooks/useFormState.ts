import { useState } from 'react';

export function useModalState() {
  const [isVisible, setIsVisible] = useState(false);

  return {
    isVisible,
    setIsVisible,
    open: () => setIsVisible(true),
    close: () => setIsVisible(false),
    toggle: () => setIsVisible(prev => !prev)
  };
}