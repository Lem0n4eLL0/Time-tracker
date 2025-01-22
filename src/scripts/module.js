function openPopupHandler(popup) {
    openPopup(popup);
    document.addEventListener('keydown', closingOnEsc);
    popup.addEventListener('click', closingOnOverlay);
  }
  
  function closePopupHandler(popup) {
    removeClosingOnEsc();
    removeClosingOnOverlay(popup);
    closePopup(popup);
  }
  
  function removeClosingOnEsc() {
    document.removeEventListener('keydown', closingOnEsc);
  }
  
  function removeClosingOnOverlay(popup) {
    popup.removeEventListener('click', closingOnOverlay);
  }
  
  function closingOnEsc(evt) {
    if (evt.key === 'Escape') {
      const openPopupElement = document.querySelector('.popup_is-opened');
      if (openPopupElement) {
        closePopupHandler(openPopupElement);
      }
    }
  }
  
  function closingOnOverlay(evt) {
    if (evt.target === evt.currentTarget) {
      closePopupHandler(evt.currentTarget);
    }
  }
  
  function openPopup(popup) {
    popup.classList.add('popup_is-opened');
  }
  
  function closePopup(popup) {
    popup.classList.remove('popup_is-opened');
  }

  export {openPopupHandler as open, closePopupHandler as close}